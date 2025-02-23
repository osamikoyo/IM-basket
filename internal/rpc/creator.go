package rpc

import (
	"context"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/osamikoyo/IM-basket/internal/config"
	"github.com/osamikoyo/IM-basket/internal/data"
	"github.com/osamikoyo/IM-basket/internal/data/models"
	"github.com/osamikoyo/IM-basket/pkg/loger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RpcServer struct{
	amqpChanel *amqp.Channel
	storage *data.Storage
	loger loger.Logger
	cfg *config.Config
}

func New(cfg *config.Config) (*RpcServer, error) {
	storage, err := data.New(cfg)
	if err != nil{
		return nil, err
	}

	conn, err := amqp.Dial(cfg.AmqpUrl)
	if err != nil{
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		return nil, err
	}

	return &RpcServer{
		storage: storage,
		cfg: cfg,
		amqpChanel: ch,
		loger: loger.New(),
	}, nil
}

func (s *RpcServer) Listen(ch chan error) {
	q, err := s.amqpChanel.QueueDeclare(
		s.cfg.RpcQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		ch <- err
	}

	err = s.amqpChanel.Qos(
		1,
		0,
		false,
	)
	if err != nil{
		ch <- err
	}

	msgs, err := s.amqpChanel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		ch <- err
	}

	var forever chan struct{}

	go func ()  {
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		for d := range msgs{
			var req models.Req

			if err = sonic.Unmarshal(d.Body, &req);err != nil{
				s.loger.Error().Err(err)

				bodys, _ := sonic.Marshal(models.NewResp(err, http.StatusInternalServerError)) 

				s.amqpChanel.PublishWithContext(ctx,
					"",
					d.ReplyTo,
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						CorrelationId: d.CorrelationId,
						Body: bodys,
					},	
				)
			}
			
			err = s.storage.Create(req.ID)
			if err != nil{
				s.loger.Error().Err(err)

				bodys, _ := sonic.Marshal(models.NewResp(err, http.StatusInternalServerError)) 

				s.amqpChanel.PublishWithContext(ctx,
					"",
					d.ReplyTo,
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						CorrelationId: d.CorrelationId,
						Body: bodys,
					},
				)
			}

			resp, _ := sonic.Marshal(models.NewResp(nil, http.StatusOK))

			err = s.amqpChanel.PublishWithContext(ctx,
				"",
				d.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					CorrelationId: d.CorrelationId,
					Body: resp,
				})
		}
	}()
	s.loger.Info().Msg("RPC server listening...")
	<- forever
}