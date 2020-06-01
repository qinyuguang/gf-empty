package book

var Ins IModel = new(model)

type IModel interface {
	Create(*Entity) error
	GetByID(int64) (*Entity, error)
}

type model struct{}

func (s *model) Create(entity *Entity) error {
	_, err := entity.Insert()
	return err
}

func (s *model) GetByID(id int64) (*Entity, error) {
	return FindOne("id=?", id)
}
