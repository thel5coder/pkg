package postgresql

type Config struct {
	host                    string
	dbName                  string
	user                    string
	password                string
	port                    string
	location                string
	sslMode                 string
	dBMaxConnection         int
	dBMAxIdleConnection     int
	dBMaxLifeTimeConnection int
}

func NewConfig(host,dbName,user,password,port,location,sslMode string,maxConnection,maxIdleConnection,maxLifeTimeConnection int) *Config{
	return &Config{
		host:                    host,
		dbName:                  dbName,
		user:                    user,
		password:                password,
		port:                    port,
		location:                location,
		sslMode:                 sslMode,
		dBMaxConnection:         maxConnection,
		dBMAxIdleConnection:     maxIdleConnection,
		dBMaxLifeTimeConnection: maxLifeTimeConnection,
	}
}

func (c *Config) Host() string {
	return c.host
}

func (c *Config) DbName() string {
	return c.dbName
}

func (c *Config) User() string {
	return c.user
}

func (c *Config) Password() string {
	return c.password
}

func (c *Config) Port() string {
	return c.port
}

func (c *Config) Location() string {
	return c.location
}

func (c *Config) SslMode() string {
	return c.sslMode
}

func (c *Config) DBMaxConnection() int {
	return c.dBMaxConnection
}

func (c *Config) DBMAxIdleConnection() int {
	return c.dBMAxIdleConnection
}

func (c *Config) DBMaxLifeTimeConnection() int {
	return c.dBMaxLifeTimeConnection
}

