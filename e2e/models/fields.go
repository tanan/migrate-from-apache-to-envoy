package models

type Fields struct {
	Scheme          string
	Host            string
	Path            string
	Status          int
	RequestCookie   map[string]string
	RespHeader      map[string]string
	RespCookieNames []string
	Service         string
	Message         string
}
