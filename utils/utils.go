package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func JSON(w http.ResponseWriter, status int, d interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(d)
}

// GetEnvVar obtiene una variable de entorno o regresa el string especificado
func GetEnvVar(this string, orThis string) string {
	d := os.Getenv(this)
	if d == "" {
		return orThis
	}
	return d
}

// GetPortFromEnvVar valida la existencia del puerto en el enviroment y lo regresa con el formato necesario
func GetPortFromEnvVar(this, orThis string) string {
	p := GetEnvVar(this, orThis)
	if p != orThis {

		return fmt.Sprintf(":%s", p)
	}
	return orThis
}
