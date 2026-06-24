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
		// 仅支持元素为 string/int*/uint* 的切片；其它元素类型暂不支持
		return newReflectSliceValue(elem, defaultValue)
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
		return sliceUsageType(v.elem.Type().Elem())
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

// reflectSliceValue 通过反射绑定到一个切片，行为与 *Value 系列一致：
// 逗号分隔解析 + 多次传参累加，支持 []string/[]int*/[]uint*
type reflectSliceValue struct {
	elem     reflect.Value
	elemType reflect.Type
	parse    elemParser
}

// elemParser 解析单个字符串为匹配 elemType 的 reflect.Value
type elemParser func(string) (reflect.Value, error)

// sliceElemParser 按 elemType.Kind 返回对应的元素解析器
func sliceElemParser(t reflect.Type) (elemParser, error) {
	switch t.Kind() {
	case reflect.String:
		return func(s string) (reflect.Value, error) {
			return reflect.ValueOf(s), nil
		}, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(s string) (reflect.Value, error) {
			n, err := strconv.ParseInt(strings.TrimSpace(s), 0, 64)
			if err != nil {
				return reflect.Value{}, err
			}
			v := reflect.New(t).Elem()
			v.SetInt(n)
			return v, nil
		}, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return func(s string) (reflect.Value, error) {
			n, err := strconv.ParseUint(strings.TrimSpace(s), 0, 64)
			if err != nil {
				return reflect.Value{}, err
			}
			v := reflect.New(t).Elem()
			v.SetUint(n)
			return v, nil
		}, nil
	default:
		return nil, fmt.Errorf("unsupported slice element type %s", t)
	}
}

// newReflectSliceValue 创建一个绑定到反射切片的 reflectSliceValue
func newReflectSliceValue(elem reflect.Value, defaultValue interface{}) (Value, error) {
	parse, err := sliceElemParser(elem.Type().Elem())
	if err != nil {
		return nil, fmt.Errorf("unsupported flag value type %s", elem.Type())
	}
	v := &reflectSliceValue{elem: elem, elemType: elem.Type().Elem(), parse: parse}
	v.setDefault(defaultValue)
	return v, nil
}

// setDefault 将默认值转换并写入切片元素
func (v *reflectSliceValue) setDefault(defaultValue interface{}) {
	slice := reflect.MakeSlice(reflect.SliceOf(v.elemType), 0, 0)
	switch data := defaultValue.(type) {
	case nil:
		// 空切片
	case string:
		for _, p := range splitStrings(data) {
			if ev, err := v.parse(p); err == nil {
				slice = reflect.Append(slice, ev)
			}
		}
	default:
		// 反射遍历任意切片/数组
		rv := reflect.ValueOf(defaultValue)
		if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
			for i := 0; i < rv.Len(); i++ {
				ev := rv.Index(i)
				if ev.Type() == v.elemType {
					slice = reflect.Append(slice, ev)
				} else if ev.Type().ConvertibleTo(v.elemType) {
					slice = reflect.Append(slice, ev.Convert(v.elemType))
				}
			}
		}
	}
	v.elem.Set(slice)
}

// Set 将逗号分隔的字符串解析并累加到切片
func (v *reflectSliceValue) Set(s string) error {
	parts := splitStrings(s)
	slice := v.elem
	for _, p := range parts {
		ev, err := v.parse(p)
		if err != nil {
			return err
		}
		slice = reflect.Append(slice, ev)
	}
	v.elem.Set(slice)
	return nil
}

// Get 返回切片值
func (v *reflectSliceValue) Get() interface{} { return v.elem.Interface() }

// String 返回切片的逗号连接字符串
func (v *reflectSliceValue) String() string {
	if v.elem.Len() == 0 {
		return ""
	}
	parts := make([]string, 0, v.elem.Len())
	for i := 0; i < v.elem.Len(); i++ {
		parts = append(parts, fmt.Sprintf("%v", v.elem.Index(i).Interface()))
	}
	return strings.Join(parts, ",")
}

// UsageType 返回类型显示名
func (v *reflectSliceValue) UsageType() string {
	return sliceUsageType(v.elemType)
}

// sliceUsageType 按 elemType.Kind 返回 UsageType 显示名
func sliceUsageType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "strings"
	case reflect.Int:
		return "ints"
	case reflect.Int8:
		return "int8s"
	case reflect.Int16:
		return "int16s"
	case reflect.Int32:
		return "int32s"
	case reflect.Int64:
		return "int64s"
	case reflect.Uint:
		return "uints"
	case reflect.Uint8:
		return "uint8s"
	case reflect.Uint16:
		return "uint16s"
	case reflect.Uint32:
		return "uint32s"
	case reflect.Uint64:
		return "uint64s"
	default:
		return "value"
	}
}
