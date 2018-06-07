package function

import (
	"encoding/json"
	"log"
)

// OK string to be returned
const OK = "OK"

// Handle a serverless request
func Handle(req []byte) string {
	if resp := azureValidationEvent(req); resp != nil {
		return *resp
	}

	c, err := getCloudEvent(req)
	if err != nil {
		log.Fatal(err)
	}

	switch c.EventType {
	case MicrosoftStorageBlobCreatedType:
		d := MicrosoftStorageBlobCreated{}
		if err := json.Unmarshal(c.Data, &d); err != nil {
			log.Fatalf("Unable to unmarshal object for: %s", MicrosoftStorageBlobCreatedType)
		}
		sendMessage(d.Url, MicrosoftStorageBlobCreatedType, string(req))
		return OK
	default:
		log.Fatalf("Unsupported eventType received: %s", c.EventType)
	}
	return OK
}
