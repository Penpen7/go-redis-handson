package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserWithCreatedAt struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type WrapUserWithCreatedAt UserWithCreatedAt

func (s UserWithCreatedAt) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(WrapUserWithCreatedAt(s)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *UserWithCreatedAt) UnmarshalBinary(data []byte) error {
	var wu WrapUserWithCreatedAt
	buf := bytes.NewReader(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&wu); err != nil {
		return err
	}
	*s = UserWithCreatedAt(wu)
	return nil
}

func workWithBinaryStruct(ctx context.Context, redisClient *redis.Client) {
	if err := redisClient.Set(ctx, "user", UserWithCreatedAt{
		Name:      "John Doe",
		Age:       12,
		CreatedAt: time.Now(),
	},
		0,
	).Err(); err != nil {
		panic(err)
	}

	var u UserWithCreatedAt
	if err := redisClient.Get(ctx, "user").Scan(&u); err != nil {
		if err == redis.Nil {
			fmt.Println("key 'name' does not exist")
			return
		}
		panic(err)
	}

	fmt.Println(u)
}

type User struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

func workWithStruct(ctx context.Context, redisClient *redis.Client) {
	// 書き込み
	if err := redisClient.HSet(ctx, "user", User{Name: "John Doe", Age: 12}).Err(); err != nil {
		panic(err)
	}
	var u User

	// 読み込み
	if err := redisClient.HGetAll(ctx, "user").Scan(&u); err != nil {
		if err == redis.Nil {
			fmt.Println("key 'name' does not exist")
			return
		}
		panic(err)
	}
	fmt.Println(u)
}

func workWithString(ctx context.Context, redisClient *redis.Client) {
	if err := redisClient.Set(ctx, "name", "John Doe", 0).Err(); err != nil {
		panic(err)
	}

	var name string
	if err := redisClient.Get(ctx, "name").Scan(&name); err != nil {
		if err == redis.Nil {
			fmt.Println("key 'name' does not exist")
			return
		}
		panic(err)
	}
	fmt.Println("name", name)

	if err := redisClient.Del(ctx, "name").Err(); err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	redisClient.FlushAll(ctx)
	workWithString(ctx, redisClient)
	workWithStruct(ctx, redisClient)
	workWithBinaryStruct(ctx, redisClient)
}
