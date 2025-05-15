package kafkalib

import (
	"github.com/sambatechno/kafkalib/gen/kafkalib/opts"
	"google.golang.org/protobuf/proto"
)

func getTopicName(m proto.Message) (string, bool) {
	desc := m.ProtoReflect().Descriptor()
	msgOpts := desc.Options()

	if proto.HasExtension(msgOpts, opts.E_TopicName) {
		ext := proto.GetExtension(msgOpts, opts.E_TopicName)
		return ext.(string), true
	}
	return "", false
}
