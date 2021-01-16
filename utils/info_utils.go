package utils

import (
"io/ioutil"
"log"
"os"

)

func GetBanner() string {
	if _, e := os.Stat("banner.txt"); os.IsNotExist(e) {
		return getStringFromFile("bannDef.txt")
	}
	return getStringFromFile("bannDef.txt")
}

func getStringFromFile(path string) string {
	data, e := ioutil.ReadFile(path)
	if e != nil {
		log.Fatal("Error al obtener el banner. Detalles: ", e)
	}

	return string(data)
}