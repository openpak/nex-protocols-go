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

func TestTitleContent_NoHeader(t *testing.T) {
	want := func() TitleContent {
		v := NewTitleContent()
		v.Id = types.NewUInt32(2023)
		v.Name = types.NewString("s3")
		v.Description = types.NewString("s4")
		v.Size = types.NewInt64(-20000000036)
		v.ContentType = types.NewString("s6")
		v.Flags = types.NewUInt32(7068)
		v.ExtendedFlags = types.NewUInt32(8077)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0xdc, 0x37, 0xe8, 0x57, 0xfb, 0xff, 0xff, 0xff, 0x03, 0x00, 0x73, 0x36, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00}

	var got TitleContent

	roundTrip(t, "TitleContent", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded TitleContent differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestTitleContent_Header(t *testing.T) {
	want := func() TitleContent {
		v := NewTitleContent()
		v.Id = types.NewUInt32(2023)
		v.Name = types.NewString("s3")
		v.Description = types.NewString("s4")
		v.Size = types.NewInt64(-20000000036)
		v.ContentType = types.NewString("s6")
		v.Flags = types.NewUInt32(7068)
		v.ExtendedFlags = types.NewUInt32(8077)
		return v
	}()

	data := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x23, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0xdc, 0x37, 0xe8, 0x57, 0xfb, 0xff, 0xff, 0xff, 0x03, 0x00, 0x73, 0x36, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00}

	var got TitleContent

	roundTrip(t, "TitleContent", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded TitleContent differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}
