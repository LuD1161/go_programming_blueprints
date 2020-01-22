package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
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
	var email string
	var ok bool
	if email, ok = c.userData["email"].(string); !ok {
		return "", ErrNoAvatarURL
	}
	m := md5.New()
	io.WriteString(m, strings.ToLower(email))
	return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
}
