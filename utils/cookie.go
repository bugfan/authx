package utils

import (
	"net/http"
	"time"
)

const cookieName = "_auth"
const cacheMaxMem = 1 * 1024 * 1024

func SetJWTDataInCookie(res http.ResponseWriter, attr map[string]string) error {
	jwtStr, err := GenJWT(attr)
	if err != nil {
		return err
	}
	expire := time.Now().Add(24 * 60 * time.Minute)
	cookie := &http.Cookie{Name: cookieName, Value: jwtStr, Expires: expire, Path: "/", HttpOnly: true}
	http.SetCookie(res, cookie)
	return nil
}

func GetJWTDataFromCookie(req *http.Request) (map[string]string, error) {
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	jwtString := cookie.Value
	return VerifyJWT(jwtString)
}

func DeleteJWTCookie(res http.ResponseWriter) {
	cookie := http.Cookie{Name: cookieName, Path: "/", MaxAge: -1}
	http.SetCookie(res, &cookie)
}
