package flag

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type CmdTable struct {
	t      *table.Table
	maxlen int
}

func newCmdTable(head ...string) *CmdTable {
	var tab CmdTable
	t := table.Table{}
	var row table.Row
	for _, v := range head {
		row = append(row, Green(v))
	}

	t.AppendHeader(row)
	t.Style().Options.SeparateRows = true
	tab.t = &t
	return &tab
}

// 添加新的表头
func (t *CmdTable) AppendHeader(head ...string) {
	t.t.AppendHeader(table.Row{head})
}

func (t *CmdTable) formatOutput(output string) string {
	if t.maxlen == 0 || len(output) <= t.maxlen {
		return output
	}
	var sb strings.Builder
	length := len(output)

	for i := 0; i < length; i += t.maxlen {
		end := i + t.maxlen
		if end > length {
			end = length
		}
		sb.WriteString(output[i:end])
		sb.WriteString("\n")
	}

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
	t.t.SetTitle(Green(title))
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
