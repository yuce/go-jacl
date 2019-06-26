package jacl

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/yuce/go-jacl/parser"
)

const maxStack = 16

type jaclMap map[string]interface{}

type errorListener struct {
	*antlr.DefaultErrorListener
}

func newErrorListener() *errorListener {
	return new(errorListener)
}

func (el *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("syntax error: %s line: %d column: %d symbol: %s", msg, line, column, offendingSymbol))
}

// Unmarshal decodes the Jacl configuration given in `text` to v.
// `v` should be either a map[string]interface{},
// or it should be a pointer to a struct.
func Unmarshal(text string, v interface{}) (err error) {
	defer func(err *error) {
		if r := recover(); r != nil {
			switch rt := r.(type) {
			case error:
				*err = rt
			case string:
				*err = errors.New(rt)
			default:
				panic(rt)
			}
		}
	}(&err)
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("non-nil pointer is required")
	}

	var rm map[string]interface{}
	var isStruct bool

	// if v is map[string]interface{} decode directly in it.
	switch m := rv.Elem().Interface().(type) {
	case map[string]interface{}:
		rm = m
	default:
		if rv.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("unmarshal error: can only unmarshal maps with string keys or structs, not: %s", reflect.TypeOf(rv.Elem()))
		}
		isStruct = true
		rm = map[string]interface{}{}
	}

	lexer := parser.NewJaclLexer(antlr.NewInputStream(text))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJaclParser(stream)
	p.AddErrorListener(newErrorListener())
	p.BuildParseTrees = true
	listener := newJaclListener(rm)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Config())

	if isStruct {
		err = unmarshalStruct(rm, rv)
	}

	return err
}

func unmarshalStruct(rm map[string]interface{}, value reflect.Value) error {
	elem := value.Elem()
	elemType := elem.Type()
	for i := 0; i < elemType.NumField(); i++ {
		value := elem.Field(i)
		if !value.CanSet() {
			continue
		}
		field := elemType.Field(i)
		propertyName := field.Tag.Get("jacl")
		if propertyName == "-" {
			// skip this field
			continue
		}
		if propertyName == "" {
			propertyName = field.Name
		}
		if configValue, ok := rm[propertyName]; ok {
			switch t := configValue.(type) {
			case string:
				value.SetString(t)
			case bool:
				value.SetBool(t)
			case int64:
				value.SetInt(t)
			case uint64:
				value.SetUint(t)
			case float64:
				value.SetFloat(t)
			case []interface{}:
				elemType := field.Type.Elem()
				switch elemType.Kind() {
				case reflect.Interface:
					// if this is a slice of interface{}, just assign it.
					value.Set(reflect.ValueOf(t))
				default:
					sl := reflect.MakeSlice(field.Type, len(t), len(t))
					// otherwise create and assign a slice of the given type.
					if elemType.Kind() == reflect.Struct {
						for i, tv := range t {
							if tvm, ok := tv.(map[string]interface{}); ok {
								st := reflect.New(elemType)
								unmarshalStruct(tvm, st)
								sl.Index(i).Set(st.Elem())
							} else {
								return fmt.Errorf("array %s doesn't contain maps", propertyName)
							}
						}
					} else {
						for i, tv := range t {
							sl.Index(i).Set(reflect.ValueOf(tv).Convert(elemType))
						}
					}
					value.Set(sl)
				}
			case map[string]interface{}:
				switch field.Type.Kind() {
				case reflect.Struct:
					st := reflect.New(field.Type)
					unmarshalStruct(t, st)
					value.Set(st.Elem())
				case reflect.Map:
					elemType := field.Type.Elem()
					switch elemType.Kind() {
					case reflect.Interface:
						// if this is a map of interface{}, just assign it.
						value.Set(reflect.ValueOf(t))
					case reflect.Struct:
						// this is a map of struct
						sm := reflect.MakeMapWithSize(field.Type, len(t))
						for tk, tv := range t {
							tkv := reflect.ValueOf(tk)
							if tvm, ok := tv.(map[string]interface{}); ok {
								st := reflect.New(elemType)
								unmarshalStruct(tvm, st)
								sm.SetMapIndex(tkv, st.Elem())
							} else {
								return fmt.Errorf("map %s doesn't contain maps", propertyName)
							}
						}
						value.Set(sm)
					default:
						// otherwise create and assign a map of the given type.
						sm := reflect.MakeMapWithSize(field.Type, len(t))
						for tk, tv := range t {
							tkv := reflect.ValueOf(tk)
							tvv := reflect.ValueOf(tv).Convert(elemType)
							sm.SetMapIndex(tkv, tvv)
						}
						value.Set(sm)
					}
				default:
					return fmt.Errorf("jacl unmarshal error: don't know how to unmarshal field: '%s'", field.Name)
				}
			default:
				return fmt.Errorf("jacl unmarshal error: don't know how to unmarshal field: '%s'", field.Name)
			}
		} else {
			return fmt.Errorf("jacl unmarshal error: property not found: '%s'", propertyName)
		}
	}

	return nil
}
