package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/yants95/imersao-full-cycle-cartola-fc/internal/usecase"
	"github.com/yants95/imersao-full-cycle-cartola-fc/pkg/uow"
)

type ProcessNewPlayer struct{}

func (p ProcessNewPlayer) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.AddPlayerInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	addNewPlayerUseCase := usecase.NewAddPlayerUseCase(uow)
	err = addNewPlayerUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
