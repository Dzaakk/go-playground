package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis ping failed: %v", err)
	}
}

func createUser(ctx context.Context, key string, user User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, key, userJSON, 0).Err()
	return err
}

func readUser(ctx context.Context, key string) (User, error) {
	userJSON, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return User{}, nil
	} else if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal([]byte(userJSON), &user)
	return user, err
}

func updateUser(ctx context.Context, key string, user User) error {
	exists, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return fmt.Errorf("key %s does not exist", key)
	}

	return createUser(ctx, key, user)
}

func deleteUser(ctx context.Context, key string) error {
	err := rdb.Del(ctx, key).Err()
	return err
}

func main() {
	ctx := context.Background()
	user := User{Name: "Alice", Age: 30}
	key := "user:1"

	err := createUser(ctx, key, user)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	readUserResult, err := readUser(ctx, key)
	if err != nil {
		log.Fatalf("Failed to read user: %v", err)
	}

	fmt.Println("Read User: ", readUserResult)

	updatedUser := User{Name: "Alice Blue", Age: 31}
	err = updateUser(ctx, key, updatedUser)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	updatedUserResult, err := readUser(ctx, key)
	if err != nil {
		log.Fatalf("Failed to read updated user: %v", err)
	}

	fmt.Println("Updated User: ", updatedUserResult)

	err = deleteUser(ctx, key)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}

	deletedUserResult, err := readUser(ctx, key)
	if err != nil {
		log.Fatalf("Failed to read deleted user: %v", err)
	}

	fmt.Println("Deleted User: ", deletedUserResult)
}
