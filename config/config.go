package config

// Config contains all directives necessary to
// run a protokollamt application.
type Config struct {
	DeployStage string
	PublicAddr  string
	ListenAddr  string
	Database    Database
	LDAP        LDAP
	Mail        Mail
}

// Database specifies connection details to the
// application's database.
type Database struct {
	ServiceAddr string
	User        string
	Password    string `toml:"-"`
	DBName      string
	SSLMode     string
}

// LDAP holds configuration for authorizing users
// by means of a LDAP infrastructure.
type LDAP struct {
	ServiceAddr string
	ServerName  string
	BindDN      string
}

// Mail defines how a connected mail server can
// be contacted for the purpose of sending protocols.
type Mail struct {
	ServiceAddr string
	User        string
	Password    string `toml:"-"`
}

// Define a receiver to retrieve the config's
// value for connecting to the LDAP service.
func (c *Config) GetServiceAddr() string {
	return c.LDAP.ServiceAddr
}

// Define a receiver to retrieve the LDAP service's
// server name for connecting with TLS.
func (c *Config) GetServerName() string {
	return c.LDAP.ServerName
}

// Define a receiver to retrieve defined part
// of request dinstinguished name, BindDN.
func (c *Config) GetBindDN() string {
	return c.LDAP.BindDN
}
