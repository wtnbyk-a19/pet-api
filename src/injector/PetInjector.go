package injector

import (
	"pet-api/src/application/usecase"
	"pet-api/src/domain/repository"
	"pet-api/src/infrastructure/database"
	"pet-api/src/interfaces/controller"
	"pet-api/src/interfaces/gateway"
)

func injectDB() database.DbConnection {
	dbConnection := database.NewDbConnection()
	return *dbConnection
}

func injectRepository() repository.IPetRepository {
	dbConnection := injectDB()
	return gateway.NewPetRepository(dbConnection)
}

func injectUsecase() usecase.IPetUsecase {
	petRepository := injectRepository()
	return usecase.NewPetUsecase(petRepository)
}

func InjectPetController() controller.PetController {
	return controller.NewPetController(injectUsecase())
}
