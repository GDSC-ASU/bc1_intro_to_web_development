package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func calculateBMI(height, wight float64) (string, float64) {
	bmi := wight / (height * height)
	var index string

	if bmi < 18.5 {
		index = "Underweight"
	} else if bmi >= 18.5 && bmi <= 24.9 {
		index = "Normal Weight"
	} else if bmi >= 25 && bmi <= 29.9 {
		index = "Overweight"
	} else {
		index = "Obesity"
	}

	return index, bmi
}

// handleCalculateBMI calculates Body Mass Index, and returns a response
// with BMI, and status based on that index.
func handleCalculateBMI(res http.ResponseWriter, req *http.Request) {
	height, _ := strconv.ParseFloat(req.URL.Query().Get("height"), 64)
	weight, _ := strconv.ParseFloat(req.URL.Query().Get("weight"), 64)

	index, bmi := calculateBMI(height, weight)

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(map[string]any{
		"index": index,
		"bmi":   bmi,
	})
}

// handleListBMI returns a response with a map with BMI categories
func handleListBMI(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(map[string]string{
		"Underweight":   "< 18.5",
		"Normal weight": "18.5 - 24.9",
		"Overweight":    "25 - 29.9",
		"Obesity":       "> 30",
	})
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/bmi/calc", handleCalculateBMI)
	router.HandleFunc("/bmi/all", handleListBMI)

	http.ListenAndServe(":8080", router)
}
