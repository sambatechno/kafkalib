package kevt

import (
	"github.com/sambatechno/kafkalib/gen/kafkalib/msg"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewUserEvent generates new user event struct with the default fields filled, so only `Body` needs to be filled.
func NewUserEvent() *msg.UserEvent {
	return &msg.UserEvent{
		CreateTimestamp: timestamppb.Now(),
	}
}
