package jacl_test

import (
	"reflect"
	"testing"

	"github.com/yuce/jacl"
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
