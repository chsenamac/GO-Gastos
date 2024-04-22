package main

import (
	"fmt"
	
	//own libraries
	"github.com/chsenamac/GO-Gastos/customlibs"
)

func main() {
	ReadDataXLSX()
}

func ReadDataXLSX() {
	fmt.Println("Reading data from an Excel file...")
	filepath := "./data/pruebaxlsx.xlsx"
	customlibs.ReadDataXLSX(filepath)
}

