package buyer

import (
	"context"
	"fmt"

	"github.com/extmatperez/courseGo/5_Sprints/meli_bootcamp/internal/domain"
)

type Service interface {
	GetTodos() ([]domain.Buyer, error)
	Get(cardNumberID string) (domain.Buyer, error)
	Crear(cardNumberID, firstName, lastName string) (domain.Buyer, error)
	Actualizar(cardNumberID, firstName, lastName string) (domain.Buyer, error)
	Eliminar(cardNumberID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTodos() ([]domain.Buyer, error) {
	ps, err := s.repository.GetAll(context.Background())
	if err != nil {
		return nil, err
	}

	if len(ps) == 0 {
		return []domain.Buyer{}, fmt.Errorf("No hay buyers cargados")
	}

	return ps, nil
}

func (s *service) Get(cardNumberID string) (domain.Buyer, error) {
	ps, err := s.repository.Get(context.Background(), cardNumberID)
	if err != nil {
		return domain.Buyer{}, err
	}

	if ps.CardNumberID != cardNumberID {
		return domain.Buyer{}, fmt.Errorf("El buyer no existe")
	}

	return ps, nil
}

func (s *service) Crear(cardNumberID, firstName, lastName string) (domain.Buyer, error) {
	existe := s.repository.Exists(context.Background(), cardNumberID)

	if existe {
		return domain.Buyer{}, fmt.Errorf("El buyer %v ya existe", cardNumberID)
	}

	buyer := domain.Buyer{CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName}
	buyerId, err := s.repository.Save(context.Background(), buyer)

	if err != nil {
		return domain.Buyer{}, fmt.Errorf("El buyer no se pudo cargar")
	} else {
		buyer.ID = buyerId
	}
	return buyer, nil
}

func (s *service) Actualizar(cardNumberID, firstName, lastName string) (domain.Buyer, error) {
	existe := s.repository.Exists(context.Background(), cardNumberID)
	if !existe {
		return domain.Buyer{}, fmt.Errorf("El buyer %s no existe", cardNumberID)
	}

	buyer := domain.Buyer{
		CardNumberID: cardNumberID,
		FirstName:    firstName,
		LastName:     lastName,
	}

	if err := s.repository.Update(context.Background(), buyer); err != nil {
		return domain.Buyer{}, nil
	}
	updatedBuyer, err := s.repository.Get(context.Background(), cardNumberID)
	if err != nil {
		return domain.Buyer{}, nil
	} else {
		return updatedBuyer, nil
	}
}

func (s *service) Eliminar(cardNumberID string) error {

	existe := s.repository.Exists(context.Background(), cardNumberID)
	if !existe {
		return fmt.Errorf("El buyer %s no existe", cardNumberID)
	}
	err := s.repository.Delete(context.Background(), cardNumberID)
	if err != nil {
		return fmt.Errorf("El buyer %s no se pudo eliminar", cardNumberID)
	}

	// If buyer with card_number_id was not deleted, log error
	existe = s.repository.Exists(context.Background(), cardNumberID)
	if existe {
		return fmt.Errorf("buyer for card_number_id %s exists after being deleted", cardNumberID)
	}

	return nil
}
