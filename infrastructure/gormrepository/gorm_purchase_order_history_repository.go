package gormrepository

import (
	"com.ai.orch-purchase-order-inquiry/query"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

// Initialize initializes the database connection.
func Initialize() {
	database, err := gorm.Open(postgres.Open(fmt.Sprintf("host=localhost user=postgres password=34647de759 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok")))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = database
}

// GetDB returns the database instance.
func GetDB() *gorm.DB {
	return db
}

func GetByOrderId(orderId int) (*query.PurchaseOrderHistory, error) {
	var partnerUser query.PurchaseOrderHistory
	err := GetDB().Model(&query.PurchaseOrderHistory{}).Where(&query.PurchaseOrderHistory{ItemId: orderId}).First(&partnerUser).Error
	return &partnerUser, err
}
