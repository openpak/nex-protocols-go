package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	tracking_3_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/tracking-3/types"
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

func TestSendTagRequest_NoHeader(t *testing.T) {
	want := SendTagRequest{TrackingID: types.NewUInt32(1014), Tag: types.NewString("s2"), Attributes: types.NewString("s3"), DeltaTime: types.NewUInt32(4041)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got SendTagRequest

	roundTrip(t, "SendTagRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TrackingID.Equals(want.TrackingID)) {
		t.Fatalf("SendTagRequest.TrackingID differs")
	}
	if !(got.Tag.Equals(want.Tag)) {
		t.Fatalf("SendTagRequest.Tag differs")
	}
	if !(got.Attributes.Equals(want.Attributes)) {
		t.Fatalf("SendTagRequest.Attributes differs")
	}
	if !(got.DeltaTime.Equals(want.DeltaTime)) {
		t.Fatalf("SendTagRequest.DeltaTime differs")
	}
}

func TestSendTagRequest_Header(t *testing.T) {
	want := SendTagRequest{TrackingID: types.NewUInt32(1014), Tag: types.NewString("s2"), Attributes: types.NewString("s3"), DeltaTime: types.NewUInt32(4041)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0xc9, 0x0f, 0x00, 0x00}

	var got SendTagRequest

	roundTrip(t, "SendTagRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TrackingID.Equals(want.TrackingID)) {
		t.Fatalf("SendTagRequest.TrackingID differs")
	}
	if !(got.Tag.Equals(want.Tag)) {
		t.Fatalf("SendTagRequest.Tag differs")
	}
	if !(got.Attributes.Equals(want.Attributes)) {
		t.Fatalf("SendTagRequest.Attributes differs")
	}
	if !(got.DeltaTime.Equals(want.DeltaTime)) {
		t.Fatalf("SendTagRequest.DeltaTime differs")
	}
}

func TestSendTagAndUpdateUserInfoRequest_NoHeader(t *testing.T) {
	want := SendTagAndUpdateUserInfoRequest{TrackingID: types.NewUInt32(1014), Tag: types.NewString("s2"), Attributes: types.NewString("s3"), DeltaTime: types.NewUInt32(4041), UserID: types.NewString("s5")}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got SendTagAndUpdateUserInfoRequest

	roundTrip(t, "SendTagAndUpdateUserInfoRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TrackingID.Equals(want.TrackingID)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.TrackingID differs")
	}
	if !(got.Tag.Equals(want.Tag)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.Tag differs")
	}
	if !(got.Attributes.Equals(want.Attributes)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.Attributes differs")
	}
	if !(got.DeltaTime.Equals(want.DeltaTime)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.DeltaTime differs")
	}
	if !(got.UserID.Equals(want.UserID)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.UserID differs")
	}
}

func TestSendTagAndUpdateUserInfoRequest_Header(t *testing.T) {
	want := SendTagAndUpdateUserInfoRequest{TrackingID: types.NewUInt32(1014), Tag: types.NewString("s2"), Attributes: types.NewString("s3"), DeltaTime: types.NewUInt32(4041), UserID: types.NewString("s5")}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got SendTagAndUpdateUserInfoRequest

	roundTrip(t, "SendTagAndUpdateUserInfoRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TrackingID.Equals(want.TrackingID)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.TrackingID differs")
	}
	if !(got.Tag.Equals(want.Tag)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.Tag differs")
	}
	if !(got.Attributes.Equals(want.Attributes)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.Attributes differs")
	}
	if !(got.DeltaTime.Equals(want.DeltaTime)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.DeltaTime differs")
	}
	if !(got.UserID.Equals(want.UserID)) {
		t.Fatalf("SendTagAndUpdateUserInfoRequest.UserID differs")
	}
}

func TestSendUserInfoRequest_NoHeader(t *testing.T) {
	want := SendUserInfoRequest{UserInfo: func() tracking_3_types.TrackingInformation {
		v := tracking_3_types.NewTrackingInformation()
		v.Ipn = types.NewUInt32(2023)
		v.UserID = types.NewString("s3")
		v.MachineID = types.NewString("s4")
		v.VisitorID = types.NewString("s5")
		v.UtsVersion = types.NewString("s6")
		return v
	}(), DeltaTime: types.NewUInt32(7068)}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x9c, 0x1b, 0x00, 0x00}

	var got SendUserInfoRequest

	roundTrip(t, "SendUserInfoRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.UserInfo.Equals(want.UserInfo)) {
		t.Fatalf("SendUserInfoRequest.UserInfo differs")
	}
	if !(got.DeltaTime.Equals(want.DeltaTime)) {
		t.Fatalf("SendUserInfoRequest.DeltaTime differs")
	}
}

func TestSendUserInfoRequest_Header(t *testing.T) {
	want := SendUserInfoRequest{UserInfo: func() tracking_3_types.TrackingInformation {
		v := tracking_3_types.NewTrackingInformation()
		v.Ipn = types.NewUInt32(2023)
		v.UserID = types.NewString("s3")
		v.MachineID = types.NewString("s4")
		v.VisitorID = types.NewString("s5")
		v.UtsVersion = types.NewString("s6")
		return v
	}(), DeltaTime: types.NewUInt32(7068)}

	data := []byte{0x00, 0x18, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x9c, 0x1b, 0x00, 0x00}

	var got SendUserInfoRequest

	roundTrip(t, "SendUserInfoRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.UserInfo.Equals(want.UserInfo)) {
		t.Fatalf("SendUserInfoRequest.UserInfo differs")
	}
	if !(got.DeltaTime.Equals(want.DeltaTime)) {
		t.Fatalf("SendUserInfoRequest.DeltaTime differs")
	}
}

func TestSendUserInfoResponse_NoHeader(t *testing.T) {
	want := SendUserInfoResponse{UserInfo: func() tracking_3_types.TrackingInformation {
		v := tracking_3_types.NewTrackingInformation()
		v.Ipn = types.NewUInt32(2023)
		v.UserID = types.NewString("s3")
		v.MachineID = types.NewString("s4")
		v.VisitorID = types.NewString("s5")
		v.UtsVersion = types.NewString("s6")
		return v
	}(), TrackingID: types.NewUInt32(7068)}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x9c, 0x1b, 0x00, 0x00}

	var got SendUserInfoResponse

	roundTrip(t, "SendUserInfoResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.UserInfo.Equals(want.UserInfo)) {
		t.Fatalf("SendUserInfoResponse.UserInfo differs")
	}
	if !(got.TrackingID.Equals(want.TrackingID)) {
		t.Fatalf("SendUserInfoResponse.TrackingID differs")
	}
}

func TestSendUserInfoResponse_Header(t *testing.T) {
	want := SendUserInfoResponse{UserInfo: func() tracking_3_types.TrackingInformation {
		v := tracking_3_types.NewTrackingInformation()
		v.Ipn = types.NewUInt32(2023)
		v.UserID = types.NewString("s3")
		v.MachineID = types.NewString("s4")
		v.VisitorID = types.NewString("s5")
		v.UtsVersion = types.NewString("s6")
		return v
	}(), TrackingID: types.NewUInt32(7068)}

	data := []byte{0x00, 0x18, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x9c, 0x1b, 0x00, 0x00}

	var got SendUserInfoResponse

	roundTrip(t, "SendUserInfoResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.UserInfo.Equals(want.UserInfo)) {
		t.Fatalf("SendUserInfoResponse.UserInfo differs")
	}
	if !(got.TrackingID.Equals(want.TrackingID)) {
		t.Fatalf("SendUserInfoResponse.TrackingID differs")
	}
}

func TestGetConfigurationRequest_NoHeader(t *testing.T) {
	want := GetConfigurationRequest{}

	data := []byte{}

	var got GetConfigurationRequest

	roundTrip(t, "GetConfigurationRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

}

func TestGetConfigurationResponse_NoHeader(t *testing.T) {
	want := GetConfigurationResponse{Tags: types.List[types.String]{types.NewString("s2"), types.NewString("s3")}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00}

	var got GetConfigurationResponse

	roundTrip(t, "GetConfigurationResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Tags.Equals(want.Tags)) {
		t.Fatalf("GetConfigurationResponse.Tags differs")
	}
}

func TestGetConfigurationResponse_Header(t *testing.T) {
	want := GetConfigurationResponse{Tags: types.List[types.String]{types.NewString("s2"), types.NewString("s3")}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00}

	var got GetConfigurationResponse

	roundTrip(t, "GetConfigurationResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Tags.Equals(want.Tags)) {
		t.Fatalf("GetConfigurationResponse.Tags differs")
	}
}

func TestSendTagsRequest_NoHeader(t *testing.T) {
	want := SendTagsRequest{TagData: types.List[tracking_3_types.TrackingTag]{func() tracking_3_types.TrackingTag {
		v := tracking_3_types.NewTrackingTag()
		v.TrackingID = types.NewUInt32(3032)
		v.Tag = types.NewString("s4")
		v.Attributes = types.NewString("s5")
		v.DeltaTime = types.NewUInt32(6059)
		v.NewUserId = types.NewString("s7")
		return v
	}(), func() tracking_3_types.TrackingTag {
		v := tracking_3_types.NewTrackingTag()
		v.TrackingID = types.NewUInt32(9086)
		v.Tag = types.NewString("s10")
		v.Attributes = types.NewString("s11")
		v.DeltaTime = types.NewUInt32(12113)
		v.NewUserId = types.NewString("s13")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x33, 0x00}

	var got SendTagsRequest

	roundTrip(t, "SendTagsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TagData.Equals(want.TagData)) {
		t.Fatalf("SendTagsRequest.TagData differs")
	}
}

func TestSendTagsRequest_Header(t *testing.T) {
	want := SendTagsRequest{TagData: types.List[tracking_3_types.TrackingTag]{func() tracking_3_types.TrackingTag {
		v := tracking_3_types.NewTrackingTag()
		v.TrackingID = types.NewUInt32(3032)
		v.Tag = types.NewString("s4")
		v.Attributes = types.NewString("s5")
		v.DeltaTime = types.NewUInt32(6059)
		v.NewUserId = types.NewString("s7")
		return v
	}(), func() tracking_3_types.TrackingTag {
		v := tracking_3_types.NewTrackingTag()
		v.TrackingID = types.NewUInt32(9086)
		v.Tag = types.NewString("s10")
		v.Attributes = types.NewString("s11")
		v.DeltaTime = types.NewUInt32(12113)
		v.NewUserId = types.NewString("s13")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x33, 0x00}

	var got SendTagsRequest

	roundTrip(t, "SendTagsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TagData.Equals(want.TagData)) {
		t.Fatalf("SendTagsRequest.TagData differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodSendTag != 1 {
		t.Fatalf("MethodSendTag = %d, want 1", MethodSendTag)
	}
	if MethodSendTagAndUpdateUserInfo != 2 {
		t.Fatalf("MethodSendTagAndUpdateUserInfo = %d, want 2", MethodSendTagAndUpdateUserInfo)
	}
	if MethodSendUserInfo != 3 {
		t.Fatalf("MethodSendUserInfo = %d, want 3", MethodSendUserInfo)
	}
	if MethodGetConfiguration != 4 {
		t.Fatalf("MethodGetConfiguration = %d, want 4", MethodGetConfiguration)
	}
	if MethodSendTags != 5 {
		t.Fatalf("MethodSendTags = %d, want 5", MethodSendTags)
	}
	if ProtocolID != 36 {
		t.Fatalf("ProtocolID = %d, want 36", ProtocolID)
	}
}
