// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package users

import (
	"github.com/google/wire"
	"github.com/leozhao0709/golang/bookstore_user_api/app/db"
	"github.com/leozhao0709/golang/bookstore_user_api/ent"
)

// Injectors from wire.go:

func InjectHandler() IHandler {
	client := db.GetEntClient()
	userClient := client.User
	iRepository := GetRepository(userClient)
	iService := GetService(iRepository)
	iHandler := GetHandler(iService)
	return iHandler
}

// wire.go:

var handlerSet = wire.NewSet(GetHandler, GetService, GetRepository, db.GetEntClient, wire.FieldsOf(new(*ent.Client), "User"))