package app

import (
	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/handlers"
	"github.com/Archiker-715/e-commerce-api/internal/kafka"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
)

var (
	productService  *uc.ProductService
	userCartService *uc.UserCartService
	orderService    *uc.OrderService
)

func (a *app) InitServices() {
	a.startAuthService()
	a.startProductService()
	a.startUserCartService()
	a.startMarketService()
	a.startOrderService()
	a.startPaymentService()
}

func (a *app) startAuthService() {
	authService := auth.NewAuthService(pg.NewAuthRepo(a.DB))
	authHandler := handlers.NewAuthHandler(authService)
	a.authRouter(authHandler)
}

func (a *app) startProductService() {
	productService := uc.NewProductService(pg.NewProductRepo(a.DB), pg.NewProductRulesRepo(a.DB))
	productHandler := handlers.NewProductHandler(productService)
	a.productRouter(productHandler)
}
func (a *app) startUserCartService() {
	userCartService := uc.NewUserCartService(pg.NewUserCartRepo(a.DB))
	userCartHandler := handlers.NewCartHandler(userCartService)
	a.userCartRouter(userCartHandler)
}
func (a *app) startMarketService() {
	marketService := uc.NewMarketService(pg.NewMarketRepo(a.DB))
	marketHandler := handlers.NewMarketHandler(marketService)
	a.marketRouter(marketHandler)
}
func (a *app) startOrderService() {
	orderService := uc.NewOrderService(pg.NewOrderRepo(a.DB), productService, userCartService)
	orderHandler := handlers.NewOrderHandler(orderService)
	a.orderRouter(orderHandler)
}
func (a *app) startPaymentService() {
	kafka.NewKafkaOrderHandler(orderService, kafka.NewKafkaConsumerClient(kafka.InitReader().NewKafkaReader("paid-orders", "localhost", "test-group"))).Start()
	orderWriter := kafka.NewKafkaOrderWriter(kafka.NewKafkaProducerClient(kafka.InitWriter().NewKafkaWriter("paid-orders", "localhost")))
	paymentService := uc.NewPaymentService(orderService, orderWriter)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	a.paymentRouter(paymentHandler)
}
