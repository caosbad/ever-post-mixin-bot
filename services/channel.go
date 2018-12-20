package services

// import (
// 	"context"
// 	"encoding/base64"
// 	"fmt"
// 	"time"

// 	client "github.com/MixinNetwork/bot-api-go-client"
// 	"github.com/caosbad/ever-post-mixin-bot/models"
// 	"github.com/caosbad/ever-post-mixin-bot/session"
// )

// type ChannelService struct{}

// type ChannelMessage struct {
// 	blaze *client.BlazeClient
// }


// func (service *ChannelService) Run(ctx context.Context) error {
// 	bots, err := models.ListBots(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	stopped := make(chan bool, 1)
// 	for _, bot := range bots {
// 		go func(bot *models.Bot) {
// 			for {
// 				blazeClient := client.NewBlazeClient(bot.ClientId, bot.SessionId, bot.PrivateKey)
// 				channelMessage := ChannelMessage{blaze: blazeClient}
// 				if err := blazeClient.Loop(ctx, channelMessage); err != nil {
// 					session.Logger(ctx).Error(err)
// 				}
// 				session.Logger(ctx).Info("connection loop end")
// 				time.Sleep(time.Second)
// 			}
// 		}(bot)
// 	}
// 	<-stopped
// 	return nil
// }

// func (c ChannelMessage) OnMessage(ctx context.Context, msg client.MessageView, uid string) error {
// 	if msg.Category == client.MessageCategorySystemAccountSnapshot || msg.Category == client.MessageCategorySystemConversation || msg.ConversationId != client.UniqueConversationId(uid, msg.UserId) {
// 		return nil
// 	}

// 	bot, err := models.FindBotByClientId(ctx, uid)
// 	if err != nil {
// 		return nil
// 	}
// 	if bot.UserId == msg.UserId {
// 		subscribers, err := models.ListSubscribers(ctx, uid)
// 		if err != nil {
// 			return nil
// 		}
// 		content, _ := base64.StdEncoding.DecodeString(msg.Data)
// 		for _, subscriber := range subscribers {
// 			conversationId := client.UniqueConversationId(uid, subscriber.SubscriberId)
// 			c.blaze.SendMessage(ctx, conversationId, subscriber.SubscriberId, msg.Category, string(content), "")
// 		}
// 		return nil
// 	}
// 	// TODO change to subcribe users
// 	if msg.Category == "PLAIN_TEXT" {
// 		data, err := base64.StdEncoding.DecodeString(msg.Data)
// 		if err != nil {
// 			return client.BlazeServerError(ctx, err)
// 		}
// 		if "/start" == string(data) {
// 			_, err = models.CreateSubscriber(ctx, uid, msg.UserId)
// 			if err != nil {
// 				return err
// 			}
// 			if err := c.blaze.SendPlainText(ctx, msg, "订阅成功"); err != nil {
// 				return client.BlazeServerError(ctx, err)
// 			}

// 			conversationId := client.UniqueConversationId(uid, bot.UserId)
// 			count := models.CountSubscribers(ctx, uid)
// 			content := fmt.Sprintf("已订阅, 订阅人数: %d", count)
// 			err = c.blaze.SendMessage(ctx, conversationId, bot.UserId, "PLAIN_TEXT", content, msg.UserId)
// 			return nil
// 		} else if "/stop" == string(data) {
// 			err = models.RemoveSubscriber(ctx, uid, msg.UserId)
// 			if err != nil {
// 				println(err)
// 			}
// 			if err := c.blaze.SendPlainText(ctx, msg, "已取消订阅"); err != nil {
// 				return client.BlazeServerError(ctx, err)
// 			}
// 			return nil
// 		}
// 	}
// 	_, err = models.FindSubscriber(ctx, uid, msg.UserId)
// 	if err != nil {
// 		content := `发送 /start 订阅消息
// 发送 /stop 取消订阅`
// 		if err := c.blaze.SendPlainText(ctx, msg, content); err != nil {
// 			return client.BlazeServerError(ctx, err)
// 		}
// 	} else {
// 		content := `已订阅，发送 /stop 取消订阅，有任何问题，请直接回复，只有频道创建人可以看到`
// 		if err := c.blaze.SendPlainText(ctx, msg, content); err != nil {
// 			return client.BlazeServerError(ctx, err)
// 		}
// 		conversationId := client.UniqueConversationId(uid, bot.UserId)
// 		data, _ := base64.StdEncoding.DecodeString(msg.Data)
// 		if err = c.blaze.SendMessage(ctx, conversationId, bot.UserId, msg.Category, string(data), msg.UserId); err != nil {
// 			println(err)
// 		}
// 	}

// 	return nil
// }
