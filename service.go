package main

import (
	pb "github.com/agxp/cloudflix/video-encoding-svc/proto"
	"golang.org/x/net/context"
	"log"
	"os"
)

type service struct {
	repo Repository
}

func (srv *service) Encode(ctx context.Context, req *pb.Request, res *pb.Response) error {
	log.SetOutput(os.Stdout)

	filenames, err := srv.repo.Encode(req.video_id)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	res = filenames

	return nil
}
