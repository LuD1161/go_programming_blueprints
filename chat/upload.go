package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/stretchr/objx"
)

func uploaderHandler(w http.ResponseWriter, req *http.Request) {
	// BUG : userID should not be user trusted value
	// Solution : Take this value from cookie, which should be signed
	// use JWT
	// userId := req.FormValue("userId")
	authCookie, err := req.Cookie("auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userData := objx.MustFromBase64(authCookie.Value)
	var userId string
	var ok bool
	if userId, ok = userData["userId"].(string); !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filename := path.Join("avatars", userId+path.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Successful")
}
