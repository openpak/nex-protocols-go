package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
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

func TestWriteStatsRequest_NoHeader(t *testing.T) {
	want := WriteStatsRequest{PlayerStatUpdates: types.List[player_stats_types.PlayerStatUpdate]{func() player_stats_types.PlayerStatUpdate {
		v := player_stats_types.NewPlayerStatUpdate()
		v.BoardID = types.NewUInt32(3032)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}
		v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() player_stats_types.PlayerStatUpdate {
		v := player_stats_types.NewPlayerStatUpdate()
		v.BoardID = types.NewUInt32(15140)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(17158), types.NewUInt32(18167)}
		v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got WriteStatsRequest

	roundTrip(t, "WriteStatsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerStatUpdates.Equals(want.PlayerStatUpdates)) {
		t.Fatalf("WriteStatsRequest.PlayerStatUpdates differs")
	}
}

func TestWriteStatsRequest_Header(t *testing.T) {
	want := WriteStatsRequest{PlayerStatUpdates: types.List[player_stats_types.PlayerStatUpdate]{func() player_stats_types.PlayerStatUpdate {
		v := player_stats_types.NewPlayerStatUpdate()
		v.BoardID = types.NewUInt32(3032)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(5050), types.NewUInt32(6059)}
		v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(9086)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-15)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(12113)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v13")}
			return v
		}()}
		return v
	}(), func() player_stats_types.PlayerStatUpdate {
		v := player_stats_types.NewPlayerStatUpdate()
		v.BoardID = types.NewUInt32(15140)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(17158), types.NewUInt32(18167)}
		v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(21194)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(24221)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x01, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x33, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00}

	var got WriteStatsRequest

	roundTrip(t, "WriteStatsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerStatUpdates.Equals(want.PlayerStatUpdates)) {
		t.Fatalf("WriteStatsRequest.PlayerStatUpdates differs")
	}
}

func TestReadStatsByPlayersRequest_NoHeader(t *testing.T) {
	want := ReadStatsByPlayersRequest{PlayerPIDs: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}, Queries: types.List[player_stats_types.StatboardQuery]{func() player_stats_types.StatboardQuery {
		v := player_stats_types.NewStatboardQuery()
		v.BoardID = types.NewUInt32(6059)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		v.ResetFrequency = types.NewUInt32(10095)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(16149)
			v.SortOrder = types.NewUInt32(17158)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(19176)
			v.SortOrder = types.NewUInt32(20185)
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardQuery {
		v := player_stats_types.NewStatboardQuery()
		v.BoardID = types.NewUInt32(22203)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(24221), types.NewUInt32(25230)}
		v.ResetFrequency = types.NewUInt32(26239)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(28257), types.NewUInt32(29266)}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(32293)
			v.SortOrder = types.NewUInt32(33302)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(35320)
			v.SortOrder = types.NewUInt32(36329)
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xbb, 0x56, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x7f, 0x66, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x61, 0x6e, 0x00, 0x00, 0x52, 0x72, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x25, 0x7e, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00}

	var got ReadStatsByPlayersRequest

	roundTrip(t, "ReadStatsByPlayersRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerPIDs.Equals(want.PlayerPIDs)) {
		t.Fatalf("ReadStatsByPlayersRequest.PlayerPIDs differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadStatsByPlayersRequest.Queries differs")
	}
}

func TestReadStatsByPlayersRequest_Header(t *testing.T) {
	want := ReadStatsByPlayersRequest{PlayerPIDs: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}, Queries: types.List[player_stats_types.StatboardQuery]{func() player_stats_types.StatboardQuery {
		v := player_stats_types.NewStatboardQuery()
		v.BoardID = types.NewUInt32(6059)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		v.ResetFrequency = types.NewUInt32(10095)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(16149)
			v.SortOrder = types.NewUInt32(17158)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(19176)
			v.SortOrder = types.NewUInt32(20185)
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardQuery {
		v := player_stats_types.NewStatboardQuery()
		v.BoardID = types.NewUInt32(22203)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(24221), types.NewUInt32(25230)}
		v.ResetFrequency = types.NewUInt32(26239)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(28257), types.NewUInt32(29266)}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(32293)
			v.SortOrder = types.NewUInt32(33302)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(35320)
			v.SortOrder = types.NewUInt32(36329)
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0xbb, 0x56, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x7f, 0x66, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x61, 0x6e, 0x00, 0x00, 0x52, 0x72, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x25, 0x7e, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00}

	var got ReadStatsByPlayersRequest

	roundTrip(t, "ReadStatsByPlayersRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerPIDs.Equals(want.PlayerPIDs)) {
		t.Fatalf("ReadStatsByPlayersRequest.PlayerPIDs differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadStatsByPlayersRequest.Queries differs")
	}
}

func TestReadStatsByPlayersResponse_NoHeader(t *testing.T) {
	want := ReadStatsByPlayersResponse{Results: types.List[player_stats_types.StatboardResult]{func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(8077)
			v.PlayerName = types.NewString("s9")
			v.SubmittedTime = types.NewDateTime(10000107)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(13122)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(16149)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(19176)
			v.PlayerName = types.NewString("s20")
			v.SubmittedTime = types.NewDateTime(21000140)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(24221)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(27248)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-33)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(31284)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-37)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(34311)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v35")}
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(37338)
		v.ContextID = types.NewUInt32(38347)
		v.ResetFrequency = types.NewUInt32(39356)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(42383)
			v.PlayerName = types.NewString("s43")
			v.SubmittedTime = types.NewDateTime(44000209)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(47428)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-53)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(50455)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v51")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(53482)
			v.PlayerName = types.NewString("s54")
			v.SubmittedTime = types.NewDateTime(55000242)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(58527)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v59")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(61554)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-67)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(65590)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-71)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(68617)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v69")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x01, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x01, 0xdb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x07, 0x86, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x35, 0x00, 0xda, 0x91, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x33, 0x00, 0xd1, 0x63, 0x9f, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x44, 0xb9, 0x00, 0x00, 0x01, 0xcb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x17, 0xc5, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x31, 0x00, 0xea, 0xd0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x35, 0x34, 0x00, 0xb2, 0x3c, 0x47, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9f, 0xe4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x39, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x01, 0xbd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x36, 0x00, 0x01, 0x00, 0x01, 0xb9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x09, 0x0c, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x39, 0x00}

	var got ReadStatsByPlayersResponse

	roundTrip(t, "ReadStatsByPlayersResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadStatsByPlayersResponse.Results differs")
	}
}

func TestReadStatsByPlayersResponse_Header(t *testing.T) {
	want := ReadStatsByPlayersResponse{Results: types.List[player_stats_types.StatboardResult]{func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(8077)
			v.PlayerName = types.NewString("s9")
			v.SubmittedTime = types.NewDateTime(10000107)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(13122)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(16149)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(19176)
			v.PlayerName = types.NewString("s20")
			v.SubmittedTime = types.NewDateTime(21000140)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(24221)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(27248)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-33)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(31284)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-37)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(34311)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v35")}
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(37338)
		v.ContextID = types.NewUInt32(38347)
		v.ResetFrequency = types.NewUInt32(39356)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(42383)
			v.PlayerName = types.NewString("s43")
			v.SubmittedTime = types.NewDateTime(44000209)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(47428)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-53)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(50455)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v51")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(53482)
			v.PlayerName = types.NewString("s54")
			v.SubmittedTime = types.NewDateTime(55000242)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(58527)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v59")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(61554)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-67)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(65590)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-71)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(68617)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v69")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xaf, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x37, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x01, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x01, 0xdb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x35, 0x00, 0x00, 0xb0, 0x00, 0x00, 0x00, 0xda, 0x91, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x33, 0x00, 0xd1, 0x63, 0x9f, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x44, 0xb9, 0x00, 0x00, 0x01, 0xcb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x17, 0xc5, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x31, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0xea, 0xd0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x35, 0x34, 0x00, 0xb2, 0x3c, 0x47, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9f, 0xe4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x39, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x01, 0xbd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x36, 0x00, 0x01, 0x00, 0x01, 0xb9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x39, 0x00}

	var got ReadStatsByPlayersResponse

	roundTrip(t, "ReadStatsByPlayersResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadStatsByPlayersResponse.Results differs")
	}
}

func TestReadLeaderboardsNearPlayerRequest_NoHeader(t *testing.T) {
	want := ReadLeaderboardsNearPlayerRequest{PlayerPID: types.NewUInt32(1014), Count: types.NewUInt32(2023), Queries: types.List[player_stats_types.LeaderboardQuery]{func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResetFrequency = types.NewUInt32(7068)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		return v
	}(), func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(12113)
		v.ContextID = types.NewUInt32(13122)
		v.ResetFrequency = types.NewUInt32(14131)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(16149), types.NewUInt32(17158)}
		return v
	}()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00}

	var got ReadLeaderboardsNearPlayerRequest

	roundTrip(t, "ReadLeaderboardsNearPlayerRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerPID.Equals(want.PlayerPID)) {
		t.Fatalf("ReadLeaderboardsNearPlayerRequest.PlayerPID differs")
	}
	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("ReadLeaderboardsNearPlayerRequest.Count differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsNearPlayerRequest.Queries differs")
	}
}

func TestReadLeaderboardsNearPlayerRequest_Header(t *testing.T) {
	want := ReadLeaderboardsNearPlayerRequest{PlayerPID: types.NewUInt32(1014), Count: types.NewUInt32(2023), Queries: types.List[player_stats_types.LeaderboardQuery]{func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResetFrequency = types.NewUInt32(7068)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		return v
	}(), func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(12113)
		v.ContextID = types.NewUInt32(13122)
		v.ResetFrequency = types.NewUInt32(14131)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(16149), types.NewUInt32(17158)}
		return v
	}()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00}

	var got ReadLeaderboardsNearPlayerRequest

	roundTrip(t, "ReadLeaderboardsNearPlayerRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerPID.Equals(want.PlayerPID)) {
		t.Fatalf("ReadLeaderboardsNearPlayerRequest.PlayerPID differs")
	}
	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("ReadLeaderboardsNearPlayerRequest.Count differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsNearPlayerRequest.Queries differs")
	}
}

func TestReadLeaderboardsNearPlayerResponse_NoHeader(t *testing.T) {
	want := ReadLeaderboardsNearPlayerResponse{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsNearPlayerResponse

	roundTrip(t, "ReadLeaderboardsNearPlayerResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsNearPlayerResponse.Results differs")
	}
}

func TestReadLeaderboardsNearPlayerResponse_Header(t *testing.T) {
	want := ReadLeaderboardsNearPlayerResponse{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsNearPlayerResponse

	roundTrip(t, "ReadLeaderboardsNearPlayerResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsNearPlayerResponse.Results differs")
	}
}

func TestReadLeaderboardsByRankRequest_NoHeader(t *testing.T) {
	want := ReadLeaderboardsByRankRequest{StartingRank: types.NewUInt32(1014), Count: types.NewUInt32(2023), Queries: types.List[player_stats_types.LeaderboardQuery]{func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResetFrequency = types.NewUInt32(7068)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		return v
	}(), func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(12113)
		v.ContextID = types.NewUInt32(13122)
		v.ResetFrequency = types.NewUInt32(14131)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(16149), types.NewUInt32(17158)}
		return v
	}()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00}

	var got ReadLeaderboardsByRankRequest

	roundTrip(t, "ReadLeaderboardsByRankRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.StartingRank.Equals(want.StartingRank)) {
		t.Fatalf("ReadLeaderboardsByRankRequest.StartingRank differs")
	}
	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("ReadLeaderboardsByRankRequest.Count differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsByRankRequest.Queries differs")
	}
}

func TestReadLeaderboardsByRankRequest_Header(t *testing.T) {
	want := ReadLeaderboardsByRankRequest{StartingRank: types.NewUInt32(1014), Count: types.NewUInt32(2023), Queries: types.List[player_stats_types.LeaderboardQuery]{func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResetFrequency = types.NewUInt32(7068)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		return v
	}(), func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(12113)
		v.ContextID = types.NewUInt32(13122)
		v.ResetFrequency = types.NewUInt32(14131)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(16149), types.NewUInt32(17158)}
		return v
	}()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00}

	var got ReadLeaderboardsByRankRequest

	roundTrip(t, "ReadLeaderboardsByRankRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.StartingRank.Equals(want.StartingRank)) {
		t.Fatalf("ReadLeaderboardsByRankRequest.StartingRank differs")
	}
	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("ReadLeaderboardsByRankRequest.Count differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsByRankRequest.Queries differs")
	}
}

func TestReadLeaderboardsByRankResponse_NoHeader(t *testing.T) {
	want := ReadLeaderboardsByRankResponse{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsByRankResponse

	roundTrip(t, "ReadLeaderboardsByRankResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsByRankResponse.Results differs")
	}
}

func TestReadLeaderboardsByRankResponse_Header(t *testing.T) {
	want := ReadLeaderboardsByRankResponse{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsByRankResponse

	roundTrip(t, "ReadLeaderboardsByRankResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsByRankResponse.Results differs")
	}
}

func TestReadLeaderboardsByPlayersRequest_NoHeader(t *testing.T) {
	want := ReadLeaderboardsByPlayersRequest{PlayerPIDs: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}, Queries: types.List[player_stats_types.LeaderboardQuery]{func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(6059)
		v.ContextID = types.NewUInt32(7068)
		v.ResetFrequency = types.NewUInt32(8077)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(10095), types.NewUInt32(11104)}
		return v
	}(), func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(13122)
		v.ContextID = types.NewUInt32(14131)
		v.ResetFrequency = types.NewUInt32(15140)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(17158), types.NewUInt32(18167)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00}

	var got ReadLeaderboardsByPlayersRequest

	roundTrip(t, "ReadLeaderboardsByPlayersRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerPIDs.Equals(want.PlayerPIDs)) {
		t.Fatalf("ReadLeaderboardsByPlayersRequest.PlayerPIDs differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsByPlayersRequest.Queries differs")
	}
}

func TestReadLeaderboardsByPlayersRequest_Header(t *testing.T) {
	want := ReadLeaderboardsByPlayersRequest{PlayerPIDs: types.List[types.UInt32]{types.NewUInt32(2023), types.NewUInt32(3032)}, Queries: types.List[player_stats_types.LeaderboardQuery]{func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(6059)
		v.ContextID = types.NewUInt32(7068)
		v.ResetFrequency = types.NewUInt32(8077)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(10095), types.NewUInt32(11104)}
		return v
	}(), func() player_stats_types.LeaderboardQuery {
		v := player_stats_types.NewLeaderboardQuery()
		v.BoardID = types.NewUInt32(13122)
		v.ContextID = types.NewUInt32(14131)
		v.ResetFrequency = types.NewUInt32(15140)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(17158), types.NewUInt32(18167)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00}

	var got ReadLeaderboardsByPlayersRequest

	roundTrip(t, "ReadLeaderboardsByPlayersRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerPIDs.Equals(want.PlayerPIDs)) {
		t.Fatalf("ReadLeaderboardsByPlayersRequest.PlayerPIDs differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsByPlayersRequest.Queries differs")
	}
}

func TestReadLeaderboardsByPlayersResponse_NoHeader(t *testing.T) {
	want := ReadLeaderboardsByPlayersResponse{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsByPlayersResponse

	roundTrip(t, "ReadLeaderboardsByPlayersResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsByPlayersResponse.Results differs")
	}
}

func TestReadLeaderboardsByPlayersResponse_Header(t *testing.T) {
	want := ReadLeaderboardsByPlayersResponse{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsByPlayersResponse

	roundTrip(t, "ReadLeaderboardsByPlayersResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsByPlayersResponse.Results differs")
	}
}

func TestReadStatboardHistoryRequest_NoHeader(t *testing.T) {
	want := ReadStatboardHistoryRequest{Queries: types.List[player_stats_types.StatboardHistoryQuery]{func() player_stats_types.StatboardHistoryQuery {
		v := player_stats_types.NewStatboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.BoardID = types.NewUInt32(6059)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(11104), types.NewUInt32(12113)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(39), Length: types.NewUInt32(49)}
		v.PlayerStats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(16149)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(19176)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-25)}
			return v
		}()}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(23000146)
			v.EndingDatetime = types.NewDateTime(24000149)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(26000155)
			v.EndingDatetime = types.NewDateTime(27000158)
			return v
		}()}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(30275)
			v.SortOrder = types.NewUInt32(31284)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(33302)
			v.SortOrder = types.NewUInt32(34311)
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardHistoryQuery {
		v := player_stats_types.NewStatboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(37338), types.NewUInt32(38347)}
		v.BoardID = types.NewUInt32(39356)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(41374), types.NewUInt32(42383)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(44401), types.NewUInt32(45410)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(138), Length: types.NewUInt32(148)}
		v.PlayerStats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(49446)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-55)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(52473)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v53")}
			return v
		}()}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(56000245)
			v.EndingDatetime = types.NewDateTime(57000248)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(59000254)
			v.EndingDatetime = types.NewDateTime(60000257)
			return v
		}()}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(63572)
			v.SortOrder = types.NewUInt32(64581)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(66599)
			v.SortOrder = types.NewUInt32(67608)
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x01, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x95, 0x36, 0x6e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xda, 0x91, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x71, 0xad, 0x00, 0x00, 0x62, 0xb1, 0x00, 0x00, 0x8a, 0x00, 0x00, 0x00, 0x94, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x26, 0xc1, 0x00, 0x00, 0x01, 0xc9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf9, 0xcc, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x33, 0x00, 0x02, 0x00, 0x00, 0x00, 0xf5, 0x7e, 0x56, 0x03, 0x00, 0x00, 0x00, 0x00, 0x38, 0xc1, 0x65, 0x03, 0x00, 0x00, 0x00, 0x00, 0xbe, 0x45, 0x84, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01, 0x88, 0x93, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x54, 0xf8, 0x00, 0x00, 0x45, 0xfc, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x18, 0x08, 0x01, 0x00}

	var got ReadStatboardHistoryRequest

	roundTrip(t, "ReadStatboardHistoryRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadStatboardHistoryRequest.Queries differs")
	}
}

func TestReadStatboardHistoryRequest_Header(t *testing.T) {
	want := ReadStatboardHistoryRequest{Queries: types.List[player_stats_types.StatboardHistoryQuery]{func() player_stats_types.StatboardHistoryQuery {
		v := player_stats_types.NewStatboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.BoardID = types.NewUInt32(6059)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(11104), types.NewUInt32(12113)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(39), Length: types.NewUInt32(49)}
		v.PlayerStats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(16149)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(19176)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-25)}
			return v
		}()}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(23000146)
			v.EndingDatetime = types.NewDateTime(24000149)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(26000155)
			v.EndingDatetime = types.NewDateTime(27000158)
			return v
		}()}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(30275)
			v.SortOrder = types.NewUInt32(31284)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(33302)
			v.SortOrder = types.NewUInt32(34311)
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardHistoryQuery {
		v := player_stats_types.NewStatboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(37338), types.NewUInt32(38347)}
		v.BoardID = types.NewUInt32(39356)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(41374), types.NewUInt32(42383)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(44401), types.NewUInt32(45410)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(138), Length: types.NewUInt32(148)}
		v.PlayerStats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(49446)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-55)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(52473)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v53")}
			return v
		}()}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(56000245)
			v.EndingDatetime = types.NewDateTime(57000248)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(59000254)
			v.EndingDatetime = types.NewDateTime(60000257)
			return v
		}()}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(63572)
			v.SortOrder = types.NewUInt32(64581)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(66599)
			v.SortOrder = types.NewUInt32(67608)
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xa7, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x01, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x95, 0x36, 0x6e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00, 0x00, 0xa7, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xda, 0x91, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x71, 0xad, 0x00, 0x00, 0x62, 0xb1, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x8a, 0x00, 0x00, 0x00, 0x94, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x26, 0xc1, 0x00, 0x00, 0x01, 0xc9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf9, 0xcc, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x33, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xf5, 0x7e, 0x56, 0x03, 0x00, 0x00, 0x00, 0x00, 0x38, 0xc1, 0x65, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xbe, 0x45, 0x84, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01, 0x88, 0x93, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x54, 0xf8, 0x00, 0x00, 0x45, 0xfc, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x18, 0x08, 0x01, 0x00}

	var got ReadStatboardHistoryRequest

	roundTrip(t, "ReadStatboardHistoryRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadStatboardHistoryRequest.Queries differs")
	}
}

func TestReadStatboardHistoryResponse_NoHeader(t *testing.T) {
	want := ReadStatboardHistoryResponse{Results: types.List[player_stats_types.StatboardResult]{func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(8077)
			v.PlayerName = types.NewString("s9")
			v.SubmittedTime = types.NewDateTime(10000107)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(13122)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(16149)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(19176)
			v.PlayerName = types.NewString("s20")
			v.SubmittedTime = types.NewDateTime(21000140)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(24221)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(27248)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-33)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(31284)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-37)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(34311)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v35")}
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(37338)
		v.ContextID = types.NewUInt32(38347)
		v.ResetFrequency = types.NewUInt32(39356)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(42383)
			v.PlayerName = types.NewString("s43")
			v.SubmittedTime = types.NewDateTime(44000209)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(47428)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-53)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(50455)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v51")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(53482)
			v.PlayerName = types.NewString("s54")
			v.SubmittedTime = types.NewDateTime(55000242)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(58527)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v59")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(61554)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-67)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(65590)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-71)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(68617)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v69")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x01, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x01, 0xdb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x07, 0x86, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x35, 0x00, 0xda, 0x91, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x33, 0x00, 0xd1, 0x63, 0x9f, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x44, 0xb9, 0x00, 0x00, 0x01, 0xcb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x17, 0xc5, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x31, 0x00, 0xea, 0xd0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x35, 0x34, 0x00, 0xb2, 0x3c, 0x47, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9f, 0xe4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x39, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x01, 0xbd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x36, 0x00, 0x01, 0x00, 0x01, 0xb9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x09, 0x0c, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x39, 0x00}

	var got ReadStatboardHistoryResponse

	roundTrip(t, "ReadStatboardHistoryResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadStatboardHistoryResponse.Results differs")
	}
}

func TestReadStatboardHistoryResponse_Header(t *testing.T) {
	want := ReadStatboardHistoryResponse{Results: types.List[player_stats_types.StatboardResult]{func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(8077)
			v.PlayerName = types.NewString("s9")
			v.SubmittedTime = types.NewDateTime(10000107)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(13122)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(16149)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(19176)
			v.PlayerName = types.NewString("s20")
			v.SubmittedTime = types.NewDateTime(21000140)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(24221)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(27248)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-33)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(31284)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-37)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(34311)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v35")}
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(37338)
		v.ContextID = types.NewUInt32(38347)
		v.ResetFrequency = types.NewUInt32(39356)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(42383)
			v.PlayerName = types.NewString("s43")
			v.SubmittedTime = types.NewDateTime(44000209)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(47428)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-53)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(50455)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v51")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(53482)
			v.PlayerName = types.NewString("s54")
			v.SubmittedTime = types.NewDateTime(55000242)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(58527)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v59")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(61554)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-67)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(65590)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-71)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(68617)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v69")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xaf, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x37, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x01, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x01, 0xdb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x35, 0x00, 0x00, 0xb0, 0x00, 0x00, 0x00, 0xda, 0x91, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x33, 0x00, 0xd1, 0x63, 0x9f, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x44, 0xb9, 0x00, 0x00, 0x01, 0xcb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x17, 0xc5, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x31, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0xea, 0xd0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x35, 0x34, 0x00, 0xb2, 0x3c, 0x47, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9f, 0xe4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x39, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x01, 0xbd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x36, 0x00, 0x01, 0x00, 0x01, 0xb9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x39, 0x00}

	var got ReadStatboardHistoryResponse

	roundTrip(t, "ReadStatboardHistoryResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadStatboardHistoryResponse.Results differs")
	}
}

func TestReadLeaderboardHistoryRequest_NoHeader(t *testing.T) {
	want := ReadLeaderboardHistoryRequest{Queries: types.List[player_stats_types.LeaderboardHistoryQuery]{func() player_stats_types.LeaderboardHistoryQuery {
		v := player_stats_types.NewLeaderboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.BoardID = types.NewUInt32(6059)
		v.ContextID = types.NewUInt32(7068)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(24), Length: types.NewUInt32(34)}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(11000110)
			v.EndingDatetime = types.NewDateTime(12000113)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(14000119)
			v.EndingDatetime = types.NewDateTime(15000122)
			return v
		}()}
		v.FilterOption = types.NewUInt32(16149)
		return v
	}(), func() player_stats_types.LeaderboardHistoryQuery {
		v := player_stats_types.NewLeaderboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(19176), types.NewUInt32(20185)}
		v.BoardID = types.NewUInt32(21194)
		v.ContextID = types.NewUInt32(22203)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(69), Length: types.NewUInt32(79)}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(26000155)
			v.EndingDatetime = types.NewDateTime(27000158)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(29000164)
			v.EndingDatetime = types.NewDateTime(30000167)
			return v
		}()}
		v.FilterOption = types.NewUInt32(31284)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x2e, 0xd9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf7, 0x9f, 0xd5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3a, 0xe2, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0xbb, 0x56, 0x00, 0x00, 0x45, 0x00, 0x00, 0x00, 0x4f, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0xe4, 0x81, 0xba, 0x01, 0x00, 0x00, 0x00, 0x00, 0x27, 0xc4, 0xc9, 0x01, 0x00, 0x00, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00}

	var got ReadLeaderboardHistoryRequest

	roundTrip(t, "ReadLeaderboardHistoryRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardHistoryRequest.Queries differs")
	}
}

func TestReadLeaderboardHistoryRequest_Header(t *testing.T) {
	want := ReadLeaderboardHistoryRequest{Queries: types.List[player_stats_types.LeaderboardHistoryQuery]{func() player_stats_types.LeaderboardHistoryQuery {
		v := player_stats_types.NewLeaderboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.BoardID = types.NewUInt32(6059)
		v.ContextID = types.NewUInt32(7068)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(24), Length: types.NewUInt32(34)}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(11000110)
			v.EndingDatetime = types.NewDateTime(12000113)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(14000119)
			v.EndingDatetime = types.NewDateTime(15000122)
			return v
		}()}
		v.FilterOption = types.NewUInt32(16149)
		return v
	}(), func() player_stats_types.LeaderboardHistoryQuery {
		v := player_stats_types.NewLeaderboardHistoryQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(19176), types.NewUInt32(20185)}
		v.BoardID = types.NewUInt32(21194)
		v.ContextID = types.NewUInt32(22203)
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(69), Length: types.NewUInt32(79)}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(26000155)
			v.EndingDatetime = types.NewDateTime(27000158)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(29000164)
			v.EndingDatetime = types.NewDateTime(30000167)
			return v
		}()}
		v.FilterOption = types.NewUInt32(31284)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x53, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x2e, 0xd9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xf7, 0x9f, 0xd5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3a, 0xe2, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x00, 0x53, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0xbb, 0x56, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x45, 0x00, 0x00, 0x00, 0x4f, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xe4, 0x81, 0xba, 0x01, 0x00, 0x00, 0x00, 0x00, 0x27, 0xc4, 0xc9, 0x01, 0x00, 0x00, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00}

	var got ReadLeaderboardHistoryRequest

	roundTrip(t, "ReadLeaderboardHistoryRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardHistoryRequest.Queries differs")
	}
}

func TestReadLeaderboardHistoryResponse_NoHeader(t *testing.T) {
	want := ReadLeaderboardHistoryResponse{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardHistoryResponse

	roundTrip(t, "ReadLeaderboardHistoryResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardHistoryResponse.Results differs")
	}
}

func TestReadLeaderboardHistoryResponse_Header(t *testing.T) {
	want := ReadLeaderboardHistoryResponse{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardHistoryResponse

	roundTrip(t, "ReadLeaderboardHistoryResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardHistoryResponse.Results differs")
	}
}

func TestReadStatboardHistoryAggregatedRequest_NoHeader(t *testing.T) {
	want := ReadStatboardHistoryAggregatedRequest{Queries: types.List[player_stats_types.StatboardHistoryAggregatedQuery]{func() player_stats_types.StatboardHistoryAggregatedQuery {
		v := player_stats_types.NewStatboardHistoryAggregatedQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.BoardID = types.NewUInt32(6059)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(11104), types.NewUInt32(12113)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(39), Length: types.NewUInt32(49)}
		v.PlayerStats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(16149)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(19176)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-25)}
			return v
		}()}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(23000146)
			v.EndingDatetime = types.NewDateTime(24000149)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(26000155)
			v.EndingDatetime = types.NewDateTime(27000158)
			return v
		}()}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(30275)
			v.SortOrder = types.NewUInt32(31284)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(33302)
			v.SortOrder = types.NewUInt32(34311)
			return v
		}()}
		v.FilterOption = types.NewUInt32(35320)
		return v
	}(), func() player_stats_types.StatboardHistoryAggregatedQuery {
		v := player_stats_types.NewStatboardHistoryAggregatedQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(38347), types.NewUInt32(39356)}
		v.BoardID = types.NewUInt32(40365)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(42383), types.NewUInt32(43392)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(45410), types.NewUInt32(46419)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(141), Length: types.NewUInt32(151)}
		v.PlayerStats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(50455)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v51")}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(53482)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-59)}
			return v
		}()}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(57000248)
			v.EndingDatetime = types.NewDateTime(58000251)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(60000257)
			v.EndingDatetime = types.NewDateTime(61000260)
			return v
		}()}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(64581)
			v.SortOrder = types.NewUInt32(65590)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(67608)
			v.SortOrder = types.NewUInt32(68617)
			return v
		}()}
		v.FilterOption = types.NewUInt32(69626)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x01, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x95, 0x36, 0x6e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00, 0xf8, 0x89, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x80, 0xa9, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x62, 0xb1, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x8d, 0x00, 0x00, 0x00, 0x97, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x17, 0xc5, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x31, 0x00, 0xea, 0xd0, 0x00, 0x00, 0x01, 0xc5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x38, 0xc1, 0x65, 0x03, 0x00, 0x00, 0x00, 0x00, 0x7b, 0x03, 0x75, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01, 0x88, 0x93, 0x03, 0x00, 0x00, 0x00, 0x00, 0x44, 0xca, 0xa2, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x45, 0xfc, 0x00, 0x00, 0x36, 0x00, 0x01, 0x00, 0x18, 0x08, 0x01, 0x00, 0x09, 0x0c, 0x01, 0x00, 0xfa, 0x0f, 0x01, 0x00}

	var got ReadStatboardHistoryAggregatedRequest

	roundTrip(t, "ReadStatboardHistoryAggregatedRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadStatboardHistoryAggregatedRequest.Queries differs")
	}
}

func TestReadStatboardHistoryAggregatedRequest_Header(t *testing.T) {
	want := ReadStatboardHistoryAggregatedRequest{Queries: types.List[player_stats_types.StatboardHistoryAggregatedQuery]{func() player_stats_types.StatboardHistoryAggregatedQuery {
		v := player_stats_types.NewStatboardHistoryAggregatedQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(4041), types.NewUInt32(5050)}
		v.BoardID = types.NewUInt32(6059)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(8077), types.NewUInt32(9086)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(11104), types.NewUInt32(12113)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(39), Length: types.NewUInt32(49)}
		v.PlayerStats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(16149)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(19176)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-25)}
			return v
		}()}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(23000146)
			v.EndingDatetime = types.NewDateTime(24000149)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(26000155)
			v.EndingDatetime = types.NewDateTime(27000158)
			return v
		}()}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(30275)
			v.SortOrder = types.NewUInt32(31284)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(33302)
			v.SortOrder = types.NewUInt32(34311)
			return v
		}()}
		v.FilterOption = types.NewUInt32(35320)
		return v
	}(), func() player_stats_types.StatboardHistoryAggregatedQuery {
		v := player_stats_types.NewStatboardHistoryAggregatedQuery()
		v.PlayerPIDs = types.List[types.UInt32]{types.NewUInt32(38347), types.NewUInt32(39356)}
		v.BoardID = types.NewUInt32(40365)
		v.ContextIDs = types.List[types.UInt32]{types.NewUInt32(42383), types.NewUInt32(43392)}
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(45410), types.NewUInt32(46419)}
		v.ResultRange = types.ResultRange{Offset: types.NewUInt32(141), Length: types.NewUInt32(151)}
		v.PlayerStats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(50455)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v51")}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(53482)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-59)}
			return v
		}()}
		v.DateRanges = types.List[player_stats_types.DateRange]{func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(57000248)
			v.EndingDatetime = types.NewDateTime(58000251)
			return v
		}(), func() player_stats_types.DateRange {
			v := player_stats_types.NewDateRange()
			v.StartingDatetime = types.NewDateTime(60000257)
			v.EndingDatetime = types.NewDateTime(61000260)
			return v
		}()}
		v.SortCriterias = types.List[player_stats_types.PlayerStatSortCriteria]{func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(64581)
			v.SortOrder = types.NewUInt32(65590)
			return v
		}(), func() player_stats_types.PlayerStatSortCriteria {
			v := player_stats_types.NewPlayerStatSortCriteria()
			v.SortStatID = types.NewUInt32(67608)
			v.SortOrder = types.NewUInt32(68617)
			return v
		}()}
		v.FilterOption = types.NewUInt32(69626)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xab, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x01, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x52, 0xf4, 0x5e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x95, 0x36, 0x6e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x1b, 0xbb, 0x8c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00, 0xf8, 0x89, 0x00, 0x00, 0x00, 0xab, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x80, 0xa9, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x62, 0xb1, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x8d, 0x00, 0x00, 0x00, 0x97, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x17, 0xc5, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xea, 0xd0, 0x00, 0x00, 0x01, 0xc5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x38, 0xc1, 0x65, 0x03, 0x00, 0x00, 0x00, 0x00, 0x7b, 0x03, 0x75, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01, 0x88, 0x93, 0x03, 0x00, 0x00, 0x00, 0x00, 0x44, 0xca, 0xa2, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x45, 0xfc, 0x00, 0x00, 0x36, 0x00, 0x01, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x18, 0x08, 0x01, 0x00, 0x09, 0x0c, 0x01, 0x00, 0xfa, 0x0f, 0x01, 0x00}

	var got ReadStatboardHistoryAggregatedRequest

	roundTrip(t, "ReadStatboardHistoryAggregatedRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadStatboardHistoryAggregatedRequest.Queries differs")
	}
}

func TestReadStatboardHistoryAggregatedResponse_NoHeader(t *testing.T) {
	want := ReadStatboardHistoryAggregatedResponse{Results: types.List[player_stats_types.StatboardResult]{func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(8077)
			v.PlayerName = types.NewString("s9")
			v.SubmittedTime = types.NewDateTime(10000107)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(13122)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(16149)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(19176)
			v.PlayerName = types.NewString("s20")
			v.SubmittedTime = types.NewDateTime(21000140)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(24221)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(27248)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-33)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(31284)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-37)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(34311)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v35")}
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(37338)
		v.ContextID = types.NewUInt32(38347)
		v.ResetFrequency = types.NewUInt32(39356)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(42383)
			v.PlayerName = types.NewString("s43")
			v.SubmittedTime = types.NewDateTime(44000209)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(47428)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-53)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(50455)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v51")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(53482)
			v.PlayerName = types.NewString("s54")
			v.SubmittedTime = types.NewDateTime(55000242)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(58527)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v59")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(61554)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-67)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(65590)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-71)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(68617)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v69")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x01, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x01, 0xdb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x07, 0x86, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x35, 0x00, 0xda, 0x91, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x33, 0x00, 0xd1, 0x63, 0x9f, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x44, 0xb9, 0x00, 0x00, 0x01, 0xcb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x17, 0xc5, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x31, 0x00, 0xea, 0xd0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x35, 0x34, 0x00, 0xb2, 0x3c, 0x47, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9f, 0xe4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x39, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x01, 0xbd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x36, 0x00, 0x01, 0x00, 0x01, 0xb9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x09, 0x0c, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x39, 0x00}

	var got ReadStatboardHistoryAggregatedResponse

	roundTrip(t, "ReadStatboardHistoryAggregatedResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadStatboardHistoryAggregatedResponse.Results differs")
	}
}

func TestReadStatboardHistoryAggregatedResponse_Header(t *testing.T) {
	want := ReadStatboardHistoryAggregatedResponse{Results: types.List[player_stats_types.StatboardResult]{func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(8077)
			v.PlayerName = types.NewString("s9")
			v.SubmittedTime = types.NewDateTime(10000107)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(13122)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-19)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(16149)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v17")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(19176)
			v.PlayerName = types.NewString("s20")
			v.SubmittedTime = types.NewDateTime(21000140)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(24221)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v25")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(27248)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-33)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(31284)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-37)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(34311)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v35")}
			return v
		}()}
		return v
	}(), func() player_stats_types.StatboardResult {
		v := player_stats_types.NewStatboardResult()
		v.BoardID = types.NewUInt32(37338)
		v.ContextID = types.NewUInt32(38347)
		v.ResetFrequency = types.NewUInt32(39356)
		v.PlayerStatSets = types.List[player_stats_types.PlayerStatSet]{func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(42383)
			v.PlayerName = types.NewString("s43")
			v.SubmittedTime = types.NewDateTime(44000209)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(47428)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-53)}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(50455)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v51")}
				return v
			}()}
			return v
		}(), func() player_stats_types.PlayerStatSet {
			v := player_stats_types.NewPlayerStatSet()
			v.PlayerPID = types.NewUInt32(53482)
			v.PlayerName = types.NewString("s54")
			v.SubmittedTime = types.NewDateTime(55000242)
			v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(58527)
				v.Value = types.Variant{TypeID: 4, Type: types.NewString("v59")}
				return v
			}(), func() player_stats_types.PropertyVariant {
				v := player_stats_types.NewPropertyVariant()
				v.ID = types.NewUInt32(61554)
				v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-67)}
				return v
			}()}
			return v
		}()}
		v.DefaultStatValues = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(65590)
			v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-71)}
			return v
		}(), func() player_stats_types.PropertyVariant {
			v := player_stats_types.NewPropertyVariant()
			v.ID = types.NewUInt32(68617)
			v.Value = types.Variant{TypeID: 4, Type: types.NewString("v69")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xaf, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x37, 0x00, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x03, 0x00, 0x73, 0x39, 0x00, 0xeb, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x01, 0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x37, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x30, 0x00, 0xcc, 0x6f, 0x40, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9d, 0x5e, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x32, 0x35, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x70, 0x6a, 0x00, 0x00, 0x01, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x34, 0x7a, 0x00, 0x00, 0x01, 0xdb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x07, 0x86, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x35, 0x00, 0x00, 0xb0, 0x00, 0x00, 0x00, 0xda, 0x91, 0x00, 0x00, 0xcb, 0x95, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x33, 0x00, 0xd1, 0x63, 0x9f, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x44, 0xb9, 0x00, 0x00, 0x01, 0xcb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x17, 0xc5, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x31, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0xea, 0xd0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x35, 0x34, 0x00, 0xb2, 0x3c, 0x47, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x9f, 0xe4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x39, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x01, 0xbd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x36, 0x00, 0x01, 0x00, 0x01, 0xb9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x39, 0x00}

	var got ReadStatboardHistoryAggregatedResponse

	roundTrip(t, "ReadStatboardHistoryAggregatedResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadStatboardHistoryAggregatedResponse.Results differs")
	}
}

func TestGetStatboardNextPurgeDateRequest_NoHeader(t *testing.T) {
	want := GetStatboardNextPurgeDateRequest{BoardID: types.NewUInt32(1014), ResetFrequency: types.NewUInt32(2023)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00}

	var got GetStatboardNextPurgeDateRequest

	roundTrip(t, "GetStatboardNextPurgeDateRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.BoardID.Equals(want.BoardID)) {
		t.Fatalf("GetStatboardNextPurgeDateRequest.BoardID differs")
	}
	if !(got.ResetFrequency.Equals(want.ResetFrequency)) {
		t.Fatalf("GetStatboardNextPurgeDateRequest.ResetFrequency differs")
	}
}

func TestGetStatboardNextPurgeDateRequest_Header(t *testing.T) {
	want := GetStatboardNextPurgeDateRequest{BoardID: types.NewUInt32(1014), ResetFrequency: types.NewUInt32(2023)}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00}

	var got GetStatboardNextPurgeDateRequest

	roundTrip(t, "GetStatboardNextPurgeDateRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.BoardID.Equals(want.BoardID)) {
		t.Fatalf("GetStatboardNextPurgeDateRequest.BoardID differs")
	}
	if !(got.ResetFrequency.Equals(want.ResetFrequency)) {
		t.Fatalf("GetStatboardNextPurgeDateRequest.ResetFrequency differs")
	}
}

func TestGetStatboardNextPurgeDateResponse_NoHeader(t *testing.T) {
	want := GetStatboardNextPurgeDateResponse{PurgeDate: types.NewDateTime(1000080)}

	data := []byte{0x90, 0x42, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GetStatboardNextPurgeDateResponse

	roundTrip(t, "GetStatboardNextPurgeDateResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PurgeDate.Equals(want.PurgeDate)) {
		t.Fatalf("GetStatboardNextPurgeDateResponse.PurgeDate differs")
	}
}

func TestGetStatboardNextPurgeDateResponse_Header(t *testing.T) {
	want := GetStatboardNextPurgeDateResponse{PurgeDate: types.NewDateTime(1000080)}

	data := []byte{0x90, 0x42, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00}

	var got GetStatboardNextPurgeDateResponse

	roundTrip(t, "GetStatboardNextPurgeDateResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PurgeDate.Equals(want.PurgeDate)) {
		t.Fatalf("GetStatboardNextPurgeDateResponse.PurgeDate differs")
	}
}

func TestReadLeaderboardsNearPlayer2Request_NoHeader(t *testing.T) {
	want := ReadLeaderboardsNearPlayer2Request{PlayerPID: types.NewUInt32(1014), Count: types.NewUInt32(2023), Queries: types.List[player_stats_types.LeaderboardQuery2]{func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResetFrequency = types.NewUInt32(7068)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}(), func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(15140)
		v.ContextID = types.NewUInt32(16149)
		v.ResetFrequency = types.NewUInt32(17158)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(19176), types.NewUInt32(20185)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(22203), types.NewUInt32(23212)}
		return v
	}()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xbb, 0x56, 0x00, 0x00, 0xac, 0x5a, 0x00, 0x00}

	var got ReadLeaderboardsNearPlayer2Request

	roundTrip(t, "ReadLeaderboardsNearPlayer2Request", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerPID.Equals(want.PlayerPID)) {
		t.Fatalf("ReadLeaderboardsNearPlayer2Request.PlayerPID differs")
	}
	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("ReadLeaderboardsNearPlayer2Request.Count differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsNearPlayer2Request.Queries differs")
	}
}

func TestReadLeaderboardsNearPlayer2Request_Header(t *testing.T) {
	want := ReadLeaderboardsNearPlayer2Request{PlayerPID: types.NewUInt32(1014), Count: types.NewUInt32(2023), Queries: types.List[player_stats_types.LeaderboardQuery2]{func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResetFrequency = types.NewUInt32(7068)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}(), func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(15140)
		v.ContextID = types.NewUInt32(16149)
		v.ResetFrequency = types.NewUInt32(17158)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(19176), types.NewUInt32(20185)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(22203), types.NewUInt32(23212)}
		return v
	}()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xbb, 0x56, 0x00, 0x00, 0xac, 0x5a, 0x00, 0x00}

	var got ReadLeaderboardsNearPlayer2Request

	roundTrip(t, "ReadLeaderboardsNearPlayer2Request", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.PlayerPID.Equals(want.PlayerPID)) {
		t.Fatalf("ReadLeaderboardsNearPlayer2Request.PlayerPID differs")
	}
	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("ReadLeaderboardsNearPlayer2Request.Count differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsNearPlayer2Request.Queries differs")
	}
}

func TestReadLeaderboardsNearPlayer2Response_NoHeader(t *testing.T) {
	want := ReadLeaderboardsNearPlayer2Response{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsNearPlayer2Response

	roundTrip(t, "ReadLeaderboardsNearPlayer2Response", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsNearPlayer2Response.Results differs")
	}
}

func TestReadLeaderboardsNearPlayer2Response_Header(t *testing.T) {
	want := ReadLeaderboardsNearPlayer2Response{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsNearPlayer2Response

	roundTrip(t, "ReadLeaderboardsNearPlayer2Response", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsNearPlayer2Response.Results differs")
	}
}

func TestReadLeaderboardsByRank2Request_NoHeader(t *testing.T) {
	want := ReadLeaderboardsByRank2Request{StartingRank: types.NewUInt32(1014), Count: types.NewUInt32(2023), Queries: types.List[player_stats_types.LeaderboardQuery2]{func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResetFrequency = types.NewUInt32(7068)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}(), func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(15140)
		v.ContextID = types.NewUInt32(16149)
		v.ResetFrequency = types.NewUInt32(17158)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(19176), types.NewUInt32(20185)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(22203), types.NewUInt32(23212)}
		return v
	}()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xbb, 0x56, 0x00, 0x00, 0xac, 0x5a, 0x00, 0x00}

	var got ReadLeaderboardsByRank2Request

	roundTrip(t, "ReadLeaderboardsByRank2Request", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.StartingRank.Equals(want.StartingRank)) {
		t.Fatalf("ReadLeaderboardsByRank2Request.StartingRank differs")
	}
	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("ReadLeaderboardsByRank2Request.Count differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsByRank2Request.Queries differs")
	}
}

func TestReadLeaderboardsByRank2Request_Header(t *testing.T) {
	want := ReadLeaderboardsByRank2Request{StartingRank: types.NewUInt32(1014), Count: types.NewUInt32(2023), Queries: types.List[player_stats_types.LeaderboardQuery2]{func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(5050)
		v.ContextID = types.NewUInt32(6059)
		v.ResetFrequency = types.NewUInt32(7068)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(9086), types.NewUInt32(10095)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}(), func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(15140)
		v.ContextID = types.NewUInt32(16149)
		v.ResetFrequency = types.NewUInt32(17158)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(19176), types.NewUInt32(20185)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(22203), types.NewUInt32(23212)}
		return v
	}()}}

	data := []byte{0xf6, 0x03, 0x00, 0x00, 0xe7, 0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x15, 0x3f, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xe8, 0x4a, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xbb, 0x56, 0x00, 0x00, 0xac, 0x5a, 0x00, 0x00}

	var got ReadLeaderboardsByRank2Request

	roundTrip(t, "ReadLeaderboardsByRank2Request", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.StartingRank.Equals(want.StartingRank)) {
		t.Fatalf("ReadLeaderboardsByRank2Request.StartingRank differs")
	}
	if !(got.Count.Equals(want.Count)) {
		t.Fatalf("ReadLeaderboardsByRank2Request.Count differs")
	}
	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsByRank2Request.Queries differs")
	}
}

func TestReadLeaderboardsByRank2Response_NoHeader(t *testing.T) {
	want := ReadLeaderboardsByRank2Response{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsByRank2Response

	roundTrip(t, "ReadLeaderboardsByRank2Response", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsByRank2Response.Results differs")
	}
}

func TestReadLeaderboardsByRank2Response_Header(t *testing.T) {
	want := ReadLeaderboardsByRank2Response{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsByRank2Response

	roundTrip(t, "ReadLeaderboardsByRank2Response", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsByRank2Response.Results differs")
	}
}

func TestReadLeaderboardsByPlayers2Request_NoHeader(t *testing.T) {
	want := ReadLeaderboardsByPlayers2Request{Queries: types.List[player_stats_types.LeaderboardQuery2]{func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(7068), types.NewUInt32(8077)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(10095), types.NewUInt32(11104)}
		return v
	}(), func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(13122)
		v.ContextID = types.NewUInt32(14131)
		v.ResetFrequency = types.NewUInt32(15140)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(17158), types.NewUInt32(18167)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(20185), types.NewUInt32(21194)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00}

	var got ReadLeaderboardsByPlayers2Request

	roundTrip(t, "ReadLeaderboardsByPlayers2Request", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsByPlayers2Request.Queries differs")
	}
}

func TestReadLeaderboardsByPlayers2Request_Header(t *testing.T) {
	want := ReadLeaderboardsByPlayers2Request{Queries: types.List[player_stats_types.LeaderboardQuery2]{func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(7068), types.NewUInt32(8077)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(10095), types.NewUInt32(11104)}
		return v
	}(), func() player_stats_types.LeaderboardQuery2 {
		v := player_stats_types.NewLeaderboardQuery2()
		v.BoardID = types.NewUInt32(13122)
		v.ContextID = types.NewUInt32(14131)
		v.ResetFrequency = types.NewUInt32(15140)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(17158), types.NewUInt32(18167)}
		v.EstimatedPIDs = types.List[types.UInt32]{types.NewUInt32(20185), types.NewUInt32(21194)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x8d, 0x1f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00, 0x33, 0x37, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x43, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00}

	var got ReadLeaderboardsByPlayers2Request

	roundTrip(t, "ReadLeaderboardsByPlayers2Request", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadLeaderboardsByPlayers2Request.Queries differs")
	}
}

func TestReadLeaderboardsByPlayers2Response_NoHeader(t *testing.T) {
	want := ReadLeaderboardsByPlayers2Response{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsByPlayers2Response

	roundTrip(t, "ReadLeaderboardsByPlayers2Response", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsByPlayers2Response.Results differs")
	}
}

func TestReadLeaderboardsByPlayers2Response_Header(t *testing.T) {
	want := ReadLeaderboardsByPlayers2Response{Results: types.List[player_stats_types.LeaderboardResult]{func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.ResetFrequency = types.NewUInt32(5050)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(6059)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(10095)
				v.PlayerName = types.NewString("s11")
				v.SubmittedTime = types.NewDateTime(12000113)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(15140)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-21)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(18167)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v19")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(20185)
			v.Rank = types.NewUInt32(21194)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-27)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(25230)
				v.PlayerName = types.NewString("s26")
				v.SubmittedTime = types.NewDateTime(27000158)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(30275)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v31")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(33302)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-39)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(35320)
			v.Rank = types.NewUInt32(36329)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v37")}
			return v
		}()}
		return v
	}(), func() player_stats_types.LeaderboardResult {
		v := player_stats_types.NewLeaderboardResult()
		v.BoardID = types.NewUInt32(39356)
		v.ContextID = types.NewUInt32(40365)
		v.ResetFrequency = types.NewUInt32(41374)
		v.LeaderboardTotalPlayerCount = types.NewUInt32(42383)
		v.PlayerRanks = types.List[player_stats_types.PlayerRank]{func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(46419)
				v.PlayerName = types.NewString("s47")
				v.SubmittedTime = types.NewDateTime(48000221)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(51464)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-57)}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(54491)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v55")}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(56509)
			v.Rank = types.NewUInt32(57518)
			v.Score = types.Variant{TypeID: 1, Type: types.NewInt64(-63)}
			return v
		}(), func() player_stats_types.PlayerRank {
			v := player_stats_types.NewPlayerRank()
			v.PlayerStatSet = func() player_stats_types.PlayerStatSet {
				v := player_stats_types.NewPlayerStatSet()
				v.PlayerPID = types.NewUInt32(61554)
				v.PlayerName = types.NewString("s62")
				v.SubmittedTime = types.NewDateTime(63000266)
				v.Stats = types.List[player_stats_types.PropertyVariant]{func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(66599)
					v.Value = types.Variant{TypeID: 4, Type: types.NewString("v67")}
					return v
				}(), func() player_stats_types.PropertyVariant {
					v := player_stats_types.NewPropertyVariant()
					v.ID = types.NewUInt32(69626)
					v.Value = types.Variant{TypeID: 1, Type: types.NewInt64(-75)}
					return v
				}()}
				return v
			}()
			v.RankStatus = types.NewUInt32(71644)
			v.Rank = types.NewUInt32(72653)
			v.Score = types.Variant{TypeID: 4, Type: types.NewString("v73")}
			return v
		}()}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x04, 0x00, 0x73, 0x31, 0x31, 0x00, 0x71, 0x1b, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x24, 0x3b, 0x00, 0x00, 0x01, 0xeb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xf7, 0x46, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x31, 0x39, 0x00, 0xd9, 0x4e, 0x00, 0x00, 0xca, 0x52, 0x00, 0x00, 0x01, 0xe5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x8e, 0x62, 0x00, 0x00, 0x04, 0x00, 0x73, 0x32, 0x36, 0x00, 0x5e, 0xfd, 0x9b, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x43, 0x76, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x31, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x16, 0x82, 0x00, 0x00, 0x01, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x89, 0x00, 0x00, 0xe9, 0x8d, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x33, 0x37, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0xbc, 0x99, 0x00, 0x00, 0xad, 0x9d, 0x00, 0x00, 0x9e, 0xa1, 0x00, 0x00, 0x8f, 0xa5, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x53, 0xb5, 0x00, 0x00, 0x04, 0x00, 0x73, 0x34, 0x37, 0x00, 0xdd, 0x6c, 0xdc, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x08, 0xc9, 0x00, 0x00, 0x01, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x0b, 0x00, 0x00, 0x00, 0xdb, 0xd4, 0x00, 0x00, 0x04, 0x04, 0x00, 0x76, 0x35, 0x35, 0x00, 0xbd, 0xdc, 0x00, 0x00, 0xae, 0xe0, 0x00, 0x00, 0x01, 0xc1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x72, 0xf0, 0x00, 0x00, 0x04, 0x00, 0x73, 0x36, 0x32, 0x00, 0xca, 0x4e, 0xc1, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x27, 0x04, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x36, 0x37, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x0f, 0x01, 0x00, 0x01, 0xb5, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdc, 0x17, 0x01, 0x00, 0xcd, 0x1b, 0x01, 0x00, 0x04, 0x04, 0x00, 0x76, 0x37, 0x33, 0x00}

	var got ReadLeaderboardsByPlayers2Response

	roundTrip(t, "ReadLeaderboardsByPlayers2Response", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadLeaderboardsByPlayers2Response.Results differs")
	}
}

func TestReadPopulationStatsRequest_NoHeader(t *testing.T) {
	want := ReadPopulationStatsRequest{Queries: types.List[player_stats_types.PopulationStatQuery]{func() player_stats_types.PopulationStatQuery {
		v := player_stats_types.NewPopulationStatQuery()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		return v
	}(), func() player_stats_types.PopulationStatQuery {
		v := player_stats_types.NewPopulationStatQuery()
		v.BoardID = types.NewUInt32(9086)
		v.ContextID = types.NewUInt32(10095)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00}

	var got ReadPopulationStatsRequest

	roundTrip(t, "ReadPopulationStatsRequest", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadPopulationStatsRequest.Queries differs")
	}
}

func TestReadPopulationStatsRequest_Header(t *testing.T) {
	want := ReadPopulationStatsRequest{Queries: types.List[player_stats_types.PopulationStatQuery]{func() player_stats_types.PopulationStatQuery {
		v := player_stats_types.NewPopulationStatQuery()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(6059), types.NewUInt32(7068)}
		return v
	}(), func() player_stats_types.PopulationStatQuery {
		v := player_stats_types.NewPopulationStatQuery()
		v.BoardID = types.NewUInt32(9086)
		v.ContextID = types.NewUInt32(10095)
		v.StatIDs = types.List[types.UInt32]{types.NewUInt32(12113), types.NewUInt32(13122)}
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xab, 0x17, 0x00, 0x00, 0x9c, 0x1b, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x7e, 0x23, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x42, 0x33, 0x00, 0x00}

	var got ReadPopulationStatsRequest

	roundTrip(t, "ReadPopulationStatsRequest", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Queries.Equals(want.Queries)) {
		t.Fatalf("ReadPopulationStatsRequest.Queries differs")
	}
}

func TestReadPopulationStatsResponse_NoHeader(t *testing.T) {
	want := ReadPopulationStatsResponse{Results: types.List[player_stats_types.PopulationStatResult]{func() player_stats_types.PopulationStatResult {
		v := player_stats_types.NewPopulationStatResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.StatID = types.NewUInt32(5050)
		v.Sum = types.NewDouble(6.5)
		v.Average = types.NewDouble(7.5)
		v.StandardDeviation = types.NewDouble(8.5)
		return v
	}(), func() player_stats_types.PopulationStatResult {
		v := player_stats_types.NewPopulationStatResult()
		v.BoardID = types.NewUInt32(10095)
		v.ContextID = types.NewUInt32(11104)
		v.StatID = types.NewUInt32(12113)
		v.Sum = types.NewDouble(13.5)
		v.Average = types.NewDouble(14.5)
		v.StandardDeviation = types.NewDouble(15.5)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1a, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x40, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2b, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2d, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2f, 0x40}

	var got ReadPopulationStatsResponse

	roundTrip(t, "ReadPopulationStatsResponse", false, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadPopulationStatsResponse.Results differs")
	}
}

func TestReadPopulationStatsResponse_Header(t *testing.T) {
	want := ReadPopulationStatsResponse{Results: types.List[player_stats_types.PopulationStatResult]{func() player_stats_types.PopulationStatResult {
		v := player_stats_types.NewPopulationStatResult()
		v.BoardID = types.NewUInt32(3032)
		v.ContextID = types.NewUInt32(4041)
		v.StatID = types.NewUInt32(5050)
		v.Sum = types.NewDouble(6.5)
		v.Average = types.NewDouble(7.5)
		v.StandardDeviation = types.NewDouble(8.5)
		return v
	}(), func() player_stats_types.PopulationStatResult {
		v := player_stats_types.NewPopulationStatResult()
		v.BoardID = types.NewUInt32(10095)
		v.ContextID = types.NewUInt32(11104)
		v.StatID = types.NewUInt32(12113)
		v.Sum = types.NewDouble(13.5)
		v.Average = types.NewDouble(14.5)
		v.StandardDeviation = types.NewDouble(15.5)
		return v
	}()}}

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0xd8, 0x0b, 0x00, 0x00, 0xc9, 0x0f, 0x00, 0x00, 0xba, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1a, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x40, 0x00, 0x24, 0x00, 0x00, 0x00, 0x6f, 0x27, 0x00, 0x00, 0x60, 0x2b, 0x00, 0x00, 0x51, 0x2f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2b, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2d, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2f, 0x40}

	var got ReadPopulationStatsResponse

	roundTrip(t, "ReadPopulationStatsResponse", true, data, func(in *nex.ByteStreamIn) error { return got.ExtractFrom(in) }, func(out *nex.ByteStreamOut) { want.WriteTo(out) })

	if !(got.Results.Equals(want.Results)) {
		t.Fatalf("ReadPopulationStatsResponse.Results differs")
	}
}

var _ Interface = (*Protocol)(nil)

func TestMethodIDs(t *testing.T) {
	if MethodWriteStats != 1 {
		t.Fatalf("MethodWriteStats = %d, want 1", MethodWriteStats)
	}
	if MethodReadStatsByPlayers != 2 {
		t.Fatalf("MethodReadStatsByPlayers = %d, want 2", MethodReadStatsByPlayers)
	}
	if MethodReadLeaderboardsNearPlayer != 3 {
		t.Fatalf("MethodReadLeaderboardsNearPlayer = %d, want 3", MethodReadLeaderboardsNearPlayer)
	}
	if MethodReadLeaderboardsByRank != 4 {
		t.Fatalf("MethodReadLeaderboardsByRank = %d, want 4", MethodReadLeaderboardsByRank)
	}
	if MethodReadLeaderboardsByPlayers != 5 {
		t.Fatalf("MethodReadLeaderboardsByPlayers = %d, want 5", MethodReadLeaderboardsByPlayers)
	}
	if MethodReadStatboardHistory != 6 {
		t.Fatalf("MethodReadStatboardHistory = %d, want 6", MethodReadStatboardHistory)
	}
	if MethodReadLeaderboardHistory != 7 {
		t.Fatalf("MethodReadLeaderboardHistory = %d, want 7", MethodReadLeaderboardHistory)
	}
	if MethodReadStatboardHistoryAggregated != 8 {
		t.Fatalf("MethodReadStatboardHistoryAggregated = %d, want 8", MethodReadStatboardHistoryAggregated)
	}
	if MethodGetStatboardNextPurgeDate != 9 {
		t.Fatalf("MethodGetStatboardNextPurgeDate = %d, want 9", MethodGetStatboardNextPurgeDate)
	}
	if MethodReadLeaderboardsNearPlayer2 != 10 {
		t.Fatalf("MethodReadLeaderboardsNearPlayer2 = %d, want 10", MethodReadLeaderboardsNearPlayer2)
	}
	if MethodReadLeaderboardsByRank2 != 11 {
		t.Fatalf("MethodReadLeaderboardsByRank2 = %d, want 11", MethodReadLeaderboardsByRank2)
	}
	if MethodReadLeaderboardsByPlayers2 != 12 {
		t.Fatalf("MethodReadLeaderboardsByPlayers2 = %d, want 12", MethodReadLeaderboardsByPlayers2)
	}
	if MethodReadPopulationStats != 13 {
		t.Fatalf("MethodReadPopulationStats = %d, want 13", MethodReadPopulationStats)
	}
	if ProtocolID != 55 {
		t.Fatalf("ProtocolID = %d, want 55", ProtocolID)
	}
}
