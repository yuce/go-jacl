# Go-Jacl

This module implements the base and the extended specifications of the [Jacl configuration language](https://github.com/yuce/jacl).

## Change Log

### 0.1.0 (2019-06-30)

* Initial release.

## Installation

Requirements:

* Go 1.12 or better (not tried with versions below).

The following should be enough to install it:

    go get github.com/yuce/go-jacl

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
type Config struct {
    DataDir string `jacl:"data-dir"`
    Bind string `jacl:"bind"`
    Cluster ClusterConfig `jacl:"cluster"`
    TLS TLSConfig `jacl:"tls"`
    Gossip GossipConfig `jacl:"gossip"`
}

type ClusterConfig struct {
    Coordinator `jacl:"coordinator"`
    SomeLegacyField `jacl:"-"` // This field is skipped.
}

type TLSConfig struct {
    CertificatePath `jacl:"certificate"`
    KeyPath `jacl:"key"`
    SkipVerify `jacl:"skip-verify"`
}

type GossipConfig struct {
    Seeds []string `jacl:"seeds"`
    Port int `jacl:"port"`
    KeyPath string `jacl:"key"`
}

config := Config{}
err := jacl.Unmarshal(text, &config)
if err != nil {
    // handle the error
}
```

When decoding into structs:

* Only exported fields of the struct are considered.
* All fields of the struct must have corresponding values in the configuration, otherwise an error is returned.
* The field name must match the property/map key, unless `jacl:"KEY_NAME"` used in the field definition. In that case the configuration key `KEY_NAME` is matched to the corresponding field.
* Use `jacl:"-"` in order to skip a field.

See the [Jacl configuration language](https://github.com/yuce/jacl) for information about the configuration language.

## TODO

* Maps of typed slices and slices of typed maps are not yet supported when unmarshalling struct fields. E.g., `Field1 []map[string]interface{}` is OK, but `Field2 []map[string]int` is not supported yet.

## License

Copyright (c) 2019 Yuce Tekol. Licensed under [MIT](LICENSE).
