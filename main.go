package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	//"golang.org/x/net/context"
	pruney "github.com/emdantrim/pruney/pruney"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "pruneymcprunetweets"
	app.Usage = "prunes yr unloved tweets"
	app.Author = "emdantrim"

	var consumerKey, consumerSecret, userKey, userSecret, dbFile string
	var actuallyPrune bool

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.Info("setting up pruney")

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
		cli.StringFlag{
			Name:        "db-file",
			Usage:       "SQLite3 database file location",
			EnvVar:      "PRUNEY_DB_FILE",
			Destination: &dbFile,
		},
		cli.BoolFlag{
			Name:        "actually-prune",
			Usage:       "actually prune tweets (opposite of dry-run)",
			EnvVar:      "PRUNEY_ACTUALLY_PRUNE",
			Destination: &actuallyPrune,
		},
	}

	app.Action = func(c *cli.Context) {
		goForthAndPrune(consumerKey, consumerSecret, userKey, userSecret)
	}

	app.Run(os.Args)
}

func goForthAndPrune(consumerKey, consumerSecret, userKey, userSecret string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(userKey, userSecret)

	pruney.PruneTweets(api, pruney.GetTweets(api))
}
