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

If running for the first time on a new machine (or the config file is not present in the directory your running the binary from) then run the command `./main -cfg` to generate a default configuration file.
I recommend then editing the configuration file to your needs.

Config fields:

```
output: string - the output directory you want to point your types to default is ./ (cwd)
file_name: string - the filename WITHOUT file extension. Because we only generate a typescript file, the tool appends .ts to the end by itself
prefixes: list of prefixes, each prefix is defined as such: name: string - the name of the cube that is known in metadata, prefix: string - the new prefix to give to the type.
```

Depending on the size of your cubejs cubes and the count of measures/dimensions, this tool should be finished within 1 second. So its "blazingly fast!!!!" but then written in Go.

If you are only using a binary then place the binary wherever you want and navigate to the binary using a terminal and execute it by using `./cube_type_extractor` or `./main` depending on if you downloaded the binary or built it from source.

#### Caveats

For now I have made the decision that we will ignore `error` cubes. Mostly because for us, we barely use the error cubes that come out of it. If it turns out to be that we need the error cube types then I will consider adding this.

## Requirements

- Cubejs running on a local machine (hardcoded url is set to localhost:4000, which is the default)
- Go installed on your machine (If running/building from source)
- You must edit your configuration file to your needs, otherwise it wont work with just the default configuration.

## Use case

For us, its mostly to create dynamically typed queries and to have typechecking that we are using the valid dimensions and or measures for our dashboarding using our own server action tooling for generic fetch methods.
For you probably not much other than its a fun little go script.

### Contributing

If you have a cool feature, or want to fix a bug, feel free to submit a PR.
