package protocol

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

func TestGetLocaleCodeRequest_NoHeader(t *testing.T) {
	want := GetLocaleCodeRequest{}

	data := []byte{}

	var got GetLocaleCodeRequest

	roundTrip(t, "GetLocaleCodeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

func TestGetLocaleCodeResponse_NoHeader(t *testing.T) {
	want := GetLocaleCodeResponse{LocalCode: types.NewString("s1")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00}

	var got GetLocaleCodeResponse

	roundTrip(t, "GetLocaleCodeResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.LocalCode.Equals(want.LocalCode)) {
		t.Fatalf("GetLocaleCodeResponse.LocalCode differs")
	}
}

func TestGetLocaleCodeResponse_Header(t *testing.T) {
	want := GetLocaleCodeResponse{LocalCode: types.NewString("s1")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00}

	var got GetLocaleCodeResponse

	roundTrip(t, "GetLocaleCodeResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.LocalCode.Equals(want.LocalCode)) {
		t.Fatalf("GetLocaleCodeResponse.LocalCode differs")
	}
}

func TestSetLocaleCodeRequest_NoHeader(t *testing.T) {
	want := SetLocaleCodeRequest{LocalCode: types.NewString("s1")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00}

	var got SetLocaleCodeRequest

	roundTrip(t, "SetLocaleCodeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.LocalCode.Equals(want.LocalCode)) {
		t.Fatalf("SetLocaleCodeRequest.LocalCode differs")
	}
}

func TestSetLocaleCodeRequest_Header(t *testing.T) {
	want := SetLocaleCodeRequest{LocalCode: types.NewString("s1")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00}

	var got SetLocaleCodeRequest

	roundTrip(t, "SetLocaleCodeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.LocalCode.Equals(want.LocalCode)) {
		t.Fatalf("SetLocaleCodeRequest.LocalCode differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodGetLocaleCode != 1 {
		t.Fatalf("MethodGetLocaleCode = %d, want 1", MethodGetLocaleCode)
	}
	if MethodSetLocaleCode != 2 {
		t.Fatalf("MethodSetLocaleCode = %d, want 2", MethodSetLocaleCode)
	}
	if ProtocolID != 39 {
		t.Fatalf("ProtocolID = %d, want 39", ProtocolID)
	}
}
