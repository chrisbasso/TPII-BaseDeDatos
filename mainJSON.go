package main

import (
	"./menuJSON"
	"bufio"
	"encoding/json"
	"fmt"
	bolt "github.com/coreos/bbolt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

type Clientes struct {
	Nrocliente int
	Nombre     string
	Apellido   string
	Domicilio  string
	Telefono   string
}

type Tarjetas struct {
	Nrotarjeta   int
	Nrocliente   int
	Validadesde  string
	Validahasta  string
	Codseguridad string
	Limitecompra int
	Estado       string
}

type Comercios struct {
	Nrocomercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

type Compras struct {
	Nrooperacion int
	Nrotarjeta   int
	Nrocomercio  int
	Fecha        string
	Monto        int
	Pagado       bool
}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transaccion de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
	err = b.Put(key, val)
	if err != nil {
		return err
	}
	// cierra transacción
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
	var buf []byte
	// abre una transacción de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})
	return buf, err
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
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Plataforma no soportada, no se puede limpiar la pantalla:(")
	}
}

func main() {

	db, err := bolt.Open("tarjetas.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	CallClear()

	menu.MostrarMenu()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "5" {
			os.Exit(0)
		    }
		if input == "1" {

			fmt.Println(" \nCLIENTES\n")
			clientes := [3]Clientes{Clientes{1, "Bruce", "Willis", "Calle Duro de Matar 2", "45678977"},
								  Clientes{2, "Silvester", "Stalone", "Calle Rocky 4", "35448977"},
								  Clientes{3, "Arnold", "Chuarseneguer", "Calle Terminator 2", "5567877"},}

			for i := range(clientes) {

			data, err := json.Marshal(clientes[i])
			if err != nil {
				log.Fatal(err)
				}
			CreateUpdate(db, "clientes", []byte(strconv.Itoa(clientes[i].Nrocliente)), data)
			resultado, err := ReadUnique(db, "clientes", []byte(strconv.Itoa(clientes[i].Nrocliente)))
			fmt.Printf("%s\n", resultado)

    			} 
			time.Sleep(6 * time.Second)
			CallClear()
			menu.MostrarMenu()
		    }
		
		if input == "2" {
			fmt.Println("\n TARJETAS\n")
			tarjetas := [3]Tarjetas{Tarjetas{4567215368912, 1, "2017-06", "2017-12", "4567", 50000, "vigente"},
								  	Tarjetas{4564578912410, 2, "2017-06", "2017-12", "4545", 50000, "vigente"},
								 	Tarjetas{4567214578789, 3, "2017-06", "2017-12", "4247", 50000, "vigente"},}

			for i := range(tarjetas) {

			data, err := json.Marshal(tarjetas[i])
			if err != nil {
				log.Fatal(err)
				}
			CreateUpdate(db, "tarjetas", []byte(strconv.Itoa(tarjetas[i].Nrotarjeta)), data)
			resultado, err := ReadUnique(db, "tarjetas", []byte(strconv.Itoa(tarjetas[i].Nrotarjeta)))
			fmt.Printf("%s\n", resultado)

			    } 

			time.Sleep(6 * time.Second)
			CallClear()
			menu.MostrarMenu()
		    }
		if input == "3" {

			fmt.Println("\nCOMERCIOS\n")

			comercios := [3]Comercios{Comercios{3040, "Dexter", "Calderon de la Barca 2548", "1612", "45898542"},
									  Comercios{4040, "Tower Records", "Av.Cabildo 2540", "1425", "45898457"},
									  Comercios{3040, "Lomitos", "Alberdi 2584", "1611", "44598548"},}

			for i := range(comercios) {

			data, err := json.Marshal(comercios[i])
			if err != nil {
				log.Fatal(err)
				}
			CreateUpdate(db, "comercios", []byte(strconv.Itoa(comercios[i].Nrocomercio)), data)
			resultado, err := ReadUnique(db, "comercios", []byte(strconv.Itoa(comercios[i].Nrocomercio)))
			fmt.Printf("%s\n", resultado)

			} 

			time.Sleep(5 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}

		if input == "4" {

			fmt.Println("\nCOMPRAS\n")

			datosTarjeta, error := ReadUnique(db, "tarjetas", []byte(strconv.Itoa(4567215368912)))
			if error != nil {
				log.Fatal(error)
				}
			var tarjeta Tarjetas
			json.Unmarshal(datosTarjeta, &tarjeta)
			
			datosComercio, error := ReadUnique(db, "comercios", []byte(strconv.Itoa(4040)))
			if error != nil {
				log.Fatal(error)
				}
			var comercio Comercios
			json.Unmarshal(datosComercio, &comercio)

			compras := [3]Compras{Compras{1, tarjeta.Nrotarjeta, comercio.Nrocomercio, "2018-07-01", 1000, false},
						  		  Compras{2, tarjeta.Nrotarjeta, comercio.Nrocomercio, "2018-07-01", 1000, false},
						  		  Compras{3, tarjeta.Nrotarjeta, comercio.Nrocomercio, "2018-07-01", 1000, false},}

			for i := range(compras) {

			data, err := json.Marshal(compras[i])
			if err != nil {
				log.Fatal(err)
				}
			CreateUpdate(db, "compras", []byte(strconv.Itoa(compras[i].Nrooperacion)), data)
			resultado, err := ReadUnique(db, "compras", []byte(strconv.Itoa(compras[i].Nrooperacion)))
			fmt.Printf("%s\n", resultado)

			} 

			time.Sleep(6 * time.Second)
			CallClear()
			menu.MostrarMenu()
		    }
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

}
