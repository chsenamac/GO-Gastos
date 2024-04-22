package customlibs

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)


//detects the operating system and clears the console
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

