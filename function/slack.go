package function

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

// sendMessage to the Slack bot and room configured in stack.yml
func sendMessage(imgURL, eventType, cloudEvent string) {
	api := slack.New(getSlackToken())
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		ImageURL: imgURL,
		Color:    "#36a64f",
		Pretext:  fmt.Sprintf("Received CloudEvent Type: %s", eventType),
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Raw CloudEvent",
				Value: fmt.Sprintf("```%s```", cloudEvent),
				Short: false,
			},
		},
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage(getSlackRoom(), "", params)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

// getSlackToken returns the Slack Bot User OAuth Access Token
// from the configured secret file
func getSlackToken() string {
	return strings.TrimSpace(string(getSecret(os.Getenv("slack_token"))))
}

// getSecret returns the secret found at the first valid secret mountpoint
// endpoints currently support OpenFaaS on Kubernetes and Swarm
func getSecret(name string) []byte {
	mounts := []string{"/var/openfaas/secrets/", "/run/secrets/"}
	var b []byte
	var err error
	for _, m := range mounts {
		if b, err = ioutil.ReadFile(m + name); err == nil {
			return b
		}
	}
	log.Fatal(err)
	return nil
}

// getSlackRoom returns the Slack room ID stored in the functions
// environment variables
func getSlackRoom() string {
	return strings.TrimSpace(os.Getenv("slack_room"))
}
