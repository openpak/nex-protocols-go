package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
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

func TestGetContentsToModerateRequest_NoHeader(t *testing.T) {
	want := GetContentsToModerateRequest{TypeID: types.NewUInt32(1014), Offset: types.NewUInt32(2023), Size: types.NewUInt32(3032)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetContentsToModerateRequest

	roundTrip(t, "GetContentsToModerateRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("GetContentsToModerateRequest.TypeID differs")
	}
	if !(got.Offset.Equals(want.Offset)) {
		t.Fatalf("GetContentsToModerateRequest.Offset differs")
	}
	if !(got.Size.Equals(want.Size)) {
		t.Fatalf("GetContentsToModerateRequest.Size differs")
	}
}

func TestGetContentsToModerateRequest_Header(t *testing.T) {
	want := GetContentsToModerateRequest{TypeID: types.NewUInt32(1014), Offset: types.NewUInt32(2023), Size: types.NewUInt32(3032)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetContentsToModerateRequest

	roundTrip(t, "GetContentsToModerateRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("GetContentsToModerateRequest.TypeID differs")
	}
	if !(got.Offset.Equals(want.Offset)) {
		t.Fatalf("GetContentsToModerateRequest.Offset differs")
	}
	if !(got.Size.Equals(want.Size)) {
		t.Fatalf("GetContentsToModerateRequest.Size differs")
	}
}

func TestGetContentsToModerateResponse_NoHeader(t *testing.T) {
	want := GetContentsToModerateResponse{Contents: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}, TotalResults: types.NewUInt32(26239)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x7f, 0x66, 0x00, 0x00}

	var got GetContentsToModerateResponse

	roundTrip(t, "GetContentsToModerateResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("GetContentsToModerateResponse.Contents differs")
	}
	if !(got.TotalResults.Equals(want.TotalResults)) {
		t.Fatalf("GetContentsToModerateResponse.TotalResults differs")
	}
}

func TestGetContentsToModerateResponse_Header(t *testing.T) {
	want := GetContentsToModerateResponse{Contents: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}, TotalResults: types.NewUInt32(26239)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x7f, 0x66, 0x00, 0x00}

	var got GetContentsToModerateResponse

	roundTrip(t, "GetContentsToModerateResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("GetContentsToModerateResponse.Contents differs")
	}
	if !(got.TotalResults.Equals(want.TotalResults)) {
		t.Fatalf("GetContentsToModerateResponse.TotalResults differs")
	}
}

func TestFlagContentAsVerifiedRequest_NoHeader(t *testing.T) {
	want := FlagContentAsVerifiedRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got FlagContentAsVerifiedRequest

	roundTrip(t, "FlagContentAsVerifiedRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("FlagContentAsVerifiedRequest.ContentKey differs")
	}
}

func TestFlagContentAsVerifiedRequest_Header(t *testing.T) {
	want := FlagContentAsVerifiedRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got FlagContentAsVerifiedRequest

	roundTrip(t, "FlagContentAsVerifiedRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("FlagContentAsVerifiedRequest.ContentKey differs")
	}
}

func TestBanContentRequest_NoHeader(t *testing.T) {
	want := BanContentRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got BanContentRequest

	roundTrip(t, "BanContentRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("BanContentRequest.ContentKey differs")
	}
}

func TestBanContentRequest_Header(t *testing.T) {
	want := BanContentRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got BanContentRequest

	roundTrip(t, "BanContentRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("BanContentRequest.ContentKey differs")
	}
}

func TestBanUserRequest_NoHeader(t *testing.T) {
	want := BanUserRequest{Pid: types.NewUInt32(1014), Reason: types.NewString("s2"), BanContents: types.NewBool(true), ExpireDate: types.NewDateTime(4000089)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x01, 0x59, 0x09, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got BanUserRequest

	roundTrip(t, "BanUserRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Pid.Equals(want.Pid)) {
		t.Fatalf("BanUserRequest.Pid differs")
	}
	if !(got.Reason.Equals(want.Reason)) {
		t.Fatalf("BanUserRequest.Reason differs")
	}
	if !(got.BanContents.Equals(want.BanContents)) {
		t.Fatalf("BanUserRequest.BanContents differs")
	}
	if !(got.ExpireDate.Equals(want.ExpireDate)) {
		t.Fatalf("BanUserRequest.ExpireDate differs")
	}
}

func TestBanUserRequest_Header(t *testing.T) {
	want := BanUserRequest{Pid: types.NewUInt32(1014), Reason: types.NewString("s2"), BanContents: types.NewBool(true), ExpireDate: types.NewDateTime(4000089)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x01, 0x59, 0x09, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got BanUserRequest

	roundTrip(t, "BanUserRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Pid.Equals(want.Pid)) {
		t.Fatalf("BanUserRequest.Pid differs")
	}
	if !(got.Reason.Equals(want.Reason)) {
		t.Fatalf("BanUserRequest.Reason differs")
	}
	if !(got.BanContents.Equals(want.BanContents)) {
		t.Fatalf("BanUserRequest.BanContents differs")
	}
	if !(got.ExpireDate.Equals(want.ExpireDate)) {
		t.Fatalf("BanUserRequest.ExpireDate differs")
	}
}

func TestBanUserFromContentTypeRequest_NoHeader(t *testing.T) {
	want := BanUserFromContentTypeRequest{TypeID: types.NewUInt32(1014), Pid: types.NewUInt32(2023), Reason: types.NewString("s3"), BanContents: types.NewBool(false), ExpireDate: types.NewDateTime(5000092)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x00, 0x9c, 0x4b, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got BanUserFromContentTypeRequest

	roundTrip(t, "BanUserFromContentTypeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("BanUserFromContentTypeRequest.TypeID differs")
	}
	if !(got.Pid.Equals(want.Pid)) {
		t.Fatalf("BanUserFromContentTypeRequest.Pid differs")
	}
	if !(got.Reason.Equals(want.Reason)) {
		t.Fatalf("BanUserFromContentTypeRequest.Reason differs")
	}
	if !(got.BanContents.Equals(want.BanContents)) {
		t.Fatalf("BanUserFromContentTypeRequest.BanContents differs")
	}
	if !(got.ExpireDate.Equals(want.ExpireDate)) {
		t.Fatalf("BanUserFromContentTypeRequest.ExpireDate differs")
	}
}

func TestBanUserFromContentTypeRequest_Header(t *testing.T) {
	want := BanUserFromContentTypeRequest{TypeID: types.NewUInt32(1014), Pid: types.NewUInt32(2023), Reason: types.NewString("s3"), BanContents: types.NewBool(false), ExpireDate: types.NewDateTime(5000092)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x00, 0x9c, 0x4b, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got BanUserFromContentTypeRequest

	roundTrip(t, "BanUserFromContentTypeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("BanUserFromContentTypeRequest.TypeID differs")
	}
	if !(got.Pid.Equals(want.Pid)) {
		t.Fatalf("BanUserFromContentTypeRequest.Pid differs")
	}
	if !(got.Reason.Equals(want.Reason)) {
		t.Fatalf("BanUserFromContentTypeRequest.Reason differs")
	}
	if !(got.BanContents.Equals(want.BanContents)) {
		t.Fatalf("BanUserFromContentTypeRequest.BanContents differs")
	}
	if !(got.ExpireDate.Equals(want.ExpireDate)) {
		t.Fatalf("BanUserFromContentTypeRequest.ExpireDate differs")
	}
}

func TestUnbanUserRequest_NoHeader(t *testing.T) {
	want := UnbanUserRequest{Pid: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got UnbanUserRequest

	roundTrip(t, "UnbanUserRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Pid.Equals(want.Pid)) {
		t.Fatalf("UnbanUserRequest.Pid differs")
	}
}

func TestUnbanUserRequest_Header(t *testing.T) {
	want := UnbanUserRequest{Pid: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got UnbanUserRequest

	roundTrip(t, "UnbanUserRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Pid.Equals(want.Pid)) {
		t.Fatalf("UnbanUserRequest.Pid differs")
	}
}

func TestUnbanUserFromContentTypeRequest_NoHeader(t *testing.T) {
	want := UnbanUserFromContentTypeRequest{TypeID: types.NewUInt32(1014), Pid: types.NewUInt32(2023)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00}

	var got UnbanUserFromContentTypeRequest

	roundTrip(t, "UnbanUserFromContentTypeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("UnbanUserFromContentTypeRequest.TypeID differs")
	}
	if !(got.Pid.Equals(want.Pid)) {
		t.Fatalf("UnbanUserFromContentTypeRequest.Pid differs")
	}
}

func TestUnbanUserFromContentTypeRequest_Header(t *testing.T) {
	want := UnbanUserFromContentTypeRequest{TypeID: types.NewUInt32(1014), Pid: types.NewUInt32(2023)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00}

	var got UnbanUserFromContentTypeRequest

	roundTrip(t, "UnbanUserFromContentTypeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("UnbanUserFromContentTypeRequest.TypeID differs")
	}
	if !(got.Pid.Equals(want.Pid)) {
		t.Fatalf("UnbanUserFromContentTypeRequest.Pid differs")
	}
}

func TestGetContentsToModerateWithThresholdRequest_NoHeader(t *testing.T) {
	want := GetContentsToModerateWithThresholdRequest{TypeID: types.NewUInt32(1014), Threshold: types.NewUInt32(2023), Offset: types.NewUInt32(3032), Size: types.NewUInt32(4041)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got GetContentsToModerateWithThresholdRequest

	roundTrip(t, "GetContentsToModerateWithThresholdRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("GetContentsToModerateWithThresholdRequest.TypeID differs")
	}
	if !(got.Threshold.Equals(want.Threshold)) {
		t.Fatalf("GetContentsToModerateWithThresholdRequest.Threshold differs")
	}
	if !(got.Offset.Equals(want.Offset)) {
		t.Fatalf("GetContentsToModerateWithThresholdRequest.Offset differs")
	}
	if !(got.Size.Equals(want.Size)) {
		t.Fatalf("GetContentsToModerateWithThresholdRequest.Size differs")
	}
}

func TestGetContentsToModerateWithThresholdRequest_Header(t *testing.T) {
	want := GetContentsToModerateWithThresholdRequest{TypeID: types.NewUInt32(1014), Threshold: types.NewUInt32(2023), Offset: types.NewUInt32(3032), Size: types.NewUInt32(4041)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got GetContentsToModerateWithThresholdRequest

	roundTrip(t, "GetContentsToModerateWithThresholdRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("GetContentsToModerateWithThresholdRequest.TypeID differs")
	}
	if !(got.Threshold.Equals(want.Threshold)) {
		t.Fatalf("GetContentsToModerateWithThresholdRequest.Threshold differs")
	}
	if !(got.Offset.Equals(want.Offset)) {
		t.Fatalf("GetContentsToModerateWithThresholdRequest.Offset differs")
	}
	if !(got.Size.Equals(want.Size)) {
		t.Fatalf("GetContentsToModerateWithThresholdRequest.Size differs")
	}
}

func TestGetContentsToModerateWithThresholdResponse_NoHeader(t *testing.T) {
	want := GetContentsToModerateWithThresholdResponse{Contents: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}, TotalResults: types.NewUInt32(26239)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x7f, 0x66, 0x00, 0x00}

	var got GetContentsToModerateWithThresholdResponse

	roundTrip(t, "GetContentsToModerateWithThresholdResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("GetContentsToModerateWithThresholdResponse.Contents differs")
	}
	if !(got.TotalResults.Equals(want.TotalResults)) {
		t.Fatalf("GetContentsToModerateWithThresholdResponse.TotalResults differs")
	}
}

func TestGetContentsToModerateWithThresholdResponse_Header(t *testing.T) {
	want := GetContentsToModerateWithThresholdResponse{Contents: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}, TotalResults: types.NewUInt32(26239)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x7f, 0x66, 0x00, 0x00}

	var got GetContentsToModerateWithThresholdResponse

	roundTrip(t, "GetContentsToModerateWithThresholdResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("GetContentsToModerateWithThresholdResponse.Contents differs")
	}
	if !(got.TotalResults.Equals(want.TotalResults)) {
		t.Fatalf("GetContentsToModerateWithThresholdResponse.TotalResults differs")
	}
}

func TestUpdateMetaDataRequest_NoHeader(t *testing.T) {
	want := UpdateMetaDataRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}(), Properties: types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(6059)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v7")}
		return v
	}(), func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(9086)
		v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
		return v
	}()}}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x37, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got UpdateMetaDataRequest

	roundTrip(t, "UpdateMetaDataRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UpdateMetaDataRequest.ContentKey differs")
	}
	if !(got.Properties.Equals(want.Properties)) {
		t.Fatalf("UpdateMetaDataRequest.Properties differs")
	}
}

func TestUpdateMetaDataRequest_Header(t *testing.T) {
	want := UpdateMetaDataRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}(), Properties: types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(6059)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v7")}
		return v
	}(), func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(9086)
		v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
		return v
	}()}}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got UpdateMetaDataRequest

	roundTrip(t, "UpdateMetaDataRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UpdateMetaDataRequest.ContentKey differs")
	}
	if !(got.Properties.Equals(want.Properties)) {
		t.Fatalf("UpdateMetaDataRequest.Properties differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodGetContentsToModerate != 1 {
		t.Fatalf("MethodGetContentsToModerate = %d, want 1", MethodGetContentsToModerate)
	}
	if MethodFlagContentAsVerified != 2 {
		t.Fatalf("MethodFlagContentAsVerified = %d, want 2", MethodFlagContentAsVerified)
	}
	if MethodBanContent != 3 {
		t.Fatalf("MethodBanContent = %d, want 3", MethodBanContent)
	}
	if MethodBanUser != 4 {
		t.Fatalf("MethodBanUser = %d, want 4", MethodBanUser)
	}
	if MethodBanUserFromContentType != 5 {
		t.Fatalf("MethodBanUserFromContentType = %d, want 5", MethodBanUserFromContentType)
	}
	if MethodUnbanUser != 6 {
		t.Fatalf("MethodUnbanUser = %d, want 6", MethodUnbanUser)
	}
	if MethodUnbanUserFromContentType != 7 {
		t.Fatalf("MethodUnbanUserFromContentType = %d, want 7", MethodUnbanUserFromContentType)
	}
	if MethodGetContentsToModerateWithThreshold != 8 {
		t.Fatalf("MethodGetContentsToModerateWithThreshold = %d, want 8", MethodGetContentsToModerateWithThreshold)
	}
	if MethodUpdateMetaData != 9 {
		t.Fatalf("MethodUpdateMetaData = %d, want 9", MethodUpdateMetaData)
	}
}
