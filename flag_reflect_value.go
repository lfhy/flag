package flag

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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
	case reflect.Slice:
		// 仅支持 []string 切片；其它元素类型暂不支持
		if elem.Type().Elem().Kind() != reflect.String {
			return nil, fmt.Errorf("unsupported flag value type %T", target)
		}
		return newReflectStringsValue(elem, defaultValue), nil
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
	case reflect.Slice:
		if v.elem.Type().Elem().Kind() == reflect.String {
			return "strings"
		}
		return "value"
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

// reflectStringsValue 通过反射绑定到一个 []string 切片，
// 行为与 stringsValue 一致：逗号分隔解析 + 多次传参累加
type reflectStringsValue struct {
	elem reflect.Value
}

// newReflectStringsValue 创建一个绑定到反射 []string 元素的 reflectStringsValue，
// 并把默认值（逗号分隔字符串或 []string）写入元素
func newReflectStringsValue(elem reflect.Value, defaultValue interface{}) *reflectStringsValue {
	v := &reflectStringsValue{elem: elem}
	v.setDefault(defaultValue)
	return v
}

// setDefault 将默认值转换并写入切片元素
func (v *reflectStringsValue) setDefault(defaultValue interface{}) {
	switch data := defaultValue.(type) {
	case nil:
		v.elem.Set(reflect.ValueOf([]string{}))
	case []string:
		v.elem.Set(reflect.ValueOf(data))
	case []interface{}:
		out := make([]string, 0, len(data))
		for _, d := range data {
			out = append(out, anyToString(d))
		}
		v.elem.Set(reflect.ValueOf(out))
	case string:
		v.elem.Set(reflect.ValueOf(splitStrings(data)))
	default:
		v.elem.Set(reflect.ValueOf([]string{}))
	}
}

// Set 将逗号分隔的字符串解析并累加到切片
func (v *reflectStringsValue) Set(s string) error {
	parts := splitStrings(s)
	if v.elem.Len() == 0 {
		v.elem.Set(reflect.ValueOf(parts))
	} else {
		v.elem.Set(reflect.AppendSlice(v.elem, reflect.ValueOf(parts)))
	}
	return nil
}

// Get 返回切片值
func (v *reflectStringsValue) Get() interface{} { return v.elem.Interface() }

// String 返回切片的逗号连接字符串
func (v *reflectStringsValue) String() string {
	if v.elem.Len() == 0 {
		return ""
	}
	parts := make([]string, 0, v.elem.Len())
	for i := 0; i < v.elem.Len(); i++ {
		parts = append(parts, v.elem.Index(i).String())
	}
	return strings.Join(parts, ",")
}

// UsageType 返回类型显示名
func (v *reflectStringsValue) UsageType() string { return "strings" }
