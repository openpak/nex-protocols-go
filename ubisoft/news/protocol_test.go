package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
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

func TestGetChannelsRequest_NoHeader(t *testing.T) {
	want := GetChannelsRequest{ResultRange: types.ResultRange{Offset: types.NewUInt32(3), Length: types.NewUInt32(13)}}

	data := []byte{0x03, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00}

	var got GetChannelsRequest

	roundTrip(t, "GetChannelsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetChannelsRequest.ResultRange differs")
	}
}

func TestGetChannelsRequest_Header(t *testing.T) {
	want := GetChannelsRequest{ResultRange: types.ResultRange{Offset: types.NewUInt32(3), Length: types.NewUInt32(13)}}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00}

	var got GetChannelsRequest

	roundTrip(t, "GetChannelsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetChannelsRequest.ResultRange differs")
	}
}

func TestGetChannelsResponse_NoHeader(t *testing.T) {
	want := GetChannelsResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetChannelsResponse

	roundTrip(t, "GetChannelsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetChannelsResponse.Channels differs")
	}
}

func TestGetChannelsResponse_Header(t *testing.T) {
	want := GetChannelsResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x00, 0x31, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetChannelsResponse

	roundTrip(t, "GetChannelsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetChannelsResponse.Channels differs")
	}
}

func TestGetChannelsByTypesRequest_NoHeader(t *testing.T) {
	want := GetChannelsByTypesRequest{NewsChannelTypes: types.List[types.String]{types.NewString("s2"), types.NewString("s3")}, ResultRange: types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00}

	var got GetChannelsByTypesRequest

	roundTrip(t, "GetChannelsByTypesRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelTypes.Equals(want.NewsChannelTypes)) {
		t.Fatalf("GetChannelsByTypesRequest.NewsChannelTypes differs")
	}
	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetChannelsByTypesRequest.ResultRange differs")
	}
}

func TestGetChannelsByTypesRequest_Header(t *testing.T) {
	want := GetChannelsByTypesRequest{NewsChannelTypes: types.List[types.String]{types.NewString("s2"), types.NewString("s3")}, ResultRange: types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00}

	var got GetChannelsByTypesRequest

	roundTrip(t, "GetChannelsByTypesRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelTypes.Equals(want.NewsChannelTypes)) {
		t.Fatalf("GetChannelsByTypesRequest.NewsChannelTypes differs")
	}
	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetChannelsByTypesRequest.ResultRange differs")
	}
}

func TestGetChannelsByTypesResponse_NoHeader(t *testing.T) {
	want := GetChannelsByTypesResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetChannelsByTypesResponse

	roundTrip(t, "GetChannelsByTypesResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetChannelsByTypesResponse.Channels differs")
	}
}

func TestGetChannelsByTypesResponse_Header(t *testing.T) {
	want := GetChannelsByTypesResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x00, 0x31, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetChannelsByTypesResponse

	roundTrip(t, "GetChannelsByTypesResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetChannelsByTypesResponse.Channels differs")
	}
}

func TestGetSubscribableChannelsRequest_NoHeader(t *testing.T) {
	want := GetSubscribableChannelsRequest{ResultRange: types.ResultRange{Offset: types.NewUInt32(3), Length: types.NewUInt32(13)}}

	data := []byte{0x03, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00}

	var got GetSubscribableChannelsRequest

	roundTrip(t, "GetSubscribableChannelsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetSubscribableChannelsRequest.ResultRange differs")
	}
}

func TestGetSubscribableChannelsRequest_Header(t *testing.T) {
	want := GetSubscribableChannelsRequest{ResultRange: types.ResultRange{Offset: types.NewUInt32(3), Length: types.NewUInt32(13)}}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00}

	var got GetSubscribableChannelsRequest

	roundTrip(t, "GetSubscribableChannelsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetSubscribableChannelsRequest.ResultRange differs")
	}
}

func TestGetSubscribableChannelsResponse_NoHeader(t *testing.T) {
	want := GetSubscribableChannelsResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetSubscribableChannelsResponse

	roundTrip(t, "GetSubscribableChannelsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetSubscribableChannelsResponse.Channels differs")
	}
}

func TestGetSubscribableChannelsResponse_Header(t *testing.T) {
	want := GetSubscribableChannelsResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x00, 0x31, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetSubscribableChannelsResponse

	roundTrip(t, "GetSubscribableChannelsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetSubscribableChannelsResponse.Channels differs")
	}
}

func TestGetChannelsByIDsRequest_NoHeader(t *testing.T) {
	want := GetChannelsByIDsRequest{NewsChannelIDs: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetChannelsByIDsRequest

	roundTrip(t, "GetChannelsByIDsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelIDs.Equals(want.NewsChannelIDs)) {
		t.Fatalf("GetChannelsByIDsRequest.NewsChannelIDs differs")
	}
}

func TestGetChannelsByIDsRequest_Header(t *testing.T) {
	want := GetChannelsByIDsRequest{NewsChannelIDs: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetChannelsByIDsRequest

	roundTrip(t, "GetChannelsByIDsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelIDs.Equals(want.NewsChannelIDs)) {
		t.Fatalf("GetChannelsByIDsRequest.NewsChannelIDs differs")
	}
}

func TestGetChannelsByIDsResponse_NoHeader(t *testing.T) {
	want := GetChannelsByIDsResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetChannelsByIDsResponse

	roundTrip(t, "GetChannelsByIDsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetChannelsByIDsResponse.Channels differs")
	}
}

func TestGetChannelsByIDsResponse_Header(t *testing.T) {
	want := GetChannelsByIDsResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x00, 0x31, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetChannelsByIDsResponse

	roundTrip(t, "GetChannelsByIDsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetChannelsByIDsResponse.Channels differs")
	}
}

func TestGetSubscribedChannelsRequest_NoHeader(t *testing.T) {
	want := GetSubscribedChannelsRequest{ResultRange: types.ResultRange{Offset: types.NewUInt32(3), Length: types.NewUInt32(13)}}

	data := []byte{0x03, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00}

	var got GetSubscribedChannelsRequest

	roundTrip(t, "GetSubscribedChannelsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetSubscribedChannelsRequest.ResultRange differs")
	}
}

func TestGetSubscribedChannelsRequest_Header(t *testing.T) {
	want := GetSubscribedChannelsRequest{ResultRange: types.ResultRange{Offset: types.NewUInt32(3), Length: types.NewUInt32(13)}}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00}

	var got GetSubscribedChannelsRequest

	roundTrip(t, "GetSubscribedChannelsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ResultRange == want.ResultRange) {
		t.Fatalf("GetSubscribedChannelsRequest.ResultRange differs")
	}
}

func TestGetSubscribedChannelsResponse_NoHeader(t *testing.T) {
	want := GetSubscribedChannelsResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetSubscribedChannelsResponse

	roundTrip(t, "GetSubscribedChannelsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetSubscribedChannelsResponse.Channels differs")
	}
}

func TestGetSubscribedChannelsResponse_Header(t *testing.T) {
	want := GetSubscribedChannelsResponse{Channels: types.List[news_types.NewsChannel]{func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(3032)
		v.OwnerPID = types.NewUInt32(4041)
		v.Name = types.NewString("s5")
		v.Description = types.NewString("s6")
		v.CreationTime = types.NewDateTime(7000098)
		v.ExpirationTime = types.NewDateTime(8000101)
		v.Type = types.NewString("s9")
		v.Locale = types.NewString("s10")
		v.Subscribable = types.NewBool(true)
		return v
	}(), func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
		v.ID = types.NewUInt32(13122)
		v.OwnerPID = types.NewUInt32(14131)
		v.Name = types.NewString("s15")
		v.Description = types.NewString("s16")
		v.CreationTime = types.NewDateTime(17000128)
		v.ExpirationTime = types.NewDateTime(18000131)
		v.Type = types.NewString("s19")
		v.Locale = types.NewString("s20")
		v.Subscribable = types.NewBool(true)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0x03, 0x00, 0x73, 0x36, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x01, 0x00, 0x31, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x35, 0x00, 0x04, 0x00, 0x73, 0x31, 0x36, 0x00, 0xc0, 0x66, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0xa9, 0x12, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0x01}

	var got GetSubscribedChannelsResponse

	roundTrip(t, "GetSubscribedChannelsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channels.Equals(want.Channels)) {
		t.Fatalf("GetSubscribedChannelsResponse.Channels differs")
	}
}

func TestSubscribeChannelRequest_NoHeader(t *testing.T) {
	want := SubscribeChannelRequest{NewsChannelID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got SubscribeChannelRequest

	roundTrip(t, "SubscribeChannelRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelID.Equals(want.NewsChannelID)) {
		t.Fatalf("SubscribeChannelRequest.NewsChannelID differs")
	}
}

func TestSubscribeChannelRequest_Header(t *testing.T) {
	want := SubscribeChannelRequest{NewsChannelID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got SubscribeChannelRequest

	roundTrip(t, "SubscribeChannelRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelID.Equals(want.NewsChannelID)) {
		t.Fatalf("SubscribeChannelRequest.NewsChannelID differs")
	}
}

func TestUnsubscribeChannelRequest_NoHeader(t *testing.T) {
	want := UnsubscribeChannelRequest{NewsChannelID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got UnsubscribeChannelRequest

	roundTrip(t, "UnsubscribeChannelRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelID.Equals(want.NewsChannelID)) {
		t.Fatalf("UnsubscribeChannelRequest.NewsChannelID differs")
	}
}

func TestUnsubscribeChannelRequest_Header(t *testing.T) {
	want := UnsubscribeChannelRequest{NewsChannelID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got UnsubscribeChannelRequest

	roundTrip(t, "UnsubscribeChannelRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelID.Equals(want.NewsChannelID)) {
		t.Fatalf("UnsubscribeChannelRequest.NewsChannelID differs")
	}
}

func TestGetNewsHeadersRequest_NoHeader(t *testing.T) {
	want := GetNewsHeadersRequest{Recipient: func() news_types.NewsRecipient {
		v := news_types.NewNewsRecipient()
		v.RecipientID = types.NewUInt32(2023)
		v.RecipientType = types.NewUInt32(3032)
		return v
	}(), Range: types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00}

	var got GetNewsHeadersRequest

	roundTrip(t, "GetNewsHeadersRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Recipient.Equals(want.Recipient)) {
		t.Fatalf("GetNewsHeadersRequest.Recipient differs")
	}
	if !(got.Range == want.Range) {
		t.Fatalf("GetNewsHeadersRequest.Range differs")
	}
}

func TestGetNewsHeadersRequest_Header(t *testing.T) {
	want := GetNewsHeadersRequest{Recipient: func() news_types.NewsRecipient {
		v := news_types.NewNewsRecipient()
		v.RecipientID = types.NewUInt32(2023)
		v.RecipientType = types.NewUInt32(3032)
		return v
	}(), Range: types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00}

	var got GetNewsHeadersRequest

	roundTrip(t, "GetNewsHeadersRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Recipient.Equals(want.Recipient)) {
		t.Fatalf("GetNewsHeadersRequest.Recipient differs")
	}
	if !(got.Range == want.Range) {
		t.Fatalf("GetNewsHeadersRequest.Range differs")
	}
}

func TestGetNewsHeadersResponse_NoHeader(t *testing.T) {
	want := GetNewsHeadersResponse{NewsHeaders: types.List[news_types.NewsHeader]{func() news_types.NewsHeader {
		v := news_types.NewNewsHeader()
		v.ID = types.NewUInt32(3032)
		v.RecipientID = types.NewUInt32(4041)
		v.RecipientType = types.NewUInt32(5050)
		v.PublisherPID = types.NewUInt32(6059)
		v.PublisherName = types.NewString("s7")
		v.PublicationTime = types.NewDateTime(8000101)
		v.DisplayTime = types.NewDateTime(9000104)
		v.ExpirationTime = types.NewDateTime(10000107)
		v.Title = types.NewString("s11")
		v.Link = types.NewString("s12")
		return v
	}(), func() news_types.NewsHeader {
		v := news_types.NewNewsHeader()
		v.ID = types.NewUInt32(14131)
		v.RecipientID = types.NewUInt32(15140)
		v.RecipientType = types.NewUInt32(16149)
		v.PublisherPID = types.NewUInt32(17158)
		v.PublisherName = types.NewString("s18")
		v.PublicationTime = types.NewDateTime(19000134)
		v.DisplayTime = types.NewDateTime(20000137)
		v.ExpirationTime = types.NewDateTime(21000140)
		v.Title = types.NewString("s22")
		v.Link = types.NewString("s23")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x38, 0x00, 0x46, 0xeb, 0x21, 0x01, 0x00, 0x00, 0x00, 0x00, 0x89, 0x2d, 0x31, 0x01, 0x00, 0x00, 0x00, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x32, 0x00, 0x04, 0x00, 0x73, 0x32, 0x33, 0x00}

	var got GetNewsHeadersResponse

	roundTrip(t, "GetNewsHeadersResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsHeaders.Equals(want.NewsHeaders)) {
		t.Fatalf("GetNewsHeadersResponse.NewsHeaders differs")
	}
}

func TestGetNewsHeadersResponse_Header(t *testing.T) {
	want := GetNewsHeadersResponse{NewsHeaders: types.List[news_types.NewsHeader]{func() news_types.NewsHeader {
		v := news_types.NewNewsHeader()
		v.ID = types.NewUInt32(3032)
		v.RecipientID = types.NewUInt32(4041)
		v.RecipientType = types.NewUInt32(5050)
		v.PublisherPID = types.NewUInt32(6059)
		v.PublisherName = types.NewString("s7")
		v.PublicationTime = types.NewDateTime(8000101)
		v.DisplayTime = types.NewDateTime(9000104)
		v.ExpirationTime = types.NewDateTime(10000107)
		v.Title = types.NewString("s11")
		v.Link = types.NewString("s12")
		return v
	}(), func() news_types.NewsHeader {
		v := news_types.NewNewsHeader()
		v.ID = types.NewUInt32(14131)
		v.RecipientID = types.NewUInt32(15140)
		v.RecipientType = types.NewUInt32(16149)
		v.PublisherPID = types.NewUInt32(17158)
		v.PublisherName = types.NewString("s18")
		v.PublicationTime = types.NewDateTime(19000134)
		v.DisplayTime = types.NewDateTime(20000137)
		v.ExpirationTime = types.NewDateTime(21000140)
		v.Title = types.NewString("s22")
		v.Link = types.NewString("s23")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x39, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x00, 0x3a, 0x00, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x38, 0x00, 0x46, 0xeb, 0x21, 0x01, 0x00, 0x00, 0x00, 0x00, 0x89, 0x2d, 0x31, 0x01, 0x00, 0x00, 0x00, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x32, 0x00, 0x04, 0x00, 0x73, 0x32, 0x33, 0x00}

	var got GetNewsHeadersResponse

	roundTrip(t, "GetNewsHeadersResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsHeaders.Equals(want.NewsHeaders)) {
		t.Fatalf("GetNewsHeadersResponse.NewsHeaders differs")
	}
}

func TestGetNewsMessagesRequest_NoHeader(t *testing.T) {
	want := GetNewsMessagesRequest{NewsMessageIDs: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetNewsMessagesRequest

	roundTrip(t, "GetNewsMessagesRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsMessageIDs.Equals(want.NewsMessageIDs)) {
		t.Fatalf("GetNewsMessagesRequest.NewsMessageIDs differs")
	}
}

func TestGetNewsMessagesRequest_Header(t *testing.T) {
	want := GetNewsMessagesRequest{NewsMessageIDs: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetNewsMessagesRequest

	roundTrip(t, "GetNewsMessagesRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsMessageIDs.Equals(want.NewsMessageIDs)) {
		t.Fatalf("GetNewsMessagesRequest.NewsMessageIDs differs")
	}
}

func TestGetNewsMessagesResponse_NoHeader(t *testing.T) {
	want := GetNewsMessagesResponse{NewsMessages: types.List[news_types.NewsMessage]{func() news_types.NewsMessage {
		v := news_types.NewNewsMessage()
		v.Body = types.NewString("s3")
		return v
	}(), func() news_types.NewsMessage {
		v := news_types.NewNewsMessage()
		v.Body = types.NewString("s5")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got GetNewsMessagesResponse

	roundTrip(t, "GetNewsMessagesResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsMessages.Equals(want.NewsMessages)) {
		t.Fatalf("GetNewsMessagesResponse.NewsMessages differs")
	}
}

func TestGetNewsMessagesResponse_Header(t *testing.T) {
	want := GetNewsMessagesResponse{NewsMessages: types.List[news_types.NewsMessage]{func() news_types.NewsMessage {
		v := news_types.NewNewsMessage()
		v.Body = types.NewString("s3")
		return v
	}(), func() news_types.NewsMessage {
		v := news_types.NewNewsMessage()
		v.Body = types.NewString("s5")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got GetNewsMessagesResponse

	roundTrip(t, "GetNewsMessagesResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsMessages.Equals(want.NewsMessages)) {
		t.Fatalf("GetNewsMessagesResponse.NewsMessages differs")
	}
}

func TestGetNumberOfNewsRequest_NoHeader(t *testing.T) {
	want := GetNumberOfNewsRequest{Recipient: func() news_types.NewsRecipient {
		v := news_types.NewNewsRecipient()
		v.RecipientID = types.NewUInt32(2023)
		v.RecipientType = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetNumberOfNewsRequest

	roundTrip(t, "GetNumberOfNewsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Recipient.Equals(want.Recipient)) {
		t.Fatalf("GetNumberOfNewsRequest.Recipient differs")
	}
}

func TestGetNumberOfNewsRequest_Header(t *testing.T) {
	want := GetNumberOfNewsRequest{Recipient: func() news_types.NewsRecipient {
		v := news_types.NewNewsRecipient()
		v.RecipientID = types.NewUInt32(2023)
		v.RecipientType = types.NewUInt32(3032)
		return v
	}()}

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetNumberOfNewsRequest

	roundTrip(t, "GetNumberOfNewsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Recipient.Equals(want.Recipient)) {
		t.Fatalf("GetNumberOfNewsRequest.Recipient differs")
	}
}

func TestGetNumberOfNewsResponse_NoHeader(t *testing.T) {
	want := GetNumberOfNewsResponse{NumberOfNews: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetNumberOfNewsResponse

	roundTrip(t, "GetNumberOfNewsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NumberOfNews.Equals(want.NumberOfNews)) {
		t.Fatalf("GetNumberOfNewsResponse.NumberOfNews differs")
	}
}

func TestGetNumberOfNewsResponse_Header(t *testing.T) {
	want := GetNumberOfNewsResponse{NumberOfNews: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetNumberOfNewsResponse

	roundTrip(t, "GetNumberOfNewsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NumberOfNews.Equals(want.NumberOfNews)) {
		t.Fatalf("GetNumberOfNewsResponse.NumberOfNews differs")
	}
}

func TestGetChannelByTypeRequest_NoHeader(t *testing.T) {
	want := GetChannelByTypeRequest{NewsChannelType: types.NewString("s1")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00}

	var got GetChannelByTypeRequest

	roundTrip(t, "GetChannelByTypeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelType.Equals(want.NewsChannelType)) {
		t.Fatalf("GetChannelByTypeRequest.NewsChannelType differs")
	}
}

func TestGetChannelByTypeRequest_Header(t *testing.T) {
	want := GetChannelByTypeRequest{NewsChannelType: types.NewString("s1")}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00}

	var got GetChannelByTypeRequest

	roundTrip(t, "GetChannelByTypeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelType.Equals(want.NewsChannelType)) {
		t.Fatalf("GetChannelByTypeRequest.NewsChannelType differs")
	}
}

func TestGetChannelByTypeResponse_NoHeader(t *testing.T) {
	want := GetChannelByTypeResponse{Channel: func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
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
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0xdf, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x00}

	var got GetChannelByTypeResponse

	roundTrip(t, "GetChannelByTypeResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channel.Equals(want.Channel)) {
		t.Fatalf("GetChannelByTypeResponse.Channel differs")
	}
}

func TestGetChannelByTypeResponse_Header(t *testing.T) {
	want := GetChannelByTypeResponse{Channel: func() news_types.NewsChannel {
		v := news_types.NewNewsChannel()
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
	}()}

	data := []byte{0x00, 0x2d, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00, 0xdf, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xd0, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0x00}

	var got GetChannelByTypeResponse

	roundTrip(t, "GetChannelByTypeResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Channel.Equals(want.Channel)) {
		t.Fatalf("GetChannelByTypeResponse.Channel differs")
	}
}

func TestGetNewsHeadersByTypeRequest_NoHeader(t *testing.T) {
	want := GetNewsHeadersByTypeRequest{NewsChannelType: types.NewString("s1"), Range: types.ResultRange{Offset: types.NewUInt32(6), Length: types.NewUInt32(16)}}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0x06, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00}

	var got GetNewsHeadersByTypeRequest

	roundTrip(t, "GetNewsHeadersByTypeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelType.Equals(want.NewsChannelType)) {
		t.Fatalf("GetNewsHeadersByTypeRequest.NewsChannelType differs")
	}
	if !(got.Range == want.Range) {
		t.Fatalf("GetNewsHeadersByTypeRequest.Range differs")
	}
}

func TestGetNewsHeadersByTypeRequest_Header(t *testing.T) {
	want := GetNewsHeadersByTypeRequest{NewsChannelType: types.NewString("s1"), Range: types.ResultRange{Offset: types.NewUInt32(6), Length: types.NewUInt32(16)}}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00}

	var got GetNewsHeadersByTypeRequest

	roundTrip(t, "GetNewsHeadersByTypeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelType.Equals(want.NewsChannelType)) {
		t.Fatalf("GetNewsHeadersByTypeRequest.NewsChannelType differs")
	}
	if !(got.Range == want.Range) {
		t.Fatalf("GetNewsHeadersByTypeRequest.Range differs")
	}
}

func TestGetNewsHeadersByTypeResponse_NoHeader(t *testing.T) {
	want := GetNewsHeadersByTypeResponse{NewsHeaders: types.List[news_types.NewsHeader]{func() news_types.NewsHeader {
		v := news_types.NewNewsHeader()
		v.ID = types.NewUInt32(3032)
		v.RecipientID = types.NewUInt32(4041)
		v.RecipientType = types.NewUInt32(5050)
		v.PublisherPID = types.NewUInt32(6059)
		v.PublisherName = types.NewString("s7")
		v.PublicationTime = types.NewDateTime(8000101)
		v.DisplayTime = types.NewDateTime(9000104)
		v.ExpirationTime = types.NewDateTime(10000107)
		v.Title = types.NewString("s11")
		v.Link = types.NewString("s12")
		return v
	}(), func() news_types.NewsHeader {
		v := news_types.NewNewsHeader()
		v.ID = types.NewUInt32(14131)
		v.RecipientID = types.NewUInt32(15140)
		v.RecipientType = types.NewUInt32(16149)
		v.PublisherPID = types.NewUInt32(17158)
		v.PublisherName = types.NewString("s18")
		v.PublicationTime = types.NewDateTime(19000134)
		v.DisplayTime = types.NewDateTime(20000137)
		v.ExpirationTime = types.NewDateTime(21000140)
		v.Title = types.NewString("s22")
		v.Link = types.NewString("s23")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x38, 0x00, 0x46, 0xeb, 0x21, 0x01, 0x00, 0x00, 0x00, 0x00, 0x89, 0x2d, 0x31, 0x01, 0x00, 0x00, 0x00, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x32, 0x00, 0x04, 0x00, 0x73, 0x32, 0x33, 0x00}

	var got GetNewsHeadersByTypeResponse

	roundTrip(t, "GetNewsHeadersByTypeResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsHeaders.Equals(want.NewsHeaders)) {
		t.Fatalf("GetNewsHeadersByTypeResponse.NewsHeaders differs")
	}
}

func TestGetNewsHeadersByTypeResponse_Header(t *testing.T) {
	want := GetNewsHeadersByTypeResponse{NewsHeaders: types.List[news_types.NewsHeader]{func() news_types.NewsHeader {
		v := news_types.NewNewsHeader()
		v.ID = types.NewUInt32(3032)
		v.RecipientID = types.NewUInt32(4041)
		v.RecipientType = types.NewUInt32(5050)
		v.PublisherPID = types.NewUInt32(6059)
		v.PublisherName = types.NewString("s7")
		v.PublicationTime = types.NewDateTime(8000101)
		v.DisplayTime = types.NewDateTime(9000104)
		v.ExpirationTime = types.NewDateTime(10000107)
		v.Title = types.NewString("s11")
		v.Link = types.NewString("s12")
		return v
	}(), func() news_types.NewsHeader {
		v := news_types.NewNewsHeader()
		v.ID = types.NewUInt32(14131)
		v.RecipientID = types.NewUInt32(15140)
		v.RecipientType = types.NewUInt32(16149)
		v.PublisherPID = types.NewUInt32(17158)
		v.PublisherName = types.NewString("s18")
		v.PublicationTime = types.NewDateTime(19000134)
		v.DisplayTime = types.NewDateTime(20000137)
		v.ExpirationTime = types.NewDateTime(21000140)
		v.Title = types.NewString("s22")
		v.Link = types.NewString("s23")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x39, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x65, 0x12, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x04, 0x00, 0x73, 0x31, 0x32, 0x00, 0x00, 0x3a, 0x00, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x38, 0x00, 0x46, 0xeb, 0x21, 0x01, 0x00, 0x00, 0x00, 0x00, 0x89, 0x2d, 0x31, 0x01, 0x00, 0x00, 0x00, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x32, 0x00, 0x04, 0x00, 0x73, 0x32, 0x33, 0x00}

	var got GetNewsHeadersByTypeResponse

	roundTrip(t, "GetNewsHeadersByTypeResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsHeaders.Equals(want.NewsHeaders)) {
		t.Fatalf("GetNewsHeadersByTypeResponse.NewsHeaders differs")
	}
}

func TestGetNewsMessagesByTypeRequest_NoHeader(t *testing.T) {
	want := GetNewsMessagesByTypeRequest{NewsChannelType: types.NewString("s1"), Range: types.ResultRange{Offset: types.NewUInt32(6), Length: types.NewUInt32(16)}}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0x06, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00}

	var got GetNewsMessagesByTypeRequest

	roundTrip(t, "GetNewsMessagesByTypeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelType.Equals(want.NewsChannelType)) {
		t.Fatalf("GetNewsMessagesByTypeRequest.NewsChannelType differs")
	}
	if !(got.Range == want.Range) {
		t.Fatalf("GetNewsMessagesByTypeRequest.Range differs")
	}
}

func TestGetNewsMessagesByTypeRequest_Header(t *testing.T) {
	want := GetNewsMessagesByTypeRequest{NewsChannelType: types.NewString("s1"), Range: types.ResultRange{Offset: types.NewUInt32(6), Length: types.NewUInt32(16)}}

	data := []byte{0x03, 0x00, 0x73, 0x31, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00}

	var got GetNewsMessagesByTypeRequest

	roundTrip(t, "GetNewsMessagesByTypeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsChannelType.Equals(want.NewsChannelType)) {
		t.Fatalf("GetNewsMessagesByTypeRequest.NewsChannelType differs")
	}
	if !(got.Range == want.Range) {
		t.Fatalf("GetNewsMessagesByTypeRequest.Range differs")
	}
}

func TestGetNewsMessagesByTypeResponse_NoHeader(t *testing.T) {
	want := GetNewsMessagesByTypeResponse{NewsMessages: types.List[news_types.NewsMessage]{func() news_types.NewsMessage {
		v := news_types.NewNewsMessage()
		v.Body = types.NewString("s3")
		return v
	}(), func() news_types.NewsMessage {
		v := news_types.NewNewsMessage()
		v.Body = types.NewString("s5")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got GetNewsMessagesByTypeResponse

	roundTrip(t, "GetNewsMessagesByTypeResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsMessages.Equals(want.NewsMessages)) {
		t.Fatalf("GetNewsMessagesByTypeResponse.NewsMessages differs")
	}
}

func TestGetNewsMessagesByTypeResponse_Header(t *testing.T) {
	want := GetNewsMessagesByTypeResponse{NewsMessages: types.List[news_types.NewsMessage]{func() news_types.NewsMessage {
		v := news_types.NewNewsMessage()
		v.Body = types.NewString("s3")
		return v
	}(), func() news_types.NewsMessage {
		v := news_types.NewNewsMessage()
		v.Body = types.NewString("s5")
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x35, 0x00}

	var got GetNewsMessagesByTypeResponse

	roundTrip(t, "GetNewsMessagesByTypeResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.NewsMessages.Equals(want.NewsMessages)) {
		t.Fatalf("GetNewsMessagesByTypeResponse.NewsMessages differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodGetChannels != 1 {
		t.Fatalf("MethodGetChannels = %d, want 1", MethodGetChannels)
	}
	if MethodGetChannelsByTypes != 2 {
		t.Fatalf("MethodGetChannelsByTypes = %d, want 2", MethodGetChannelsByTypes)
	}
	if MethodGetSubscribableChannels != 3 {
		t.Fatalf("MethodGetSubscribableChannels = %d, want 3", MethodGetSubscribableChannels)
	}
	if MethodGetChannelsByIDs != 4 {
		t.Fatalf("MethodGetChannelsByIDs = %d, want 4", MethodGetChannelsByIDs)
	}
	if MethodGetSubscribedChannels != 5 {
		t.Fatalf("MethodGetSubscribedChannels = %d, want 5", MethodGetSubscribedChannels)
	}
	if MethodSubscribeChannel != 6 {
		t.Fatalf("MethodSubscribeChannel = %d, want 6", MethodSubscribeChannel)
	}
	if MethodUnsubscribeChannel != 7 {
		t.Fatalf("MethodUnsubscribeChannel = %d, want 7", MethodUnsubscribeChannel)
	}
	if MethodGetNewsHeaders != 8 {
		t.Fatalf("MethodGetNewsHeaders = %d, want 8", MethodGetNewsHeaders)
	}
	if MethodGetNewsMessages != 9 {
		t.Fatalf("MethodGetNewsMessages = %d, want 9", MethodGetNewsMessages)
	}
	if MethodGetNumberOfNews != 10 {
		t.Fatalf("MethodGetNumberOfNews = %d, want 10", MethodGetNumberOfNews)
	}
	if MethodGetChannelByType != 11 {
		t.Fatalf("MethodGetChannelByType = %d, want 11", MethodGetChannelByType)
	}
	if MethodGetNewsHeadersByType != 12 {
		t.Fatalf("MethodGetNewsHeadersByType = %d, want 12", MethodGetNewsHeadersByType)
	}
	if MethodGetNewsMessagesByType != 13 {
		t.Fatalf("MethodGetNewsMessagesByType = %d, want 13", MethodGetNewsMessagesByType)
	}
	if ProtocolID != 31 {
		t.Fatalf("ProtocolID = %d, want 31", ProtocolID)
	}
}
