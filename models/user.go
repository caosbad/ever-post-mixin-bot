package models

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"

	"github.com/MixinNetwork/bot-api-go-client"
	"github.com/caosbad/ever-post-mixin-bot/config"
	"github.com/caosbad/ever-post-mixin-bot/session"
	jwt "github.com/dgrijalva/jwt-go"
	"gitlab.com/toby3d/telegraph"
)

const users_DDL = `
CREATE TABLE IF NOT EXISTS users (
	user_id	          VARCHAR(36) PRIMARY KEY,
	access_token      VARCHAR(512),
	full_name         VARCHAR(512),
	avatar_url        VARCHAR(1024),
	telegraph_token   VARCHAR(64),
	created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
  );
CREATE INDEX ON users (user_id);
`

const (
	DefaultAvatar = "https://images.mixin.one/E2y0BnTopFK9qey0YI-8xV3M82kudNnTaGw0U5SU065864SsewNUo6fe9kDF1HIzVYhXqzws4lBZnLj1lPsjk-0=s128"
)

type User struct {
	UserId              string    `sql:"user_id,pk"`
	AccessToken         string    `sql:"access_token"`
	FullName            string    `sql:"full_name"`
	AvatarURL           string    `sql:"avatar_url"`
	CreatedAt           time.Time `sql:"created_at,notnull"`
	TelegraphToken      string 		`sql:"telegraph_token"`
	AuthenticationToken string    `sql:"-"`
}

var userCols = []string{"user_id", "access_token", "full_name", "avatar_url", "telegraph_token"}

func AuthenticateUserByOAuth(ctx context.Context, authorizationCode string) (*User, error) {
	accessToken, scope, err := bot.OAuthGetAccessToken(ctx, config.ClientId, config.ClientSecret, authorizationCode, "")
	if err != nil {
	    fmt.Print(err)
		return nil, session.ServerError(ctx, err)
	}
	if !strings.Contains(scope, "PROFILE:READ") {
		return nil, session.ForbiddenError(ctx)
	}
	me, err := bot.UserMe(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	user, err := FindUserById(ctx, me.UserId)
	if err != nil {
		println(err)
		return nil, err
	}

	var telegraphAccount telegraph.Account
	if user == nil {
		telegraphAccount, err := telegraph.CreateAccount(&telegraph.Account{
			ShortName:  me.FullName,
			AuthorName: me.FullName,
		})
		if err != nil {
			println(err)
			return nil, err
		}
		user, err = createUser(ctx, me.UserId, accessToken, me.FullName, me.AvatarURL, telegraphAccount.AccessToken)
	} else {
		// todo
		telegraphAccount.AccessToken = user.TelegraphToken
		_, err := telegraphAccount.EditAccountInfo(&telegraph.Account{
			ShortName:  me.FullName,
			AuthorName: me.FullName,
		})
	if err != nil {
			return nil, err
		}
		user, err = user.updateProfile(ctx, accessToken, me.FullName, me.AvatarURL)
	}
	if err != nil {
		return nil, err
	}
	authenticationToken, err := generateAuthenticationToken(ctx, user.UserId, accessToken)
	if err != nil {
		return nil, session.BadDataError(ctx)
	}
	user.AuthenticationToken = authenticationToken
	return user, nil
}

func generateAuthenticationToken(ctx context.Context, userId, accessToken string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
	})
	sum := sha256.Sum256([]byte(accessToken))
	return token.SignedString(sum[:])
}
// find user info with authenticationToken and query user info 
func AuthenticateUserByToken(ctx context.Context, authenticationToken string) (*User, error) {
	var user *User
	var queryErr error
	token, err := jwt.Parse(authenticationToken, func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, session.BadDataError(ctx)
		}
		_, ok = token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, session.BadDataError(ctx)
		}
		user, queryErr = FindUserById(ctx, fmt.Sprint(claims["jti"]))
		if queryErr != nil {
			return nil, queryErr
		}
		if user == nil {
			return nil, nil
		}
		sum := sha256.Sum256([]byte(user.AccessToken))
		return sum[:], nil
	})

	if queryErr != nil {
		return nil, queryErr
	}
	if err != nil || !token.Valid {
		return nil, nil
	}
	return user, nil
}

func createUser(ctx context.Context, userId, accessToken, fullName, avatarURL, telegraphToken string) (*User, error) {
	if _, err := uuid.FromString(userId); err != nil {
		return nil, session.BadDataError(ctx)
	}
	if avatarURL == "" {
		avatarURL = DefaultAvatar
	}
	user := &User{
		UserId:      userId,
		FullName:    fullName,
		AvatarURL:   avatarURL,
		TelegraphToken: telegraphToken,
		AccessToken: accessToken,
		CreatedAt:   time.Now(),
	}
	if err := session.Database(ctx).Insert(user); err != nil {
		return nil, session.TransactionError(ctx, err)
	}
	return user, nil
}

func (user *User) updateProfile(ctx context.Context, accessToken, fullName, avatarURL string) (*User, error) {
	if avatarURL == "" {
		avatarURL = DefaultAvatar
	}
	user.AccessToken, user.FullName, user.AvatarURL = accessToken, fullName, avatarURL
	if _, err := session.Database(ctx).Model(user).Column("access_token", "full_name", "avatar_url").WherePK().Update(); err != nil {
		return nil, session.TransactionError(ctx, err)
	}
	return user, nil
}

func FindUserById(ctx context.Context, userId string) (*User, error) {
	user := &User{
		UserId: userId,
	}
	if err := session.Database(ctx).Model(user).Column(userCols...).WherePK().Select(user); err == pg.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, session.TransactionError(ctx, err)
	}
	return user, nil
}
