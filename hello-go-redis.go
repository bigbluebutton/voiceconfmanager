package main

import "fmt"
import "flag"
import "github.com/garyburd/redigo/redis"
import "sync"

var (
	redisPool      *redis.Pool
	redisAddress   = flag.String("redis-address", ":6379", "Address to the Redis server")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

func publish(channel, value interface{}) {
	/*
		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			panic(err)
		}
		defer c.Close()
		c.Do("PUBLISH", channel, value)
	*/
	conn := redisPool.Get()
	defer conn.Close()
	conn.Do("PUBLISH", channel, value)
}

func main() {
	flag.Parse()

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)

		if err != nil {
			return nil, err
		}

		return c, err
	}, *maxConnections)

	defer redisPool.Close()

	/*
		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			panic(err)
		}
		defer c.Close()
	*/

	var wg sync.WaitGroup
	wg.Add(2)

	//	psc := redis.PubSubConn{Conn: c}
	psc := redis.PubSubConn{Conn: redisPool.Get()}

	// This goroutine receives and prints pushed notifications from the server.
	// The goroutine exits when the connection is unsubscribed from all
	// channels or there is an error.
	go func() {
		defer wg.Done()
		for {
			switch n := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("Message: %s %s\n", n.Channel, n.Data)
			case redis.PMessage:
				fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
			case redis.Subscription:
				fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
				if n.Count == 0 {
					return
				}
			case error:
				fmt.Printf("error: %v\n", n)
				return
			}
		}
	}()

	// This goroutine manages subscriptions for the connection.
	go func() {
		defer wg.Done()

		psc.Subscribe("example")
		psc.PSubscribe("p*")
		psc.PSubscribe("bigbluebutton:to-bbb-apps:system")

		// The following function calls publish a message using another
		// connection to the Redis server.
		publish("example", "hello")
		publish("example", "world")
		publish("pexample", "foo")
		publish("pexample", "bar")

		// Unsubscribe from all connections. This will cause the receiving
		// goroutine to exit.
		psc.Unsubscribe()
		//psc.PUnsubscribe()
	}()

	wg.Wait()

	// Output:
	// Subscription: subscribe example 1
	// Subscription: psubscribe p* 2
	// Message: example hello
	// Message: example world
	// PMessage: p* pexample foo
	// PMessage: p* pexample bar
	// Subscription: unsubscribe example 1
	// Subscription: punsubscribe p* 0
}

func main3() {
	//INIT OMIT
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	//set
	c.Do("SET", "message1", "Hello World")

	//get
	world, err := redis.String(c.Do("GET", "message1"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
	//ENDINIT OMIT

	psc := redis.PubSubConn{c}
	psc.PSubscribe("bigbluebutton:to-bbb-apps:system")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.PMessage:
			fmt.Printf("PMessage: %s %s %s\n", v.Pattern, v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Printf("error: %v\n", v)
		}
	}
}
