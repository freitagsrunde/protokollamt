package handlers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"crypto/tls"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/ldap.v2"
)

// SessionCreater specifies what functionality
// is needed to communicate with and authenticate
// against configured LDAP service, and for creating
// signed session objects.
type SessionCreater interface {
	GetJWTSigningSecret() string
	GetJWTValidFor() time.Duration
	GetLDAPServiceAddr() string
	GetLDAPServerName() string
	GetLDAPBindDN() string
}

// LoginPayload represents the values an user
// can supply to protokollamt in order to
// authenticate against configured LDAP.
type LoginPayload struct {
	Name     string `form:"login-name"`
	Password string `form:"login-password"`
}

// CreateSession produces a JSON Web Token (JWT) with
// authenticated claims based on supplied user values
// and saves it inside the session storage.
func CreateSession(name string, jwtSignSecret []byte, jwtValidFor time.Duration) (string, error) {

	// Save current timestamp.
	nowTime := time.Now()
	expTime := nowTime.Add(jwtValidFor)

	// Create a JWT with claims to identify user.
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss": name,
		"iat": nowTime.Unix(),
		"nbf": nowTime.Add((-1 * time.Minute)).Unix(),
		"exp": expTime.Unix(),
	})

	// Obtain the signed string.
	signedToken, err := token.SignedString(jwtSignSecret)
	if err != nil {
		return "", fmt.Errorf("failed to create signed JWT string: %v", err)
	}

	return signedToken, nil
}

// Index delivers the first page of protokollamt,
// a login form to authenticate via LDAP.
func Index() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"PageTitle": "Protokollamt der Freitagsrunde",
			"MainTitle": "Protokollamt",
		})
	}
}

// IndexLogin accepts user supplied LDAP credentials,
// asks the LDAP service to verify them, and creates
// a new session for respective user.
func IndexLogin(sessCreater SessionCreater) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Payload LoginPayload

		// Parse login form data into above defined payload.
		err := c.BindWith(&Payload, binding.FormPost)
		if err != nil {

			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"PageTitle":  "Protokollamt der Freitagsrunde",
				"MainTitle":  "Protokollamt",
				"FatalError": "Verarbeitungsfehler. Bitte erneut versuchen.",
			})
			c.Abort()
			return
		}

		Payload.Name = strings.TrimSpace(Payload.Name)

		// Connect to LDAP service configured in
		// protokollamt's config file.
		l, err := ldap.Dial("tcp", sessCreater.GetLDAPServiceAddr())
		if err != nil {

			log.Printf("Connecting to configured LDAP service failed: %v", err)

			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"PageTitle":  "Protokollamt der Freitagsrunde",
				"MainTitle":  "Protokollamt",
				"FatalError": "Verarbeitungsfehler. Bitte erneut versuchen.",
			})
			c.Abort()
			return
		}
		defer l.Close()

		// Upgrade current unencrypted session with
		// LDAP service to TLS with StartTLS.
		err = l.StartTLS(&tls.Config{
			ServerName:         sessCreater.GetLDAPServerName(),
			InsecureSkipVerify: false,
		})
		if err != nil {

			log.Printf("Upgrading LDAP connection to TLS failed: %v", err)

			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"PageTitle":  "Protokollamt der Freitagsrunde",
				"MainTitle":  "Protokollamt",
				"FatalError": "Verarbeitungsfehler. Bitte erneut versuchen.",
			})
			c.Abort()
			return
		}

		// Bind with name of user and password supplied
		// via validated login form values.
		err = l.Bind(fmt.Sprintf("uid=%s,%s", Payload.Name, sessCreater.GetLDAPBindDN()), Payload.Password)
		if err != nil {

			// Check if user supplied invalid credentials.
			if strings.Contains(err.Error(), "Code 49 \"Invalid Credentials\"") {

				c.HTML(http.StatusBadRequest, "index.html", gin.H{
					"PageTitle":  "Protokollamt der Freitagsrunde",
					"MainTitle":  "Protokollamt",
					"FatalError": "Benutzername oder Passwort falsch.",
				})
				c.Abort()
				return
			}

			// If not, this is an internal error.
			log.Printf("Binding with supplied user name and password failed: %v", err)

			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"PageTitle":  "Protokollamt der Freitagsrunde",
				"MainTitle":  "Protokollamt",
				"FatalError": "Verarbeitungsfehler. Bitte erneut versuchen.",
			})
			c.Abort()
			return
		}

		// Create a new session and obtain a signed JWT.
		token, err := CreateSession(Payload.Name, []byte(sessCreater.GetJWTSigningSecret()), sessCreater.GetJWTValidFor())
		if err != nil {

			log.Printf("JWT creation failure: %v", err)

			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"PageTitle":  "Protokollamt der Freitagsrunde",
				"MainTitle":  "Protokollamt",
				"FatalError": "Verarbeitungsfehler. Bitte erneut versuchen.",
			})
			c.Abort()
			return
		}

		// Insert cookie into client's headers.
		c.SetCookie("Protokollamt", token, int(sessCreater.GetJWTValidFor()), "", "", false, true)

		c.Redirect(http.StatusFound, "/protocols")
	}
}

// IndexLogout destroys the active session of
// requesting user, logging the user out.
func IndexLogout() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}
