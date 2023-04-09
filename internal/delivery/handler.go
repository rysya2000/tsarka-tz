package delivery

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest/internal/model"
	"rest/internal/usecase"

	"github.com/go-chi/chi"
)

type Handler struct {
	usecase usecase.Rester
}

func New(usecase usecase.Rester) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) SubStrHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res, err := h.usecase.FindSubStr(context.TODO(), string(body))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	data := model.SubStr{
		Substring: res,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error sending response body", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) EmailHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res, err := h.usecase.CheckEmail(context.TODO(), string(body))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	data := model.Email{
		Emails: res,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error sending response body", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) SelfHandler(w http.ResponseWriter, r *http.Request) {
	str := chi.URLParam(r, "str")

	res, err := h.usecase.FindSelf(context.TODO(), str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := model.Self{
		Self: res,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error sending response body", http.StatusInternalServerError)
		return
	}

}
