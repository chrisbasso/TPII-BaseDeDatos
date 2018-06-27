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
	_ "./menu"
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
const (
	DB_USER = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME = "prueba5"
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

	
	
	menu.MostrarMenu()

	

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "15" {
			os.Exit(0)
		}
		if input == "1"{
			fmt.Println("")
			fmt.Println(" TABLAS CREADAS!")
			_, err = db.Exec(crearTablas())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "2"{
			fmt.Println("")
			fmt.Println(" PKs y FKs ESTABLECIDAS!")
			_, err = db.Exec(establecerPKyFK())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "3"{
			fmt.Println("")
			fmt.Println(" FUNCIONES CARGADAS!")
			_, err = db.Exec(cargarFunciones())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}

		if input == "4"{
			fmt.Println("")
			fmt.Println(" DATOS CARGADOS!")
			_, err = db.Exec(cargarDatos())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "5"{
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
			menu.MostrarMenu()
		}
		if input == "6"{
			fmt.Println("")
			fmt.Println("DATOS CLIENTES")
			fmt.Println("")
			rows, err := db.Query("SELECT nombre,apellido,domicilio,telefono FROM cliente;")
			mostrarError(err)
			defer rows.Close()
			for rows.Next() {
				var nombre string
				var apellido string
				var domicilio string
				var telefono string
				err := rows.Scan(&nombre,&apellido,&domicilio,&telefono)
				if err != nil {
					log.Fatal(err)
					}
				fmt.Println(nombre + " - " + apellido + " - " + domicilio + " - " + telefono)
				}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "7"{
			fmt.Println("")
			fmt.Println("DATOS TARJETAS")
			fmt.Println("")
			rows, err := db.Query("SELECT nrotarjeta,validadesde,validahasta,codseguridad,limitecompra,estado FROM tarjeta;")
			mostrarError(err)
			defer rows.Close()
			for rows.Next() {
				var nrotarjeta string
				var validadesde string
				var validahasta string
				var codseguridad string
				var limitecompra string
				var estado string
				err := rows.Scan(&nrotarjeta,&validadesde,&validahasta,&codseguridad,&limitecompra,&estado)
				if err != nil {
					log.Fatal(err)
					}
				fmt.Println(nrotarjeta + " - " + validadesde + " - " + validahasta + " - " + codseguridad + " - " + limitecompra + " - " + estado)
				}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		
		if input == "8"{
			fmt.Println("")
			fmt.Println("DATOS COMERCIOS")
			fmt.Println("")
			rows, err := db.Query("SELECT nombre,domicilio,codigopostal,telefono FROM comercio;")
			mostrarError(err)
			defer rows.Close()
			for rows.Next() {
				var nombre string
				var domicilio string
				var codigopostal string
				var telefono string
				err := rows.Scan(&nombre,&domicilio,&codigopostal,&telefono)
				if err != nil {
					log.Fatal(err)
					}
				fmt.Println(nombre + " - " + domicilio + " - " + codigopostal + " - " + telefono)
				}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		
		if input == "9"{
			fmt.Println("")
			fmt.Println("LISTA COMPRAS")
			fmt.Println("")
			rows, err := db.Query("SELECT nrooperacion,nrotarjeta,nrocomercio,fecha,monto FROM compra;")
			mostrarError(err)
			defer rows.Close()
			for rows.Next() {
				var nrooperacion string
				var nrocomercio string
				var nrotarjeta string
				var fecha string
				var monto string
				err := rows.Scan(&nrooperacion,&nrotarjeta,&nrocomercio,&fecha,&monto)
				if err != nil {
					log.Fatal(err)
					}
				fmt.Println(nrooperacion + " - " + nrotarjeta + " - " + nrocomercio + " - " + fecha + " - " + monto)
				}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		
		if input == "10"{
			fmt.Println("")
			fmt.Println("LISTA RECHAZOS")
			fmt.Println("")
			rows, err := db.Query("SELECT * FROM rechazo;")
			mostrarError(err)
			defer rows.Close()
			for rows.Next() {
				var nrorechazo string
				var nrotarjeta string
				var nrocomercio string
				var fecha string
				var monto string
				var motivo string
				err := rows.Scan(&nrorechazo,&nrotarjeta,&nrocomercio,&fecha,&monto,&motivo)
				if err != nil {
					log.Fatal(err)
					}
				fmt.Println(nrorechazo + " - " + nrotarjeta + " - " + nrocomercio + " - " + fecha + " - " + monto + " - " + motivo)
				}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "11"{
			fmt.Println("")
			fmt.Println("COMPRA AUTORIZADA")
			fmt.Println("")
			
			_, err := db.Query(`SELECT autorizarCompra('475913199634','2516',111.00,3050);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
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
func cargarFunciones() string{

	datos, errorDeLectura := ioutil.ReadFile("funciones.sql")
	mostrarError(errorDeLectura)
	ret := string(datos)
	return ret

}

/*
func menu(){



	fmt.Println("")
	fmt.Println("             \x1b[32;1m  --BASE DE DATOS TARJETAS--\x1b[0m")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("  * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("  *                                                   *")
	fmt.Println("  *  \x1b[33;1m01 - CREAR TABLAS\x1b[0m                                *")
	fmt.Println("  *  \x1b[33;1m02 - ESTABLECER PKs y FKs\x1b[0m                        *")
	fmt.Println("  *  \x1b[33;1m03 - CARGAR FUNCIONES\x1b[0m                            *")
	fmt.Println("  *  \x1b[33;1m04 - CARGAR DATOS\x1b[0m                                *")
	fmt.Println("  *  \x1b[33;1m05 - MOSTRAR TABLAS\x1b[0m                              *")
	fmt.Println("  *  \x1b[33;1m06 - MOSTRAR CLIENTES\x1b[0m                            *")
	fmt.Println("  *  \x1b[33;1m07 - MOSTRAR TARJETAS\x1b[0m                            *")
	fmt.Println("  *  \x1b[33;1m08 - MOSTRAR COMERCIOS\x1b[0m                           *")
	fmt.Println("  *  \x1b[33;1m09 - MOSTRAR COMPRAS\x1b[0m                             *")
	fmt.Println("  *  \x1b[33;1m10 - MOSTRAR RECHAZOS\x1b[0m                            *")
	fmt.Println("  *  \x1b[33;1m11 - REALIZAR COMPRA AUTORIZADA\x1b[0m                  *")
	fmt.Println("  *  \x1b[33;1m12 - REALIZAR COMPRA CON TARJETA ANULADA\x1b[0m         *")
	fmt.Println("  *  \x1b[33;1m13 - REALIZAR COMPRA CON TARJETA SUSPENDIDA\x1b[0m      *")
	fmt.Println("  *  \x1b[33;1m14 - REALIZAR COMPRA SUPERANDO EL LIMITE\x1b[0m         *")
	fmt.Println("  *  \x1b[33;1m15 - SALIR\x1b[0m                                       *")
	fmt.Println("  *                                                   *")
	fmt.Println("  * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("")
	fmt.Println("   ELIJA UNA OPCIÓN")
	fmt.Println("----------------------------------------------------------------------------------------")


}
*/

func mostrarError(e error) {
	if e != nil{
		log.Fatal(e)
	}
}
