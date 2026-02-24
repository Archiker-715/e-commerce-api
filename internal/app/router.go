package app

import (
	"github.com/Archiker-715/e-commerce-api/internal/handlers"
	"github.com/Archiker-715/e-commerce-api/internal/middleware"
)

func (a *app) authRouter(authHandler *handlers.AuthHandler) {
	authRouter := a.Router.PathPrefix("/sso").Subrouter()
	authRouter.HandleFunc("/login", authHandler.Authorize).Methods("POST")
	authRouter.HandleFunc("/registration", authHandler.Registration).Methods("POST")
}

func (a *app) productRouter(productHandler *handlers.ProductHandler) {
	productRouter := a.Router.PathPrefix("products").Subrouter()
	productRouter.HandleFunc("/products", productHandler.GetProduct).Methods("GET")
	productRouter.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	productRouter.HandleFunc("/products/{productId}", productHandler.UpdateProduct).Methods("PUT")
	productRouter.HandleFunc("/products/{productId}", productHandler.DeleteProduct).Methods("DELETE")
	productRouter.Use(middleware.AuthMiddleware)
}
func (a *app) userCartRouter(userCartHandler *handlers.CartHandler) {
	// TODO: чекнуть реалиазацию query params
	userCartRouter := a.Router.PathPrefix("cart").Subrouter()
	userCartRouter.HandleFunc("/cart", userCartHandler.GetUserCart).Methods("GET")
	userCartRouter.HandleFunc("/cart", userCartHandler.AddProductToCart).Methods("POST")
	userCartRouter.HandleFunc("/cart", userCartHandler.ChangeProductCount).Methods("PUT")
	userCartRouter.HandleFunc("/cart", userCartHandler.DeleteProductFromCart).Methods("DELETE")
	userCartRouter.Use(middleware.AuthMiddleware)
}
func (a *app) marketRouter(marketHandler *handlers.MarketHandler) {
	marketRouter := a.Router.PathPrefix("markets").Subrouter()
	marketRouter.HandleFunc("/create-market", marketHandler.AddMarket).Methods("POST")
	marketRouter.HandleFunc("/link-user-market", marketHandler.LinkUserMarket).Methods("POST")
	marketRouter.Use(middleware.AuthMiddleware)
}
func (a *app) orderRouter(orderHandler *handlers.OrderHandler) {
	orderRouter := a.Router.PathPrefix("orders").Subrouter()
	orderRouter.HandleFunc("/create-order", orderHandler.CreateOrder).Methods("POST")
	orderRouter.Use(middleware.AuthMiddleware)
}
func (a *app) paymentRouter(paymentHandler *handlers.PaymentHandler) {
	paymentRouter := a.Router.PathPrefix("payments").Subrouter()
	paymentRouter.HandleFunc("/create-payment", paymentHandler.DoPayment).Methods("POST")
	paymentRouter.Use(middleware.AuthMiddleware)
}
