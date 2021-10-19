package warehouse

import (
	"context"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w domain.Warehouse) (int, error)
	Update(ctx context.Context, w domain.Warehouse) error
	Delete(ctx context.Context, id int) error
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) Exists(ctx context.Context, warehouseCode string) bool {
	return s.repository.Exists(ctx, warehouseCode)
}

func (s *Service) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	return s.repository.Save(ctx, w)
}

func (s *Service) Update(ctx context.Context, w domain.Warehouse) error {
	return s.repository.Update(ctx, w)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
