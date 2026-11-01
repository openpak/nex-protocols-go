package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

func testStreamIn(data []byte, header bool) *nex.ByteStreamIn {
	versions := nex.NewLibraryVersions()
	versions.Main = nex.NewLibraryVersion(3, 0, 0)
	settings := nex.NewByteStreamSettings()
	settings.UseStructureHeader = header

	return nex.NewByteStreamIn(data, versions, settings)
}

func testStreamOut(header bool) *nex.ByteStreamOut {
	versions := nex.NewLibraryVersions()
	versions.Main = nex.NewLibraryVersion(3, 0, 0)
	settings := nex.NewByteStreamSettings()
	settings.UseStructureHeader = header

	return nex.NewByteStreamOut(versions, settings)
}

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

// roundTrip decodes want with decode, checks everything was consumed, re-encodes with encode and compares bytes.
func roundTrip(t *testing.T, name string, header bool, want []byte, decode func(*nex.ByteStreamIn) error, encode func(*nex.ByteStreamOut)) {
	t.Helper()

	in := testStreamIn(want, header)
	if err := decode(in); err != nil {
		t.Fatalf("%s (header=%v): decode: %s", name, header, err)
	}

	if in.Remaining() != 0 {
		t.Fatalf("%s (header=%v): %d bytes left over", name, header, in.Remaining())
	}

	out := testStreamOut(header)
	encode(out)

	if !bytes.Equal(out.Bytes(), want) {
		t.Fatalf("%s (header=%v): encoded bytes differ\n got %x\nwant %x", name, header, out.Bytes(), want)
	}
}

func TestCreateSessionRequest_NoHeader(t *testing.T) {
	want := CreateSessionRequest{GameSession: func() game_session_types.GameSession {
		v := game_session_types.NewGameSession()
		v.TypeID = types.NewUInt32(2023)
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(5050)
			v.Value = types.NewUInt32(6059)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(8077)
			v.Value = types.NewUInt32(9086)
			return v
		}()}
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got CreateSessionRequest

	roundTrip(t, "CreateSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSession.Equals(want.GameSession)) {
		t.Fatalf("CreateSessionRequest.GameSession differs")
	}
}

func TestCreateSessionRequest_Header(t *testing.T) {
	want := CreateSessionRequest{GameSession: func() game_session_types.GameSession {
		v := game_session_types.NewGameSession()
		v.TypeID = types.NewUInt32(2023)
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(5050)
			v.Value = types.NewUInt32(6059)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(8077)
			v.Value = types.NewUInt32(9086)
			return v
		}()}
		return v
	}()}

	data := []byte{0x00, 0x22, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got CreateSessionRequest

	roundTrip(t, "CreateSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSession.Equals(want.GameSession)) {
		t.Fatalf("CreateSessionRequest.GameSession differs")
	}
}

func TestCreateSessionResponse_NoHeader(t *testing.T) {
	want := CreateSessionResponse{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got CreateSessionResponse

	roundTrip(t, "CreateSessionResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("CreateSessionResponse.GameSessionKey differs")
	}
}

func TestCreateSessionResponse_Header(t *testing.T) {
	want := CreateSessionResponse{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got CreateSessionResponse

	roundTrip(t, "CreateSessionResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("CreateSessionResponse.GameSessionKey differs")
	}
}

func TestUpdateSessionRequest_NoHeader(t *testing.T) {
	want := UpdateSessionRequest{GameSessionUpdate: func() game_session_types.GameSessionUpdate {
		v := game_session_types.NewGameSessionUpdate()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.NewUInt32(8077)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.NewUInt32(11104)
			return v
		}()}
		return v
	}()}

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00}

	var got UpdateSessionRequest

	roundTrip(t, "UpdateSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionUpdate.Equals(want.GameSessionUpdate)) {
		t.Fatalf("UpdateSessionRequest.GameSessionUpdate differs")
	}
}

func TestUpdateSessionRequest_Header(t *testing.T) {
	want := UpdateSessionRequest{GameSessionUpdate: func() game_session_types.GameSessionUpdate {
		v := game_session_types.NewGameSessionUpdate()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.NewUInt32(8077)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.NewUInt32(11104)
			return v
		}()}
		return v
	}()}

	data := []byte{0x00, 0x2b, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00}

	var got UpdateSessionRequest

	roundTrip(t, "UpdateSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionUpdate.Equals(want.GameSessionUpdate)) {
		t.Fatalf("UpdateSessionRequest.GameSessionUpdate differs")
	}
}

func TestDeleteSessionRequest_NoHeader(t *testing.T) {
	want := DeleteSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got DeleteSessionRequest

	roundTrip(t, "DeleteSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("DeleteSessionRequest.GameSessionKey differs")
	}
}

func TestDeleteSessionRequest_Header(t *testing.T) {
	want := DeleteSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got DeleteSessionRequest

	roundTrip(t, "DeleteSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("DeleteSessionRequest.GameSessionKey differs")
	}
}

func TestMigrateSessionRequest_NoHeader(t *testing.T) {
	want := MigrateSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got MigrateSessionRequest

	roundTrip(t, "MigrateSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("MigrateSessionRequest.GameSessionKey differs")
	}
}

func TestMigrateSessionRequest_Header(t *testing.T) {
	want := MigrateSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got MigrateSessionRequest

	roundTrip(t, "MigrateSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("MigrateSessionRequest.GameSessionKey differs")
	}
}

func TestMigrateSessionResponse_NoHeader(t *testing.T) {
	want := MigrateSessionResponse{GameSessionKeyMigrated: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got MigrateSessionResponse

	roundTrip(t, "MigrateSessionResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKeyMigrated.Equals(want.GameSessionKeyMigrated)) {
		t.Fatalf("MigrateSessionResponse.GameSessionKeyMigrated differs")
	}
}

func TestMigrateSessionResponse_Header(t *testing.T) {
	want := MigrateSessionResponse{GameSessionKeyMigrated: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got MigrateSessionResponse

	roundTrip(t, "MigrateSessionResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKeyMigrated.Equals(want.GameSessionKeyMigrated)) {
		t.Fatalf("MigrateSessionResponse.GameSessionKeyMigrated differs")
	}
}

func TestLeaveSessionRequest_NoHeader(t *testing.T) {
	want := LeaveSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got LeaveSessionRequest

	roundTrip(t, "LeaveSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("LeaveSessionRequest.GameSessionKey differs")
	}
}

func TestLeaveSessionRequest_Header(t *testing.T) {
	want := LeaveSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got LeaveSessionRequest

	roundTrip(t, "LeaveSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("LeaveSessionRequest.GameSessionKey differs")
	}
}

func TestGetSessionRequest_NoHeader(t *testing.T) {
	want := GetSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetSessionRequest

	roundTrip(t, "GetSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("GetSessionRequest.GameSessionKey differs")
	}
}

func TestGetSessionRequest_Header(t *testing.T) {
	want := GetSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetSessionRequest

	roundTrip(t, "GetSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("GetSessionRequest.GameSessionKey differs")
	}
}

func TestGetSessionResponse_NoHeader(t *testing.T) {
	want := GetSessionResponse{SearchResult: func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.HostPID = types.NewUInt32(5050)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.7"), types.NewStationURL("prudp:/address=10.0.0.8")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(11104)
			v.Value = types.NewUInt32(12113)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(14131)
			v.Value = types.NewUInt32(15140)
			return v
		}()}
		return v
	}()}

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x37, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x38, 0x00, 0x02, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00}

	var got GetSessionResponse

	roundTrip(t, "GetSessionResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResult.Equals(want.SearchResult)) {
		t.Fatalf("GetSessionResponse.SearchResult differs")
	}
}

func TestGetSessionResponse_Header(t *testing.T) {
	want := GetSessionResponse{SearchResult: func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.HostPID = types.NewUInt32(5050)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.7"), types.NewStationURL("prudp:/address=10.0.0.8")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(11104)
			v.Value = types.NewUInt32(12113)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(14131)
			v.Value = types.NewUInt32(15140)
			return v
		}()}
		return v
	}()}

	data := []byte{0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x37, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x38, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00}

	var got GetSessionResponse

	roundTrip(t, "GetSessionResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResult.Equals(want.SearchResult)) {
		t.Fatalf("GetSessionResponse.SearchResult differs")
	}
}

func TestSearchSessionsRequest_NoHeader(t *testing.T) {
	want := SearchSessionsRequest{GameSessionQuery: func() game_session_types.GameSessionQuery {
		v := game_session_types.NewGameSessionQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.Parameters = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(6059)
			v.Value = types.NewUInt32(7068)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.NewUInt32(10095)
			return v
		}()}
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00}

	var got SearchSessionsRequest

	roundTrip(t, "SearchSessionsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionQuery.Equals(want.GameSessionQuery)) {
		t.Fatalf("SearchSessionsRequest.GameSessionQuery differs")
	}
}

func TestSearchSessionsRequest_Header(t *testing.T) {
	want := SearchSessionsRequest{GameSessionQuery: func() game_session_types.GameSessionQuery {
		v := game_session_types.NewGameSessionQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.Parameters = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(6059)
			v.Value = types.NewUInt32(7068)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.NewUInt32(10095)
			return v
		}()}
		return v
	}()}

	data := []byte{0x00, 0x26, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00}

	var got SearchSessionsRequest

	roundTrip(t, "SearchSessionsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionQuery.Equals(want.GameSessionQuery)) {
		t.Fatalf("SearchSessionsRequest.GameSessionQuery differs")
	}
}

func TestSearchSessionsResponse_NoHeader(t *testing.T) {
	want := SearchSessionsResponse{SearchResults: types.List[game_session_types.GameSessionSearchResult]{func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.HostPID = types.NewUInt32(6059)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.8"), types.NewStationURL("prudp:/address=10.0.0.9")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.NewUInt32(13122)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(15140)
			v.Value = types.NewUInt32(16149)
			return v
		}()}
		return v
	}(), func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(19176)
			v.SessionID = types.NewUInt32(20185)
			return v
		}()
		v.HostPID = types.NewUInt32(21194)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.23"), types.NewStationURL("prudp:/address=10.0.0.24")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(27248)
			v.Value = types.NewUInt32(28257)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(30275)
			v.Value = types.NewUInt32(31284)
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x38, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x39, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x33, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x34, 0x00, 0x02, 0x00, 0x00, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x61, 0x6e, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00}

	var got SearchSessionsResponse

	roundTrip(t, "SearchSessionsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchSessionsResponse.SearchResults differs")
	}
}

func TestSearchSessionsResponse_Header(t *testing.T) {
	want := SearchSessionsResponse{SearchResults: types.List[game_session_types.GameSessionSearchResult]{func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.HostPID = types.NewUInt32(6059)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.8"), types.NewStationURL("prudp:/address=10.0.0.9")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.NewUInt32(13122)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(15140)
			v.Value = types.NewUInt32(16149)
			return v
		}()}
		return v
	}(), func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(19176)
			v.SessionID = types.NewUInt32(20185)
			return v
		}()
		v.HostPID = types.NewUInt32(21194)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.23"), types.NewStationURL("prudp:/address=10.0.0.24")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(27248)
			v.Value = types.NewUInt32(28257)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(30275)
			v.Value = types.NewUInt32(31284)
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x38, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x39, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x33, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x34, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x61, 0x6e, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00}

	var got SearchSessionsResponse

	roundTrip(t, "SearchSessionsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchSessionsResponse.SearchResults differs")
	}
}

func TestAddParticipantsRequest_NoHeader(t *testing.T) {
	want := AddParticipantsRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}(), PublicParticipantIDs: types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}, PrivateParticipantIDs: types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got AddParticipantsRequest

	roundTrip(t, "AddParticipantsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("AddParticipantsRequest.GameSessionKey differs")
	}
	if !(got.PublicParticipantIDs.Equals(want.PublicParticipantIDs)) {
		t.Fatalf("AddParticipantsRequest.PublicParticipantIDs differs")
	}
	if !(got.PrivateParticipantIDs.Equals(want.PrivateParticipantIDs)) {
		t.Fatalf("AddParticipantsRequest.PrivateParticipantIDs differs")
	}
}

func TestAddParticipantsRequest_Header(t *testing.T) {
	want := AddParticipantsRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}(), PublicParticipantIDs: types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}, PrivateParticipantIDs: types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got AddParticipantsRequest

	roundTrip(t, "AddParticipantsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("AddParticipantsRequest.GameSessionKey differs")
	}
	if !(got.PublicParticipantIDs.Equals(want.PublicParticipantIDs)) {
		t.Fatalf("AddParticipantsRequest.PublicParticipantIDs differs")
	}
	if !(got.PrivateParticipantIDs.Equals(want.PrivateParticipantIDs)) {
		t.Fatalf("AddParticipantsRequest.PrivateParticipantIDs differs")
	}
}

func TestRemoveParticipantsRequest_NoHeader(t *testing.T) {
	want := RemoveParticipantsRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}(), ParticipantIDs: types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00}

	var got RemoveParticipantsRequest

	roundTrip(t, "RemoveParticipantsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("RemoveParticipantsRequest.GameSessionKey differs")
	}
	if !(got.ParticipantIDs.Equals(want.ParticipantIDs)) {
		t.Fatalf("RemoveParticipantsRequest.ParticipantIDs differs")
	}
}

func TestRemoveParticipantsRequest_Header(t *testing.T) {
	want := RemoveParticipantsRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}(), ParticipantIDs: types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00}

	var got RemoveParticipantsRequest

	roundTrip(t, "RemoveParticipantsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("RemoveParticipantsRequest.GameSessionKey differs")
	}
	if !(got.ParticipantIDs.Equals(want.ParticipantIDs)) {
		t.Fatalf("RemoveParticipantsRequest.ParticipantIDs differs")
	}
}

func TestGetParticipantCountRequest_NoHeader(t *testing.T) {
	want := GetParticipantCountRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetParticipantCountRequest

	roundTrip(t, "GetParticipantCountRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("GetParticipantCountRequest.GameSessionKey differs")
	}
}

func TestGetParticipantCountRequest_Header(t *testing.T) {
	want := GetParticipantCountRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetParticipantCountRequest

	roundTrip(t, "GetParticipantCountRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("GetParticipantCountRequest.GameSessionKey differs")
	}
}

func TestGetParticipantCountResponse_NoHeader(t *testing.T) {
	want := GetParticipantCountResponse{Count: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetParticipantCountResponse

	roundTrip(t, "GetParticipantCountResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("GetParticipantCountResponse.Count differs")
	}
}

func TestGetParticipantCountResponse_Header(t *testing.T) {
	want := GetParticipantCountResponse{Count: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetParticipantCountResponse

	roundTrip(t, "GetParticipantCountResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("GetParticipantCountResponse.Count differs")
	}
}

func TestGetParticipantsRequest_NoHeader(t *testing.T) {
	want := GetParticipantsRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}(), ResultRange: types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00}

	var got GetParticipantsRequest

	roundTrip(t, "GetParticipantsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("GetParticipantsRequest.GameSessionKey differs")
	}
	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetParticipantsRequest.ResultRange differs")
	}
}

func TestGetParticipantsRequest_Header(t *testing.T) {
	want := GetParticipantsRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}(), ResultRange: types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00}

	var got GetParticipantsRequest

	roundTrip(t, "GetParticipantsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("GetParticipantsRequest.GameSessionKey differs")
	}
	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetParticipantsRequest.ResultRange differs")
	}
}

func TestGetParticipantsResponse_NoHeader(t *testing.T) {
	want := GetParticipantsResponse{Participants: types.List[game_session_types.GameSessionParticipant]{func() game_session_types.GameSessionParticipant {
		v := game_session_types.NewGameSessionParticipant()
		v.PID = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.6"), types.NewStationURL("prudp:/address=10.0.0.7")}
		return v
	}(), func() game_session_types.GameSessionParticipant {
		v := game_session_types.NewGameSessionParticipant()
		v.PID = types.NewUInt32(9086)
		v.Name = types.NewString("s10")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.12"), types.NewStationURL("prudp:/address=10.0.0.13")}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x36, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x37, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x02, 0x00, 0x00, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x32, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x33, 0x00}

	var got GetParticipantsResponse

	roundTrip(t, "GetParticipantsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Participants.Equals(want.Participants)) {
		t.Fatalf("GetParticipantsResponse.Participants differs")
	}
}

func TestGetParticipantsResponse_Header(t *testing.T) {
	want := GetParticipantsResponse{Participants: types.List[game_session_types.GameSessionParticipant]{func() game_session_types.GameSessionParticipant {
		v := game_session_types.NewGameSessionParticipant()
		v.PID = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.6"), types.NewStationURL("prudp:/address=10.0.0.7")}
		return v
	}(), func() game_session_types.GameSessionParticipant {
		v := game_session_types.NewGameSessionParticipant()
		v.PID = types.NewUInt32(9086)
		v.Name = types.NewString("s10")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.12"), types.NewStationURL("prudp:/address=10.0.0.13")}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x36, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x37, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x02, 0x00, 0x00, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x32, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x33, 0x00}

	var got GetParticipantsResponse

	roundTrip(t, "GetParticipantsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Participants.Equals(want.Participants)) {
		t.Fatalf("GetParticipantsResponse.Participants differs")
	}
}

func TestSendInvitationRequest_NoHeader(t *testing.T) {
	want := SendInvitationRequest{Invitation: func() game_session_types.GameSessionInvitation {
		v := game_session_types.NewGameSessionInvitation()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.RecipientPIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		v.Message = types.NewString("s8")
		return v
	}()}

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00}

	var got SendInvitationRequest

	roundTrip(t, "SendInvitationRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Invitation.Equals(want.Invitation)) {
		t.Fatalf("SendInvitationRequest.Invitation differs")
	}
}

func TestSendInvitationRequest_Header(t *testing.T) {
	want := SendInvitationRequest{Invitation: func() game_session_types.GameSessionInvitation {
		v := game_session_types.NewGameSessionInvitation()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.RecipientPIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		v.Message = types.NewString("s8")
		return v
	}()}

	data := []byte{0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00}

	var got SendInvitationRequest

	roundTrip(t, "SendInvitationRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Invitation.Equals(want.Invitation)) {
		t.Fatalf("SendInvitationRequest.Invitation differs")
	}
}

func TestGetInvitationReceivedCountRequest_NoHeader(t *testing.T) {
	want := GetInvitationReceivedCountRequest{GameSessionTypeID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetInvitationReceivedCountRequest

	roundTrip(t, "GetInvitationReceivedCountRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("GetInvitationReceivedCountRequest.GameSessionTypeID differs")
	}
}

func TestGetInvitationReceivedCountRequest_Header(t *testing.T) {
	want := GetInvitationReceivedCountRequest{GameSessionTypeID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetInvitationReceivedCountRequest

	roundTrip(t, "GetInvitationReceivedCountRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("GetInvitationReceivedCountRequest.GameSessionTypeID differs")
	}
}

func TestGetInvitationReceivedCountResponse_NoHeader(t *testing.T) {
	want := GetInvitationReceivedCountResponse{Count: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetInvitationReceivedCountResponse

	roundTrip(t, "GetInvitationReceivedCountResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("GetInvitationReceivedCountResponse.Count differs")
	}
}

func TestGetInvitationReceivedCountResponse_Header(t *testing.T) {
	want := GetInvitationReceivedCountResponse{Count: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetInvitationReceivedCountResponse

	roundTrip(t, "GetInvitationReceivedCountResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("GetInvitationReceivedCountResponse.Count differs")
	}
}

func TestGetInvitationsReceivedRequest_NoHeader(t *testing.T) {
	want := GetInvitationsReceivedRequest{GameSessionTypeID: types.NewUInt32(1014), ResultRange: types.ResultRange{Offset: types.NewUInt32(6), Length: types.NewUInt32(16)}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00}

	var got GetInvitationsReceivedRequest

	roundTrip(t, "GetInvitationsReceivedRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("GetInvitationsReceivedRequest.GameSessionTypeID differs")
	}
	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetInvitationsReceivedRequest.ResultRange differs")
	}
}

func TestGetInvitationsReceivedRequest_Header(t *testing.T) {
	want := GetInvitationsReceivedRequest{GameSessionTypeID: types.NewUInt32(1014), ResultRange: types.ResultRange{Offset: types.NewUInt32(6), Length: types.NewUInt32(16)}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00}

	var got GetInvitationsReceivedRequest

	roundTrip(t, "GetInvitationsReceivedRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("GetInvitationsReceivedRequest.GameSessionTypeID differs")
	}
	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetInvitationsReceivedRequest.ResultRange differs")
	}
}

func TestGetInvitationsReceivedResponse_NoHeader(t *testing.T) {
	want := GetInvitationsReceivedResponse{Invitations: types.List[game_session_types.GameSessionInvitationReceived]{func() game_session_types.GameSessionInvitationReceived {
		v := game_session_types.NewGameSessionInvitationReceived()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.SenderPID = types.NewUInt32(6059)
		v.Message = types.NewString("s7")
		v.CreationTime = types.NewDateTime(8000101)
		return v
	}(), func() game_session_types.GameSessionInvitationReceived {
		v := game_session_types.NewGameSessionInvitationReceived()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(11104)
			v.SessionID = types.NewUInt32(12113)
			return v
		}()
		v.SenderPID = types.NewUInt32(13122)
		v.Message = types.NewString("s14")
		v.CreationTime = types.NewDateTime(15000122)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x34, 0x00, 0x3a, 0xe2, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GetInvitationsReceivedResponse

	roundTrip(t, "GetInvitationsReceivedResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Invitations.Equals(want.Invitations)) {
		t.Fatalf("GetInvitationsReceivedResponse.Invitations differs")
	}
}

func TestGetInvitationsReceivedResponse_Header(t *testing.T) {
	want := GetInvitationsReceivedResponse{Invitations: types.List[game_session_types.GameSessionInvitationReceived]{func() game_session_types.GameSessionInvitationReceived {
		v := game_session_types.NewGameSessionInvitationReceived()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.SenderPID = types.NewUInt32(6059)
		v.Message = types.NewString("s7")
		v.CreationTime = types.NewDateTime(8000101)
		return v
	}(), func() game_session_types.GameSessionInvitationReceived {
		v := game_session_types.NewGameSessionInvitationReceived()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(11104)
			v.SessionID = types.NewUInt32(12113)
			return v
		}()
		v.SenderPID = types.NewUInt32(13122)
		v.Message = types.NewString("s14")
		v.CreationTime = types.NewDateTime(15000122)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x34, 0x00, 0x3a, 0xe2, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GetInvitationsReceivedResponse

	roundTrip(t, "GetInvitationsReceivedResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Invitations.Equals(want.Invitations)) {
		t.Fatalf("GetInvitationsReceivedResponse.Invitations differs")
	}
}

func TestGetInvitationSentCountRequest_NoHeader(t *testing.T) {
	want := GetInvitationSentCountRequest{GameSessionTypeID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetInvitationSentCountRequest

	roundTrip(t, "GetInvitationSentCountRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("GetInvitationSentCountRequest.GameSessionTypeID differs")
	}
}

func TestGetInvitationSentCountRequest_Header(t *testing.T) {
	want := GetInvitationSentCountRequest{GameSessionTypeID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetInvitationSentCountRequest

	roundTrip(t, "GetInvitationSentCountRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("GetInvitationSentCountRequest.GameSessionTypeID differs")
	}
}

func TestGetInvitationSentCountResponse_NoHeader(t *testing.T) {
	want := GetInvitationSentCountResponse{Count: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetInvitationSentCountResponse

	roundTrip(t, "GetInvitationSentCountResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("GetInvitationSentCountResponse.Count differs")
	}
}

func TestGetInvitationSentCountResponse_Header(t *testing.T) {
	want := GetInvitationSentCountResponse{Count: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetInvitationSentCountResponse

	roundTrip(t, "GetInvitationSentCountResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("GetInvitationSentCountResponse.Count differs")
	}
}

func TestGetInvitationsSentRequest_NoHeader(t *testing.T) {
	want := GetInvitationsSentRequest{GameSessionTypeID: types.NewUInt32(1014), ResultRange: types.ResultRange{Offset: types.NewUInt32(6), Length: types.NewUInt32(16)}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00}

	var got GetInvitationsSentRequest

	roundTrip(t, "GetInvitationsSentRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("GetInvitationsSentRequest.GameSessionTypeID differs")
	}
	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetInvitationsSentRequest.ResultRange differs")
	}
}

func TestGetInvitationsSentRequest_Header(t *testing.T) {
	want := GetInvitationsSentRequest{GameSessionTypeID: types.NewUInt32(1014), ResultRange: types.ResultRange{Offset: types.NewUInt32(6), Length: types.NewUInt32(16)}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00}

	var got GetInvitationsSentRequest

	roundTrip(t, "GetInvitationsSentRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("GetInvitationsSentRequest.GameSessionTypeID differs")
	}
	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetInvitationsSentRequest.ResultRange differs")
	}
}

func TestGetInvitationsSentResponse_NoHeader(t *testing.T) {
	want := GetInvitationsSentResponse{Invitations: types.List[game_session_types.GameSessionInvitationSent]{func() game_session_types.GameSessionInvitationSent {
		v := game_session_types.NewGameSessionInvitationSent()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.RecipientPID = types.NewUInt32(6059)
		v.Message = types.NewString("s7")
		v.CreationTime = types.NewDateTime(8000101)
		return v
	}(), func() game_session_types.GameSessionInvitationSent {
		v := game_session_types.NewGameSessionInvitationSent()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(11104)
			v.SessionID = types.NewUInt32(12113)
			return v
		}()
		v.RecipientPID = types.NewUInt32(13122)
		v.Message = types.NewString("s14")
		v.CreationTime = types.NewDateTime(15000122)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x34, 0x00, 0x3a, 0xe2, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GetInvitationsSentResponse

	roundTrip(t, "GetInvitationsSentResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Invitations.Equals(want.Invitations)) {
		t.Fatalf("GetInvitationsSentResponse.Invitations differs")
	}
}

func TestGetInvitationsSentResponse_Header(t *testing.T) {
	want := GetInvitationsSentResponse{Invitations: types.List[game_session_types.GameSessionInvitationSent]{func() game_session_types.GameSessionInvitationSent {
		v := game_session_types.NewGameSessionInvitationSent()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.RecipientPID = types.NewUInt32(6059)
		v.Message = types.NewString("s7")
		v.CreationTime = types.NewDateTime(8000101)
		return v
	}(), func() game_session_types.GameSessionInvitationSent {
		v := game_session_types.NewGameSessionInvitationSent()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(11104)
			v.SessionID = types.NewUInt32(12113)
			return v
		}()
		v.RecipientPID = types.NewUInt32(13122)
		v.Message = types.NewString("s14")
		v.CreationTime = types.NewDateTime(15000122)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x34, 0x00, 0x3a, 0xe2, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GetInvitationsSentResponse

	roundTrip(t, "GetInvitationsSentResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Invitations.Equals(want.Invitations)) {
		t.Fatalf("GetInvitationsSentResponse.Invitations differs")
	}
}

func TestAcceptInvitationRequest_NoHeader(t *testing.T) {
	want := AcceptInvitationRequest{GameSessionInvitation: func() game_session_types.GameSessionInvitationReceived {
		v := game_session_types.NewGameSessionInvitationReceived()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.SenderPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()}

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got AcceptInvitationRequest

	roundTrip(t, "AcceptInvitationRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionInvitation.Equals(want.GameSessionInvitation)) {
		t.Fatalf("AcceptInvitationRequest.GameSessionInvitation differs")
	}
}

func TestAcceptInvitationRequest_Header(t *testing.T) {
	want := AcceptInvitationRequest{GameSessionInvitation: func() game_session_types.GameSessionInvitationReceived {
		v := game_session_types.NewGameSessionInvitationReceived()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.SenderPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()}

	data := []byte{0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got AcceptInvitationRequest

	roundTrip(t, "AcceptInvitationRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionInvitation.Equals(want.GameSessionInvitation)) {
		t.Fatalf("AcceptInvitationRequest.GameSessionInvitation differs")
	}
}

func TestDeclineInvitationRequest_NoHeader(t *testing.T) {
	want := DeclineInvitationRequest{GameSessionInvitation: func() game_session_types.GameSessionInvitationReceived {
		v := game_session_types.NewGameSessionInvitationReceived()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.SenderPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()}

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got DeclineInvitationRequest

	roundTrip(t, "DeclineInvitationRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionInvitation.Equals(want.GameSessionInvitation)) {
		t.Fatalf("DeclineInvitationRequest.GameSessionInvitation differs")
	}
}

func TestDeclineInvitationRequest_Header(t *testing.T) {
	want := DeclineInvitationRequest{GameSessionInvitation: func() game_session_types.GameSessionInvitationReceived {
		v := game_session_types.NewGameSessionInvitationReceived()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.SenderPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()}

	data := []byte{0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got DeclineInvitationRequest

	roundTrip(t, "DeclineInvitationRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionInvitation.Equals(want.GameSessionInvitation)) {
		t.Fatalf("DeclineInvitationRequest.GameSessionInvitation differs")
	}
}

func TestCancelInvitationRequest_NoHeader(t *testing.T) {
	want := CancelInvitationRequest{GameSessionInvitation: func() game_session_types.GameSessionInvitationSent {
		v := game_session_types.NewGameSessionInvitationSent()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.RecipientPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()}

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got CancelInvitationRequest

	roundTrip(t, "CancelInvitationRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionInvitation.Equals(want.GameSessionInvitation)) {
		t.Fatalf("CancelInvitationRequest.GameSessionInvitation differs")
	}
}

func TestCancelInvitationRequest_Header(t *testing.T) {
	want := CancelInvitationRequest{GameSessionInvitation: func() game_session_types.GameSessionInvitationSent {
		v := game_session_types.NewGameSessionInvitationSent()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.RecipientPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()}

	data := []byte{0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got CancelInvitationRequest

	roundTrip(t, "CancelInvitationRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionInvitation.Equals(want.GameSessionInvitation)) {
		t.Fatalf("CancelInvitationRequest.GameSessionInvitation differs")
	}
}

func TestSendTextMessageRequest_NoHeader(t *testing.T) {
	want := SendTextMessageRequest{GameSessionMessage: func() game_session_types.GameSessionMessage {
		v := game_session_types.NewGameSessionMessage()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.Message = types.NewString("s5")
		return v
	}()}

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got SendTextMessageRequest

	roundTrip(t, "SendTextMessageRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionMessage.Equals(want.GameSessionMessage)) {
		t.Fatalf("SendTextMessageRequest.GameSessionMessage differs")
	}
}

func TestSendTextMessageRequest_Header(t *testing.T) {
	want := SendTextMessageRequest{GameSessionMessage: func() game_session_types.GameSessionMessage {
		v := game_session_types.NewGameSessionMessage()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.Message = types.NewString("s5")
		return v
	}()}

	data := []byte{0x00, 0x12, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got SendTextMessageRequest

	roundTrip(t, "SendTextMessageRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionMessage.Equals(want.GameSessionMessage)) {
		t.Fatalf("SendTextMessageRequest.GameSessionMessage differs")
	}
}

func TestRegisterURLsRequest_NoHeader(t *testing.T) {
	want := RegisterURLsRequest{StationURLs: types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.2"), types.NewStationURL("prudp:/address=10.0.0.3")}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x33, 0x00}

	var got RegisterURLsRequest

	roundTrip(t, "RegisterURLsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.StationURLs.Equals(want.StationURLs)) {
		t.Fatalf("RegisterURLsRequest.StationURLs differs")
	}
}

func TestRegisterURLsRequest_Header(t *testing.T) {
	want := RegisterURLsRequest{StationURLs: types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.2"), types.NewStationURL("prudp:/address=10.0.0.3")}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x33, 0x00}

	var got RegisterURLsRequest

	roundTrip(t, "RegisterURLsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.StationURLs.Equals(want.StationURLs)) {
		t.Fatalf("RegisterURLsRequest.StationURLs differs")
	}
}

func TestJoinSessionRequest_NoHeader(t *testing.T) {
	want := JoinSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got JoinSessionRequest

	roundTrip(t, "JoinSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("JoinSessionRequest.GameSessionKey differs")
	}
}

func TestJoinSessionRequest_Header(t *testing.T) {
	want := JoinSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got JoinSessionRequest

	roundTrip(t, "JoinSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("JoinSessionRequest.GameSessionKey differs")
	}
}

func TestAbandonSessionRequest_NoHeader(t *testing.T) {
	want := AbandonSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got AbandonSessionRequest

	roundTrip(t, "AbandonSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("AbandonSessionRequest.GameSessionKey differs")
	}
}

func TestAbandonSessionRequest_Header(t *testing.T) {
	want := AbandonSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got AbandonSessionRequest

	roundTrip(t, "AbandonSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("AbandonSessionRequest.GameSessionKey differs")
	}
}

func TestSearchSessionsWithParticipantsRequest_NoHeader(t *testing.T) {
	want := SearchSessionsWithParticipantsRequest{GameSessionTypeID: types.NewUInt32(1014), ParticipantIDs: types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got SearchSessionsWithParticipantsRequest

	roundTrip(t, "SearchSessionsWithParticipantsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("SearchSessionsWithParticipantsRequest.GameSessionTypeID differs")
	}
	if !(got.ParticipantIDs.Equals(want.ParticipantIDs)) {
		t.Fatalf("SearchSessionsWithParticipantsRequest.ParticipantIDs differs")
	}
}

func TestSearchSessionsWithParticipantsRequest_Header(t *testing.T) {
	want := SearchSessionsWithParticipantsRequest{GameSessionTypeID: types.NewUInt32(1014), ParticipantIDs: types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got SearchSessionsWithParticipantsRequest

	roundTrip(t, "SearchSessionsWithParticipantsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionTypeID.Equals(want.GameSessionTypeID)) {
		t.Fatalf("SearchSessionsWithParticipantsRequest.GameSessionTypeID differs")
	}
	if !(got.ParticipantIDs.Equals(want.ParticipantIDs)) {
		t.Fatalf("SearchSessionsWithParticipantsRequest.ParticipantIDs differs")
	}
}

func TestSearchSessionsWithParticipantsResponse_NoHeader(t *testing.T) {
	want := SearchSessionsWithParticipantsResponse{SearchResults: types.List[game_session_types.GameSessionSearchWithParticipantsResult]{func() game_session_types.GameSessionSearchWithParticipantsResult {
		v := game_session_types.NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		return v
	}(), func() game_session_types.GameSessionSearchWithParticipantsResult {
		v := game_session_types.NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got SearchSessionsWithParticipantsResponse

	roundTrip(t, "SearchSessionsWithParticipantsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchSessionsWithParticipantsResponse.SearchResults differs")
	}
}

func TestSearchSessionsWithParticipantsResponse_Header(t *testing.T) {
	want := SearchSessionsWithParticipantsResponse{SearchResults: types.List[game_session_types.GameSessionSearchWithParticipantsResult]{func() game_session_types.GameSessionSearchWithParticipantsResult {
		v := game_session_types.NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		return v
	}(), func() game_session_types.GameSessionSearchWithParticipantsResult {
		v := game_session_types.NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got SearchSessionsWithParticipantsResponse

	roundTrip(t, "SearchSessionsWithParticipantsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchSessionsWithParticipantsResponse.SearchResults differs")
	}
}

func TestGetSessionsRequest_NoHeader(t *testing.T) {
	want := GetSessionsRequest{GameSessionKeys: types.List[game_session_types.GameSessionKey]{func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(3032)
		v.SessionID = types.NewUInt32(4041)
		return v
	}(), func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(6059)
		v.SessionID = types.NewUInt32(7068)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00}

	var got GetSessionsRequest

	roundTrip(t, "GetSessionsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKeys.Equals(want.GameSessionKeys)) {
		t.Fatalf("GetSessionsRequest.GameSessionKeys differs")
	}
}

func TestGetSessionsRequest_Header(t *testing.T) {
	want := GetSessionsRequest{GameSessionKeys: types.List[game_session_types.GameSessionKey]{func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(3032)
		v.SessionID = types.NewUInt32(4041)
		return v
	}(), func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(6059)
		v.SessionID = types.NewUInt32(7068)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00}

	var got GetSessionsRequest

	roundTrip(t, "GetSessionsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKeys.Equals(want.GameSessionKeys)) {
		t.Fatalf("GetSessionsRequest.GameSessionKeys differs")
	}
}

func TestGetSessionsResponse_NoHeader(t *testing.T) {
	want := GetSessionsResponse{SearchResults: types.List[game_session_types.GameSessionSearchResult]{func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.HostPID = types.NewUInt32(6059)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.8"), types.NewStationURL("prudp:/address=10.0.0.9")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.NewUInt32(13122)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(15140)
			v.Value = types.NewUInt32(16149)
			return v
		}()}
		return v
	}(), func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(19176)
			v.SessionID = types.NewUInt32(20185)
			return v
		}()
		v.HostPID = types.NewUInt32(21194)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.23"), types.NewStationURL("prudp:/address=10.0.0.24")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(27248)
			v.Value = types.NewUInt32(28257)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(30275)
			v.Value = types.NewUInt32(31284)
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x38, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x39, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x33, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x34, 0x00, 0x02, 0x00, 0x00, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x61, 0x6e, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00}

	var got GetSessionsResponse

	roundTrip(t, "GetSessionsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("GetSessionsResponse.SearchResults differs")
	}
}

func TestGetSessionsResponse_Header(t *testing.T) {
	want := GetSessionsResponse{SearchResults: types.List[game_session_types.GameSessionSearchResult]{func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.HostPID = types.NewUInt32(6059)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.8"), types.NewStationURL("prudp:/address=10.0.0.9")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.NewUInt32(13122)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(15140)
			v.Value = types.NewUInt32(16149)
			return v
		}()}
		return v
	}(), func() game_session_types.GameSessionSearchResult {
		v := game_session_types.NewGameSessionSearchResult()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(19176)
			v.SessionID = types.NewUInt32(20185)
			return v
		}()
		v.HostPID = types.NewUInt32(21194)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.23"), types.NewStationURL("prudp:/address=10.0.0.24")}
		v.Attributes = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(27248)
			v.Value = types.NewUInt32(28257)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(30275)
			v.Value = types.NewUInt32(31284)
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x38, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x39, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x33, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x32, 0x34, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x61, 0x6e, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00}

	var got GetSessionsResponse

	roundTrip(t, "GetSessionsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("GetSessionsResponse.SearchResults differs")
	}
}

func TestGetParticipantsURLsRequest_NoHeader(t *testing.T) {
	want := GetParticipantsURLsRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}(), ParticipantIDs: types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00}

	var got GetParticipantsURLsRequest

	roundTrip(t, "GetParticipantsURLsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("GetParticipantsURLsRequest.GameSessionKey differs")
	}
	if !(got.ParticipantIDs.Equals(want.ParticipantIDs)) {
		t.Fatalf("GetParticipantsURLsRequest.ParticipantIDs differs")
	}
}

func TestGetParticipantsURLsRequest_Header(t *testing.T) {
	want := GetParticipantsURLsRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}(), ParticipantIDs: types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00}

	var got GetParticipantsURLsRequest

	roundTrip(t, "GetParticipantsURLsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("GetParticipantsURLsRequest.GameSessionKey differs")
	}
	if !(got.ParticipantIDs.Equals(want.ParticipantIDs)) {
		t.Fatalf("GetParticipantsURLsRequest.ParticipantIDs differs")
	}
}

func TestGetParticipantsURLsResponse_NoHeader(t *testing.T) {
	want := GetParticipantsURLsResponse{Participants: types.List[game_session_types.GameSessionParticipant]{func() game_session_types.GameSessionParticipant {
		v := game_session_types.NewGameSessionParticipant()
		v.PID = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.6"), types.NewStationURL("prudp:/address=10.0.0.7")}
		return v
	}(), func() game_session_types.GameSessionParticipant {
		v := game_session_types.NewGameSessionParticipant()
		v.PID = types.NewUInt32(9086)
		v.Name = types.NewString("s10")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.12"), types.NewStationURL("prudp:/address=10.0.0.13")}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x36, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x37, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x02, 0x00, 0x00, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x32, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x33, 0x00}

	var got GetParticipantsURLsResponse

	roundTrip(t, "GetParticipantsURLsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Participants.Equals(want.Participants)) {
		t.Fatalf("GetParticipantsURLsResponse.Participants differs")
	}
}

func TestGetParticipantsURLsResponse_Header(t *testing.T) {
	want := GetParticipantsURLsResponse{Participants: types.List[game_session_types.GameSessionParticipant]{func() game_session_types.GameSessionParticipant {
		v := game_session_types.NewGameSessionParticipant()
		v.PID = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.6"), types.NewStationURL("prudp:/address=10.0.0.7")}
		return v
	}(), func() game_session_types.GameSessionParticipant {
		v := game_session_types.NewGameSessionParticipant()
		v.PID = types.NewUInt32(9086)
		v.Name = types.NewString("s10")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.12"), types.NewStationURL("prudp:/address=10.0.0.13")}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x36, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x37, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x02, 0x00, 0x00, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x32, 0x00, 0x19, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x33, 0x00}

	var got GetParticipantsURLsResponse

	roundTrip(t, "GetParticipantsURLsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Participants.Equals(want.Participants)) {
		t.Fatalf("GetParticipantsURLsResponse.Participants differs")
	}
}

func TestMigrateSessionHostRequest_NoHeader(t *testing.T) {
	want := MigrateSessionHostRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got MigrateSessionHostRequest

	roundTrip(t, "MigrateSessionHostRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("MigrateSessionHostRequest.GameSessionKey differs")
	}
}

func TestMigrateSessionHostRequest_Header(t *testing.T) {
	want := MigrateSessionHostRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got MigrateSessionHostRequest

	roundTrip(t, "MigrateSessionHostRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("MigrateSessionHostRequest.GameSessionKey differs")
	}
}

func TestSplitSessionRequest_NoHeader(t *testing.T) {
	want := SplitSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got SplitSessionRequest

	roundTrip(t, "SplitSessionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("SplitSessionRequest.GameSessionKey differs")
	}
}

func TestSplitSessionRequest_Header(t *testing.T) {
	want := SplitSessionRequest{GameSessionKey: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got SplitSessionRequest

	roundTrip(t, "SplitSessionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKey.Equals(want.GameSessionKey)) {
		t.Fatalf("SplitSessionRequest.GameSessionKey differs")
	}
}

func TestSplitSessionResponse_NoHeader(t *testing.T) {
	want := SplitSessionResponse{GameSessionKeyMigrated: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got SplitSessionResponse

	roundTrip(t, "SplitSessionResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKeyMigrated.Equals(want.GameSessionKeyMigrated)) {
		t.Fatalf("SplitSessionResponse.GameSessionKeyMigrated differs")
	}
}

func TestSplitSessionResponse_Header(t *testing.T) {
	want := SplitSessionResponse{GameSessionKeyMigrated: func() game_session_types.GameSessionKey {
		v := game_session_types.NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got SplitSessionResponse

	roundTrip(t, "SplitSessionResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionKeyMigrated.Equals(want.GameSessionKeyMigrated)) {
		t.Fatalf("SplitSessionResponse.GameSessionKeyMigrated differs")
	}
}

func TestSearchSocialSessionsRequest_NoHeader(t *testing.T) {
	want := SearchSocialSessionsRequest{GameSessionSocialQuery: func() game_session_types.GameSessionSocialQuery {
		v := game_session_types.NewGameSessionSocialQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.Parameters = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(6059)
			v.Value = types.NewUInt32(7068)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.NewUInt32(10095)
			return v
		}()}
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00}

	var got SearchSocialSessionsRequest

	roundTrip(t, "SearchSocialSessionsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionSocialQuery.Equals(want.GameSessionSocialQuery)) {
		t.Fatalf("SearchSocialSessionsRequest.GameSessionSocialQuery differs")
	}
}

func TestSearchSocialSessionsRequest_Header(t *testing.T) {
	want := SearchSocialSessionsRequest{GameSessionSocialQuery: func() game_session_types.GameSessionSocialQuery {
		v := game_session_types.NewGameSessionSocialQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.Parameters = types.List[game_session_types.Property]{func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(6059)
			v.Value = types.NewUInt32(7068)
			return v
		}(), func() game_session_types.Property {
			v := game_session_types.NewProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.NewUInt32(10095)
			return v
		}()}
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}()}

	data := []byte{0x00, 0x32, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00}

	var got SearchSocialSessionsRequest

	roundTrip(t, "SearchSocialSessionsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.GameSessionSocialQuery.Equals(want.GameSessionSocialQuery)) {
		t.Fatalf("SearchSocialSessionsRequest.GameSessionSocialQuery differs")
	}
}

func TestSearchSocialSessionsResponse_NoHeader(t *testing.T) {
	want := SearchSocialSessionsResponse{SearchResults: types.List[game_session_types.GameSessionSearchWithParticipantsResult]{func() game_session_types.GameSessionSearchWithParticipantsResult {
		v := game_session_types.NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		return v
	}(), func() game_session_types.GameSessionSearchWithParticipantsResult {
		v := game_session_types.NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got SearchSocialSessionsResponse

	roundTrip(t, "SearchSocialSessionsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchSocialSessionsResponse.SearchResults differs")
	}
}

func TestSearchSocialSessionsResponse_Header(t *testing.T) {
	want := SearchSocialSessionsResponse{SearchResults: types.List[game_session_types.GameSessionSearchWithParticipantsResult]{func() game_session_types.GameSessionSearchWithParticipantsResult {
		v := game_session_types.NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		return v
	}(), func() game_session_types.GameSessionSearchWithParticipantsResult {
		v := game_session_types.NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got SearchSocialSessionsResponse

	roundTrip(t, "SearchSocialSessionsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchSocialSessionsResponse.SearchResults differs")
	}
}

func TestReportUnsuccessfulJoinSessionsRequest_NoHeader(t *testing.T) {
	want := ReportUnsuccessfulJoinSessionsRequest{UnsuccessfulJoinSessions: types.List[game_session_types.GameSessionUnsuccessfulJoinSession]{func() game_session_types.GameSessionUnsuccessfulJoinSession {
		v := game_session_types.NewGameSessionUnsuccessfulJoinSession()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.ErrorCategory = types.NewUInt32(6059)
		v.ErrorCode = types.NewInt32(-218)
		return v
	}(), func() game_session_types.GameSessionUnsuccessfulJoinSession {
		v := game_session_types.NewGameSessionUnsuccessfulJoinSession()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(10095)
			v.SessionID = types.NewUInt32(11104)
			return v
		}()
		v.ErrorCategory = types.NewUInt32(12113)
		v.ErrorCode = types.NewInt32(-404)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x26, 0xff, 0xff, 0xff, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x6c, 0xfe, 0xff, 0xff}

	var got ReportUnsuccessfulJoinSessionsRequest

	roundTrip(t, "ReportUnsuccessfulJoinSessionsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.UnsuccessfulJoinSessions.Equals(want.UnsuccessfulJoinSessions)) {
		t.Fatalf("ReportUnsuccessfulJoinSessionsRequest.UnsuccessfulJoinSessions differs")
	}
}

func TestReportUnsuccessfulJoinSessionsRequest_Header(t *testing.T) {
	want := ReportUnsuccessfulJoinSessionsRequest{UnsuccessfulJoinSessions: types.List[game_session_types.GameSessionUnsuccessfulJoinSession]{func() game_session_types.GameSessionUnsuccessfulJoinSession {
		v := game_session_types.NewGameSessionUnsuccessfulJoinSession()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(4041)
			v.SessionID = types.NewUInt32(5050)
			return v
		}()
		v.ErrorCategory = types.NewUInt32(6059)
		v.ErrorCode = types.NewInt32(-218)
		return v
	}(), func() game_session_types.GameSessionUnsuccessfulJoinSession {
		v := game_session_types.NewGameSessionUnsuccessfulJoinSession()
		v.SessionKey = func() game_session_types.GameSessionKey {
			v := game_session_types.NewGameSessionKey()
			v.TypeID = types.NewUInt32(10095)
			v.SessionID = types.NewUInt32(11104)
			return v
		}()
		v.ErrorCategory = types.NewUInt32(12113)
		v.ErrorCode = types.NewInt32(-404)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x26, 0xff, 0xff, 0xff, 0x00, 0x15, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x6c, 0xfe, 0xff, 0xff}

	var got ReportUnsuccessfulJoinSessionsRequest

	roundTrip(t, "ReportUnsuccessfulJoinSessionsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.UnsuccessfulJoinSessions.Equals(want.UnsuccessfulJoinSessions)) {
		t.Fatalf("ReportUnsuccessfulJoinSessionsRequest.UnsuccessfulJoinSessions differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodCreateSession != 1 {
		t.Fatalf("MethodCreateSession = %d, want 1", MethodCreateSession)
	}
	if MethodUpdateSession != 2 {
		t.Fatalf("MethodUpdateSession = %d, want 2", MethodUpdateSession)
	}
	if MethodDeleteSession != 3 {
		t.Fatalf("MethodDeleteSession = %d, want 3", MethodDeleteSession)
	}
	if MethodMigrateSession != 4 {
		t.Fatalf("MethodMigrateSession = %d, want 4", MethodMigrateSession)
	}
	if MethodLeaveSession != 5 {
		t.Fatalf("MethodLeaveSession = %d, want 5", MethodLeaveSession)
	}
	if MethodGetSession != 6 {
		t.Fatalf("MethodGetSession = %d, want 6", MethodGetSession)
	}
	if MethodSearchSessions != 7 {
		t.Fatalf("MethodSearchSessions = %d, want 7", MethodSearchSessions)
	}
	if MethodAddParticipants != 8 {
		t.Fatalf("MethodAddParticipants = %d, want 8", MethodAddParticipants)
	}
	if MethodRemoveParticipants != 9 {
		t.Fatalf("MethodRemoveParticipants = %d, want 9", MethodRemoveParticipants)
	}
	if MethodGetParticipantCount != 10 {
		t.Fatalf("MethodGetParticipantCount = %d, want 10", MethodGetParticipantCount)
	}
	if MethodGetParticipants != 11 {
		t.Fatalf("MethodGetParticipants = %d, want 11", MethodGetParticipants)
	}
	if MethodSendInvitation != 12 {
		t.Fatalf("MethodSendInvitation = %d, want 12", MethodSendInvitation)
	}
	if MethodGetInvitationReceivedCount != 13 {
		t.Fatalf("MethodGetInvitationReceivedCount = %d, want 13", MethodGetInvitationReceivedCount)
	}
	if MethodGetInvitationsReceived != 14 {
		t.Fatalf("MethodGetInvitationsReceived = %d, want 14", MethodGetInvitationsReceived)
	}
	if MethodGetInvitationSentCount != 15 {
		t.Fatalf("MethodGetInvitationSentCount = %d, want 15", MethodGetInvitationSentCount)
	}
	if MethodGetInvitationsSent != 16 {
		t.Fatalf("MethodGetInvitationsSent = %d, want 16", MethodGetInvitationsSent)
	}
	if MethodAcceptInvitation != 17 {
		t.Fatalf("MethodAcceptInvitation = %d, want 17", MethodAcceptInvitation)
	}
	if MethodDeclineInvitation != 18 {
		t.Fatalf("MethodDeclineInvitation = %d, want 18", MethodDeclineInvitation)
	}
	if MethodCancelInvitation != 19 {
		t.Fatalf("MethodCancelInvitation = %d, want 19", MethodCancelInvitation)
	}
	if MethodSendTextMessage != 20 {
		t.Fatalf("MethodSendTextMessage = %d, want 20", MethodSendTextMessage)
	}
	if MethodRegisterURLs != 21 {
		t.Fatalf("MethodRegisterURLs = %d, want 21", MethodRegisterURLs)
	}
	if MethodJoinSession != 22 {
		t.Fatalf("MethodJoinSession = %d, want 22", MethodJoinSession)
	}
	if MethodAbandonSession != 23 {
		t.Fatalf("MethodAbandonSession = %d, want 23", MethodAbandonSession)
	}
	if MethodSearchSessionsWithParticipants != 24 {
		t.Fatalf("MethodSearchSessionsWithParticipants = %d, want 24", MethodSearchSessionsWithParticipants)
	}
	if MethodGetSessions != 25 {
		t.Fatalf("MethodGetSessions = %d, want 25", MethodGetSessions)
	}
	if MethodGetParticipantsURLs != 26 {
		t.Fatalf("MethodGetParticipantsURLs = %d, want 26", MethodGetParticipantsURLs)
	}
	if MethodMigrateSessionHost != 27 {
		t.Fatalf("MethodMigrateSessionHost = %d, want 27", MethodMigrateSessionHost)
	}
	if MethodSplitSession != 28 {
		t.Fatalf("MethodSplitSession = %d, want 28", MethodSplitSession)
	}
	if MethodSearchSocialSessions != 29 {
		t.Fatalf("MethodSearchSocialSessions = %d, want 29", MethodSearchSocialSessions)
	}
	if MethodReportUnsuccessfulJoinSessions != 30 {
		t.Fatalf("MethodReportUnsuccessfulJoinSessions = %d, want 30", MethodReportUnsuccessfulJoinSessions)
	}
	if ProtocolID != 42 {
		t.Fatalf("ProtocolID = %d, want 42", ProtocolID)
	}
}
