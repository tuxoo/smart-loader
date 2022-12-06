package app

import (
	"fmt"
	"github.com/tuxoo/smart-loader/facade-service/internal/di"
)

func Run() {
	fmt.Println(`
  ####~~##~~~#~~####~~#####~~######~~~~##~~~~~~####~~~####~~#####~~#####~#####
 ##~~~~~###~##~##~~##~##~~##~~~##~~~~~~##~~~~~##~~##~##~~##~##~~##~##~~~~##~~##
  ####~~##~#~#~######~#####~~~~##~~~~~~##~~~~~##~~##~######~##~~##~####~~#####
 ~~~~##~##~~~#~##~~##~##~~##~~~##~~~~~~##~~~~~##~~##~##~~##~##~~##~##~~~~##~~##
  ####~~##~~~#~##~~##~##~~##~~~##~~~~~~######~~####~~##~~##~#####~~#####~##~~##
	`)

	di.App.Run()

	//nc, err := nats.Connect("nats://host.docker.internal:4222")
	//if err != nil {
	//	logrus.Fatalf("error initializing nats: %s", err.Error())
	//}

	//err = nc.Publish("foo", []byte("Hello World"))

	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//<-quit
}
