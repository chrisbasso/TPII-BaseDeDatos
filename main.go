package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
	} else { //unsupported platform
		panic("Plataforma no soportada, no se puede limpiar la pantalla:(")
	}
}

func main() {

	time.Sleep(1 * time.Second)
	CallClear()
	menu()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "7" {
			os.Exit(0)
		}
		if input == "1"{
			fmt.Println("")
			fmt.Println(" TABLAS CREADAS!")
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}
		if input == "2"{
			fmt.Println("")
			fmt.Println(" PKs y FKs ESTABLECIDAS!")
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}
		if input == "3"{
			fmt.Println("")
			fmt.Println(" DATOS CARGADOS!")
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}
		if input == "4"{
			fmt.Println("")
			fmt.Println(" CONSUMO REALIZADO!")
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}
		if input == "5"{
			fmt.Println("")
			fmt.Println(" CONSUMO REALIZADO!")
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}
		if input == "6"{
			fmt.Println("")
			fmt.Println(" FACTURA GENERADA!")
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func menu(){

	fmt.Println("")
	fmt.Println("            --BASE DE DATOS TARJETAS--")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("  * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  1 - CREAR TABLAS                           *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  2 - ESTABLECER PKs y FKs                   *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  3 - CARGAR DATOS                           *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  4 - REALIZAR CONSUMO PERMITIDO             *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  5 - REALIZAR CONSUMO NO PERMITIDO          *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  6 - GENERAR FACTURA                        *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  7 - SALIR                                  *")
	fmt.Println("  *                                             *")
	fmt.Println("  * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("")
	fmt.Println("  ELIJA UNA OPCIÓN")
	fmt.Println("-----------------------------------------------------")


}