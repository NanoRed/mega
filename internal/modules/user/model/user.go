package model

import (
	"context"
	"math/rand"
	"time"

	"github.com/RedAFD/mega/internal/config"
	"github.com/RedAFD/mega/internal/storage/postgres"
	"github.com/RedAFD/mega/internal/storage/redis"
	"github.com/RedAFD/mega/internal/utils/i18n"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	jsoniter "github.com/json-iterator/go"
)

type Gender uint8

func (g Gender) String() (str string) {
	switch g {
	case GenderUnknown:
		str = i18n.Sprintf("未知")
	case GenderMale:
		str = i18n.Sprintf("男")
	case GenderFemale:
		str = i18n.Sprintf("女")
	}
	return
}

const (
	GenderUnknown Gender = iota
	GenderMale
	GenderFemale
)

type User struct {
	ID       uint64    `pg:",pk,notnull"`
	Email    string    `pg:"type:varchar(255),unique,notnull"`
	Password string    `pg:"type:char(60),notnull"`
	CreateAt time.Time `pg:"default:now(),notnull"`
}

func (u *User) CreateTable() error {
	return postgres.DB().
		Model(u).
		CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
}

func (u *User) FindUserByEmail() (ok bool, err error) {
	ctx := context.Background()
	key := "user_" + u.Email
	var res string
	res, err = redis.DB().Get(ctx, key).Result()
	if err != nil {
		err = postgres.DB().Model(u).
			Where("email = ?email").
			Select()
		if err != nil && err != pg.ErrNoRows {
			return
		} else {
			ok = true
		}
		var b []byte
		b, err = jsoniter.Marshal(u)
		if err == nil {
			rand.Seed(time.Now().UnixNano())
			salt := time.Second * time.Duration(rand.Intn(config.RedisExpirationSecondSaltRange))
			redis.DB().SetEX(ctx, key, b, config.RedisDefaultExpiration+salt)
		}
	} else {
		err = jsoniter.Unmarshal([]byte(res), u)
		if err == nil && len(u.Email) > 0 {
			ok = true
		}
	}
	return
}
