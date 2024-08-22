# Cube Type Extractor

A small utility tool I built for work to extract all dimensions and measures from our cubejs cubes.

This is a generic tool that should work for any use case.

## Usage

Simply run the command with go: `go run main.go` or run it remotely through `go run github.com/mikkurogue/cube-type-extract/blob/master/main.go`
It should then create a typescript file containing the string union types of each dimension and measure you have, prefixed with the cube name they are related to.

## Use case

For us, its mostly to create dynamically typed queries and to have typechecking that we are using the valid dimensions and or measures for our dashboarding.
For you probably not much.

## Planned

I'm planning on discussing internally with the team what property we would like to add to our dimensions to try and make the union types be less bloaty.
I'll then be updating and adding a flag to this tool to run the tool but fetching only the dimensions and measures with that specific property set to true.

For now I'm thinking of making the property be called `extractable` but a nicer name would be welcome too.
