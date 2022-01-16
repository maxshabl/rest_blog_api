package login

import (
	"blog/internal/domain/user/entity"
	"blog/internal/domain/user/repo"
	"blog/internal/domain/user/service/auth"
	"blog/internal/errors"
	"strconv"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type Command struct {
	Login    string `valid:"alphanum,length(3|255),required"`
	Password string `valid:"length(8|255),required"`
}

func (cmd Command) Validate() (bool, error) {
	return govalidator.ValidateStruct(cmd)
}

func Handler(repo repo.Repo, cmd Command) (entity.User, string, error) {
	var user entity.User

	if _, err := cmd.Validate(); err != nil {
		return user, "", err
	}

	user, err := repo.FindUserByName(cmd.Login)

	if err != nil {
		return user, "", errors.New(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(cmd.Password)); err != nil {
		return entity.User{}, "", errors.NewBadRequest(err.Error())
	}

	manager, err := auth.NewAuthManager(user.PasswordHash)

	if err != nil {
		return user, "", errors.New(err.Error())
	}

	token, err := manager.NewJWT(strconv.FormatInt(user.Id, 10), 1)

	if err != nil {
		return user, "", errors.New(err.Error())
	}
	return user, token, nil
}
