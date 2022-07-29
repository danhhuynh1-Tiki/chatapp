package main

import (
	"chat/pkg/cron"
	userDelivery "chat/services/domain/user/delivery/http"
	userUC "chat/services/domain/user/usecase"
	repository "chat/services/repository"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"

	"chat/pkg/config"
	authDelivery "chat/services/domain/auth/delivery/http"
	authUC "chat/services/domain/auth/usecase"

	roomDelivery "chat/services/domain/room/delivery/http"
	roomUC "chat/services/domain/room/usecase"

	messageDelivery "chat/services/domain/message/delivery/http"
	messageUC "chat/services/domain/message/usecase"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client
	//redisClient *redis.Client

	userRepository repository.UserRepository
	userUseCase    userUC.UserUseCase
	userHandler    userDelivery.UserHandler
	userRouter     userDelivery.UserRouter

	authRepository repository.AuthRepository
	authUseCase    authUC.AuthUseCase
	authHandler    authDelivery.AuthHandler
	authRouter     authDelivery.AuthRouter

	roomRepository repository.RoomRepository
	roomUseCase    roomUC.RoomUsecase
	roomHandler    roomDelivery.RoomHandler
	roomRouter     roomDelivery.RoomRouter

	messageRepository repository.MessageRepository
	messageUseCase    messageUC.MessageUseCase
	messageHandler    messageDelivery.MessageHandler
	messageRouter     messageDelivery.MessageRouter
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoClient, err = mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Connect to Redis
	//redisClient = redis.NewClient(&redis.Options{
	//	Addr: config.RedisUri,
	//})
	//
	//if _, err := redisClient.Ping(ctx).Result(); err != nil {
	//	panic(err)
	//}

	//err = redisClient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client connected successfully...")

	userCollection := mongoClient.Database("chat").Collection("users")
	roomCollection := mongoClient.Database("chat").Collection("room")
	messageCollection := mongoClient.Database("chat").Collection("message")

	userRepository = repository.NewUserRepository(ctx, userCollection)
	go cron.CronStatus(userRepository)
	userUseCase = userUC.NewUserUseCase(userRepository)
	userHandler = userDelivery.NewUserHandler(userUseCase)
	userRouter = userDelivery.NewUserRouter(userHandler)

	authRepository = repository.NewAuthRepository(ctx, userCollection)
	authUseCase = authUC.NewAuthUseCase(authRepository)
	authHandler = authDelivery.NewAuthHandler(authUseCase, userUseCase)
	authRouter = authDelivery.NewAuthRouter(authHandler)

	roomRepository = repository.NewRoomRepository(ctx, roomCollection)
	roomUseCase = roomUC.NewRoomUsecase(roomRepository, userRepository)
	roomHandler = roomDelivery.NewRoomHandler(roomUseCase)
	roomRouter = roomDelivery.NewRoomRouter(roomHandler)

	messageRepository = repository.NewMessageRepository(ctx, messageCollection)
	messageUseCase = messageUC.NewMessageUseCase(messageRepository, userRepository)
	messageHandler = messageDelivery.NewMessageHandler(messageUseCase)
	messageRouter = messageDelivery.NewRoomMessageRouter(messageHandler)
	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoClient.Disconnect(ctx)

	//value, err := redisClient.Get(ctx, "test").Result()

	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		panic(err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 300 * time.Hour,
	//}))
	//router.GET("/healthchecker", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	//})

	authRouter.AuthRoute(router, userUseCase)
	userRouter.UserRoute(router, userUseCase)
	roomRouter.RoomRoute(router, userUseCase)
	messageRouter.MessageRoute(router, userUseCase)
	log.Fatal(server.Run(":" + config.Port))
}
