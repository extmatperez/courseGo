package employee

import (
	"context"
	"fmt"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"
)

type Service interface {
	GetAll() ([]domain.Employee, error)
	Store(card_number_id, first_name, last_name string, warehouse_id int) (domain.Employee, error)
	Delete(card_number_id string) error
	Get(card_number_id string) (domain.Employee, error)
	Update(card_number_id, first_name, last_name string, warehouse_id int) (domain.Employee, error)
}

type service struct {
	repository Repository
}

//NewService
//create Service
//@return Service
func NewService(r Repository) Service {
	return &service{repository: r}
}

//GetAll
//get all employees
//@return []domain.Employee, error
func (s *service) GetAll() ([]domain.Employee, error) {
	var ctx context.Context = context.Background()
	employees, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	//validate no repeat ids
	err = validateCardNumberId(employees)
	if err != nil {
		return nil, err
	}

	return employees, err
}

//validateCardNumberId
//validate that there aren't repeat ids in employees slice
//@return error
func validateCardNumberId(employees []domain.Employee) error {

	idMap := make(map[string]domain.Employee)

	if len(employees) == 0 {
		return nil
	}

	for _, employee := range employees {
		if user, exist := idMap[employee.CardNumberID]; exist {
			return web.NewErrorf(404, "el usuario: %s y el usuario: %s tienen el mismo identificador: %s", user.FirstName, employee.FirstName, employee.CardNumberID)

		}
		idMap[employee.CardNumberID] = employee
	}

	return nil
}

//Store
//create employee
//@return domain.Employee, error
func (s *service) Store(card_number_id, first_name, last_name string, warehouse_id int) (domain.Employee, error) {
	var ctx context.Context = context.Background()

	if exist := s.repository.Exists(ctx, card_number_id); exist {
		return domain.Employee{}, fmt.Errorf("el usuario con identificador: %s ya existe", card_number_id)
	}
	employee := domain.Employee{CardNumberID: card_number_id, FirstName: first_name, LastName: last_name, WarehouseID: warehouse_id}

	id, err := s.repository.Save(ctx, employee)

	if err != nil {
		return domain.Employee{}, err
	}

	employee_validate, err := s.Get(card_number_id)

	if err != nil {
		return domain.Employee{}, err
	}

	if employee_validate.FirstName != employee.FirstName ||
		employee_validate.LastName != employee.LastName ||
		employee_validate.WarehouseID != employee.WarehouseID {
		return domain.Employee{}, fmt.Errorf("el usuario no ha sido actualizado correctamente")
	}

	employee.ID = id

	return employee, nil
}

//Get
//get one employees
//@return domain.Employee, error
func (s *service) Get(card_number_id string) (domain.Employee, error) {
	var ctx context.Context = context.Background()
	if exist := s.repository.Exists(ctx, card_number_id); !exist {
		return domain.Employee{}, fmt.Errorf("el usuario con identificador: %s no existe", card_number_id)
	}
	employee, err := s.repository.Get(ctx, card_number_id)
	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

//Delete
//delete one employee by card_number_id
//@return error
func (s *service) Delete(card_number_id string) error {
	var ctx context.Context = context.Background()
	if exist := s.repository.Exists(ctx, card_number_id); !exist {
		return fmt.Errorf("el usuario con identificador: %s no existe", card_number_id)
	}
	//validate the user isnt used in other table
	if userIsUsed := isUsed(card_number_id); userIsUsed {
		return fmt.Errorf("el usuario con identificador: %s esta en uso", card_number_id)
	}

	if err := s.repository.Delete(ctx, card_number_id); err != nil {
		return err
	}

	return nil
}

//Update
//update one employee
//@return domain.Employee, error
func (s *service) Update(card_number_id, first_name, last_name string, warehouse_id int) (domain.Employee, error) {
	var ctx context.Context = context.Background()
	if exist := s.repository.Exists(ctx, card_number_id); !exist {
		return domain.Employee{}, fmt.Errorf("el usuario con identificador: %s no existe", card_number_id)
	}
	employee := domain.Employee{CardNumberID: card_number_id, FirstName: first_name, LastName: last_name, WarehouseID: warehouse_id}

	err := s.repository.Update(ctx, employee)

	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

//isUsed
//validate the user isnt used in other table
//@return bool
func isUsed(card_number_id string) bool {
	return false
}
