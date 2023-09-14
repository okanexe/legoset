package handler

import (
	"encoding/json"
	"net/http"

	"legoapi/internal/app"
	"legoapi/internal/service"

	"github.com/gorilla/mux"
)

type LegoSetHandler struct {
	LegoSetService service.Service
}

func NewLegoSetHandler(legoSetService service.Service) *LegoSetHandler {
	return &LegoSetHandler{
		LegoSetService: legoSetService,
	}
}

func (h *LegoSetHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/legosets", h.GetAllLegoSets).Methods("GET")
	router.HandleFunc("/legosets/{code}", h.GetLegoSetByCode).Methods("GET")
	router.HandleFunc("/legosets", h.CreateLegoSet).Methods("POST")
	router.HandleFunc("/legosets/{code}", h.UpdateLegoSet).Methods("PUT")
	router.HandleFunc("/legosets/{code}", h.DeleteLegoSet).Methods("DELETE")
}

func (h *LegoSetHandler) GetAllLegoSets(w http.ResponseWriter, r *http.Request) {
	legoSets, err := h.LegoSetService.GetAllLegoSets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, legoSets, http.StatusOK)
}

func (h *LegoSetHandler) GetLegoSetByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	legoSet, err := h.LegoSetService.GetLegoSetByCode(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeResponse(w, legoSet, http.StatusOK)
}

func (h *LegoSetHandler) CreateLegoSet(w http.ResponseWriter, r *http.Request) {
	var legoSet app.LegoSet

	err := json.NewDecoder(r.Body).Decode(&legoSet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.LegoSetService.CreateLegoSet(legoSet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, "LEGO set created successfully", http.StatusCreated)
}

func (h *LegoSetHandler) UpdateLegoSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var legoSet app.LegoSet

	err := json.NewDecoder(r.Body).Decode(&legoSet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.LegoSetService.UpdateLegoSet(code, legoSet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, "LEGO set updated successfully", http.StatusOK)
}

func (h *LegoSetHandler) DeleteLegoSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	err := h.LegoSetService.DeleteLegoSet(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, "LEGO set deleted successfully", http.StatusOK)
}

func writeResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
