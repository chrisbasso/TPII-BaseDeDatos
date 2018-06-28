package main

import (

	"encoding/json"
	"fmt"
	bolt "github.com/coreos/bbolt"
	"log"
	"strconv"
	"./menuJSON"
	"os"
	"os/exec"
	"runtime"
	"bufio"
	"time"
	
)

type Clientes struct {
	
	Nrocliente int
	Nombre string
	Apellido string
	Domicilio string
	Telefono string

}

type Tarjetas struct{
	
	Nrotarjeta int
	Nrocliente int
	Validadesde string
	Validahasta string
	Codseguridad string
	Limitecompra int
	Estado string
}

type Comercios struct {

	Nrocomercio int
	Nombre string
	Domicilio string
	Codigopostal string
	Telefono string

}

type Compras struct {
	
	Nrooperacion int
	Nrotarjeta int 
	Nrocomercio int
	Fecha string
	Monto int
	Pagado bool

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
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
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
		if input == "1"{
				
			
			fmt.Println("")
			fmt.Println(" CLIENTES")
			fmt.Println("")
			cliente1 := Clientes{1, "Bruce", "Willis", "Calle Duro de Matar 2","45678977"}
			cliente2 := Clientes{2, "Silvester", "Stalone", "Calle Rocky 4","35448977"}
			cliente3 := Clientes{3, "Arnold", "Chuarseneguer", "Calle Terminator 2","5567877"}
			
		    data, err := json.Marshal(cliente1)
			if err != nil {
			log.Fatal(err)
			}
			CreateUpdate(db, "clientes", []byte(strconv.Itoa(cliente1.Nrocliente)), data)
			resultado, err := ReadUnique(db, "clientes", []byte(strconv.Itoa(cliente1.Nrocliente)))
			fmt.Printf("%s\n", resultado)

			data2, err2 := json.Marshal(cliente2)
			if err2 != nil {
			log.Fatal(err2)
			}
			CreateUpdate(db, "clientes", []byte(strconv.Itoa(cliente2.Nrocliente)), data2)
			resultado2, err2 := ReadUnique(db, "clientes", []byte(strconv.Itoa(cliente2.Nrocliente)))
			fmt.Printf("%s\n", resultado2)

			data3, err3 := json.Marshal(cliente3)
			if err3 != nil {
			log.Fatal(err3)
			}
			CreateUpdate(db, "clientes", []byte(strconv.Itoa(cliente3.Nrocliente)), data3)
			resultado3, err3 := ReadUnique(db, "clientes", []byte(strconv.Itoa(cliente3.Nrocliente)))
			fmt.Printf("%s\n", resultado3)

			time.Sleep(6 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "2"{
				
			
			fmt.Println("")
			fmt.Println(" TARJETAS")
			fmt.Println("")
			tarjeta1 := Tarjetas{4567215368912, 1, "2017-06", "2017-12","4567",50000, "vigente"}
			tarjeta2 := Tarjetas{4564578912410, 2, "2017-06", "2017-12","4545",50000, "vigente"}
			tarjeta3 := Tarjetas{4567214578789, 3, "2017-06", "2017-12","4247",50000, "vigente"}

			
		    data, err := json.Marshal(tarjeta1)
			if err != nil {
			log.Fatal(err)
			}
			CreateUpdate(db, "tarjetas", []byte(strconv.Itoa(tarjeta1.Nrotarjeta)), data)
			resultado, err := ReadUnique(db, "tarjetas", []byte(strconv.Itoa(tarjeta1.Nrotarjeta)))
			fmt.Printf("%s\n", resultado)

			data2, err2 := json.Marshal(tarjeta2)
			if err2 != nil {
			log.Fatal(err2)
			}
			CreateUpdate(db, "tarjetas", []byte(strconv.Itoa(tarjeta2.Nrotarjeta)), data2)
			resultado2, err2 := ReadUnique(db, "tarjetas", []byte(strconv.Itoa(tarjeta2.Nrotarjeta)))
			fmt.Printf("%s\n", resultado2)

			data3, err3 := json.Marshal(tarjeta3)
			if err3 != nil {
			log.Fatal(err3)
			}
			CreateUpdate(db, "tarjetas", []byte(strconv.Itoa(tarjeta3.Nrotarjeta)), data3)
			resultado3, err3 := ReadUnique(db, "tarjetas", []byte(strconv.Itoa(tarjeta3.Nrotarjeta)))
			fmt.Printf("%s\n", resultado3)

			time.Sleep(6 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "3"{
				
			
			fmt.Println("")
			fmt.Println(" COMERCIOS")
			fmt.Println("")
			comercio1 := Comercios{5850, "Dexter", "Einstein 3584", "1425","45645777"}
			comercio2 := Comercios{4040, "Dexter", "Anchorena 4587", "1428","45644787"}
			comercio3 := Comercios{1245, "Dexter", "Almafuerte 8954", "1434","45645879"}

			
		    data, err := json.Marshal(comercio1)
			if err != nil {
			log.Fatal(err)
			}
			CreateUpdate(db, "comercios", []byte(strconv.Itoa(comercio1.Nrocomercio)), data)
			resultado, err := ReadUnique(db, "comercios", []byte(strconv.Itoa(comercio1.Nrocomercio)))
			fmt.Printf("%s\n", resultado)

			data2, err2 := json.Marshal(comercio2)
			if err2 != nil {
			log.Fatal(err2)
			}
			CreateUpdate(db, "comercios", []byte(strconv.Itoa(comercio2.Nrocomercio)), data2)
			resultado2, err2 := ReadUnique(db, "comercios", []byte(strconv.Itoa(comercio2.Nrocomercio)))
			fmt.Printf("%s\n", resultado2)

			data3, err3 := json.Marshal(comercio3)
			if err3 != nil {
			log.Fatal(err3)
			}
			CreateUpdate(db, "comercios", []byte(strconv.Itoa(comercio3.Nrocomercio)), data3)
			resultado3, err3 := ReadUnique(db, "comercios", []byte(strconv.Itoa(comercio3.Nrocomercio)))
			fmt.Printf("%s\n", resultado3)

			time.Sleep(6 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if input == "4"{
				
			
			fmt.Println("")
			fmt.Println(" COMPRAS")
			fmt.Println("")
			compra1 := Compras{8748, 457894587456, 3050, "2018-06-15",125,true}
			compra2 := Compras{8124, 457894587889, 4040, "2018-06-18",189,false}
			compra3 := Compras{8747, 457894587456, 3082, "2018-06-05",1124,true}
			
			
		    data, err := json.Marshal(compra1)
			if err != nil {
			log.Fatal(err)
			}
			CreateUpdate(db, "compras", []byte(strconv.Itoa(compra1.Nrooperacion)), data)
			resultado, err := ReadUnique(db, "compras", []byte(strconv.Itoa(compra1.Nrooperacion)))
			fmt.Printf("%s\n", resultado)

			data2, err2 := json.Marshal(compra2)
			if err2 != nil {
			log.Fatal(err2)
			}
			CreateUpdate(db, "compras", []byte(strconv.Itoa(compra2.Nrooperacion)), data2)
			resultado2, err2 := ReadUnique(db, "compras", []byte(strconv.Itoa(compra2.Nrooperacion)))
			fmt.Printf("%s\n", resultado2)

			data3, err3 := json.Marshal(compra3)
			if err3 != nil {
			log.Fatal(err3)
			}
			CreateUpdate(db, "compras", []byte(strconv.Itoa(compra3.Nrooperacion)), data3)
			resultado3, err3 := ReadUnique(db, "compras", []byte(strconv.Itoa(compra3.Nrooperacion)))
			fmt.Printf("%s\n", resultado3)

			time.Sleep(6 * time.Second)
			CallClear()
			menu.MostrarMenu()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

}
