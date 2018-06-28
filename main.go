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
	DB_NAME = "tarjetas"
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
		if input == "20" {
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
			fmt.Println(" PKs y FKs BORRADAS!")
			_, err = db.Exec(borrarPKyFK())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "4"{
			fmt.Println("")
			fmt.Println(" FUNCIONES CARGADAS!")
			_, err = db.Exec(cargarFunciones())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}

		if input == "5"{
			fmt.Println("")
			fmt.Println(" DATOS CARGADOS!")
			_, err = db.Exec(cargarDatos())
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "6"{
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
		if input == "7"{
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
		if input == "8"{
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
		
		if input == "9"{
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
		
		if input == "10"{
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
		
		if input == "11"{
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
		if input == "12"{
			fmt.Println("")
			fmt.Println("COMPRA AUTORIZADA")
			fmt.Println("")
			_, err := db.Query(`SELECT autorizarCompra('475913199634','2516',111.00,3050);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "13"{
			fmt.Println("")
			fmt.Println("COMPRA CON TARJETA ANULADA")
			fmt.Println("")
			_, err := db.Query(`SELECT autorizarCompra('437501035853','8764',1810.00,4040);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "14"{
			fmt.Println("")
			fmt.Println("COMPRA CON TARJETA SUSPENDIDA")
			fmt.Println("")
			_, err := db.Query(`SELECT autorizarCompra('488207236937','2650',230.00,3050);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "15"{
			fmt.Println("")
			fmt.Println("COMPRA SUPERANDO LIMITE")
			fmt.Println("")
			_, err := db.Query(`SELECT autorizarCompra('485834874942','1505',60000.00,3050);`)
			mostrarError(err)
			time.Sleep(2 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "16"{
			fmt.Println("")
			fmt.Println("DOS COMPRAS EN MISMO CP EN MENOS DE 1 MINUTO")
			fmt.Println("")
			time.Sleep(2 * time.Second)
			_, err := db.Query(`SELECT autorizarCompra('489419235332','5820',2000,9604);`)
			mostrarError(err)
			fmt.Println("PRIMERA COMPRA REALIZADA")
			time.Sleep(3 * time.Second)
			_, err2 := db.Query(`SELECT autorizarCompra('489419235332','5820',100,3050);`)
			mostrarError(err2)
			fmt.Println("SEGUNDA COMPRA REALIZADA")
			time.Sleep(3 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "17"{
			fmt.Println("")
			fmt.Println("DOS COMPRAS EN DISTINTO CP EN MENOS DE 5 MINUTOS")
			fmt.Println("")
			time.Sleep(2 * time.Second)
			_, err := db.Query(`SELECT autorizarCompra('489419235332','5820',100,3050);`)
			mostrarError(err)
			fmt.Println("PRIMERA COMPRA REALIZADA")
			time.Sleep(3 * time.Second)
			_, err2 := db.Query(`SELECT autorizarCompra('489419235332','5820',100,4040);`)
			mostrarError(err2)
			fmt.Println("SEGUNDA COMPRA REALIZADA")
			time.Sleep(3 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "18"{
			fmt.Println("")
			fmt.Println("ALERTAS")
			fmt.Println("")
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
		if input == "19"{
			fmt.Println("")
			fmt.Println("FACTURA GENERADA")
			fmt.Println("")
			_, err := db.Query(`SELECT generarFactura('27944','2018-06-25');`)
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
func borrarPKyFK() string{

	datos, errorDeLectura := ioutil.ReadFile("DROP PK y FK.sql")
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

func mostrarError(e error) {
	if e != nil{
		log.Fatal(e)
	}
}
