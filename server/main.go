package main

import (
	"github.com/spilth/ogay"
	"github.com/spilth/ogay-grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type translatorServer struct{}

func (*translatorServer) TranslateWord(context context.Context, translationRequest *ogay_grpc.TranslationRequest) (*ogay_grpc.TranslationResponse, error) {
	pigLatinWord := ogay.TranslateWord(translationRequest.EnglishWord)

	return &ogay_grpc.TranslationResponse{PigLatinWord: pigLatinWord}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	ogay_grpc.RegisterTranslatorServer(grpcServer, &translatorServer{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
