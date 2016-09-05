package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/tompscanlan/labreserved"
	"github.com/tompscanlan/labreserved/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.LabreservedAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LabreservedAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetItemsHandler = operations.GetItemsHandlerFunc(func(params operations.GetItemsParams) middleware.Responder {
		items := operations.NewGetItemsOK()
		items.SetPayload(labreserved.AllItems)
		return items
	})
	api.PostItemHandler = operations.PostItemHandlerFunc(func(params operations.PostItemParams) middleware.Responder {
		labreserved.AllItems[*params.Additem.Name] = *params.Additem
		item := operations.NewPostItemOK()
		i, ok := labreserved.AllItems[*params.Additem.Name]
		if ok {
			item.SetPayload(&i)
			return item
		} else {
			err := operations.NewPostItemBadRequest()
			err.SetPayload("failed to add and/or find item to lab map")
			return err
		}

	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
