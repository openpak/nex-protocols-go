package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
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

func TestSearchContentsRequest_NoHeader(t *testing.T) {
	want := SearchContentsRequest{Query: func() user_storage_types.UserStorageQuery {
		v := user_storage_types.NewUserStorageQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}
		v.Parameters = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-13)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}()}
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x01, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00}

	var got SearchContentsRequest

	roundTrip(t, "SearchContentsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Query.Equals(want.Query)) {
		t.Fatalf("SearchContentsRequest.Query differs")
	}
}

func TestSearchContentsRequest_Header(t *testing.T) {
	want := SearchContentsRequest{Query: func() user_storage_types.UserStorageQuery {
		v := user_storage_types.NewUserStorageQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}
		v.Parameters = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-13)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}()}
		return v
	}()}

	data := []byte{0x00, 0x3b, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x01, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00}

	var got SearchContentsRequest

	roundTrip(t, "SearchContentsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Query.Equals(want.Query)) {
		t.Fatalf("SearchContentsRequest.Query differs")
	}
}

func TestSearchContentsResponse_NoHeader(t *testing.T) {
	want := SearchContentsResponse{SearchResults: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got SearchContentsResponse

	roundTrip(t, "SearchContentsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchContentsResponse.SearchResults differs")
	}
}

func TestSearchContentsResponse_Header(t *testing.T) {
	want := SearchContentsResponse{SearchResults: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got SearchContentsResponse

	roundTrip(t, "SearchContentsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchContentsResponse.SearchResults differs")
	}
}

func TestSearchContentsWithTotalRequest_NoHeader(t *testing.T) {
	want := SearchContentsWithTotalRequest{Query: func() user_storage_types.UserStorageQuery {
		v := user_storage_types.NewUserStorageQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}
		v.Parameters = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-13)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}()}
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x01, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00}

	var got SearchContentsWithTotalRequest

	roundTrip(t, "SearchContentsWithTotalRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Query.Equals(want.Query)) {
		t.Fatalf("SearchContentsWithTotalRequest.Query differs")
	}
}

func TestSearchContentsWithTotalRequest_Header(t *testing.T) {
	want := SearchContentsWithTotalRequest{Query: func() user_storage_types.UserStorageQuery {
		v := user_storage_types.NewUserStorageQuery()
		v.TypeID = types.NewUInt32(2023)
		v.QueryID = types.NewUInt32(3032)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(12), Length: types.NewUInt32(22)}
		v.Parameters = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(7068)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-13)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}()}
		return v
	}()}

	data := []byte{0x00, 0x3b, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x01, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00}

	var got SearchContentsWithTotalRequest

	roundTrip(t, "SearchContentsWithTotalRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Query.Equals(want.Query)) {
		t.Fatalf("SearchContentsWithTotalRequest.Query differs")
	}
}

func TestSearchContentsWithTotalResponse_NoHeader(t *testing.T) {
	want := SearchContentsWithTotalResponse{SearchResults: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}, TotalResults: types.NewUInt32(26239)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x7f, 0x66, 0x00, 0x00}

	var got SearchContentsWithTotalResponse

	roundTrip(t, "SearchContentsWithTotalResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchContentsWithTotalResponse.SearchResults differs")
	}
	if !(got.TotalResults.Equals(want.TotalResults)) {
		t.Fatalf("SearchContentsWithTotalResponse.TotalResults differs")
	}
}

func TestSearchContentsWithTotalResponse_Header(t *testing.T) {
	want := SearchContentsWithTotalResponse{SearchResults: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}, TotalResults: types.NewUInt32(26239)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x7f, 0x66, 0x00, 0x00}

	var got SearchContentsWithTotalResponse

	roundTrip(t, "SearchContentsWithTotalResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchContentsWithTotalResponse.SearchResults differs")
	}
	if !(got.TotalResults.Equals(want.TotalResults)) {
		t.Fatalf("SearchContentsWithTotalResponse.TotalResults differs")
	}
}

func TestDeleteContentRequest_NoHeader(t *testing.T) {
	want := DeleteContentRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got DeleteContentRequest

	roundTrip(t, "DeleteContentRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("DeleteContentRequest.ContentKey differs")
	}
}

func TestDeleteContentRequest_Header(t *testing.T) {
	want := DeleteContentRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got DeleteContentRequest

	roundTrip(t, "DeleteContentRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("DeleteContentRequest.ContentKey differs")
	}
}

func TestSaveMetaDataRequest_NoHeader(t *testing.T) {
	want := SaveMetaDataRequest{Properties: types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(3032)
		v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-9)}
		return v
	}(), func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(6059)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v7")}
		return v
	}()}, ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(9086)
		v.ContentID = types.NewUInt64(1000000081)
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x01, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xab, 0x17, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x37, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x51, 0xca, 0x9a, 0x3b, 0x00, 0x00, 0x00, 0x00}

	var got SaveMetaDataRequest

	roundTrip(t, "SaveMetaDataRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Properties.Equals(want.Properties)) {
		t.Fatalf("SaveMetaDataRequest.Properties differs")
	}
	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveMetaDataRequest.ContentKey differs")
	}
}

func TestSaveMetaDataRequest_Header(t *testing.T) {
	want := SaveMetaDataRequest{Properties: types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(3032)
		v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-9)}
		return v
	}(), func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(6059)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v7")}
		return v
	}()}, ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(9086)
		v.ContentID = types.NewUInt64(1000000081)
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x01, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x37, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x51, 0xca, 0x9a, 0x3b, 0x00, 0x00, 0x00, 0x00}

	var got SaveMetaDataRequest

	roundTrip(t, "SaveMetaDataRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Properties.Equals(want.Properties)) {
		t.Fatalf("SaveMetaDataRequest.Properties differs")
	}
	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveMetaDataRequest.ContentKey differs")
	}
}

func TestSaveMetaDataResponse_NoHeader(t *testing.T) {
	want := SaveMetaDataResponse{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got SaveMetaDataResponse

	roundTrip(t, "SaveMetaDataResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveMetaDataResponse.ContentKey differs")
	}
}

func TestSaveMetaDataResponse_Header(t *testing.T) {
	want := SaveMetaDataResponse{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got SaveMetaDataResponse

	roundTrip(t, "SaveMetaDataResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveMetaDataResponse.ContentKey differs")
	}
}

func TestSaveContentDBRequest_NoHeader(t *testing.T) {
	want := SaveContentDBRequest{Properties: types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(3032)
		v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-9)}
		return v
	}(), func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(6059)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v7")}
		return v
	}()}, Data: types.NewBuffer([]byte{8, 9, 10}), ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(10095)
		v.ContentID = types.NewUInt64(1100000088)
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x01, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xab, 0x17, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x37, 0x00, 0x03, 0x00, 0x00, 0x00, 0x08, 0x09, 0x0a, 0x6f, 0x27, 0x00, 0x00, 0x58, 0xab, 0x90, 0x41, 0x00, 0x00, 0x00, 0x00}

	var got SaveContentDBRequest

	roundTrip(t, "SaveContentDBRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Properties.Equals(want.Properties)) {
		t.Fatalf("SaveContentDBRequest.Properties differs")
	}
	if !(got.Data.Equals(want.Data)) {
		t.Fatalf("SaveContentDBRequest.Data differs")
	}
	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveContentDBRequest.ContentKey differs")
	}
}

func TestSaveContentDBRequest_Header(t *testing.T) {
	want := SaveContentDBRequest{Properties: types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(3032)
		v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-9)}
		return v
	}(), func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(6059)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v7")}
		return v
	}()}, Data: types.NewBuffer([]byte{8, 9, 10}), ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(10095)
		v.ContentID = types.NewUInt64(1100000088)
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x01, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x37, 0x00, 0x03, 0x00, 0x00, 0x00, 0x08, 0x09, 0x0a, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x58, 0xab, 0x90, 0x41, 0x00, 0x00, 0x00, 0x00}

	var got SaveContentDBRequest

	roundTrip(t, "SaveContentDBRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Properties.Equals(want.Properties)) {
		t.Fatalf("SaveContentDBRequest.Properties differs")
	}
	if !(got.Data.Equals(want.Data)) {
		t.Fatalf("SaveContentDBRequest.Data differs")
	}
	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveContentDBRequest.ContentKey differs")
	}
}

func TestSaveContentDBResponse_NoHeader(t *testing.T) {
	want := SaveContentDBResponse{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got SaveContentDBResponse

	roundTrip(t, "SaveContentDBResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveContentDBResponse.ContentKey differs")
	}
}

func TestSaveContentDBResponse_Header(t *testing.T) {
	want := SaveContentDBResponse{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got SaveContentDBResponse

	roundTrip(t, "SaveContentDBResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveContentDBResponse.ContentKey differs")
	}
}

func TestSaveContentAndGetUploadInfoRequest_NoHeader(t *testing.T) {
	want := SaveContentAndGetUploadInfoRequest{Properties: types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(3032)
		v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-9)}
		return v
	}(), func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(6059)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v7")}
		return v
	}()}, Size: types.NewUInt32(8077), ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(10095)
		v.ContentID = types.NewUInt64(1100000088)
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x01, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xab, 0x17, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x58, 0xab, 0x90, 0x41, 0x00, 0x00, 0x00, 0x00}

	var got SaveContentAndGetUploadInfoRequest

	roundTrip(t, "SaveContentAndGetUploadInfoRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Properties.Equals(want.Properties)) {
		t.Fatalf("SaveContentAndGetUploadInfoRequest.Properties differs")
	}
	if !(got.Size.Equals(want.Size)) {
		t.Fatalf("SaveContentAndGetUploadInfoRequest.Size differs")
	}
	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveContentAndGetUploadInfoRequest.ContentKey differs")
	}
}

func TestSaveContentAndGetUploadInfoRequest_Header(t *testing.T) {
	want := SaveContentAndGetUploadInfoRequest{Properties: types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(3032)
		v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-9)}
		return v
	}(), func() user_storage_types.ContentProperty {
		v := user_storage_types.NewContentProperty()
		v.ID = types.NewUInt32(6059)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v7")}
		return v
	}()}, Size: types.NewUInt32(8077), ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(10095)
		v.ContentID = types.NewUInt64(1100000088)
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x01, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x37, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x58, 0xab, 0x90, 0x41, 0x00, 0x00, 0x00, 0x00}

	var got SaveContentAndGetUploadInfoRequest

	roundTrip(t, "SaveContentAndGetUploadInfoRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Properties.Equals(want.Properties)) {
		t.Fatalf("SaveContentAndGetUploadInfoRequest.Properties differs")
	}
	if !(got.Size.Equals(want.Size)) {
		t.Fatalf("SaveContentAndGetUploadInfoRequest.Size differs")
	}
	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("SaveContentAndGetUploadInfoRequest.ContentKey differs")
	}
}

func TestSaveContentAndGetUploadInfoResponse_NoHeader(t *testing.T) {
	want := SaveContentAndGetUploadInfoResponse{UploadInfo: func() user_storage_types.UserContentURL {
		v := user_storage_types.NewUserContentURL()
		v.Protocol = types.NewString("s2")
		v.Host = types.NewString("s3")
		v.Path = types.NewString("s4")
		return v
	}(), PendingID: types.NewUInt64(500000046), Headers: types.List[types.String]{types.NewString("s7"), types.NewString("s8")}}

	data := []byte{0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00}

	var got SaveContentAndGetUploadInfoResponse

	roundTrip(t, "SaveContentAndGetUploadInfoResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.UploadInfo.Equals(want.UploadInfo)) {
		t.Fatalf("SaveContentAndGetUploadInfoResponse.UploadInfo differs")
	}
	if !(got.PendingID.Equals(want.PendingID)) {
		t.Fatalf("SaveContentAndGetUploadInfoResponse.PendingID differs")
	}
	if !(got.Headers.Equals(want.Headers)) {
		t.Fatalf("SaveContentAndGetUploadInfoResponse.Headers differs")
	}
}

func TestSaveContentAndGetUploadInfoResponse_Header(t *testing.T) {
	want := SaveContentAndGetUploadInfoResponse{UploadInfo: func() user_storage_types.UserContentURL {
		v := user_storage_types.NewUserContentURL()
		v.Protocol = types.NewString("s2")
		v.Host = types.NewString("s3")
		v.Path = types.NewString("s4")
		return v
	}(), PendingID: types.NewUInt64(500000046), Headers: types.List[types.String]{types.NewString("s7"), types.NewString("s8")}}

	data := []byte{0x00, 0x0f, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x37, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00}

	var got SaveContentAndGetUploadInfoResponse

	roundTrip(t, "SaveContentAndGetUploadInfoResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.UploadInfo.Equals(want.UploadInfo)) {
		t.Fatalf("SaveContentAndGetUploadInfoResponse.UploadInfo differs")
	}
	if !(got.PendingID.Equals(want.PendingID)) {
		t.Fatalf("SaveContentAndGetUploadInfoResponse.PendingID differs")
	}
	if !(got.Headers.Equals(want.Headers)) {
		t.Fatalf("SaveContentAndGetUploadInfoResponse.Headers differs")
	}
}

func TestUploadEndRequest_NoHeader(t *testing.T) {
	want := UploadEndRequest{PendingID: types.NewUInt64(100000018), Result: types.NewBool(false), ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(4041)
		v.ContentID = types.NewUInt64(500000046)
		return v
	}()}

	data := []byte{0x12, 0xe1, 0xf5, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00}

	var got UploadEndRequest

	roundTrip(t, "UploadEndRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PendingID.Equals(want.PendingID)) {
		t.Fatalf("UploadEndRequest.PendingID differs")
	}
	if !(got.Result.Equals(want.Result)) {
		t.Fatalf("UploadEndRequest.Result differs")
	}
	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UploadEndRequest.ContentKey differs")
	}
}

func TestUploadEndRequest_Header(t *testing.T) {
	want := UploadEndRequest{PendingID: types.NewUInt64(100000018), Result: types.NewBool(false), ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(4041)
		v.ContentID = types.NewUInt64(500000046)
		return v
	}()}

	data := []byte{0x12, 0xe1, 0xf5, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00}

	var got UploadEndRequest

	roundTrip(t, "UploadEndRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PendingID.Equals(want.PendingID)) {
		t.Fatalf("UploadEndRequest.PendingID differs")
	}
	if !(got.Result.Equals(want.Result)) {
		t.Fatalf("UploadEndRequest.Result differs")
	}
	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UploadEndRequest.ContentKey differs")
	}
}

func TestUploadEndResponse_NoHeader(t *testing.T) {
	want := UploadEndResponse{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got UploadEndResponse

	roundTrip(t, "UploadEndResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UploadEndResponse.ContentKey differs")
	}
}

func TestUploadEndResponse_Header(t *testing.T) {
	want := UploadEndResponse{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got UploadEndResponse

	roundTrip(t, "UploadEndResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UploadEndResponse.ContentKey differs")
	}
}

func TestGetContentDBRequest_NoHeader(t *testing.T) {
	want := GetContentDBRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetContentDBRequest

	roundTrip(t, "GetContentDBRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetContentDBRequest.ContentKey differs")
	}
}

func TestGetContentDBRequest_Header(t *testing.T) {
	want := GetContentDBRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetContentDBRequest

	roundTrip(t, "GetContentDBRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetContentDBRequest.ContentKey differs")
	}
}

func TestGetContentDBResponse_NoHeader(t *testing.T) {
	want := GetContentDBResponse{Data: types.NewBuffer([]byte{1, 2, 3})}

	data := []byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03}

	var got GetContentDBResponse

	roundTrip(t, "GetContentDBResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Data.Equals(want.Data)) {
		t.Fatalf("GetContentDBResponse.Data differs")
	}
}

func TestGetContentDBResponse_Header(t *testing.T) {
	want := GetContentDBResponse{Data: types.NewBuffer([]byte{1, 2, 3})}

	data := []byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03}

	var got GetContentDBResponse

	roundTrip(t, "GetContentDBResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Data.Equals(want.Data)) {
		t.Fatalf("GetContentDBResponse.Data differs")
	}
}

func TestGetContentURLRequest_NoHeader(t *testing.T) {
	want := GetContentURLRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetContentURLRequest

	roundTrip(t, "GetContentURLRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetContentURLRequest.ContentKey differs")
	}
}

func TestGetContentURLRequest_Header(t *testing.T) {
	want := GetContentURLRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetContentURLRequest

	roundTrip(t, "GetContentURLRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetContentURLRequest.ContentKey differs")
	}
}

func TestGetContentURLResponse_NoHeader(t *testing.T) {
	want := GetContentURLResponse{DownloadInfo: func() user_storage_types.UserContentURL {
		v := user_storage_types.NewUserContentURL()
		v.Protocol = types.NewString("s2")
		v.Host = types.NewString("s3")
		v.Path = types.NewString("s4")
		return v
	}()}

	data := []byte{0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00}

	var got GetContentURLResponse

	roundTrip(t, "GetContentURLResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.DownloadInfo.Equals(want.DownloadInfo)) {
		t.Fatalf("GetContentURLResponse.DownloadInfo differs")
	}
}

func TestGetContentURLResponse_Header(t *testing.T) {
	want := GetContentURLResponse{DownloadInfo: func() user_storage_types.UserContentURL {
		v := user_storage_types.NewUserContentURL()
		v.Protocol = types.NewString("s2")
		v.Host = types.NewString("s3")
		v.Path = types.NewString("s4")
		return v
	}()}

	data := []byte{0x00, 0x0f, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x32, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00}

	var got GetContentURLResponse

	roundTrip(t, "GetContentURLResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.DownloadInfo.Equals(want.DownloadInfo)) {
		t.Fatalf("GetContentURLResponse.DownloadInfo differs")
	}
}

func TestGetSlotCountRequest_NoHeader(t *testing.T) {
	want := GetSlotCountRequest{TypeID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetSlotCountRequest

	roundTrip(t, "GetSlotCountRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("GetSlotCountRequest.TypeID differs")
	}
}

func TestGetSlotCountRequest_Header(t *testing.T) {
	want := GetSlotCountRequest{TypeID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetSlotCountRequest

	roundTrip(t, "GetSlotCountRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("GetSlotCountRequest.TypeID differs")
	}
}

func TestGetSlotCountResponse_NoHeader(t *testing.T) {
	want := GetSlotCountResponse{SlotCount: func() user_storage_types.UserSlotCount {
		v := user_storage_types.NewUserSlotCount()
		v.Pid = types.NewUInt32(2023)
		v.TypeID = types.NewUInt32(3032)
		v.UsedSlots = types.NewUInt32(4041)
		v.TotalSlots = types.NewUInt32(5050)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00}

	var got GetSlotCountResponse

	roundTrip(t, "GetSlotCountResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SlotCount.Equals(want.SlotCount)) {
		t.Fatalf("GetSlotCountResponse.SlotCount differs")
	}
}

func TestGetSlotCountResponse_Header(t *testing.T) {
	want := GetSlotCountResponse{SlotCount: func() user_storage_types.UserSlotCount {
		v := user_storage_types.NewUserSlotCount()
		v.Pid = types.NewUInt32(2023)
		v.TypeID = types.NewUInt32(3032)
		v.UsedSlots = types.NewUInt32(4041)
		v.TotalSlots = types.NewUInt32(5050)
		return v
	}()}

	data := []byte{0x00, 0x10, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00}

	var got GetSlotCountResponse

	roundTrip(t, "GetSlotCountResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SlotCount.Equals(want.SlotCount)) {
		t.Fatalf("GetSlotCountResponse.SlotCount differs")
	}
}

func TestGetMetaDataRequest_NoHeader(t *testing.T) {
	want := GetMetaDataRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetMetaDataRequest

	roundTrip(t, "GetMetaDataRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetMetaDataRequest.ContentKey differs")
	}
}

func TestGetMetaDataRequest_Header(t *testing.T) {
	want := GetMetaDataRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetMetaDataRequest

	roundTrip(t, "GetMetaDataRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetMetaDataRequest.ContentKey differs")
	}
}

func TestGetMetaDataResponse_NoHeader(t *testing.T) {
	want := GetMetaDataResponse{Content: func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(3032)
			v.ContentID = types.NewUInt64(400000039)
			return v
		}()
		v.Pid = types.NewUInt32(5050)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(8077)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v9")}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(11104)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-17)}
			return v
		}()}
		return v
	}()}

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0x27, 0x84, 0xd7, 0x17, 0x00, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x39, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x01, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got GetMetaDataResponse

	roundTrip(t, "GetMetaDataResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Content.Equals(want.Content)) {
		t.Fatalf("GetMetaDataResponse.Content differs")
	}
}

func TestGetMetaDataResponse_Header(t *testing.T) {
	want := GetMetaDataResponse{Content: func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(3032)
			v.ContentID = types.NewUInt64(400000039)
			return v
		}()
		v.Pid = types.NewUInt32(5050)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(8077)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v9")}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(11104)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-17)}
			return v
		}()}
		return v
	}()}

	data := []byte{0x00, 0x3a, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x27, 0x84, 0xd7, 0x17, 0x00, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x39, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x01, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got GetMetaDataResponse

	roundTrip(t, "GetMetaDataResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Content.Equals(want.Content)) {
		t.Fatalf("GetMetaDataResponse.Content differs")
	}
}

func TestLikeRequest_NoHeader(t *testing.T) {
	want := LikeRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got LikeRequest

	roundTrip(t, "LikeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("LikeRequest.ContentKey differs")
	}
}

func TestLikeRequest_Header(t *testing.T) {
	want := LikeRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got LikeRequest

	roundTrip(t, "LikeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("LikeRequest.ContentKey differs")
	}
}

func TestUnlikeRequest_NoHeader(t *testing.T) {
	want := UnlikeRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got UnlikeRequest

	roundTrip(t, "UnlikeRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UnlikeRequest.ContentKey differs")
	}
}

func TestUnlikeRequest_Header(t *testing.T) {
	want := UnlikeRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got UnlikeRequest

	roundTrip(t, "UnlikeRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UnlikeRequest.ContentKey differs")
	}
}

func TestIsLikedRequest_NoHeader(t *testing.T) {
	want := IsLikedRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got IsLikedRequest

	roundTrip(t, "IsLikedRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("IsLikedRequest.ContentKey differs")
	}
}

func TestIsLikedRequest_Header(t *testing.T) {
	want := IsLikedRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got IsLikedRequest

	roundTrip(t, "IsLikedRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("IsLikedRequest.ContentKey differs")
	}
}

func TestIsLikedResponse_NoHeader(t *testing.T) {
	want := IsLikedResponse{Liked: types.NewBool(true)}

	data := []byte{0x01}

	var got IsLikedResponse

	roundTrip(t, "IsLikedResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Liked.Equals(want.Liked)) {
		t.Fatalf("IsLikedResponse.Liked differs")
	}
}

func TestIsLikedResponse_Header(t *testing.T) {
	want := IsLikedResponse{Liked: types.NewBool(true)}

	data := []byte{0x01}

	var got IsLikedResponse

	roundTrip(t, "IsLikedResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Liked.Equals(want.Liked)) {
		t.Fatalf("IsLikedResponse.Liked differs")
	}
}

func TestGetFavouritesRequest_NoHeader(t *testing.T) {
	want := GetFavouritesRequest{ContentTypes: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetFavouritesRequest

	roundTrip(t, "GetFavouritesRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentTypes.Equals(want.ContentTypes)) {
		t.Fatalf("GetFavouritesRequest.ContentTypes differs")
	}
}

func TestGetFavouritesRequest_Header(t *testing.T) {
	want := GetFavouritesRequest{ContentTypes: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetFavouritesRequest

	roundTrip(t, "GetFavouritesRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentTypes.Equals(want.ContentTypes)) {
		t.Fatalf("GetFavouritesRequest.ContentTypes differs")
	}
}

func TestGetFavouritesResponse_NoHeader(t *testing.T) {
	want := GetFavouritesResponse{Favourites: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got GetFavouritesResponse

	roundTrip(t, "GetFavouritesResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Favourites.Equals(want.Favourites)) {
		t.Fatalf("GetFavouritesResponse.Favourites differs")
	}
}

func TestGetFavouritesResponse_Header(t *testing.T) {
	want := GetFavouritesResponse{Favourites: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got GetFavouritesResponse

	roundTrip(t, "GetFavouritesResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Favourites.Equals(want.Favourites)) {
		t.Fatalf("GetFavouritesResponse.Favourites differs")
	}
}

func TestMakeFavouriteRequest_NoHeader(t *testing.T) {
	want := MakeFavouriteRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got MakeFavouriteRequest

	roundTrip(t, "MakeFavouriteRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("MakeFavouriteRequest.ContentKey differs")
	}
}

func TestMakeFavouriteRequest_Header(t *testing.T) {
	want := MakeFavouriteRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got MakeFavouriteRequest

	roundTrip(t, "MakeFavouriteRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("MakeFavouriteRequest.ContentKey differs")
	}
}

func TestRemoveFromFavouritesRequest_NoHeader(t *testing.T) {
	want := RemoveFromFavouritesRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got RemoveFromFavouritesRequest

	roundTrip(t, "RemoveFromFavouritesRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("RemoveFromFavouritesRequest.ContentKey differs")
	}
}

func TestRemoveFromFavouritesRequest_Header(t *testing.T) {
	want := RemoveFromFavouritesRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got RemoveFromFavouritesRequest

	roundTrip(t, "RemoveFromFavouritesRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("RemoveFromFavouritesRequest.ContentKey differs")
	}
}

func TestReportInappropriateRequest_NoHeader(t *testing.T) {
	want := ReportInappropriateRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}(), Reason: types.NewString("s4")}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00}

	var got ReportInappropriateRequest

	roundTrip(t, "ReportInappropriateRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("ReportInappropriateRequest.ContentKey differs")
	}
	if !(got.Reason.Equals(want.Reason)) {
		t.Fatalf("ReportInappropriateRequest.Reason differs")
	}
}

func TestReportInappropriateRequest_Header(t *testing.T) {
	want := ReportInappropriateRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}(), Reason: types.NewString("s4")}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00}

	var got ReportInappropriateRequest

	roundTrip(t, "ReportInappropriateRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("ReportInappropriateRequest.ContentKey differs")
	}
	if !(got.Reason.Equals(want.Reason)) {
		t.Fatalf("ReportInappropriateRequest.Reason differs")
	}
}

func TestIncrementPlayCountRequest_NoHeader(t *testing.T) {
	want := IncrementPlayCountRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got IncrementPlayCountRequest

	roundTrip(t, "IncrementPlayCountRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("IncrementPlayCountRequest.ContentKey differs")
	}
}

func TestIncrementPlayCountRequest_Header(t *testing.T) {
	want := IncrementPlayCountRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got IncrementPlayCountRequest

	roundTrip(t, "IncrementPlayCountRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("IncrementPlayCountRequest.ContentKey differs")
	}
}

func TestUpdateCustomStatRequest_NoHeader(t *testing.T) {
	want := UpdateCustomStatRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}(), StatID: types.NewUInt16(407), IncValue: types.NewInt64(-20000000036)}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00, 0x97, 0x01, 0xdc, 0x37, 0xe8, 0x57, 0xfb, 0xff, 0xff, 0xff}

	var got UpdateCustomStatRequest

	roundTrip(t, "UpdateCustomStatRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UpdateCustomStatRequest.ContentKey differs")
	}
	if !(got.StatID.Equals(want.StatID)) {
		t.Fatalf("UpdateCustomStatRequest.StatID differs")
	}
	if !(got.IncValue.Equals(want.IncValue)) {
		t.Fatalf("UpdateCustomStatRequest.IncValue differs")
	}
}

func TestUpdateCustomStatRequest_Header(t *testing.T) {
	want := UpdateCustomStatRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}(), StatID: types.NewUInt16(407), IncValue: types.NewInt64(-20000000036)}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00, 0x97, 0x01, 0xdc, 0x37, 0xe8, 0x57, 0xfb, 0xff, 0xff, 0xff}

	var got UpdateCustomStatRequest

	roundTrip(t, "UpdateCustomStatRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("UpdateCustomStatRequest.ContentKey differs")
	}
	if !(got.StatID.Equals(want.StatID)) {
		t.Fatalf("UpdateCustomStatRequest.StatID differs")
	}
	if !(got.IncValue.Equals(want.IncValue)) {
		t.Fatalf("UpdateCustomStatRequest.IncValue differs")
	}
}

func TestGetOwnContentsRequest_NoHeader(t *testing.T) {
	want := GetOwnContentsRequest{TypeID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetOwnContentsRequest

	roundTrip(t, "GetOwnContentsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("GetOwnContentsRequest.TypeID differs")
	}
}

func TestGetOwnContentsRequest_Header(t *testing.T) {
	want := GetOwnContentsRequest{TypeID: types.NewUInt32(1014)}

	data := []byte{0xf6, 0x03, 0x00, 0x00}

	var got GetOwnContentsRequest

	roundTrip(t, "GetOwnContentsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TypeID.Equals(want.TypeID)) {
		t.Fatalf("GetOwnContentsRequest.TypeID differs")
	}
}

func TestGetOwnContentsResponse_NoHeader(t *testing.T) {
	want := GetOwnContentsResponse{Results: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got GetOwnContentsResponse

	roundTrip(t, "GetOwnContentsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("GetOwnContentsResponse.Results differs")
	}
}

func TestGetOwnContentsResponse_Header(t *testing.T) {
	want := GetOwnContentsResponse{Results: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got GetOwnContentsResponse

	roundTrip(t, "GetOwnContentsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("GetOwnContentsResponse.Results differs")
	}
}

func TestGetMostPopularTagsRequest_NoHeader(t *testing.T) {
	want := GetMostPopularTagsRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetMostPopularTagsRequest

	roundTrip(t, "GetMostPopularTagsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetMostPopularTagsRequest.ContentKey differs")
	}
}

func TestGetMostPopularTagsRequest_Header(t *testing.T) {
	want := GetMostPopularTagsRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetMostPopularTagsRequest

	roundTrip(t, "GetMostPopularTagsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetMostPopularTagsRequest.ContentKey differs")
	}
}

func TestGetMostPopularTagsResponse_NoHeader(t *testing.T) {
	want := GetMostPopularTagsResponse{Tags: types.List[user_storage_types.WeightedTag]{func() user_storage_types.WeightedTag {
		v := user_storage_types.NewWeightedTag()
		v.Id = types.NewUInt32(3032)
		v.NumberOfOccurences = types.NewUInt32(4041)
		return v
	}(), func() user_storage_types.WeightedTag {
		v := user_storage_types.NewWeightedTag()
		v.Id = types.NewUInt32(6059)
		v.NumberOfOccurences = types.NewUInt32(7068)
		return v
	}()}, TotalNumberOfTaggings: types.NewUInt32(8077)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00}

	var got GetMostPopularTagsResponse

	roundTrip(t, "GetMostPopularTagsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Tags.Equals(want.Tags)) {
		t.Fatalf("GetMostPopularTagsResponse.Tags differs")
	}
	if !(got.TotalNumberOfTaggings.Equals(want.TotalNumberOfTaggings)) {
		t.Fatalf("GetMostPopularTagsResponse.TotalNumberOfTaggings differs")
	}
}

func TestGetMostPopularTagsResponse_Header(t *testing.T) {
	want := GetMostPopularTagsResponse{Tags: types.List[user_storage_types.WeightedTag]{func() user_storage_types.WeightedTag {
		v := user_storage_types.NewWeightedTag()
		v.Id = types.NewUInt32(3032)
		v.NumberOfOccurences = types.NewUInt32(4041)
		return v
	}(), func() user_storage_types.WeightedTag {
		v := user_storage_types.NewWeightedTag()
		v.Id = types.NewUInt32(6059)
		v.NumberOfOccurences = types.NewUInt32(7068)
		return v
	}()}, TotalNumberOfTaggings: types.NewUInt32(8077)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00}

	var got GetMostPopularTagsResponse

	roundTrip(t, "GetMostPopularTagsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Tags.Equals(want.Tags)) {
		t.Fatalf("GetMostPopularTagsResponse.Tags differs")
	}
	if !(got.TotalNumberOfTaggings.Equals(want.TotalNumberOfTaggings)) {
		t.Fatalf("GetMostPopularTagsResponse.TotalNumberOfTaggings differs")
	}
}

func TestGetTagsRequest_NoHeader(t *testing.T) {
	want := GetTagsRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetTagsRequest

	roundTrip(t, "GetTagsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetTagsRequest.ContentKey differs")
	}
}

func TestGetTagsRequest_Header(t *testing.T) {
	want := GetTagsRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}()}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00}

	var got GetTagsRequest

	roundTrip(t, "GetTagsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("GetTagsRequest.ContentKey differs")
	}
}

func TestGetTagsResponse_NoHeader(t *testing.T) {
	want := GetTagsResponse{TagIds: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetTagsResponse

	roundTrip(t, "GetTagsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TagIds.Equals(want.TagIds)) {
		t.Fatalf("GetTagsResponse.TagIds differs")
	}
}

func TestGetTagsResponse_Header(t *testing.T) {
	want := GetTagsResponse{TagIds: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got GetTagsResponse

	roundTrip(t, "GetTagsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.TagIds.Equals(want.TagIds)) {
		t.Fatalf("GetTagsResponse.TagIds differs")
	}
}

func TestTagContentRequest_NoHeader(t *testing.T) {
	want := TagContentRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}(), NewTagIds: types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}}

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00}

	var got TagContentRequest

	roundTrip(t, "TagContentRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("TagContentRequest.ContentKey differs")
	}
	if !(got.NewTagIds.Equals(want.NewTagIds)) {
		t.Fatalf("TagContentRequest.NewTagIds differs")
	}
}

func TestTagContentRequest_Header(t *testing.T) {
	want := TagContentRequest{ContentKey: func() user_storage_types.UserContentKey {
		v := user_storage_types.NewUserContentKey()
		v.TypeID = types.NewUInt32(2023)
		v.ContentID = types.NewUInt64(300000032)
		return v
	}(), NewTagIds: types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}}

	data := []byte{0x00, 0x0c, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x20, 0xa3, 0xe1, 0x11, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00}

	var got TagContentRequest

	roundTrip(t, "TagContentRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.ContentKey.Equals(want.ContentKey)) {
		t.Fatalf("TagContentRequest.ContentKey differs")
	}
	if !(got.NewTagIds.Equals(want.NewTagIds)) {
		t.Fatalf("TagContentRequest.NewTagIds differs")
	}
}

func TestSearchContentsByPlayersRequest_NoHeader(t *testing.T) {
	want := SearchContentsByPlayersRequest{Pids: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}, Query: func() user_storage_types.UserStorageQuery {
		v := user_storage_types.NewUserStorageQuery()
		v.TypeID = types.NewUInt32(5050)
		v.QueryID = types.NewUInt32(6059)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(21), Length: types.NewUInt32(31)}
		v.Parameters = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(13122)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
			return v
		}()}
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got SearchContentsByPlayersRequest

	roundTrip(t, "SearchContentsByPlayersRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Pids.Equals(want.Pids)) {
		t.Fatalf("SearchContentsByPlayersRequest.Pids differs")
	}
	if !(got.Query.Equals(want.Query)) {
		t.Fatalf("SearchContentsByPlayersRequest.Query differs")
	}
}

func TestSearchContentsByPlayersRequest_Header(t *testing.T) {
	want := SearchContentsByPlayersRequest{Pids: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}, Query: func() user_storage_types.UserStorageQuery {
		v := user_storage_types.NewUserStorageQuery()
		v.TypeID = types.NewUInt32(5050)
		v.QueryID = types.NewUInt32(6059)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(21), Length: types.NewUInt32(31)}
		v.Parameters = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(13122)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
			return v
		}()}
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got SearchContentsByPlayersRequest

	roundTrip(t, "SearchContentsByPlayersRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Pids.Equals(want.Pids)) {
		t.Fatalf("SearchContentsByPlayersRequest.Pids differs")
	}
	if !(got.Query.Equals(want.Query)) {
		t.Fatalf("SearchContentsByPlayersRequest.Query differs")
	}
}

func TestSearchContentsByPlayersResponse_NoHeader(t *testing.T) {
	want := SearchContentsByPlayersResponse{SearchResults: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got SearchContentsByPlayersResponse

	roundTrip(t, "SearchContentsByPlayersResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchContentsByPlayersResponse.SearchResults differs")
	}
}

func TestSearchContentsByPlayersResponse_Header(t *testing.T) {
	want := SearchContentsByPlayersResponse{SearchResults: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got SearchContentsByPlayersResponse

	roundTrip(t, "SearchContentsByPlayersResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchContentsByPlayersResponse.SearchResults differs")
	}
}

func TestSearchContentsByPlayersWithTotalRequest_NoHeader(t *testing.T) {
	want := SearchContentsByPlayersWithTotalRequest{Pids: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}, Query: func() user_storage_types.UserStorageQuery {
		v := user_storage_types.NewUserStorageQuery()
		v.TypeID = types.NewUInt32(5050)
		v.QueryID = types.NewUInt32(6059)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(21), Length: types.NewUInt32(31)}
		v.Parameters = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(13122)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
			return v
		}()}
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got SearchContentsByPlayersWithTotalRequest

	roundTrip(t, "SearchContentsByPlayersWithTotalRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Pids.Equals(want.Pids)) {
		t.Fatalf("SearchContentsByPlayersWithTotalRequest.Pids differs")
	}
	if !(got.Query.Equals(want.Query)) {
		t.Fatalf("SearchContentsByPlayersWithTotalRequest.Query differs")
	}
}

func TestSearchContentsByPlayersWithTotalRequest_Header(t *testing.T) {
	want := SearchContentsByPlayersWithTotalRequest{Pids: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}, Query: func() user_storage_types.UserStorageQuery {
		v := user_storage_types.NewUserStorageQuery()
		v.TypeID = types.NewUInt32(5050)
		v.QueryID = types.NewUInt32(6059)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(21), Length: types.NewUInt32(31)}
		v.Parameters = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(13122)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
			return v
		}()}
		return v
	}()}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got SearchContentsByPlayersWithTotalRequest

	roundTrip(t, "SearchContentsByPlayersWithTotalRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Pids.Equals(want.Pids)) {
		t.Fatalf("SearchContentsByPlayersWithTotalRequest.Pids differs")
	}
	if !(got.Query.Equals(want.Query)) {
		t.Fatalf("SearchContentsByPlayersWithTotalRequest.Query differs")
	}
}

func TestSearchContentsByPlayersWithTotalResponse_NoHeader(t *testing.T) {
	want := SearchContentsByPlayersWithTotalResponse{SearchResults: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}, TotalResults: types.NewUInt32(26239)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x7f, 0x66, 0x00, 0x00}

	var got SearchContentsByPlayersWithTotalResponse

	roundTrip(t, "SearchContentsByPlayersWithTotalResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchContentsByPlayersWithTotalResponse.SearchResults differs")
	}
	if !(got.TotalResults.Equals(want.TotalResults)) {
		t.Fatalf("SearchContentsByPlayersWithTotalResponse.TotalResults differs")
	}
}

func TestSearchContentsByPlayersWithTotalResponse_Header(t *testing.T) {
	want := SearchContentsByPlayersWithTotalResponse{SearchResults: types.List[user_storage_types.UserContent]{func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(4041)
			v.ContentID = types.NewUInt64(500000046)
			return v
		}()
		v.Pid = types.NewUInt32(6059)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() user_storage_types.UserContent {
		v := user_storage_types.NewUserContent()
		v.Key = func() user_storage_types.UserContentKey {
			v := user_storage_types.NewUserContentKey()
			v.TypeID = types.NewUInt32(16149)
			v.ContentID = types.NewUInt64(1700000130)
			return v
		}()
		v.Pid = types.NewUInt32(18167)
		v.Properties = types.List[user_storage_types.ContentProperty]{func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() user_storage_types.ContentProperty {
			v := user_storage_types.NewContentProperty()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}, TotalResults: types.NewUInt32(26239)}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x2e, 0x65, 0xcd, 0x1d, 0x00, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x82, 0xf1, 0x53, 0x65, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x7f, 0x66, 0x00, 0x00}

	var got SearchContentsByPlayersWithTotalResponse

	roundTrip(t, "SearchContentsByPlayersWithTotalResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.SearchResults.Equals(want.SearchResults)) {
		t.Fatalf("SearchContentsByPlayersWithTotalResponse.SearchResults differs")
	}
	if !(got.TotalResults.Equals(want.TotalResults)) {
		t.Fatalf("SearchContentsByPlayersWithTotalResponse.TotalResults differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodSearchContents != 1 {
		t.Fatalf("MethodSearchContents = %d, want 1", MethodSearchContents)
	}
	if MethodSearchContentsWithTotal != 2 {
		t.Fatalf("MethodSearchContentsWithTotal = %d, want 2", MethodSearchContentsWithTotal)
	}
	if MethodDeleteContent != 3 {
		t.Fatalf("MethodDeleteContent = %d, want 3", MethodDeleteContent)
	}
	if MethodSaveMetaData != 4 {
		t.Fatalf("MethodSaveMetaData = %d, want 4", MethodSaveMetaData)
	}
	if MethodSaveContentDB != 5 {
		t.Fatalf("MethodSaveContentDB = %d, want 5", MethodSaveContentDB)
	}
	if MethodSaveContentAndGetUploadInfo != 6 {
		t.Fatalf("MethodSaveContentAndGetUploadInfo = %d, want 6", MethodSaveContentAndGetUploadInfo)
	}
	if MethodUploadEnd != 7 {
		t.Fatalf("MethodUploadEnd = %d, want 7", MethodUploadEnd)
	}
	if MethodGetContentDB != 8 {
		t.Fatalf("MethodGetContentDB = %d, want 8", MethodGetContentDB)
	}
	if MethodGetContentURL != 9 {
		t.Fatalf("MethodGetContentURL = %d, want 9", MethodGetContentURL)
	}
	if MethodGetSlotCount != 10 {
		t.Fatalf("MethodGetSlotCount = %d, want 10", MethodGetSlotCount)
	}
	if MethodGetMetaData != 11 {
		t.Fatalf("MethodGetMetaData = %d, want 11", MethodGetMetaData)
	}
	if MethodLike != 12 {
		t.Fatalf("MethodLike = %d, want 12", MethodLike)
	}
	if MethodUnlike != 13 {
		t.Fatalf("MethodUnlike = %d, want 13", MethodUnlike)
	}
	if MethodIsLiked != 14 {
		t.Fatalf("MethodIsLiked = %d, want 14", MethodIsLiked)
	}
	if MethodGetFavourites != 15 {
		t.Fatalf("MethodGetFavourites = %d, want 15", MethodGetFavourites)
	}
	if MethodMakeFavourite != 16 {
		t.Fatalf("MethodMakeFavourite = %d, want 16", MethodMakeFavourite)
	}
	if MethodRemoveFromFavourites != 17 {
		t.Fatalf("MethodRemoveFromFavourites = %d, want 17", MethodRemoveFromFavourites)
	}
	if MethodReportInappropriate != 18 {
		t.Fatalf("MethodReportInappropriate = %d, want 18", MethodReportInappropriate)
	}
	if MethodIncrementPlayCount != 19 {
		t.Fatalf("MethodIncrementPlayCount = %d, want 19", MethodIncrementPlayCount)
	}
	if MethodUpdateCustomStat != 20 {
		t.Fatalf("MethodUpdateCustomStat = %d, want 20", MethodUpdateCustomStat)
	}
	if MethodGetOwnContents != 21 {
		t.Fatalf("MethodGetOwnContents = %d, want 21", MethodGetOwnContents)
	}
	if MethodGetMostPopularTags != 22 {
		t.Fatalf("MethodGetMostPopularTags = %d, want 22", MethodGetMostPopularTags)
	}
	if MethodGetTags != 23 {
		t.Fatalf("MethodGetTags = %d, want 23", MethodGetTags)
	}
	if MethodTagContent != 24 {
		t.Fatalf("MethodTagContent = %d, want 24", MethodTagContent)
	}
	if MethodSearchContentsByPlayers != 25 {
		t.Fatalf("MethodSearchContentsByPlayers = %d, want 25", MethodSearchContentsByPlayers)
	}
	if MethodSearchContentsByPlayersWithTotal != 26 {
		t.Fatalf("MethodSearchContentsByPlayersWithTotal = %d, want 26", MethodSearchContentsByPlayersWithTotal)
	}
	if ProtocolID != 53 {
		t.Fatalf("ProtocolID = %d, want 53", ProtocolID)
	}
}
