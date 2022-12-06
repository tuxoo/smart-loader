package app

import (
	"fmt"
	"github.com/tuxoo/smart-loader/loader-service/internal/di"
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
}
