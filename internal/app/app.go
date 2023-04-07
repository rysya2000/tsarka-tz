package app

import (
	"rest/config"
	"rest/internal/usecase"	
	"rest/internal/delivery"
)

func Run(cfg *config.Config) {

	usecase := usecase.New()

	handler := delivery.New(usecase)


//  HTTP
	

}
