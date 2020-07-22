package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//hola

type EndPoint struct {
	IpAddress string `json: "ipAddress"`
	Grade     string `json: "ssl_grade"`
	Country   string `json: "country"`
	Owner     string `json: "owner"`
}

type EndPoints1 struct {
	endPoints []EndPoint
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

func GetInfo(h string) []byte {
	pasado := getDomainInfo(h + "&fromCache=on&maxAge=6")
	reciente := getDomainInfo(h)
	if len(pasado.Endpoints) != len(reciente.Endpoints) {
		reciente.Servers_changed = true
	} else {
		for i := 0; i < len(reciente.Endpoints); i++ {
			if reciente.Endpoints[i].IpAddress != pasado.Endpoints[i].IpAddress {
				reciente.Servers_changed = true
				break
			}
		}
	}
	reciente.P_ssl_grade = pasado.Ssl_grade
	getWebpageInfo(reciente)
	respuesta, _ := json.Marshal(reciente)
	return respuesta
}

func getDomainInfo(h string) *Host {

	url := "https://api.ssllabs.com/api/v3/analyze?host=" + h

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error)
	}

	defer res.Body.Close()

	byteValue, _ := ioutil.ReadAll(res.Body)
	var data interface{}
	json.Unmarshal(byteValue, &data)
	grado := ""
	//fmt.Println(data)
	info := data.(map[string]interface{})
	if info != nil {
		endpoints := info["endpoints"]
		agregar := EndPoint{}
		ep := []EndPoint{}
		j := 0
		size := 0
		if reflect.TypeOf(endpoints) != nil {
			size = int(reflect.TypeOf(endpoints).Size())
			size = (size - 2) / 10
		}
		for l := 0; l < size; l++ {
			for k, i := range endpoints.([]interface{})[l].(map[string]interface{}) {

				c := k
				if c == "grade" {
					agregar.Grade = i.(string)
					compareGrade(&agregar, &grado)
				}
				if c == "ipAddress" {
					agregar.IpAddress = i.(string)
					getServerInfo(&agregar)
				}
				j++
				if j == 10 {
					ep = append(ep, agregar)
					j = 0
					agregar = EndPoint{}
				}

			}
		}
		return &Host{
			Endpoints: ep,
			Host:      info["host"].(string),
			Ssl_grade: grado,
		}

	}
	return nil
}

func getServerInfo(str *EndPoint) {
	url := "https://who.is/whois-ip/ip-address/" + str.IpAddress
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err.Error)
	}

	// Save each .post-title as a list
	doc.Find(".queryResponseBodyKey").Each(func(i int, s *goquery.Selection) {
		info := strings.Split(s.Text(), "\n")

		for _, element := range info {
			element = strings.Join(strings.Fields(element), "")

			if strings.Contains(element, "OrgName:") {
				dato := strings.Split(element, ":")
				str.Owner = dato[1]
			} else if strings.Contains(element, "Country") {
				dato := strings.Split(element, ":")
				str.Country = dato[1]
			}
		}
	})
}

func compareGrade(s1 *EndPoint, s2 *string) {
	if *s2 == "" {
		*s2 = s1.Grade
	} else {
		switch strings.Compare(s1.Grade, *s2) {
		case 1:
			*s2 = s1.Grade
			break
		default:
			break
		}
	}
}
func getWebpageInfo(r *Host) {
	url := "https://" + r.Host
	getTitle(url, r)
	getLogo(url, r)
	r.Is_down = PingServer(r.Host)

}
func getTitle(url string, r *Host) {

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error)
	}
	defer res.Body.Close()

	dataBytes, err := ioutil.ReadAll(res.Body)
	contenido := string(dataBytes)

	indiceIniTitulo := strings.Index(contenido, "<title>")
	if indiceIniTitulo == -1 {
		fmt.Println("no entro")
	}
	indiceIniTitulo += 7

	indiceFInTitulo := strings.Index(contenido, "</title>")
	if indiceFInTitulo == -1 {
		fmt.Println("falta cerrar")
	}

	r.Title = string([]byte(contenido[indiceIniTitulo:indiceFInTitulo]))
}

func getLogo(url string, r *Host) {

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {

		panic(err.Error)
	}
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		linea := s
		la, _ := linea.Attr("href")
		if strings.Contains(la, ".png") && r.Logo == "" {
			r.Logo = url + la
		}
	})
	if r.Logo == "" {
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			linea := s
			la, _ := linea.Attr("src")
			if strings.Contains(la, ".png") && r.Logo == "" {
				r.Logo = la
			}
		})
	}
}
