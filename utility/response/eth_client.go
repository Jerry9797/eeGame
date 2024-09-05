package response

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/v2/frame/g"
)

var EthClient *ethclient.Client

// 启动eth client
func init() {
	url, err := g.Cfg().Get(context.Background(), "ETH.url")
	if err != nil {
		panic(err)
	}
	client, err := ethclient.DialContext(context.Background(), url.String())
	if err != nil {
		panic(err)
	}
	EthClient = client
}
