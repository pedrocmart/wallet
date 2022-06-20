package controllers

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/pedrocmart/wallet/internal/wallet"
	"github.com/shopspring/decimal"
)

const cacheSeconds = 0

var ctx = context.TODO()

// Store is an interface that implements functions
// to be used for cache and storage alike
type Store interface {
	GetBalance(wid string) (decimal.Decimal, error)
	Credit(wid, amount string) (decimal.Decimal, error)
	Debit(wid, amount string) (decimal.Decimal, error)
}

// NewCacheService returns a new cache service object
func NewCacheService(addr, pw string, db int) *CacheService {

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		DB:   0,
	})
	return &CacheService{
		Cache: rdb,
	}
}

// CacheService represents the cache connection service
type CacheService struct {
	Cache *redis.Client
}

// CacheStore wraps any object that satisfies the Store interface
// with caching capabilities
type CacheStore struct {
	*CacheService
	Store
}

// GetBalance attempts to fetch a wid's balance from cache, gets it from storage otherwise
func (chs *CacheStore) GetBalance(wid string) (decimal.Decimal, error) {
	// Check cache
	b, err := chs.Cache.Get(ctx, wid).Result()
	switch {
	case err == redis.Nil:
		// Not found in cache
		var bal decimal.Decimal
		// Get the balance from storage
		bal, err = chs.Store.GetBalance(wid)
		if err != nil {
			return none, err
		}

		// Set wid and balance in cache
		err = chs.Cache.Set(ctx, wid, bal.StringFixed(wallet.Fixed), cacheSeconds).Err()
		if err != nil {
			return none, err
		}
		return bal, nil
	case err != nil:
		// Some error
		return none, err
	default:
		// Found in cache
		var balance decimal.Decimal
		balance, err = decimal.NewFromString(b)

		return balance, nil
	}
}

// Credit updates the balance in storage and cache
func (chs *CacheStore) Credit(wid, amount string) (decimal.Decimal, error) {
	// Set the balance in the storage
	bal, err := chs.Store.Credit(wid, amount)
	if err != nil {
		return none, err
	}

	// Update the cache
	err = chs.Cache.Set(ctx, wid, bal.StringFixed(wallet.Fixed), cacheSeconds).Err()
	if err != nil {
		return none, err
	}

	return bal, nil
}

// Debit updates the balance in storage and cache
func (chs *CacheStore) Debit(wid, amount string) (decimal.Decimal, error) {
	// Set the balance in the storage
	bal, err := chs.Store.Debit(wid, amount)
	if err != nil {
		return none, err
	}

	// Update the cache
	err = chs.Cache.Set(ctx, wid, bal.StringFixed(wallet.Fixed), cacheSeconds).Err()
	if err != nil {
		return none, err
	}

	return bal, nil
}
