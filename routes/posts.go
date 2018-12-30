package routes

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MixinNetwork/bot-api-go-client"
	"net/url"

	//"github.com/crossle/hacker-news-mixin-bot/config"
	"github.com/caosbad/ever-post-mixin-bot/config"
	"github.com/caosbad/ever-post-mixin-bot/middlewares"
	"github.com/caosbad/ever-post-mixin-bot/models"
	"github.com/caosbad/ever-post-mixin-bot/session"
	"github.com/caosbad/ever-post-mixin-bot/views"
	"github.com/dimfeld/httptreemux"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

type postsImpl struct{}

func registerPosts(router *httptreemux.TreeMux) {
	impl := &postsImpl{}
	router.GET("/api/posts/:id", impl.getPost)
	router.PUT("/api/posts/:id", impl.updatePost)
	router.POST("/api/posts/:id", impl.publishPost)
	router.GET("/api/verify/:id", impl.verifyTrace)
	router.GET("/api/myPosts/:type", impl.getUserPosts)
	router.GET("/api/posts", impl.getAllPosts)

	router.PUT("/api/drafts/:id", impl.updateDraft)
	router.POST("/api/drafts", impl.createDraft)
	router.DELETE("/api/drafts/:id", impl.deleteDraft)
	router.GET("/api/drafts/:id", impl.getUserDraft)
	router.GET("/api/drafts", impl.getUserDrafts)

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
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.TraceId != post.TraceId {
			views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
			return
		}
		current := middlewares.CurrentUser(r)
		body.PostId = post.PostId
		if post, err := models.PublishPost(r.Context(), current, body); err != nil {
			views.RenderErrorResponse(w, r, err)
		} else {
			views.RenderPost(w, r, post)
			//sendNotifyToSubscribers(r.Context(), post)
		}
	}

}

func sendNotifyToSubscribers(ctx context.Context, post *models.Post) {
	subscribers, err := models.ListSubscribers(ctx, post.UserId)
	if err == nil {
		for _, sub := range subscribers {
			conversationId := bot.UniqueConversationId(config.ClientId, sub.SubscriberId)
			data := base64.StdEncoding.EncodeToString([]byte(post.Title + " " + "https://ipfs.io/ipfs/" + post.IpfsId))
			err := bot.PostMessage(ctx, conversationId, sub.SubscriberId, bot.UuidNewV4().String(), "PLAIN_TEXT", data, config.ClientId, config.SessionId, config.PrivateKey)
			if err != nil {
				fmt.Print(err)
			}
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
		if body.IpfsId != "" {
			sendNotifyToSubscribers(r.Context(), post)
		}
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
	limit, offset, err := getUrlParams(r)
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
	limit, offset, err := getUrlParams(r)
	if err != nil {
		views.RenderErrorResponse(w, r, errors.New("Params error"))
		return
	}
	var postType = "draft"
	if val, ok := params["type"]; ok {
		postType = val
	}

	current := middlewares.CurrentUser(r)

	if list, count, err := models.FindPostsByUser(r.Context(), current, offset, limit, postType); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if list == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderPosts(w, r, list, count)
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
	limit, offset, err := getUrlParams(r)
	if err != nil {
		views.RenderErrorResponse(w, r, errors.New("Params error"))
		return
	}

	current := middlewares.CurrentUser(r)
	if list, count, err := models.FindDraftsByUser(r.Context(), current, offset, limit); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if list == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderPosts(w, r, list, count)
	}
}

func (impl *postsImpl) getAllPosts(w http.ResponseWriter, r *http.Request, params map[string]string) {
	limit, offset, err := getUrlParams(r)
	if err != nil {
		views.RenderErrorResponse(w, r, errors.New("Params error"))
	}
	if list, count, err := models.FindAllPosts(r.Context(), offset, limit); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else if list == nil {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context()))
	} else {
		views.RenderAllPosts(w, r, list, count)
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

func getUrlParams(r *http.Request) (int, int, error) {
	var offset, limit = 0, 20
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return 0, 0, err
	}
	if val, ok := query["offset"]; ok {
		offset, err = strconv.Atoi(val[0])
	}
	if val, ok := query["limit"]; ok {
		limit, err = strconv.Atoi(val[0])
	}
	return limit, offset, nil
}
