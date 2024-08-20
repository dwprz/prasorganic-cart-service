package test

import (
	"context"
	"testing"

	"github.com/dwprz/prasorganic-cart-service/src/common/errors"
	"github.com/dwprz/prasorganic-cart-service/src/common/log"
	"github.com/dwprz/prasorganic-cart-service/src/infrastructure/database"
	"github.com/dwprz/prasorganic-cart-service/src/interface/repository"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/dwprz/prasorganic-cart-service/src/model/entity"
	repositoryimpl "github.com/dwprz/prasorganic-cart-service/src/repository"
	"github.com/dwprz/prasorganic-cart-service/test/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_DeleteItem$ -v ./src/repository/test/ -count=1

type DeleteItemTestSuite struct {
	suite.Suite
	cart        *entity.Cart
	cartRepo    repository.Cart
	postgresDB  *gorm.DB
	cartTestUtl *util.CartTest
}

func (c *DeleteItemTestSuite) SetupSuite() {
	c.postgresDB = database.NewPostgres()
	c.cartRepo = repositoryimpl.NewCart(c.postgresDB)

	c.cartTestUtl = util.NewCartTest(c.postgresDB)
}

func (c *DeleteItemTestSuite) SetupTest() {
	c.cart = c.cartTestUtl.Create()
}

func (c *DeleteItemTestSuite) TearDownTest() {
	c.cartTestUtl.Delete()
}

func (c *DeleteItemTestSuite) TearDownSuite() {
	sqlDB, err := c.postgresDB.DB()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.DeleteItemTestSuite/TearDownSuite", "section": "postgresDB.DB"}).Fatal(err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.DeleteItemTestSuite/TearDownSuite", "section": "sqlDB.Close"}).Fatal(err)
	}
}

func (c *DeleteItemTestSuite) Test_Success() {
	err := c.cartRepo.DeleteItem(context.Background(), &dto.DeleteItemCartReq{
		UserId:    c.cart.UserId,
		ProductId: c.cart.ProductId,
	})

	assert.NoError(c.T(), err)
}

func (c *DeleteItemTestSuite) Test_NotFound() {
	err := c.cartRepo.DeleteItem(context.Background(), &dto.DeleteItemCartReq{
		UserId:    "not-found",
		ProductId: 10,
	})
	assert.Error(c.T(), err)

	resErr, ok := err.(*errors.Response)
	assert.True(c.T(), ok)

	assert.Equal(c.T(), 404, resErr.HttpCode)
}

func TestRepository_DeleteItem(t *testing.T) {
	suite.Run(t, new(DeleteItemTestSuite))
}
