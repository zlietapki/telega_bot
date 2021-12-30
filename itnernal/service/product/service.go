package product

import (
	"errors"
	"strconv"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if idx > len(allProducts)-1 {
		return nil, errors.New("no such index " + strconv.Itoa(idx))
	}
	return &allProducts[idx], nil
}
