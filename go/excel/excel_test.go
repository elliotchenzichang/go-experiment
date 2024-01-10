package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"testing"
)

func TestCreateExcel(t *testing.T) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A2", "Hello world.")
	//设置单元格样式
	style, err := f.NewStyle(`{
    "font":
    {
        "bold": true,
        "family": "font-family",
        "size": 20,
        "color": "#777777"
    }
}`)
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellStyle("Sheet1", "B1", "B1", style)

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func TestReadExcelValue(t *testing.T) {
	fd, err := os.OpenFile("Book1.xlsx", os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	f, err := excelize.OpenReader(fd)
	value, _ := f.GetCellValue("Sheet1", "A2")
	fmt.Println(value)
}

func TestStreamWriter(t *testing.T) {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	for i := 1; i < 1000000; i++ {
		row := []interface{}{1}
		cell, _ := excelize.CoordinatesToCellName(1, i)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
