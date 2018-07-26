package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type formattedDescription struct {
	Original    string   `json:"original"`
	Primary     string   `json:"primary"`
	Secondary   []string `json:"secondary"`
	Consistency string   `json:"consistency"`
}

// Describe takes a soil description string and outputs a more structured response
func (s *Server) Describe(w http.ResponseWriter, req *http.Request) {
	defer func(t time.Time) {
		log.Printf("%s: %s (%v): %s", req.Method, req.URL.Path, time.Since(t), req.FormValue("desc"))
	}(time.Now())

	desc := req.FormValue("desc")

	response, err := json.Marshal(parse(desc))
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func parse(orig string) formattedDescription {
	d := formattedDescription{}
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
	baseType["sandy"] = "sand"
	baseType["silty"] = "silt"
	baseType["clayey"] = "clay"

	terms := make(map[string][]string)

	terms["primary"] = []string{"gravel", "sand", "clay", "silt"}
	terms["secondary"] = []string{"sandy", "gravelly", "silty", "clayey"}

	// consistency terms
	terms["consistency"] = []string{"loose", "soft", "firm", "compact", "hard", "dense"}

	var prev string

	// determine primary constituent
primary:
	for _, word := range singleWords {
		for _, term := range terms["primary"] {
			// select first matching term and check that it is not part of "some gravel", "trace silt" etc.
			if word == term && prev != "some" && prev != "trace" {
				d.Primary = strings.ToUpper(word)
				break primary
			}
			prev = word
		}
	}

	// determine secondary constituent(s) for -ly terms (gravelly etc.)
	for _, word := range singleWords {
		for _, term := range terms["secondary"] {
			if word == term {
				d.Secondary = append(d.Secondary, baseType[word])
			}
		}
	}

consistency:
	// determine consistency
	for _, word := range singleWords {
		for _, term := range terms["consistency"] {
			if word == term {
				d.Consistency = word
				break consistency
			}
		}
	}

	return d

}
