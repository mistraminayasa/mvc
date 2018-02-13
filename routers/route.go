package routers

import (
	c "mvc/services"

	"github.com/kataras/iris"
)

// GetEngine ==> create a route
func GetEngine() *iris.Application {
	app := iris.New()

	app.Get("/customers", c.GetCustomers)
	app.Get("/customer/:id", c.GetCustomer)
	app.Post("/customers", c.CreateCustomer)
	app.Put("/customer/:id", c.UpdateCustomer)
	app.Delete("/customer/:id", c.DeleteCustomer)

	return app
}
