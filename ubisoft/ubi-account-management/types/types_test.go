package types

import (
	"testing"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

func TestUbiAccountRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "040073353300040073353400040073353500391000000100010400733631003e776655443322113f100000040073363400000104007336370004007336380004007336390002000000481000000400733733000400733734004c100000040073373700040073373800"}, {true, "00780000000400733533000400733534000400733535000007000000391000000100010400733631003e776655443322113f10000004007336340000010400733637000400733638000400733639000200000000100000004810000004007337330004007337340000100000004c100000040073373700040073373800"}}

	for _, c := range cases {
		want := func() UbiAccount {
			v := NewUbiAccount()
			v.UbiAccountID = types.NewString("s53")
			v.Username = types.NewString("s54")
			v.Password = types.NewString("s55")
			v.Status = func() UbiAccountStatus {
				v := NewUbiAccountStatus()
				v.BasicStatus = types.NewUInt32(4153)
				v.MissingRequiredInformations = types.NewBool(true)
				v.RecoveringPassword = types.NewBool(false)
				v.PendingDeactivation = types.NewBool(true)
				return v
			}()
			v.Email = types.NewString("s61")
			v.DateOfBirth = types.NewDateTime(1234605616436508478)
			v.Gender = types.NewUInt32(4159)
			v.CountryCode = types.NewString("s64")
			v.OptIn = types.NewBool(false)
			v.ThirdPartyOptIn = types.NewBool(true)
			v.FirstName = types.NewString("s67")
			v.LastName = types.NewString("s68")
			v.PreferredLanguage = types.NewString("s69")
			v.ExternalAccounts = types.List[ExternalAccount]{func() ExternalAccount {
				v := NewExternalAccount()
				v.AccountType = types.NewUInt32(4168)
				v.ID = types.NewString("s73")
				v.Username = types.NewString("s74")
				return v
			}(), func() ExternalAccount {
				v := NewExternalAccount()
				v.AccountType = types.NewUInt32(4172)
				v.ID = types.NewString("s77")
				v.Username = types.NewString("s78")
				return v
			}()}
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUbiAccount()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UbiAccount)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestTOSRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "040073383000040073383100040073383200"}, {true, "0012000000040073383000040073383100040073383200"}}

	for _, c := range cases {
		want := func() TOS {
			v := NewTOS()
			v.LocaleCode = types.NewString("s80")
			v.Content = types.NewString("s81")
			v.StoringInfoQuestion = types.NewString("s82")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewTOS()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(TOS)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestTOSExRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "04007338340002000000040073383600040073383700040073383800"}, {true, "001c00000004007338340002000000040073383600040073383700040073383800"}}

	for _, c := range cases {
		want := func() TOSEx {
			v := NewTOSEx()
			v.LocaleCode = types.NewString("s84")
			v.Contents = types.List[types.String]{types.NewString("s86"), types.NewString("s87")}
			v.StoringInfoQuestion = types.NewString("s88")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewTOSEx()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(TOSEx)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestValidationFailureReasonRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "5a100000040073393100"}, {true, "000a0000005a100000040073393100"}}

	for _, c := range cases {
		want := func() ValidationFailureReason {
			v := NewValidationFailureReason()
			v.ValidationID = types.NewUInt32(4186)
			v.Description = types.NewString("s91")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewValidationFailureReason()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(ValidationFailureReason)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestUsernameValidationRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0200000004007339340004007339350002000000621000000400733939006510000005007331303200"}, {true, "00330000000200000004007339340004007339350002000000000a00000062100000040073393900000b0000006510000005007331303200"}}

	for _, c := range cases {
		want := func() UsernameValidation {
			v := NewUsernameValidation()
			v.Suggestions = types.List[types.String]{types.NewString("s94"), types.NewString("s95")}
			v.Reasons = types.List[ValidationFailureReason]{func() ValidationFailureReason {
				v := NewValidationFailureReason()
				v.ValidationID = types.NewUInt32(4194)
				v.Description = types.NewString("s99")
				return v
			}(), func() ValidationFailureReason {
				v := NewValidationFailureReason()
				v.ValidationID = types.NewUInt32(4197)
				v.Description = types.NewString("s102")
				return v
			}()}
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUsernameValidation()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UsernameValidation)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestCountryRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733130340005007331303500"}, {true, "000e0000000500733130340005007331303500"}}

	for _, c := range cases {
		want := func() Country {
			v := NewCountry()
			v.Code = types.NewString("s104")
			v.Name = types.NewString("s105")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewCountry()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(Country)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestExternalAccountRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "6b1000000500733130380005007331303900"}, {true, "00120000006b1000000500733130380005007331303900"}}

	for _, c := range cases {
		want := func() ExternalAccount {
			v := NewExternalAccount()
			v.AccountType = types.NewUInt32(4203)
			v.ID = types.NewString("s108")
			v.Username = types.NewString("s109")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewExternalAccount()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(ExternalAccount)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestUbiAccountStatusRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "6f100000010001"}, {true, "00070000006f100000010001"}}

	for _, c := range cases {
		want := func() UbiAccountStatus {
			v := NewUbiAccountStatus()
			v.BasicStatus = types.NewUInt32(4207)
			v.MissingRequiredInformations = types.NewBool(true)
			v.RecoveringPassword = types.NewBool(false)
			v.PendingDeactivation = types.NewBool(true)
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUbiAccountStatus()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UbiAccountStatus)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}
