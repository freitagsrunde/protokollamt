package handlers

import (
	"log"
	"strings"

	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v2"
)

type LDAPService interface {
	GetServiceAddr() string
	GetServerName() string
	GetBindDN() string
}

func Index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func IndexLogin(ldapService LDAPService) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Connect to LDAP service configured in
		// by protokollamt config file.
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
		err = l.Bind("uid=USER,ou=users,dc=freitagsrunde,dc=org", "PASSWORD")
		if err != nil {

			// Check if user supplied invalid credentials.
			if strings.Contains(err.Error(), "Code 49 \"Invalid Credentials\"") {
				c.JSON(http.StatusBadRequest, gin.H{
					"reason": "you are wrong",
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

func IndexLogout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}
