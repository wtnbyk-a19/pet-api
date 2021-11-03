package injector

import (
	"../application/usecase"
	"../domain/repository"
	"../infrastructure/mysql"
	"../interfaces/controller"
	"../interfaces/gateway"
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
