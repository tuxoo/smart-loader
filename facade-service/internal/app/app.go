package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
)

func Run(configPath string) {
	fmt.Println(`
 ####~~##~~~#~~####~~#####~~######~~~~##~~~~~~####~~~####~~#####~~#####~#####
##~~~~~###~##~##~~##~##~~##~~~##~~~~~~##~~~~~##~~##~##~~##~##~~##~##~~~~##~~##
 ####~~##~#~#~######~#####~~~~##~~~~~~##~~~~~##~~##~######~##~~##~####~~#####
~~~~##~##~~~#~##~~##~##~~##~~~##~~~~~~##~~~~~##~~##~##~~##~##~~##~##~~~~##~~##
 ####~~##~~~#~##~~##~##~~##~~~##~~~~~~######~~####~~##~~##~#####~~#####~##~~##
	`)

	_, err := config.InitConfig(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	logrus.Println("Hello")
}
