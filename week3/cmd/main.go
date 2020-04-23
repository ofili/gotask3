package main

import (
	"encoding/json"
	"gotask3/week3/exporter"
	"gotask3/week3/facebook"
	"gotask3/week3/linkedin"
	"gotask3/week3/twitter"
)

func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)
	err := exporter.Export(twtr, "twtrdata.txt")
	err = exporter.Export(fb, "fbdata.txt")
	err = exporter.Export(lnkdin, "lkdata.txt")
	_, _ = json.Marshal(exporter.Export(fb, "fbdata.json"))
	_, _ = json.Marshal(exporter.Export(twtr, "twtrdata.json"))

	if err != nil {
		panic(err)
	}

}
