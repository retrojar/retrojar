package user

type LoginUrlRepsonse struct {
	URL string `json:"url" example:"http://google.com/...."`
}

type JWTRepsonse struct {
	JWT string `json:"jwt" example:"JhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOi...."`
}
