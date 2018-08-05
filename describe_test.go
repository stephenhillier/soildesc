package main

import (
	"testing"
)

func TestParseDescription(t *testing.T) {
	cases := []struct {
		desc string
		want description
	}{
		{"wet gravel", description{Primary: "gravel", Moisture: "wet"}},
		{"compact silty sand, some clay, wet", description{Primary: "sand", Secondary: "silt", Consistency: "compact", Moisture: "wet"}},
		{"water bearing sands, trace gravel, loose", description{Primary: "sand", Secondary: "gravel", Consistency: "loose", Moisture: "wet"}},
		{"silty sand and gravel", description{Primary: "sand", Secondary: "gravel"}},
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
	}
}
