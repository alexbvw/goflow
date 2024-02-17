package controllers

import (
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

func GenerateEtherWallet(c *fiber.Ctx) error {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not generate private key",
		})
	}
	privateKeyData := crypto.FromECDSA(privateKey)
	publicKeyData := crypto.FromECDSAPub(&privateKey.PublicKey)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address":    crypto.PubkeyToAddress(privateKey.PublicKey).Hex(),
		"publicKey":  hexutil.Encode(publicKeyData),
		"privateKey": hexutil.Encode(privateKeyData),
	})
}

func CreateEncrpytedWallets(c *fiber.Ctx) error {
	password := c.Query("password")
	keyStore := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := keyStore.NewAccount(password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not generate wallet",
		})
	}
	_, err = keyStore.NewAccount(password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not generate wallet",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "wallets have been generated",
	})
}

// func CreateEtherClient(c *fiber.Ctx) error {
// 	address := c.Params("address")
// 	client, err :=	ethclient.Dial(url)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "could not dial client",
// 		})
// 	}
// 	defer client.Close()
// 	common.HexToAddress(address)
// 	client.BalanceAt(c.Context(), )
// }

func GetWalletAddressBalance(c *fiber.Ctx) error {
	address1 := c.Params("address1")
	address2 := c.Params("address2")
	client, err := ethclient.Dial(os.Getenv("INFURAURL"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not dial client",
		})
	}
	defer client.Close()
	a1 := common.HexToAddress(address1)
	a2 := common.HexToAddress(address2)
	b1, err1 := client.BalanceAt(c.Context(), a1, nil)
	b2, err2 := client.BalanceAt(c.Context(), a2, nil)
	if err1 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not retrieve address balance",
		})
	}
	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not retrieve address balance",
		})
	}
	//1 ether = 10^18 wei
	floatBalance1 := new(big.Float)
	floatBalance1.SetString(b1.String())
	value1 := new(big.Float).Quo(floatBalance1, big.NewFloat(math.Pow10(18)))
	//1 ether = 10^18 wei
	floatBalance2 := new(big.Float)
	floatBalance2.SetString(b2.String())
	value2 := new(big.Float).Quo(floatBalance2, big.NewFloat(math.Pow10(18)))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":            false,
		"address_balance1": value1,
		"address_balance2": value2,
	})
}

func GetLastEtherBlock(c *fiber.Ctx) error {
	client, err := ethclient.DialContext(c.Context(), os.Getenv("INFURAURL"))
	if err != nil {
		defer client.Close()
	}
	block, err := client.BlockByNumber(c.Context(), nil)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not get block",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":             false,
		"last_block_number": block.Number().String(),
	})
}

func RetriveEtherhWalletBalance(c *fiber.Ctx) error {
	client, err := ethclient.DialContext(c.Context(), os.Getenv("INFURAURL"))
	if err != nil {
		defer client.Close()
	}
	addr := c.Query("address")
	address := common.HexToAddress(addr)
	balance, err := client.BalanceAt(c.Context(), address, nil)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not get balance",
		})
	}
	//1 ether = 10^18 wei
	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	value := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":         false,
		"ether_balance": value,
	})
}
