package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	carnet := "201903022"
	mensaje := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Sistemas Operativos 1 - 2025</title>
		</head>
		<body>
			<h1>Sistemas Operativos 1 - 2025</h1>
			<p>Número de Carnet: %s</p>
		</body>
		</html>
	`, carnet)

	fmt.Fprint(w, mensaje)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
