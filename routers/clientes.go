package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"test/models"

	"github.com/gorilla/mux"
)

type Compras struct {
	Items []models.Cliente
}

func DataCliente() []models.Cliente {
	c := []models.Cliente{
		{1000223, "992003040", "Juan Mata", true, "gold", 123.20, "2019-12-01"},
		{1000224, "3232891714", "Flabio Hinestroza", true, "silver", 234.20, "2022-03-04"},
		{1000225, "3232891714", "Jhon", false, "gold", 112.20, "2022-03-07"},
		{1000226, "3232891714", "Yuli", false, "gold", 112.20, "2022-03-07"},
		{1000227, "3232891714", "Juan", false, "gold", 112.20, "2022-03-06"},
		{1000228, "3232891714", "Sofia", false, "gold", 112.20, "2022-03-10"},
	}
	return c
}
func GetCliente(w http.ResponseWriter, r *http.Request) {

	result := DataCliente()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		Data []models.Cliente `json:"data"`
	}{
		Data: result,
	})

}
func GetClienteCompras(w http.ResponseWriter, r *http.Request) {
	fecha := mux.Vars(r)["fecha"]
	items := []models.Cliente{}
	box := Compras{items}

	result := DataCliente()
	if len(fecha) < 1 {
		http.Error(w, "Debe enviar el parámetro fecha", http.StatusBadRequest)
		return
	}
	for i := range result {

		if result[i].Date == fecha {

			box.AddItem(result[i])
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		Data []models.Cliente `json:"data"`
	}{
		Data: box.Items,
	})

}

func GetClienteResumen(w http.ResponseWriter, r *http.Request) {
	dias := r.URL.Query().Get("dias")
	fecha := mux.Vars(r)["fecha"]
	items := []models.Cliente{}
	box := Compras{items}
	result := DataCliente()
	var total = 0
	var nocompraron = 0

	var comprasPorTDC = make(map[string]int)
	var CalcularMaxMayor = 0.0
	fechaTemp := strings.Split(fecha, "-")

	diasTemp := fechaTemp[2]
	dT, _ := strconv.ParseInt(diasTemp, 10, 8)
	dP, _ := strconv.ParseInt(dias, 10, 8)
	resulDTP := int(dT + dP)
	fechaTemp[2] = strconv.Itoa(resulDTP)
	fechaNew := fechaTemp[0] + "-" + fechaTemp[1] + "-0" + fechaTemp[2]

	if len(fecha) < 1 {
		http.Error(w, "Debe enviar el parámetro fecha", http.StatusBadRequest)
		return
	}
	if len(dias) < 1 {
		http.Error(w, "Debe enviar el parámetro dias", http.StatusBadRequest)
		return
	}
	for i := range result {

		if result[i].Date >= fecha && result[i].Date <= fechaNew {

			total += int(result[i].Monto)

			CalcularMaxMayor = result[i].Monto

			if result[i].Monto > CalcularMaxMayor {
				CalcularMaxMayor = result[i].Monto
			}

			if !result[i].Compro {
				nocompraron += int(result[i].Monto)
			}
			if result[i].Tdc == "gold" {
				comprasPorTDC["oro"] += int(result[i].Monto)
			}
			box.AddItem(result[i])
		}
	}
	fmt.Println(CalcularMaxMayor)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		Total         int            `json:"total"`
		CompraMasAlta float64        `json:"compraMasAlta"`
		Nocompraron   float64        `json:"nocompraron"`
		ComprasPorTDC map[string]int `json:"comprasPorTDC"`
	}{

		Total:         total,
		CompraMasAlta: CalcularMaxMayor,
		Nocompraron:   float64(nocompraron),
		ComprasPorTDC: comprasPorTDC,
	})

}

func (box *Compras) AddItem(item models.Cliente) []models.Cliente {
	box.Items = append(box.Items, item)

	return box.Items
}
