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
	fmt.Println("            \x1b[32;1m--BASE DE DATOS TARJETAS--\x1b[0m")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("  * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  \x1b[33;1m1 - CREAR TABLAS\x1b[0m                           *")
	fmt.Println("  *                                             *")
	fmt.Println("  * \x1b[33;1m 2 - ESTABLECER PKs y FKs\x1b[0m                   *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  \x1b[33;1m3 - CARGAR DATOS\x1b[0m                           *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  \x1b[33;1m4 - REALIZAR CONSUMO PERMITIDO\x1b[0m             *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  \x1b[33;1m5 - REALIZAR CONSUMO NO PERMITIDO\x1b[0m          *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  \x1b[33;1m6 - GENERAR FACTURA\x1b[0m                        *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  \x1b[33;1m7 - SALIR\x1b[0m                                  *")
	fmt.Println("  *                                             *")
	fmt.Println("  * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("")
	fmt.Println("  ELIJA UNA OPCIÓN")
	fmt.Println("-----------------------------------------------------")


}