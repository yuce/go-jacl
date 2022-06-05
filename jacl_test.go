package jacl_test

import (
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/yuce/go-jacl"
)

func TestEmpty(t *testing.T) {
	target := map[string]interface{}{}
	compare(t, "empty", "", target)
	compare(t, "ignore comments", `
		// Comments are ignored

		/* Multiline comments
			too
		*/`, target)
}

func TestPropertyName(t *testing.T) {
	target := map[string]interface{}{"foo": "bar"}
	compare(t, "simple property name", `foo: "bar"`, target)

	target = map[string]interface{}{"foo": "bar"}
	compare(t, "quoted property name", `"foo": "bar"`, target)

	target = map[string]interface{}{"foo\\n\\t1": "bar"}
	compare(t, "escapes are not expanded", `"foo\n\t1": "bar"`, target)
}

func TestSetString(t *testing.T) {
	target := map[string]interface{}{"foo": "bar"}
	compare(t, "set string", `foo: "bar"`, target)

	target = map[string]interface{}{"fooz": `this is a
		multiline
		raw string`}
	compare(t, "set raw string", `fooz: '''this is a
		multiline
		raw string'''`, target)

	target = map[string]interface{}{"foo": "bar\nqoox"}
	compare(t, "set string with escape", `foo: "bar\nqoox"`, target)

	target = map[string]interface{}{"fooz": `this is a
		multiline raw string.
		It keeps escape\n\tcharacters.`}
	compare(t, "set raw string with single quotes", `fooz: '''this is a
		multiline raw string.
		It keeps escape\n\tcharacters.'''`, target)

	target = map[string]interface{}{"fooz": `this is another
		multiline raw string.
		It keeps escape\n\tcharacters.`}
	compare(t, "set raw string with double quotes", `fooz: """this is another
		multiline raw string.
		It keeps escape\n\tcharacters."""`, target)

	target = map[string]interface{}{
		"fooz": `this is a
		multiline '''raw''' string.
		It keeps escape\n\tcharacters.`,
		"barz": "yet another string",
	}
	compare(t, "set raw string with double quotes 2",
		`fooz: """this is a
		multiline '''raw''' string.
		It keeps escape\n\tcharacters."""
		barz: "yet another string"`, target)
}

func TestSetUnsignedInteger(t *testing.T) {
	target := map[string]interface{}{"mynum": uint64(2)}
	compare(t, "set bin uint", `mynum: 0b10`, target)

	target = map[string]interface{}{"mynum": uint64(2)}
	compare(t, "set bin uint with dash", `mynum: 0b1_0`, target)

	target = map[string]interface{}{"mynum": uint64(012345670)}
	compare(t, "set octal uint", `mynum: 0o12345670`, target)

	target = map[string]interface{}{"mynum": uint64(012345670)}
	compare(t, "set octal uint with dash", `mynum: 0o12_345_670`, target)

	target = map[string]interface{}{"mynum": uint64(123467890)}
	compare(t, "set decimal uint", `mynum: 0d123467890`, target)

	target = map[string]interface{}{"mynum": uint64(123467890)}
	compare(t, "set decimal uint with dash", `mynum: 0d123_467_890`, target)

	target = map[string]interface{}{"mynum": uint64(0x1234567890ABCDEF)}
	compare(t, "set hex uint", `mynum: 0x1234567890ABCDEF`, target)

	target = map[string]interface{}{"mynum": uint64(0x1234567890ABCDEF)}
	compare(t, "set hex uint with dash", `mynum: 0x1234_5678_90AB_CDEF`, target)
}

func TestSetSignedInteger(t *testing.T) {
	target := map[string]interface{}{"mynum": int64(123467890)}
	compare(t, "set int", `mynum: 123467890`, target)

	target = map[string]interface{}{"mynum": int64(123467890)}
	compare(t, "set int with dash", `mynum: 123_467_890`, target)

	target = map[string]interface{}{"mynum": int64(-123467890)}
	compare(t, "set int negative", `mynum: -123467890`, target)

	target = map[string]interface{}{"mynum": int64(123467890)}
	compare(t, "set int positive", `mynum: +123467890`, target)
}

func TestSetFloat(t *testing.T) {
	target := map[string]interface{}{"pi": float64(3.14159265358979323846264338327950288419716939937510582097494459)}
	compare(t, "set float", `pi: 3.14159265358979323846264338327950288419716939937510582097494459`, target)

	target = map[string]interface{}{"pi": float64(1233123.141592)}
	compare(t, "set float with dash", `pi: 1_233_123.141_592`, target)

	target = map[string]interface{}{"minus_pi": float64(-3.14159265358979323846264338327950288419716939937510582097494459)}
	compare(t, "set minus float", `minus_pi: -3.14159265358979323846264338327950288419716939937510582097494459`, target)

	target = map[string]interface{}{"some_num": float64(21.4e12)}
	compare(t, "set float with exp", `some_num: 21.4e12`, target)

	target = map[string]interface{}{"some_num": float64(21.4e12)}
	compare(t, "set float with exp", `some_num: 21.4e12`, target)
}

func TestSetBool(t *testing.T) {
	target := map[string]interface{}{"this_is_true": true}
	compare(t, "set true", `this_is_true: true`, target)

	target = map[string]interface{}{"this_is_false": false}
	compare(t, "set false", `this_is_false: false`, target)
}

func TestSetList(t *testing.T) {
	target := map[string]interface{}{"empty": []interface{}{}}
	compare(t, "set empty list", `empty: []`, target)

	target = map[string]interface{}{"ls": []interface{}{int64(1)}}
	compare(t, "set list with one item", `ls: [1]`, target)

	target = map[string]interface{}{"ls": []interface{}{int64(1), "foo"}}
	compare(t, "set list with two items", `ls: [1 "foo"]`, target)

	target = map[string]interface{}{"ls": []interface{}{int64(1), "foo"}}
	compare(t, "set list with two items with comma", `ls: [1, "foo"]`, target)

	target = map[string]interface{}{"ls": []interface{}{int64(1), "foo"}}
	compare(t, "set list with two items multiline", `ls: [
		1
		"foo"
	]`, target)

	target = map[string]interface{}{"ls": []interface{}{[]interface{}{"foo"}, "bar"}}
	compare(t, "set list in list", `ls: [
		["foo"]
		"bar"
	]`, target)
}

func TestSetMap(t *testing.T) {
	target := map[string]interface{}{"empty": map[string]interface{}{}}
	compare(t, "set empty map", `empty: {}`, target)

	target = map[string]interface{}{"mymap": map[string]interface{}{"key1": "value1"}}
	compare(t, "set map with one key", `mymap: {key1: "value1"}`, target)

	target = map[string]interface{}{"mymap": map[string]interface{}{"key with space in it": "value1"}}
	compare(t, "set map with one quoted key", `mymap: {"key with space in it": "value1"}`, target)

	target = map[string]interface{}{"mymap": map[string]interface{}{
		"key1": "value1",
		"key2": int64(44),
	}}
	compare(t, "set map with two keys", `mymap: {key1: "value1", "key2": 44}`, target)

	target = map[string]interface{}{"mymap": map[string]interface{}{
		"key1": "value1",
		"key2": int64(44),
	}}
	compare(t, "set map with two keys", `mymap: {
		key1: "value1"
		"key2": 44
	}`, target)

	target = map[string]interface{}{"mymap": map[string]interface{}{
		"key0": map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": "qoox",
			},
		},
		"key1": "value1",
	}}
	compare(t, "set map in map", `mymap: {
		key0: {
			foo: {
				bar: "qoox"
			}
		}
		key1: "value1"
	}`, target)

	target = map[string]interface{}{
		"ma1": map[string]interface{}{
			"key1": "value1",
			"key2": int64(44),
		},
		"ma2": map[string]interface{}{
			"key3": "value3",
		},
	}
	compare(t, "set map with two keys", `
		ma1: {
			key1: "value1"
			"key2": 44
		}
		ma2: {
			key3: "value3"
		}
	`, target)
}

func TestDuplicateKey(t *testing.T) {
	text := `
		a: 42
		a: 43
	`
	v := map[string]interface{}{}
	mustNotUnmarshal(t, text, &v)
}

func TestSyntaxError(t *testing.T) {
	text := `invalid key: "value"`
	m := map[string]interface{}{}
	err := jacl.Unmarshal(text, &m)
	if err == nil {
		t.Fatalf("should have failed")
	}
}

func TestDecodeStruct(t *testing.T) {
	type address struct {
		Street string
		Zip    int
	}

	type person struct {
		privateField string

		String string
		Int    int `jacl:"int"`
		Uint   uint
		Bool   bool
		Float  float64

		Slice       []interface{} `jacl:"slice"`
		StringSlice []string
		IntSlice    []int
		UintSlice   []uint
		BoolSlice   []bool
		FloatSlice  []float64
		MapSlice    []map[string]interface{}

		Map       map[string]interface{}
		StringMap map[string]string
		IntMap    map[string]int
		UintMap   map[string]uint
		BoolMap   map[string]bool
		FloatMap  map[string]float64
		ListMap   map[string][]interface{}

		// StringSliceMap map[string][]string

		Struct     address
		ListStruct []address
		MapStruct  map[string]address

		SkippedField string `jacl:"-"`
	}
	p := person{}
	text := `
		String: "Jane Doe"
		int: 22
		Uint: 0xBEE423
		Bool: false
		Float: -3.14e10

		slice: [1, true, "foo"]
		StringSlice: ["bar", "qoox"]
		IntSlice: [10, 20, 30]
		UintSlice: [0b10, 0o20, 0d30]
		BoolSlice: [false, true]
		FloatSlice: [100.2, -50.1e20]
		MapSlice: [
			{a: 1, b: 2}
			{c: "foo"}
		]

		Map: {key: true, x: 1}
		StringMap: {a: "foo", b: "bar"}
		IntMap: {a: 1, b: -2}
		UintMap: {a: 1, b: 2}
		BoolMap: {a: true, b: false}
		FloatMap: {a: -2.019, b:2.0e20}
		ListMap: {
			a: [1, 2, 3]
			b: ["foo", "bar"]
		}

		/*
		StringSliceMap: {
			a: ["foo", "bar"]
		}
		*/

		Struct: {
			Street: "Wonder Street"
			Zip: 34001
		}

		ListStruct: [
			{
				Street: "Under Street"
				Zip: 38004
			}
			{
				Street: "Over Street"
				Zip: 38005
			}
		]

		MapStruct: {
			primary: {
				Street: "Under Street"
				Zip: 38004
			}
			secondary: {
				Street: "Over Street"
				Zip: 38005
			}
		}
	`
	err := jacl.Unmarshal(text, &p)
	if err != nil {
		t.Fatal(err)
	}

	compareValue(t, "String", "Jane Doe", p.String)
	compareValue(t, "Int", 22, p.Int)
	compareValue(t, "Uint", uint(0xBEE423), p.Uint)
	compareValue(t, "Bool", false, p.Bool)
	compareValue(t, "Float", -3.14e10, p.Float)

	compareValue(t, "Slice", []interface{}{int64(1), true, "foo"}, p.Slice)
	compareValue(t, "StringSlice", []string{"bar", "qoox"}, p.StringSlice)
	compareValue(t, "IntSlice", []int{10, 20, 30}, p.IntSlice)
	compareValue(t, "UintSlice", []uint{uint(2), uint(020), uint(30)}, p.UintSlice)
	compareValue(t, "BoolSlice", []bool{false, true}, p.BoolSlice)
	compareValue(t, "FloatSlice", []float64{100.2, -50.1e20}, p.FloatSlice)
	compareValue(t, "MapSlice", []map[string]interface{}{
		map[string]interface{}{"a": int64(1), "b": int64(2)},
		map[string]interface{}{"c": "foo"},
	}, p.MapSlice)

	compareValue(t, "Map", map[string]interface{}{"key": true, "x": int64(1)}, p.Map)
	compareValue(t, "StringMap", map[string]string{"a": "foo", "b": "bar"}, p.StringMap)
	compareValue(t, "IntMap", map[string]int{"a": 1, "b": -2}, p.IntMap)
	compareValue(t, "UintMap", map[string]uint{"a": 1, "b": 2}, p.UintMap)
	compareValue(t, "BoolMap", map[string]bool{"a": true, "b": false}, p.BoolMap)
	compareValue(t, "FloatMap", map[string]float64{"a": -2.019, "b": 2.0e20}, p.FloatMap)
	compareValue(t, "ListMap", map[string][]interface{}{
		"a": []interface{}{int64(1), int64(2), int64(3)},
		"b": []interface{}{"foo", "bar"},
	}, p.ListMap)

	/*
		compareValue(t, "StringSliceMap", map[string][]string{
			"a": []string{"foo", "bar"},
		}, p.StringSliceMap)
	*/

	compareValue(t, "Struct", address{Street: "Wonder Street", Zip: 34001}, p.Struct)

	compareValue(t, "ListStruct", []address{
		address{Street: "Under Street", Zip: 38004},
		address{Street: "Over Street", Zip: 38005},
	}, p.ListStruct)

	compareValue(t, "MapStruct", map[string]address{
		"primary":   address{Street: "Under Street", Zip: 38004},
		"secondary": address{Street: "Over Street", Zip: 38005},
	}, p.MapStruct)

	compareValue(t, "SkippedField", "", p.SkippedField)
}

func TestReadmeSample(t *testing.T) {
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
	config1 := map[string]interface{}{}
	err := jacl.Unmarshal(text, config1)
	if err != nil {
		t.Fatal(err)
	}

	target1 := map[string]interface{}{
		"bind":     "https://01.pilosa.local:10101",
		"data-dir": "/tmp/data",
		"cluster": map[string]interface{}{
			"coordinator": true,
		},
		"tls": map[string]interface{}{
			"certificate": "pilosa.local.crt",
			"key":         "pilosa.local.key",
			"skip-verify": true,
		},
		"gossip": map[string]interface{}{
			"seeds": []interface{}{
				"01.pilosa.local:15000",
				"02.pilosa.local:15000",
			},
			"port": int64(15000),
			"key":  "pilosa.local.gossip32",
		},
	}
	if !reflect.DeepEqual(target1, config1) {
		t.Fatalf("\n%v\n!=\n%v", target1, config1)
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

	config2 := Config{}
	err = jacl.Unmarshal(text, &config2)
	if err != nil {
		t.Fatal(err)
	}

	target2 := Config{
		DataDir: "/tmp/data",
		Bind:    "https://01.pilosa.local:10101",
		Cluster: ClusterConfig{
			Coordinator: true,
		},
		TLS: TLSConfig{
			CertificatePath: "pilosa.local.crt",
			KeyPath:         "pilosa.local.key",
			SkipVerify:      true,
		},
		Gossip: GossipConfig{
			Seeds: []string{
				"01.pilosa.local:15000",
				"02.pilosa.local:15000",
			},
			Port:    15000,
			KeyPath: "pilosa.local.gossip32",
		},
	}

	if !reflect.DeepEqual(target2, config2) {
		t.Fatalf("\n%v\n!=\n%v", target2, config2)
	}
}

func TestUnmarshalStruct(t *testing.T) {
	type A struct {
		A1 string
		A2 int
	}
	type B struct {
		B bool
	}
	type Config struct {
		PropA A
		PropB bool
	}

	text1 := `
		PropA: {A1: "some-string", A2: 55}
	`
	text2 := `
		PropB: true
	`

	m := map[string]interface{}{}
	err := jacl.Unmarshal(text1, m)
	if err != nil {
		t.Fatal(err)
	}
	err = jacl.Unmarshal(text2, m)
	if err != nil {
		t.Fatal(err)
	}
	config := Config{}
	err = jacl.UnmarshalStruct(m, &config)
	if err != nil {
		t.Fatal(err)
	}
	target := Config{
		PropA: A{
			A1: "some-string",
			A2: 55,
		},
		PropB: true,
	}
	if !reflect.DeepEqual(target, config) {
		t.Fatalf("\n%v\n!=\n%v", target, config)
	}
}

func TestUnmarshalBoolField(t *testing.T) {
	type C struct {
		A bool
	}
	c := C{}

	mustUnmarshal(t, "A: true", &c)
	compareValue(t, "unmarshal bool true", C{A: true}, c)

	mustUnmarshal(t, "A: false", &c)
	compareValue(t, "unmarshal bool false", C{A: false}, c)

	mustNotUnmarshal(t, "A: 1", &c)
}

func TestUnmarshalIntField(t *testing.T) {
	type C struct {
		A int
	}
	c := C{}

	mustUnmarshal(t, "A: 0", &c)
	compareValue(t, "unmarshal int 0", C{A: 0}, c)

	mustUnmarshal(t, "A: 9223372036854775807", &c)
	compareValue(t, "unmarshal int MaxInt", C{A: jacl.MaxInt}, c)

	mustUnmarshal(t, "A: -9223372036854775808", &c)
	compareValue(t, "unmarshal int MinInt", C{A: jacl.MinInt}, c)

	mustNotUnmarshal(t, "A: 0d10", &c)

	type C8 struct {
		A int8
	}
	c8 := C8{}

	mustUnmarshal(t, "A: 0", &c8)
	compareValue(t, "unmarshal int8 0", C8{A: 0}, c8)

	mustUnmarshal(t, "A: 127", &c8)
	compareValue(t, "unmarshal int8 MaxInt8", C8{A: math.MaxInt8}, c8)

	mustUnmarshal(t, "A: -128", &c8)
	compareValue(t, "unmarshal int8 MinInt8", C8{A: math.MinInt8}, c8)

	mustNotUnmarshal(t, "A: 0d10", &c8)
	mustNotUnmarshal(t, "A: 128", &c8)
	mustNotUnmarshal(t, "A: -129", &c8)

	type C16 struct {
		A int16
	}
	c16 := C16{}

	mustUnmarshal(t, "A: 0", &c16)
	compareValue(t, "unmarshal int16 0", C16{A: 0}, c16)

	mustUnmarshal(t, "A: 32767", &c16)
	compareValue(t, "unmarshal int16 MaxInt16", C16{A: math.MaxInt16}, c16)

	mustUnmarshal(t, "A: -32768", &c16)
	compareValue(t, "unmarshal int16 MinInt16", C16{A: math.MinInt16}, c16)

	mustNotUnmarshal(t, "A: 0d10", &c16)
	mustNotUnmarshal(t, "A: 32768", &c16)
	mustNotUnmarshal(t, "A: -32769", &c16)

	type C32 struct {
		A int32
	}
	c32 := C32{}

	mustUnmarshal(t, "A: 0", &c32)
	compareValue(t, "unmarshal int32 0", C32{A: 0}, c32)

	mustUnmarshal(t, "A: 2147483647", &c32)
	compareValue(t, "unmarshal int32 MaxInt32", C32{A: math.MaxInt32}, c32)

	mustUnmarshal(t, "A: -2147483648", &c32)
	compareValue(t, "unmarshal int32 MinInt32", C32{A: math.MinInt32}, c32)

	mustNotUnmarshal(t, "A: 0d10", &c32)
	mustNotUnmarshal(t, "A: 2147483648", &c32)
	mustNotUnmarshal(t, "A: -2147483649", &c32)

	type C64 struct {
		A int64
	}
	c64 := C64{}

	mustUnmarshal(t, "A: 0", &c64)
	compareValue(t, "unmarshal int64 0", C64{A: 0}, c64)

	mustUnmarshal(t, "A: 9223372036854775807", &c64)
	compareValue(t, "unmarshal int64 MaxInt", C64{A: math.MaxInt64}, c64)

	mustUnmarshal(t, "A: -9223372036854775808", &c64)
	compareValue(t, "unmarshal int64 MinInt", C64{A: math.MinInt64}, c64)

	mustNotUnmarshal(t, "A: 0d10", &c64)
}

func TestUnmarshalUintField(t *testing.T) {
	type C struct {
		A uint
	}
	c := C{}

	mustUnmarshal(t, "A: 0d0", &c)
	compareValue(t, "unmarshal uint 0", C{A: 0}, c)

	mustUnmarshal(t, "A: 0xFFFFFFFFFFFFFFFF", &c)
	compareValue(t, "unmarshal uint MaxUint", C{A: jacl.MaxUint}, c)

	mustNotUnmarshal(t, "A: 10", &c)

	type C8 struct {
		A uint8
	}
	c8 := C8{}

	mustUnmarshal(t, "A: 0d0", &c8)
	compareValue(t, "unmarshal uint8 0", C8{A: 0}, c8)

	mustUnmarshal(t, "A: 0xFF", &c8)
	compareValue(t, "unmarshal uint8 MaxUint8", C8{A: math.MaxUint8}, c8)

	mustNotUnmarshal(t, "A: 10", &c8)
	mustNotUnmarshal(t, "A: 0x100", &c8)

	type C16 struct {
		A uint16
	}
	c16 := C16{}

	mustUnmarshal(t, "A: 0d0", &c16)
	compareValue(t, "unmarshal uint16 0", C16{A: 0}, c16)

	mustUnmarshal(t, "A: 0xFFFF", &c16)
	compareValue(t, "unmarshal uint16 MaxUint16", C16{A: math.MaxUint16}, c16)

	mustNotUnmarshal(t, "A: 10", &c16)
	mustNotUnmarshal(t, "A: 0x10000", &c16)

	type C32 struct {
		A uint32
	}
	c32 := C32{}

	mustUnmarshal(t, "A: 0d0", &c32)
	compareValue(t, "unmarshal uint32 0", C32{A: 0}, c32)

	mustUnmarshal(t, "A: 0xFFFFFFFF", &c32)
	compareValue(t, "unmarshal uint32 MaxUint32", C32{A: math.MaxUint32}, c32)

	mustNotUnmarshal(t, "A: 10", &c32)
	mustNotUnmarshal(t, "A: 0x100000000", &c32)

	type C64 struct {
		A uint64
	}
	c64 := C64{}

	mustUnmarshal(t, "A: 0d0", &c64)
	compareValue(t, "unmarshal uint 0", C64{A: 0}, c64)

	mustUnmarshal(t, "A: 0xFFFFFFFFFFFFFFFF", &c64)
	compareValue(t, "unmarshal uint64 MaxUint", C64{A: math.MaxUint64}, c64)

	mustNotUnmarshal(t, "A: 10", &c64)
}

func TestUnmarshalFloatField(t *testing.T) {
	type C32 struct {
		A float32
	}
	c32 := C32{}

	mustUnmarshal(t, "A: 0.0", &c32)
	compareValue(t, "unmarshal float32 0", C32{A: 0}, c32)

	mustUnmarshal(t, "A: 3.40282346638528859811704183484516925440e+38", &c32)
	compareValue(t, "unmarshal float32 MaxFloat", C32{A: math.MaxFloat32}, c32)

	mustUnmarshal(t, "A: -3.40282346638528859811704183484516925440e+38", &c32)
	compareValue(t, "unmarshal float32 MinFloat", C32{A: -math.MaxFloat32}, c32)

	mustNotUnmarshal(t, "A: 1", &c32)
	mustNotUnmarshal(t, "A: 1.797693134862315708145274237317043567981e+308", &c32)
	mustNotUnmarshal(t, "A: -1.797693134862315708145274237317043567981e+308", &c32)

	type C64 struct {
		A float64
	}
	c64 := C64{}

	mustUnmarshal(t, "A: 0.0", &c64)
	compareValue(t, "unmarshal float64 0", C64{A: 0}, c64)

	mustUnmarshal(t, "A: 1.797693134862315708145274237317043567981e+308", &c64)
	compareValue(t, "unmarshal float64 MaxFloat", C64{A: math.MaxFloat64}, c64)

	mustUnmarshal(t, "A: -1.797693134862315708145274237317043567981e+308", &c64)
	compareValue(t, "unmarshal float64 MinFloat", C64{A: -math.MaxFloat64}, c64)

	mustNotUnmarshal(t, "A: 1", &c64)
}

func TestStructDefaults(t *testing.T) {
	type C struct {
		F1 string
		F2 int
		F3 bool
	}

	defaults := map[string]interface{}{
		"F1": "default string",
		"F2": 54,
		"F3": true,
	}
	err := jacl.Unmarshal(`
		F1: "modified string"
	`, defaults)
	if err != nil {
		t.Fatal(err)
	}

	c := C{}
	err = jacl.UnmarshalStruct(defaults, &c)
	if err != nil {
		t.Fatal(err)
	}
	target := C{
		F1: "modified string",
		F2: 54,
		F3: true,
	}
	compareValue(t, "struct defaults", target, c)
}

func TestMultipleText(t *testing.T) {
	type C struct {
		F1 string
		F2 string
		F3 string
	}

	texts := []string{
		`F1: "field 1"`,
		`F2: "field 2"`,
		`F3: "field 3"`,
	}

	props := map[string]interface{}{}
	for _, text := range texts {
		err := jacl.Unmarshal(text, props)
		if err != nil {
			t.Fatal(err)
		}
	}

	config := C{}
	err := jacl.UnmarshalStruct(props, &config)
	if err != nil {
		t.Fatal(err)
	}

	target := C{
		F1: "field 1",
		F2: "field 2",
		F3: "field 3",
	}
	compareValue(t, "unmarshal multiple text", target, config)
}

func TestTrimText(t *testing.T) {
	text := `
		source: trim"""
			from __future__ import print_function
			import sys

			def main():
				args = sys.argv[1:]
				if args:
					print("%s arguments passed." % len(args))
					for i, arg in enumerate(args):
						print("  %s. %s" % (i, arg))
				else:
					print("No arguments passed.")

			if __name__ == "__main__":
				main()
			"""
		`
	type Doc struct {
		Source string `jacl:"source"`
	}
	doc := Doc{}
	err := jacl.Unmarshal(text, &doc)
	if err != nil {
		t.Fatal(err)
	}

	target := `from __future__ import print_function
import sys

def main():
	args = sys.argv[1:]
	if args:
		print("%s arguments passed." % len(args))
		for i, arg in enumerate(args):
			print("  %s. %s" % (i, arg))
	else:
		print("No arguments passed.")

if __name__ == "__main__":
	main()`
	if target != doc.Source {
		t.Fatalf("trim:\n%s\n!=\n%s", target, doc.Source)
	}
}

func TestPinText(t *testing.T) {
	text := "\n\t\tText: pin\"\"\"\n\n\t\t\t^\n\t\t\t\tSome text.\n\t\t\t\t\n\t\t\tMore text.\n\"\"\"\n"

	type Doc struct {
		Text string
	}
	doc := Doc{}
	err := jacl.Unmarshal(text, &doc)
	if err != nil {
		t.Fatal(err)
	}
	doc.Text = strings.Replace(doc.Text, "\t", ">", -1)

	target := "\tSome text.\n\t\nMore text.\n"
	target = strings.Replace(target, "\t", ">", -1)

	if target != doc.Text {
		t.Fatalf("pin:\n%s\n!=\n%s", target, doc.Text)
	}
}

func compare(t *testing.T, testName string, text string, target map[string]interface{}) {
	m := map[string]interface{}{}
	err := jacl.Unmarshal(text, m)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(target, m) {
		t.Fatalf("%s:\n%#v\n!=\n%#v", testName, target, m)
	}
}

func compareValue(t *testing.T, testName string, target, value interface{}) {
	if !reflect.DeepEqual(target, value) {
		t.Fatalf("%s:\n%#v (%s)\n!=\n%#v (%s)",
			testName, target, reflect.TypeOf(target), value, reflect.TypeOf(value))
	}
}

func mustUnmarshal(t *testing.T, text string, v interface{}) {
	err := jacl.Unmarshal(text, v)
	if err != nil {
		t.Fatal(err)
	}
}

func mustNotUnmarshal(t *testing.T, text string, v interface{}) {
	err := jacl.Unmarshal(text, v)
	if err == nil {
		t.Fatalf("should have failed")
	}
}
