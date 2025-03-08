package flag

import (
	"fmt"
	"runtime"
)

// 定义颜色常量
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

// 将字符串转换为黑色
func Black(str string) string {
	return textColor(TextBlack, str)
}

// 将字符串转换为红色
func Red(str string) string {
	return textColor(TextRed, str)
}

// 将字符串转换为绿色
func Green(str string) string {
	return textColor(TextGreen, str)
}

// 将字符串转换为黄色
func Yellow(str string) string {
	return textColor(TextYellow, str)
}

// 将字符串转换为蓝色
func Blue(str string) string {
	return textColor(TextBlue, str)
}

// 将字符串转换为品红色
func Magenta(str string) string {
	return textColor(TextMagenta, str)
}

// 将字符串转换为青色
func Cyan(str string) string {
	return textColor(TextCyan, str)
}

// 将字符串转换为白色
func White(str string) string {
	return textColor(TextWhite, str)
}

// 根据颜色将字符串转换为对应的颜色
func textColor(color int, str string) string {
	if runtime.GOOS == "windows" {
		return str
	}

	switch color {
	case TextBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
	case TextMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
	case TextCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}

// 打印错误信息
func StdErrorln(a ...any) {
	fmt.Print("\033[01;31m错误:\033[0m")
	fmt.Println(a...)
}

// 打印错误信息
func StdError(format string, args ...interface{}) error {
	fmt.Print("\033[01;31m错误:\033[0m")
	fmt.Printf(format, args...)
	return fmt.Errorf(format, args...)
}

// 打印错误信息
func StdErrorf(format string, args ...interface{}) error {
	fmt.Print("\033[01;31m错误:\033[0m")
	fmt.Printf(format, args...)
	return fmt.Errorf(format, args...)
}

// 打印提示信息
func StdInfof(format string, a ...any) {
	fmt.Print("\033[01;32m提示:\033[0m")
	fmt.Printf(format, a...)
}

// 打印提示信息
func StdInfoln(a ...any) {
	fmt.Print("\033[01;32m提示:\033[0m")
	fmt.Println(a...)
}

// 打印绿色的前缀
func Greenf(Title string, format string, a ...any) {
	fmt.Print("\033[01;32m" + Title + ":\033[0m")
	fmt.Printf(format, a...)
}

// 打印绿色的前缀
func Greenln(Title string, a ...any) {
	fmt.Print("\033[01;32m" + Title + ":\033[0m")
	fmt.Println(a...)
}

// 打印红色的前缀
func Redln(Title string, a ...any) {
	fmt.Print("\033[01;31m" + Title + ":\033[0m")
	fmt.Println(a...)
}

// 打印红色的前缀
func Redf(Title string, format string, args ...interface{}) {
	fmt.Print("\033[01;31m" + Title + ":\033[0m")
	fmt.Printf(format, args...)
}
