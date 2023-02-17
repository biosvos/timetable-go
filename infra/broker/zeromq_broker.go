package broker

import (
	"encoding/json"
	"fmt"
	"github.com/pebbe/zmq4"
	"github.com/pkg/errors"
	"timetable-go/adapter"
	"timetable-go/domain"
)

type ZeroMessageQueueBroker struct {
	ctx    *zmq4.Context
	socket *zmq4.Socket
}

func NewZeroMessageQueueBroker(port int) (adapter.Broker, func(), error) {
	ctx, err := zmq4.NewContext()
	const errMessage = "failed to new zeromq"
	if err != nil {
		return nil, nil, errors.Wrap(err, errMessage)
	}
	socket, err := ctx.NewSocket(zmq4.PUB)
	if err != nil {
		return nil, nil, errors.Wrap(err, errMessage)
	}

	err = socket.Bind(fmt.Sprintf("tcp://*:%v", port))
	if err != nil {
		return nil, nil, errors.Wrap(err, errMessage)
	}
	return &ZeroMessageQueueBroker{
			ctx:    ctx,
			socket: socket,
		}, func() {
			_ = socket.Close()
		},
		nil
}

type body struct {
	Id    string  `json:"id"`
	Start string  `json:"start"`
	End   *string `json:"end,omitempty"`
	Memo  string  `json:"memo,omitempty"`
}

func (z *ZeroMessageQueueBroker) Publish(topic adapter.Topic, record *domain.TimeRecord) error {
	marshal, err := json.Marshal(&body{
		Id:    record.Id(),
		Start: record.StartString(),
		End:   record.EndString(),
		Memo:  record.Memo(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to publish")
	}

	_, err = z.socket.SendMessage(topic, marshal)
	if err != nil {
		return errors.Wrap(err, "failed to publish")
	}
	return nil
}
