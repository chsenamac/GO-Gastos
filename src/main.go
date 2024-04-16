package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("Sistema de analisis, categorizacion y optimizacion de Gastos")

	salida := true

	for salida {
		switch opcion := showOptions(); opcion {
		case 1:
			limpiarConsola("")
			fmt.Println("se leera desde una hoja de calculo")
			ReadDataXLSX()
			salida = false // Esta es la condicion de salida del programa
			//Agregar la logica que corresponda
		case 2:
			fmt.Println("se leera desde una base de datos")
			limpiarConsola("")
			salida = false // Esta es la condicion de salida del programa
			//Agregar la logica que corresponda

		case 3:
			fmt.Println("se leera desde un OCR")
			salida = false // Esta es la condicion de salida del programa
			//Agregar la logica que corresponda

		case 4:
			fmt.Println("Saliendo del programa...")
			salida = false // Esta es la condicion de salida del programa
			//Agregar la logica que corresponda

			//return
		default:
			limpiarConsola("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}

func showOptions() int {
	fmt.Println("1. Hoja de calculo")
	fmt.Println("2. Base de datos")
	fmt.Println("3. OCR")
	fmt.Println("4. Salir")
	var option int
	fmt.Scanln(&option)
	return option
}

func limpiarConsola(mensaje string) {

	var limpiarCmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		limpiarCmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		limpiarCmd = exec.Command("clear")
	default:
		fmt.Println("No se pudo determinar el sistema operativo para limpiar la consola.")
		return
	}

	limpiarCmd.Stdout = os.Stdout
	if err := limpiarCmd.Run(); err != nil {
		fmt.Println("Error al limpiar la consola:", err)
	}

	fmt.Println(mensaje)
}

func ReadDataXLSX() {
	// Abrir el archivo de Excel
	f, err := excelize.OpenFile("DatosPrueba.xlsx")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Archivo abierto")
	}

	// Obtener las filas de la hoja de cálculo
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("sheet1 abierto")
	}

	fmt.Println("Numero de filas: ", len(rows))

	// Iterar sobre las filas y mostrar los datos
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		for _, cell := range row {
			fmt.Printf("%v\t", cell)
		}
		fmt.Println()
	}
}

func Suma() {
		
	var num1, num2 float64

	fmt.Print("Ingrese el primer número: ")
	if _, err := fmt.Scan(&num1); err != nil {
		fmt.Println("Error: Ingrese un número válido")
		return
	}

	fmt.Print("Ingrese el segundo número: ")
	if _, err := fmt.Scan(&num2); err != nil {
		fmt.Println("Error: Ingrese un número válido")
		return
	}

	resultado := num1 + num2
	fmt.Println("El resultado es:", resultado)
}
