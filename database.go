package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go"
	"log"
	"os"
	"github.com/go-redis/redis"

)

func ConnectToS3() (*minio.Client, error) {
	endpoint := os.Getenv("MINIO_URL")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatal("failed to connect to minio", err)
	}

	minioClient.SetAppInfo("video-hosting-svc", "1.0.0")
	minioClient.TraceOn(nil)

	log.Printf("%#v\n", minioClient) // minioClient is now setup
	log.Printf("%#v\n", "hello, i just finished printing minioclient")

	buckets, err := minioClient.ListBuckets()
	if err != nil {
		fmt.Println("failed to list minio buckets", err)
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
	log.Printf("%#v\n", "hello, i just finished printing all the buckets")

	return minioClient, nil
}

func ConnectToPostgres() (*sql.DB, error) {
	PG_HOST := os.Getenv("POSTGRES_POSTGRESQL_SERVICE_HOST")
	PG_USER := os.Getenv("PG_USER")
	PG_PASSWORD := os.Getenv("PG_PASSWORD")
	DB_NAME := "videos"
	sslmode := "disable"

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		PG_HOST, PG_USER, PG_PASSWORD, DB_NAME, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping postgres", err)
		return nil, err
	}
	log.Print("successfully connected to postgres")

	return db, nil
}

func ConnectToRedis() (*redis.Client, error) {
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr: REDIS_HOST,
		Password: REDIS_PASSWORD,
		DB: 0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("failed to connect to redis", err)
		return nil, err
	}
	log.Print(pong, "successfully connected to redis")

	return client, nil
}
