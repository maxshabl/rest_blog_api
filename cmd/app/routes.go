package main

import (
	"blog/internal/http/v1/action"
	"blog/internal/middelware"
	"blog/internal/router"
	"net/http"
)

var routes = []router.Route{
	{
		Path:        "/hello",
		Action:      http.HandlerFunc(action.GetUserInfo),
		Method:      http.MethodGet,
		Middlewares: []middelware.Middleware{middelware.MiddlewareOne, middelware.MiddlewareTwo},
	},
	{
		Path:        "/login",
		Action:      http.HandlerFunc(action.GetTokenByPassword),
		Method:      http.MethodPost,
		Middlewares: []middelware.Middleware{middelware.MiddlewareOne, middelware.MiddlewareTwo},
	},

	{
		Path:        "/register",
		Action:      http.HandlerFunc(action.RegisterUser),
		Method:      http.MethodPost,
		Middlewares: []middelware.Middleware{middelware.MiddlewareOne, middelware.MiddlewareTwo},
	},
}
