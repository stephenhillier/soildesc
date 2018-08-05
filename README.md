# soildesc
Generate database-friendly soil descriptions from a field or lab visual description


## Visual descriptions
Visual descriptions help engineering technologists and field engineers log soil conditions at a drilling site. A relatively plain-english, descriptive format is used :
> sand, some gravel, wet

Using well-known terms (soil types like "sand", "gravel", "silt", and adjectives like "wet", "compact", "jagged") allows field descriptions
to be easily understood by other engineers who can get a basic, general understanding of the lithology of a project site. 

Even though visual descriptions work well for humans, they are not always database friendly. `soildesc` aims to parse field descriptions
into a JSON-formatted object containing consistent terminology:

```
> sand, silty, wet

{
    "primary": "sand",
    "secondary": "silt",
    "moisture": "wet"
}
```

```
> loose water bearing silts, sands

{
    "primary": "silt",
    "secondary": "sand",
    "consistency": "loose",
    "moisture": "wet"
}
```
## Methodology
`soildesc` does not take an elegant approach. The text string is simply scanned, comparing each word
(as well as the one before it) with some pre-defined lists of common terms.

## Starting the API
The API is written in Go; building the server requires having the [Go language](http://www.golang.org/) installed.
It is recommended that you follow the instructions for setting up a Go environment and clone
the repository into your `src` dir:
```
$GOPATH/src/github.com/<username>/soildesc
```
The soildesc API has no dependencies aside from the standard library.

Running the unit tests:
```
go test
```

Building the binary:
```
go build
```

Running the server:
```
./soildesc
```

Send a test request:
```
curl -X POST localhost:8000/describe -d desc="very wet sand and gravel"
```
