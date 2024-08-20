package util

import (
	"github.com/dwprz/prasorganic-cart-service/src/common/log"
	"github.com/dwprz/prasorganic-cart-service/src/model/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CartTest struct {
	db *gorm.DB
}

func NewCartTest(db *gorm.DB) *CartTest {
	return &CartTest{
		db: db,
	}
}

func (c *CartTest) Create() *entity.Cart {
	cart := &entity.Cart{
		UserId:    "hyfa_5Sq7nQcaY6ACksXP",
		ProductId: 10,
		Quantity:  5,
	}

	if err := c.db.Create(cart).Error; err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.NewCartTest/Delete", "section": "db.Exec"}).Error(err.Error())
	}

	return cart
}

func (c *CartTest) Delete() {
	if err := c.db.Exec("DELETE FROM carts").Error; err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.NewCartTest/Delete", "section": "db.Exec"}).Error(err.Error())
	}
}
