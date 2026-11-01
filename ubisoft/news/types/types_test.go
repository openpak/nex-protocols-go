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

func TestNewsChannel_NoHeader(t *testing.T) {
	want := func() NewsChannel {
		v := NewNewsChannel()
		v.ID = types.NewUInt32(2023)
		v.OwnerPID = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.Description = types.NewString("s5")
		v.CreationTime = types.NewDateTime(6000095)
		v.ExpirationTime = types.NewDateTime(7000098)
		v.Type = types.NewString("s8")
		v.Locale = types.NewString("s9")
		v.Subscribable = types.NewBool(false)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0xdf, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x00}

	var got NewsChannel

	roundTrip(t, "NewsChannel", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded NewsChannel differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestNewsChannel_Header(t *testing.T) {
	want := func() NewsChannel {
		v := NewNewsChannel()
		v.ID = types.NewUInt32(2023)
		v.OwnerPID = types.NewUInt32(3032)
		v.Name = types.NewString("s4")
		v.Description = types.NewString("s5")
		v.CreationTime = types.NewDateTime(6000095)
		v.ExpirationTime = types.NewDateTime(7000098)
		v.Type = types.NewString("s8")
		v.Locale = types.NewString("s9")
		v.Subscribable = types.NewBool(false)
		return v
	}()

	data := []byte{0x00, 0x2d, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0xdf, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x00}

	var got NewsChannel

	roundTrip(t, "NewsChannel", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded NewsChannel differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestNewsHeader_NoHeader(t *testing.T) {
	want := func() NewsHeader {
		v := NewNewsHeader()
		v.ID = types.NewUInt32(2023)
		v.RecipientID = types.NewUInt32(3032)
		v.RecipientType = types.NewUInt32(4041)
		v.PublisherPID = types.NewUInt32(5050)
		v.PublisherName = types.NewString("s6")
		v.PublicationTime = types.NewDateTime(7000098)
		v.DisplayTime = types.NewDateTime(8000101)
		v.ExpirationTime = types.NewDateTime(9000104)
		v.Title = types.NewString("s10")
		v.Link = types.NewString("s11")
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00}

	var got NewsHeader

	roundTrip(t, "NewsHeader", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded NewsHeader differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestNewsHeader_Header(t *testing.T) {
	want := func() NewsHeader {
		v := NewNewsHeader()
		v.ID = types.NewUInt32(2023)
		v.RecipientID = types.NewUInt32(3032)
		v.RecipientType = types.NewUInt32(4041)
		v.PublisherPID = types.NewUInt32(5050)
		v.PublisherName = types.NewString("s6")
		v.PublicationTime = types.NewDateTime(7000098)
		v.DisplayTime = types.NewDateTime(8000101)
		v.ExpirationTime = types.NewDateTime(9000104)
		v.Title = types.NewString("s10")
		v.Link = types.NewString("s11")
		return v
	}()

	data := []byte{0x00, 0x39, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00}

	var got NewsHeader

	roundTrip(t, "NewsHeader", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded NewsHeader differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestNewsMessage_NoHeader(t *testing.T) {
	want := func() NewsMessage { v := NewNewsMessage(); v.Body = types.NewString("s2"); return v }()

	data := []byte{0x03, 0x00, 0x73, 0x32, 0x00}

	var got NewsMessage

	roundTrip(t, "NewsMessage", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded NewsMessage differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestNewsMessage_Header(t *testing.T) {
	want := func() NewsMessage { v := NewNewsMessage(); v.Body = types.NewString("s2"); return v }()

	data := []byte{0x00, 0x05, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00}

	var got NewsMessage

	roundTrip(t, "NewsMessage", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded NewsMessage differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestNewsRecipient_NoHeader(t *testing.T) {
	want := func() NewsRecipient {
		v := NewNewsRecipient()
		v.RecipientID = types.NewUInt32(2023)
		v.RecipientType = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got NewsRecipient

	roundTrip(t, "NewsRecipient", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded NewsRecipient differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestNewsRecipient_Header(t *testing.T) {
	want := func() NewsRecipient {
		v := NewNewsRecipient()
		v.RecipientID = types.NewUInt32(2023)
		v.RecipientType = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got NewsRecipient

	roundTrip(t, "NewsRecipient", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded NewsRecipient differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}
