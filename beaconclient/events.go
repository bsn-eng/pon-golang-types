package beaconclient

type PayloadAttributesEvent struct {
	Version string                     `json:"version"`
	Data *PayloadAttributesEventData `json:"data"`
}

type HeadEvent struct {
	Data *HeadEventData `json:"data"`
}