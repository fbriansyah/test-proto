package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	pb "github.com/fbriansyah/test-proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	Id    uint64 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	data = flag.String("data", "{\"name\": \"Febriansyah\", \"email\": \"febrian@mail.com\"}", "data user")
	mode = flag.String("mode", "get-all-user", "mode: get-user, create-user, get-all-user")
)

func route(mode string, ctx context.Context, c pb.UserApiClient) {
	switch mode {
	case "get-all-user":
		getAllUser(ctx, c)
	case "create-user":
		createUser(ctx, c)
	case "get-user":
		getUser(ctx, c)
	default:
		log.Fatalln("Unknown Mode:", mode)
	}
}

func getAllUser(ctx context.Context, c pb.UserApiClient) {
	r, err := c.GetAllUser(ctx, &pb.UsersRequest{})
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}
	users := r.GetUsers()
	for _, user := range users {
		if user.GetId() == 0 {
			continue
		}
		log.Println(user.GetId(), user.GetEmail(), user.GetName())
	}
}

func getUser(ctx context.Context, c pb.UserApiClient) {
	user := User{}
	err := json.Unmarshal([]byte(*data), &user)
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	r, err := c.GetUser(ctx, &pb.UserRequest{Id: user.Id, Email: user.Email})
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}
	log.Println(r.String())
}

func createUser(ctx context.Context, c pb.UserApiClient) {
	user := User{}
	err := json.Unmarshal([]byte(*data), &user)
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	r, err := c.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}
	log.Println(r.String())
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserApiClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	route(*mode, ctx, c)

}
