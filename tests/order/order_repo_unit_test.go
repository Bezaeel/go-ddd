package order_test

import (
	"go-ddd/src/order"
	"go-ddd/tests"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type InitRepoTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func TestOrderRepoUnitTest(t *testing.T) {
	suite.Run(t, &InitRepoTestSuite{})
}

func (its *InitRepoTestSuite) SetupSuite() {
	_db := tests.SetupTestDB()
	db, err := tests.ConvertToGormDB(_db)
	if err != nil {
		its.T().Log("Error converting to GORM DB", err)
	}

	its.db = db
}

func (its *InitRepoTestSuite) TestCreateOrder() {
	repo := order.NewOrderRepository(its.db)

	newOrder := order.Order{
		Id:             1,
		ShipmentNumber: 1,
		CargoId:        200,
		IsShipped:      false,
		CreatedAt:      time.Now().UTC(),
	}
	expectedOrder, err := repo.CreateOrder(newOrder)
	if err != nil {
		its.FailNowf("Failed to create user", err.Error())
	}

	var actualOrder order.Order
	its.db.First(&actualOrder, expectedOrder.Id)
	its.Require().Equal(newOrder, actualOrder)
}

func (its *InitRepoTestSuite) TeardownTestSuite() {
	its.T().Log("tearing down...")
	db, _ := its.db.DB()
	tests.TearDown(db)
}
