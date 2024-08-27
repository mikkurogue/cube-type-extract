# Cube Type Extractor

A small utility tool I built for work to extract all dimensions and measures from our cubejs cubes.

This is a generic tool that should work for any use case.

## Usage

Make sure you have alocal instance of CUbeJs running, and update your metadata for your dimensions and measures.

Thsi tool will only extract and create a union type based on if your meta property in your schemas has the property `extractable: true`. E.g.:

```
userName: {
    sql: (CUBE) => `${CUBE}.user_name`,
    type: `string`,
    meta: {
      extractable: true,
      // whatever other meta props you may have
    },
  },
```

Simply run the command with go: `go run main.go` or run it remotely through `go run github.com/mikkurogue/cube-type-extract/blob/master/main.go` (will create a binary thats built for all platforms at some point soon tm)
It should then create a typescript file containing the string union types of each dimension and measure you have, prefixed with the cube name they are related to.

Depending on the size of your cubejs cubes and the count of measures/dimensions, this tool should be finished within 1 second. So its "blazingly fast!!!!" but then written in Go.

If you are only using a binary then place the binary wherever you want and navigate to the binary using a terminal and execute it by using `./cube_type_extractor` or `./main` depending on if you downloaded the binary or built it from source.

### Supported input flags:

`-file="<FILENAME>"` Optional flag to rename the file to something you want. Do not provide a file extension as it always is a .ts file. If left empty will use the default name "cubejs-types"

`-rename=<true/false>` Optional flag to rename each cube to provide a new Suffix for the types. Default false.

## Requirements

- Cubejs running on a local machine (hardcoded url is set to localhost:4000, which is the default)
- Go installed on your machine (If running/building from source)

## Use case

For us, its mostly to create dynamically typed queries and to have typechecking that we are using the valid dimensions and or measures for our dashboarding.
For you probably not much.

### Contributing

If you have a cool feature, or want to fix a bug, feel free to submit a PR.
