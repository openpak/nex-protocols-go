package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	title_storage_admin_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/title-storage-admin/types"
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

func TestGetContentTypesDescriptionRequest_NoHeader(t *testing.T) {
	want := GetContentTypesDescriptionRequest{}

	data := []byte{}

	var got GetContentTypesDescriptionRequest

	roundTrip(t, "GetContentTypesDescriptionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

func TestGetContentTypesDescriptionResponse_NoHeader(t *testing.T) {
	want := GetContentTypesDescriptionResponse{ContentDescriptions: types.List[title_storage_admin_types.ContentTypeDescription]{func() title_storage_admin_types.ContentTypeDescription {
		v := title_storage_admin_types.NewContentTypeDescription()
		v.ContentType = types.NewString("s3")
		v.ExtensionFields = types.List[title_storage_admin_types.ContentTypeExtensionField]{func() title_storage_admin_types.ContentTypeExtensionField {
			v := title_storage_admin_types.NewContentTypeExtensionField()
			v.FieldName = types.NewString("s6")
			v.FieldType = types.NewString("s7")
			return v
		}(), func() title_storage_admin_types.ContentTypeExtensionField {
			v := title_storage_admin_types.NewContentTypeExtensionField()
			v.FieldName = types.NewString("s9")
			v.FieldType = types.NewString("s10")
			return v
		}()}
		return v
	}(), func() title_storage_admin_types.ContentTypeDescription {
		v := title_storage_admin_types.NewContentTypeDescription()
		v.ContentType = types.NewString("s12")
		v.ExtensionFields = types.List[title_storage_admin_types.ContentTypeExtensionField]{func() title_storage_admin_types.ContentTypeExtensionField {
			v := title_storage_admin_types.NewContentTypeExtensionField()
			v.FieldName = types.NewString("s15")
			v.FieldType = types.NewString("s16")
			return v
		}(), func() title_storage_admin_types.ContentTypeExtensionField {
			v := title_storage_admin_types.NewContentTypeExtensionField()
			v.FieldName = types.NewString("s18")
			v.FieldType = types.NewString("s19")
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x02, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0x04, 0x00, 0x73, 0x31, 0x38, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00}

	var got GetContentTypesDescriptionResponse

	roundTrip(t, "GetContentTypesDescriptionResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentDescriptions.Equals(want.ContentDescriptions)) {
		t.Fatalf("GetContentTypesDescriptionResponse.ContentDescriptions differs")
	}
}

func TestGetContentTypesDescriptionResponse_Header(t *testing.T) {
	want := GetContentTypesDescriptionResponse{ContentDescriptions: types.List[title_storage_admin_types.ContentTypeDescription]{func() title_storage_admin_types.ContentTypeDescription {
		v := title_storage_admin_types.NewContentTypeDescription()
		v.ContentType = types.NewString("s3")
		v.ExtensionFields = types.List[title_storage_admin_types.ContentTypeExtensionField]{func() title_storage_admin_types.ContentTypeExtensionField {
			v := title_storage_admin_types.NewContentTypeExtensionField()
			v.FieldName = types.NewString("s6")
			v.FieldType = types.NewString("s7")
			return v
		}(), func() title_storage_admin_types.ContentTypeExtensionField {
			v := title_storage_admin_types.NewContentTypeExtensionField()
			v.FieldName = types.NewString("s9")
			v.FieldType = types.NewString("s10")
			return v
		}()}
		return v
	}(), func() title_storage_admin_types.ContentTypeDescription {
		v := title_storage_admin_types.NewContentTypeDescription()
		v.ContentType = types.NewString("s12")
		v.ExtensionFields = types.List[title_storage_admin_types.ContentTypeExtensionField]{func() title_storage_admin_types.ContentTypeExtensionField {
			v := title_storage_admin_types.NewContentTypeExtensionField()
			v.FieldName = types.NewString("s15")
			v.FieldType = types.NewString("s16")
			return v
		}(), func() title_storage_admin_types.ContentTypeExtensionField {
			v := title_storage_admin_types.NewContentTypeExtensionField()
			v.FieldName = types.NewString("s18")
			v.FieldType = types.NewString("s19")
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x38, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00}

	var got GetContentTypesDescriptionResponse

	roundTrip(t, "GetContentTypesDescriptionResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentDescriptions.Equals(want.ContentDescriptions)) {
		t.Fatalf("GetContentTypesDescriptionResponse.ContentDescriptions differs")
	}
}

func TestGetContentTypeDescriptionRequest_NoHeader(t *testing.T) {
	want := GetContentTypeDescriptionRequest{ContentType: types.NewString("s1")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00}

	var got GetContentTypeDescriptionRequest

	roundTrip(t, "GetContentTypeDescriptionRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentType.Equals(want.ContentType)) {
		t.Fatalf("GetContentTypeDescriptionRequest.ContentType differs")
	}
}

func TestGetContentTypeDescriptionRequest_Header(t *testing.T) {
	want := GetContentTypeDescriptionRequest{ContentType: types.NewString("s1")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00}

	var got GetContentTypeDescriptionRequest

	roundTrip(t, "GetContentTypeDescriptionRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentType.Equals(want.ContentType)) {
		t.Fatalf("GetContentTypeDescriptionRequest.ContentType differs")
	}
}

func TestGetContentTypeDescriptionResponse_NoHeader(t *testing.T) {
	want := GetContentTypeDescriptionResponse{ExtensionFields: types.List[title_storage_admin_types.ContentTypeExtensionField]{func() title_storage_admin_types.ContentTypeExtensionField {
		v := title_storage_admin_types.NewContentTypeExtensionField()
		v.FieldName = types.NewString("s3")
		v.FieldType = types.NewString("s4")
		return v
	}(), func() title_storage_admin_types.ContentTypeExtensionField {
		v := title_storage_admin_types.NewContentTypeExtensionField()
		v.FieldName = types.NewString("s6")
		v.FieldType = types.NewString("s7")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00}

	var got GetContentTypeDescriptionResponse

	roundTrip(t, "GetContentTypeDescriptionResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ExtensionFields.Equals(want.ExtensionFields)) {
		t.Fatalf("GetContentTypeDescriptionResponse.ExtensionFields differs")
	}
}

func TestGetContentTypeDescriptionResponse_Header(t *testing.T) {
	want := GetContentTypeDescriptionResponse{ExtensionFields: types.List[title_storage_admin_types.ContentTypeExtensionField]{func() title_storage_admin_types.ContentTypeExtensionField {
		v := title_storage_admin_types.NewContentTypeExtensionField()
		v.FieldName = types.NewString("s3")
		v.FieldType = types.NewString("s4")
		return v
	}(), func() title_storage_admin_types.ContentTypeExtensionField {
		v := title_storage_admin_types.NewContentTypeExtensionField()
		v.FieldName = types.NewString("s6")
		v.FieldType = types.NewString("s7")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00}

	var got GetContentTypeDescriptionResponse

	roundTrip(t, "GetContentTypeDescriptionResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ExtensionFields.Equals(want.ExtensionFields)) {
		t.Fatalf("GetContentTypeDescriptionResponse.ExtensionFields differs")
	}
}

func TestBrowseContentsRequest_NoHeader(t *testing.T) {
	want := BrowseContentsRequest{}

	data := []byte{}

	var got BrowseContentsRequest

	roundTrip(t, "BrowseContentsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

func TestBrowseContentsResponse_NoHeader(t *testing.T) {
	want := BrowseContentsResponse{Contents: types.List[title_storage_admin_types.AdminTitleContent]{func() title_storage_admin_types.AdminTitleContent {
		v := title_storage_admin_types.NewAdminTitleContent()
		v.Id = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.ContentDescription = func() title_storage_admin_types.ContentDescription {
			v := title_storage_admin_types.NewContentDescription()
			v.ContentDisplayName = types.NewString("s6")
			v.ContentDescription = types.NewString("s7")
			return v
		}()
		v.Flags = types.NewUInt32(8077)
		v.ExtendedFlags = types.NewUInt32(9086)
		v.DateAvailable = types.NewDateTime(10000107)
		v.DateExpired = types.NewDateTime(11000110)
		v.ContentType = types.NewString("s12")
		v.DownloadURL = types.NewString("s13")
		return v
	}(), func() title_storage_admin_types.AdminTitleContent {
		v := title_storage_admin_types.NewAdminTitleContent()
		v.Id = types.NewUInt32(15140)
		v.Name = types.NewString("s16")
		v.ContentDescription = func() title_storage_admin_types.ContentDescription {
			v := title_storage_admin_types.NewContentDescription()
			v.ContentDisplayName = types.NewString("s18")
			v.ContentDescription = types.NewString("s19")
			return v
		}()
		v.Flags = types.NewUInt32(20185)
		v.ExtendedFlags = types.NewUInt32(21194)
		v.DateAvailable = types.NewDateTime(22000143)
		v.DateExpired = types.NewDateTime(23000146)
		v.ContentType = types.NewString("s24")
		v.DownloadURL = types.NewString("s25")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xd9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x04, 0x00, 0x73, 0x31, 0x33, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0x04, 0x00, 0x73, 0x31, 0x38, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x0f, 0xb2, 0x4f, 0x01, 0x00, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x34, 0x00, 0x04, 0x00, 0x73, 0x32, 0x35, 0x00}

	var got BrowseContentsResponse

	roundTrip(t, "BrowseContentsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("BrowseContentsResponse.Contents differs")
	}
}

func TestBrowseContentsResponse_Header(t *testing.T) {
	want := BrowseContentsResponse{Contents: types.List[title_storage_admin_types.AdminTitleContent]{func() title_storage_admin_types.AdminTitleContent {
		v := title_storage_admin_types.NewAdminTitleContent()
		v.Id = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.ContentDescription = func() title_storage_admin_types.ContentDescription {
			v := title_storage_admin_types.NewContentDescription()
			v.ContentDisplayName = types.NewString("s6")
			v.ContentDescription = types.NewString("s7")
			return v
		}()
		v.Flags = types.NewUInt32(8077)
		v.ExtendedFlags = types.NewUInt32(9086)
		v.DateAvailable = types.NewDateTime(10000107)
		v.DateExpired = types.NewDateTime(11000110)
		v.ContentType = types.NewString("s12")
		v.DownloadURL = types.NewString("s13")
		return v
	}(), func() title_storage_admin_types.AdminTitleContent {
		v := title_storage_admin_types.NewAdminTitleContent()
		v.Id = types.NewUInt32(15140)
		v.Name = types.NewString("s16")
		v.ContentDescription = func() title_storage_admin_types.ContentDescription {
			v := title_storage_admin_types.NewContentDescription()
			v.ContentDisplayName = types.NewString("s18")
			v.ContentDescription = types.NewString("s19")
			return v
		}()
		v.Flags = types.NewUInt32(20185)
		v.ExtendedFlags = types.NewUInt32(21194)
		v.DateAvailable = types.NewDateTime(22000143)
		v.DateExpired = types.NewDateTime(23000146)
		v.ContentType = types.NewString("s24")
		v.DownloadURL = types.NewString("s25")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xd9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x04, 0x00, 0x73, 0x31, 0x33, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x38, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x0f, 0xb2, 0x4f, 0x01, 0x00, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x34, 0x00, 0x04, 0x00, 0x73, 0x32, 0x35, 0x00}

	var got BrowseContentsResponse

	roundTrip(t, "BrowseContentsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Contents.Equals(want.Contents)) {
		t.Fatalf("BrowseContentsResponse.Contents differs")
	}
}

func TestGetContentDetailsRequest_NoHeader(t *testing.T) {
	want := GetContentDetailsRequest{ContentId: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetContentDetailsRequest

	roundTrip(t, "GetContentDetailsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("GetContentDetailsRequest.ContentId differs")
	}
}

func TestGetContentDetailsRequest_Header(t *testing.T) {
	want := GetContentDetailsRequest{ContentId: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetContentDetailsRequest

	roundTrip(t, "GetContentDetailsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("GetContentDetailsRequest.ContentId differs")
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

func TestCreateContentAndGetUploadInformationRequest_NoHeader(t *testing.T) {
	want := CreateContentAndGetUploadInformationRequest{Name: types.NewString("s1"), Description: types.NewString("s2"), Key: types.NewString("s3"), DateAvailable: types.NewDateTime(4000089), DateExpired: types.NewDateTime(5000092), Flags: types.NewUInt32(6059), ExtendedFlags: types.NewUInt32(7068), ContentType: types.NewString("s8"), Extension: types.DataHolder{Object: types.NewData()}, Mime: types.NewString("s10")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x59, 0x09, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9c, 0x4b, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00}

	var got CreateContentAndGetUploadInformationRequest

	roundTrip(t, "CreateContentAndGetUploadInformationRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Name.Equals(want.Name)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Name differs")
	}
	if !(got.Description.Equals(want.Description)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Description differs")
	}
	if !(got.Key.Equals(want.Key)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Key differs")
	}
	if !(got.DateAvailable.Equals(want.DateAvailable)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.DateAvailable differs")
	}
	if !(got.DateExpired.Equals(want.DateExpired)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.DateExpired differs")
	}
	if !(got.Flags.Equals(want.Flags)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Flags differs")
	}
	if !(got.ExtendedFlags.Equals(want.ExtendedFlags)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.ExtendedFlags differs")
	}
	if !(got.ContentType.Equals(want.ContentType)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.ContentType differs")
	}
	if !(got.Extension.Equals(want.Extension)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Extension differs")
	}
	if !(got.Mime.Equals(want.Mime)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Mime differs")
	}
}

func TestCreateContentAndGetUploadInformationRequest_Header(t *testing.T) {
	want := CreateContentAndGetUploadInformationRequest{Name: types.NewString("s1"), Description: types.NewString("s2"), Key: types.NewString("s3"), DateAvailable: types.NewDateTime(4000089), DateExpired: types.NewDateTime(5000092), Flags: types.NewUInt32(6059), ExtendedFlags: types.NewUInt32(7068), ContentType: types.NewString("s8"), Extension: types.DataHolder{Object: types.NewData()}, Mime: types.NewString("s10")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x59, 0x09, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9c, 0x4b, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x09, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00}

	var got CreateContentAndGetUploadInformationRequest

	roundTrip(t, "CreateContentAndGetUploadInformationRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Name.Equals(want.Name)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Name differs")
	}
	if !(got.Description.Equals(want.Description)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Description differs")
	}
	if !(got.Key.Equals(want.Key)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Key differs")
	}
	if !(got.DateAvailable.Equals(want.DateAvailable)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.DateAvailable differs")
	}
	if !(got.DateExpired.Equals(want.DateExpired)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.DateExpired differs")
	}
	if !(got.Flags.Equals(want.Flags)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Flags differs")
	}
	if !(got.ExtendedFlags.Equals(want.ExtendedFlags)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.ExtendedFlags differs")
	}
	if !(got.ContentType.Equals(want.ContentType)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.ContentType differs")
	}
	if !(got.Extension.Equals(want.Extension)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Extension differs")
	}
	if !(got.Mime.Equals(want.Mime)) {
		t.Fatalf("CreateContentAndGetUploadInformationRequest.Mime differs")
	}
}

func TestCreateContentAndGetUploadInformationResponse_NoHeader(t *testing.T) {
	want := CreateContentAndGetUploadInformationResponse{ContentId: types.NewUInt32(1014), UploadURL: types.NewString("s2"), UploadMimeType: types.NewString("s3"), UploadBody: types.Map[types.String, types.String]{types.NewString("s5"): types.NewString("s6")}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00}

	var got CreateContentAndGetUploadInformationResponse

	roundTrip(t, "CreateContentAndGetUploadInformationResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("CreateContentAndGetUploadInformationResponse.ContentId differs")
	}
	if !(got.UploadURL.Equals(want.UploadURL)) {
		t.Fatalf("CreateContentAndGetUploadInformationResponse.UploadURL differs")
	}
	if !(got.UploadMimeType.Equals(want.UploadMimeType)) {
		t.Fatalf("CreateContentAndGetUploadInformationResponse.UploadMimeType differs")
	}
	if !(got.UploadBody.Equals(want.UploadBody)) {
		t.Fatalf("CreateContentAndGetUploadInformationResponse.UploadBody differs")
	}
}

func TestCreateContentAndGetUploadInformationResponse_Header(t *testing.T) {
	want := CreateContentAndGetUploadInformationResponse{ContentId: types.NewUInt32(1014), UploadURL: types.NewString("s2"), UploadMimeType: types.NewString("s3"), UploadBody: types.Map[types.String, types.String]{types.NewString("s5"): types.NewString("s6")}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00}

	var got CreateContentAndGetUploadInformationResponse

	roundTrip(t, "CreateContentAndGetUploadInformationResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("CreateContentAndGetUploadInformationResponse.ContentId differs")
	}
	if !(got.UploadURL.Equals(want.UploadURL)) {
		t.Fatalf("CreateContentAndGetUploadInformationResponse.UploadURL differs")
	}
	if !(got.UploadMimeType.Equals(want.UploadMimeType)) {
		t.Fatalf("CreateContentAndGetUploadInformationResponse.UploadMimeType differs")
	}
	if !(got.UploadBody.Equals(want.UploadBody)) {
		t.Fatalf("CreateContentAndGetUploadInformationResponse.UploadBody differs")
	}
}

func TestRemoveContentRequest_NoHeader(t *testing.T) {
	want := RemoveContentRequest{ContentId: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got RemoveContentRequest

	roundTrip(t, "RemoveContentRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("RemoveContentRequest.ContentId differs")
	}
}

func TestRemoveContentRequest_Header(t *testing.T) {
	want := RemoveContentRequest{ContentId: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got RemoveContentRequest

	roundTrip(t, "RemoveContentRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("RemoveContentRequest.ContentId differs")
	}
}

func TestRemoveContentResponse_NoHeader(t *testing.T) {
	want := RemoveContentResponse{Retval: types.NewBool(true)}

	data := []byte{0x01}

	var got RemoveContentResponse

	roundTrip(t, "RemoveContentResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Retval.Equals(want.Retval)) {
		t.Fatalf("RemoveContentResponse.Retval differs")
	}
}

func TestRemoveContentResponse_Header(t *testing.T) {
	want := RemoveContentResponse{Retval: types.NewBool(true)}

	data := []byte{0x01}

	var got RemoveContentResponse

	roundTrip(t, "RemoveContentResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Retval.Equals(want.Retval)) {
		t.Fatalf("RemoveContentResponse.Retval differs")
	}
}

func TestUpdateContentMetaDataRequest_NoHeader(t *testing.T) {
	want := UpdateContentMetaDataRequest{ContentId: types.NewUInt32(1014), Name: types.NewString("s2"), Description: types.NewString("s3"), DateAvailable: types.NewDateTime(4000089), DateExpired: types.NewDateTime(5000092), Flags: types.NewUInt32(6059), ExtendedFlags: types.NewUInt32(7068), Extension: types.DataHolder{Object: types.NewData()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x59, 0x09, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9c, 0x4b, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got UpdateContentMetaDataRequest

	roundTrip(t, "UpdateContentMetaDataRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("UpdateContentMetaDataRequest.ContentId differs")
	}
	if !(got.Name.Equals(want.Name)) {
		t.Fatalf("UpdateContentMetaDataRequest.Name differs")
	}
	if !(got.Description.Equals(want.Description)) {
		t.Fatalf("UpdateContentMetaDataRequest.Description differs")
	}
	if !(got.DateAvailable.Equals(want.DateAvailable)) {
		t.Fatalf("UpdateContentMetaDataRequest.DateAvailable differs")
	}
	if !(got.DateExpired.Equals(want.DateExpired)) {
		t.Fatalf("UpdateContentMetaDataRequest.DateExpired differs")
	}
	if !(got.Flags.Equals(want.Flags)) {
		t.Fatalf("UpdateContentMetaDataRequest.Flags differs")
	}
	if !(got.ExtendedFlags.Equals(want.ExtendedFlags)) {
		t.Fatalf("UpdateContentMetaDataRequest.ExtendedFlags differs")
	}
	if !(got.Extension.Equals(want.Extension)) {
		t.Fatalf("UpdateContentMetaDataRequest.Extension differs")
	}
}

func TestUpdateContentMetaDataRequest_Header(t *testing.T) {
	want := UpdateContentMetaDataRequest{ContentId: types.NewUInt32(1014), Name: types.NewString("s2"), Description: types.NewString("s3"), DateAvailable: types.NewDateTime(4000089), DateExpired: types.NewDateTime(5000092), Flags: types.NewUInt32(6059), ExtendedFlags: types.NewUInt32(7068), Extension: types.DataHolder{Object: types.NewData()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x59, 0x09, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9c, 0x4b, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x05, 0x00, 0x44, 0x61, 0x74, 0x61, 0x00, 0x09, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got UpdateContentMetaDataRequest

	roundTrip(t, "UpdateContentMetaDataRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("UpdateContentMetaDataRequest.ContentId differs")
	}
	if !(got.Name.Equals(want.Name)) {
		t.Fatalf("UpdateContentMetaDataRequest.Name differs")
	}
	if !(got.Description.Equals(want.Description)) {
		t.Fatalf("UpdateContentMetaDataRequest.Description differs")
	}
	if !(got.DateAvailable.Equals(want.DateAvailable)) {
		t.Fatalf("UpdateContentMetaDataRequest.DateAvailable differs")
	}
	if !(got.DateExpired.Equals(want.DateExpired)) {
		t.Fatalf("UpdateContentMetaDataRequest.DateExpired differs")
	}
	if !(got.Flags.Equals(want.Flags)) {
		t.Fatalf("UpdateContentMetaDataRequest.Flags differs")
	}
	if !(got.ExtendedFlags.Equals(want.ExtendedFlags)) {
		t.Fatalf("UpdateContentMetaDataRequest.ExtendedFlags differs")
	}
	if !(got.Extension.Equals(want.Extension)) {
		t.Fatalf("UpdateContentMetaDataRequest.Extension differs")
	}
}

func TestUpdateContentMetaDataResponse_NoHeader(t *testing.T) {
	want := UpdateContentMetaDataResponse{Retval: types.NewBool(true)}

	data := []byte{0x01}

	var got UpdateContentMetaDataResponse

	roundTrip(t, "UpdateContentMetaDataResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Retval.Equals(want.Retval)) {
		t.Fatalf("UpdateContentMetaDataResponse.Retval differs")
	}
}

func TestUpdateContentMetaDataResponse_Header(t *testing.T) {
	want := UpdateContentMetaDataResponse{Retval: types.NewBool(true)}

	data := []byte{0x01}

	var got UpdateContentMetaDataResponse

	roundTrip(t, "UpdateContentMetaDataResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Retval.Equals(want.Retval)) {
		t.Fatalf("UpdateContentMetaDataResponse.Retval differs")
	}
}

func TestContentInsertionCompletedRequest_NoHeader(t *testing.T) {
	want := ContentInsertionCompletedRequest{ContentId: types.NewUInt32(1014), ContentSize: types.NewUInt32(2023)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00}

	var got ContentInsertionCompletedRequest

	roundTrip(t, "ContentInsertionCompletedRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("ContentInsertionCompletedRequest.ContentId differs")
	}
	if !(got.ContentSize.Equals(want.ContentSize)) {
		t.Fatalf("ContentInsertionCompletedRequest.ContentSize differs")
	}
}

func TestContentInsertionCompletedRequest_Header(t *testing.T) {
	want := ContentInsertionCompletedRequest{ContentId: types.NewUInt32(1014), ContentSize: types.NewUInt32(2023)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00}

	var got ContentInsertionCompletedRequest

	roundTrip(t, "ContentInsertionCompletedRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentId.Equals(want.ContentId)) {
		t.Fatalf("ContentInsertionCompletedRequest.ContentId differs")
	}
	if !(got.ContentSize.Equals(want.ContentSize)) {
		t.Fatalf("ContentInsertionCompletedRequest.ContentSize differs")
	}
}

func TestContentInsertionCompletedResponse_NoHeader(t *testing.T) {
	want := ContentInsertionCompletedResponse{Retval: types.NewBool(true)}

	data := []byte{0x01}

	var got ContentInsertionCompletedResponse

	roundTrip(t, "ContentInsertionCompletedResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Retval.Equals(want.Retval)) {
		t.Fatalf("ContentInsertionCompletedResponse.Retval differs")
	}
}

func TestContentInsertionCompletedResponse_Header(t *testing.T) {
	want := ContentInsertionCompletedResponse{Retval: types.NewBool(true)}

	data := []byte{0x01}

	var got ContentInsertionCompletedResponse

	roundTrip(t, "ContentInsertionCompletedResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Retval.Equals(want.Retval)) {
		t.Fatalf("ContentInsertionCompletedResponse.Retval differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodGetContentTypesDescription != 1 {
		t.Fatalf("MethodGetContentTypesDescription = %d, want 1", MethodGetContentTypesDescription)
	}
	if MethodGetContentTypeDescription != 2 {
		t.Fatalf("MethodGetContentTypeDescription = %d, want 2", MethodGetContentTypeDescription)
	}
	if MethodBrowseContents != 3 {
		t.Fatalf("MethodBrowseContents = %d, want 3", MethodBrowseContents)
	}
	if MethodGetContentDetails != 4 {
		t.Fatalf("MethodGetContentDetails = %d, want 4", MethodGetContentDetails)
	}
	if MethodCreateContentAndGetUploadInformation != 5 {
		t.Fatalf("MethodCreateContentAndGetUploadInformation = %d, want 5", MethodCreateContentAndGetUploadInformation)
	}
	if MethodRemoveContent != 6 {
		t.Fatalf("MethodRemoveContent = %d, want 6", MethodRemoveContent)
	}
	if MethodUpdateContentMetaData != 7 {
		t.Fatalf("MethodUpdateContentMetaData = %d, want 7", MethodUpdateContentMetaData)
	}
	if MethodContentInsertionCompleted != 8 {
		t.Fatalf("MethodContentInsertionCompleted = %d, want 8", MethodContentInsertionCompleted)
	}
}
