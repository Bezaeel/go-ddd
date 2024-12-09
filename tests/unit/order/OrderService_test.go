package unit_order_tests

import (
	"go-ddd/src/modules/order"
	"go-ddd/tests/unit/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrderServiceShould struct {
	suite.Suite
	repoMock *mocks.OrderRepository
	sut      order.OrderService
}

func TestOrderServiceShouldTestSuite(t *testing.T) {
	suite.Run(t, &OrderServiceShould{})
}

func (uts *OrderServiceShould) SetupTest() {
	repoMock := mocks.OrderRepository{}
	uts.repoMock = &repoMock

	sut := order.NewOrderService(uts.repoMock)
	uts.sut = sut

}

func (uts *OrderServiceShould) Test_CreateOrder_ShouldCreateOrder_WhenCommandIsValid() {
	// arrange
	orderFixture := order.OrderEntity{
		Id:             900,
		CargoId:        201,
		ShipmentNumber: 1001,
		IsShipped:      false,
	}

	orderCommandFixture := &order.CreateOrderCommand{
		CargoId:        201,
		ShipmentNumber: 1001,
	}

	toOrder := order.MapToOrder(*orderCommandFixture)

	uts.repoMock.On("CreateOrder", toOrder).Return(&orderFixture, nil)

	// act
	expected := uts.sut.CreateOrder(*orderCommandFixture)

	// assert
	uts.repoMock.AssertNumberOfCalls(uts.T(), "CreateOrder", 1)
	uts.Require().Equal(orderFixture.Id, expected.Value.OrderId)
}
