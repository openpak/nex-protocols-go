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

func TestContentTypeExtensionField_NoHeader(t *testing.T) {
	want := func() ContentTypeExtensionField {
		v := NewContentTypeExtensionField()
		v.FieldName = types.NewString("s2")
		v.FieldType = types.NewString("s3")
		return v
	}()

	data := []byte{0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00}

	var got ContentTypeExtensionField

	roundTrip(t, "ContentTypeExtensionField", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded ContentTypeExtensionField differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestContentTypeExtensionField_Header(t *testing.T) {
	want := func() ContentTypeExtensionField {
		v := NewContentTypeExtensionField()
		v.FieldName = types.NewString("s2")
		v.FieldType = types.NewString("s3")
		return v
	}()

	data := []byte{0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00}

	var got ContentTypeExtensionField

	roundTrip(t, "ContentTypeExtensionField", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded ContentTypeExtensionField differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestContentTypeDescription_NoHeader(t *testing.T) {
	want := func() ContentTypeDescription {
		v := NewContentTypeDescription()
		v.ContentType = types.NewString("s2")
		v.ExtensionFields = types.List[ContentTypeExtensionField]{func() ContentTypeExtensionField {
			v := NewContentTypeExtensionField()
			v.FieldName = types.NewString("s5")
			v.FieldType = types.NewString("s6")
			return v
		}(), func() ContentTypeExtensionField {
			v := NewContentTypeExtensionField()
			v.FieldName = types.NewString("s8")
			v.FieldType = types.NewString("s9")
			return v
		}()}
		return v
	}()

	data := []byte{0x03, 0x00, 0x73, 0x32, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00}

	var got ContentTypeDescription

	roundTrip(t, "ContentTypeDescription", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded ContentTypeDescription differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestContentTypeDescription_Header(t *testing.T) {
	want := func() ContentTypeDescription {
		v := NewContentTypeDescription()
		v.ContentType = types.NewString("s2")
		v.ExtensionFields = types.List[ContentTypeExtensionField]{func() ContentTypeExtensionField {
			v := NewContentTypeExtensionField()
			v.FieldName = types.NewString("s5")
			v.FieldType = types.NewString("s6")
			return v
		}(), func() ContentTypeExtensionField {
			v := NewContentTypeExtensionField()
			v.FieldName = types.NewString("s8")
			v.FieldType = types.NewString("s9")
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x27, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00}

	var got ContentTypeDescription

	roundTrip(t, "ContentTypeDescription", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded ContentTypeDescription differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestContentDescription_NoHeader(t *testing.T) {
	want := func() ContentDescription {
		v := NewContentDescription()
		v.ContentDisplayName = types.NewString("s2")
		v.ContentDescription = types.NewString("s3")
		return v
	}()

	data := []byte{0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00}

	var got ContentDescription

	roundTrip(t, "ContentDescription", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded ContentDescription differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestContentDescription_Header(t *testing.T) {
	want := func() ContentDescription {
		v := NewContentDescription()
		v.ContentDisplayName = types.NewString("s2")
		v.ContentDescription = types.NewString("s3")
		return v
	}()

	data := []byte{0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00}

	var got ContentDescription

	roundTrip(t, "ContentDescription", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded ContentDescription differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestAdminTitleContent_NoHeader(t *testing.T) {
	want := func() AdminTitleContent {
		v := NewAdminTitleContent()
		v.Id = types.NewUInt32(2023)
		v.Name = types.NewString("s3")
		v.ContentDescription = func() ContentDescription {
			v := NewContentDescription()
			v.ContentDisplayName = types.NewString("s5")
			v.ContentDescription = types.NewString("s6")
			return v
		}()
		v.Flags = types.NewUInt32(7068)
		v.ExtendedFlags = types.NewUInt32(8077)
		v.DateAvailable = types.NewDateTime(9000104)
		v.DateExpired = types.NewDateTime(10000107)
		v.ContentType = types.NewString("s11")
		v.DownloadURL = types.NewString("s12")
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00}

	var got AdminTitleContent

	roundTrip(t, "AdminTitleContent", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded AdminTitleContent differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestAdminTitleContent_Header(t *testing.T) {
	want := func() AdminTitleContent {
		v := NewAdminTitleContent()
		v.Id = types.NewUInt32(2023)
		v.Name = types.NewString("s3")
		v.ContentDescription = func() ContentDescription {
			v := NewContentDescription()
			v.ContentDisplayName = types.NewString("s5")
			v.ContentDescription = types.NewString("s6")
			return v
		}()
		v.Flags = types.NewUInt32(7068)
		v.ExtendedFlags = types.NewUInt32(8077)
		v.DateAvailable = types.NewDateTime(9000104)
		v.DateExpired = types.NewDateTime(10000107)
		v.ContentType = types.NewString("s11")
		v.DownloadURL = types.NewString("s12")
		return v
	}()

	data := []byte{0x00, 0x3c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00}

	var got AdminTitleContent

	roundTrip(t, "AdminTitleContent", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded AdminTitleContent differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}
