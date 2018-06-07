package function

import (
	"encoding/json"
	"log"
)

const MicrosoftStorageBlobCreatedType = "Microsoft.Storage.BlobCreated"

// SubscriptionValidationEvent received from Azure EventGrid
// when it tests webhooks during event subscription creation
type SubscriptionValidationEvent struct {
	Id              string
	Topic           string
	Subject         string
	Data            SubscriptionValidationData
	EventType       string
	EventTime       string
	MetadataVersion string
	DataVersion     string
}

// SubscriptionValidationData contains a ValidationCode that
// must be returned in a response, and a ValidationUrl that
// can be used to manually validate the subscription
type SubscriptionValidationData struct {
	ValidationCode string
	ValidationUrl  string
}

// SubscriptionValidationResp returned to EventGrid in order
// to validate the event subscription
type SubscriptionValidationResp struct {
	ValidationResponse string
}

// MicrosoftStorageBlobCreated event used to demonstrate
// consumption of CloudEvents, this is included in the
// CloudEvent Data field when a blob is added to an Azure
// Storage container
type MicrosoftStorageBlobCreated struct {
	Api                string
	ClientRequestId    string
	RequestId          string
	ETag               string
	ContentType        string
	ContentLength      int
	BlobType           string
	Url                string
	Sequencer          string
	StorageDiagnostics StorageDiagnostics
}

// StorageDiagnostics occasionally included by the Azure
// Storage service. This property should be ignored by
// event consumers.
type StorageDiagnostics struct {
	BatchId string
}

// azureValidationEvent handles a received Azure Subscription Validation
// event, returning the ValidationCode extracted from the request
func azureValidationEvent(req []byte) *string {
	v := []SubscriptionValidationEvent{}
	if err := json.Unmarshal(req, &v); err == nil {
		if len(v) > 0 && len(v[0].Data.ValidationCode) > 0 {
			r := SubscriptionValidationResp{}
			r.ValidationResponse = v[0].Data.ValidationCode
			if b, err := json.Marshal(r); err == nil {
				resp := string(b)
				return &resp
			}
			log.Fatalf("Unable to marshal Azure validation response")
		}
	}
	return nil
}
