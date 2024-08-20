package test

import (
	"context"
	"testing"

	"github.com/dwprz/prasorganic-cart-service/src/common/log"
	"github.com/dwprz/prasorganic-cart-service/src/infrastructure/database"
	"github.com/dwprz/prasorganic-cart-service/src/interface/repository"
	"github.com/dwprz/prasorganic-cart-service/src/model/entity"
	repositoryimpl "github.com/dwprz/prasorganic-cart-service/src/repository"
	"github.com/dwprz/prasorganic-cart-service/test/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_CountByUserId$ -v ./src/repository/test/ -count=1

type CountByUserIdTestSuite struct {
	suite.Suite
	cart        *entity.Cart
	cartRepo    repository.Cart
	postgresDB  *gorm.DB
	cartTestUtl *util.CartTest
}

func (c *CountByUserIdTestSuite) SetupSuite() {
	c.postgresDB = database.NewPostgres()
	c.cartRepo = repositoryimpl.NewCart(c.postgresDB)

	c.cartTestUtl = util.NewCartTest(c.postgresDB)

	c.cart = c.cartTestUtl.Create()
}

func (c *CountByUserIdTestSuite) TearDownSuite() {
	c.cartTestUtl.Delete()

	sqlDB, err := c.postgresDB.DB()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.CountByUserIdTestSuite/TearDownSuite", "section": "postgresDB.DB"}).Fatal(err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.CountByUserIdTestSuite/TearDownSuite", "section": "sqlDB.Close"}).Fatal(err)
	}
}

func (c *CountByUserIdTestSuite) Test_Success() {
	userId := c.cart.UserId

	totalCart, err := c.cartRepo.CountByUserId(context.Background(), userId)
	assert.NoError(c.T(), err)

	assert.Equal(c.T(), int64(1), totalCart)
}

func (c *CountByUserIdTestSuite) Test_NotFound() {
	userId := "not-found"

	totalCart, err := c.cartRepo.CountByUserId(context.Background(), userId)
	assert.NoError(c.T(), err)

	assert.Equal(c.T(), int64(0), totalCart)
}

func TestRepository_CountByUserId(t *testing.T) {
	suite.Run(t, new(CountByUserIdTestSuite))
}
