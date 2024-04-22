package customlibs

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

//Read data from an Excel file
func ReadDataXLSX(filepath string) {
    f, err := excelize.OpenFile(filepath)
    if err != nil {
        log.Fatal(err)
    }

    // gets the rows of the sheet
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        log.Fatal(err)
    }

    // Iterar sobre las filas y mostrar los datos
    for _, row := range rows {
        for _, cell := range row {
            fmt.Printf("%s\t", cell)
        }
        fmt.Println()
    }
}
