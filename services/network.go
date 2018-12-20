package services

// import (
// 	"context"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	client "github.com/MixinNetwork/bot-api-go-client"
// 	"github.com/caosbad/ever-post-mixin-bot/models"
// )

// type MixinNetworkService struct {
// }

// func (service *MixinNetworkService) Run(ctx context.Context) error {
// 	bot, err := models.FindBotByBotId(ctx, "abc167a8-1e8c-4eef-bfc7-c1475107eced")
// 	if err != nil {
// 		return err
// 	}
// 	for {
// 		chains, err := fetchChain(ctx)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		for _, chain := range chains {
// 			if !chain.IsSynchronized {
// 				subscribers, err := models.ListSubscribers(ctx, bot.ClientId)
// 				if err != nil {
// 					return nil
// 				}

// 				content := base64.StdEncoding.EncodeToString([]byte(chain.Name + " 区块同步异常"))
// 				for _, subscriber := range subscribers {
// 					conversationId := client.UniqueConversationId(bot.ClientId, subscriber.SubscriberId)
// 					client.PostMessage(ctx, conversationId, subscriber.SubscriberId, client.UuidNewV4().String(), "PLAIN_TEXT",
// 						content, bot.ClientId, bot.SessionId, bot.PrivateKey)
// 				}
// 			}
// 		}
// 		time.Sleep(1 * time.Minute)
// 	}
// 	return nil
// }

// type NetworkChain struct {
// 	ChainId        string `json:"chain_id"`
// 	Name           string `json:"name"`
// 	IsSynchronized bool   `json:"is_synchronized"`
// }

// func fetchChain(ctx context.Context) ([]NetworkChain, error) {
// 	resp, err := http.Get("https://api.mixin.one/network")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	var block struct {
// 		Data struct {
// 			Chains []NetworkChain `json:"chains"`
// 		} `json:"data"`
// 	}
// 	err = json.NewDecoder(resp.Body).Decode(&block)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return block.Data.Chains, nil
// }
