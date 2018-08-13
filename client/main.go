package main

import (
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	"time"
	"os"
	"github.com/spilth/ogay-grpc"
)

func main() {
	englishWord := "Hello"
	if len(os.Args) > 1 {
		englishWord = os.Args[1]
	}

	connection, err := grpc.Dial("localhost:4200", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	translatorClient := ogay_grpc.NewTranslatorClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	translationRequest := &ogay_grpc.TranslationRequest{EnglishWord: englishWord}
	translationResponse, err := translatorClient.TranslateWord(ctx, translationRequest)
	if err != nil {
		log.Fatalf("could not translate: %v", err)
	}

	log.Printf("'%s' translates to '%s'", englishWord, translationResponse.PigLatinWord)
}
