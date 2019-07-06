package jacl

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/yuce/go-jacl/parser"
)

const maxStack = 16

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

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
			return fmt.Errorf("unmarshal error: can only unmarshal to maps with string keys or structs, not: %s", reflect.TypeOf(rv.Elem()))
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

// Unmarshal decodes the map given in `m` to v.
// `v` should be a pointer to a struct.
func UnmarshalStruct(m map[string]interface{}, v interface{}) (err error) {
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
	if rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("unmarshal error: can only unmarshal to structs, not: %s", reflect.TypeOf(rv.Elem()))
	}
	return unmarshalStruct(m, rv)
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
				if value.Kind() == reflect.String {
					value.SetString(t)
				} else {
					return fmt.Errorf("field %s is not string", field.Name)
				}
			case bool:
				if value.Kind() == reflect.Bool {
					value.SetBool(t)
				} else {
					return fmt.Errorf("field %s is not bool", field.Name)
				}
			case int:
				if err := setInt64(int64(t), value, field.Name); err != nil {
					return err
				}
			case int8:
				if err := setInt64(int64(t), value, field.Name); err != nil {
					return err
				}
			case int16:
				if err := setInt64(int64(t), value, field.Name); err != nil {
					return err
				}
			case int32:
				if err := setInt64(int64(t), value, field.Name); err != nil {
					return err
				}
			case int64:
				if err := setInt64(t, value, field.Name); err != nil {
					return err
				}
			case uint:
				if err := setUint64(uint64(t), value, field.Name); err != nil {
					return err
				}
			case uint8:
				if err := setUint64(uint64(t), value, field.Name); err != nil {
					return err
				}
			case uint16:
				if err := setUint64(uint64(t), value, field.Name); err != nil {
					return err
				}
			case uint32:
				if err := setUint64(uint64(t), value, field.Name); err != nil {
					return err
				}
			case uint64:
				if err := setUint64(t, value, field.Name); err != nil {
					return err
				}
			case float32:
				if err := setFloat64(float64(t), value, field.Name); err != nil {
					return err
				}
			case float64:
				if err := setFloat64(t, value, field.Name); err != nil {
					return err
				}
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

func setInt64(t int64, value reflect.Value, name string) error {
	switch value.Kind() {
	case reflect.Int:
		if t < int64(MinInt) || t > int64(MaxInt) {
			return fmt.Errorf("field %s cannot store %d without under/overflow", name, t)
		}
	case reflect.Int8:
		if t < int64(math.MinInt8) || t > int64(math.MaxInt8) {
			return fmt.Errorf("field %s cannot store %d without under/overflow", name, t)
		}
	case reflect.Int16:
		if t < int64(math.MinInt16) || t > int64(math.MaxInt16) {
			return fmt.Errorf("field %s cannot store %d without under/overflow", name, t)
		}
	case reflect.Int32:
		if t < int64(math.MinInt32) || t > int64(math.MaxInt32) {
			return fmt.Errorf("field %s cannot store %d without under/overflow", name, t)
		}
	case reflect.Int64:
		break
	default:
		return fmt.Errorf("field %s is not int", name)
	}
	value.SetInt(t)
	return nil
}

func setUint64(t uint64, value reflect.Value, name string) error {
	switch value.Kind() {
	case reflect.Uint:
		if t > uint64(MaxUint) {
			return fmt.Errorf("field %s cannot store %d without overflow", name, t)
		}
	case reflect.Uint8:
		if t > uint64(math.MaxUint8) {
			return fmt.Errorf("field %s cannot store %d without overflow", name, t)
		}
	case reflect.Uint16:
		if t > uint64(math.MaxUint16) {
			return fmt.Errorf("field %s cannot store %d without overflow", name, t)
		}
	case reflect.Uint32:
		if t > uint64(math.MaxUint32) {
			return fmt.Errorf("field %s cannot store %d without overflow", name, t)
		}
	case reflect.Uint64:
		break
	default:
		return fmt.Errorf("field %s is not uint", name)
	}
	value.SetUint(t)
	return nil
}

func setFloat64(t float64, value reflect.Value, name string) error {
	switch value.Kind() {
	case reflect.Float32:
		if t < float64(-math.MaxFloat32) || t > float64(math.MaxFloat32) {
			return fmt.Errorf("field %s cannot store %f without under/overflow", name, t)
		}

	case reflect.Float64:
		break
	default:
		return fmt.Errorf("field %s is not float", name)
	}
	value.SetFloat(t)
	return nil
}
