package main

import (
	"errors"
	"io/ioutil"
	"path"
)

// ErrNoAvator is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

// Avatar represents types capable of representing
// user profile pictures.
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client,
	// or returns an error if something goes wrong.
	// ErrNoAvatarURL is returned if the object is unable to get
	// a URL for the specified client
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

// GetAvatarURL : To get the avatar url
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	// defining these vars to maintain
	// line of sight
	var urlStr string
	var ok bool
	if urlStr, ok = c.userData["avatar_url"].(string); !ok {
		return "", ErrNoAvatarURL
	}
	return urlStr, nil
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

// GetAvatarURL : To get the avatar url
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	// defining these vars to maintain
	// line of sight
	var userId string
	var ok bool
	if userId, ok = c.userData["userId"].(string); !ok {
		return "", ErrNoAvatarURL
	}
	return "//www.gravatar.com/avatar/" + userId, nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL : To get the avatar url
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	// defining these vars to maintain
	// line of sight
	var userId string
	var ok bool
	if userId, ok = c.userData["userId"].(string); !ok {
		return "", ErrNoAvatarURL
	}
	files, err := ioutil.ReadDir("avatars")
	if err != nil {
		return "", ErrNoAvatarURL
	}
	// Done to follow clean line-of-sight
	found, fileName := false, ""
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if match, _ := path.Match(userId+"*", file.Name()); match {
			// Done to follow clean line-of-sight
			found = true
			fileName = file.Name()
			break
		}
	}
	if !found {
		return "", ErrNoAvatarURL
	}
	return "/avatars/" + fileName, nil
}
