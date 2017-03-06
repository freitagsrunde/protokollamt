package handlers

import (
	"fmt"
	"log"
	"strings"

	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v2"
)

// LDAPService specifies what functionality
// is needed to communicate with and authenticate
// against configured LDAP service.
type LDAPService interface {
	GetServiceAddr() string
	GetServerName() string
	GetBindDN() string
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
func IndexLogin(ldapService LDAPService) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Connect to LDAP service configured in
		// protokollamt's config file.
		l, err := ldap.Dial("tcp", ldapService.GetServiceAddr())
		if err != nil {
			log.Printf("Connecting to configured LDAP service failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"reason": "internal issue",
			})
			c.Abort()
			return
		}
		defer l.Close()

		// Upgrade current unencrypted session with
		// LDAP service to TLS with StartTLS.
		err = l.StartTLS(&tls.Config{
			ServerName:         ldapService.GetServerName(),
			InsecureSkipVerify: false,
		})
		if err != nil {
			log.Printf("Upgrading LDAP connection to TLS failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"reason": "internal issue",
			})
			c.Abort()
			return
		}

		// Bind with name of user and password supplied
		// via validated login form values.
		err = l.Bind(fmt.Sprintf("uid=%s,%s", "USER", ldapService.GetBindDN()), "PASSWORD")
		if err != nil {

			// Check if user supplied invalid credentials.
			if strings.Contains(err.Error(), "Code 49 \"Invalid Credentials\"") {
				c.JSON(http.StatusBadRequest, gin.H{
					"reason": "you provided wrong credentials",
				})
				c.Abort()
				return
			}

			// If not, this is an internal error.
			log.Printf("Binding with supplied user name and password failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"reason": "internal issue",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
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
