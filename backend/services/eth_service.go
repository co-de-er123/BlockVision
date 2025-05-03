package services

import (
	"log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial(os.Getenv("ETH_NODE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client:", err)
	}
}

func GetBlockDetails(blockNumber uint64) (*ethereum.Block, error) {
	block, err := client.BlockByNumber(nil, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, err
	}
	return block, nil
}

func GetTransactionDetails(txHash string) (*types.Transaction, error) {
	tx, pending, err := client.TransactionByHash(common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	log.Println("Pending:", pending)
	return tx, nil
}

func GetNetworkHealth() (map[string]interface{}, error) {
	// Here you can add metrics like network block height, gas used, etc.
	// For now, returning basic info.
	info := map[string]interface{}{
		"chain":   "Ethereum",
		"network": "Mainnet",
	}
	return info, nil
}
