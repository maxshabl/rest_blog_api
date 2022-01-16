package action

import (
	"blog/internal/domain/user/repo"
	"blog/internal/domain/user/usecase/login"
	"blog/internal/errors"
	"blog/internal/http/v1/request"
	"blog/internal/http/v1/response"
	"blog/internal/store"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	params := request.Login{}
	err := decoder.Decode(&params)

	if err != nil {
		errors.AddErrorContext(err, "", "wrong input params")
		fmt.Fprintf(os.Stderr, "input format errors %s", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, jwt, err := login.Handler(repo.Init(store.GetDB()), login.Command(params))

	if err != nil {
		errors.HandleError(w, err)
		return
	}

	result, _ := json.Marshal(response.LoginResponse{
		Login:      user.Login,
		Email:      user.Email,
		JWT:        jwt,
		ResetToken: user.PasswordResetToken,
	})

	w.Write(result)
	println("Params:", user.Login, user.PasswordHash)

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
