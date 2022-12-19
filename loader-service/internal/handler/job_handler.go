package handler

import (
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/tuxoo/smart-loader/loader-service/internal/client"
	"github.com/tuxoo/smart-loader/loader-service/internal/service"
)

const NEW_JOB = "job.new"

type JobHandler struct {
	client     *client.NatsClient
	jobService service.IJobService
}

func NewJobHandler(client *client.NatsClient, jobService service.IJobService) *JobHandler {
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

		if err = h.jobService.ProcessJob(jobId); err != nil {
			return
		}
	}); err != nil {
		return err
	}

	return nil
}
