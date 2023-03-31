package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tr_js struct {
	Tr []ttd `json:"tr"`
}

type ttd struct {
	Td    []string `json:"td"`
	Align string   `json:"@align"`
}

func main() {
	raw, _ := os.ReadFile("E:\\js.json")
	var tr Tr_js
	err := json.Unmarshal(raw, &tr)
	if err != nil {
		fmt.Println(err)
	}
	var res_js []string
	for i, td := range tr.Tr {
		if i == 0 {
			continue
		}
		res := td.Td[0] + " " + td.Td[1]
		res_js = append(res_js, res)
	}
	res_raw, _ := json.Marshal(res_js)
	os.WriteFile("E:\\resjs.json", res_raw, os.ModePerm)
	fmt.Println(tr)

}
