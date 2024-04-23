package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

const SESSION_ID = "1111"

func newCookieStore() *sessions.CookieStore {
	authKey := []byte("12345")
	encryptionKey := []byte("secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}
func main() {
	store := newCookieStore()

	e := echo.New()

	e.Use(echo.WrapMiddleware(context.ClearHandler))

	e.GET("/set", func(c echo.Context) error {
		session, err := store.Get(c.Request(), SESSION_ID)
		if err != nil {
			return err
		}

		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"

		// Simpan sesi
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			return err
		}

		// Buat cookie dari sesi
		cookie := &http.Cookie{
			Name:  SESSION_ID,
			Value: session.ID,
			Path:  "/",
		}

		// Set cookie ke respons HTTP
		http.SetCookie(c.Response().Writer, cookie)

		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.GET("/get", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)

		if len(session.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}

		return c.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})

	e.GET("/delete", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.Start(":8000")
}
