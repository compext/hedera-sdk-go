package hedera

import (
	"time"
)

type TopicMessage struct {
	ConsensusTimestamp time.Time
	Contents           []byte
	RunningHash        []byte
	SequenceNumber     uint64
	Chunks             []TopicMessageChunk
	TransactionID      *TransactionID
}
