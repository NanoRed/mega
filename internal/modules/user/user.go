package user

import (
	"github.com/RedAFD/mega/internal/modules/user/model"
)

func init() {
	(*model.User)(nil).CreateTable()
}
