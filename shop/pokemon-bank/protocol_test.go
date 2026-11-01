package protocol

import (
	"bytes"
	"testing"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	shop_types "github.com/PretendoNetwork/nex-protocols-go/v2/shop/types"
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

func expectBytes(t *testing.T, name string, got, want []byte) {
	t.Helper()

	if !bytes.Equal(got, want) {
		t.Fatalf("%s bytes\n got  %x\n want %x", name, got, want)
	}
}

// wiki encodings: String = u16 length (with NUL) + chars + NUL, qBuffer = u16 length + data,
// List = u32 count, DateTime = u64, everything little-endian
var (
	strAB      = []byte{0x03, 0x00, 'a', 'b', 0x00}
	qbufDEADBE = []byte{0x03, 0x00, 0xDE, 0xAD, 0xBE}
)

func TestMethodIDs(t *testing.T) {
	want := []uint32{
		MethodGetItems, MethodGetChallengeBlob, MethodGetRivToken, MethodGetRivTokenByItemID,
		MethodGetItemRights, MethodVerifyAndRegisterTicket, MethodDebugSetExpireTime,
		MethodPrincipalIDToSupportNumber, MethodSupportNumberToPrincipalID, MethodGetGameServerTime,
	}

	for i, id := range want {
		if id != uint32(i+1) {
			t.Errorf("method %d has ID %d", i+1, id)
		}
	}

	if ProtocolID != 200 || len(patchedMethods) != 10 {
		t.Errorf("ProtocolID %d, %d patched methods", ProtocolID, len(patchedMethods))
	}
}

func TestGetRivTokenRequest(t *testing.T) {
	itemCode, referenceID, err := extractGetRivTokenParameters(newIn(append(append([]byte{}, strAB...), qbufDEADBE...)))
	if err != nil {
		t.Fatal(err)
	}

	if string(itemCode) != "ab" || !bytes.Equal(referenceID, []byte{0xDE, 0xAD, 0xBE}) {
		t.Fatalf("got %q %x", string(itemCode), []byte(referenceID))
	}

	if _, _, err = extractGetRivTokenParameters(newIn(strAB)); err == nil {
		t.Fatal("expected error on truncated request")
	}
}

func TestGetRivTokenByItemIDRequest(t *testing.T) {
	itemID, err := extractGetRivTokenByItemIDParameters(newIn([]byte{0x78, 0x56, 0x34, 0x12}))
	if err != nil || itemID != 0x12345678 {
		t.Fatalf("got %x %v", uint32(itemID), err)
	}
}

func TestGetItemRightsRequest(t *testing.T) {
	ticket, err := extractGetItemRightsParameters(newIn(qbufDEADBE))
	if err != nil || !bytes.Equal(ticket, []byte{0xDE, 0xAD, 0xBE}) {
		t.Fatalf("got %x %v", []byte(ticket), err)
	}
}

func TestVerifyAndRegisterTicketRequest(t *testing.T) {
	data := append(append([]byte{}, qbufDEADBE...), 1, 2, 3, 4, 5, 6, 7, 8)

	ticket, purchased, err := extractVerifyAndRegisterTicketParameters(newIn(data))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(ticket, []byte{0xDE, 0xAD, 0xBE}) || uint64(purchased) != 0x0807060504030201 {
		t.Fatalf("got %x %x", []byte(ticket), uint64(purchased))
	}
}

func TestDebugSetExpireTimeRequest(t *testing.T) {
	expire, err := extractDebugSetExpireTimeParameters(newIn([]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	if err != nil || uint64(expire) != 0x0102030405060708 {
		t.Fatalf("got %x %v", uint64(expire), err)
	}
}

func TestPrincipalIDToSupportNumberRequest(t *testing.T) {
	pid, err := extractPrincipalIDToSupportNumberParameters(newIn([]byte{0x01, 0x00, 0x00, 0x80}))
	if err != nil || pid != 0x80000001 {
		t.Fatalf("got %x %v", uint32(pid), err)
	}
}

func TestSupportNumberToPrincipalIDRequest(t *testing.T) {
	supportNumber, err := extractSupportNumberToPrincipalIDParameters(newIn(strAB))
	if err != nil || string(supportNumber) != "ab" {
		t.Fatalf("got %q %v", string(supportNumber), err)
	}
}

// Responses: the handlers hand the body to the callback owner, so these check the
// wire layout of the response types against the wiki tables.

func TestGetItemsResponse(t *testing.T) {
	item := shop_types.NewShopItem()
	item.ItemID = types.NewUInt32(7)
	item.ReferenceID = types.NewQBuffer([]byte{0xAA})
	item.ServiceName = types.NewString("a")
	item.ItemCode = types.NewString("b")

	items := types.NewList[shop_types.ShopItem]()
	items = append(items, item)

	out := newOut()
	items.WriteTo(out)

	content := []byte{
		7, 0, 0, 0, // itemId
		0x01, 0x00, 0xAA, // referenceId
		0x02, 0x00, 'a', 0x00, // serviceName
		0x02, 0x00, 'b', 0x00, // itemCode
	}
	want := append([]byte{1, 0, 0, 0, 0 /* structure version */, byte(len(content)), 0, 0, 0}, content...)
	expectBytes(t, "GetItems response", out.Bytes(), want)

	var back types.List[shop_types.ShopItem]
	if err := back.ExtractFrom(newIn(want)); err != nil || !back.Equals(items) {
		t.Fatalf("round trip failed: %v", err)
	}
}

func TestGetChallengeBlobResponse(t *testing.T) {
	out := newOut()
	types.NewBuffer([]byte{1, 2, 3}).WriteTo(out)
	expectBytes(t, "GetChallengeBlob response", out.Bytes(), []byte{3, 0, 0, 0, 1, 2, 3})
}

func TestStringResponses(t *testing.T) {
	// GetRivToken, GetRivTokenByItemID and PrincipalIDToSupportNumber all return a single String
	out := newOut()
	types.NewString("ab").WriteTo(out)
	expectBytes(t, "String response", out.Bytes(), strAB)
}

func TestGetItemRightsResponse(t *testing.T) {
	right := shop_types.NewShopItemRights()
	right.ReferenceID = types.NewQBuffer([]byte{0xAA})
	right.ItemType = types.NewInt8(-1)
	right.Attribute = types.NewUInt32(0x01020304)

	rights := types.NewList[shop_types.ShopItemRights]()
	rights = append(rights, right)

	out := newOut()
	rights.WriteTo(out)

	content := []byte{0x01, 0x00, 0xAA, 0xFF, 4, 3, 2, 1}
	want := append([]byte{1, 0, 0, 0, 0, byte(len(content)), 0, 0, 0}, content...)
	expectBytes(t, "GetItemRights response", out.Bytes(), want)

	var back types.List[shop_types.ShopItemRights]
	if err := back.ExtractFrom(newIn(want)); err != nil || !back.Equals(rights) {
		t.Fatalf("round trip failed: %v", err)
	}
}

func TestDateTimeResponses(t *testing.T) {
	// VerifyAndRegisterTicket and GetGameServerTime both return a single DateTime
	out := newOut()
	types.NewDateTime(0x0102030405060708).WriteTo(out)
	expectBytes(t, "DateTime response", out.Bytes(), []byte{8, 7, 6, 5, 4, 3, 2, 1})
}

func TestSupportNumberToPrincipalIDResponse(t *testing.T) {
	out := newOut()
	types.NewUInt32(0x80000001).WriteTo(out)
	expectBytes(t, "SupportNumberToPrincipalID response", out.Bytes(), []byte{1, 0, 0, 0x80})
}
