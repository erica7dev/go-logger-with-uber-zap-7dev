package main 

import (
	_"log"
	_"os"
	"go.uber.org/zap"
)

func main(){
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("Starting application")
	logger.Warn("Beware of the monsters")
	logger.Error("Something went wrong")
}