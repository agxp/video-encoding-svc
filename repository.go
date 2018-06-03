package main

import (
	"context"
	"database/sql"
	pb "github.com/agxp/cloudflix/video-encoding-svc/proto"
	"github.com/minio/minio-go"
	"github.com/opentracing/opentracing-go"
	"log"
	"github.com/go-redis/redis"
	"github.com/xfrr/goffmpeg/transcoder"
	"strings"
	"strconv"
)

type Repository interface {
	Encode(ctx context.Context, request *pb.Request) (*pb.Response, error)
}

type EncodeRepository struct {
	s3     *minio.Client
	pg     *sql.DB
	cache  *redis.Client
	tracer *opentracing.Tracer
}

func GetResolution(video_path string) string {
	return "1280x720"
}

func (repo *EncodeRepository) GenerateThumb(video_id string, video_path string)  {
	thumb_path := video_path + ".jpg"

	trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( video_path, thumb_path )
	// Handle error...
	if err != nil {
		log.Println(err)
	}

	trans.MediaFile().SetFrameRate(1)
	trans.MediaFile().SetSeekTimeInput("00:00:04")
	trans.MediaFile().SetDurationInput("00:00:01")

	// Start transcoder process
	done, err := trans.Run()
	if err != nil {
		log.Println(err)
	}

	progress, err := trans.Output()
	if err != nil {
		log.Println(err)
	}

	for msg := range progress {
		log.Println(msg)
	}

	// This channel is used to wait for the process to end
	<-done

	repo.s3.FPutObject("thumb", video_id + ".jpg", thumb_path, minio.PutObjectOptions{})

}

func (repo *EncodeRepository) Encode144p(video_id string, video_path string) {
	// encode 144p
	trans := new(transcoder.Transcoder)

	outputVideoPath := video_id + "_144.mp4"

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( video_path, "/tmp/" + outputVideoPath )
	// Handle error...
	if err != nil {
		log.Println(err)
	}
	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetPreset("veryfast")
	trans.MediaFile().SetResolution("256x144")
	trans.MediaFile().SetQuality(26)
	trans.MediaFile().SetFrameRate(30)

	// Start transcoder process
	done, err := trans.Run()
	if err != nil {
		log.Println(err)
	}

	progress, err := trans.Output()
	if err != nil {
		log.Println(err)
	}

	for msg := range progress {
		log.Println(msg)
	}

	// This channel is used to wait for the process to end
	<-done

	repo.s3.FPutObject("videos", video_id + "/" + outputVideoPath, "/tmp/" + outputVideoPath, minio.PutObjectOptions{})
}

func (repo *EncodeRepository) Encode240p(video_id string, video_path string) {
	// encode 240p
	trans := new(transcoder.Transcoder)

	outputVideoPath := video_id + "_240.mp4"

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( video_path, "/tmp/" + outputVideoPath )
	// Handle error...
	if err != nil {
		log.Println(err)
	}
	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetPreset("veryfast")
	trans.MediaFile().SetResolution("426x240")
	trans.MediaFile().SetQuality(26)
	trans.MediaFile().SetFrameRate(30)

	// Start transcoder process
	done, err := trans.Run()
	if err != nil {
		log.Println(err)
	}

	progress, err := trans.Output()
	if err != nil {
		log.Println(err)
	}

	for msg := range progress {
		log.Println(msg)
	}

	// This channel is used to wait for the process to end
	<-done

	repo.s3.FPutObject("videos", video_id + "/" + outputVideoPath, "/tmp/" + outputVideoPath, minio.PutObjectOptions{})

}

func (repo *EncodeRepository) Encode360p(video_id string, video_path string) {
	// encode 360p
	trans := new(transcoder.Transcoder)

	outputVideoPath := video_id + "_360.mp4"

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( video_path, "/tmp/" + outputVideoPath )
	// Handle error...
	if err != nil {
		log.Println(err)
	}
	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetPreset("veryfast")
	trans.MediaFile().SetResolution("640x360")
	trans.MediaFile().SetQuality(26)
	trans.MediaFile().SetFrameRate(30)

	// Start transcoder process
	done, err := trans.Run()
	if err != nil {
		log.Println(err)
	}

	progress, err := trans.Output()
	if err != nil {
		log.Println(err)
	}

	for msg := range progress {
		log.Println(msg)
	}

	// This channel is used to wait for the process to end
	<-done

	repo.s3.FPutObject("videos", video_id + "/" + outputVideoPath, "/tmp/" + outputVideoPath, minio.PutObjectOptions{})

}

func (repo *EncodeRepository) Encode480p(video_id string, video_path string) {
	// encode 480p
	trans := new(transcoder.Transcoder)

	outputVideoPath := video_id + "_480.mp4"

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( video_path, "/tmp/" + outputVideoPath)
	// Handle error...
	if err != nil {
		log.Println(err)
	}
	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetPreset("veryfast")
	trans.MediaFile().SetResolution("854x480")
	trans.MediaFile().SetQuality(26)
	trans.MediaFile().SetFrameRate(30)

	// Start transcoder process
	done, err := trans.Run()
	if err != nil {
		log.Println(err)
	}

	progress, err := trans.Output()
	if err != nil {
		log.Println(err)
	}

	for msg := range progress {
		log.Println(msg)
	}

	// This channel is used to wait for the process to end
	<-done

	repo.s3.FPutObject("videos", video_id + "/" + outputVideoPath, "/tmp/" + outputVideoPath, minio.PutObjectOptions{})

}

func (repo *EncodeRepository) Encode720p(video_id string, video_path string) {
	// encode 720
	trans := new(transcoder.Transcoder)

	outputVideoPath := video_id + "_720.mp4"

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( video_path, "/tmp/" + outputVideoPath )
	// Handle error...
	if err != nil {
		log.Println(err)
	}
	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetPreset("veryfast")
	trans.MediaFile().SetResolution("1280x720")
	trans.MediaFile().SetQuality(26)
	trans.MediaFile().SetFrameRate(30)

	// Start transcoder process
	done, err := trans.Run()
	if err != nil {
		log.Println(err)
	}

	progress, err := trans.Output()
	if err != nil {
		log.Println(err)
	}

	for msg := range progress {
		log.Println(msg)
	}

	// This channel is used to wait for the process to end
	<-done

	repo.s3.FPutObject("videos", video_id + "/" + outputVideoPath, "/tmp/" + outputVideoPath, minio.PutObjectOptions{})

}

func (repo *EncodeRepository) Encode1080p(video_id string, video_path string) {
	// encode 1080p
	trans := new(transcoder.Transcoder)

	outputVideoPath := video_id + "_1080.mp4"

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( video_path, "/tmp/" + outputVideoPath )
	// Handle error...
	if err != nil {
		log.Println(err)
	}

	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetPreset("veryfast")
	trans.MediaFile().SetResolution("1920x1080")
	trans.MediaFile().SetQuality(26)
	trans.MediaFile().SetFrameRate(30)

	// Start transcoder process
	done, err := trans.Run()
	if err != nil {
		log.Println(err)
	}

	progress, err := trans.Output()
	if err != nil {
		log.Println(err)
	}

	for msg := range progress {
		log.Println(msg)
	}

	// This channel is used to wait for the process to end
	<-done

	repo.s3.FPutObject("videos", video_id + "/" + outputVideoPath, "/tmp/" + outputVideoPath, minio.PutObjectOptions{})

}

func (repo *EncodeRepository) Encode(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	video_id := request.VideoId
	if video_id == "" {
		log.Print("invalid id")
		return nil, nil
	}
	sp, _ := opentracing.StartSpanFromContext(context.Background(), "Encode_Repo")

	sp.LogKV("video_id", video_id)

	defer sp.Finish()

	// if already encoded, exit
	// else:

	// get filepath
	psSP, _ := opentracing.StartSpanFromContext(context.Background(), "PG_EncodeGetFilePath", opentracing.ChildOf(sp.Context()))

	var file_path string

	selectQuery := `select file_path from videos where id=$1`
	err := repo.pg.QueryRow(selectQuery, video_id).Scan(&file_path)
	if err != nil {
		log.Print(err)
		psSP.Finish()
		return nil, err
	}
	psSP.Finish()

	sp.LogKV("file_path", file_path)

	video_path := "/tmp/" + video_id

	// pull file from S3
	err = repo.s3.FGetObject("videos", file_path, video_path, minio.GetObjectOptions{})
	if err != nil {
		log.Print(err)
		psSP.Finish()
		return nil, err
	}

	resolution := GetResolution(video_path)

	sp.LogKV("resolution", resolution)

	//var width float64
	var height float64

	if resolution != "" {
		resolution := strings.Split(resolution, "x")
		if len(resolution) != 0 {
			//width, _ = strconv.ParseFloat(resolution[0], 64)
			height, _ = strconv.ParseFloat(resolution[1], 64)
		}
	}

	sp.LogKV("height", height)

	repo.GenerateThumb(video_id, video_path)

	if height >= 144 {
		repo.Encode144p(video_id, video_path)
	}

	if height >= 240 {
		repo.Encode240p(video_id, video_path)
	}

	if height >= 360 {
		repo.Encode360p(video_id, video_path)
	}

	if height >= 480 {
		repo.Encode480p(video_id, video_path)
	}

	if height >= 720 {
		repo.Encode720p(video_id, video_path)
	}

	if height >= 1080 {
		repo.Encode1080p(video_id, video_path)
	}


	return &pb.Response{Filenames:nil}, nil
}
