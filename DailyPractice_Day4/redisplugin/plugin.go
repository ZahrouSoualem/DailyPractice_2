package redisplug

// docker run -d -p 6379:6379 --name redis-db redis

type Client struct {
	address string
}

func NewClient() *Client {
	return &Client{
		address: "127.0.0.1",
	}
}

func initClient() *Client {

	return NewClient()
}
