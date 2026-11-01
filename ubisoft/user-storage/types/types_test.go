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

func TestContentProperty_NoHeader(t *testing.T) {
	want := func() ContentProperty {
		v := NewContentProperty()
		v.ID = types.NewUInt32(2023)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v3")}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x33, 0x00}

	var got ContentProperty

	roundTrip(t, "ContentProperty", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded ContentProperty differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestContentProperty_Header(t *testing.T) {
	want := func() ContentProperty {
		v := NewContentProperty()
		v.ID = types.NewUInt32(2023)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v3")}
		return v
	}()

	data := []byte{0x00, 0x0a, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x33, 0x00}

	var got ContentProperty

	roundTrip(t, "ContentProperty", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded ContentProperty differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserContentKey_NoHeader(t *testing.T) {
	want := func() UserContentKey {
		v := NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got UserContentKey

	roundTrip(t, "UserContentKey", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserContentKey differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserContentKey_Header(t *testing.T) {
	want := func() UserContentKey {
		v := NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got UserContentKey

	roundTrip(t, "UserContentKey", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserContentKey differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserContent_NoHeader(t *testing.T) {
	want := func() UserContent {
		v := NewUserContent()
		v.Key = func() UserContentKey {
			v := NewUserContentKey()
			v.TypeID = types.NewUInt32(3032)
			v.ContentID = types.NewUInt64(400000039)
			return v
		}()
		v.Pid = types.NewUInt32(5050)
		v.Properties = types.List[ContentProperty]{func() ContentProperty {
			v := NewContentProperty()
			v.ID = types.NewUInt32(8077)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v9")}
			return v
		}(), func() ContentProperty {
			v := NewContentProperty()
			v.ID = types.NewUInt32(11104)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-17)}
			return v
		}()}
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0x27, 0x84, 0xd7, 0x17, 0x00, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x39, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x01, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got UserContent

	roundTrip(t, "UserContent", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserContent differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserContent_Header(t *testing.T) {
	want := func() UserContent {
		v := NewUserContent()
		v.Key = func() UserContentKey {
			v := NewUserContentKey()
			v.TypeID = types.NewUInt32(3032)
			v.ContentID = types.NewUInt64(400000039)
			return v
		}()
		v.Pid = types.NewUInt32(5050)
		v.Properties = types.List[ContentProperty]{func() ContentProperty {
			v := NewContentProperty()
			v.ID = types.NewUInt32(8077)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v9")}
			return v
		}(), func() ContentProperty {
			v := NewContentProperty()
			v.ID = types.NewUInt32(11104)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-17)}
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x3a, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x27, 0x84, 0xd7, 0x17, 0x00, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x39, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x01, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got UserContent

	roundTrip(t, "UserContent", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserContent differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserStorageQuery_NoHeader(t *testing.T) {
	want := func() UserStorageQuery {
		v := NewUserStorageQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}
		v.Parameters = types.List[ContentProperty]{func() ContentProperty {
			v := NewContentProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-13)}
			return v
		}(), func() ContentProperty {
			v := NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}()}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x01, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00}

	var got UserStorageQuery

	roundTrip(t, "UserStorageQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserStorageQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserStorageQuery_Header(t *testing.T) {
	want := func() UserStorageQuery {
		v := NewUserStorageQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}
		v.Parameters = types.List[ContentProperty]{func() ContentProperty {
			v := NewContentProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-13)}
			return v
		}(), func() ContentProperty {
			v := NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x3b, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x01, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00}

	var got UserStorageQuery

	roundTrip(t, "UserStorageQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserStorageQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserSlotCount_NoHeader(t *testing.T) {
	want := func() UserSlotCount {
		v := NewUserSlotCount()
		v.Pid = types.NewUInt32(2023)
		v.TypeID = types.NewUInt32(3032)
		v.UsedSlots = types.NewUInt32(4041)
		v.TotalSlots = types.NewUInt32(5050)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00}

	var got UserSlotCount

	roundTrip(t, "UserSlotCount", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserSlotCount differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserSlotCount_Header(t *testing.T) {
	want := func() UserSlotCount {
		v := NewUserSlotCount()
		v.Pid = types.NewUInt32(2023)
		v.TypeID = types.NewUInt32(3032)
		v.UsedSlots = types.NewUInt32(4041)
		v.TotalSlots = types.NewUInt32(5050)
		return v
	}()

	data := []byte{0x00, 0x10, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00}

	var got UserSlotCount

	roundTrip(t, "UserSlotCount", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserSlotCount differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserContentURL_NoHeader(t *testing.T) {
	want := func() UserContentURL {
		v := NewUserContentURL()
		v.Protocol = types.NewString("s2")
		v.Host = types.NewString("s3")
		v.Path = types.NewString("s4")
		return v
	}()

	data := []byte{0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00}

	var got UserContentURL

	roundTrip(t, "UserContentURL", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserContentURL differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestUserContentURL_Header(t *testing.T) {
	want := func() UserContentURL {
		v := NewUserContentURL()
		v.Protocol = types.NewString("s2")
		v.Host = types.NewString("s3")
		v.Path = types.NewString("s4")
		return v
	}()

	data := []byte{0x00, 0x0f, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00}

	var got UserContentURL

	roundTrip(t, "UserContentURL", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded UserContentURL differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestWeightedTag_NoHeader(t *testing.T) {
	want := func() WeightedTag {
		v := NewWeightedTag()
		v.Id = types.NewUInt32(2023)
		v.NumberOfOccurences = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got WeightedTag

	roundTrip(t, "WeightedTag", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded WeightedTag differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestWeightedTag_Header(t *testing.T) {
	want := func() WeightedTag {
		v := NewWeightedTag()
		v.Id = types.NewUInt32(2023)
		v.NumberOfOccurences = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got WeightedTag

	roundTrip(t, "WeightedTag", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded WeightedTag differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}
