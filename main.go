package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("TOKEN_KEY"), os.Getenv("TOKEN_SECRET_KEY"))
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	tweet, _, err := client.Statuses.Update("Tweetando de um script feito em Go #nerd", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(tweet.Text)

}
