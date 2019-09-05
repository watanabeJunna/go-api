package main

import (
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "log"
    "net"
    pb "user/pb"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type userService struct{}

const (
    dbUrl = "mongodb://172.17.0.3:27017"
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

func (e *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
    user := &pb.User { Name: "WatanabeJunna" }

    _, err := collection.InsertOne(context.Background(), user)

    if err != nil {
        return nil, err
    }

    document, err := getUser()

    if err != nil {
        return nil, err
    }

    return &pb.User{ Name: (*document)["name"].(string) }, nil
}

func getUser() (*bson.M, error) {
    var document bson.M

    err := collection.FindOne(context.Background(), bson.D{}).Decode(&document)

    if err != nil {
        return nil, err
    }

    return &document, nil
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