package uc

import (
	"context"
	"errors"
	"fmt"

	ctxpkg "github.com/Archiker-715/e-commerce-api/internal/auth/ctx"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
)

type UserCartService struct {
	repo           *pg.UserCartRepo
	productService ProdSrv
}

type ProdSrv interface {
	GetProductsByIds(productsId []uint) (products []entity.Product, err error)
}

func NewUserCartService(repo *pg.UserCartRepo) *UserCartService {
	return &UserCartService{repo: repo}
}

func (c *UserCartService) GetUserCart(ctx context.Context) (userCart []entity.UserCart, err error) {
	if userCart, err = c.repo.GetUserCart(ctxpkg.UserFromCtx(ctx)); err != nil {
		return []entity.UserCart{}, fmt.Errorf("get userCart error: %w", err)
	}

	if err = c.checkProductsQuantityInCart(userCart); err != nil {
		return []entity.UserCart{}, fmt.Errorf("checkProductsQuantityInCart error: %w", err)
	}
	return
}

func (c *UserCartService) GetProductsFromCartByIds(ctx context.Context, productsId []uint) (userCart []entity.UserCart, err error) {
	if userCart, err = c.repo.GetProductsFromCartByIds(ctxpkg.UserFromCtx(ctx), productsId); err != nil {
		return []entity.UserCart{}, fmt.Errorf("GetProductsFromCartByIds error: %w", err)
	}

	if err = c.checkProductsQuantityInCart(userCart); err != nil {
		return []entity.UserCart{}, fmt.Errorf("checkProductsQuantityInCart error: %w", err)
	}
	return
}

func (c *UserCartService) checkProductsQuantityInCart(userCart []entity.UserCart) error {
	prIds := make([]uint, len(userCart))
	for _, ucProduct := range userCart {
		prIds = append(prIds, ucProduct.ProductId)
	}

	products, err := c.productService.GetProductsByIds(prIds)
	if err != nil {
		return fmt.Errorf("get productsById error: %w", err)
	}

	productsQuantityInCartExceed := func(available uint) error {
		return fmt.Errorf("The permissible quantity has been exceeded. Available: %v", available)
	}

	if len(products) == len(userCart) {
		for i := 0; i < len(userCart); i++ {
			for _, product := range products {
				if userCart[i].ProductId == product.ProductID {
					if userCart[i].Count <= product.Count {
						continue
					} else {
						userCart[i].PurchaseAvailable = false
						userCart[i].NotAvailableReason = productsQuantityInCartExceed(product.Count)
					}
				}
				continue
			}
		}
	} else {
		return errors.New("not found all products from cart")
	}
	return nil
}

func (c *UserCartService) AddProductToCart(ctx context.Context, productId uint) error {
	return c.repo.AddProductToCart(productId, ctxpkg.UserFromCtx(ctx))
}

func (c *UserCartService) DeleteProductFromCart(ctx context.Context, productId uint) error {
	return c.repo.DeleteProductFromCart(productId, ctxpkg.UserFromCtx(ctx))
}

func (c *UserCartService) ChangeProductCount(ctx context.Context, productId uint, action string) error {
	switch action {
	case "increase":
		return c.repo.IncreaseProductInCart(productId, ctxpkg.UserFromCtx(ctx))
	case "decrease":
		return c.repo.DecreaseProductInCart(productId, ctxpkg.UserFromCtx(ctx))
	default:
		return errors.New("change param is not in increase or decrease")
	}
}

func (c *UserCartService) DeleteProductsFromCart(ctx context.Context, prIds []uint) error {
	return c.repo.DeleteProductsFromCart(ctxpkg.UserFromCtx(ctx), prIds)
}
