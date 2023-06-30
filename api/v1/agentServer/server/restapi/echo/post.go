// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostHandlerFunc turns a function with the right signature into a post handler
type PostHandlerFunc func(PostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostHandlerFunc) Handle(params PostParams) middleware.Responder {
	return fn(params)
}

// PostHandler interface for that can handle valid post params
type PostHandler interface {
	Handle(PostParams) middleware.Responder
}

// NewPost creates a new http.Handler for the post operation
func NewPost(ctx *middleware.Context, handler PostHandler) *Post {
	return &Post{Context: ctx, Handler: handler}
}

/*
	Post swagger:route POST / echo post

echo http request counts

echo http request counts
*/
type Post struct {
	Context *middleware.Context
	Handler PostHandler
}

func (o *Post) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}