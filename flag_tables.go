package flag

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

// 定义CmdTable结构体，包含一个table.Table类型的指针t和一个int类型的maxlen
type CmdTable struct {
	t      *table.Table
	maxlen int
}

// 创建一个新的CmdTable实例，参数为head，类型为...string，即可变参数
func newCmdTable(head ...string) *CmdTable {
	// 声明一个CmdTable类型的变量tab
	var tab CmdTable
	// 声明一个table.Table类型的变量t
	t := table.Table{}
	// 声明一个table.Row类型的变量row
	var row table.Row
	// 遍历head中的每一个元素
	for _, v := range head {
		// 将Green(v)追加到row中
		row = append(row, Green(v))
	}

	// 将row作为表头追加到t中
	t.AppendHeader(row)
	// 设置t的样式，将SeparateRows设置为true
	t.Style().Options.SeparateRows = true
	// 将t的指针赋值给tab的t
	tab.t = &t
	// 返回tab的指针
	return &tab
}

// 添加新的表头
func (t *CmdTable) AppendHeader(head ...string) {
	t.t.AppendHeader(table.Row{head})
}

// formatOutput函数用于格式化输出字符串
func (t *CmdTable) formatOutput(output string) string {
	// 如果最大长度为0或者输出字符串长度小于等于最大长度，则直接返回输出字符串
	if t.maxlen == 0 || len(output) <= t.maxlen {
		return output
	}
	// 创建一个字符串构建器
	var sb strings.Builder
	// 获取输出字符串的长度
	length := len(output)

	// 循环遍历输出字符串
	for i := 0; i < length; i += t.maxlen {
		// 计算当前截取的结束位置
		end := i + t.maxlen
		// 如果结束位置大于输出字符串长度，则将结束位置设置为输出字符串长度
		if end > length {
			end = length
		}
		// 将截取的字符串写入字符串构建器
		sb.WriteString(output[i:end])
		// 添加换行符
		sb.WriteString("\n")
	}

	// 返回格式化后的字符串
	return sb.String()
}

// 添加新行 返回最后ID号
func (t *CmdTable) Add(data ...string) int {
	var row table.Row
	for _, v := range data {
		row = append(row, t.formatOutput(v))
	}
	t.t.AppendRow(row)
	return t.t.Length()
}

// 添加索引
func (t *CmdTable) EnableIndex() {
	t.t.SetAutoIndex(true)
}

// 关闭索引
func (t *CmdTable) DisableIndex() {
	t.t.SetAutoIndex(false)
}

// 设置表头
func (t *CmdTable) SetTitle(title string) {
	t.t.SetTitle("%s", Green(title))
}

// 获取最后一个索引
func (t *CmdTable) LastIndex() int {
	return t.t.Length()
}

// 插入汇总合并行
func (t *CmdTable) AppendMergeRow(data ...string) {
	t.t.AppendFooter(table.Row{data}, table.RowConfig{AutoMerge: true})
}

// 打印表格
func (t *CmdTable) Print() {
	fmt.Println(t.t.Render())
}

// 设置参数最长长度，大于该长度会进行分割
func (t *CmdTable) SetValueMaxLen(length int) {
	t.maxlen = length
}
