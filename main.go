package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pruneymcprunetweets"
	app.Usage = "prunes yr unloved tweets"
	app.Author = "emdantrim"
	app.Action = goForthAndPrune

	var consumerKey, consumerSecret, userKey, userSecret string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "consumer-key",
			Usage:       "API Consumer Key",
			EnvVar:      "TWITTER_CONSUMER_KEY",
			Destination: &consumerKey,
		},
		cli.StringFlag{
			Name:        "consumer-secret",
			Usage:       "API Consumer Secret",
			EnvVar:      "TWITTER_CONSUMER_SECRET",
			Destination: &consumerSecret,
		},
		cli.StringFlag{
			Name:        "user-key",
			Usage:       "API User Key",
			EnvVar:      "TWITTER_USER_KEY",
			Destination: &userKey,
		},
		cli.StringFlag{
			Name:        "user-secret",
			Usage:       "API User Secret",
			EnvVar:      "TWITTER_USER_SECRET",
			Destination: &userSecret,
		},
	}

	app.Run(os.Args)
}

func goForthAndPrune(c *cli.Context) {
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.WithFields(logrus.Fields{"omg": "hi"}).Info("testy test")
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(userKey, userSecret)

	fmt.Println("test")
}
