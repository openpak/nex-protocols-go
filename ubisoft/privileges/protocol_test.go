package protocol

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	"testing"
)

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

func TestMethodIDGetPrivileges(t *testing.T) {
	if MethodGetPrivileges != 1 {
		t.Fatalf("MethodGetPrivileges = %d, want 1", MethodGetPrivileges)
	}
}

func TestParseGetPrivilegesRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007334353300"}, {true, "05007334353300"}}

	for _, c := range cases {
		w0 := types.NewString("s453")

		in := newIn(t, c.data, c.header)
		g0, err := parseGetPrivilegesRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("localeCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseGetPrivilegesRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDActivateKey(t *testing.T) {
	if MethodActivateKey != 2 {
		t.Fatalf("MethodActivateKey = %d, want 2", MethodActivateKey)
	}
}

func TestParseActivateKeyRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733435340005007334353500"}, {true, "0500733435340005007334353500"}}

	for _, c := range cases {
		w0 := types.NewString("s454")
		w1 := types.NewString("s455")

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseActivateKeyRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("uniqueKey differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("languageCode differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseActivateKeyRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDActivateKeyWithExpectedPrivileges(t *testing.T) {
	if MethodActivateKeyWithExpectedPrivileges != 3 {
		t.Fatalf("MethodActivateKeyWithExpectedPrivileges = %d, want 3", MethodActivateKeyWithExpectedPrivileges)
	}
}

func TestParseActivateKeyWithExpectedPrivilegesRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "050073343536000500733435370002000000cb110000cc110000"}, {true, "050073343536000500733435370002000000cb110000cc110000"}}

	for _, c := range cases {
		w0 := types.NewString("s456")
		w1 := types.NewString("s457")
		w2 := types.List[types.UInt32]{types.NewUInt32(4555), types.NewUInt32(4556)}

		in := newIn(t, c.data, c.header)
		g0, g1, g2, err := parseActivateKeyWithExpectedPrivilegesRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("uniqueKey differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("languageCode differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("expectedPrivilegeIDs differs (header=%v): %v vs %v", c.header, g2, w2)
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
		if _, _, _, err := parseActivateKeyWithExpectedPrivilegesRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetPrivilegeRemainDuration(t *testing.T) {
	if MethodGetPrivilegeRemainDuration != 4 {
		t.Fatalf("MethodGetPrivilegeRemainDuration = %d, want 4", MethodGetPrivilegeRemainDuration)
	}
}

func TestParseGetPrivilegeRemainDurationRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "cd110000"}, {true, "cd110000"}}

	for _, c := range cases {
		w0 := types.NewUInt32(4557)

		in := newIn(t, c.data, c.header)
		g0, err := parseGetPrivilegeRemainDurationRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("privilegeID differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseGetPrivilegeRemainDurationRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetExpiredPrivileges(t *testing.T) {
	if MethodGetExpiredPrivileges != 5 {
		t.Fatalf("MethodGetExpiredPrivileges = %d, want 5", MethodGetExpiredPrivileges)
	}
}

func TestMethodIDGetPrivilegesEx(t *testing.T) {
	if MethodGetPrivilegesEx != 6 {
		t.Fatalf("MethodGetPrivilegesEx = %d, want 6", MethodGetPrivilegesEx)
	}
}

func TestParseGetPrivilegesExRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007334363200"}, {true, "05007334363200"}}

	for _, c := range cases {
		w0 := types.NewString("s462")

		in := newIn(t, c.data, c.header)
		g0, err := parseGetPrivilegesExRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("localeCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseGetPrivilegesExRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}
