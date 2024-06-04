package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SimulacionCompraRequest struct {
	IDUsuario      int     `json:"id_usuario"`
	IDMoneda       int     `json:"id_moneda"`
	Cantidad       float64 `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`
}

func handleSimulacionCompra(w http.ResponseWriter, r *http.Request) {
	var request SimulacionCompraRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Validar los datos de entrada
	if request.IDUsuario == 0 || request.IDMoneda == 0 || request.Cantidad <= 0 || request.PrecioUnitario <= 0 {
		http.Error(w, "Todos los campos son requeridos y deben ser valores positivos", http.StatusBadRequest)
		return
	}

	// Guardar la simulación de compra en la base de datos
	_, err = db.Exec("INSERT INTO SimulacionCompra (id_usuario, id_moneda, cantidad, precio_unitario) VALUES ($1, $2, $3, $4)",
		request.IDUsuario, request.IDMoneda, request.Cantidad, request.PrecioUnitario)
	if err != nil {
		http.Error(w, "Error al guardar la simulación de compra", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Simulación de compra guardada correctamente"))
}

func main() {
	http.HandleFunc("/simulacion-compra", handleSimulacionCompra)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
