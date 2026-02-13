package handlers

import (
	"fmt"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/errs"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
)

type OrderHandler struct {
	order *uc.OrderService
}

func NewOrderHandler(service *uc.OrderService) *OrderHandler {
	return &OrderHandler{order: service}
}

func (o *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder []entity.ProductsToOrder
	if err := httpsrv.JsonDecode(w, r, &newOrder, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("failed to parse input: %v", err))
		return
	}

}

// пользак заходит в корзину (GET user-cart) DONE
// выбирает только то что хочет оплатить либо всё - неважно DONE
// далее идёт логика оплаты
// формирование temp_order с id userId+timestamp на выходе (POST TempOrder) DONE
// горутина которая будет удалять заказ при неоплате через 15 мин, или чекать если оплачено (если оплачен двигаем temp в false)
// оплата по id (POST payment) (поход во внешнюю систему) (пока заглушка)
// успешная оплата -  двигаем temp в false
// DELETE from user-cart то что заказано (ordered true - скрываются из выборки GET user-cart. Заказ оплачен - удаляем WHERE prIds = ? AND order_id = ?. Не оплаечн - ordered = false )
