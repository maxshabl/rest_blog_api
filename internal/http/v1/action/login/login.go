package login

import (
	"blog/internal/domain/user/repo"
	"blog/internal/domain/user/usecase/login"
	"blog/internal/errors"
	"blog/internal/store"
	"encoding/json"
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func Final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func GetTokenByPassword(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing GetTokenByPassword")

	decoder := json.NewDecoder(r.Body)
	params := Request{}
	err := decoder.Decode(&params)

	if err != nil {
		err := errors.AddErrorContext(err, errors.ContextField, "wrong input")
		errors.HandleError(w, err)
		return
	}

	user, jwt, err := login.Handler(repo.Init(store.GetDB()), login.Command(params))

	if err != nil {
		errors.HandleError(w, err)
		return
	}

	result, _ := json.Marshal(Response{
		Login:      user.Login,
		Email:      user.Email,
		JWT:        jwt,
		ResetToken: user.PasswordResetToken,
	})

	w.Write(result)
}

//func GetTokenByRefresh() (w http.ResponseWriter, r *http.Request) {
// принять реквест, загрузить модель
// проверить токен, получить юзера
/*
	Токен:
	  сгенерировать два токена
	  сохранить
	  вернуть
*/
//}
