package main

import (
	"log"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/handlers"
	"github.com/Archiker-715/e-commerce-api/internal/kafka"
	"github.com/Archiker-715/e-commerce-api/internal/middleware"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	pg.Connect()

	authService := auth.NewAuthService(pg.NewAuthRepo(pg.DB))
	authHandler := handlers.NewAuthHandler(*authService)

	productService := uc.NewProductService(pg.NewProductRepo(pg.DB), pg.NewProductRulesRepo(pg.DB))
	productHandler := handlers.NewProductHandler(*productService)

	userCartService := uc.NewUserCartService(pg.NewUserCartRepo(pg.DB))
	userCartHandler := handlers.NewCartHandler(userCartService)

	marketService := uc.NewMarketService(pg.NewMarketRepo(pg.DB))
	marketHandler := handlers.NewMarketHandler(marketService)

	orderService := uc.NewOrderService(pg.NewOrderRepo(pg.DB), productService, userCartService)

	kafka.NewKafkaOrderHandler(orderService, kafka.NewKafkaConsumerClient(kafka.InitReader().NewKafkaReader("paid-orders", "localhost", "test-group"))).Start()
	orderHandler := handlers.NewOrderHandler(orderService)

	paymentService := uc.NewPaymentService(orderService, kafka.NewKafkaOrderWriter(kafka.NewKafkaProducerClient(kafka.InitWriter().NewKafkaWriter("paid-orders", "localhost"))))
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	r := mux.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %s %s", r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
		})
	})
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/api/v1").Subrouter()
	authRouter := api.PathPrefix("/sso").Subrouter()
	authRouter.HandleFunc("/login", authHandler.Authorize).Methods("POST")
	authRouter.HandleFunc("/registration", authHandler.Registration).Methods("POST")

	productRouter := api.PathPrefix("products").Subrouter()
	productRouter.HandleFunc("/products", productHandler.GetProduct).Methods("GET")
	productRouter.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	productRouter.HandleFunc("/products/{productId}", productHandler.UpdateProduct).Methods("PUT")
	productRouter.HandleFunc("/products/{productId}", productHandler.DeleteProduct).Methods("DELETE")
	productRouter.Use(middleware.AuthMiddleware)

	// TODO: чекнуть реалиазацию query params
	userCartRouter := api.PathPrefix("cart").Subrouter()
	userCartRouter.HandleFunc("/cart", userCartHandler.GetUserCart).Methods("GET")
	userCartRouter.HandleFunc("/cart", userCartHandler.AddProductToCart).Methods("POST")
	userCartRouter.HandleFunc("/cart", userCartHandler.ChangeProductCount).Methods("PUT")
	userCartRouter.HandleFunc("/cart", userCartHandler.DeleteProductFromCart).Methods("DELETE")
	userCartRouter.Use(middleware.AuthMiddleware)

	marketRouter := api.PathPrefix("markets").Subrouter()
	marketRouter.HandleFunc("/create-market", marketHandler.AddMarket).Methods("POST")
	marketRouter.HandleFunc("/link-user-market", marketHandler.LinkUserMarket).Methods("POST")
	marketRouter.Use(middleware.AuthMiddleware)

	orderRouter := api.PathPrefix("orders").Subrouter()
	orderRouter.HandleFunc("/create-order", orderHandler.CreateOrder).Methods("POST")
	orderRouter.Use(middleware.AuthMiddleware)

	paymentRouter := api.PathPrefix("payments").Subrouter()
	paymentRouter.HandleFunc("/create-payment", paymentHandler.DoPayment).Methods("POST")
	paymentRouter.Use(middleware.AuthMiddleware)
}
