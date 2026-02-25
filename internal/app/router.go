package app

import (
	"log"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/handler"
	"github.com/Archiker-715/e-commerce-api/internal/middleware"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %s %s", r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
		})
	})
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return r.PathPrefix("/api/v1").Subrouter()
}

func (a *app) initRoutes() {
	a.authRouter(a.Handlers.auth)
	a.productRouter(a.Handlers.product)
	a.userCartRouter(a.Handlers.cart)
	a.marketRouter(a.Handlers.market)
	a.orderRouter(a.Handlers.order)
	a.paymentRouter(a.Handlers.payment)
}

func (a *app) authRouter(authHandler *handler.AuthHandler) {
	authRouter := a.Router.PathPrefix("/sso").Subrouter()
	authRouter.HandleFunc("/login", authHandler.Authorize).Methods("POST")
	authRouter.HandleFunc("/registration", authHandler.Registration).Methods("POST")
}

func (a *app) productRouter(productHandler *handler.ProductHandler) {
	productRouter := a.Router.PathPrefix("products").Subrouter()
	productRouter.HandleFunc("/products", productHandler.GetProduct).Methods("GET")
	productRouter.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	productRouter.HandleFunc("/products/{productId}", productHandler.UpdateProduct).Methods("PUT")
	productRouter.HandleFunc("/products/{productId}", productHandler.DeleteProduct).Methods("DELETE")
	productRouter.Use(middleware.AuthMiddleware)
}
func (a *app) userCartRouter(userCartHandler *handler.CartHandler) {
	// TODO: чекнуть реалиазацию query params
	userCartRouter := a.Router.PathPrefix("cart").Subrouter()
	userCartRouter.HandleFunc("/cart", userCartHandler.GetUserCart).Methods("GET")
	userCartRouter.HandleFunc("/cart", userCartHandler.AddProductToCart).Methods("POST")
	userCartRouter.HandleFunc("/cart", userCartHandler.ChangeProductCount).Methods("PUT")
	userCartRouter.HandleFunc("/cart", userCartHandler.DeleteProductFromCart).Methods("DELETE")
	userCartRouter.Use(middleware.AuthMiddleware)
}
func (a *app) marketRouter(marketHandler *handler.MarketHandler) {
	marketRouter := a.Router.PathPrefix("markets").Subrouter()
	marketRouter.HandleFunc("/create-market", marketHandler.AddMarket).Methods("POST")
	marketRouter.HandleFunc("/link-user-market", marketHandler.LinkUserMarket).Methods("POST")
	marketRouter.Use(middleware.AuthMiddleware)
}
func (a *app) orderRouter(orderHandler *handler.OrderHandler) {
	orderRouter := a.Router.PathPrefix("orders").Subrouter()
	orderRouter.HandleFunc("/create-order", orderHandler.CreateOrder).Methods("POST")
	orderRouter.Use(middleware.AuthMiddleware)
}
func (a *app) paymentRouter(paymentHandler *handler.PaymentHandler) {
	paymentRouter := a.Router.PathPrefix("payments").Subrouter()
	paymentRouter.HandleFunc("/create-payment", paymentHandler.DoPayment).Methods("POST")
	paymentRouter.Use(middleware.AuthMiddleware)
}
