package fullsort

import "time"

type Response struct {
	Swaps		uint32			`json:"swaps,omitempty"`
	Accesses 	uint32			`json:"accesses,omitempty"`
	Duration	time.Duration	`json:"duration,omitempty"`
}