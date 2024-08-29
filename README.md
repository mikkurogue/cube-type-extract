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


# Configuration File for Type Generator CLI Tool

This document provides an overview of the required configuration for the Type Generator CLI tool. The configuration is defined in a JSON format and is used to specify the output location, file name, and prefixes for the generated types.

## Configuration Structure

The configuration file should follow the structure outlined below:

```json
{
  "cube_url": "http://localhost:4000",
  "output": "./",
  "file_name": "cubejs-types",
  "prefixes": [
    {
      "name": "Placeholder",
      "prefix": "Main"
    }
  ]
}
```

## Fields
    
    cube_url:
        Type: string
        Description: The url for your cubejs api. Default is localhost:4000

    output:
        Type: string
        Description: Specifies the output directory where the generated type file will be saved. In this example, "./" indicates that the file will be saved in the current working directory.

    file_name:
        Type: string
        Description: Defines the name of the generated type file. In this example, the file will be named "cubejs-types".

    prefixes:

        Type: array of objects with keys name and prefix (strings) 

        Description: A list of prefix configurations that map entities to the prefixes you want to apply in the generated type file. Each entry in the list is an object with the following fields:
            name:
                Type: string
                Description: The name of the entity for which the types are being generated. In the example, "Placeholder" is used as a placeholder entity name.
            prefix:
                Type: string
                Description: The prefix to be applied to the entity's types in the generated file. In this example, "Main" is the prefix that will replace the entity name in the type definitions.


Depending on the size of your cubejs cubes and the count of measures/dimensions, this tool should be finished within 1 second. So its "blazingly fast!!!!" but then written in Go.

If you are only using a binary then place the binary wherever you want and navigate to the binary using a terminal and execute it by using `./cube_type_extractor` or `./main` depending on if you downloaded the binary or built it from source.

#### Caveats

For now I have made the decision that we will ignore `error` cubes. Mostly because for us, we barely use the error cubes that come out of it. If it turns out to be that we need the error cube types then I will consider adding this.

## Requirements

- [Cube](https://cube.dev/). Recommended to use a local instance of Cube, configuration defaults to `http://localhost:4000` as this is the default. Your mileage may vary.
- [Go](https://go.dev/) installed on your machine (If running/building from source)
- You must edit your configuration file to your needs, otherwise it wont work with just the default configuration.

## Use case

For us, its mostly to create dynamically typed queries and to have typechecking that we are using the valid dimensions and or measures for our dashboarding using our own server action tooling for generic fetch methods.
For you probably not much other than its a fun little go script.

### Contributing

If you have a cool feature, or want to fix a bug, feel free to submit a PR.

Feature requests are also possible, however as this is a side project I wont make many promises (unless I am already working on a feature for work and it fits the scope)
