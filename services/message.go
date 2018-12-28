package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/MixinNetwork/go-number"
	"github.com/caosbad/ever-post-mixin-bot/models"
	"github.com/satori/go.uuid"
	"time"

	"github.com/MixinNetwork/bot-api-go-client"
	"github.com/caosbad/ever-post-mixin-bot/config"
	"github.com/caosbad/ever-post-mixin-bot/session"
	// uuid "github.com/satori/go.uuid"
)

type BlazeService struct {
}

type ResponseMessage struct {
	blazeClient *bot.BlazeClient
}

func (service *BlazeService) Run(ctx context.Context) error {
	for {
		blazeClient := bot.NewBlazeClient(config.ClientId, config.SessionId, config.PrivateKey)
		response := ResponseMessage{
			blazeClient: blazeClient,
		}
		if err := blazeClient.Loop(ctx, response); err != nil {
			session.Logger(ctx).Error(err)
		}
		session.Logger(ctx).Info("connection loop end")
		time.Sleep(time.Second)
	}
	return nil
}

func (r ResponseMessage) OnMessage(ctx context.Context, msg bot.MessageView, uid string) error {
	if msg.Category == bot.MessageCategorySystemAccountSnapshot {
		data, err := base64.StdEncoding.DecodeString(msg.Data)
		if err != nil {
			return bot.BlazeServerError(ctx, err)
		}
		var transfer bot.TransferView
		err = json.Unmarshal(data, &transfer)
		if err != nil {
			return bot.BlazeServerError(ctx, err)
		}
		res, err := handleTransfer(ctx, transfer, msg.UserId, r.blazeClient)
		if err != nil {
			return bot.BlazeServerError(ctx, err)
		} else {
			return sendText(ctx, msg, res)
		}
	}
	if msg.Category == bot.MessageCategoryPlainText {

		if _, err := base64.StdEncoding.DecodeString(msg.Data); err != nil {
			return bot.BlazeServerError(ctx, err)
		} else if err = handleText(ctx, r.blazeClient, msg); err != nil {
			return bot.BlazeServerError(ctx, err)
		}
	}
	return nil
}

// text message
func handleText(ctx context.Context, client *bot.BlazeClient, msg bot.MessageView) error {
	// _, err := uuid.FromString(transfer.TraceId)
	// if err != nil {
	// 	return nil
	// }
	// if transfer.Amount != config.PayAmount || transfer.AssetId != config.PayAssetId {
	// 	return session.BadDataError(ctx)
	// }
	// _, err = models.CreateChannelBot(ctx, userId, transfer.TraceId)
	if err := client.SendPlainText(ctx, msg, "订阅成功"); err != nil {
		return bot.BlazeServerError(ctx, err)
	}
	return nil
}

// pay message incliude post,
func handleTransfer(ctx context.Context, transfer bot.TransferView, userId string, client *bot.BlazeClient) (string, error) {

	if _, err := uuid.FromString(transfer.Memo); err == nil {
		if post, err := models.FindPostByPostId(ctx, transfer.Memo); err != nil {
			return "", bot.BlazeServerError(ctx, err)
		} else if post.UserId == transfer.CounterUserId { // publish or ipfs article payment
			if post.Path == "" && transfer.Amount == config.CNBAmount && transfer.AssetId == config.CNBAssetId { // publish
				post.TraceId = transfer.TraceId
				_, err = models.UpdatePostTraceId(ctx, post)
				if err != nil {
					return "", bot.BlazeServerError(ctx, err)
				}
				return "payment successfully...", nil
			} else if post.Path != "" && post.TraceId != "" && transfer.Amount == config.CNBAmount && transfer.AssetId == config.CNBAssetId {
				post.TraceId = transfer.TraceId
				_, err = models.UpdatePostTraceId(ctx, post)
				if err != nil {
					return "", bot.BlazeServerError(ctx, err)
				}
				return "payment successfully...", nil
			} else {
				return "Donation successfully", nil
			}

		} else { // vote article payment
			if err := sendTransfer(ctx, &bot.TransferInput{
				AssetId:     transfer.AssetId,
				RecipientId: post.UserId,
				Amount:      number.FromString(transfer.Amount),
				TraceId:     uuid.Must(uuid.NewV4()).String(),
				Memo:        "Donate from your reader.",
			}); err != nil {
				return "", bot.BlazeServerError(ctx, err)
			} else if support, err := models.CreateSupport(ctx, post, transfer); err != nil || support == nil {
				return "", bot.BlazeServerError(ctx, err)
			} else {
				return "Donation successfully", nil
			}
		}

	} else { // no pain no gain
		//if err := sendTransfer(ctx, &bot.TransferInput{
		//	AssetId:     transfer.AssetId,
		//	RecipientId: transfer.CounterUserId,
		//	Amount:      number.FromString(transfer.Amount),
		//	TraceId:     uuid.Must(uuid.NewV4()).String(),
		//	Memo:        "invalid transfer",
		//}); err != nil {
		//	return nil, bot.BlazeServerError(ctx, err)
		//}
		//return "Thanks for your donation...", nil
		return "Thx for your donation...", nil
	}

}

func sendTransfer(ctx context.Context, transfer *bot.TransferInput) error {

	err := bot.CreateTransfer(ctx, transfer, config.ClientId, config.SessionId, config.PrivateKey, config.Pin, config.PinToken)
	if err != nil {
		return session.ServerError(ctx, err)
	}
	return nil
}

func sendText(ctx context.Context, msg bot.MessageView, content string) error {
	//conversationId := bot.UniqueConversationId(config.ClientId, sub.UserId)
	data := base64.StdEncoding.EncodeToString([]byte(content))
	err := bot.PostMessage(ctx, msg.ConversationId, msg.UserId, bot.UuidNewV4().String(), "PLAIN_TEXT", data, config.ClientId, config.SessionId, config.PrivateKey)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func sendPaidMessage(ctx context.Context, client *bot.BlazeClient, msg bot.MessageView, content string) error {
	if err := client.SendPlainText(ctx, msg, content); err != nil {
		return bot.BlazeServerError(ctx, err)
	}
	//imageMap := map[string]interface{}{
	//	"attachment_id": "2cd57e11-a58e-4705-a47e-77f0586c915e",
	//	"size":          316047,
	//	"width":         1532,
	//	"mime_type":     "image/jpeg",
	//	"height":        1098,
	//}
	//imageData, _ := json.Marshal(imageMap)
	//if err := client.SendMessage(ctx, msg.ConversationId, msg.UserId, bot.MessageCategoryPlainImage, string(imageData), ""); err != nil {
	//	return bot.BlazeServerError(ctx, err)
	//}
	return nil
}
