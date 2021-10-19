package seller

import (
	"context"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
)

type IService interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, cid int) (domain.Seller, error)
	ExistsCid(ctx context.Context, cid int) bool
	ExistsId(ctx context.Context, id int) bool
	Save(ctx context.Context, seller domain.SellerToSave) (int, error)
	Update(ctx context.Context, seller domain.Seller) error
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

func (s *Service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) GetByCid(ctx context.Context, cid int) (domain.Seller, error) {
	return s.repository.GetByCid(ctx, cid)
}

func (s *Service) GetById(ctx context.Context, id int) (domain.Seller, error) {
	return s.repository.GetById(ctx, id)
}

func (s *Service) ExistsCid(ctx context.Context, cid int) bool {
	return s.repository.ExistsCid(ctx, cid)
}

func (s *Service) ExistsId(ctx context.Context, id int) bool {
	return s.repository.ExistsId(ctx, id)
}

func (s *Service) Save(ctx context.Context, sellerToSave domain.SellerToSave) (int, error) {
	return s.repository.Save(ctx, sellerToSave)
}

func (s *Service) Update(ctx context.Context, seller domain.Seller) error {
	return s.repository.Update(ctx, seller)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
