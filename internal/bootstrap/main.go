package bootstrap

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Pauloo27/logger"
	"github.com/Pauloo27/vcserver/internal/vcserver"
)

func Start() {
	loadEnv()
	loadTarget()

	if err := vcserver.Instance.Start(); err != nil {
		logger.Fatal(err)
	}

	stop := make(chan os.Signal, 1)
	//lint:ignore SA1016 i dont know, it just works lol
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-stop
	_ = vcserver.Instance.Stop()
}
