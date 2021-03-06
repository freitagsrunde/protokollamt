package middleware

import (
	"fmt"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTSigningSecretGetter defines the one method
// needed to obtain initially set JWT signing secret.
type JWTSigningSecretGetter interface {
	GetJWTSigningSecret() string
}

// CookieRemoveRedirect deletes this application's
// authentication cookie and redirects to index page.
func CookieRemoveRedirect(c *gin.Context) {

	// Instruct client to invalidate set cookie.
	c.SetCookie("Protokollamt", "", -1, "", "", false, true)

	// Redirect to index page.
	c.Redirect(http.StatusFound, "/")

	// Tell gin not to proceed in call stack.
	c.Abort()
}

// NotAuthorized ensures that only clients with no
// 'Protokollamt' cookie present continue to the index
// page again. This allows to directly redirect to
// restricted pages if a correctly named cookie is
// found. Its validity is checked on a restricted page.
func NotAuthorized() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Check if a cookie of name 'Protokollamt' is
		// present in the client's request headers.
		_, err := c.Request.Cookie("Protokollamt")
		if err == nil {
			c.Redirect(http.StatusFound, "/protocols")
			c.Abort()
			return
		}
	}
}

// Authorized is run as a middleware in every callstack
// of an HTTP service endpoint that requires authentication.
// It looks for, parses, and validates the JWT supplied
// by requesting client.
func Authorized(jwtGetter JWTSigningSecretGetter) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Extract cookie containing JWT from request.
		cookie, err := c.Request.Cookie("Protokollamt")
		if err != nil {
			CookieRemoveRedirect(c)
			return
		}

		// Parse authorization token.
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {

			// Verify that JWT was signed with correct algorithm.
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if ok != true {
				return nil, fmt.Errorf("Wrong signature algorithm: %v", token.Header["alg"])
			}

			return []byte(jwtGetter.GetJWTSigningSecret()), nil
		})

		// Check for parsing errors.
		if err != nil {
			CookieRemoveRedirect(c)
			return
		}

		// Check if JWT is valid.
		if token.Valid != true {
			CookieRemoveRedirect(c)
			return
		}

		// Obtain claims from token.
		claims := token.Claims.(jwt.MapClaims)
		err = claims.Valid()
		if err != nil {
			CookieRemoveRedirect(c)
			return
		}

		c.Next()
	}
}
