package types

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
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

func TestGameSessionKey_NoHeader(t *testing.T) {
	want := func() GameSessionKey {
		v := NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GameSessionKey

	roundTrip(t, "GameSessionKey", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionKey differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionKey_Header(t *testing.T) {
	want := func() GameSessionKey {
		v := NewGameSessionKey()
		v.TypeID = types.NewUInt32(2023)
		v.SessionID = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GameSessionKey

	roundTrip(t, "GameSessionKey", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionKey differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSession_NoHeader(t *testing.T) {
	want := func() GameSession {
		v := NewGameSession()
		v.TypeID = types.NewUInt32(2023)
		v.Attributes = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(5050)
			v.Value = types.NewUInt32(6059)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(8077)
			v.Value = types.NewUInt32(9086)
			return v
		}()}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got GameSession

	roundTrip(t, "GameSession", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSession differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSession_Header(t *testing.T) {
	want := func() GameSession {
		v := NewGameSession()
		v.TypeID = types.NewUInt32(2023)
		v.Attributes = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(5050)
			v.Value = types.NewUInt32(6059)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(8077)
			v.Value = types.NewUInt32(9086)
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x22, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00}

	var got GameSession

	roundTrip(t, "GameSession", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSession differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionSearchResult_NoHeader(t *testing.T) {
	want := func() GameSessionSearchResult {
		v := NewGameSessionSearchResult()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.HostPID = types.NewUInt32(5050)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.7"), types.NewStationURL("prudp:/address=10.0.0.8")}
		v.Attributes = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(11104)
			v.Value = types.NewUInt32(12113)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(14131)
			v.Value = types.NewUInt32(15140)
			return v
		}()}
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x37, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x38, 0x00, 0x02, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00}

	var got GameSessionSearchResult

	roundTrip(t, "GameSessionSearchResult", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionSearchResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionSearchResult_Header(t *testing.T) {
	want := func() GameSessionSearchResult {
		v := NewGameSessionSearchResult()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.HostPID = types.NewUInt32(5050)
		v.HostURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.7"), types.NewStationURL("prudp:/address=10.0.0.8")}
		v.Attributes = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(11104)
			v.Value = types.NewUInt32(12113)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(14131)
			v.Value = types.NewUInt32(15140)
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x37, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x38, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00}

	var got GameSessionSearchResult

	roundTrip(t, "GameSessionSearchResult", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionSearchResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionUpdate_NoHeader(t *testing.T) {
	want := func() GameSessionUpdate {
		v := NewGameSessionUpdate()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.Attributes = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.NewUInt32(8077)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.NewUInt32(11104)
			return v
		}()}
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00}

	var got GameSessionUpdate

	roundTrip(t, "GameSessionUpdate", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionUpdate differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionUpdate_Header(t *testing.T) {
	want := func() GameSessionUpdate {
		v := NewGameSessionUpdate()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.Attributes = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.NewUInt32(8077)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.NewUInt32(11104)
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x2b, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00}

	var got GameSessionUpdate

	roundTrip(t, "GameSessionUpdate", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionUpdate differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionParticipant_NoHeader(t *testing.T) {
	want := func() GameSessionParticipant {
		v := NewGameSessionParticipant()
		v.PID = types.NewUInt32(2023)
		v.Name = types.NewString("s3")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.5"), types.NewStationURL("prudp:/address=10.0.0.6")}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x35, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x36, 0x00}

	var got GameSessionParticipant

	roundTrip(t, "GameSessionParticipant", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionParticipant differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionParticipant_Header(t *testing.T) {
	want := func() GameSessionParticipant {
		v := NewGameSessionParticipant()
		v.PID = types.NewUInt32(2023)
		v.Name = types.NewString("s3")
		v.StationURLs = types.List[types.StationURL]{types.NewStationURL("prudp:/address=10.0.0.5"), types.NewStationURL("prudp:/address=10.0.0.6")}
		return v
	}()

	data := []byte{0x00, 0x41, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x35, 0x00, 0x18, 0x00, 0x70, 0x72, 0x75, 0x64, 0x70, 0x3a, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x36, 0x00}

	var got GameSessionParticipant

	roundTrip(t, "GameSessionParticipant", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionParticipant differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionInvitation_NoHeader(t *testing.T) {
	want := func() GameSessionInvitation {
		v := NewGameSessionInvitation()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.RecipientPIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		v.Message = types.NewString("s8")
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00}

	var got GameSessionInvitation

	roundTrip(t, "GameSessionInvitation", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionInvitation differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionInvitation_Header(t *testing.T) {
	want := func() GameSessionInvitation {
		v := NewGameSessionInvitation()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.RecipientPIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		v.Message = types.NewString("s8")
		return v
	}()

	data := []byte{0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00}

	var got GameSessionInvitation

	roundTrip(t, "GameSessionInvitation", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionInvitation differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionInvitationSent_NoHeader(t *testing.T) {
	want := func() GameSessionInvitationSent {
		v := NewGameSessionInvitationSent()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.RecipientPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GameSessionInvitationSent

	roundTrip(t, "GameSessionInvitationSent", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionInvitationSent differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionInvitationSent_Header(t *testing.T) {
	want := func() GameSessionInvitationSent {
		v := NewGameSessionInvitationSent()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.RecipientPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()

	data := []byte{0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GameSessionInvitationSent

	roundTrip(t, "GameSessionInvitationSent", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionInvitationSent differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionInvitationReceived_NoHeader(t *testing.T) {
	want := func() GameSessionInvitationReceived {
		v := NewGameSessionInvitationReceived()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.SenderPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GameSessionInvitationReceived

	roundTrip(t, "GameSessionInvitationReceived", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionInvitationReceived differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionInvitationReceived_Header(t *testing.T) {
	want := func() GameSessionInvitationReceived {
		v := NewGameSessionInvitationReceived()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.SenderPID = types.NewUInt32(5050)
		v.Message = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		return v
	}()

	data := []byte{0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GameSessionInvitationReceived

	roundTrip(t, "GameSessionInvitationReceived", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionInvitationReceived differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionQuery_NoHeader(t *testing.T) {
	want := func() GameSessionQuery {
		v := NewGameSessionQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.Parameters = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(6059)
			v.Value = types.NewUInt32(7068)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.NewUInt32(10095)
			return v
		}()}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00}

	var got GameSessionQuery

	roundTrip(t, "GameSessionQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionQuery_Header(t *testing.T) {
	want := func() GameSessionQuery {
		v := NewGameSessionQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.Parameters = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(6059)
			v.Value = types.NewUInt32(7068)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.NewUInt32(10095)
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x26, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00}

	var got GameSessionQuery

	roundTrip(t, "GameSessionQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionSocialQuery_NoHeader(t *testing.T) {
	want := func() GameSessionSocialQuery {
		v := NewGameSessionSocialQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.Parameters = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(6059)
			v.Value = types.NewUInt32(7068)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.NewUInt32(10095)
			return v
		}()}
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00}

	var got GameSessionSocialQuery

	roundTrip(t, "GameSessionSocialQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionSocialQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionSocialQuery_Header(t *testing.T) {
	want := func() GameSessionSocialQuery {
		v := NewGameSessionSocialQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.Parameters = types.List[Property]{func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(6059)
			v.Value = types.NewUInt32(7068)
			return v
		}(), func() Property {
			v := NewProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.NewUInt32(10095)
			return v
		}()}
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}()

	data := []byte{0x00, 0x32, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00}

	var got GameSessionSocialQuery

	roundTrip(t, "GameSessionSocialQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionSocialQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionMessage_NoHeader(t *testing.T) {
	want := func() GameSessionMessage {
		v := NewGameSessionMessage()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.Message = types.NewString("s5")
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got GameSessionMessage

	roundTrip(t, "GameSessionMessage", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionMessage differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionMessage_Header(t *testing.T) {
	want := func() GameSessionMessage {
		v := NewGameSessionMessage()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.Message = types.NewString("s5")
		return v
	}()

	data := []byte{0x00, 0x12, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got GameSessionMessage

	roundTrip(t, "GameSessionMessage", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionMessage differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionSearchWithParticipantsResult_NoHeader(t *testing.T) {
	want := func() GameSessionSearchWithParticipantsResult {
		v := NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}
		return v
	}()

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got GameSessionSearchWithParticipantsResult

	roundTrip(t, "GameSessionSearchWithParticipantsResult", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionSearchWithParticipantsResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionSearchWithParticipantsResult_Header(t *testing.T) {
	want := func() GameSessionSearchWithParticipantsResult {
		v := NewGameSessionSearchWithParticipantsResult()
		v.ParticipantIDs = types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}
		return v
	}()

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got GameSessionSearchWithParticipantsResult

	roundTrip(t, "GameSessionSearchWithParticipantsResult", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionSearchWithParticipantsResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionUnsuccessfulJoinSession_NoHeader(t *testing.T) {
	want := func() GameSessionUnsuccessfulJoinSession {
		v := NewGameSessionUnsuccessfulJoinSession()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.ErrorCategory = types.NewUInt32(5050)
		v.ErrorCode = types.NewInt32(-187)
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x45, 0xff, 0xff, 0xff}

	var got GameSessionUnsuccessfulJoinSession

	roundTrip(t, "GameSessionUnsuccessfulJoinSession", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionUnsuccessfulJoinSession differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestGameSessionUnsuccessfulJoinSession_Header(t *testing.T) {
	want := func() GameSessionUnsuccessfulJoinSession {
		v := NewGameSessionUnsuccessfulJoinSession()
		v.SessionKey = func() GameSessionKey {
			v := NewGameSessionKey()
			v.TypeID = types.NewUInt32(3032)
			v.SessionID = types.NewUInt32(4041)
			return v
		}()
		v.ErrorCategory = types.NewUInt32(5050)
		v.ErrorCode = types.NewInt32(-187)
		return v
	}()

	data := []byte{0x00, 0x15, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x45, 0xff, 0xff, 0xff}

	var got GameSessionUnsuccessfulJoinSession

	roundTrip(t, "GameSessionUnsuccessfulJoinSession", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded GameSessionUnsuccessfulJoinSession differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestProperty_NoHeader(t *testing.T) {
	want := func() Property {
		v := NewProperty()
		v.ID = types.NewUInt32(2023)
		v.Value = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got Property

	roundTrip(t, "Property", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded Property differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestProperty_Header(t *testing.T) {
	want := func() Property {
		v := NewProperty()
		v.ID = types.NewUInt32(2023)
		v.Value = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got Property

	roundTrip(t, "Property", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded Property differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}
