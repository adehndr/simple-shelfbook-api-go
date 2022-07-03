package exception

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/shelfbook-api/helper"
	"example.com/shelfbook-api/model/web"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Status:  strconv.Itoa(http.StatusInternalServerError),
		Message: "INTERNAL SERVER ERROR",
		Data:    err,
	}
	encoder := json.NewEncoder(w)
	err2 := encoder.Encode(webResponse)
	helper.PanicIfError(err2)
}
