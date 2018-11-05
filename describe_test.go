package main

import (
	"testing"
)

func TestParseDescription(t *testing.T) {
	cases := []struct {
		desc string
		want description
	}{
		{"wet gravel", description{Primary: "gravel", Moisture: "wet", Ordered: []string{"gravel"}}},
		{"compact silty sand, some clay, wet", description{Primary: "sand", Secondary: "silt", Consistency: "compact", Moisture: "wet", Ordered: []string{"sand", "silt", "clay"}}},
		{"water bearing sands, trace gravel, loose", description{Primary: "sand", Secondary: "gravel", Consistency: "loose", Moisture: "wet", Ordered: []string{"sand", "gravel"}}},

		// note: this is a poor description.  silty should come last.  here we just make sure it is handled as if silty came after gravel
		{"silty sand and gravel", description{Primary: "sand", Secondary: "gravel", Ordered: []string{"sand", "gravel", "silt"}}},
		{"sand and gravel, silty", description{Primary: "sand", Secondary: "gravel", Ordered: []string{"sand", "gravel", "silt"}}},
	}

	for _, test := range cases {
		desc := parseDescription(test.desc)
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
