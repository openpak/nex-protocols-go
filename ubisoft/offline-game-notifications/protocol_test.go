package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/notifications/types"
	offline_game_notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/offline-game-notifications/types"
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

func TestPollNotificationsRequest_NoHeader(t *testing.T) {
	want := PollNotificationsRequest{}

	data := []byte{}

	var got PollNotificationsRequest

	roundTrip(t, "PollNotificationsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

func TestPollNotificationsResponse_NoHeader(t *testing.T) {
	want := PollNotificationsResponse{ListNotifications: types.List[notifications_types.NotificationEvent]{func() notifications_types.NotificationEvent {
		v := notifications_types.NewNotificationEvent()
		v.PIDSource = types.NewPID(22)
		v.Type = 26
		v.Param1 = types.NewUInt64(34)
		v.Param2 = types.NewUInt64(38)
		v.StrParam = types.NewString("n2")
		return v
	}(), func() notifications_types.NotificationEvent {
		v := notifications_types.NewNotificationEvent()
		v.PIDSource = types.NewPID(33)
		v.Type = 39
		v.Param1 = types.NewUInt64(51)
		v.Param2 = types.NewUInt64(57)
		v.StrParam = types.NewString("n3")
		return v
	}()}, NbRemainingNotifs: types.NewUInt32(4041)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x26, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x32, 0x00, 0x21, 0x00, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x39, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x33, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got PollNotificationsResponse

	roundTrip(t, "PollNotificationsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ListNotifications.Equals(want.ListNotifications)) {
		t.Fatalf("PollNotificationsResponse.ListNotifications differs")
	}
	if !(got.NbRemainingNotifs.Equals(want.NbRemainingNotifs)) {
		t.Fatalf("PollNotificationsResponse.NbRemainingNotifs differs")
	}
}

func TestPollNotificationsResponse_Header(t *testing.T) {
	want := PollNotificationsResponse{ListNotifications: types.List[notifications_types.NotificationEvent]{func() notifications_types.NotificationEvent {
		v := notifications_types.NewNotificationEvent()
		v.PIDSource = types.NewPID(22)
		v.Type = 26
		v.Param1 = types.NewUInt64(34)
		v.Param2 = types.NewUInt64(38)
		v.StrParam = types.NewString("n2")
		return v
	}(), func() notifications_types.NotificationEvent {
		v := notifications_types.NewNotificationEvent()
		v.PIDSource = types.NewPID(33)
		v.Type = 39
		v.Param1 = types.NewUInt64(51)
		v.Param2 = types.NewUInt64(57)
		v.StrParam = types.NewString("n3")
		return v
	}()}, NbRemainingNotifs: types.NewUInt32(4041)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x26, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x32, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x39, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x33, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got PollNotificationsResponse

	roundTrip(t, "PollNotificationsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ListNotifications.Equals(want.ListNotifications)) {
		t.Fatalf("PollNotificationsResponse.ListNotifications differs")
	}
	if !(got.NbRemainingNotifs.Equals(want.NbRemainingNotifs)) {
		t.Fatalf("PollNotificationsResponse.NbRemainingNotifs differs")
	}
}

func TestPollSpecificOfflineNotificationsRequest_NoHeader(t *testing.T) {
	want := PollSpecificOfflineNotificationsRequest{Majortype: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got PollSpecificOfflineNotificationsRequest

	roundTrip(t, "PollSpecificOfflineNotificationsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Majortype.Equals(want.Majortype)) {
		t.Fatalf("PollSpecificOfflineNotificationsRequest.Majortype differs")
	}
}

func TestPollSpecificOfflineNotificationsRequest_Header(t *testing.T) {
	want := PollSpecificOfflineNotificationsRequest{Majortype: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got PollSpecificOfflineNotificationsRequest

	roundTrip(t, "PollSpecificOfflineNotificationsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Majortype.Equals(want.Majortype)) {
		t.Fatalf("PollSpecificOfflineNotificationsRequest.Majortype differs")
	}
}

func TestPollSpecificOfflineNotificationsResponse_NoHeader(t *testing.T) {
	want := PollSpecificOfflineNotificationsResponse{ListTimedNotification: types.List[offline_game_notifications_types.TimedNotification]{func() offline_game_notifications_types.TimedNotification {
		v := offline_game_notifications_types.NewTimedNotification()
		v.Timestamp = types.NewDateTime(3000086)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(44)
			v.Type = 52
			v.Param1 = types.NewUInt64(68)
			v.Param2 = types.NewUInt64(76)
			v.StrParam = types.NewString("n4")
			return v
		}()
		return v
	}(), func() offline_game_notifications_types.TimedNotification {
		v := offline_game_notifications_types.NewTimedNotification()
		v.Timestamp = types.NewDateTime(6000095)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(77)
			v.Type = 91
			v.Param1 = types.NewUInt64(119)
			v.Param2 = types.NewUInt64(133)
			v.StrParam = types.NewString("n7")
			return v
		}()
		return v
	}()}, Ret: types.NewUInt32(8077)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x16, 0xc7, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x34, 0x00, 0xdf, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x00, 0x00, 0x5b, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x85, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00}

	var got PollSpecificOfflineNotificationsResponse

	roundTrip(t, "PollSpecificOfflineNotificationsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ListTimedNotification.Equals(want.ListTimedNotification)) {
		t.Fatalf("PollSpecificOfflineNotificationsResponse.ListTimedNotification differs")
	}
	if !(got.Ret.Equals(want.Ret)) {
		t.Fatalf("PollSpecificOfflineNotificationsResponse.Ret differs")
	}
}

func TestPollSpecificOfflineNotificationsResponse_Header(t *testing.T) {
	want := PollSpecificOfflineNotificationsResponse{ListTimedNotification: types.List[offline_game_notifications_types.TimedNotification]{func() offline_game_notifications_types.TimedNotification {
		v := offline_game_notifications_types.NewTimedNotification()
		v.Timestamp = types.NewDateTime(3000086)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(44)
			v.Type = 52
			v.Param1 = types.NewUInt64(68)
			v.Param2 = types.NewUInt64(76)
			v.StrParam = types.NewString("n4")
			return v
		}()
		return v
	}(), func() offline_game_notifications_types.TimedNotification {
		v := offline_game_notifications_types.NewTimedNotification()
		v.Timestamp = types.NewDateTime(6000095)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(77)
			v.Type = 91
			v.Param1 = types.NewUInt64(119)
			v.Param2 = types.NewUInt64(133)
			v.StrParam = types.NewString("n7")
			return v
		}()
		return v
	}()}, Ret: types.NewUInt32(8077)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x16, 0xc7, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x34, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0xdf, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x00, 0x00, 0x5b, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x85, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00}

	var got PollSpecificOfflineNotificationsResponse

	roundTrip(t, "PollSpecificOfflineNotificationsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ListTimedNotification.Equals(want.ListTimedNotification)) {
		t.Fatalf("PollSpecificOfflineNotificationsResponse.ListTimedNotification differs")
	}
	if !(got.Ret.Equals(want.Ret)) {
		t.Fatalf("PollSpecificOfflineNotificationsResponse.Ret differs")
	}
}

func TestPollAnyOfflineNotificationsRequest_NoHeader(t *testing.T) {
	want := PollAnyOfflineNotificationsRequest{}

	data := []byte{}

	var got PollAnyOfflineNotificationsRequest

	roundTrip(t, "PollAnyOfflineNotificationsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

func TestPollAnyOfflineNotificationsResponse_NoHeader(t *testing.T) {
	want := PollAnyOfflineNotificationsResponse{ListTimedNotification: types.List[offline_game_notifications_types.TimedNotification]{func() offline_game_notifications_types.TimedNotification {
		v := offline_game_notifications_types.NewTimedNotification()
		v.Timestamp = types.NewDateTime(3000086)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(44)
			v.Type = 52
			v.Param1 = types.NewUInt64(68)
			v.Param2 = types.NewUInt64(76)
			v.StrParam = types.NewString("n4")
			return v
		}()
		return v
	}(), func() offline_game_notifications_types.TimedNotification {
		v := offline_game_notifications_types.NewTimedNotification()
		v.Timestamp = types.NewDateTime(6000095)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(77)
			v.Type = 91
			v.Param1 = types.NewUInt64(119)
			v.Param2 = types.NewUInt64(133)
			v.StrParam = types.NewString("n7")
			return v
		}()
		return v
	}()}, NbRemainingNotifs: types.NewUInt32(8077)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x16, 0xc7, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x34, 0x00, 0xdf, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x00, 0x00, 0x5b, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x85, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00}

	var got PollAnyOfflineNotificationsResponse

	roundTrip(t, "PollAnyOfflineNotificationsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ListTimedNotification.Equals(want.ListTimedNotification)) {
		t.Fatalf("PollAnyOfflineNotificationsResponse.ListTimedNotification differs")
	}
	if !(got.NbRemainingNotifs.Equals(want.NbRemainingNotifs)) {
		t.Fatalf("PollAnyOfflineNotificationsResponse.NbRemainingNotifs differs")
	}
}

func TestPollAnyOfflineNotificationsResponse_Header(t *testing.T) {
	want := PollAnyOfflineNotificationsResponse{ListTimedNotification: types.List[offline_game_notifications_types.TimedNotification]{func() offline_game_notifications_types.TimedNotification {
		v := offline_game_notifications_types.NewTimedNotification()
		v.Timestamp = types.NewDateTime(3000086)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(44)
			v.Type = 52
			v.Param1 = types.NewUInt64(68)
			v.Param2 = types.NewUInt64(76)
			v.StrParam = types.NewString("n4")
			return v
		}()
		return v
	}(), func() offline_game_notifications_types.TimedNotification {
		v := offline_game_notifications_types.NewTimedNotification()
		v.Timestamp = types.NewDateTime(6000095)
		v.Notification = func() notifications_types.NotificationEvent {
			v := notifications_types.NewNotificationEvent()
			v.PIDSource = types.NewPID(77)
			v.Type = 91
			v.Param1 = types.NewUInt64(119)
			v.Param2 = types.NewUInt64(133)
			v.StrParam = types.NewString("n7")
			return v
		}()
		return v
	}()}, NbRemainingNotifs: types.NewUInt32(8077)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x16, 0xc7, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x34, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0xdf, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x00, 0x00, 0x5b, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x85, 0x00, 0x00, 0x00, 0x03, 0x00, 0x6e, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00}

	var got PollAnyOfflineNotificationsResponse

	roundTrip(t, "PollAnyOfflineNotificationsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ListTimedNotification.Equals(want.ListTimedNotification)) {
		t.Fatalf("PollAnyOfflineNotificationsResponse.ListTimedNotification differs")
	}
	if !(got.NbRemainingNotifs.Equals(want.NbRemainingNotifs)) {
		t.Fatalf("PollAnyOfflineNotificationsResponse.NbRemainingNotifs differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodPollNotifications != 1 {
		t.Fatalf("MethodPollNotifications = %d, want 1", MethodPollNotifications)
	}
	if MethodPollSpecificOfflineNotifications != 2 {
		t.Fatalf("MethodPollSpecificOfflineNotifications = %d, want 2", MethodPollSpecificOfflineNotifications)
	}
	if MethodPollAnyOfflineNotifications != 3 {
		t.Fatalf("MethodPollAnyOfflineNotifications = %d, want 3", MethodPollAnyOfflineNotifications)
	}
	if ProtocolID != 71 {
		t.Fatalf("ProtocolID = %d, want 71", ProtocolID)
	}
}
