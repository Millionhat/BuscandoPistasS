package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	D "../DBM"
	S "../search"
	"github.com/buaazp/fasthttprouter"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
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

func Historial(r *fasthttp.RequestCtx) {
	//db := connectDB()
	strContentType := []byte("Content-Type")
	strApplicationJSON := []byte("application/json")
	r.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	r.SetContentType("application/json; charset=UTF-8")
	r.Response.Header.Add("Access-Control-Allow-Origin", "*")
	/*rows, err := db.Query("Select host,datos From infositios")
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
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
	r.Response.SetBody(agregar)
	db.Close()*/
	r.Response.SetBody(D.FetchInfo())
}

func CrearDato(r *fasthttp.RequestCtx) {
	host := r.UserValue("host").(string)
	host = strings.ToLower(host)
	//db := connectDB()
	strContentType := []byte("Content-Type")
	strApplicationJSON := []byte("application/json")
	r.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	r.SetContentType("application/json; charset=UTF-8")
	r.Response.Header.Add("Access-Control-Allow-Origin", "*")
	r.SetStatusCode(http.StatusOK)
	d := S.GetInfo(host)
	D.GuardarDato(host, d)
	/*
		statement, err := db.Prepare(`INSERT INTO infositios(host,datos) VALUES ($1,$2)`)
		if err != nil {
			panic(err.Error)
		}
		defer statement.Close()
		guardo, err := statement.Exec(&host, &d)
		if err != nil {
			panic(err.Error)
		}
		afectada, err := guardo.RowsAffected()
		if err != nil {
			panic(err.Error)
		}
		fmt.Println(afectada)
	*/
	r.Response.SetBody(d)

	//db.Close()
}

func main() {
	/*d := connectDB()

	d.Close()*/
	router := fasthttprouter.New()
	router.GET("/historial", Historial)
	router.GET("/nuevo/:host", CrearDato)
	server := fasthttp.Server{
		Handler: router.Handler,
		Name:    "Doofenshmirtz",
	}
	if err := server.ListenAndServe(":8087"); err != nil {
		server.Shutdown()
		log.Fatal(server.ListenAndServe(":8087"))
		//log.Println(d)
	}

}

/*
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
*/
