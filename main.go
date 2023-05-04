package main

import (
	"github.com/tealeg/xlsx/v3"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"strings"
)

func main() {
	m := make(map[string]bool)
	file, _ := os.Open("acode.xlsx")
	all, err := io.ReadAll(file)
	if err != nil {
		return
	}
	binary, err := xlsx.OpenBinary(all)
	if err != nil {
		return
	}
	sheet := binary.Sheets[0]
	rowNum := sheet.MaxRow //行
	//简单校验一下excel里内容是否是测试用例模板
	checkRow, _ := sheet.Row(0)
	if checkRow.GetCell(0).Value != "中文名" {
		slog.Info("not equal")
	} else {
		slog.Info("equal")
	}
	openFile, err := os.OpenFile("exam.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	defer openFile.Close()

	for i := 1; i < rowNum; i++ {
		row, _ := sheet.Row(i)
		name := row.GetCell(0).Value
		code := row.GetCell(1).Value
		slog.Info("获取当前行", slog.String("name", name), slog.String("code", code))
		s := strings.Join([]string{"\"", name, "\":\"", code, "\",\n"}, "")
		if _, ok := m[name]; !ok {
			openFile.WriteString(s)
			m[name] = true
		}

	}
}
