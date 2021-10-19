package section

import (
	"context"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"
)

// Interface that represents a Service
type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Section) (int, error)
	Update(ctx context.Context, s domain.Section, id int) error
	Delete(ctx context.Context, id int) error
}

// Structure representing a Service
type Service struct {
	repository Repository
}

// This creates a new Section Service
// Params: repository to get the repository associated with the service
// Returns a new Section service
func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

// Gets all the section stored in the database
// Params: context that has the information for the request
// Returns all sections in a JSON response
func (s *Service) GetAll(ctx context.Context) ([]domain.Section, error) {
	sections, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

// Gets a specific section based on the section_number
// Params: context that has the information for the request and the section id
// Returns specific section with the id passed as path params
func (s *Service) Get(ctx context.Context, id int) (domain.Section, error) {
	if exists := s.Exists(ctx, id); !exists {
		return domain.Section{}, web.NewErrorf(404, "section with section_number %d hasn't been found", id)
	}
	return s.repository.Get(ctx, id)
}

// Params: context that has the information for the request and the section id
func (s *Service) Exists(ctx context.Context, cid int) bool {
	return s.repository.Exists(ctx, cid)
}

// Inserts a new section into the database with all of its fields
// Params: context that has the information for the request and the section structure for the request
// Section has fields such as: SectionNumber, CurrentTemperature, MinTemperature, CurrentCapacity, MinCapacity, MaxCapacity, WarehouseID, ProductTypeID
func (s *Service) Save(ctx context.Context, sec domain.Section) (int, error) {
	if exists := s.Exists(ctx, sec.SectionNumber); exists {
		return 0, web.NewErrorf(409, "section with section_number %d already exists", sec.SectionNumber)
	}
	return s.repository.Save(ctx, sec)
}

// Updates all of the fields that a Section passed by path params has
// Params: context that has the information for the request and the section structure for the request
func (s *Service) Update(ctx context.Context, sec domain.Section, id int) error {
	if exists := s.Exists(ctx, id); !exists {
		return web.NewErrorf(404, "section with section_number %d doesn't exist", id)
	}
	if exists := s.Exists(ctx, sec.SectionNumber); exists {
		return web.NewErrorf(400, "new section number %d already exists", sec.SectionNumber)
	}
	return s.repository.Update(ctx, sec, id)
}

// Deletes a section held within the database
// Params: context that has the information for the request and the section id
func (s *Service) Delete(ctx context.Context, id int) error {
	if exists := s.Exists(ctx, id); !exists {
		return web.NewErrorf(404, "section with section_number %d doesn't exist", id)
	}
	return s.repository.Delete(ctx, id)
}
