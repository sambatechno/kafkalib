package kafkalib

import "google.golang.org/protobuf/proto"

type ErrMissingTopic struct {
	MsgName string
}

func (e ErrMissingTopic) Error() string {
	return "missing topic for message " + e.MsgName
}

func NewErrMissingTopic(m proto.Message) ErrMissingTopic {
	return ErrMissingTopic{MsgName: string(m.ProtoReflect().Descriptor().FullName())}
}
