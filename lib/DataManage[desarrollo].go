package data

import (
    "fmt"
    "log"

    "github.com/xuri/excelize/v2"
)


func ReadDataXLSX() {
    // Abrir el archivo de Excel
    f, err := excelize.OpenFile("ejemplo.xlsx")
    if err != nil {
        log.Fatal(err)
    }

    // Obtener las filas de la hoja de cálculo
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
