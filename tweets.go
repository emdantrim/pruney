package pruneymcprunetweets

import (
	"fmt"
	"time"
)

func getTweets() {
	result, err := api.GetSearch("emdantrim", nil)
	if err == nil {
		for _, tweet := range result.Statuses {
			tweetTime, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweet.CreatedAt)
			if err == nil {
				fmt.Println(tweet.User.ScreenName, (time.Now().Sub(tweetTime)), tweet.Text, tweet.FavoriteCount, tweet.RetweetCount)
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
	}
}
