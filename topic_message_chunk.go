package hedera

import (
	"time"
)

type TopicMessageChunk struct {
	ConsensusTimestamp time.Time
	ContentSize        uint64
	RunningHash        []byte
	SequenceNumber     uint64
}
