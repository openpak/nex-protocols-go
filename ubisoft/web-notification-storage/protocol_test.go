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

func TestRegisterUserRequest_NoHeader(t *testing.T) {
	want := RegisterUserRequest{}

	data := []byte{}

	var got RegisterUserRequest

	roundTrip(t, "RegisterUserRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

func TestPollNotificationsRequest_NoHeader(t *testing.T) {
	want := PollNotificationsRequest{}

	data := []byte{}

	var got PollNotificationsRequest

	roundTrip(t, "PollNotificationsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

func TestPollNotificationsResponse_NoHeader(t *testing.T) {
	want := PollNotificationsResponse{ListNotifications: types.NewString("s1"), NbNotifications: types.NewInt32(-63)}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0xc1, 0xff, 0xff, 0xff}

	var got PollNotificationsResponse

	roundTrip(t, "PollNotificationsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ListNotifications.Equals(want.ListNotifications)) {
		t.Fatalf("PollNotificationsResponse.ListNotifications differs")
	}
	if !(got.NbNotifications.Equals(want.NbNotifications)) {
		t.Fatalf("PollNotificationsResponse.NbNotifications differs")
	}
}

func TestPollNotificationsResponse_Header(t *testing.T) {
	want := PollNotificationsResponse{ListNotifications: types.NewString("s1"), NbNotifications: types.NewInt32(-63)}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0xc1, 0xff, 0xff, 0xff}

	var got PollNotificationsResponse

	roundTrip(t, "PollNotificationsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ListNotifications.Equals(want.ListNotifications)) {
		t.Fatalf("PollNotificationsResponse.ListNotifications differs")
	}
	if !(got.NbNotifications.Equals(want.NbNotifications)) {
		t.Fatalf("PollNotificationsResponse.NbNotifications differs")
	}
}

func TestUnregisterUserRequest_NoHeader(t *testing.T) {
	want := UnregisterUserRequest{}

	data := []byte{}

	var got UnregisterUserRequest

	roundTrip(t, "UnregisterUserRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodRegisterUser != 1 {
		t.Fatalf("MethodRegisterUser = %d, want 1", MethodRegisterUser)
	}
	if MethodPollNotifications != 2 {
		t.Fatalf("MethodPollNotifications = %d, want 2", MethodPollNotifications)
	}
	if MethodUnregisterUser != 3 {
		t.Fatalf("MethodUnregisterUser = %d, want 3", MethodUnregisterUser)
	}
}
