package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

type IConnection interface {
	Connect() IConnection

	StoreWithExpired(key string, val interface{}, duration string) error

	Store(key string, val interface{}) error

	GetData(key string, data interface{}) error

	Remove(key string) error
}

type Connection struct {
	address  string
	password string
	db       int
	client   *redis.Client
}

func NewConnection(address, password string) IConnection {
	return &Connection{
		address:  address,
		password: password,
		db:       0,
	}
}

func (c *Connection) Connect() IConnection {
	redisOption := &redis.Options{
		Addr:     c.address,
		Password: c.password,
	}
	c.client = redis.NewClient(redisOption)

	return c
}

func (c *Connection) StoreWithExpired(key string, val interface{}, duration string) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = c.client.Set(key, string(b), dur).Err()

	return err
}

func (c *Connection) Store(key string, val interface{}) error {
	b, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = c.client.Set(key, string(b), 0).Err()

	return err
}

func (c *Connection) GetData(key string, data interface{}) error {
	res, err := c.client.Get(key).Result()
	if err != nil {
		return err
	}

	if res == "" {
		return errors.New("[Redis] Value of " + key + " is empty.")
	}

	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return err
	}

	return err
}

func (c *Connection) Remove(key string) error {
	return c.client.Del(key).Err()
}
