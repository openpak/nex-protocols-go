package protocol

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	"testing"
)

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

func TestMethodIDGetActions(t *testing.T) {
	if MethodGetActions != 1 {
		t.Fatalf("MethodGetActions = %d, want 1", MethodGetActions)
	}
}

func TestParseGetActionsRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "a8fdffffa7fdffff05007336303200050073363033000500733630340005007336303500"}, {true, "a8fdffffa7fdffff05007336303200050073363033000500733630340005007336303500"}}

	for _, c := range cases {
		w0 := types.NewInt32(-600)
		w1 := types.NewInt32(-601)
		w2 := types.NewString("s602")
		w3 := types.NewString("s603")
		w4 := types.NewString("s604")
		w5 := types.NewString("s605")

		in := newIn(t, c.data, c.header)
		g0, g1, g2, g3, g4, g5, err := parseGetActionsRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("startRowIndex differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("maximumRows differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("sortExpression differs (header=%v): %v vs %v", c.header, g2, w2)
		}
		if !w3.Equals(g3) {
			t.Fatalf("cultureName differs (header=%v): %v vs %v", c.header, g3, w3)
		}
		if !w4.Equals(g4) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g4, w4)
		}
		if !w5.Equals(g5) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g5, w5)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		w2.WriteTo(out)
		w3.WriteTo(out)
		w4.WriteTo(out)
		w5.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, _, _, _, _, err := parseGetActionsRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetActionsCompleted(t *testing.T) {
	if MethodGetActionsCompleted != 2 {
		t.Fatalf("MethodGetActionsCompleted = %d, want 2", MethodGetActionsCompleted)
	}
}

func TestParseGetActionsCompletedRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "a2fdffffa1fdffff05007336303800050073363039000500733631300005007336313100"}, {true, "a2fdffffa1fdffff05007336303800050073363039000500733631300005007336313100"}}

	for _, c := range cases {
		w0 := types.NewInt32(-606)
		w1 := types.NewInt32(-607)
		w2 := types.NewString("s608")
		w3 := types.NewString("s609")
		w4 := types.NewString("s610")
		w5 := types.NewString("s611")

		in := newIn(t, c.data, c.header)
		g0, g1, g2, g3, g4, g5, err := parseGetActionsCompletedRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("startRowIndex differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("maximumRows differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("sortExpression differs (header=%v): %v vs %v", c.header, g2, w2)
		}
		if !w3.Equals(g3) {
			t.Fatalf("cultureName differs (header=%v): %v vs %v", c.header, g3, w3)
		}
		if !w4.Equals(g4) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g4, w4)
		}
		if !w5.Equals(g5) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g5, w5)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		w2.WriteTo(out)
		w3.WriteTo(out)
		w4.WriteTo(out)
		w5.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, _, _, _, _, err := parseGetActionsCompletedRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetActionsCount(t *testing.T) {
	if MethodGetActionsCount != 3 {
		t.Fatalf("MethodGetActionsCount = %d, want 3", MethodGetActionsCount)
	}
}

func TestParseGetActionsCountRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733631320005007336313300"}, {true, "0500733631320005007336313300"}}

	for _, c := range cases {
		w0 := types.NewString("s612")
		w1 := types.NewString("s613")

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseGetActionsCountRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseGetActionsCountRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetActionsCompletedCount(t *testing.T) {
	if MethodGetActionsCompletedCount != 4 {
		t.Fatalf("MethodGetActionsCompletedCount = %d, want 4", MethodGetActionsCompletedCount)
	}
}

func TestParseGetActionsCompletedCountRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733631340005007336313500"}, {true, "0500733631340005007336313500"}}

	for _, c := range cases {
		w0 := types.NewString("s614")
		w1 := types.NewString("s615")

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseGetActionsCompletedCountRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseGetActionsCompletedCountRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetRewards(t *testing.T) {
	if MethodGetRewards != 5 {
		t.Fatalf("MethodGetRewards = %d, want 5", MethodGetRewards)
	}
}

func TestParseGetRewardsRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "98fdffff97fdffff05007336313800050073363139000500733632300005007336323100"}, {true, "98fdffff97fdffff05007336313800050073363139000500733632300005007336323100"}}

	for _, c := range cases {
		w0 := types.NewInt32(-616)
		w1 := types.NewInt32(-617)
		w2 := types.NewString("s618")
		w3 := types.NewString("s619")
		w4 := types.NewString("s620")
		w5 := types.NewString("s621")

		in := newIn(t, c.data, c.header)
		g0, g1, g2, g3, g4, g5, err := parseGetRewardsRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("startRowIndex differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("maximumRows differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("sortExpression differs (header=%v): %v vs %v", c.header, g2, w2)
		}
		if !w3.Equals(g3) {
			t.Fatalf("cultureName differs (header=%v): %v vs %v", c.header, g3, w3)
		}
		if !w4.Equals(g4) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g4, w4)
		}
		if !w5.Equals(g5) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g5, w5)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		w2.WriteTo(out)
		w3.WriteTo(out)
		w4.WriteTo(out)
		w5.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, _, _, _, _, err := parseGetRewardsRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetRewardsPurchased(t *testing.T) {
	if MethodGetRewardsPurchased != 6 {
		t.Fatalf("MethodGetRewardsPurchased = %d, want 6", MethodGetRewardsPurchased)
	}
}

func TestParseGetRewardsPurchasedRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "92fdffff91fdffff05007336323400050073363235000500733632360005007336323700"}, {true, "92fdffff91fdffff05007336323400050073363235000500733632360005007336323700"}}

	for _, c := range cases {
		w0 := types.NewInt32(-622)
		w1 := types.NewInt32(-623)
		w2 := types.NewString("s624")
		w3 := types.NewString("s625")
		w4 := types.NewString("s626")
		w5 := types.NewString("s627")

		in := newIn(t, c.data, c.header)
		g0, g1, g2, g3, g4, g5, err := parseGetRewardsPurchasedRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("startRowIndex differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("maximumRows differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("sortExpression differs (header=%v): %v vs %v", c.header, g2, w2)
		}
		if !w3.Equals(g3) {
			t.Fatalf("cultureName differs (header=%v): %v vs %v", c.header, g3, w3)
		}
		if !w4.Equals(g4) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g4, w4)
		}
		if !w5.Equals(g5) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g5, w5)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		w2.WriteTo(out)
		w3.WriteTo(out)
		w4.WriteTo(out)
		w5.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, _, _, _, _, err := parseGetRewardsPurchasedRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDUplayWelcome(t *testing.T) {
	if MethodUplayWelcome != 7 {
		t.Fatalf("MethodUplayWelcome = %d, want 7", MethodUplayWelcome)
	}
}

func TestParseUplayWelcomeRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "050073363238000500733632390005007336333000"}, {true, "050073363238000500733632390005007336333000"}}

	for _, c := range cases {
		w0 := types.NewString("s628")
		w1 := types.NewString("s629")
		w2 := types.NewString("s630")

		in := newIn(t, c.data, c.header)
		g0, g1, g2, err := parseUplayWelcomeRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("cultureName differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g2, w2)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		w2.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, _, err := parseUplayWelcomeRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDSetActionCompleted(t *testing.T) {
	if MethodSetActionCompleted != 8 {
		t.Fatalf("MethodSetActionCompleted = %d, want 8", MethodSetActionCompleted)
	}
}

func TestParseSetActionCompletedRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007336333100050073363332000500733633330005007336333400"}, {true, "05007336333100050073363332000500733633330005007336333400"}}

	for _, c := range cases {
		w0 := types.NewString("s631")
		w1 := types.NewString("s632")
		w2 := types.NewString("s633")
		w3 := types.NewString("s634")

		in := newIn(t, c.data, c.header)
		g0, g1, g2, g3, err := parseSetActionCompletedRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("actionCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("cultureName differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g2, w2)
		}
		if !w3.Equals(g3) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g3, w3)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		w2.WriteTo(out)
		w3.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, _, _, err := parseSetActionCompletedRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDSetActionsCompleted(t *testing.T) {
	if MethodSetActionsCompleted != 9 {
		t.Fatalf("MethodSetActionsCompleted = %d, want 9", MethodSetActionsCompleted)
	}
}

func TestParseSetActionsCompletedRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733633360005007336333700050073363338000500733633390005007336343000"}, {true, "020000000500733633360005007336333700050073363338000500733633390005007336343000"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s636"), types.NewString("s637")}
		w1 := types.NewString("s638")
		w2 := types.NewString("s639")
		w3 := types.NewString("s640")

		in := newIn(t, c.data, c.header)
		g0, g1, g2, g3, err := parseSetActionsCompletedRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("actionCodeList differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("cultureName differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g2, w2)
		}
		if !w3.Equals(g3) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g3, w3)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		w2.WriteTo(out)
		w3.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, _, _, err := parseSetActionsCompletedRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetUserToken(t *testing.T) {
	if MethodGetUserToken != 10 {
		t.Fatalf("MethodGetUserToken = %d, want 10", MethodGetUserToken)
	}
}

func TestMethodIDGetVirtualCurrencyUserBalance(t *testing.T) {
	if MethodGetVirtualCurrencyUserBalance != 11 {
		t.Fatalf("MethodGetVirtualCurrencyUserBalance = %d, want 11", MethodGetVirtualCurrencyUserBalance)
	}
}

func TestParseGetVirtualCurrencyUserBalanceRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007336343100"}, {true, "05007336343100"}}

	for _, c := range cases {
		w0 := types.NewString("s641")

		in := newIn(t, c.data, c.header)
		g0, err := parseGetVirtualCurrencyUserBalanceRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseGetVirtualCurrencyUserBalanceRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetSectionsByKey(t *testing.T) {
	if MethodGetSectionsByKey != 12 {
		t.Fatalf("MethodGetSectionsByKey = %d, want 12", MethodGetSectionsByKey)
	}
}

func TestParseGetSectionsByKeyRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007336343200050073363433000500733634340005007336343500"}, {true, "05007336343200050073363433000500733634340005007336343500"}}

	for _, c := range cases {
		w0 := types.NewString("s642")
		w1 := types.NewString("s643")
		w2 := types.NewString("s644")
		w3 := types.NewString("s645")

		in := newIn(t, c.data, c.header)
		g0, g1, g2, g3, err := parseGetSectionsByKeyRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("cultureName differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("sectionKey differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("platformCode differs (header=%v): %v vs %v", c.header, g2, w2)
		}
		if !w3.Equals(g3) {
			t.Fatalf("gameCode differs (header=%v): %v vs %v", c.header, g3, w3)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		w2.WriteTo(out)
		w3.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, _, _, err := parseGetSectionsByKeyRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}
