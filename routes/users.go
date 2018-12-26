package routes

import (
	"encoding/json"
	"errors"
	"github.com/satori/go.uuid"
	"net/http"

	"github.com/caosbad/ever-post-mixin-bot/middlewares"
	"github.com/caosbad/ever-post-mixin-bot/models"
	"github.com/caosbad/ever-post-mixin-bot/session"
	"github.com/caosbad/ever-post-mixin-bot/views"
	"github.com/dimfeld/httptreemux"
)

type usersImpl struct{}

func registerUsers(router *httptreemux.TreeMux) {
	impl := &usersImpl{}
	router.POST("/auth", impl.authenticate)
	router.GET("/me", impl.me)
	router.GET("/assets", impl.assets)
	router.GET("/user/:id", impl.getUser)
	router.GET("/subscriber/:id", impl.isSubscribers)
	router.POST("/subscriber/:id", impl.subscribeUser)
	router.DELETE("/subscriber/:id", impl.cancelSubscribe)
	// router.GET("/tele", impl.registerTele)
}

func (impl *usersImpl) authenticate(w http.ResponseWriter, r *http.Request, params map[string]string) {
	// validate params
	var body struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil { // decode body post params
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
	} else if user, err := models.AuthenticateUserByOAuth(r.Context(), body.Code); err != nil { // get user by code 
		views.RenderErrorResponse(w, r, err)
	} else {
		views.RenderUserWithAuthentication(w, r, user)
	}
}

func (impl *usersImpl) me(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	views.RenderUserWithAuthentication(w, r, middlewares.CurrentUser(r))
}

func (impl *usersImpl) assets(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	current := middlewares.CurrentUser(r)
	if assets, err := models.Assets(r.Context(), current.AccessToken);
		err != nil {
		views.RenderErrorResponse(w, r, err)
	} else {
		views.RenderDataResponse(w, r, map[string]interface{}{"success": true, "assets": assets})
	}

}

func (impl *usersImpl) subscribeUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
	//var body struct {
	//	UserId string `json:"id"`
	//}
	userId := params["id"]
	if _, err := uuid.FromString(userId); err != nil {
		views.RenderErrorResponse(w, r, err)
		return
	}

	current := middlewares.CurrentUser(r)
	if current.UserId == userId {
		views.RenderErrorResponse(w, r, errors.New("Can not subscribe yourself"))
	} else if subscriber, err := models.CreateSubscriber(r.Context(), userId, current.UserId); err == nil {
		views.RenderDataResponse(w, r, map[string]interface{}{"isSub": true, "subscriber": subscriber})
	} else {
		views.RenderErrorResponse(w, r, err)
	}
}

func (impl *usersImpl) isSubscribers(w http.ResponseWriter, r *http.Request, params map[string]string) {
	//var body struct {
	//	UserId string `json:"id"`
	//}
	userId := params["id"]
	if _, err := uuid.FromString(userId); err != nil {
		views.RenderErrorResponse(w, r, err)
		return
	}

	current := middlewares.CurrentUser(r)
	if current.UserId == userId {
		views.RenderDataResponse(w, r, map[string]bool{"isSelf": true})
	} else if subscriber, err := models.FindSubscriber(r.Context(), userId, current.UserId); err == nil {
		views.RenderDataResponse(w, r, map[string]interface{}{"isSub": true, "subscriber": subscriber})
	} else {
		views.RenderDataResponse(w, r, map[string]bool{"isSub": false})
	}
}

func (impl *usersImpl) cancelSubscribe(w http.ResponseWriter, r *http.Request, params map[string]string) {
	userId := params["id"]
	if _, err := uuid.FromString(userId); err != nil {
		views.RenderErrorResponse(w, r, err)
		return
	}

	current := middlewares.CurrentUser(r)
	if err := models.RemoveSubscriber(r.Context(), userId, current.UserId); err == nil {
		views.RenderDataResponse(w, r, map[string]interface{}{
			"success": true,
		})
	} else {
		views.RenderErrorResponse(w, r, err)
	}

}

func (impl *usersImpl) getUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
	userId := params["id"]
	if _, err := uuid.FromString(userId); err != nil {
		views.RenderErrorResponse(w, r, err)
		return
	}
	if user, err := models.FindUserById(r.Context(), userId); err == nil {
		views.RenderUserView(w, r, user)
	} else {
		views.RenderErrorResponse(w, r, err)
	}

}

// func (impl *usersImpl) registerTele(w http.ResponseWriter, r *http.Request, _ map[string]string) {
// 	if err: = models.RegisterTelegraphAccessToken(r.Context(), body.Code) {
// 		views.RenderErrorResponse(w, r, err)
// 	} else {
// 		views.RenderUserWithAuthentication(w, r, user)
// 	}
// }
