package router

import (
	"goflow/interface/controllers"
	"goflow/interface/middleware"

	"github.com/gofiber/fiber/v2"
)

// AccountRoutes func for describe group of private routes.
func AcountRoutes(a *fiber.App) {
	//Create routes group.
	account := a.Group("/v1/account")

	//Ethereum Blockchain
	account.Get("/last-ether-block", controllers.GetLastEtherBlock)

	//Ether Wallet
	account.Post("/generate-wallet", middleware.AdminProtected(), controllers.GenerateEtherWallet)
	account.Get("/address-balance", middleware.AdminProtected(), controllers.RetriveEtherhWalletBalance)

	//Test Wallets
	account.Get("/wallets-address-balance", middleware.AdminProtected(), controllers.GetWalletAddressBalance)
	account.Post("/generate-encrpyted-wallets", middleware.AdminProtected(), controllers.CreateEncrpytedWallets)
}
