package protocol

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	"testing"
)

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

func TestMethodIDLookupSceNpIds(t *testing.T) {
	if MethodLookupSceNpIds != 1 {
		t.Fatalf("MethodLookupSceNpIds = %d, want 1", MethodLookupSceNpIds)
	}
}

func TestParseLookupSceNpIdsRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000008b1100008c110000"}, {true, "020000008b1100008c110000"}}

	for _, c := range cases {
		w0 := types.List[types.UInt32]{types.NewUInt32(4491), types.NewUInt32(4492)}

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupSceNpIdsRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("pids differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLookupSceNpIdsRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupPrincipalIDs(t *testing.T) {
	if MethodLookupPrincipalIDs != 2 {
		t.Fatalf("MethodLookupPrincipalIDs = %d, want 2", MethodLookupPrincipalIDs)
	}
}

func TestParseLookupPrincipalIDsRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "02000000050073333938000500733339390090110000"}, {true, "02000000050073333938000500733339390090110000"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s398"), types.NewString("s399")}
		w1 := types.NewUInt32(4496)

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseLookupPrincipalIDsRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("firstPartyIds differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("platformId differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseLookupPrincipalIDsRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupFirstPartyIds(t *testing.T) {
	if MethodLookupFirstPartyIds != 3 {
		t.Fatalf("MethodLookupFirstPartyIds = %d, want 3", MethodLookupFirstPartyIds)
	}
}

func TestParseLookupFirstPartyIdsRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "02000000921100009311000094110000"}, {true, "02000000921100009311000094110000"}}

	for _, c := range cases {
		w0 := types.List[types.UInt32]{types.NewUInt32(4498), types.NewUInt32(4499)}
		w1 := types.NewUInt32(4500)

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseLookupFirstPartyIdsRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("pids differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("platformId differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseLookupFirstPartyIdsRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDUserHasPlayed(t *testing.T) {
	if MethodUserHasPlayed != 4 {
		t.Fatalf("MethodUserHasPlayed = %d, want 4", MethodUserHasPlayed)
	}
}

func TestParseUserHasPlayedRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "02000000050073343036000500733430370098110000"}, {true, "02000000050073343036000500733430370098110000"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s406"), types.NewString("s407")}
		w1 := types.NewUInt32(4504)

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseUserHasPlayedRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("firstPartyIds differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("platformId differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseUserHasPlayedRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDIsUserPlaying(t *testing.T) {
	if MethodIsUserPlaying != 5 {
		t.Fatalf("MethodIsUserPlaying = %d, want 5", MethodIsUserPlaying)
	}
}

func TestParseIsUserPlayingRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0200000005007334313000050073343131009c110000"}, {true, "0200000005007334313000050073343131009c110000"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s410"), types.NewString("s411")}
		w1 := types.NewUInt32(4508)

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseIsUserPlayingRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("firstPartyIds differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("platformId differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseIsUserPlayingRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDUpdateSonyAccountInfo(t *testing.T) {
	if MethodUpdateSonyAccountInfo != 6 {
		t.Fatalf("MethodUpdateSonyAccountInfo = %d, want 6", MethodUpdateSonyAccountInfo)
	}
}

func TestParseUpdateSonyAccountInfoRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "03009d02039e110000"}, {true, "03009d02039e110000"}}

	for _, c := range cases {
		w0 := types.NewQBuffer([]byte{157, 2, 3})
		w1 := types.NewUInt32(4510)

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseUpdateSonyAccountInfoRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ticketData differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("ticketSize differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseUpdateSonyAccountInfoRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupUsernames(t *testing.T) {
	if MethodLookupUsernames != 7 {
		t.Fatalf("MethodLookupUsernames = %d, want 7", MethodLookupUsernames)
	}
}

func TestParseLookupUsernamesRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "02000000a0110000a1110000a2110000"}, {true, "02000000a0110000a1110000a2110000"}}

	for _, c := range cases {
		w0 := types.List[types.UInt32]{types.NewUInt32(4512), types.NewUInt32(4513)}
		w1 := types.NewUInt32(4514)

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseLookupUsernamesRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("pids differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("platformId differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseLookupUsernamesRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}
