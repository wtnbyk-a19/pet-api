package injector

import (
	"pet-api/src/application/usecase"
	"pet-api/src/domain/repository"
	"pet-api/src/infrastructure/database"
	"pet-api/src/interfaces/controller"
	"pet-api/src/interfaces/gateway"
)

func injectRepository() repository.IPetRepository {
	return gateway.NewPetRepository(database.DbConn)
}

func injectUsecase() usecase.IPetUsecase {
	petRepository := injectRepository()
	return usecase.NewPetUsecase(petRepository)
}

func InjectPetController() controller.PetController {
	return controller.NewPetController(injectUsecase())
}
