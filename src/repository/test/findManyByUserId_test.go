package test

import (
	"context"
	"testing"

	"github.com/dwprz/prasorganic-cart-service/src/common/errors"
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
// go test -run ^TestRepository_FindManyByUserId$ -v ./src/repository/test/ -count=1

type FindManyByUserIdTestSuite struct {
	suite.Suite
	cart        *entity.Cart
	cartRepo    repository.Cart
	postgresDB  *gorm.DB
	cartTestUtl *util.CartTest
}

func (c *FindManyByUserIdTestSuite) SetupSuite() {
	c.postgresDB = database.NewPostgres()
	c.cartRepo = repositoryimpl.NewCart(c.postgresDB)

	c.cartTestUtl = util.NewCartTest(c.postgresDB)

	c.cart = c.cartTestUtl.Create()
}

func (c *FindManyByUserIdTestSuite) TearDownSuite() {
	c.cartTestUtl.Delete()

	sqlDB, err := c.postgresDB.DB()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.FindManyByUserIdTestSuite/TearDownSuite", "section": "postgresDB.DB"}).Fatal(err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.FindManyByUserIdTestSuite/TearDownSuite", "section": "sqlDB.Close"}).Fatal(err)
	}
}

func (c *FindManyByUserIdTestSuite) Test_Success() {
	userId := c.cart.UserId
	limit := 20
	offset := 0

	res, err := c.cartRepo.FindManyByUserId(context.Background(), userId, limit, offset)
	assert.NoError(c.T(), err)

	assert.NotEmpty(c.T(), res.Cart)
}

func (c *FindManyByUserIdTestSuite) Test_NotFound() {
	userId := "not-found"
	limit := 20
	offset := 0

	res, err := c.cartRepo.FindManyByUserId(context.Background(), userId, limit, offset)
	assert.Error(c.T(), err)

	resErr, ok := err.(*errors.Response)
	assert.True(c.T(), ok)

	assert.Equal(c.T(), 404, resErr.HttpCode)
	assert.Nil(c.T(), res)
}

func TestRepository_FindManyByUserId(t *testing.T) {
	suite.Run(t, new(FindManyByUserIdTestSuite))
}
