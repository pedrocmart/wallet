package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pedrocmart/wallet/initdb"
	"github.com/pedrocmart/wallet/internal/controllers"
	"github.com/pedrocmart/wallet/internal/handlers"
	"github.com/pedrocmart/wallet/internal/models"
	"github.com/sirupsen/logrus"
)

const (
	dbuser        = "root"
	dbpassword    = "root"
	dbname        = "wallet"
	cacheserver   = "localhost:6379"
	cachepassword = ""
	cachedb       = 0
)

var (
	err     error
	service *models.DBService
	dbInfo  = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		dbuser, dbpassword, dbname,
	)
)

func main() {
	// Initialize logger
	log := logrus.New()
	log.SetOutput(os.Stdout)

	// Initialize DB connection
	service, err = models.NewDBService()
	if err != nil {
		panic(err)
	}
	defer service.Close()
	service.AutoMigrate() // Attempt to initialize with wallet table

	// Insert wallets with IDs 123 and 456
	initdb.InitDB()

	// Wallet controller that uses a cache service on top of underlying storage
	walletC := &controllers.CacheStore{
		controllers.NewCacheService(cacheserver, cachepassword, cachedb),
		controllers.NewWalletController(service),
	}

	// Initialize handlers
	handlrs := handlers.NewHandlers(log, walletC)

	r := gin.Default()
	// GET the wallet balance
	r.GET(handlers.EndpointGETBalance, handlrs.NewGetBalanceHandler())
	// POST credit to the wallet balance
	r.POST(handlers.EndpointPOSTCredit, handlrs.NewPostCreditHandler())
	// POST debit to the wallet balance
	r.POST(handlers.EndpointPOSTDebit, handlrs.NewPostDebitHandler())
	r.Run(":5001")
}
