package main

import (
	pb "github.com/agxp/cloudflix/video-encoding-svc/proto"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"github.com/micro/protobuf/proto"
)

type service struct {
	repo   Repository
	tracer *opentracing.Tracer
	logger *zap.Logger
}


func (srv *service) Encoder(ctx context.Context, req *pb.Request, res *pb.Response) error {
	sp, _ := opentracing.StartSpanFromContext(ctx, "Encoder_Service")

	logger.Info("Request for Encode_Service received")
	defer sp.Finish()

	rsp, err := srv.repo.Encode(sp.Context(), req.VideoId)
	if err != nil {
		logger.Error("failed Encode", zap.Error(err))
		return err
	}

	data, err := proto.Marshal(rsp)
	if err != nil {
		logger.Error("marshal error", zap.Error(err))
	}

	err = proto.Unmarshal(data, res)
	if err != nil {
		logger.Error("unmarshal error", zap.Error(err))
		return err
	}

	return nil
}