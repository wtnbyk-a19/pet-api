package injector

import (
	PetApp "pet-api/src/application/Pet"
	"pet-api/src/domain/Pet"
	"pet-api/src/infrastructure/database"
	"pet-api/src/interfaces/controller"
	"pet-api/src/interfaces/gateway"
)

func injectRepository() Pet.IPetRepository {
	return gateway.NewPetRepository(database.DbConn)
}

func injectUsecase() PetApp.IPetUsecase {
	petRepository := injectRepository()
	return PetApp.NewPetUsecase(petRepository)
}

func InjectPetController() controller.PetController {
	return controller.NewPetController(injectUsecase())
}
