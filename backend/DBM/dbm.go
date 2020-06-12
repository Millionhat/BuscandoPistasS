package DBM

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Basket struct {
	Info []DBrecord
}

type DBrecord struct {
	Nombre string
	Datos  Host
}

func (a DBrecord) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *DBrecord) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type EndPoint struct {
	IpAddress string `json: "ipAddress"`
	Grade     string `json: "ssl_grade"`
	Country   string `json: "country"`
	Owner     string `json: "owner"`
}

type Host struct {
	Endpoints       []EndPoint `json: "endpoints"`
	Host            string     `json: "host"`
	Servers_changed bool       `json: "servers_changed"`
	Ssl_grade       string     `json: "ssl_grade"`
	P_ssl_grade     string     `json: "previous_ssl_grade"`
	Logo            string     `json: "logo"`
	Title           string     `json: "title"`
	Is_down         bool       `json: "is_down"`
}

func (a Host) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Host) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

func (a EndPoint) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *EndPoint) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
func connectDB() *sql.DB {
	conStr := "postgresql://millionhat@localhost:8081/infobusqueda?sslmode=disable"
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		panic(err.Error)
	}
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS infositios (host STRING PRIMARY KEY, datos JSONB )"); err != nil {
		log.Fatal(err)
	}
	return db
}
func FetchInfo() []byte {
	db := connectDB()
	rows, err := db.Query("Select host,datos From infositios")
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	db.Close()
	conjunto := []DBrecord{}
	for rows.Next() {
		estructura := DBrecord{}
		err = rows.Scan(&estructura.Nombre, &estructura.Datos)
		if err != nil {
			panic(err.Error)
		}
		conjunto = append(conjunto, estructura)
	}
	body := Basket{}
	body.Info = conjunto
	agregar, _ := json.Marshal(conjunto)
	return agregar
}

func GuardarDato(t string, d []byte) {
	db := connectDB()
	verificacion := `SELECT * FROM Infositios WHERE host=$1`
	row := db.QueryRow(verificacion, t)
	err := row.Scan(&t)
	if err != sql.ErrNoRows {
		actualizar := `UPDATE infositios SET datos=$1 WHERE host=$2`
		_, err := db.Exec(actualizar, d, t)
		if err != nil {
			panic(err.Error)
		}
		fmt.Println("actualizado")
	} else {
		statement, err := db.Prepare(`INSERT INTO infositios(host,datos) VALUES ($1,$2)`)
		if err != nil {
			panic(err.Error)
		}
		defer statement.Close()
		guardo, err := statement.Exec(&t, &d)
		if err != nil {
			panic(err.Error)
		}
		afectada, err := guardo.RowsAffected()
		if err != nil {
			panic(err.Error)
		}
		fmt.Println(afectada, "agregado")
	}
	db.Close()
}
