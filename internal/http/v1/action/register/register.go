package register

import (
	"blog/internal/domain/user/repo"
	"blog/internal/domain/user/usecase/register"
	"blog/internal/errors"
	"blog/internal/store"
	"encoding/json"
	"net/http"
	"strconv"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := Request{}
	err := decoder.Decode(&params)
	if err != nil {
		err := errors.NewBadRequest("action RegisterUser invalid param")
		errors.HandleError(w, errors.AddErrorContext(err, "", "invalid input data"))

		return
	}
	loginCommand := register.Command(params)

	userId, err := register.Handler(repo.Init(store.GetDB()), loginCommand)

	if err != nil {
		err := errors.AddErrorContext(err, errors.ContextField, "wrong input")
		errors.HandleError(w, err)
		return
	}

	w.Write([]byte(strconv.FormatInt(userId, 10)))
}
