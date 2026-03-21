package flag

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

var defaultTablePrint func(head ...string) CmdTablePrint

func SetTablePrint(print func(head ...string) CmdTablePrint) {
	defaultTablePrint = print
}

func newCmdTable(head ...string) CmdTablePrint {
	if defaultTablePrint != nil {
		return defaultTablePrint(head...)
	}
	return NewMarkdwonTable(head...)
}

type CmdTablePrint interface {
	Add(data ...string)
	Print()
}
type MarkdwonTable struct {
	head []string
	rows [][]string
}

var ansiColorRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func NewMarkdwonTable(head ...string) *MarkdwonTable {
	return &MarkdwonTable{
		head: append([]string(nil), head...),
	}
}

func (m *MarkdwonTable) Add(data ...string) {
	row := make([]string, len(data))
	copy(row, data)
	m.rows = append(m.rows, row)
}

func (m *MarkdwonTable) Print() {
	if len(m.head) == 0 {
		return
	}

	colNum := len(m.head)
	widths := make([]int, colNum)

	for i, h := range m.head {
		widths[i] = visibleWidth(h)
	}

	for _, row := range m.rows {
		for i := 0; i < colNum && i < len(row); i++ {
			w := visibleWidth(row[i])
			if w > widths[i] {
				widths[i] = w
			}
		}
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	idFmt := color.New(color.FgYellow).SprintfFunc()

	printRow := func(cols []string, isHeader bool) {
		for i := 0; i < colNum; i++ {
			var cell string
			if i < len(cols) {
				cell = cols[i]
			}

			if isHeader {
				cell = headerFmt("%s", cell)
			} else if i == 0 {
				cell = idFmt("%s", cell)
			}

			fmt.Print(padRightDisplay(cell, widths[i]))
			if i != colNum-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	printRow(m.head, true)
	for _, row := range m.rows {
		printRow(row, false)
	}
}

func padRightDisplay(s string, width int) string {
	w := visibleWidth(s)
	if w >= width {
		return s
	}
	return s + strings.Repeat(" ", width-w)
}

func visibleWidth(s string) int {
	return runewidth.StringWidth(stripANSI(s))
}

func stripANSI(s string) string {
	return ansiColorRegex.ReplaceAllString(s, "")
}
