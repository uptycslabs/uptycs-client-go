## Uptycs Client (Go)

A Go library for [uptycs](https://uptycs.io)

## Contributing

Addition of a new config type requires the baseline addition of several files and a modification to a few other files.

Each config type needs
- an entry (or multiple) in models.go describing the exact structure of that object type
- a file in fixtures file with the expected output from the get call to the api
- a go file in the uptycs folder following the naming convention of the rest of the config types
- a second go file in the uptycs folder named identically to the previous file with _test added

The entry within the models.go file will come from the output of the get call for that particular config type. Note, if there are options to make calls to the api for one or multiple items, the singular and plural versions of these structs should be made seperately. User and Users / Role and Roles can be used as examples for this structure.



### Examples ###

See the `_examples/` dir
