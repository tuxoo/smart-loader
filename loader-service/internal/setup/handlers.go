package setup

import (
	"github.com/tuxoo/smart-loader/loader-service/internal/client"
	"github.com/tuxoo/smart-loader/loader-service/internal/handler"
	"github.com/tuxoo/smart-loader/loader-service/internal/service"
)

func provideJobHandler(client *client.NatsClient, jobService service.IJobService) *handler.JobHandler {
	return handler.NewJobHandler(client, jobService)
}
