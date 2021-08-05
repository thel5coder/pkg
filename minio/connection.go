package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type IConnection interface {
	SetAccessKey(accessKey string) IConnection
	SetSecretKey(secretKey string) IConnection
	SetUseSsl(ssl bool) IConnection
	SetEndPoint(endPoint string) IConnection
	Connect() (err error)
	GetClient() *minio.Client
}

type Connection struct {
	accessKey string
	secretKey string
	useSSL    bool
	endPoint  string
	duration  int
	client    *minio.Client
}

func NewConnection() IConnection {
	return &Connection{}
}

func (c *Connection) SetAccessKey(accessKey string) IConnection {
	c.accessKey = accessKey

	return c
}

func (c *Connection) SetSecretKey(secretKey string) IConnection {
	c.secretKey = secretKey

	return c
}

func (c *Connection) SetUseSsl(ssl bool) IConnection {
	c.useSSL = ssl

	return c
}

func (c *Connection) SetEndPoint(endPoint string) IConnection {
	c.endPoint = endPoint

	return c
}

func (c *Connection) Connect() (err error) {
	c.client, err = minio.New(c.endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.accessKey, c.secretKey, ""),
		Secure: c.useSSL,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) GetClient() *minio.Client {
	return c.client
}
