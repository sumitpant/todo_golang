package logger;

import (
	"log"
	"go.uber.org/zap"
)


func Logs (info string){

	logger ,err:=zap.NewProduction();
	if err != nil{
		log.Fatal(err)
	}

	logger.Sugar().Info(info)

}