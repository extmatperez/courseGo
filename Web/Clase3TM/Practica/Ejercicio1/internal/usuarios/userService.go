package internal

type Service interface {
	GetAll() ([]Usuario, error)
	LoadFile() ([]Usuario, error)
	Store(nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error)
	UpdateNombre(id int, nombre string) (Usuario, error)
	Delete(id int) error
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Usuario, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}
func (s *service) LoadFile() ([]Usuario, error) {
	ps, err := s.repository.LoadFile()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Usuario{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, activo, fechaDeCreacion)
	if err != nil {
		return Usuario{}, err
	}

	return producto, nil
}
func (s *service) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion)
}

func (s *service) UpdateNombre(id int, nombre string) (Usuario, error) {
	return s.repository.UpdateNombre(id, nombre)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
