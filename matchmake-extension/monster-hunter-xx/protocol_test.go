package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	matchmake_extension_monster_hunter_x_x_types "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/monster-hunter-xx/types"
)

func settings() *nex.ByteStreamSettings {
	s := nex.NewByteStreamSettings()
	s.UseStructureHeader = true
	s.PIDSize = 8 // the wiki documents the friend PIDs as Uint64

	return s
}

func newIn(data []byte) *nex.ByteStreamIn {
	return nex.NewByteStreamIn(data, nex.NewLibraryVersions(), settings())
}

func newOut() *nex.ByteStreamOut {
	return nex.NewByteStreamOut(nex.NewLibraryVersions(), settings())
}

// TestMethodIDs checks the method table against the Matchmake Extension (MHXX) wiki page (IDs 54-59)
func TestMethodIDs(t *testing.T) {
	want := []uint32{
		MethodUpdateFriendUserProfile, MethodGetFriendUserProfiles, MethodGetFriends,
		MethodAddFriends, MethodRemoveFriend, MethodFindCommunityByOwner,
	}

	for i, id := range want {
		if id != uint32(54+i) {
			t.Errorf("method %d has ID %d", 54+i, id)
		}
	}

	if len(patchedMethods) != len(want) {
		t.Errorf("%d patched methods, want %d", len(patchedMethods), len(want))
	}
}

// GetFriends response: List<FriendUserInfo> infos
func TestGetFriendsResponse(t *testing.T) {
	info := matchmake_extension_monster_hunter_x_x_types.NewFriendUserInfo()
	info.PID = types.NewPID(0x0102030405060708)
	info.Name = types.NewString("ab")
	info.Presence = types.NewUInt32(5)

	infos := types.NewList[matchmake_extension_monster_hunter_x_x_types.FriendUserInfo]()
	infos = append(infos, info)

	out := newOut()
	infos.WriteTo(out)

	content := []byte{
		8, 7, 6, 5, 4, 3, 2, 1, // pid
		0x03, 0x00, 'a', 'b', 0x00, // name
		5, 0, 0, 0, // presence
	}
	want := append([]byte{1, 0, 0, 0, 0, byte(len(content)), 0, 0, 0}, content...)

	if !bytes.Equal(out.Bytes(), want) {
		t.Fatalf("response bytes\n got  %x\n want %x", out.Bytes(), want)
	}

	var back types.List[matchmake_extension_monster_hunter_x_x_types.FriendUserInfo]
	if err := back.ExtractFrom(newIn(want)); err != nil || !back.Equals(infos) {
		t.Fatalf("round trip failed: %v", err)
	}
}
