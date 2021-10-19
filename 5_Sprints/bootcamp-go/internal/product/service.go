package product

import (
	"context"
	"errors"
	"net/http"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Get(id int) (domain.Product, error)
	Save(p domain.Product) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)
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

//Save: recibe un producto y lo guarda si el productCode no existe. Devuelve el producto creado o un error
func (s *service) Save(p domain.Product) (domain.Product, error) {
	//Si existe el producto se devuelve un error
	if s.repository.Exists(context.Background(), p.ProductCode) {
		return domain.Product{}, web.NewError(http.StatusConflict, "el produto ya existe")
	}

	//Se guarda el producto
	id, err := s.repository.Save(context.Background(), p)

	//si hay error, se devuelve
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}

	//se asigna el id
	p.ID = id

	//se retorna el producto creado
	return p, nil

}

//GetAll: busca y devuelve lista de productos que no tengan ProductCode repetido o un error
func (s *service) GetAll() ([]domain.Product, error) {

	//Se obtienen todos los productos
	products, err := s.repository.GetAll(context.Background())

	if err != nil {
		return nil, errors.New(err.Error())
	}
	//Se filtran para obtener productos con ProductCode unico
	elementMap := make(map[string]domain.Product)
	for _, product := range products {
		elementMap[product.ProductCode] = product
	}

	var productsFilter []domain.Product
	for _, value := range elementMap {
		productsFilter = append(productsFilter, value)
	}

	//se retorna la lista filtrada
	return productsFilter, nil
}

//Get: recibe el id de un producto, lo obtiene y lo devuelve o retorna error
func (s *service) Get(id int) (domain.Product, error) {
	//se obtiene un producto por su id
	product, err := s.repository.Get(context.Background(), id)

	//si hay error, se devuelve el error
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}

	//se retorna el producto buscado
	return product, nil
}

//Update: recibe el producto actualizado, lo actualiza completo y lo devuelve, o un error
func (s *service) Update(p domain.Product) (domain.Product, error) {

	//se busca el producto con el id
	product, err := s.Get(p.ID)
	//Si no existe el producto se devuelve un error
	if err != nil {
		return domain.Product{}, web.NewError(http.StatusBadRequest, "el producto no existe")
	}

	//si el ProductCode es distinto se verifica que no exista otro producto con el ProductCode a actuializar
	if p.ProductCode != product.ProductCode {
		if s.repository.Exists(context.Background(), p.ProductCode) {
			//si existe devuelve error
			return domain.Product{}, web.NewError(http.StatusBadRequest, "el ProductCode ingresado ya existe")
		}
	}
	//se actualiza el producto
	err = s.repository.Update(context.Background(), p)

	//si hay error se devuelve
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}

	//se retorna el producto
	return p, nil
}

//Delete: recibe el id del producto y lo elimina, si hay error lo devuelve
func (s *service) Delete(id int) error {

	//busca el producto por id
	_, err := s.Get(id)

	//si hay error lo retorna
	if err != nil {
		return err
	}

	//elimina el producto por su id
	err = s.repository.Delete(context.Background(), id)

	return err

}
