package test

import (
	"context"
	"testing"

	"github.com/dwprz/prasorganic-cart-service/src/common/log"
	"github.com/dwprz/prasorganic-cart-service/src/infrastructure/database"
	"github.com/dwprz/prasorganic-cart-service/src/interface/repository"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	repositoryimpl "github.com/dwprz/prasorganic-cart-service/src/repository"
	"github.com/dwprz/prasorganic-cart-service/test/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_CreateCart$ -v ./src/repository/test/ -count=1

type CreateCartTestSuite struct {
	suite.Suite
	cartRepo    repository.Cart
	postgresDB  *gorm.DB
	cartTestUtl *util.CartTest
}

func (c *CreateCartTestSuite) SetupSuite() {
	c.postgresDB = database.NewPostgres()
	c.cartRepo = repositoryimpl.NewCart(c.postgresDB)

	c.cartTestUtl = util.NewCartTest(c.postgresDB)
}

func (c *CreateCartTestSuite) TearDownTest() {
	c.cartTestUtl.Delete()
}

func (c *CreateCartTestSuite) TearDownSuite() {
	sqlDB, err := c.postgresDB.DB()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.CreateCartTestSuite/TearDownSuite", "section": "postgresDB.DB"}).Fatal(err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.CreateCartTestSuite/TearDownSuite", "section": "sqlDB.Close"}).Fatal(err)
	}
}

func (c *CreateCartTestSuite) Test_Success() {
	err := c.cartRepo.Create(context.Background(), &dto.CreateCartReq{
		UserId:    "hyfa_5Sq7nQcaY6ACksXP",
		ProductId: 10,
		Quantity:  5,
	})

	assert.NoError(c.T(), err)
}

func (c *CreateCartTestSuite) Test_AlreadyExists() {
	c.cartRepo.Create(context.Background(), &dto.CreateCartReq{
		UserId:    "hyfa_5Sq7nQcaY6ACksXb",
		ProductId: 10,
		Quantity:  5,
	})

	err := c.cartRepo.Create(context.Background(), &dto.CreateCartReq{
		UserId:    "hyfa_5Sq7nQcaY6ACksXb",
		ProductId: 10,
		Quantity:  5,
	})

	assert.Error(c.T(), err)
}

func TestRepository_CreateCart(t *testing.T) {
	suite.Run(t, new(CreateCartTestSuite))
}
