package hedera

import (
	"regexp"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var rstStream = regexp.MustCompile("(?i)\\brst[^0-9a-zA-Z]stream\\b") //nolint

type TopicMessageQuery struct {
	errorHandler      func(stat status.Status)
	completionHandler func()
	retryHandler      func(err error) bool
	attempt           uint64
	maxAttempts       uint64
	topicID           *TopicID
	startTime         *time.Time
	endTime           *time.Time
	limit             uint64
}

func NewTopicMessageQuery() *TopicMessageQuery {
	return &TopicMessageQuery{
		maxAttempts:       maxAttempts,
		errorHandler:      _DefaultErrorHandler,
		retryHandler:      _DefaultRetryHandler,
		completionHandler: _DefaultCompletionHandler,
	}
}

func (query *TopicMessageQuery) SetTopicID(topicID TopicID) *TopicMessageQuery {
	query.topicID = &topicID
	return query
}

func (query *TopicMessageQuery) GetTopicID() TopicID {
	if query.topicID == nil {
		return TopicID{}
	}

	return *query.topicID
}

func (query *TopicMessageQuery) SetStartTime(startTime time.Time) *TopicMessageQuery {
	query.startTime = &startTime
	return query
}

func (query *TopicMessageQuery) GetStartTime() time.Time {
	if query.startTime != nil {
		return *query.startTime
	}

	return time.Time{}
}

func (query *TopicMessageQuery) SetEndTime(endTime time.Time) *TopicMessageQuery {
	query.endTime = &endTime
	return query
}

func (query *TopicMessageQuery) GetEndTime() time.Time {
	if query.endTime != nil {
		return *query.endTime
	}

	return time.Time{}
}

func (query *TopicMessageQuery) SetLimit(limit uint64) *TopicMessageQuery {
	query.limit = limit
	return query
}

func (query *TopicMessageQuery) GetLimit() uint64 {
	return query.limit
}

func (query *TopicMessageQuery) SetMaxAttempts(maxAttempts uint64) *TopicMessageQuery {
	query.maxAttempts = maxAttempts
	return query
}

func (query *TopicMessageQuery) GetMaxAttempts() uint64 {
	return query.maxAttempts
}

func (query *TopicMessageQuery) SetErrorHandler(errorHandler func(stat status.Status)) *TopicMessageQuery {
	query.errorHandler = errorHandler
	return query
}

func (query *TopicMessageQuery) SetCompletionHandler(completionHandler func()) *TopicMessageQuery {
	query.completionHandler = completionHandler
	return query
}

func (query *TopicMessageQuery) SetRetryHandler(retryHandler func(err error) bool) *TopicMessageQuery {
	query.retryHandler = retryHandler
	return query
}

func (query *TopicMessageQuery) _ValidateNetworkOnIDs(client *Client) error {
	if client == nil || !client.autoValidateChecksums {
		return nil
	}

	if query.topicID != nil {
		if err := query.topicID.ValidateChecksum(client); err != nil {
			return err
		}
	}

	return nil
}

func (query *TopicMessageQuery) Subscribe(client *Client, onNext func(TopicMessage)) (SubscriptionHandle, error) {
	handle := SubscriptionHandle{}

	err := query._ValidateNetworkOnIDs(client)
	if err != nil {
		return SubscriptionHandle{}, err
	}

	return handle, nil
}

func _DefaultErrorHandler(stat status.Status) {
	println("Failed to subscribe to topic with status", stat.Code().String())
}

func _DefaultCompletionHandler() {
	println("Subscription to topic finished")
}

func _DefaultRetryHandler(err error) bool {
	code := status.Code(err)

	switch code {
	case codes.NotFound, codes.ResourceExhausted, codes.Unavailable:
		return true
	case codes.Internal:
		grpcErr, ok := status.FromError(err)

		if !ok {
			return false
		}

		return rstStream.FindIndex([]byte(grpcErr.Message())) != nil
	default:
		return false
	}
}
