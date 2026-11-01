package types

import (
	"testing"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

func init() {
	types.RegisterObjectHolderType(types.NewData())
}

func TestUplayActionRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007335303600050073353037000500733530380003feffff0500733531300002000000050073353133000105007335313500050073353137000105007335313900"}, {true, "004c00000005007335303600050073353037000500733530380003feffff0500733531300002000000000f000000050073353133000105007335313500000f000000050073353137000105007335313900"}}

	for _, c := range cases {
		want := func() UplayAction {
			v := NewUplayAction()
			v.Code = types.NewString("s506")
			v.Name = types.NewString("s507")
			v.Description = types.NewString("s508")
			v.Value = types.NewInt32(-509)
			v.GameCode = types.NewString("s510")
			v.Platforms = types.List[UplayActionPlatform]{func() UplayActionPlatform {
				v := NewUplayActionPlatform()
				v.PlatformCode = types.NewString("s513")
				v.Completed = types.NewBool(true)
				v.SpecificKey = types.NewString("s515")
				return v
			}(), func() UplayActionPlatform {
				v := NewUplayActionPlatform()
				v.PlatformCode = types.NewString("s517")
				v.Completed = types.NewBool(true)
				v.SpecificKey = types.NewString("s519")
				return v
			}()}
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUplayAction()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UplayAction)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestUplayRewardRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "050073353231000500733532320005007335323300f4fdffff05007335323500050073353236000200000005007335323900010500733533320000"}, {true, "0045000000050073353231000500733532320005007335323300f4fdffff0500733532350005007335323600020000000008000000050073353239000100080000000500733533320000"}}

	for _, c := range cases {
		want := func() UplayReward {
			v := NewUplayReward()
			v.Code = types.NewString("s521")
			v.Name = types.NewString("s522")
			v.Description = types.NewString("s523")
			v.Value = types.NewInt32(-524)
			v.RewardTypeName = types.NewString("s525")
			v.GameCode = types.NewString("s526")
			v.Platforms = types.List[UplayRewardPlatform]{func() UplayRewardPlatform {
				v := NewUplayRewardPlatform()
				v.PlatformCode = types.NewString("s529")
				v.Purchased = types.NewBool(true)
				return v
			}(), func() UplayRewardPlatform {
				v := NewUplayRewardPlatform()
				v.PlatformCode = types.NewString("s532")
				v.Purchased = types.NewBool(false)
				return v
			}()}
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUplayReward()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UplayReward)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestUplaySectionRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733533350005007335333600050073353337000500733533380002000000050073353431000500733534320083020500733534340005007335343600050073353437000500733534380005007335343900dafdffff050073353531000500733535320005007335353300050073353535000500733535360091020500733535380005007335363000050073353631000500733536320005007335363300ccfdffff0500733536350005007335363600050073353637000500733536380005007335363900"}, {true, "00da000000050073353335000500733533360005007335333700050073353338000200000000510000000500733534310005007335343200830205007335343400003500000005007335343600050073353437000500733534380005007335343900dafdffff05007335353100050073353532000500733535330000510000000500733535350005007335353600910205007335353800003500000005007335363000050073353631000500733536320005007335363300ccfdffff0500733536350005007335363600050073353637000500733536380005007335363900"}}

	for _, c := range cases {
		want := func() UplaySection {
			v := NewUplaySection()
			v.Key = types.NewString("s535")
			v.Name = types.NewString("s536")
			v.TypeName = types.NewString("s537")
			v.MenuTypeName = types.NewString("s538")
			v.ContentList = types.List[UplaySectionContent]{func() UplaySectionContent {
				v := NewUplaySectionContent()
				v.Key = types.NewString("s541")
				v.Name = types.NewString("s542")
				v.Order = types.NewInt16(643)
				v.TypeName = types.NewString("s544")
				v.LocalizedInfo = func() UplaySectionContentLocalized {
					v := NewUplaySectionContentLocalized()
					v.Key = types.NewString("s546")
					v.Culture = types.NewString("s547")
					v.Text = types.NewString("s548")
					v.Url = types.NewString("s549")
					v.Duration = types.NewInt32(-550)
					v.Size = types.NewString("s551")
					v.Width = types.NewString("s552")
					v.Height = types.NewString("s553")
					return v
				}()
				return v
			}(), func() UplaySectionContent {
				v := NewUplaySectionContent()
				v.Key = types.NewString("s555")
				v.Name = types.NewString("s556")
				v.Order = types.NewInt16(657)
				v.TypeName = types.NewString("s558")
				v.LocalizedInfo = func() UplaySectionContentLocalized {
					v := NewUplaySectionContentLocalized()
					v.Key = types.NewString("s560")
					v.Culture = types.NewString("s561")
					v.Text = types.NewString("s562")
					v.Url = types.NewString("s563")
					v.Duration = types.NewInt32(-564)
					v.Size = types.NewString("s565")
					v.Width = types.NewString("s566")
					v.Height = types.NewString("s567")
					return v
				}()
				return v
			}()}
			v.GameCode = types.NewString("s568")
			v.PlatformCode = types.NewString("s569")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUplaySection()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UplaySection)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestUplayActionPlatformRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "050073353731000105007335373300"}, {true, "000f000000050073353731000105007335373300"}}

	for _, c := range cases {
		want := func() UplayActionPlatform {
			v := NewUplayActionPlatform()
			v.PlatformCode = types.NewString("s571")
			v.Completed = types.NewBool(true)
			v.SpecificKey = types.NewString("s573")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUplayActionPlatform()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UplayActionPlatform)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestUplayRewardPlatformRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733537350001"}, {true, "00080000000500733537350001"}}

	for _, c := range cases {
		want := func() UplayRewardPlatform {
			v := NewUplayRewardPlatform()
			v.PlatformCode = types.NewString("s575")
			v.Purchased = types.NewBool(true)
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUplayRewardPlatform()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UplayRewardPlatform)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestUplaySectionContentRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "0500733537380005007335373900a8020500733538310005007335383300050073353834000500733538350005007335383600b5fdffff050073353838000500733538390005007335393000"}, {true, "00510000000500733537380005007335373900a80205007335383100003500000005007335383300050073353834000500733538350005007335383600b5fdffff050073353838000500733538390005007335393000"}}

	for _, c := range cases {
		want := func() UplaySectionContent {
			v := NewUplaySectionContent()
			v.Key = types.NewString("s578")
			v.Name = types.NewString("s579")
			v.Order = types.NewInt16(680)
			v.TypeName = types.NewString("s581")
			v.LocalizedInfo = func() UplaySectionContentLocalized {
				v := NewUplaySectionContentLocalized()
				v.Key = types.NewString("s583")
				v.Culture = types.NewString("s584")
				v.Text = types.NewString("s585")
				v.Url = types.NewString("s586")
				v.Duration = types.NewInt32(-587)
				v.Size = types.NewString("s588")
				v.Width = types.NewString("s589")
				v.Height = types.NewString("s590")
				return v
			}()
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUplaySectionContent()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UplaySectionContent)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}

func TestUplaySectionContentLocalizedRoundTrip(t *testing.T) {
	cases := []struct {
		header bool
		data   string
	}{{false, "05007335393200050073353933000500733539340005007335393500acfdffff050073353937000500733539380005007335393900"}, {true, "003500000005007335393200050073353933000500733539340005007335393500acfdffff050073353937000500733539380005007335393900"}}

	for _, c := range cases {
		want := func() UplaySectionContentLocalized {
			v := NewUplaySectionContentLocalized()
			v.Key = types.NewString("s592")
			v.Culture = types.NewString("s593")
			v.Text = types.NewString("s594")
			v.Url = types.NewString("s595")
			v.Duration = types.NewInt32(-596)
			v.Size = types.NewString("s597")
			v.Width = types.NewString("s598")
			v.Height = types.NewString("s599")
			return v
		}()

		out := newOut(c.header)
		want.WriteTo(out)
		checkBytes(t, out.Bytes(), c.data)

		in := newIn(t, c.data, c.header)
		got := NewUplaySectionContentLocalized()
		if err := got.ExtractFrom(in); err != nil {
			t.Fatal(err)
		}

		if !want.Equals(got) {
			t.Fatalf("decoded value differs (header=%v):\n%s\nwant:\n%s", c.header, got, want)
		}

		if in.Remaining() != 0 {
			t.Fatalf("%d bytes left over", in.Remaining())
		}

		copied := want.Copy().(UplaySectionContentLocalized)
		if !want.Equals(copied) {
			t.Fatal("copy differs")
		}
	}
}
