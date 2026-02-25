package redis

import (
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
)

func (c *client) LockProducts(productsIds []uint) ([]*redsync.Mutex, error) {
	mutex := make([]*redsync.Mutex, len(productsIds))
	for _, productId := range productsIds {
		m := c.redsyncInstance.NewMutex(
			fmt.Sprintf("productId:%v", productId),
			redsync.WithExpiry(10*time.Second),
			redsync.WithTries(10),
			redsync.WithRetryDelay(100*time.Millisecond),
		)
		mutex = append(mutex, m)

		if err := m.Lock(); err != nil {
			return nil, err
		}
	}
	return mutex, nil
}

func (c *client) UnlockProducts(mutex []*redsync.Mutex) error {
	for _, m := range mutex {
		if ok, err := m.Unlock(); !ok || err != nil {
			return err
		}
	}
	return nil
}
