package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("go-mssqldb", "dipdb")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	/*
		config := oauth1.NewConfig("SSeMQtmKAWgQNY6gV7N5YyeDk", "EmiVayn5Vq9HYAUvHBqupfeTJbiD3BbunkWfdbARmt5KQRw7dY")
		token := oauth1.NewToken("941730287175782400-flygbKv9LIREFui9uN7UqNudVGmDLij", "4uwjQ5Q5bYSE62L5ZOqb6istpl5wzfDCobzwc4Fy7NyZP")
		httpClient := config.Client(oauth1.NoContext, token)

		// Twitter client
		client := twitter.NewClient(httpClient)

		// Home Timeline
		//tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		//	Count: 60,
		//})
		// Convenience Demux demultiplexed stream messages
		demux := twitter.NewSwitchDemux()
		demux.Tweet = func(tweet *twitter.Tweet) {
			fmt.Println("Text: ")
			fmt.Println(tweet.Text)
			if tweet.Place!=nil {
				fmt.Println("CountryCode: ")
				_, err := fmt.Println(tweet.Place.CountryCode)
				if err!=nil	{
					fmt.Println(err)
				}

				fmt.Println("ID, Name, FullName, create: ")
				fmt.Println(tweet.Place.ID)
				fmt.Println(tweet.Place.Name)
				fmt.Println(tweet.Place.FullName)
				fmt.Println(tweet.CreatedAt)
			}


			fmt.Println("______________________________________")
		}
		demux.DM = func(dm *twitter.DirectMessage) {
			fmt.Println("ID: ")
			fmt.Println(dm.SenderID)
			fmt.Println("______________________________________")
		}
		demux.Event = func(event *twitter.Event) {
			fmt.Println("Event: ")
			fmt.Printf("%#v\n", event)
		}

		fmt.Println("Starting Stream...")

		// FILTER
		filterParams := &twitter.StreamFilterParams{
			Track:         []string{"я"},
			StallWarnings: twitter.Bool(true),
		}
		stream, err := client.Streams.Filter(filterParams)
		if err != nil {

			log.Fatal(err)
		}

		// Receive messages until stopped or stream quits
		go demux.HandleChan(stream.Messages)

		// Wait for SIGINT and SIGTERM (HIT CTRL-C)
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		log.Println(<-ch)

		fmt.Println("Stopping Stream...")
		stream.Stop()
	*/
}
