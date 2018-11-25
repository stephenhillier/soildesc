package soildesc

import (
	"testing"
)

func TestParseDescription(t *testing.T) {
	cases := []struct {
		desc string
		want Description
	}{
		{"wet gravel", Description{Primary: "gravel", Moisture: "wet", Ordered: []string{"gravel"}}},
		{"compact silty sand, some clay, wet", Description{Primary: "sand", Secondary: "silt", Consistency: "compact", Moisture: "wet", Ordered: []string{"sand", "silt", "clay"}}},
		{"water bearing sands, trace gravel, loose", Description{Primary: "sand", Secondary: "gravel", Consistency: "loose", Moisture: "wet", Ordered: []string{"sand", "gravel"}}},

		// note: this is a poor sescription.  silty should come last.  here we just make sure it is handled as if silty came after gravel
		{"silty sand and gravel", Description{Primary: "sand", Secondary: "gravel", Ordered: []string{"sand", "gravel", "silt"}}},
		{"sand and gravel, silty", Description{Primary: "sand", Secondary: "gravel", Ordered: []string{"sand", "gravel", "silt"}}},

		// from "Applications of Artifial Intelligence in Engineering VI" by G. Rzevzky, R.A. Adey 2012
		// {
		// 	"Moist stiff reddish brown closely fissured thinly bedded silty sand CLAY with a little dark greenish grey sub-rounded fine gravel",
		// 	Description{Primary: "clay", Secondary: "sand", Ordered: []string{"clay", "sand", "silt", "gravel"}},
		// },
	}

	for _, test := range cases {
		desc, err := ParseDescription(test.desc)
		if err != nil {
			t.Errorf("Error running ParseDescription function")
		}
		if desc.Primary != test.want.Primary {
			t.Errorf("Primary soil was incorrect. got: %s, want: %s", desc.Primary, test.want.Primary)
		}
		if desc.Secondary != test.want.Secondary {
			t.Errorf("Secondary soil was incorrect. got: %s, want: %s", desc.Secondary, test.want.Secondary)
		}
		if desc.Consistency != test.want.Consistency {
			t.Errorf("Consistency was incorrect. got: %s, want: %s", desc.Consistency, test.want.Consistency)
		}
		if desc.Moisture != test.want.Moisture {
			t.Errorf("Moisture was incorrect. got: %s, want: %s", desc.Moisture, test.want.Moisture)
		}
		if !testEq(desc.Ordered, test.want.Ordered) {
			t.Errorf("Ordered terms were incorrect. got: %s, want: %s", desc.Ordered, test.want.Ordered)

		}
	}
}

func TestParseSoilTermsV2(t *testing.T) {
	cases := []struct {
		desc string
		want []string
	}{
		{"wet gravel", []string{"gravel"}},
		{"compact silty sand, some clay, wet", []string{"sand", "silt", "clay"}},
		{"water bearing sands, trace gravel, loose", []string{"sand", "gravel"}},

		// note: this is a poor sescription.  silty should come last.  here we just make sure it is handled as if silty came after gravel
		{"silty sand and gravel", []string{"sand", "gravel", "silt"}},
		{"SAND and GRAVEL, silty", []string{"sand", "gravel", "silt"}},

		// from "Applications of Artificial Intelligence in Engineering VI" by G. Rzevzky, R.A. Adey 2012
		// {
		// 	"Moist stiff reddish brown closely fissured thinly bedded silty sand CLAY with a little dark greenish grey sub-rounded fine gravel",
		// 	[]string{"clay", "sand", "silt", "gravel"},
		// },
	}

	for _, test := range cases {
		desc := ParseSoilTerms(test.desc)

		if !testEq(desc, test.want) {
			t.Errorf("Ordered terms were incorrect. got: %s, want: %s", desc, test.want)
		}
	}
}

func testEq(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
