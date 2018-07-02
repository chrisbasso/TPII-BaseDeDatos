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
	"./menu"
	
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
	DB_NAME = "postgres"
)

func main() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
	log.Fatal(err)
	}
	_, err = db.Query("DROP DATABASE IF EXISTS tarjetas;")
	mostrarError(err)
	if err != nil {
		log.Fatal(err)
		}
	_, err = db.Query("CREATE DATABASE tarjetas;")
	mostrarError(err)
	if err != nil {
		log.Fatal(err)
		}

	dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, "tarjetas")
	db, err = sql.Open("postgres", dbinfo)
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
		if input == "21" {
			os.Exit(0)
		}
		if input == "1"{
			fmt.Println(" \nTABLAS CREADAS!\n")
			_, err = db.Exec(leerArchivo("tablas.sql"))
			mostrarError(err)
			time.Sleep(1 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "2"{
			fmt.Println(" \nPKs y FKs ESTABLECIDAS!\n")
			_, err = db.Exec(leerArchivo("PK y FK.sql"))
			mostrarError(err)
			time.Sleep(1 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "3"{
			fmt.Println(" \nPKs y FKs BORRADAS!\n")
			_, err = db.Exec(leerArchivo("DROP PK y FK.sql"))
			mostrarError(err)
			time.Sleep(1 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "4"{
			fmt.Println(" \nFUNCIONES CARGADAS!\n")
			_, err = db.Exec(leerArchivo("funciones.sql"))
			mostrarError(err)
			time.Sleep(1 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}

		if input == "5"{
			fmt.Println(" \nDATOS CARGADOS!\n")
			_, err = db.Exec(leerArchivo("datos.sql"))
			mostrarError(err)
			time.Sleep(1 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "6"{
			fmt.Println("\nTABLAS\n")
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
		if input == "7"{
			fmt.Println("\nDATOS CLIENTES\n")
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
		if input == "8"{
			fmt.Println("\nDATOS TARJETAS\n")
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
		
		if input == "9"{
			fmt.Println("\nDATOS COMERCIOS\n")
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
		
		if input == "10"{
			fmt.Println("\nLISTA COMPRAS\n")
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
		
		if input == "11"{
			fmt.Println("\nLISTA RECHAZOS\n")
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
			time.Sleep(6 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "12"{
			fmt.Println("\nCOMPRA AUTORIZADA\n")
			_, err := db.Query(`SELECT autorizarCompra('475913199634','2516',111.00,3050);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "13"{
			fmt.Println("\nCOMPRA CON TARJETA ANULADA\n")
			_, err := db.Query(`SELECT autorizarCompra('437501035853','8764',1810.00,4040);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "14"{
			fmt.Println("\nCOMPRA CON TARJETA SUSPENDIDA\n")
			_, err := db.Query(`SELECT autorizarCompra('488207236937','2650',230.00,3050);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "15"{
			fmt.Println("\nCOMPRA SUPERANDO LIMITE\n")
			_, err := db.Query(`SELECT autorizarCompra('485834874942','1505',60000.00,3050);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "16"{
			fmt.Println("\nDOS COMPRAS EN MISMO CP EN MENOS DE 1 MINUTO\n")
			time.Sleep(1 * time.Second)
			_, err := db.Query(`SELECT autorizarCompra('489419235332','5820',2000,9604);`)
			mostrarError(err)
			fmt.Println("PRIMERA COMPRA REALIZADA")
			time.Sleep(2 * time.Second)
			_, err2 := db.Query(`SELECT autorizarCompra('489419235332','5820',100,3050);`)
			mostrarError(err2)
			fmt.Println("SEGUNDA COMPRA REALIZADA")
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "17"{
			fmt.Println("\nDOS COMPRAS EN DISTINTO CP EN MENOS DE 5 MINUTOS\n")
			time.Sleep(1 * time.Second)
			_, err := db.Query(`SELECT autorizarCompra('436782605294','3064',100,2782);`)
			mostrarError(err)
			fmt.Println("PRIMERA COMPRA REALIZADA")
			time.Sleep(2 * time.Second)
			_, err2 := db.Query(`SELECT autorizarCompra('436782605294','3064',100,4040);`)
			mostrarError(err2)
			fmt.Println("SEGUNDA COMPRA REALIZADA")
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "18"{
			fmt.Println("\nCOMPRAS REALIZADAS\n")
			time.Sleep(1 * time.Second)
			_, err := db.Query(`SELECT realizarConsumosTest();`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "19"{
			fmt.Println("\nALERTAS\n")
			rows, err := db.Query("SELECT nroalerta,nrotarjeta,codalerta,descripcion FROM alerta;")
			mostrarError(err)
			defer rows.Close()
			for rows.Next() {
				var nroalerta string
				var nrotarjeta string
				var codalerta string
				var descripcion string
				err := rows.Scan(&nroalerta,&nrotarjeta,&codalerta,&descripcion)
				if err != nil {
					log.Fatal(err)
					}
				fmt.Println(nroalerta + " - " + nrotarjeta + " - " + codalerta + " - " + descripcion)
				}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(8 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "20"{
			fmt.Println("\nFACTURA GENERADA PARA CLIENTE 21390\n")
			_, err := db.Query(`SELECT generarFactura('21390','2018-07-01');`)
			mostrarError(err)
			time.Sleep(3 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}


	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func leerArchivo(archivo string) string {

	datos, err := ioutil.ReadFile(archivo)
	if err != nil{
		log.Fatal(err)
	}
	ret := string(datos)
	return ret

}


func mostrarError(e error) {
	if e != nil{
		log.Fatal(e)
	}
}
