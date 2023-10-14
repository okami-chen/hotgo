package table

import (
	"hotgo/addons/bee/dao"
	"hotgo/addons/bee/model/entity"
)

type UserSocialTable struct {
	entity.UserSocial
}

func (t *UserSocialTable) Table() string {
	return dao.UserSocial.Table()
}
