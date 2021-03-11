package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler/response"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/auth"
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
		response.Error(w, http.StatusMethodNotAllowed, nil, "bad request method")
		return
	}

	xToken := r.Header.Get("x-token")
	if xToken == "" {
		response.Error(w, http.StatusUnauthorized, nil, "x-token is empty")
		return
	}

	if err := auth.Validate(xToken); err != nil {
		response.Error(w, http.StatusUnauthorized, err, "x-token is invalid")
		return
	}

	userID, err := auth.Get(xToken, "user_id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, err, "your token don't have user_id")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err, "body couldn't read")
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		response.Error(w, http.StatusBadRequest, nil, "body is empty")
		return
	}

	var jsonBody map[string]int
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		response.Error(w, http.StatusBadRequest, err, "body couldn't convert to json")
		return
	}

	times, ok := jsonBody["times"]
	if !ok || times == 0 {
		response.Error(w, http.StatusBadRequest, nil, "times is empty or not exist")
		return
	}

	results, err := g.gachaUseCase.Draw(times)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err, "internal server error")
		return
	}

	if err := g.gachaUseCase.Store(userID, results); err != nil {
		response.Error(w, http.StatusInternalServerError, err, "internal server error")
		return
	}

	type responseList struct {
		Results []*usecase.Result `json:"results"`
	}

	res := new(responseList)
	res.Results = results
	response.WriteJSON(w, res)
}
