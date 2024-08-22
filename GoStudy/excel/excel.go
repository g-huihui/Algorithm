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
