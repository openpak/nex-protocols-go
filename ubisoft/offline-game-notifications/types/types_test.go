package types

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/notifications/types"
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

func TestTimedNotification_NoHeader(t *testing.T) {
	want := func() TimedNotification {
		v := NewTimedNotification()
		v.Timestamp = types.NewDateTime(2000083)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(33)
			v.Type = 39
			v.Param1 = types.NewUInt64(51)
			v.Param2 = types.NewUInt64(57)
			v.StrParam = types.NewString("n3")
			return v
		}()
		return v
	}()

	data := []byte{0xd3, 0x84, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x39, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x33, 0x00}

	var got TimedNotification

	roundTrip(t, "TimedNotification", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded TimedNotification differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestTimedNotification_Header(t *testing.T) {
	want := func() TimedNotification {
		v := NewTimedNotification()
		v.Timestamp = types.NewDateTime(2000083)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(33)
			v.Type = 39
			v.Param1 = types.NewUInt64(51)
			v.Param2 = types.NewUInt64(57)
			v.StrParam = types.NewString("n3")
			return v
		}()
		return v
	}()

	data := []byte{0x00, 0x22, 0x00, 0x00, 0x00, 0xd3, 0x84, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x39, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x33, 0x00}

	var got TimedNotification

	roundTrip(t, "TimedNotification", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded TimedNotification differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}
