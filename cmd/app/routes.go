package main

import (
	"blog/internal/http/v1/action/login"
	"blog/internal/http/v1/action/register"
	"blog/internal/middelware"
	"blog/internal/router"
	"net/http"
)

var routes = []router.Route{
	{
		Path:        "/hello",
		Action:      http.HandlerFunc(login.GetUserInfo),
		Method:      http.MethodGet,
		Middlewares: []middelware.Middleware{middelware.MiddlewareOne, middelware.MiddlewareTwo},
	},
	{
		Path:        "/login",
		Action:      http.HandlerFunc(login.GetTokenByPassword),
		Method:      http.MethodPost,
		Middlewares: []middelware.Middleware{middelware.MiddlewareOne, middelware.MiddlewareTwo},
	},

	{
		Path:        "/register",
		Action:      http.HandlerFunc(register.RegisterUser),
		Method:      http.MethodPost,
		Middlewares: []middelware.Middleware{middelware.MiddlewareOne, middelware.MiddlewareTwo},
	},
}
