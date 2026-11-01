package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	title_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/title-storage/types"
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

func TestBrowseContentsRequest_NoHeader(t *testing.T) {
	want := BrowseContentsRequest{Flags: types.NewUInt32(1014), FlagsMaskType: types.NewUInt8(15), ExtendedFlags: types.NewUInt32(3032), ExtendedFlagsMaskType: types.NewUInt8(29)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x0f, 0xd8, 0x0b, 0x00, 0x00, 0x1d}

	var got BrowseContentsRequest

	roundTrip(t, "BrowseContentsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Flags.Equals(want.Flags)) {
		t.Fatalf("BrowseContentsRequest.Flags differs")
	}
	if !(got.FlagsMaskType.Equals(want.FlagsMaskType)) {
		t.Fatalf("BrowseContentsRequest.FlagsMaskType differs")
	}
	if !(got.ExtendedFlags.Equals(want.ExtendedFlags)) {
		t.Fatalf("BrowseContentsRequest.ExtendedFlags differs")
	}
	if !(got.ExtendedFlagsMaskType.Equals(want.ExtendedFlagsMaskType)) {
		t.Fatalf("BrowseContentsRequest.ExtendedFlagsMaskType differs")
	}
}

func TestBrowseContentsRequest_Header(t *testing.T) {
	want := BrowseContentsRequest{Flags: types.NewUInt32(1014), FlagsMaskType: types.NewUInt8(15), ExtendedFlags: types.NewUInt32(3032), ExtendedFlagsMaskType: types.NewUInt8(29)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x0f, 0xd8, 0x0b, 0x00, 0x00, 0x1d}

	var got BrowseContentsRequest

	roundTrip(t, "BrowseContentsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Flags.Equals(want.Flags)) {
		t.Fatalf("BrowseContentsRequest.Flags differs")
	}
	if !(got.FlagsMaskType.Equals(want.FlagsMaskType)) {
		t.Fatalf("BrowseContentsRequest.FlagsMaskType differs")
	}
	if !(got.ExtendedFlags.Equals(want.ExtendedFlags)) {
		t.Fatalf("BrowseContentsRequest.ExtendedFlags differs")
	}
	if !(got.ExtendedFlagsMaskType.Equals(want.ExtendedFlagsMaskType)) {
		t.Fatalf("BrowseContentsRequest.ExtendedFlagsMaskType differs")
	}
}

func TestBrowseContentsResponse_NoHeader(t *testing.T) {
	want := BrowseContentsResponse{Contents: types.List[title_storage_types.TitleContent]{func() title_storage_types.TitleContent {
		v := title_storage_types.NewTitleContent()
		v.Id = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.Description = types.NewString("s5")
		v.Size = types.NewInt64(-24000000043)
		v.ContentType = types.NewString("s7")
		v.Flags = types.NewUInt32(8077)
		v.ExtendedFlags = types.NewUInt32(9086)
		return v
	}(), func() title_storage_types.TitleContent {
		v := title_storage_types.NewTitleContent()
		v.Id = types.NewUInt32(11104)
		v.Name = types.NewString("s12")
		v.Description = types.NewString("s13")
		v.Size = types.NewInt64(-56000000099)
		v.ContentType = types.NewString("s15")
		v.Flags = types.NewUInt32(16149)
		v.ExtendedFlags = types.NewUInt32(17158)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0xd5, 0x0f, 0x7d, 0x69, 0xfa, 0xff, 0xff, 0xff, 0x03, 0x00, 0x73, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x04, 0x00, 0x73, 0x31, 0x33, 0x00, 0x9d, 0xcf, 0x23, 0xf6, 0xf2, 0xff, 0xff, 0xff, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00}

	var got BrowseContentsResponse

	roundTrip(t, "BrowseContentsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("BrowseContentsResponse.Contents differs")
	}
}

func TestBrowseContentsResponse_Header(t *testing.T) {
	want := BrowseContentsResponse{Contents: types.List[title_storage_types.TitleContent]{func() title_storage_types.TitleContent {
		v := title_storage_types.NewTitleContent()
		v.Id = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.Description = types.NewString("s5")
		v.Size = types.NewInt64(-24000000043)
		v.ContentType = types.NewString("s7")
		v.Flags = types.NewUInt32(8077)
		v.ExtendedFlags = types.NewUInt32(9086)
		return v
	}(), func() title_storage_types.TitleContent {
		v := title_storage_types.NewTitleContent()
		v.Id = types.NewUInt32(11104)
		v.Name = types.NewString("s12")
		v.Description = types.NewString("s13")
		v.Size = types.NewInt64(-56000000099)
		v.ContentType = types.NewString("s15")
		v.Flags = types.NewUInt32(16149)
		v.ExtendedFlags = types.NewUInt32(17158)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x23, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0xd5, 0x0f, 0x7d, 0x69, 0xfa, 0xff, 0xff, 0xff, 0x03, 0x00, 0x73, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x26, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x04, 0x00, 0x73, 0x31, 0x33, 0x00, 0x9d, 0xcf, 0x23, 0xf6, 0xf2, 0xff, 0xff, 0xff, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00}

	var got BrowseContentsResponse

	roundTrip(t, "BrowseContentsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("BrowseContentsResponse.Contents differs")
	}
}

func TestBrowseContentsByTypeRequest_NoHeader(t *testing.T) {
	want := BrowseContentsByTypeRequest{ContentType: types.NewString("s1"), Flags: types.NewUInt32(2023), FlagsMaskType: types.NewUInt8(22), ExtendedFlags: types.NewUInt32(4041), ExtendedFlagsMaskType: types.NewUInt8(36)}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x16, 0xc9, 0x0f, 0x00, 0x00, 0x24}

	var got BrowseContentsByTypeRequest

	roundTrip(t, "BrowseContentsByTypeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentType.Equals(want.ContentType)) {
		t.Fatalf("BrowseContentsByTypeRequest.ContentType differs")
	}
	if !(got.Flags.Equals(want.Flags)) {
		t.Fatalf("BrowseContentsByTypeRequest.Flags differs")
	}
	if !(got.FlagsMaskType.Equals(want.FlagsMaskType)) {
		t.Fatalf("BrowseContentsByTypeRequest.FlagsMaskType differs")
	}
	if !(got.ExtendedFlags.Equals(want.ExtendedFlags)) {
		t.Fatalf("BrowseContentsByTypeRequest.ExtendedFlags differs")
	}
	if !(got.ExtendedFlagsMaskType.Equals(want.ExtendedFlagsMaskType)) {
		t.Fatalf("BrowseContentsByTypeRequest.ExtendedFlagsMaskType differs")
	}
}

func TestBrowseContentsByTypeRequest_Header(t *testing.T) {
	want := BrowseContentsByTypeRequest{ContentType: types.NewString("s1"), Flags: types.NewUInt32(2023), FlagsMaskType: types.NewUInt8(22), ExtendedFlags: types.NewUInt32(4041), ExtendedFlagsMaskType: types.NewUInt8(36)}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x16, 0xc9, 0x0f, 0x00, 0x00, 0x24}

	var got BrowseContentsByTypeRequest

	roundTrip(t, "BrowseContentsByTypeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentType.Equals(want.ContentType)) {
		t.Fatalf("BrowseContentsByTypeRequest.ContentType differs")
	}
	if !(got.Flags.Equals(want.Flags)) {
		t.Fatalf("BrowseContentsByTypeRequest.Flags differs")
	}
	if !(got.FlagsMaskType.Equals(want.FlagsMaskType)) {
		t.Fatalf("BrowseContentsByTypeRequest.FlagsMaskType differs")
	}
	if !(got.ExtendedFlags.Equals(want.ExtendedFlags)) {
		t.Fatalf("BrowseContentsByTypeRequest.ExtendedFlags differs")
	}
	if !(got.ExtendedFlagsMaskType.Equals(want.ExtendedFlagsMaskType)) {
		t.Fatalf("BrowseContentsByTypeRequest.ExtendedFlagsMaskType differs")
	}
}

func TestBrowseContentsByTypeResponse_NoHeader(t *testing.T) {
	want := BrowseContentsByTypeResponse{Contents: types.List[types.DataHolder]{types.DataHolder{Object: types.NewData()}, types.DataHolder{Object: types.NewData()}}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got BrowseContentsByTypeResponse

	roundTrip(t, "BrowseContentsByTypeResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("BrowseContentsByTypeResponse.Contents differs")
	}
}

func TestBrowseContentsByTypeResponse_Header(t *testing.T) {
	want := BrowseContentsByTypeResponse{Contents: types.List[types.DataHolder]{types.DataHolder{Object: types.NewData()}, types.DataHolder{Object: types.NewData()}}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x09, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x09, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got BrowseContentsByTypeResponse

	roundTrip(t, "BrowseContentsByTypeResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("BrowseContentsByTypeResponse.Contents differs")
	}
}

func TestGetContentDetailsRequest_NoHeader(t *testing.T) {
	want := GetContentDetailsRequest{ContentID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetContentDetailsRequest

	roundTrip(t, "GetContentDetailsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentID.Equals(want.ContentID)) {
		t.Fatalf("GetContentDetailsRequest.ContentID differs")
	}
}

func TestGetContentDetailsRequest_Header(t *testing.T) {
	want := GetContentDetailsRequest{ContentID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetContentDetailsRequest

	roundTrip(t, "GetContentDetailsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentID.Equals(want.ContentID)) {
		t.Fatalf("GetContentDetailsRequest.ContentID differs")
	}
}

func TestGetContentDetailsResponse_NoHeader(t *testing.T) {
	want := GetContentDetailsResponse{Details: types.DataHolder{Object: types.NewData()}}

	data := []byte{0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GetContentDetailsResponse

	roundTrip(t, "GetContentDetailsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Details.Equals(want.Details)) {
		t.Fatalf("GetContentDetailsResponse.Details differs")
	}
}

func TestGetContentDetailsResponse_Header(t *testing.T) {
	want := GetContentDetailsResponse{Details: types.DataHolder{Object: types.NewData()}}

	data := []byte{0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x09, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GetContentDetailsResponse

	roundTrip(t, "GetContentDetailsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Details.Equals(want.Details)) {
		t.Fatalf("GetContentDetailsResponse.Details differs")
	}
}

func TestDownloadContentRequest_NoHeader(t *testing.T) {
	want := DownloadContentRequest{ContentID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got DownloadContentRequest

	roundTrip(t, "DownloadContentRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentID.Equals(want.ContentID)) {
		t.Fatalf("DownloadContentRequest.ContentID differs")
	}
}

func TestDownloadContentRequest_Header(t *testing.T) {
	want := DownloadContentRequest{ContentID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got DownloadContentRequest

	roundTrip(t, "DownloadContentRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentID.Equals(want.ContentID)) {
		t.Fatalf("DownloadContentRequest.ContentID differs")
	}
}

func TestDownloadContentResponse_NoHeader(t *testing.T) {
	want := DownloadContentResponse{Proto: types.NewString("s1"), Host: types.NewString("s2"), Path: types.NewString("s3")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00}

	var got DownloadContentResponse

	roundTrip(t, "DownloadContentResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Proto.Equals(want.Proto)) {
		t.Fatalf("DownloadContentResponse.Proto differs")
	}
	if !(got.Host.Equals(want.Host)) {
		t.Fatalf("DownloadContentResponse.Host differs")
	}
	if !(got.Path.Equals(want.Path)) {
		t.Fatalf("DownloadContentResponse.Path differs")
	}
}

func TestDownloadContentResponse_Header(t *testing.T) {
	want := DownloadContentResponse{Proto: types.NewString("s1"), Host: types.NewString("s2"), Path: types.NewString("s3")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00}

	var got DownloadContentResponse

	roundTrip(t, "DownloadContentResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Proto.Equals(want.Proto)) {
		t.Fatalf("DownloadContentResponse.Proto differs")
	}
	if !(got.Host.Equals(want.Host)) {
		t.Fatalf("DownloadContentResponse.Host differs")
	}
	if !(got.Path.Equals(want.Path)) {
		t.Fatalf("DownloadContentResponse.Path differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodBrowseContents != 1 {
		t.Fatalf("MethodBrowseContents = %d, want 1", MethodBrowseContents)
	}
	if MethodBrowseContentsByType != 2 {
		t.Fatalf("MethodBrowseContentsByType = %d, want 2", MethodBrowseContentsByType)
	}
	if MethodGetContentDetails != 3 {
		t.Fatalf("MethodGetContentDetails = %d, want 3", MethodGetContentDetails)
	}
	if MethodDownloadContent != 4 {
		t.Fatalf("MethodDownloadContent = %d, want 4", MethodDownloadContent)
	}
	if ProtocolID != 51 {
		t.Fatalf("ProtocolID = %d, want 51", ProtocolID)
	}
}
