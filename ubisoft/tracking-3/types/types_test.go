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

func TestTrackingInformation_NoHeader(t *testing.T) {
	want := func() TrackingInformation {
		v := NewTrackingInformation()
		v.Ipn = types.NewUInt32(2023)
		v.UserID = types.NewString("s3")
		v.MachineID = types.NewString("s4")
		v.VisitorID = types.NewString("s5")
		v.UtsVersion = types.NewString("s6")
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00}

	var got TrackingInformation

	roundTrip(t, "TrackingInformation", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded TrackingInformation differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestTrackingInformation_Header(t *testing.T) {
	want := func() TrackingInformation {
		v := NewTrackingInformation()
		v.Ipn = types.NewUInt32(2023)
		v.UserID = types.NewString("s3")
		v.MachineID = types.NewString("s4")
		v.VisitorID = types.NewString("s5")
		v.UtsVersion = types.NewString("s6")
		return v
	}()

	data := []byte{0x00, 0x18, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00}

	var got TrackingInformation

	roundTrip(t, "TrackingInformation", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded TrackingInformation differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestTrackingTag_NoHeader(t *testing.T) {
	want := func() TrackingTag {
		v := NewTrackingTag()
		v.TrackingID = types.NewUInt32(2023)
		v.Tag = types.NewString("s3")
		v.Attributes = types.NewString("s4")
		v.DeltaTime = types.NewUInt32(5050)
		v.NewUserId = types.NewString("s6")
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00}

	var got TrackingTag

	roundTrip(t, "TrackingTag", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded TrackingTag differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestTrackingTag_Header(t *testing.T) {
	want := func() TrackingTag {
		v := NewTrackingTag()
		v.TrackingID = types.NewUInt32(2023)
		v.Tag = types.NewString("s3")
		v.Attributes = types.NewString("s4")
		v.DeltaTime = types.NewUInt32(5050)
		v.NewUserId = types.NewString("s6")
		return v
	}()

	data := []byte{0x00, 0x17, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0xba, 0x13, 0x00, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00}

	var got TrackingTag

	roundTrip(t, "TrackingTag", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded TrackingTag differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}
