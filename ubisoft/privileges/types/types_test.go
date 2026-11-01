package types

import (
	"testing"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

func TestPrivilegeRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "b611000005007334333900"}, {true, "000b000000b611000005007334333900"}}

	for _, c := range cases {
		want := func() Privilege {
			v := NewPrivilege()
			v.ID = types.NewUInt32(4534)
			v.Description = types.NewString("s439")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewPrivilege()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(Privilege)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestPrivilegeExRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "b91100000500733434320045feffff"}, {true, "000f000000b91100000500733434320045feffff"}}

	for _, c := range cases {
		want := func() PrivilegeEx {
			v := NewPrivilegeEx()
			v.ID = types.NewUInt32(4537)
			v.Description = types.NewString("s442")
			v.Duration = types.NewInt32(-443)
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewPrivilegeEx()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(PrivilegeEx)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestPrivilegeGroupRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733434350002000000c011000005007334343900c311000005007334353200"}, {true, "002b0000000500733434350002000000000b000000c011000005007334343900000b000000c311000005007334353200"}}

	for _, c := range cases {
		want := func() PrivilegeGroup {
			v := NewPrivilegeGroup()
			v.Description = types.NewString("s445")
			v.Privileges = types.List[Privilege]{func() Privilege {
				v := NewPrivilege()
				v.ID = types.NewUInt32(4544)
				v.Description = types.NewString("s449")
				return v
			}(), func() Privilege {
				v := NewPrivilege()
				v.ID = types.NewUInt32(4547)
				v.Description = types.NewString("s452")
				return v
			}()}
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewPrivilegeGroup()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(PrivilegeGroup)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}
