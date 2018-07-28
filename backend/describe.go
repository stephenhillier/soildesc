package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/stephenhillier/soildesc/backend/models"
)

// Describe takes a soil description string and outputs a more structured response
func (s *Server) Describe(w http.ResponseWriter, req *http.Request) {
	desc := req.FormValue("desc")

	parsed := parse(desc)

	created, err := s.db.CreateDescription(parsed)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		return
	}

	response, err := json.Marshal(created)
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func parse(orig string) models.Description {
	d := models.Description{}
	d.Original = orig

	var singleWords []string

	for _, word := range strings.Split(orig, " ") {
		singleWords = append(singleWords, strings.Trim(word, ","))
	}

	// important description terms:
	// primary: gravel, sand, clay, silt
	// secondary: sandy, gravelly, silty, clayey, some gravel,
	// some sand, some silt, some clay, trace sand, trace gravel,
	// trace clay, trace silt

	baseType := make(map[string]string)
	baseType["gravelly"] = "gravel"
	baseType["gravels"] = "gravel"
	baseType["sandy"] = "sand"
	baseType["sands"] = "sand"
	baseType["silty"] = "silt"
	baseType["silts"] = "silt"
	baseType["clayey"] = "clay"
	baseType["clays"] = "clay"
	baseType["water bearing"] = "wet"
	baseType["water"] = "wet"

	terms := make(map[string][]string)

	// parsing a description works by brute force - words in the original description
	// are matched against the `terms` map.
	//
	// standard terminology is relatively limited, but this list could be stored
	// in a database in the future to allow adding more terms easily
	terms["primary"] = []string{"gravel", "sand", "clay", "silt"}
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
	terms["moisture"] = []string{"very dry", "very wet", "water bearing", "water", "dry", "damp", "moist", "wet"}

	var prev string
	var soil string
	var moisture string

	for _, word := range singleWords {

		// determine primary constituent
	primary:
		for _, term := range terms["primary"] {
			// select first matching term and check that it is not part of "some gravel", "trace silt" etc.
			if d.Primary == "" &&
				(word == term || word == term+"s") &&
				prev != "some" && prev != "trace" &&
				prev != "and" && prev != "&" {
				d.Primary = term
				break primary
			}

			// if a second "primary term" exists (e.g. sand and gravel), take the second as the secondary soil type.
			if d.Secondary == "" && (word == term || word == term+"s") {
				d.Secondary = term
				break primary
			}
		}

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
					// } else {
					// 	d.Additional = append(d.Additional, soil)
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

					log.Printf("%s %s", moisture, term)
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

	return d

}
