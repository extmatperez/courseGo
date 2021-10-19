package buyer

import (
	"context"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"fmt"
	"log"
)

type Service interface {
	Create(cardNumberID, firstName, lastName string) (domain.Buyer, error)
	GetAll() ([]domain.Buyer, error)
	Get(cardNumberID string) (domain.Buyer, error)
	Delete(cardNumberID string) error
	Update(cardNumberID, firstName, lastName string) (domain.Buyer, error)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

type service struct {
	repository Repository
}

// Given a card_number_id, first_name and last_name, creates a new buyer with those values
// If a buyer with card_number_id already exists, returns an error
func (s *service) Create(cardNumberID, firstName, lastName string) (domain.Buyer, error) {
	
	if exists, _ := s.Exists(cardNumberID); exists {
		return domain.Buyer{}, fmt.Errorf("buyer for card_number_id %s already exists", cardNumberID)
	}

	buyer := domain.Buyer{
		CardNumberID: cardNumberID,
		FirstName: firstName,
		LastName: lastName,
	}

	if buyerID, err := s.repository.Save(context.Background(), buyer); err != nil {
		log.Println("ERROR", err)
		return domain.Buyer{}, nil
	} else {
		buyer.ID = buyerID
		return buyer, nil
	}
}

// Returns the list of all the buyers
func (s *service) GetAll() ([]domain.Buyer, error) {

	buyersList, err := s.repository.GetAll(context.Background())
	
	if err != nil {
		log.Println("ERROR", err)
		return make([]domain.Buyer, 0), nil
	}
	//If buyers dont have unique card_number_id, return an empty list and log error
	if !AreCardNumberIDsUnique(buyersList) {
		log.Println("ERROR", fmt.Errorf("buyers card_number_id's are not unique"))
		return make([]domain.Buyer, 0), nil
	}

	if len(buyersList) == 0 {
		buyersList = make([]domain.Buyer, 0)
	}

	return buyersList, nil
}

// Given a card_number_id, finds the buyer for that card_number_id
// If the buyer does not exist return error
func (s *service) Get(cardNumberID string) (domain.Buyer, error) {

	if exists, _ := s.Exists(cardNumberID); !exists {
		return domain.Buyer{}, fmt.Errorf("buyer for card_number_id %s does not exist", cardNumberID)
	}

	buyer, err := s.repository.Get(context.Background(), cardNumberID)

	if err != nil {
		log.Println("ERROR", err)
		return domain.Buyer{}, nil
	}

	// If buyer result does not match card_number_id searched, return empty buyer and log error
	if buyer.CardNumberID != cardNumberID {
		log.Println("ERROR", fmt.Errorf("searched for buyer with card_number_id %s but got %s", cardNumberID, buyer.CardNumberID))
		return domain.Buyer{}, nil
	}

	return buyer, nil
}

// Given a card_number_id, deletes the buyer for that card_number_id
// If the buyer does not exist return error
func (s *service) Delete(cardNumberID string) error {
	
	if exists, _ := s.Exists(cardNumberID); !exists {
		return fmt.Errorf("buyer for card_number_id %s does not exist", cardNumberID)
	}

	if err := s.repository.Delete(context.Background(), cardNumberID); err != nil {
		log.Println("ERROR", err)
		return nil
	}

	// If buyer with card_number_id was not deleted, log error
	if exists, _ := s.Exists(cardNumberID); exists {
		log.Println("ERROR", fmt.Errorf("buyer for card_number_id %s exists after being deleted", cardNumberID))
	}

	return nil
}

// Given a card_number_id, first_name and last_name, update the buyer for that card_number_id with the values first_name and last_name
// If the buyer does not exist return error
func (s *service) Update(cardNumberID, firstName, lastName string) (domain.Buyer, error) {
	
	if exists, _ := s.Exists(cardNumberID); !exists {
		return domain.Buyer{}, fmt.Errorf("buyer for card_number_id %s does not exist", cardNumberID)
	}

	buyer := domain.Buyer{
		CardNumberID: cardNumberID,
		FirstName: firstName,
		LastName: lastName,
	}
	
	if err := s.repository.Update(context.Background(), buyer); err != nil {
		log.Println("ERROR", err)
		return domain.Buyer{}, nil
	}
	
	// If updated buyer was not updated properly, return empty buyer and log error
	if updatedBuyer, err := s.repository.Get(context.Background(), cardNumberID); err != nil {
		log.Println("ERROR", err)
		return domain.Buyer{}, nil
	} else {
		return updatedBuyer, nil
	}
}

// Given a card_number_id, return true if there exists a buyer with that card_number_id
func (s *service) Exists(cardNumberID string) (bool, error) {

	return s.repository.Exists(context.Background(), cardNumberID), nil
}

// Given a list of buyers, return true if all buyers have different card_number_id's
func AreCardNumberIDsUnique(buyers []domain.Buyer) (bool) {

	uniqueCardNumberIDs := make(map[string]bool)
	// Iterate the buyers list and store in map the card_number_id as key.
	// If at some point a buyer has a card_number_id that already exists in map it means 2 buyers share same card_number_id
	for _, buyer := range buyers {
	
		if _, exists := uniqueCardNumberIDs[buyer.CardNumberID]; exists {
			return false
		}
		uniqueCardNumberIDs[buyer.CardNumberID] = true
	}

	return true
}