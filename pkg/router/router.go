package router

import (
	"github.com/labstack/echo/v4"
	"github.com/randomtoy/go-8ball/pkg/randomizer"
	"github.com/randomtoy/go-8ball/pkg/templater"
	"github.com/randomtoy/go-8ball/pkg/validator"
)

type Router struct {
	Request  Request
	Response Response
}

type Request struct {
	Code string `json:"code"`
}

type Response struct {
	Code   string `json:"code"`
	Value  string `json:"value"`
	String string `json:"string"`
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) MainHandler(c echo.Context) error {

	err := c.Bind(&r.Request)
	if err != nil {
		return err
	}
	if !validator.IsValidCode(r.Request.Code) {
		return c.JSON(400, "invalid code")
	}
	rand := randomizer.NewRandomizer(r.Request.Code)
	random := rand.Generate()
	t := templater.NewTemplater(r.Request.Code)

	template, err := t.Generate(random.ValueInt)
	if err != nil {
		return err
	}
	r.Response.Code = template.Code
	r.Response.Value = template.Value
	r.Response.String = template.String

	return c.JSON(200, r.Response)
}
