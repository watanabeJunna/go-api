package main

import (
    "log"
    "fmt"
    "net"
    "os"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "user/pb"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "../lib/cipher"
)

type userService struct{}

var key = os.Getenv("AES_KEY")

func (e *userService) Auth(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
    username := req.Username
    password, err := cipher.Encrypt(key, req.Password)

    if err != nil {
        return nil, err
    }

    query := bson.M{ "username": username }

    user := &bson.M{}

    err = collection.FindOne(context.Background(), query).Decode(user)

    if err != nil {
        fmt.Println(err)
    }

    if len((*user)) == 0 {
        return &pb.AuthResponse{ Ok: false }, nil
    }

    dec, err := cipher.Decrypt(key, (*user)["password"].(string))

    if err != nil {
        return nil, err
    }

    ok := dec == password

    return &pb.AuthResponse{ Ok: ok }, nil
}

const (
    dbUrl = "mongodb://db:27017"
    dbName = "test"
    collectionName = "user"
)

var (
    client *mongo.Client
    collection *mongo.Collection
)

func init() {
    client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))

    if err != nil {
        panic(err)
    }

    if err = client.Connect(context.Background()); err != nil {
        panic(err)
    }
    
    collection = client.Database(dbName).Collection(collectionName)
}

func main() {
    defer client.Disconnect(context.Background())

    listen, err := net.Listen("tcp", ":9090")

    if err != nil {
        log.Fatalln(err)
        return
    }

    if err != nil {
        log.Fatalln(err)
        return
    }

    server := grpc.NewServer()
    pb.RegisterUserServiceServer(server, &userService{})
    reflection.Register(server)

    if err := server.Serve(listen); err != nil {
        log.Fatalln(err)
        return
    }
}