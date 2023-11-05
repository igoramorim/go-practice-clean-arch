package restorder

import (
	"encoding/json"
	"errors"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"net/http"
	"strconv"
)

// TODO: Add unit tests.

func NewHandler(
	createOrderUseCase dorder.CreateOrderUseCase,
	findAllOrdersByPageUseCase dorder.FindAllOrdersByPageUseCase) *Handler {

	return &Handler{
		createOrderUseCase:         createOrderUseCase,
		findAllOrdersByPageUseCase: findAllOrdersByPageUseCase,
	}
}

type Handler struct {
	createOrderUseCase         dorder.CreateOrderUseCase
	findAllOrdersByPageUseCase dorder.FindAllOrdersByPageUseCase
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input dorder.CreateOrderUseCaseInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.createOrderUseCase.Execute(r.Context(), input)
	if err != nil {
		if errors.Is(err, dorder.ErrOrderAlreadyExists) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FindAllByPage(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 10
	}

	input := dorder.FindAllOrdersByPageUseCaseInput{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
	res, err := h.findAllOrdersByPageUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
