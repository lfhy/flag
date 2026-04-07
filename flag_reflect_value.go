package flag

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var durationType = reflect.TypeOf(time.Duration(0))

type usageTypeProvider interface {
	UsageType() string
}

type reflectValue struct {
	elem     reflect.Value
	duration bool
}

type reflectBoolValue struct {
	reflectValue
}

func (b *reflectBoolValue) IsBoolFlag() bool { return true }

func newReflectValue(target interface{}, defaultValue interface{}) (Value, error) {
	if target == nil {
		return nil, fmt.Errorf("flag value is nil")
	}

	elem := reflect.ValueOf(target)
	if elem.Kind() != reflect.Ptr {
		ptr := reflect.New(elem.Type())
		ptr.Elem().Set(elem)
		elem = ptr
	}

	if elem.IsNil() {
		return nil, fmt.Errorf("flag value pointer is nil")
	}

	elem = elem.Elem()
	if !elem.CanSet() {
		return nil, fmt.Errorf("flag value %T cannot be set", target)
	}

	useDuration := shouldUseDuration(elem.Type(), defaultValue)
	value := reflectValue{elem: elem, duration: useDuration}

	switch elem.Kind() {
	case reflect.Bool:
		elem.SetBool(anyToBool(defaultValue))
		return &reflectBoolValue{reflectValue: value}, nil
	case reflect.String:
		elem.SetString(anyToString(defaultValue))
		return &value, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if useDuration {
			elem.SetInt(int64(anyToTimeDuration(defaultValue)))
		} else {
			elem.SetInt(anyToInt64(defaultValue))
		}
		return &value, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		elem.SetUint(anyToUint64(defaultValue))
		return &value, nil
	case reflect.Float32, reflect.Float64:
		elem.SetFloat(anyToFloat64(defaultValue))
		return &value, nil
	default:
		return nil, fmt.Errorf("unsupported flag value type %T", target)
	}
}

func shouldUseDuration(t reflect.Type, defaultValue interface{}) bool {
	if t == durationType {
		return true
	}

	if t.Kind() != reflect.Int64 {
		return false
	}

	if _, ok := defaultValue.(time.Duration); ok {
		return true
	}

	if s, ok := defaultValue.(string); ok {
		_, err := time.ParseDuration(s)
		return err == nil
	}

	return false
}

func (v *reflectValue) Set(s string) error {
	switch v.elem.Kind() {
	case reflect.Bool:
		data, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		v.elem.SetBool(data)
		return nil
	case reflect.String:
		v.elem.SetString(s)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.duration {
			data, err := time.ParseDuration(s)
			if err != nil {
				return err
			}
			v.elem.SetInt(int64(data))
			return nil
		}
		data, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return err
		}
		v.elem.SetInt(data)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		data, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return err
		}
		v.elem.SetUint(data)
		return nil
	case reflect.Float32, reflect.Float64:
		data, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		v.elem.SetFloat(data)
		return nil
	default:
		return fmt.Errorf("unsupported flag value kind %s", v.elem.Kind())
	}
}

func (v *reflectValue) String() string {
	if v.duration {
		return time.Duration(v.elem.Int()).String()
	}

	if v.elem.Kind() == reflect.String {
		return v.elem.String()
	}

	return fmt.Sprintf("%v", v.elem.Interface())
}

func (v *reflectValue) UsageType() string {
	if v.duration {
		return "duration"
	}

	switch v.elem.Kind() {
	case reflect.String:
		return "string"
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "int"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return "uint"
	default:
		return "value"
	}
}
