package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"io/ioutil"
)
type cliente struct {
	nombre string
	apellido string
	}

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
const (
	DB_USER = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME = "prueba4"
)

func main() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
	log.Fatal(err)
	}
	defer db.Close()

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
			_, err = db.Exec(crearTablas())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}
		if input == "2"{
			fmt.Println("")
			fmt.Println(" PKs y FKs ESTABLECIDAS!")
			_, err = db.Exec(establecerPKyFK())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}
		if input == "3"{
			fmt.Println("")
			fmt.Println(" DATOS CARGADOS!")
			_, err = db.Exec(cargarDatos())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu()
		}
		if input == "4"{
			fmt.Println("")
			fmt.Println("TABLAS")
			fmt.Println("")
			rows, err := db.Query("select table_name from information_schema.tables where table_schema = 'public' and table_type='BASE TABLE';")
			mostrarError(err)
			defer rows.Close()
			for rows.Next() {
				var nombre string
				err := rows.Scan(&nombre)
				if err != nil {
					log.Fatal(err)
					}
				fmt.Println(nombre)
				}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
			CallClear()
			menu()
		}
		if input == "5"{
			fmt.Println("")
			fmt.Println("DATOS")
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

func crearTablas() string {

	datos, errorDeLectura := ioutil.ReadFile("tablas.sql")
	mostrarError(errorDeLectura)
	ret := string(datos)
	return ret

}

func establecerPKyFK() string{

	datos, errorDeLectura := ioutil.ReadFile("PK y FK.sql")
	mostrarError(errorDeLectura)
	ret := string(datos)
	return ret

}
func cargarDatos() string{

	datos, errorDeLectura := ioutil.ReadFile("datos.sql")
	mostrarError(errorDeLectura)
	ret := string(datos)
	return ret

}


func menu(){



	fmt.Println("")
	fmt.Println("             \x1b[32;1m--BASE DE DATOS TARJETAS--\x1b[0m")
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
	fmt.Println("  *  \x1b[33;1m4 - MOSTRAR TABLAS\x1b[0m                         *")
	fmt.Println("  *                                             *")
	fmt.Println("  *  \x1b[33;1m5 - MOSTRAR DATOS\x1b[0m                          *")
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

func mostrarError(e error) {
	if e != nil{
		log.Fatal(e)
	}
}
