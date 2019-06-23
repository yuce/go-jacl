package jacl_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/yuce/go-jacl"
)

func TestEmpty(t *testing.T) {
	target := map[string]interface{}{}
	compare(t, "empty", "", target)
	compare(t, "ignore comments", `
		# Comments are ignored

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
	target := map[string]interface{}{"mynum": uint64(21)}
	compare(t, "set bin uint", `mynum: 0b10101`, target)

	target = map[string]interface{}{"mynum": uint64(0123)}
	compare(t, "set octal uint", `mynum: 0o123`, target)

	target = map[string]interface{}{"mynum": uint64(567)}
	compare(t, "set decimal uint", `mynum: 0d567`, target)

	target = map[string]interface{}{"mynum": uint64(0xBEEF)}
	compare(t, "set hex uint", `mynum: 0xBEEF`, target)
}

func TestSetSignedInteger(t *testing.T) {
	target := map[string]interface{}{"mynum": int64(21)}
	compare(t, "set int", `mynum: 21`, target)

	target = map[string]interface{}{"mynum": int64(-21)}
	compare(t, "set int negative", `mynum: -21`, target)

	target = map[string]interface{}{"mynum": int64(21)}
	compare(t, "set int positive", `mynum: +21`, target)
}

func TestSetFloat(t *testing.T) {
	target := map[string]interface{}{"pi": float64(3.14159265358979323846264338327950288419716939937510582097494459)}
	compare(t, "set float", `pi: 3.14159265358979323846264338327950288419716939937510582097494459`, target)

	target = map[string]interface{}{"minus_pi": float64(-3.14159265358979323846264338327950288419716939937510582097494459)}
	compare(t, "set minus float", `minus_pi: -3.14159265358979323846264338327950288419716939937510582097494459`, target)

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
	err := jacl.Unmarshal(text, &m)
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
