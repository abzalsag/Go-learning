package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Car struct {
	Id    int    `json:"id"`
	Model string `json:"model"`
	Brand string `json:"brand"`
	Color string `json:"color"`
}

var (
	cars []Car
	mu   sync.Mutex
)

func createCar(w http.ResponseWriter, r *http.Request) {
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	cars = append(cars, car)
	mu.Unlock()
	w.Header().Set("Content-Type", "appilication/json")
	json.NewEncoder(w).Encode(car)
}

func getCar(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if len(cars) == 0 {
		http.Error(w, "no cars found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "appilication/json")
	json.NewEncoder(w).Encode(cars)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updated Car
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Failed to parse body", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i, car := range cars {
		if car.Id == id {
			updated = cars[i]
			cars[i] = updated
			w.Header().Set("Content-Type", "appilication/json")
			json.NewEncoder(w).Encode(updated)
			return

		}
	}
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i, car := range cars {
		if car.Id == id {
			cars = append(cars[:i], cars[i+1:]...)
			fmt.Fprintln(w, "deleted")
			return
		}

	}
	http.Error(w, "car not found", http.StatusNotFound)

}
func carHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCar(w, r)
	case http.MethodPost:
		createCar(w, r)
	case http.MethodPut:
		updateCar(w, r)
	case http.MethodDelete:
		deleteCar(w, r)
	default:
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/", carHandler)
	fmt.Println("start server")
	http.ListenAndServe(":8080", nil)

}
