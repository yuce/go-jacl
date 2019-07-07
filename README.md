<a href="https://travis-ci.org/yuce/go-jacl"><img src="https://api.travis-ci.org/yuce/go-jacl.svg?branch=master" alt="Build Status"></a>
<a href="https://coveralls.io/github/yuce/go-jacl?branch=master"><img src="https://coveralls.io/repos/github/yuce/go-jacl/badge.svg?branch=master" alt="Coverage Status" /></a>

# Go-Jacl

This module implements the base and the extended specifications of the [Jacl configuration language](https://github.com/yuce/jacl).

## Change Log

### 0.2.0 (2019-07-06)

* Added `jacl.UnmarshalStruct` function, which unmarshals a struct from a `map[string]interface{}`. This is used to achieve [default struct values](#default-struct-values) and [unmarshaling from multiple texts](#unmarshaling-from-multiple-texts).
* Added underflow and overflow checks for signed/unsigned integers and floats. See: [Field Underflow/Overflow](#field-underflow/overflow).

### 0.1.0 (2019-06-30)

* Initial release.

## Installation

Requirements:

* Go 1.12 or better (not tried with versions below).

The following should be enough to install it:

    go get github.com/yuce/go-jacl

## Jacl Specification

See the [Jacl configuration language](https://github.com/yuce/jacl) for information about the configuration language.

## Usage

Go-Jacl has a single function, `jacl.Unmarshal`, to decode configuration from text into a `map[string]interface{}` or a pointer to a struct.

Example:

```go
text := `
    // Sample configuration

    bind: "https://01.pilosa.local:10101"
    data-dir: "/tmp/data"

    cluster: {
        coordinator: true    
    }

    tls: {
        certificate: "pilosa.local.crt"
        key: "pilosa.local.key"
        skip-verify: true
    }

    gossip: {
        seeds: [
            "01.pilosa.local:15000"
            "02.pilosa.local:15000"
        ]
        port: 15000
        key: "pilosa.local.gossip32"
    }
`

// Decode to a map
config := map[string]interface{}{}
err := jacl.Unmarshal(text, &config)
if err != nil {
    // handle the error
}

// Decode to a struct
type ClusterConfig struct {
    Coordinator     bool   `jacl:"coordinator"`
    SomeLegacyField string `jacl:"-"` // This field is skipped.
}

type TLSConfig struct {
    CertificatePath string `jacl:"certificate"`
    KeyPath         string `jacl:"key"`
    SkipVerify      bool   `jacl:"skip-verify"`
}

type GossipConfig struct {
    Seeds   []string `jacl:"seeds"`
    Port    int      `jacl:"port"`
    KeyPath string   `jacl:"key"`
}

type Config struct {
    DataDir string        `jacl:"data-dir"`
    Bind    string        `jacl:"bind"`
    Cluster ClusterConfig `jacl:"cluster"`
    TLS     TLSConfig     `jacl:"tls"`
    Gossip  GossipConfig  `jacl:"gossip"`
}

config := Config{}
err := jacl.Unmarshal(text, &config)
if err != nil {
    // handle the error
}
```

## Decoding Into Structs

When decoding into structs:

* Only exported fields of the struct are considered.
* All fields of the struct must have corresponding values in the configuration, otherwise an error is returned.
* The field name must match the property/map key, unless `jacl:"KEY_NAME"` used in the field definition. In that case the configuration key `KEY_NAME` is matched to the corresponding field.
* Use `jacl:"-"` in order to skip a field.

### Supported Go Data Types

The following are the Go data types which are mapped from their Jacl counterparts. Note that trying to unmarshal to a field with a different type (e.g., a signed integer to `uint` vice versa, or a float to `int`) returns an error:

Jacl Type        | Go Type                | Allowed Field Types
-----------------|------------------------|--------------------
String           | string                 | string
Unsigned integer | uint64                 | uint, uint8, uint16, uint32, uint64
Signed integer   | int64                  | int, int8, int16, int32, int64
Float            | float64                | float32, float64
Boolean          | bool                   | bool
Array            | []interface{}          | []interface{}, []T
Map              | map[string]interface{} | map[string]interface{}, map[string]T

In the table above `T` is any type.

### Field Underflow/Overflow

If unmarshalling to a field underflows or overflows the chosen data type, then an error is returned:

```go
type Config struct {
    Number int8
}
config := Config{}
err := jacl.Unmarshal("Number: 128", &config)
```

`err` above is not `nil`, since 128 is bigger than math.MaxInt8.

### Default Struct Values

Go-Jacl requires every field of a struct to be set on unmarshal unless a field is skipped with `jacl:"-"`. So, if a property is missing the configuration `jacl.Unmarshal` would return an error.

Consider the following struct:

```go
type C struct {
    F1 string
    F2 int
}
```

In order to have a default for the field `F2`, we can pass a map with defaults to `jacl.Unmarshal`:

```go
defaults := map[string]interface{}{
    "F1": "default string",
    "F2": 54,
}

text := `
    F1: "modified string"
`
err := jacl.Unmarshal(text, &defaults)
if err != nil {
    // handle the error
}
```

The `defaults` map contains the following values after unmarshalling:

```go
map[string]interface{}{
    "F1": "modified string",
    "F2": int64(54),
}
```

Since the properties for all fields are set, we can pass that map to `jacl.UnmarshalStruct`:

```go
config := C{}
err = jacl.UnmarshalStruct(defaults, &config)
if err != nil {
    t.Fatal(err)
}
```

The value of `config` is:

```go
C{
    F1: "modified string",
    F2: 54,
}
```

### Unmarshaling From Multiple Texts

Suppose we separated our configuration into multiple files since there are lots of properties to be set. Instead of having separate structs for each file, we want to have a single struct. We can use the same `jack.UnmarshalStruct` technique in the previous section to accomplish that.

This is the sample struct:

```go
type C struct {
    F1 string
    F2 string
    // ...
    F9 string
}
```

These are the contents of the configuration files:

```go
texts := []string{
    `F1: "field 1"`,
    `F2: "field 2"`,
    // ...
    `F9: "field 9"`,
}
```

We use the same map to unmarshal each text:

```go
props := map[string]interface{}{}
for _, text := range texts {
    err := jacl.Unmarshal(text, &props)
    if err != nil {
        // handle the error
    }
}
```

Finally, unmarshal the map to a struct:

```go
config := C{}
err := jacl.UnmarshalStruct(props, &config)
if err != nil {
    // handle the error
}
```

## TODO

* Maps of typed slices and slices of typed maps are not yet supported when unmarshalling struct fields. E.g., `Field1 []map[string]interface{}` is OK, but `Field2 []map[string]int` is not supported yet.

## License

Copyright (c) 2019 Yuce Tekol. Licensed under [MIT](LICENSE).
