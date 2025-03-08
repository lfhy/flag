package flag

import (
	"fmt"
	"strconv"
	"time"
)

// -- bool Value
type boolValue bool

// newBoolValue 创建一个新的boolValue类型的值，并将其赋值给指针p
func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return (*boolValue)(p)
}

// Set 将字符串s转换为bool类型，并赋值给指针b
func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

// Get 返回boolValue类型的值
func (b *boolValue) Get() interface{} { return bool(*b) }

// String 返回boolValue类型的值的字符串表示
func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

// IsBoolFlag 返回boolValue类型是否为布尔值
func (b *boolValue) IsBoolFlag() bool { return true }

// optional interface to indicate boolean flags that can be
// supplied without "=value" text
type boolFlag interface {
	Value
	IsBoolFlag() bool
}

// -- int Value
type intValue int

// newIntValue 创建一个新的intValue类型的值，并将其赋值给指针p
func newIntValue(val int, p *int) *intValue {
	*p = val
	return (*intValue)(p)
}

// Set 将字符串s转换为int类型，并赋值给指针i
func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

// Get 返回intValue类型的值
func (i *intValue) Get() interface{} { return int(*i) }

// String 返回intValue类型的值的字符串表示
func (i *intValue) String() string { return fmt.Sprintf("%v", *i) }

// -- int64 Value
type int64Value int64

// newInt64Value 创建一个新的int64Value类型的值，并将其赋值给指针p
func newInt64Value(val int64, p *int64) *int64Value {
	*p = val
	return (*int64Value)(p)
}

// Set 将字符串s转换为int64类型，并赋值给指针i
func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = int64Value(v)
	return err
}

// Get 返回int64Value类型的值
func (i *int64Value) Get() interface{} { return int64(*i) }

// String 返回int64Value类型的值的字符串表示
func (i *int64Value) String() string { return fmt.Sprintf("%v", *i) }

// -- uint Value
type uintValue uint

// newUintValue 创建一个新的uintValue类型的值，并将其赋值给指针p
func newUintValue(val uint, p *uint) *uintValue {
	*p = val
	return (*uintValue)(p)
}

// Set 将字符串s转换为uint类型，并赋值给指针i
func (i *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uintValue(v)
	return err
}

// Get 返回uintValue类型的值
func (i *uintValue) Get() interface{} { return uint(*i) }

// String 返回uintValue类型的值的字符串表示
func (i *uintValue) String() string { return fmt.Sprintf("%v", *i) }

// -- uint64 Value
type uint64Value uint64

// newUint64Value 创建一个新的uint64Value类型的值，并将其赋值给指针p
func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val
	return (*uint64Value)(p)
}

// Set 将字符串s转换为uint64类型，并赋值给指针i
func (i *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uint64Value(v)
	return err
}

// Get 返回uint64Value类型的值
func (i *uint64Value) Get() interface{} { return uint64(*i) }

// String 返回uint64Value类型的值的字符串表示
func (i *uint64Value) String() string { return fmt.Sprintf("%v", *i) }

// -- string Value
type stringValue string

// newStringValue 创建一个新的stringValue类型的值，并将其赋值给指针p
func newStringValue(val string, p *string) *stringValue {
	*p = val
	return (*stringValue)(p)
}

// Set 将字符串val赋值给指针s
func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

// Get 返回stringValue类型的值
func (s *stringValue) Get() interface{} { return string(*s) }

// String 返回stringValue类型的值的字符串表示
func (s *stringValue) String() string { return string(*s) }

// -- float64 Value
type float64Value float64

// newFloat64Value 创建一个新的float64Value类型的值，并将其赋值给指针p
func newFloat64Value(val float64, p *float64) *float64Value {
	*p = val
	return (*float64Value)(p)
}

// Set 将字符串s转换为float64类型，并赋值给指针f
func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = float64Value(v)
	return err
}

// Get 返回float64Value类型的值
func (f *float64Value) Get() interface{} { return float64(*f) }

// String 返回float64Value类型的值的字符串表示
func (f *float64Value) String() string { return fmt.Sprintf("%v", *f) }

// -- time.Duration Value
type durationValue time.Duration

// newDurationValue 创建一个新的durationValue类型的值，并将其赋值给指针p
func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val
	return (*durationValue)(p)
}

// Set 将字符串s转换为time.Duration类型，并赋值给指针d
func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

// Get 返回durationValue类型的值
func (d *durationValue) Get() interface{} { return time.Duration(*d) }

// String 返回durationValue类型的值的字符串表示
func (d *durationValue) String() string { return (*time.Duration)(d).String() }
