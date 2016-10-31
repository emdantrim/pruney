package tweets

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
)

//GetTweets does the thing
func GetTweets(api *anaconda.TwitterApi) []anaconda.Tweet {
	result, err := api.GetSearch("emdantrim", nil)
	if err != nil {
		logrus.Error(err)
	} else {
		for _, tweet := range result.Statuses {
			logrus.WithFields(logrus.Fields{
				"username": tweet.User.ScreenName,
				"id":       tweet.Id,
				"tweet":    tweet.Text,
				"created":  tweet.CreatedAt,
				"retweets": tweet.RetweetCount,
				"favs":     tweet.FavoriteCount,
			}).Info("found tweet")
		}
	}
	return result.Statuses
}

//convert tweet time to Time object
//tweetTime, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweet.CreatedAt)

// so I think what makes most sense here is to make a whole database of all
// of the user's tweets, and their content. SQLite seems like the best thing to
// use because I am lazy and this is small-ish scale and time expense is of no
// importance.

// app should take (require) a db location.

//

//PruneTweets deletes tweets that nobody likes
func PruneTweets(api *anaconda.TwitterApi, tweets []anaconda.Tweet) {
	for _, tweet := range tweets {
		if tweet.FavoriteCount == 0 {
			// DANGER
			DeleteTweet(api, tweet, false)
			// DANGER
		}
	}

}

//DeleteTweet deletes a tweet.
func DeleteTweet(api *anaconda.TwitterApi, tweet anaconda.Tweet, actuallyDelete bool) {
	logrus.WithFields(logrus.Fields{"id": tweet.Id}).Debug("deleting tweet")
	if actuallyDelete {
		_, err := api.DeleteTweet(tweet.Id, false)
		if err != nil {
			logrus.Error(err)
		}
	} else {
		logrus.Info("didn't actually delete tweet")
	}
	logrus.WithFields(logrus.Fields{"id": tweet.Id}).Info("deleted tweet")
}
