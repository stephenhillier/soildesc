# soildesc
Parse terms from a visual description of a soil sample


## Visual descriptions
Visual descriptions help engineering technologists and field engineers log soil conditions at a drilling site. A relatively plain-english, descriptive format is used :
> sand, some gravel, wet

Soildesc parses these descriptions and returns data in a format that is easier to use with a database or programming language:

```
> sand, silty, wet

{
    "ordered": ["sand", "silt"],
    "primary": "sand",
    "secondary": "silt",
    "moisture": "wet"
}
```

```
> loose water bearing silts, sands

{
    "ordered": ["silt", "sand"],
    "primary": "silt",
    "secondary": "sand",
    "consistency": "loose",
    "moisture": "wet"
}
```
## Methodology
The text string is simply scanned, comparing each word (as well as the one before it) with some pre-defined lists of common terms.

## Usage

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
