package app

import (
	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/kafka"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
)

type services struct {
	AuthService     *auth.AuthService
	ProductService  *uc.ProductService
	UserCartService *uc.UserCartService
	MarketService   *uc.MarketService
	OrderService    *uc.OrderService
	PaymentService  *uc.PaymentService
}

func newServices(r *repositories) *services {

	productService := uc.NewProductService(r.ProductRepo, r.ProductRulesRepo)
	userCartService := uc.NewUserCartService(r.UserCartRepo)
	orderService := uc.NewOrderService(
		r.OrderRepo,
		productService,
		userCartService,
	)

	orderWriter := kafka.NewKafkaOrderWriter(
		kafka.NewKafkaProducerClient(
			kafka.InitWriter().NewKafkaWriter("paid-orders", "localhost"),
		),
	)

	return &services{
		AuthService:     auth.NewAuthService(r.AuthRepo),
		ProductService:  productService,
		UserCartService: userCartService,
		MarketService:   uc.NewMarketService(r.MarketRepo),
		OrderService:    orderService,
		PaymentService:  uc.NewPaymentService(orderService, orderWriter),
	}
}
