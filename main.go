//go:generate goagen bootstrap -d github.com/yu-ichiko/goa-security-example/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/yu-ichiko/goa-security-example/app"
	"github.com/yu-ichiko/goa-security-example/controller"
	"github.com/yu-ichiko/goa-security-example/middleware/basic"
	"github.com/yu-ichiko/goa-security-example/middleware/sample"
	"github.com/yu-ichiko/goa-security-example/middleware/security"
)

func main() {
	// Create service
	service := goa.New("adder")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	seq := security.New(basic.New())
	seq.Use(sample.New())

	app.UseBasicAuthMiddleware(service, seq.Middleware())

	// Mount "operands" controller
	c := controller.NewOperandsController(service)
	app.MountOperandsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
