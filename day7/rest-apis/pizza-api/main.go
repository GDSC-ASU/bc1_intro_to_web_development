package main

import (
	"encoding/json"
	"net/http"
)

var pizzas = make(map[string]Pizza)

type Pizza struct {
	Type        string  `json:"type"`
	ExtraCheese bool    `json:"extraCheese"`
	Size        string  `json:"size"`
	Price       float64 `json:"price"`
}

func handleGetPizza(res http.ResponseWriter, req *http.Request) {
	_type := req.URL.Query().Get("type")
	pizza, ok := pizzas[_type]
	if ok {
		json.NewEncoder(res).Encode(pizza)
		return
	}

	// res.WriteHeader(404)
	res.WriteHeader(http.StatusNotFound)
}

func handleGetPizzas(res http.ResponseWriter, req *http.Request) {
	pizzasArr := make([]Pizza, 0)

	for _, pizza := range pizzas {
		pizzasArr = append(pizzasArr, pizza)
	}

	json.NewEncoder(res).Encode(pizzasArr)
}

func handleAddPizza(res http.ResponseWriter, req *http.Request) {
	var pizza Pizza
	err := json.NewDecoder(req.Body).Decode(&pizza)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	pizzas[pizza.Type] = pizza
}

func handleUpdatePizza(res http.ResponseWriter, req *http.Request) {
	_type := req.URL.Query().Get("type")

	_, exists := pizzas[_type]
	if len(_type) == 0 || !exists {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	var pizza Pizza
	err := json.NewDecoder(req.Body).Decode(&pizza)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	pizzas[_type] = pizza
}

func handleDeletePizza(res http.ResponseWriter, req *http.Request) {
	_type := req.URL.Query().Get("type")

	_, exists := pizzas[_type]
	if len(_type) == 0 || !exists {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	delete(pizzas, _type)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/pizza/get", handleGetPizza)
	router.HandleFunc("/pizza/get-all", handleGetPizzas)
	router.HandleFunc("/pizza/add", handleAddPizza)
	router.HandleFunc("/pizza/update", handleUpdatePizza)
	router.HandleFunc("/pizza/delete", handleDeletePizza)

	http.ListenAndServe(":8081", router)
}
