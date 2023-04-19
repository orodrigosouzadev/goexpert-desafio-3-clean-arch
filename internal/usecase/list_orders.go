package usecase

import "github.com/orodrigosouzadev/goexpert-desafio-3-clean-arch/internal/entity"

type OrderDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type ListOrdersOutputDTO struct {
	Orders []*OrderDTO
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() (*ListOrdersOutputDTO, error) {
	dto := &ListOrdersOutputDTO{}
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return dto, err
	}
	for _, order := range orders {
		dto.Orders = append(dto.Orders, &OrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return dto, nil
}
