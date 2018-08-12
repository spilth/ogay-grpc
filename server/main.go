package main

import (
	"google.golang.org/grpc"
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	"github.com/spilth/ogay"
	"github.com/spilth/ogay-grpc"
)

type translatorServer struct{}

func (*translatorServer) TranslateWord(context context.Context, translationRequest *ogay_grpc.TranslationRequest) (translationResponse *ogay_grpc.TranslationResponse, err error) {
	pigLatinWord := ogay.TranslateWord(translationRequest.EnglishWord)

	return &ogay_grpc.TranslationResponse{PigLatinWord: pigLatinWord},nil
}

func main() {
	listener, err := net.Listen("tcp", ":4200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	ogay_grpc.RegisterTranslatorServer(grpcServer, &translatorServer{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
