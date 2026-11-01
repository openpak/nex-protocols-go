package protocol

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	"testing"
)

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

func TestMethodIDAuthenticate(t *testing.T) {
	if MethodAuthenticate != 1 {
		t.Fatalf("MethodAuthenticate = %d, want 1", MethodAuthenticate)
	}
}

func TestParseAuthenticateRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0300733100"}, {true, "0300733100"}}

	for _, c := range cases {
		w0 := types.NewString("s1")

		in := newIn(t, c.data, c.header)
		g0, err := parseAuthenticateRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("strUserName differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseAuthenticateRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLoginWithToken(t *testing.T) {
	if MethodLoginWithToken != 2 {
		t.Fatalf("MethodLoginWithToken = %d, want 2", MethodLoginWithToken)
	}
}

func TestParseLoginWithTokenRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0300733200"}, {true, "0300733200"}}

	for _, c := range cases {
		w0 := types.NewString("s2")

		in := newIn(t, c.data, c.header)
		g0, err := parseLoginWithTokenRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("strToken differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLoginWithTokenRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLoginWithTokenEx(t *testing.T) {
	if MethodLoginWithTokenEx != 3 {
		t.Fatalf("MethodLoginWithTokenEx = %d, want 3", MethodLoginWithTokenEx)
	}
}

func TestParseLoginWithTokenExRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0300733300050044617461000400000000000000"}, {true, "03007333000500446174610009000000050000000000000000"}}

	for _, c := range cases {
		w0 := types.NewString("s3")
		w1 := types.DataHolder{Object: types.NewData()}

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseLoginWithTokenExRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("strToken differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("oAnyData differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseLoginWithTokenExRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLogin(t *testing.T) {
	if MethodLogin != 4 {
		t.Fatalf("MethodLogin = %d, want 4", MethodLogin)
	}
}

func TestParseLoginRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "03007335000300733600"}, {true, "03007335000300733600"}}

	for _, c := range cases {
		w0 := types.NewString("s5")
		w1 := types.NewString("s6")

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseLoginRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("strUsername differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("strPassword differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseLoginRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLoginWithSubAccount(t *testing.T) {
	if MethodLoginWithSubAccount != 5 {
		t.Fatalf("MethodLoginWithSubAccount = %d, want 5", MethodLoginWithSubAccount)
	}
}

func TestParseLoginWithSubAccountRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "050044617461000400000000000000"}, {true, "0500446174610009000000050000000000000000"}}

	for _, c := range cases {
		w0 := types.DataHolder{Object: types.NewData()}

		in := newIn(t, c.data, c.header)
		g0, err := parseLoginWithSubAccountRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("loginData differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLoginWithSubAccountRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDRegister(t *testing.T) {
	if MethodRegister != 6 {
		t.Fatalf("MethodRegister = %d, want 6", MethodRegister)
	}
}

func TestParseRegisterRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "02000000170070727564703a2f616464726573733d312e322e332e3400170070727564703a2f616464726573733d312e322e332e3400"}, {true, "02000000170070727564703a2f616464726573733d312e322e332e3400170070727564703a2f616464726573733d312e322e332e3400"}}

	for _, c := range cases {
		w0 := types.List[types.StationURL]{types.NewStationURL(types.NewString("prudp:/address=1.2.3.4")), types.NewStationURL(types.NewString("prudp:/address=1.2.3.4"))}

		in := newIn(t, c.data, c.header)
		g0, err := parseRegisterRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("vecMyURLs differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseRegisterRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDRegisterEx(t *testing.T) {
	if MethodRegisterEx != 7 {
		t.Fatalf("MethodRegisterEx = %d, want 7", MethodRegisterEx)
	}
}

func TestParseRegisterExRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "02000000170070727564703a2f616464726573733d312e322e332e3400170070727564703a2f616464726573733d312e322e332e3400050044617461000400000000000000"}, {true, "02000000170070727564703a2f616464726573733d312e322e332e3400170070727564703a2f616464726573733d312e322e332e34000500446174610009000000050000000000000000"}}

	for _, c := range cases {
		w0 := types.List[types.StationURL]{types.NewStationURL(types.NewString("prudp:/address=1.2.3.4")), types.NewStationURL(types.NewString("prudp:/address=1.2.3.4"))}
		w1 := types.DataHolder{Object: types.NewData()}

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseRegisterExRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("vecMyURLs differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("hCustomData differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseRegisterExRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLoginWithTokenCafe(t *testing.T) {
	if MethodLoginWithTokenCafe != 8 {
		t.Fatalf("MethodLoginWithTokenCafe = %d, want 8", MethodLoginWithTokenCafe)
	}
}

func TestParseLoginWithTokenCafeRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "040073313500"}, {true, "040073313500"}}

	for _, c := range cases {
		w0 := types.NewString("s15")

		in := newIn(t, c.data, c.header)
		g0, err := parseLoginWithTokenCafeRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("strNintendoToken differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLoginWithTokenCafeRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLoginWithTokenCafeEx(t *testing.T) {
	if MethodLoginWithTokenCafeEx != 9 {
		t.Fatalf("MethodLoginWithTokenCafeEx = %d, want 9", MethodLoginWithTokenCafeEx)
	}
}

func TestParseLoginWithTokenCafeExRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "040073313600050044617461000400000000000000"}, {true, "0400733136000500446174610009000000050000000000000000"}}

	for _, c := range cases {
		w0 := types.NewString("s16")
		w1 := types.DataHolder{Object: types.NewData()}

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseLoginWithTokenCafeExRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("strNintendoToken differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("oAnyData differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseLoginWithTokenCafeExRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}
