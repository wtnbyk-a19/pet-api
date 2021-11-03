package injector

import (
	"pet-api/src/application/usecase"
	"pet-api/src/domain/repository"
	"pet-api/src/infrastructure/mysql"
	"pet-api/src/interfaces/controller"
	"pet-api/src/interfaces/gateway"
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
