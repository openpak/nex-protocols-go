package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/pokemon-bank/types"
)

func settings() *nex.ByteStreamSettings {
	s := nex.NewByteStreamSettings()
	s.UseStructureHeader = true

	return s
}

func newIn(data []byte) *nex.ByteStreamIn {
	return nex.NewByteStreamIn(data, nex.NewLibraryVersions(), settings())
}

func newOut() *nex.ByteStreamOut {
	return nex.NewByteStreamOut(nex.NewLibraryVersions(), settings())
}

// TestMethodIDs checks the method table against the Data Store (Pokemon Bank) wiki page (IDs 40-57)
func TestMethodIDs(t *testing.T) {
	want := []uint32{
		MethodPrepareUploadPokemon, MethodUploadPokemon, MethodSearchPokemon, MethodPrepareTradePokemon,
		MethodTradePokemon, MethodDownloadOtherPokemon, MethodDownloadMyPokemon, MethodDeletePokemon,
		MethodGetTransactionParam, MethodPreparePostBankObject, MethodCompletePostBankObject,
		MethodPrepareGetBankObject, MethodPrepareUpdateBankObject, MethodCompleteUpdateBankObject,
		MethodRollbackBankObject, MethodGetUnlockKey, MethodRequestMigration, MethodGetMigrationStatus,
	}

	for i, id := range want {
		if id != uint32(40+i) {
			t.Errorf("method %d has ID %d", 40+i, id)
		}
	}

	if len(patchedMethods) != len(want) {
		t.Errorf("%d patched methods, want %d", len(patchedMethods), len(want))
	}
}

// GetMigrationStatus response: BankMigrationInfo pInfo, Uint32 detailCode
func TestGetMigrationStatusResponse(t *testing.T) {
	info := datastore_pokemon_bank_types.NewBankMigrationInfo()
	info.MigrationStatus = types.NewUInt32(2)
	info.UpdatedTime = types.NewDateTime(0x0102030405060708)

	out := newOut()
	info.WriteTo(out)
	types.NewUInt32(0x0A0B0C0D).WriteTo(out)

	content := []byte{2, 0, 0, 0, 8, 7, 6, 5, 4, 3, 2, 1}
	want := append([]byte{0, byte(len(content)), 0, 0, 0}, content...)
	want = append(want, 0x0D, 0x0C, 0x0B, 0x0A)

	if !bytes.Equal(out.Bytes(), want) {
		t.Fatalf("response bytes\n got  %x\n want %x", out.Bytes(), want)
	}

	in := newIn(want)

	var back datastore_pokemon_bank_types.BankMigrationInfo
	var detailCode types.UInt32

	if err := back.ExtractFrom(in); err != nil {
		t.Fatal(err)
	}

	if err := detailCode.ExtractFrom(in); err != nil {
		t.Fatal(err)
	}

	if !back.Equals(info) || detailCode != 0x0A0B0C0D {
		t.Fatalf("round trip mismatch: %s %x", back, uint32(detailCode))
	}
}
