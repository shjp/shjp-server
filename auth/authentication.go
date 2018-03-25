package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("very-secret"))

type LoginHandler struct {
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user")
	if err != nil {
		fmt.Printf("Cannot get session: %s", err)
		return
	}

	fmt.Printf("saving to session, foo = %s\n", session.Values["foo"])
	session.Values["foo"] = "bar"
	session.Save(r, w)
}
