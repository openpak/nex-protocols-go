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

func TestPlayerStatUpdate_NoHeader(t *testing.T) {
	want := func() PlayerStatUpdate {
		v := NewPlayerStatUpdate()
		v.BoardID = types.NewUInt32(2023)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(8077)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v9")}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(11104)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-17)}
			return v
		}()}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x39, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x01, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got PlayerStatUpdate

	roundTrip(t, "PlayerStatUpdate", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PlayerStatUpdate differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPlayerStatUpdate_Header(t *testing.T) {
	want := func() PlayerStatUpdate {
		v := NewPlayerStatUpdate()
		v.BoardID = types.NewUInt32(2023)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(8077)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v9")}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(11104)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-17)}
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x35, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x39, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x01, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got PlayerStatUpdate

	roundTrip(t, "PlayerStatUpdate", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PlayerStatUpdate differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestStatboardQuery_NoHeader(t *testing.T) {
	want := func() StatboardQuery {
		v := NewStatboardQuery()
		v.BoardID = types.NewUInt32(2023)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.ResetFrequency = types.NewUInt32(6059)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		v.SortCriterias = types.List[PlayerStatSortCriteria]{func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(12113)
			v.SortOrder = types.NewUInt32(13122)
			return v
		}(), func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(15140)
			v.SortOrder = types.NewUInt32(16149)
			return v
		}()}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00}

	var got StatboardQuery

	roundTrip(t, "StatboardQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded StatboardQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestStatboardQuery_Header(t *testing.T) {
	want := func() StatboardQuery {
		v := NewStatboardQuery()
		v.BoardID = types.NewUInt32(2023)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.ResetFrequency = types.NewUInt32(6059)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		v.SortCriterias = types.List[PlayerStatSortCriteria]{func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(12113)
			v.SortOrder = types.NewUInt32(13122)
			return v
		}(), func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(15140)
			v.SortOrder = types.NewUInt32(16149)
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x3e, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00}

	var got StatboardQuery

	roundTrip(t, "StatboardQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded StatboardQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestLeaderboardQuery_NoHeader(t *testing.T) {
	want := func() LeaderboardQuery {
		v := NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.ResetFrequency = types.NewUInt32(4041)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00}

	var got LeaderboardQuery

	roundTrip(t, "LeaderboardQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded LeaderboardQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestLeaderboardQuery_Header(t *testing.T) {
	want := func() LeaderboardQuery {
		v := NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.ResetFrequency = types.NewUInt32(4041)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		return v
	}()

	data := []byte{0x00, 0x18, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00}

	var got LeaderboardQuery

	roundTrip(t, "LeaderboardQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded LeaderboardQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestLeaderboardQuery2_NoHeader(t *testing.T) {
	want := func() LeaderboardQuery2 {
		v := NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.ResetFrequency = types.NewUInt32(4041)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00}

	var got LeaderboardQuery2

	roundTrip(t, "LeaderboardQuery2", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded LeaderboardQuery2 differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestLeaderboardQuery2_Header(t *testing.T) {
	want := func() LeaderboardQuery2 {
		v := NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.ResetFrequency = types.NewUInt32(4041)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		return v
	}()

	data := []byte{0x00, 0x24, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00}

	var got LeaderboardQuery2

	roundTrip(t, "LeaderboardQuery2", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded LeaderboardQuery2 differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestStatboardResult_NoHeader(t *testing.T) {
	want := func() StatboardResult {
		v := NewStatboardResult()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.ResetFrequency = types.NewUInt32(4041)
		v.PlayerStatSets = types.List[PlayerStatSet]{func() PlayerStatSet {
			v := NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(7068)
			v.PlayerName = types.NewString("s8")
			v.SubmittedTime = types.NewDateTime(9000104)
			v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(12113)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
				return v
			}(), func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(15140)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
				return v
			}()}
			return v
		}(), func() PlayerStatSet {
			v := NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(18167)
			v.PlayerName = types.NewString("s19")
			v.SubmittedTime = types.NewDateTime(20000137)
			v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(23212)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-29)}
				return v
			}(), func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(26239)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v27")}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(30275)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(33302)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
			return v
		}()}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x89, 0x2d, 0x31, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xac, 0x5a, 0x00, 0x00, 0x01, 0xe3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f, 0x66, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x37, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got StatboardResult

	roundTrip(t, "StatboardResult", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded StatboardResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestStatboardResult_Header(t *testing.T) {
	want := func() StatboardResult {
		v := NewStatboardResult()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.ResetFrequency = types.NewUInt32(4041)
		v.PlayerStatSets = types.List[PlayerStatSet]{func() PlayerStatSet {
			v := NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(7068)
			v.PlayerName = types.NewString("s8")
			v.SubmittedTime = types.NewDateTime(9000104)
			v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(12113)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
				return v
			}(), func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(15140)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
				return v
			}()}
			return v
		}(), func() PlayerStatSet {
			v := NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(18167)
			v.PlayerName = types.NewString("s19")
			v.SubmittedTime = types.NewDateTime(20000137)
			v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(23212)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-29)}
				return v
			}(), func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(26239)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v27")}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(30275)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(33302)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0xaf, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x37, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x38, 0x00, 0xa8, 0x54, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x38, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x39, 0x00, 0x89, 0x2d, 0x31, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xac, 0x5a, 0x00, 0x00, 0x01, 0xe3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x7f, 0x66, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x37, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got StatboardResult

	roundTrip(t, "StatboardResult", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded StatboardResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestLeaderboardResult_NoHeader(t *testing.T) {
	want := func() LeaderboardResult {
		v := NewLeaderboardResult()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.ResetFrequency = types.NewUInt32(4041)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(5050)
		v.PlayerRanks = types.List[PlayerRank]{func() PlayerRank {
			v := NewPlayerRank()
			v.PlayerStatSet = func() PlayerStatSet {
				v := NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(9086)
				v.PlayerName = types.NewString("s10")
				v.SubmittedTime = types.NewDateTime(11000110)
				v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
					v := NewPropertyVariant()
					v.ID = types.NewUInt32(14131)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v15")}
					return v
				}(), func() PropertyVariant {
					v := NewPropertyVariant()
					v.ID = types.NewUInt32(17158)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-23)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(19176)
			v.Rank = types.NewUInt32(20185)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v21")}
			return v
		}(), func() PlayerRank {
			v := NewPlayerRank()
			v.PlayerStatSet = func() PlayerStatSet {
				v := NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(24221)
				v.PlayerName = types.NewString("s25")
				v.SubmittedTime = types.NewDateTime(26000155)
				v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
					v := NewPropertyVariant()
					v.ID = types.NewUInt32(29266)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-35)}
					return v
				}(), func() PropertyVariant {
					v := NewPropertyVariant()
					v.ID = types.NewUInt32(32293)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v33")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(34311)
			v.Rank = types.NewUInt32(35320)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-41)}
			return v
		}()}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x2e, 0xd9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x35, 0x00, 0x06, 0x43, 0x00, 0x00, 0x01, 0xe9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x31, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x35, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x52, 0x72, 0x00, 0x00, 0x01, 0xdd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x25, 0x7e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x33, 0x00, 0x07, 0x86, 0x00, 0x00, 0xf8, 0x89, 0x00, 0x00, 0x01, 0xd7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got LeaderboardResult

	roundTrip(t, "LeaderboardResult", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded LeaderboardResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestLeaderboardResult_Header(t *testing.T) {
	want := func() LeaderboardResult {
		v := NewLeaderboardResult()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.ResetFrequency = types.NewUInt32(4041)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(5050)
		v.PlayerRanks = types.List[PlayerRank]{func() PlayerRank {
			v := NewPlayerRank()
			v.PlayerStatSet = func() PlayerStatSet {
				v := NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(9086)
				v.PlayerName = types.NewString("s10")
				v.SubmittedTime = types.NewDateTime(11000110)
				v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
					v := NewPropertyVariant()
					v.ID = types.NewUInt32(14131)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v15")}
					return v
				}(), func() PropertyVariant {
					v := NewPropertyVariant()
					v.ID = types.NewUInt32(17158)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-23)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(19176)
			v.Rank = types.NewUInt32(20185)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v21")}
			return v
		}(), func() PlayerRank {
			v := NewPlayerRank()
			v.PlayerStatSet = func() PlayerStatSet {
				v := NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(24221)
				v.PlayerName = types.NewString("s25")
				v.SubmittedTime = types.NewDateTime(26000155)
				v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
					v := NewPropertyVariant()
					v.ID = types.NewUInt32(29266)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-35)}
					return v
				}(), func() PropertyVariant {
					v := NewPropertyVariant()
					v.ID = types.NewUInt32(32293)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v33")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(34311)
			v.Rank = types.NewUInt32(35320)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-41)}
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0xb8, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x30, 0x00, 0x2e, 0xd9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x35, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x01, 0xe9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x31, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x35, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x52, 0x72, 0x00, 0x00, 0x01, 0xdd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x25, 0x7e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x33, 0x00, 0x07, 0x86, 0x00, 0x00, 0xf8, 0x89, 0x00, 0x00, 0x01, 0xd7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	var got LeaderboardResult

	roundTrip(t, "LeaderboardResult", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded LeaderboardResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestStatboardHistoryQuery_NoHeader(t *testing.T) {
	want := func() StatboardHistoryQuery {
		v := NewStatboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}
		v.BoardID = types.NewUInt32(5050)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(7068), types.NewUInt32(8077)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(10095), types.NewUInt32(11104)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(36), Length: types.NewUInt32(46)}
		v.PlayerStats = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(15140)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(18167)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
			return v
		}()}
		v.DateRanges = types.List[DateRange]{func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(22000143)
			v.EndingDatetime = types.NewDateTime(23000146)
			return v
		}(), func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(25000152)
			v.EndingDatetime = types.NewDateTime(26000155)
			return v
		}()}
		v.SortCriterias = types.List[PlayerStatSortCriteria]{func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(29266)
			v.SortOrder = types.NewUInt32(30275)
			return v
		}(), func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(32293)
			v.SortOrder = types.NewUInt32(33302)
			return v
		}()}
		return v
	}()

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0f, 0xb2, 0x4f, 0x01, 0x00, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0xd8, 0x78, 0x7d, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x52, 0x72, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x25, 0x7e, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00}

	var got StatboardHistoryQuery

	roundTrip(t, "StatboardHistoryQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded StatboardHistoryQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestStatboardHistoryQuery_Header(t *testing.T) {
	want := func() StatboardHistoryQuery {
		v := NewStatboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}
		v.BoardID = types.NewUInt32(5050)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(7068), types.NewUInt32(8077)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(10095), types.NewUInt32(11104)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(36), Length: types.NewUInt32(46)}
		v.PlayerStats = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(15140)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(18167)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
			return v
		}()}
		v.DateRanges = types.List[DateRange]{func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(22000143)
			v.EndingDatetime = types.NewDateTime(23000146)
			return v
		}(), func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(25000152)
			v.EndingDatetime = types.NewDateTime(26000155)
			return v
		}()}
		v.SortCriterias = types.List[PlayerStatSortCriteria]{func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(29266)
			v.SortOrder = types.NewUInt32(30275)
			return v
		}(), func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(32293)
			v.SortOrder = types.NewUInt32(33302)
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0xa7, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x0f, 0xb2, 0x4f, 0x01, 0x00, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xd8, 0x78, 0x7d, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x52, 0x72, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x25, 0x7e, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00}

	var got StatboardHistoryQuery

	roundTrip(t, "StatboardHistoryQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded StatboardHistoryQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestStatboardHistoryAggregatedQuery_NoHeader(t *testing.T) {
	want := func() StatboardHistoryAggregatedQuery {
		v := NewStatboardHistoryAggregatedQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}
		v.BoardID = types.NewUInt32(5050)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(7068), types.NewUInt32(8077)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(10095), types.NewUInt32(11104)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(36), Length: types.NewUInt32(46)}
		v.PlayerStats = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(15140)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(18167)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
			return v
		}()}
		v.DateRanges = types.List[DateRange]{func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(22000143)
			v.EndingDatetime = types.NewDateTime(23000146)
			return v
		}(), func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(25000152)
			v.EndingDatetime = types.NewDateTime(26000155)
			return v
		}()}
		v.SortCriterias = types.List[PlayerStatSortCriteria]{func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(29266)
			v.SortOrder = types.NewUInt32(30275)
			return v
		}(), func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(32293)
			v.SortOrder = types.NewUInt32(33302)
			return v
		}()}
		v.FilterOption = types.NewUInt32(34311)
		return v
	}()

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0f, 0xb2, 0x4f, 0x01, 0x00, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0xd8, 0x78, 0x7d, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x52, 0x72, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x25, 0x7e, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00}

	var got StatboardHistoryAggregatedQuery

	roundTrip(t, "StatboardHistoryAggregatedQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded StatboardHistoryAggregatedQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestStatboardHistoryAggregatedQuery_Header(t *testing.T) {
	want := func() StatboardHistoryAggregatedQuery {
		v := NewStatboardHistoryAggregatedQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}
		v.BoardID = types.NewUInt32(5050)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(7068), types.NewUInt32(8077)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(10095), types.NewUInt32(11104)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(36), Length: types.NewUInt32(46)}
		v.PlayerStats = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(15140)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(18167)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
			return v
		}()}
		v.DateRanges = types.List[DateRange]{func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(22000143)
			v.EndingDatetime = types.NewDateTime(23000146)
			return v
		}(), func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(25000152)
			v.EndingDatetime = types.NewDateTime(26000155)
			return v
		}()}
		v.SortCriterias = types.List[PlayerStatSortCriteria]{func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(29266)
			v.SortOrder = types.NewUInt32(30275)
			return v
		}(), func() PlayerStatSortCriteria {
			v := NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(32293)
			v.SortOrder = types.NewUInt32(33302)
			return v
		}()}
		v.FilterOption = types.NewUInt32(34311)
		return v
	}()

	data := []byte{0x00, 0xab, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x0f, 0xb2, 0x4f, 0x01, 0x00, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xd8, 0x78, 0x7d, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x52, 0x72, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x25, 0x7e, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00}

	var got StatboardHistoryAggregatedQuery

	roundTrip(t, "StatboardHistoryAggregatedQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded StatboardHistoryAggregatedQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestLeaderboardHistoryQuery_NoHeader(t *testing.T) {
	want := func() LeaderboardHistoryQuery {
		v := NewLeaderboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(21), Length: types.NewUInt32(31)}
		v.DateRanges = types.List[DateRange]{func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(10000107)
			v.EndingDatetime = types.NewDateTime(11000110)
			return v
		}(), func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(13000116)
			v.EndingDatetime = types.NewDateTime(14000119)
			return v
		}()}
		v.FilterOption = types.NewUInt32(15140)
		return v
	}()

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xd9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0xb4, 0x5d, 0xc6, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x9f, 0xd5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00}

	var got LeaderboardHistoryQuery

	roundTrip(t, "LeaderboardHistoryQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded LeaderboardHistoryQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestLeaderboardHistoryQuery_Header(t *testing.T) {
	want := func() LeaderboardHistoryQuery {
		v := NewLeaderboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(3032), types.NewUInt32(4041)}
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(21), Length: types.NewUInt32(31)}
		v.DateRanges = types.List[DateRange]{func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(10000107)
			v.EndingDatetime = types.NewDateTime(11000110)
			return v
		}(), func() DateRange {
			v := NewDateRange()
			v.StartingDatetime = types.NewDateTime(13000116)
			v.EndingDatetime = types.NewDateTime(14000119)
			return v
		}()}
		v.FilterOption = types.NewUInt32(15140)
		return v
	}()

	data := []byte{0x00, 0x53, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xd9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xb4, 0x5d, 0xc6, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x9f, 0xd5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00}

	var got LeaderboardHistoryQuery

	roundTrip(t, "LeaderboardHistoryQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded LeaderboardHistoryQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPopulationStatQuery_NoHeader(t *testing.T) {
	want := func() PopulationStatQuery {
		v := NewPopulationStatQuery()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00}

	var got PopulationStatQuery

	roundTrip(t, "PopulationStatQuery", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PopulationStatQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPopulationStatQuery_Header(t *testing.T) {
	want := func() PopulationStatQuery {
		v := NewPopulationStatQuery()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}
		return v
	}()

	data := []byte{0x00, 0x14, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00}

	var got PopulationStatQuery

	roundTrip(t, "PopulationStatQuery", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PopulationStatQuery differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPopulationStatResult_NoHeader(t *testing.T) {
	want := func() PopulationStatResult {
		v := NewPopulationStatResult()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.StatID = types.NewUInt32(4041)
		v.Sum = types.NewDouble(5.5)
		v.Average = types.NewDouble(6.5)
		v.StandardDeviation = types.NewDouble(7.5)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1a, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x40}

	var got PopulationStatResult

	roundTrip(t, "PopulationStatResult", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PopulationStatResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPopulationStatResult_Header(t *testing.T) {
	want := func() PopulationStatResult {
		v := NewPopulationStatResult()
		v.BoardID = types.NewUInt32(2023)
		v.ContextID = types.NewUInt32(3032)
		v.StatID = types.NewUInt32(4041)
		v.Sum = types.NewDouble(5.5)
		v.Average = types.NewDouble(6.5)
		v.StandardDeviation = types.NewDouble(7.5)
		return v
	}()

	data := []byte{0x00, 0x24, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1a, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x40}

	var got PopulationStatResult

	roundTrip(t, "PopulationStatResult", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PopulationStatResult differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPlayerStatSortCriteria_NoHeader(t *testing.T) {
	want := func() PlayerStatSortCriteria {
		v := NewPlayerStatSortCriteria()
		v.SortStatID = types.NewUInt32(2023)
		v.SortOrder = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got PlayerStatSortCriteria

	roundTrip(t, "PlayerStatSortCriteria", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PlayerStatSortCriteria differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPlayerStatSortCriteria_Header(t *testing.T) {
	want := func() PlayerStatSortCriteria {
		v := NewPlayerStatSortCriteria()
		v.SortStatID = types.NewUInt32(2023)
		v.SortOrder = types.NewUInt32(3032)
		return v
	}()

	data := []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00}

	var got PlayerStatSortCriteria

	roundTrip(t, "PlayerStatSortCriteria", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PlayerStatSortCriteria differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPlayerStatSet_NoHeader(t *testing.T) {
	want := func() PlayerStatSet {
		v := NewPlayerStatSet()
		v.PlayerPID = types.NewUInt32(2023)
		v.PlayerName = types.NewString("s3")
		v.SubmittedTime = types.NewDateTime(4000089)
		v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(7068)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-13)}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}()}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x59, 0x09, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x01, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00}

	var got PlayerStatSet

	roundTrip(t, "PlayerStatSet", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PlayerStatSet differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPlayerStatSet_Header(t *testing.T) {
	want := func() PlayerStatSet {
		v := NewPlayerStatSet()
		v.PlayerPID = types.NewUInt32(2023)
		v.PlayerName = types.NewString("s3")
		v.SubmittedTime = types.NewDateTime(4000089)
		v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(7068)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-13)}
			return v
		}(), func() PropertyVariant {
			v := NewPropertyVariant()
			v.ID = types.NewUInt32(10095)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v11")}
			return v
		}()}
		return v
	}()

	data := []byte{0x00, 0x37, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x03, 0x00, 0x73, 0x33, 0x00, 0x59, 0x09, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x01, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x31, 0x00}

	var got PlayerStatSet

	roundTrip(t, "PlayerStatSet", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PlayerStatSet differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPlayerRank_NoHeader(t *testing.T) {
	want := func() PlayerRank {
		v := NewPlayerRank()
		v.PlayerStatSet = func() PlayerStatSet {
			v := NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(3032)
			v.PlayerName = types.NewString("s4")
			v.SubmittedTime = types.NewDateTime(5000092)
			v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(8077)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v9")}
				return v
			}(), func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(11104)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-17)}
				return v
			}()}
			return v
		}()
		v.RankStatus = types.NewUInt32(13122)
		v.Rank = types.NewUInt32(14131)
		v.Score = types.Variant{TypeID: 4, Type: types.NewString("v15")}
		return v
	}()

	data := []byte{0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x9c, 0x4b, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x39, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x01, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x35, 0x00}

	var got PlayerRank

	roundTrip(t, "PlayerRank", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PlayerRank differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPlayerRank_Header(t *testing.T) {
	want := func() PlayerRank {
		v := NewPlayerRank()
		v.PlayerStatSet = func() PlayerStatSet {
			v := NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(3032)
			v.PlayerName = types.NewString("s4")
			v.SubmittedTime = types.NewDateTime(5000092)
			v.Stats = types.List[PropertyVariant]{func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(8077)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v9")}
				return v
			}(), func() PropertyVariant {
				v := NewPropertyVariant()
				v.ID = types.NewUInt32(11104)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-17)}
				return v
			}()}
			return v
		}()
		v.RankStatus = types.NewUInt32(13122)
		v.Rank = types.NewUInt32(14131)
		v.Score = types.Variant{TypeID: 4, Type: types.NewString("v15")}
		return v
	}()

	data := []byte{0x00, 0x4a, 0x00, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x03, 0x00, 0x73, 0x34, 0x00, 0x9c, 0x4b, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x39, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x01, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x35, 0x00}

	var got PlayerRank

	roundTrip(t, "PlayerRank", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PlayerRank differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestDateRange_NoHeader(t *testing.T) {
	want := func() DateRange {
		v := NewDateRange()
		v.StartingDatetime = types.NewDateTime(2000083)
		v.EndingDatetime = types.NewDateTime(3000086)
		return v
	}()

	data := []byte{0xd3, 0x84, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0xc7, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got DateRange

	roundTrip(t, "DateRange", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded DateRange differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestDateRange_Header(t *testing.T) {
	want := func() DateRange {
		v := NewDateRange()
		v.StartingDatetime = types.NewDateTime(2000083)
		v.EndingDatetime = types.NewDateTime(3000086)
		return v
	}()

	data := []byte{0x00, 0x10, 0x00, 0x00, 0x00, 0xd3, 0x84, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0xc7, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got DateRange

	roundTrip(t, "DateRange", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded DateRange differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPropertyVariant_NoHeader(t *testing.T) {
	want := func() PropertyVariant {
		v := NewPropertyVariant()
		v.ID = types.NewUInt32(2023)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v3")}
		return v
	}()

	data := []byte{0xe7, 0x07, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x33, 0x00}

	var got PropertyVariant

	roundTrip(t, "PropertyVariant", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PropertyVariant differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}

func TestPropertyVariant_Header(t *testing.T) {
	want := func() PropertyVariant {
		v := NewPropertyVariant()
		v.ID = types.NewUInt32(2023)
		v.Value = types.Variant{TypeID: 4, Type: types.NewString("v3")}
		return v
	}()

	data := []byte{0x00, 0x0a, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x04, 0x03, 0x00, 0x76, 0x33, 0x00}

	var got PropertyVariant

	roundTrip(t, "PropertyVariant", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !got.Equals(want) {
		t.Fatalf("decoded PropertyVariant differs from expected:\n%s\n%s", got, want)
	}

	_ = got.Copy()
	_ = got.String()
}
