package models

type Cliente struct {
	ClientId uint64  `json:"clientId"`
	Phone    string  `json:"phone"`
	Nombre   string  `json:"nombre"`
	Compro   bool    `json:"compro"`
	Tdc      string  `json:"tdc"`
	Monto    float64 `json:"monto"`
	Date     string  `json:"date"`
}
