package views

import (
	"github.com/caosbad/ever-post-mixin-bot/models"
	"net/http"
)

type UserView struct {
	Type   string `json:"type"`
	UserId string `json:"user_id"`
	//IdentityNumber string `json:"identity_number"`
	FullName       string `json:"full_name"`
	AvatarURL      string `json:"avatar_url"`
	TelegraphToken string `json:"telegraph_token"`
}

type UserWithAuthenticationView struct {
	UserView
	AuthenticationToken string `json:"authentication_token"`
}

type SubscriberView struct {
	UserId       string `json:"user_id"`
	SubscriberId string `json::"subscriber_id"`
}

func buildUserView(user *models.User) UserView {
	userView := UserView{
		Type:           "user",
		UserId:         user.UserId,
		FullName:       user.FullName,
		AvatarURL:      user.AvatarURL,
		TelegraphToken: user.TelegraphToken}
	return userView
}

func RenderUserView(w http.ResponseWriter, r *http.Request, user *models.User) {
	RenderDataResponse(w, r, buildUserView(user))
}
func RenderUserWithAuthentication(w http.ResponseWriter, r *http.Request, user *models.User) {
	userView := UserWithAuthenticationView{
		UserView:            buildUserView(user),
		AuthenticationToken: user.AuthenticationToken,
	}
	RenderDataResponse(w, r, userView)
}
func RenderSubscriberView(w http.ResponseWriter, r *http.Request, subscriber *models.Subscriber) {
	subscriberView := SubscriberView{
		UserId:       subscriber.UserId,
		SubscriberId: subscriber.SubscriberId,
	}
	RenderDataResponse(w, r, subscriberView)
}
