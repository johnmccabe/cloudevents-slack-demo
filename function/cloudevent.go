package function

import (
	"encoding/json"
)

// CloudEvent v0.1
// https://github.com/cloudevents/spec/blob/v0.1/json-format.md
type CloudEvent struct {
	EventType          string
	EventTypeVersion   string
	CloudEventsVersion string
	Source             string
	EventID            string
	EventTime          string
	ContentType        string
	Extensions         map[string]string
	Data               json.RawMessage
}

// getCloudEvent returns a pointer to a CloudEvent extracted from the
// request submitted to the handler
func getCloudEvent(req []byte) (*CloudEvent, error) {
	c := CloudEvent{}
	if err := json.Unmarshal(req, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
