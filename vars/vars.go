package vars

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EstrucReg struct {
	Id     int64   `json:"id"`
	Codigo string  `json:"codigo"`
	Nombre string  `json:"nombre"`
	Stock  int64   `json:"stock"`
	Precio float64 `json:"precio"` //Modificar
}

var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

const AllowedCORSDomain = "http://localhost"

const Port = ":8000" //Modificar

const NombreRuta = "articulos" //Modificar
