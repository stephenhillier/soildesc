package soildesc

import (
	"strings"
)

// Description contains information about a parsed soil visual description
type Description struct {
	Original    string   `json:"original" db:"original"`
	Ordered     []string `json:"ordered" db:"ordered"`
	Primary     string   `json:"primary" db:"primary"`
	Secondary   string   `json:"secondary" db:"secondary"`
	Consistency string   `json:"consistency" db:"consistency"`
	Moisture    string   `json:"moisture" db:"moisture"`
}

// ParseDescription takes an input string, scans it for keywords and fills
// a description struct type with a best guess for each category
// (primary, secondary soil etc)
//
// examples of input descriptions are "sandy gravel, very wet" or "water bearing silts".
// the output Description will contain consistent, standard terms for the primary soil type,
// secondary soil type, moisture content and consistency (loose, compact etc)
//
// TODO: this code started small, checking input against some limited cases
// Adding more cases and categories (e.g. moisture, consistency) has increased
// need for refactor.
func ParseDescription(orig string) (Description, error) {
	d := Description{}
	d.Original = orig

	var singleWords []string

	orig = strings.ToLower(orig)

	r := strings.NewReplacer(
		",", " ",
		"-", " ",
		":", " ",
		"\\", " ",
		"/", " ",
		";", " ",
		".", " ",
		"?", " ",
		"(", " ",
		")", " ",
	)

	orig = r.Replace(orig)

	for _, word := range strings.Split(orig, " ") {
		singleWords = append(singleWords, strings.Trim(word, ","))
	}

	// important description terms:
	// primary: gravel, sand, clay, silt
	// secondary: sandy, gravelly, silty, clayey, some gravel,
	// some sand, some silt, some clay, trace sand, trace gravel,
	// trace clay, trace silt

	baseType := map[string]string{
		"gravelly":      "gravel",
		"gravels":       "gravel",
		"sandy":         "sand",
		"sands":         "sand",
		"silty":         "silt",
		"silts":         "silt",
		"clayey":        "clay",
		"clays":         "clay",
		"water bearing": "wet",
		"water":         "wet",
	}

	terms := make(map[string][]string)

	// parsing a description works by brute force - words in the original description
	// are matched against the `terms` map.
	//
	// standard terminology is relatively limited, but this list could be stored
	// in a database in the future to allow adding more terms easily

	terms["primary"] = []string{
		"gravel",
		"sand",
		"clay",
		"silt",
		"hardpan",
		"soil",
		"bedrock",
		"muskeg",
		"topsoil",
		"mudstone",
		"granite",
		"conglomerate",
		"granodiorite",
		"basalt",
		"sandstone",
		"shale",
		"boulders",
		"cobbles",
		"gravels",
		"mud",
		"till",
		"rock",
		"gneiss",
		"quartz",
		"quartzite",
		"limestone",
		"pebbles",
		"organics",
		"volcanics",
		"feldspar"}

	terms["secondary"] = []string{
		"sandy",
		"gravelly",
		"silty",
		"clayey",
		"some sand",
		"some gravel",
		"some silt",
		"some clay",
		"trace sand",
		"trace gravel",
		"trace silt",
		"trace clay",
	}

	// consistency terms (firmness/looseness of material)
	terms["consistency"] = []string{"loose", "soft", "firm", "compact", "hard", "dense"}

	// moisture content terms
	// some terms will be converted to a more "standard" one (e.g. "water bearing" will beocme "wet")
	// via the baseType map
	terms["moisture"] = []string{"very dry", "very wet", "water bearing", "water", "dry", "damp", "moist", "wet"}

	var prev string
	var soil string
	var moisture string

	for _, word := range singleWords {
		// determine primary constituent before moving on to other properties
	primary:
		for _, term := range terms["primary"] {

			// select first matching term and check that it is not part of "some gravel", "trace silt" etc.
			if (word == term || word == term+"s") && prev != "some" && prev != "trace" {
				if d.Primary == "" && prev != "and" && prev != "&" {
					d.Primary = term
					d.Ordered = append([]string{term}, d.Ordered...)
				} else if d.Secondary == "" {
					// some secondary soil types might come in the form "sand and gravel" (e.g. gravel will be secondary)
					// we can catch these while searching for primary terms
					d.Secondary = term
					d.Ordered = append(d.Ordered, term)
					break primary
				}
			}
		}
		prev = word
	}

	prev = "" // reset prev to an empty string before iterating again

	for _, word := range singleWords {
		// determine secondary constituent(s) for -ly terms (gravelly etc.)
		for _, term := range terms["secondary"] {
			if word == term || prev+" "+word == term {
				// words like "gravelly" need to be converted
				soil = baseType[word]
				// if soil is not in baseType map, default to current word
				if soil == "" {
					soil = word
				}

				if d.Secondary == "" {
					d.Secondary = soil
					d.Ordered = append(d.Ordered, soil)
				} else {
					d.Ordered = append(d.Ordered, soil)
				}
			}
		}

		if d.Consistency == "" {
		consistency:
			// determine consistency
			for _, term := range terms["consistency"] {
				if word == term {
					d.Consistency = word
					break consistency
				}
			}
		}

		if d.Moisture == "" {
		moisture:
			for _, term := range terms["moisture"] {
				if word == term || prev+" "+word == term {
					moisture = baseType[term]

					// if soil is not in baseType map, default to current word
					if moisture == "" {
						d.Moisture = term
					} else {
						d.Moisture = moisture
					}
					break moisture
				}
			}
		}

		prev = word

	}

	return d, nil
}
