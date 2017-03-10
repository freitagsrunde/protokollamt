package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Config contains all directives necessary to
// run a protokollamt application.
type Config struct {
	DeployStage string
	PublicAddr  string
	ListenAddr  string
	Database    Database
	JWT         JWT
	LDAP        LDAP
	Mail        Mail
}

// Database specifies connection details to the
// application's database.
type Database struct {
	Conn        *gorm.DB `toml:"-"`
	ConnString  string   `toml:"-"`
	ServiceAddr string
	User        string
	Password    string `toml:"-"`
	DBName      string
	SSLMode     string
}

// GetDBConn returns the active connection
// of protokollamt to configured database.
func (c *Config) GetDBConn() *gorm.DB {
	return c.Database.Conn
}

// JWT contains configuration values for use
// of JSON Web Tokens.
type JWT struct {
	SigningSecret string `toml:"-"`
	ValidFor      int
}

// GetJWTSignSecret returns the randomly generated
// JWT signing secret.
func (c *Config) GetJWTSigningSecret() string {
	return c.JWT.SigningSecret
}

// GetJWTSignSecret returns the randomly generated
// JWT signing secret.
func (c *Config) GetJWTValidFor() time.Duration {
	return (time.Duration(c.JWT.ValidFor) * time.Second)
}

// LDAP holds configuration for authorizing users
// by means of a LDAP infrastructure.
type LDAP struct {
	ServiceAddr string
	ServerName  string
	BindDN      string
}

// GetServiceAddr is a receiver to retrieve the
// config's LDAP service address.
func (c *Config) GetLDAPServiceAddr() string {
	return c.LDAP.ServiceAddr
}

// GetServerName is a receiver to retrieve the LDAP
// service's server name for connecting with TLS.
func (c *Config) GetLDAPServerName() string {
	return c.LDAP.ServerName
}

// GetBindDN is a receiver to retrieve defined part
// of request dinstinguished name, BindDN.
func (c *Config) GetLDAPBindDN() string {
	return c.LDAP.BindDN
}

// Mail defines how a connected mail server can
// be contacted for the purpose of sending protocols.
type Mail struct {
	ServiceAddr string
	User        string
	Password    string `toml:"-"`
}
