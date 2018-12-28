package views

import (
	"net/http"
	"time"

	"github.com/caosbad/ever-post-mixin-bot/models"
	"github.com/caosbad/telegraph"
)

type PostView struct {
	PostId        string    `json:"post_id"`
	Title         string    `json:"title"`
	Path          string    `json:"path"`
	Telegraph_url string    `json:"telegraph_url"`
	Description   string    `json:"description"`
	IpfsId        string    `json:"ipfs_id"`
	Content       string    `json:"content"`
	UserId        string    `json:"user_id"`
	TraceId       string    `json:"trace_id"`
	CreatedAt     time.Time `json:"created_at"`
	AvatarURL     string    `json:"avatar_url"`
}

// type PostsView struct {
// 	TotalCount 		int  `json:"total_count"`
// 	Posts 				[]telegraph.Page `json:"posts"`
// }

func buildPostView(post *models.Post) PostView {
	view := PostView{
		PostId:        post.PostId,
		Title:         post.Title,
		Path:          post.Path,
		Telegraph_url: post.TelegraphUrl,
		Description:   post.Description,
		IpfsId:        post.IpfsId,
		Content:       post.Content,
		UserId:        post.UserId,
		TraceId:       post.TraceId,
		CreatedAt:     post.CreatedAt,
	}
	return view
}

// return telegraph post list 
// func buildTelegraphPostsView(list *telegraph.PageList) PostsView {
// 	return list
// }

func buildPostsView(list []*models.Post) []PostView {
	var posts []PostView
	for _, post := range list {
		post := PostView{
			PostId:        post.PostId,
			Title:         post.Title,
			Path:          post.Path,
			Telegraph_url: post.TelegraphUrl,
			Description:   post.Description,
			IpfsId:        post.IpfsId,
			Content:       post.Content,
			UserId:        post.UserId,
			CreatedAt:     post.CreatedAt,
		}
		posts = append(posts, post)
	}
	return posts
}

func buildPostsItemsView(list []*models.PostListItem) []PostView {
	var posts []PostView
	for _, post := range list {
		post := PostView{
			PostId:        post.PostId,
			Title:         post.Title,
			Path:          post.Path,
			Telegraph_url: post.TelegraphUrl,
			Description:   post.Description,
			IpfsId:        post.IpfsId,
			Content:       post.Content,
			UserId:        post.UserId,
			CreatedAt:     post.CreatedAt,
			AvatarURL:     post.AvatarURL,
		}
		posts = append(posts, post)
	}
	return posts
}

func RenderPost(w http.ResponseWriter, r *http.Request, post *models.Post) {
	RenderDataResponse(w, r, buildPostView(post))
}

func RenderPosts(w http.ResponseWriter, r *http.Request, posts []*models.Post) {
	RenderDataResponse(w, r, buildPostsView(posts))
}
func RenderAllPosts(w http.ResponseWriter, r *http.Request, posts []*models.PostListItem) {
	RenderDataResponse(w, r, buildPostsItemsView(posts))
}

func RenderTelegraphPosts(w http.ResponseWriter, r *http.Request, list *telegraph.PageList) {
	RenderDataResponse(w, r, list)
}
