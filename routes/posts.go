package routes

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/MixinNetwork/bot-api-go-client"
	//"github.com/crossle/hacker-news-mixin-bot/config"
	"github.com/caosbad/ever-post-mixin-bot/config"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"

	"github.com/caosbad/ever-post-mixin-bot/middlewares"
	"github.com/caosbad/ever-post-mixin-bot/models"
	"github.com/caosbad/ever-post-mixin-bot/session"
	"github.com/caosbad/ever-post-mixin-bot/views"
	"github.com/dimfeld/httptreemux"
)

type postsImpl struct{}

func registerPosts(router *httptreemux.TreeMux) {
	impl := &postsImpl{}
	router.GET("/posts/:id", impl.getPost)
	router.PUT("/posts/:id", impl.updatePost)
	router.POST("/posts/:id", impl.publishPost)
	router.GET("/verify/:id", impl.verifyTrace)
	router.GET("/myPosts", impl.getUserPosts)
	router.GET("/posts", impl.getAllPosts)

	router.PUT("/drafts/:id", impl.updateDraft)
	router.POST("/drafts", impl.createDraft)
	router.DELETE("/drafts/:id", impl.deleteDraft)
	router.GET("/drafts/:id", impl.getUserDraft)
	router.GET("/drafts", impl.getUserDrafts)

	// router.GET("/posts/:id", impl.getPost)
}

func (impl *postsImpl) createDraft(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var body struct {
		Title           string `json:"title"`
		Description     string `json:"description"`
		Content         string `json:"content"`
		MarkdownContent string `json:"markdown"`
		TraceId         string `json:"traceId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	current := middlewares.CurrentUser(r)
	// postId := params["id"]
	if post, err := models.CreateDraft(r.Context(), current, body.Title, body.Description, body.Content); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else {
		views.RenderPost(w, r, post)
	}
}

func (impl *postsImpl) publishPost(w http.ResponseWriter, r *http.Request, params map[string]string) {
	postId := params["id"]
	if _, err := uuid.FromString(postId); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	} else if post, err := models.FindPostByPostId(r.Context(), postId); err != nil || post == nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
	} else if post != nil {
		var body models.PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
			return
		}
		current := middlewares.CurrentUser(r)
		body.PostId = post.PostId
		if post, err := models.PublishPost(r.Context(), current, body); err != nil {
			views.RenderErrorResponse(w, r, err)
		} else {
			views.RenderPost(w, r, post)
			sendNotifyToSubscribers(r.Context(), post)
		}
	}

}

func sendNotifyToSubscribers(ctx context.Context, post *models.Post) {
	subscribers, err := models.ListSubscribers(ctx, post.UserId)
	if err == nil {
		for _, sub := range subscribers {
			conversationId := bot.UniqueConversationId(config.ClientId, sub.UserId)
			data := base64.StdEncoding.EncodeToString([]byte(post.Title + " " + post.TelegraphUrl))
			bot.PostMessage(ctx, conversationId, sub.UserId, bot.UuidNewV4().String(), "PLAIN_TEXT", data, config.ClientId, config.SessionId, config.PrivateKey)
		}
	}
}

func (impl *postsImpl) updatePost(w http.ResponseWriter, r *http.Request, params map[string]string) {
	postId := params["id"]
	if _, err := uuid.FromString(postId); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}

	var body models.PostBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	current := middlewares.CurrentUser(r)
	body.PostId = postId

	if post, err := models.UpdatePost(r.Context(), current, body); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else {
		views.RenderPost(w, r, post)
		sendNotifyToSubscribers(r.Context(), post)
	}
}

func (impl *postsImpl) updateDraft(w http.ResponseWriter, r *http.Request, params map[string]string) {
	postId := params["id"]
	if _, err := uuid.FromString(postId); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}

	var body models.PostBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	current := middlewares.CurrentUser(r)
	body.PostId = postId

	if post, err := models.UpdateDraft(r.Context(), current, body); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else {
		views.RenderPost(w, r, post)
	}
}

func (impl *postsImpl) deleteDraft(w http.ResponseWriter, r *http.Request, params map[string]string) {
	postId := params["id"]
	if _, err := uuid.FromString(postId); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	current := middlewares.CurrentUser(r)

	if err := models.DeleteDraft(r.Context(), current, params["id"]); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else {
		result := make(map[string]bool)
		result["success"] = true
		views.RenderDataResponse(w, r, result)
	}
}

func (impl *postsImpl) getUserTelegraphPosts(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var offset, limit = 0, 20
	var err error
	if val, ok := params["offset"]; ok {
		offset, err = strconv.Atoi(val)
	}
	if val, ok := params["limit"]; ok {
		limit, err = strconv.Atoi(val)
	}
	if err != nil {
		views.RenderErrorResponse(w, r, errors.New("Params error"))
		return
	}

	current := middlewares.CurrentUser(r)
	if list, err := models.FindTelegraphPostsByUser(r.Context(), current, offset, limit); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if list == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderTelegraphPosts(w, r, list)
	}
}

func (impl *postsImpl) getUserPosts(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var offset, limit = 0, 20
	var err error
	if val, ok := params["offset"]; ok {
		offset, err = strconv.Atoi(val)
	}
	if val, ok := params["limit"]; ok {
		limit, err = strconv.Atoi(val)
	}
	if err != nil {
		views.RenderErrorResponse(w, r, errors.New("Params error"))
		return
	}

	current := middlewares.CurrentUser(r)
	if list, err := models.FindAllPostsByUser(r.Context(), current, offset, limit); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if list == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderPosts(w, r, list)
	}
}

func (impl *postsImpl) getPost(w http.ResponseWriter, r *http.Request, params map[string]string) {
	postId := params["id"]
	if _, err := uuid.FromString(postId); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	if post, err := models.FindPostByPostId(r.Context(), postId); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if post == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderPost(w, r, post)
	}
}

func (impl *postsImpl) getUserDraft(w http.ResponseWriter, r *http.Request, params map[string]string) {
	postId := params["id"]
	if _, err := uuid.FromString(postId); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	current := middlewares.CurrentUser(r)
	if post, err := models.FindPostByPostId(r.Context(), postId); err != nil && current.UserId == post.UserId {
		views.RenderErrorResponse(w, r, err)
	} else if post == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderPost(w, r, post)
	}
}

func (impl *postsImpl) getUserDrafts(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var offset, limit = 0, 20
	var err error
	if val, ok := params["offset"]; ok {
		offset, err = strconv.Atoi(val)
	}
	if val, ok := params["limit"]; ok {
		limit, err = strconv.Atoi(val)
	}
	if err != nil {
		views.RenderErrorResponse(w, r, errors.New("Params error"))
		return
	}

	current := middlewares.CurrentUser(r)
	if list, err := models.FindDraftsByUser(r.Context(), current, offset, limit); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if list == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderPosts(w, r, list)
	}
}

func (impl *postsImpl) getAllPosts(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var offset, limit = 0, 20
	var err error
	if val, ok := params["offset"]; ok {
		offset, err = strconv.Atoi(val)
	}
	if val, ok := params["limit"]; ok {
		limit, err = strconv.Atoi(val)
	}
	if err != nil {
		views.RenderErrorResponse(w, r, errors.New("Params error"))
	}
	if list, err := models.FindAllPosts(r.Context(), offset, limit); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if list == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderPosts(w, r, list)
	}
}

// TODO
func (impl *postsImpl) verifyTrace(w http.ResponseWriter, r *http.Request, params map[string]string) {
	current := middlewares.CurrentUser(r)

	if post, err := models.VerifyTrace(r.Context(), current, params["id"]); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if post == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderPost(w, r, post)
	}
}
