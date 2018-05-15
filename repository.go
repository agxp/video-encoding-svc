package main

import (
	pb "github.com/agxp/cloudflix/video-encoding-svc/proto"
	"github.com/minio/minio-go"
	"log"
	"os"
	"time"
)

type Repository interface {
	Encode(video_id string) (*pb.Response, error)
}

type EncodeRepository struct {
	s3 *minio.Client
}

func (repo *EncodeRepository) Encode(video_id string) (*pb.Response, error) {
	log.SetOutput(os.Stdout)
	var res *pb.Response
	log.Println("video_id: ", video_id)
	tmpPath := "/tmp/" + video_id
	tempVideo, err := repo.s3.FGetObject("videos", video_id, tmpPath)
	if err != nil {
		log.Fatalln(err)
		return nil, er
	}
	
	presignedURL, err := repo.s3.PresignedPutObject("videos", filename, time.Duration(1000)*time.Second)
	if err != nil {
		log.Fatalln(filename)
		return nil, err
	}

	res.PresignedUrl = presignedURL.String()

	return res, nil
}
