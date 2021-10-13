package internal

type userService interface {
	GetAll() ([]Usuario, error)
	LoadFile() ([]Usuario, error)
	Store(nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error)
}
type service struct {
	repository userRepository
}

func NewService(r userRepository) userService {
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
