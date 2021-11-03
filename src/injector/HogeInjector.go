package injector

import (
	"docker-go-api/src/application/usecase"
	"docker-go-api/src/domain/repository"
	"docker-go-api/src/infrastructure/mysql"
	"docker-go-api/src/interfaces/controller"
	"docker-go-api/src/interfaces/gateway"
)

func injectDB() mysql.DbConnection {
	dbConnection := mysql.NewDbConnection()
	return *dbConnection
}

func injectRepository() repository.IHogeRepository {
	dbConnection := injectDB()
	return gateway.NewHogeRepository(dbConnection)
}

func injectUsecase() usecase.IHogeUsecase {
	hogeRepository := injectRepository()
	return usecase.NewHogeUsecase(hogeRepository)
}

func InjectHogeController() controller.HogeController {
	return controller.NewHogeController(injectUsecase())
}
