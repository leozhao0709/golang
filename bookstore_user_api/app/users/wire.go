// +build wireinject

package users

import (
	"github.com/google/wire"
	"github.com/leozhao0709/golang/bookstore_user_api/app/db"
	"github.com/leozhao0709/golang/bookstore_user_api/ent"
)

var handlerSet = wire.NewSet(GetHandler, GetService, GetRepository, db.GetEntClient, wire.FieldsOf(new(*ent.Client), "User"))

func InjectHandler() IHandler {
	wire.Build(handlerSet)
	return &handler{}
}
