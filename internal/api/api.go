package api

import (
	"encoding/json"
	"net/http"

	"github.com/carlosealves2/test-auto-version/internal/ports"
	"github.com/carlosealves2/test-auto-version/internal/utils"
)

type APIRouter struct {
	calculatorService ports.CalculatorService
}

func NewAPIRouter(calculatorService ports.CalculatorService) *APIRouter {
	return &APIRouter{
		calculatorService: calculatorService,
	}
}

func (ar *APIRouter) SumTwoNumbers(w http.ResponseWriter, r *http.Request) {
	intA, intB, err := utils.GetParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := ar.calculatorService.Add(intA, intB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
		"result": result,
	})

}
func (ar *APIRouter) SubtractTwoNumbers(w http.ResponseWriter, r *http.Request) {
	intA, intB, err := utils.GetParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := ar.calculatorService.Sub(intA, intB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
		"result": result,
	})
}
func (ar *APIRouter) MultiplyTwoNumbers(w http.ResponseWriter, r *http.Request) {
	intA, intB, err := utils.GetParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := ar.calculatorService.Multiply(intA, intB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
		"result": result,
	})
}
func (ar *APIRouter) DivideTwoNumbers(w http.ResponseWriter, r *http.Request) {
	intA, intB, err := utils.GetParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := ar.calculatorService.Divide(intA, intB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
		"result": result,
	})
}

func (ar *APIRouter) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /sum", ar.SumTwoNumbers)
	mux.HandleFunc("GET /subtract", ar.SubtractTwoNumbers)
	mux.HandleFunc("GET /multiply", ar.MultiplyTwoNumbers)
	mux.HandleFunc("GET /divide", ar.DivideTwoNumbers)
}
