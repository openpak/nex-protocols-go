package protocol

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	ubi_account_management_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/ubi-account-management/types"
	"testing"
)

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

func TestMethodIDCreateAccount(t *testing.T) {
	if MethodCreateAccount != 1 {
		t.Fatalf("MethodCreateAccount = %d, want 1", MethodCreateAccount)
	}
}

func TestParseCreateAccountRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007331313600050073313137000500733131380078100000000100050073313234007d776655443322117e100000050073313237000100050073313330000500733133310005007331333200020000008710000005007331333600050073313337008b1000000500733134300005007331343100"}, {true, "0084000000050073313136000500733131370005007331313800000700000078100000000100050073313234007d776655443322117e10000005007331323700010005007331333000050073313331000500733133320002000000001200000087100000050073313336000500733133370000120000008b1000000500733134300005007331343100"}}

	for _, c := range cases {
		w0 := func() ubi_account_management_types.UbiAccount {
			v := ubi_account_management_types.NewUbiAccount()
			v.UbiAccountID = types.NewString("s116")
			v.Username = types.NewString("s117")
			v.Password = types.NewString("s118")
			v.Status = func() ubi_account_management_types.UbiAccountStatus {
				v := ubi_account_management_types.NewUbiAccountStatus()
				v.BasicStatus = types.NewUInt32(4216)
				v.MissingRequiredInformations = types.NewBool(false)
				v.RecoveringPassword = types.NewBool(true)
				v.PendingDeactivation = types.NewBool(false)
				return v
			}()
			v.Email = types.NewString("s124")
			v.DateOfBirth = types.NewDateTime(1234605616436508541)
			v.Gender = types.NewUInt32(4222)
			v.CountryCode = types.NewString("s127")
			v.OptIn = types.NewBool(true)
			v.ThirdPartyOptIn = types.NewBool(false)
			v.FirstName = types.NewString("s130")
			v.LastName = types.NewString("s131")
			v.PreferredLanguage = types.NewString("s132")
			v.ExternalAccounts = types.List[ubi_account_management_types.ExternalAccount]{func() ubi_account_management_types.ExternalAccount {
				v := ubi_account_management_types.NewExternalAccount()
				v.AccountType = types.NewUInt32(4231)
				v.ID = types.NewString("s136")
				v.Username = types.NewString("s137")
				return v
			}(), func() ubi_account_management_types.ExternalAccount {
				v := ubi_account_management_types.NewExternalAccount()
				v.AccountType = types.NewUInt32(4235)
				v.ID = types.NewString("s140")
				v.Username = types.NewString("s141")
				return v
			}()}
			return v
		}()

		in := newIn(t, c.data, c.header)
		g0, err := parseCreateAccountRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ubiAccount differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseCreateAccountRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDUpdateAccount(t *testing.T) {
	if MethodUpdateAccount != 2 {
		t.Fatalf("MethodUpdateAccount = %d, want 2", MethodUpdateAccount)
	}
}

func TestParseUpdateAccountRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "050073313433000500733134340005007331343500931000000100010500733135310098776655443322119910000005007331353400000105007331353700050073313538000500733135390002000000a21000000500733136330005007331363400a61000000500733136370005007331363800"}, {true, "008400000005007331343300050073313434000500733134350000070000009310000001000105007331353100987766554433221199100000050073313534000001050073313537000500733135380005007331353900020000000012000000a210000005007331363300050073313634000012000000a61000000500733136370005007331363800"}}

	for _, c := range cases {
		w0 := func() ubi_account_management_types.UbiAccount {
			v := ubi_account_management_types.NewUbiAccount()
			v.UbiAccountID = types.NewString("s143")
			v.Username = types.NewString("s144")
			v.Password = types.NewString("s145")
			v.Status = func() ubi_account_management_types.UbiAccountStatus {
				v := ubi_account_management_types.NewUbiAccountStatus()
				v.BasicStatus = types.NewUInt32(4243)
				v.MissingRequiredInformations = types.NewBool(true)
				v.RecoveringPassword = types.NewBool(false)
				v.PendingDeactivation = types.NewBool(true)
				return v
			}()
			v.Email = types.NewString("s151")
			v.DateOfBirth = types.NewDateTime(1234605616436508568)
			v.Gender = types.NewUInt32(4249)
			v.CountryCode = types.NewString("s154")
			v.OptIn = types.NewBool(false)
			v.ThirdPartyOptIn = types.NewBool(true)
			v.FirstName = types.NewString("s157")
			v.LastName = types.NewString("s158")
			v.PreferredLanguage = types.NewString("s159")
			v.ExternalAccounts = types.List[ubi_account_management_types.ExternalAccount]{func() ubi_account_management_types.ExternalAccount {
				v := ubi_account_management_types.NewExternalAccount()
				v.AccountType = types.NewUInt32(4258)
				v.ID = types.NewString("s163")
				v.Username = types.NewString("s164")
				return v
			}(), func() ubi_account_management_types.ExternalAccount {
				v := ubi_account_management_types.NewExternalAccount()
				v.AccountType = types.NewUInt32(4262)
				v.ID = types.NewString("s167")
				v.Username = types.NewString("s168")
				return v
			}()}
			return v
		}()

		in := newIn(t, c.data, c.header)
		g0, err := parseUpdateAccountRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ubiAccount differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseUpdateAccountRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetAccount(t *testing.T) {
	if MethodGetAccount != 3 {
		t.Fatalf("MethodGetAccount = %d, want 3", MethodGetAccount)
	}
}

func TestMethodIDLinkAccount(t *testing.T) {
	if MethodLinkAccount != 4 {
		t.Fatalf("MethodLinkAccount = %d, want 4", MethodLinkAccount)
	}
}

func TestParseLinkAccountRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733136390005007331373000"}, {true, "0500733136390005007331373000"}}

	for _, c := range cases {
		w0 := types.NewString("s169")
		w1 := types.NewString("s170")

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseLinkAccountRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ubiAccountUsername differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("ubiAccountPassword differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseLinkAccountRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetTOS(t *testing.T) {
	if MethodGetTOS != 5 {
		t.Fatalf("MethodGetTOS = %d, want 5", MethodGetTOS)
	}
}

func TestParseGetTOSRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "050073313731000500733137320000"}, {true, "050073313731000500733137320000"}}

	for _, c := range cases {
		w0 := types.NewString("s171")
		w1 := types.NewString("s172")
		w2 := types.NewBool(false)

		in := newIn(t, c.data, c.header)
		g0, g1, g2, err := parseGetTOSRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("countryCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("languageCode differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("htmlVersion differs (header=%v): %v vs %v", c.header, g2, w2)
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
		if _, _, _, err := parseGetTOSRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDValidateUsername(t *testing.T) {
	if MethodValidateUsername != 6 {
		t.Fatalf("MethodValidateUsername = %d, want 6", MethodValidateUsername)
	}
}

func TestParseValidateUsernameRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007331373400"}, {true, "05007331373400"}}

	for _, c := range cases {
		w0 := types.NewString("s174")

		in := newIn(t, c.data, c.header)
		g0, err := parseValidateUsernameRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("username differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseValidateUsernameRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDValidatePassword(t *testing.T) {
	if MethodValidatePassword != 7 {
		t.Fatalf("MethodValidatePassword = %d, want 7", MethodValidatePassword)
	}
}

func TestParseValidatePasswordRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733137350005007331373600"}, {true, "0500733137350005007331373600"}}

	for _, c := range cases {
		w0 := types.NewString("s175")
		w1 := types.NewString("s176")

		in := newIn(t, c.data, c.header)
		g0, g1, err := parseValidatePasswordRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("password differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("username differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		w1.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, _, err := parseValidatePasswordRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDValidateEmail(t *testing.T) {
	if MethodValidateEmail != 8 {
		t.Fatalf("MethodValidateEmail = %d, want 8", MethodValidateEmail)
	}
}

func TestParseValidateEmailRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007331373700"}, {true, "05007331373700"}}

	for _, c := range cases {
		w0 := types.NewString("s177")

		in := newIn(t, c.data, c.header)
		g0, err := parseValidateEmailRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("email differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseValidateEmailRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetCountryList(t *testing.T) {
	if MethodGetCountryList != 9 {
		t.Fatalf("MethodGetCountryList = %d, want 9", MethodGetCountryList)
	}
}

func TestParseGetCountryListRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007331373800"}, {true, "05007331373800"}}

	for _, c := range cases {
		w0 := types.NewString("s178")

		in := newIn(t, c.data, c.header)
		g0, err := parseGetCountryListRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("languageCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseGetCountryListRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDForgetPassword(t *testing.T) {
	if MethodForgetPassword != 10 {
		t.Fatalf("MethodForgetPassword = %d, want 10", MethodForgetPassword)
	}
}

func TestParseForgetPasswordRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007331373900"}, {true, "05007331373900"}}

	for _, c := range cases {
		w0 := types.NewString("s179")

		in := newIn(t, c.data, c.header)
		g0, err := parseForgetPasswordRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("usernameOrEmail differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseForgetPasswordRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupPrincipalIds(t *testing.T) {
	if MethodLookupPrincipalIds != 11 {
		t.Fatalf("MethodLookupPrincipalIds = %d, want 11", MethodLookupPrincipalIds)
	}
}

func TestParseLookupPrincipalIdsRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733138310005007331383200"}, {true, "020000000500733138310005007331383200"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s181"), types.NewString("s182")}

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupPrincipalIdsRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ubiAccountIds differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLookupPrincipalIdsRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupUbiAccountIDsByPids(t *testing.T) {
	if MethodLookupUbiAccountIDsByPids != 12 {
		t.Fatalf("MethodLookupUbiAccountIDsByPids = %d, want 12", MethodLookupUbiAccountIDsByPids)
	}
}

func TestParseLookupUbiAccountIDsByPidsRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "02000000b8100000b9100000"}, {true, "02000000b8100000b9100000"}}

	for _, c := range cases {
		w0 := types.List[types.UInt32]{types.NewUInt32(4280), types.NewUInt32(4281)}

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupUbiAccountIDsByPidsRequest(in)
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
		if _, err := parseLookupUbiAccountIDsByPidsRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupUbiAccountIDsByUsernames(t *testing.T) {
	if MethodLookupUbiAccountIDsByUsernames != 13 {
		t.Fatalf("MethodLookupUbiAccountIDsByUsernames = %d, want 13", MethodLookupUbiAccountIDsByUsernames)
	}
}

func TestParseLookupUbiAccountIDsByUsernamesRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733138370005007331383800"}, {true, "020000000500733138370005007331383800"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s187"), types.NewString("s188")}

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupUbiAccountIDsByUsernamesRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("usernames differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLookupUbiAccountIDsByUsernamesRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupUsernamesByUbiAccountIDs(t *testing.T) {
	if MethodLookupUsernamesByUbiAccountIDs != 14 {
		t.Fatalf("MethodLookupUsernamesByUbiAccountIDs = %d, want 14", MethodLookupUsernamesByUbiAccountIDs)
	}
}

func TestParseLookupUsernamesByUbiAccountIDsRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733139300005007331393100"}, {true, "020000000500733139300005007331393100"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s190"), types.NewString("s191")}

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupUsernamesByUbiAccountIDsRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ubiAccountIds differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLookupUsernamesByUbiAccountIDsRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupUbiAccountIDsByUsernameSubString(t *testing.T) {
	if MethodLookupUbiAccountIDsByUsernameSubString != 15 {
		t.Fatalf("MethodLookupUbiAccountIDsByUsernameSubString = %d, want 15", MethodLookupUbiAccountIDsByUsernameSubString)
	}
}

func TestParseLookupUbiAccountIDsByUsernameSubStringRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007331393200"}, {true, "05007331393200"}}

	for _, c := range cases {
		w0 := types.NewString("s192")

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupUbiAccountIDsByUsernameSubStringRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("usernameSubString differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLookupUbiAccountIDsByUsernameSubStringRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDUserHasPlayed(t *testing.T) {
	if MethodUserHasPlayed != 16 {
		t.Fatalf("MethodUserHasPlayed = %d, want 16", MethodUserHasPlayed)
	}
}

func TestParseUserHasPlayedRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733139340005007331393500"}, {true, "020000000500733139340005007331393500"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s194"), types.NewString("s195")}

		in := newIn(t, c.data, c.header)
		g0, err := parseUserHasPlayedRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ubiAccountIDs differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseUserHasPlayedRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDIsUserPlaying(t *testing.T) {
	if MethodIsUserPlaying != 17 {
		t.Fatalf("MethodIsUserPlaying = %d, want 17", MethodIsUserPlaying)
	}
}

func TestParseIsUserPlayingRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733139370005007331393800"}, {true, "020000000500733139370005007331393800"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s197"), types.NewString("s198")}

		in := newIn(t, c.data, c.header)
		g0, err := parseIsUserPlayingRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ubiAccountIDs differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseIsUserPlayingRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupUbiAccountIDsByUsernamesGlobal(t *testing.T) {
	if MethodLookupUbiAccountIDsByUsernamesGlobal != 18 {
		t.Fatalf("MethodLookupUbiAccountIDsByUsernamesGlobal = %d, want 18", MethodLookupUbiAccountIDsByUsernamesGlobal)
	}
}

func TestParseLookupUbiAccountIDsByUsernamesGlobalRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733230300005007332303100"}, {true, "020000000500733230300005007332303100"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s200"), types.NewString("s201")}

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupUbiAccountIDsByUsernamesGlobalRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("usernames differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLookupUbiAccountIDsByUsernamesGlobalRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupUbiAccountIDsByEmailsGlobal(t *testing.T) {
	if MethodLookupUbiAccountIDsByEmailsGlobal != 19 {
		t.Fatalf("MethodLookupUbiAccountIDsByEmailsGlobal = %d, want 19", MethodLookupUbiAccountIDsByEmailsGlobal)
	}
}

func TestParseLookupUbiAccountIDsByEmailsGlobalRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733230330005007332303400"}, {true, "020000000500733230330005007332303400"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s203"), types.NewString("s204")}

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupUbiAccountIDsByEmailsGlobalRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("emails differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLookupUbiAccountIDsByEmailsGlobalRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDLookupUsernamesByUbiAccountIDsGlobal(t *testing.T) {
	if MethodLookupUsernamesByUbiAccountIDsGlobal != 20 {
		t.Fatalf("MethodLookupUsernamesByUbiAccountIDsGlobal = %d, want 20", MethodLookupUsernamesByUbiAccountIDsGlobal)
	}
}

func TestParseLookupUsernamesByUbiAccountIDsGlobalRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "020000000500733230360005007332303700"}, {true, "020000000500733230360005007332303700"}}

	for _, c := range cases {
		w0 := types.List[types.String]{types.NewString("s206"), types.NewString("s207")}

		in := newIn(t, c.data, c.header)
		g0, err := parseLookupUsernamesByUbiAccountIDsGlobalRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("ubiAccountIds differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		out := newOut(c.header)
		w0.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		truncated := newIn(t, c.data[:len(c.data)-2], c.header)
		if _, err := parseLookupUsernamesByUbiAccountIDsGlobalRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDGetTOSEx(t *testing.T) {
	if MethodGetTOSEx != 21 {
		t.Fatalf("MethodGetTOSEx = %d, want 21", MethodGetTOSEx)
	}
}

func TestParseGetTOSExRequest(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "050073323038000500733230390001"}, {true, "050073323038000500733230390001"}}

	for _, c := range cases {
		w0 := types.NewString("s208")
		w1 := types.NewString("s209")
		w2 := types.NewBool(true)

		in := newIn(t, c.data, c.header)
		g0, g1, g2, err := parseGetTOSExRequest(in)
		if err != nil {
			t.Fatal(err)
		}
		if !w0.Equals(g0) {
			t.Fatalf("countryCode differs (header=%v): %v vs %v", c.header, g0, w0)
		}
		if !w1.Equals(g1) {
			t.Fatalf("languageCode differs (header=%v): %v vs %v", c.header, g1, w1)
		}
		if !w2.Equals(g2) {
			t.Fatalf("htmlVersion differs (header=%v): %v vs %v", c.header, g2, w2)
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
		if _, _, _, err := parseGetTOSExRequest(truncated); err == nil {
			t.Fatal("expected an error for truncated parameters")
		}
	}
}

func TestMethodIDHasAcceptedLatestTOS(t *testing.T) {
	if MethodHasAcceptedLatestTOS != 22 {
		t.Fatalf("MethodHasAcceptedLatestTOS = %d, want 22", MethodHasAcceptedLatestTOS)
	}
}

func TestMethodIDAcceptLatestTOS(t *testing.T) {
	if MethodAcceptLatestTOS != 23 {
		t.Fatalf("MethodAcceptLatestTOS = %d, want 23", MethodAcceptLatestTOS)
	}
}
