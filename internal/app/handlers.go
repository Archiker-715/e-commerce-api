package app

import "github.com/Archiker-715/e-commerce-api/internal/handler"

type handlers struct {
	auth    *handler.AuthHandler
	product *handler.ProductHandler
	cart    *handler.CartHandler
	market  *handler.MarketHandler
	order   *handler.OrderHandler
	payment *handler.PaymentHandler
}

func (a *app) initHandlers() {
	a.Handlers = &handlers{
		auth:    handler.NewAuthHandler(a.Services.AuthService),
		product: handler.NewProductHandler(a.Services.ProductService),
		cart:    handler.NewCartHandler(a.Services.UserCartService),
		market:  handler.NewMarketHandler(a.Services.MarketService),
		order:   handler.NewOrderHandler(a.Services.OrderService),
		payment: handler.NewPaymentHandler(a.Services.PaymentService),
	}
}
