package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://api.gael.cloud/general/public/sismos"

func main() {
	fmt.Println("API Request Demo")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	//fmt.Print(content)

	sismos := sismosFromJson(content)
	fmt.Println(sismos)

	for _, sismo := range sismos {
		fmt.Println(sismo.Fecha, " ", sismo.Profundidad, " ", sismo.Magnitud, " ", sismo.RefGeografica)
	}

}

func sismosFromJson(content string) []Sismo {
	sismos := make([]Sismo, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var sismo Sismo
	for decoder.More() {
		err := decoder.Decode(&sismo)
		if err != nil {
			panic(err)
		}
		sismos = append(sismos, sismo)
	}
	return sismos
}

type Sismo struct {
	Fecha, Profundidad, Magnitud, RefGeografica string
}
