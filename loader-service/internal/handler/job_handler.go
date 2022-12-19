package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/tuxoo/smart-loader/loader-service/internal/client"
	"github.com/tuxoo/smart-loader/loader-service/internal/service"
)

const NEW_JOB = "job.new"

type JobHandler struct {
	client     *client.NatsClient
	jobService service.IJobStageService
}

func NewJobHandler(client *client.NatsClient, jobService service.IJobStageService) *JobHandler {
	return &JobHandler{
		client:     client,
		jobService: jobService,
	}
}

func (h *JobHandler) Handle() error {
	if _, err := h.client.Conn.Subscribe(NEW_JOB, func(msg *nats.Msg) {
		jobId, err := uuid.FromBytes(msg.Data)
		if err != nil {
			return
		}

		// TODO: context.Background()
		if err = h.jobService.ProcessStage(context.Background(), jobId); err != nil {
			return
		}
	}); err != nil {
		return err
	}

	return nil
}
