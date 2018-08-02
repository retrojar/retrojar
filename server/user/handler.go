package user

import (
	goContext "context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"os"
	"retro-memo/server/httputil"
	"time"
)

type handler struct {
	config   *oauth2.Config
	secret   string
	clientId string
}

func NewHandler() *handler {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	return &handler{
		config,
		os.Getenv("APP_SECRET"),
		os.Getenv("GOOGLE_CLIENT_ID"),
	}
}

// @Summary get login url
// @Accept json
// @Produce json
// @Success 200 {object} user.LoginUrlRepsonse
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/login/google [get]
func (h *handler) GetLoginUrl(context *gin.Context) {
	url := h.config.AuthCodeURL("")
	context.JSON(http.StatusOK, LoginUrlRepsonse{url})
}

// @Summary get jwt
// @Accept json
// @Produce json
// @Success 200 {object} user.JWTRepsonse
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/login/google/jwt [get]
func (h *handler) GetJWT(context *gin.Context) {
	profile, err := h.queryGoogleProfile(context.Query("code"))
	if err != nil {
		httputil.NewError(context, http.StatusBadRequest, err)
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud":   h.clientId,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
		"iat":   time.Now().Unix(),
		"name":  profile.Name,
		"pic":   profile.Picture,
		"email": profile.Email,
	})

	tokenString, err := jwtToken.SignedString([]byte(h.secret))
	if err != nil {
		httputil.NewError(context, http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, JWTRepsonse{tokenString})
}

func (h *handler) queryGoogleProfile(code string) (googleProfile, error) {
	profile := googleProfile{}

	token, err := h.config.Exchange(goContext.Background(), code)
	if err != nil {
		return profile, err
	}

	client := h.config.Client(goContext.Background(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return profile, err
	}

	defer userInfo.Body.Close()
	data, _ := ioutil.ReadAll(userInfo.Body)

	if err := json.Unmarshal(data, &profile); err != nil {
		return profile, err
	}

	return profile, nil
}

type googleProfile struct {
	Name    string `json:"Name"`
	Email   string `json:"Email"`
	Picture string `json:"Picture"`
}
