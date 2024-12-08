/**
 * @Author: Gong Yanhui
 * @Description: 使用GO操作Excel表格
 * @Date: 2024-08-22 20:46
 */

package excel

import (
	"bufio"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

func fmtExcelIWant(filePath, sheet, writeFile string) error {

	// 打开Excel文件
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	// 根据sheet名获取到对应页
	cells, err := file.GetRows(sheet)
	if err != nil {
		return err
	}

	// 准备一个写入文件
	openFile, err := os.OpenFile(writeFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(openFile)
	for i, row := range cells {
		for j, str := range row {
			if str == "Regular Expression" {
				// 将读取到的数据写入文件
				_, errIn := writer.WriteString(cells[i][j+1])
				if errIn != nil {
					fmt.Println("[ERROR]:", cells[i][j+1], errIn)
				} else {
					writer.WriteString("\n")
				}
			}
		}
	}
	writer.Flush()

	return nil
}

func TestExcel() {
	err := fmtExcelIWant(
		"/Users/gongyanhui/Downloads/0802Policies.xlsx",
		"Exceptions",
		"/Users/gongyanhui/Downloads/domain.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func TestNewExcel() {
	f := excelize.NewFile()
	// 创建一个工作表
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		panic(err)
	}
	// 设置单元格的值
	_ = f.SetCellValue("Sheet2", "A2", "Hello world.")
	_ = f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("export.xlsx"); err != nil {
		fmt.Println(err)
	}
}
