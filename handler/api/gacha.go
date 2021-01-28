package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler/auth"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/usecase"
)

type GachaHandler interface {
	Draw(w http.ResponseWriter, r *http.Request)
}

type gachaHandler struct {
	gachaUseCase usecase.GachaUseCase
}

func NewGachaHandler(ug usecase.GachaUseCase) GachaHandler {
	return &gachaHandler{
		gachaUseCase: ug,
	}
}

func (g gachaHandler) Draw(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "bad request method", http.StatusMethodNotAllowed)
		return
	}

	xToken := r.Header.Get("x-token")
	if xToken == "" {
		http.Error(w, "x-token is empty", http.StatusUnauthorized)
		return
	}

	authedUser, err := auth.ParseToken(xToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body couldn't read", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		http.Error(w, "body is empty", http.StatusBadRequest)
		return
	}

	var jsonBody map[string]int
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		http.Error(w, "body couldn't convert to json", http.StatusBadRequest)
		return
	}

	times, ok := jsonBody["times"]
	if !ok || times == 0 {
		http.Error(w, "times is empty or not exist", http.StatusBadRequest)
		return
	}

	results, err := g.gachaUseCase.Draw(times)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := g.gachaUseCase.Store(authedUser.Id, results); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	type response struct {
		Results []*usecase.Result
	}

	res := new(response)
	res.Results = results
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	return
}
