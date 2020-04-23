package main

import (
	"gotask3/week3/exporter"
	"gotask3/week3/facebook"
	"gotask3/week3/linkedin"
	"gotask3/week3/twitter"
)

/*
func main() {
	runtime.GOMAXPROCS(2)
	godur, _ := time.ParseDuration("10ms")
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello ")
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutines")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

*/

func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)
	err := exporter.Export(twtr, "twtrdata.txt")
	err = exporter.Export(fb, "fbdata.txt")
	err = exporter.Export(lnkdin, "lkdata.txt")
	err = exporter.Export(fb, "fbdata.json")

	if err != nil {
		panic(err)
	}

	// twtr := new(twitter.Twitter)
	// lnkdin := new(linkedin.Linkedin)

	// ScrollFeeds(fb, twtr, lnkdin)
}

// ScrollFeeds prints all social media feeds
/*
func ScrollFeeds(platforms ...week3.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("=================================")
	}
}
*/
