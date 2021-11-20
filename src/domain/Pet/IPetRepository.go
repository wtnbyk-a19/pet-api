package Pet

type IPetRepository interface {
	Create(pet *Pet) (error error)
}
