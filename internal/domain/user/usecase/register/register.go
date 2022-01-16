package register

import (
	"blog/internal/domain/user/repo"
	"blog/internal/errors"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type Command struct {
	Login    string `valid:"alphanum,length(3|255),required"`
	Password string `valid:"length(8|255),required"`
	Email    string `valid:"length(3|255),required"`
}

func (cmd Command) Validate() (bool, error) {
	return govalidator.ValidateStruct(cmd)

}

func Handler(repo repo.Repo, cmd Command) (int64, error) {

	if _, err := cmd.Validate(); err != nil {
		return 0, err
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), 16)

	if err != nil {
		return 0, err
	}

	if repo.IsLoginEmailExist(cmd.Login, cmd.Email) {
		return 0, errors.New("user with such params already exists")
	}
	userId, err := repo.AddUser(cmd.Login, string(passwordHash), cmd.Email, 0, 0)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
