package gormrepository

import (
	"com.ai.orch-purchase-order-inquiry/query"
	"gorm.io/gorm"
)

type GormPurchaseOrderHistoryRepository struct {
	db *gorm.DB
}

func NewPurchaseOrderHistoryRepository(gormDB *gorm.DB) *GormPurchaseOrderHistoryRepository {
	return &GormPurchaseOrderHistoryRepository{db: gormDB}
}

func (repo GormPurchaseOrderHistoryRepository) GetByOrderId(orderId int) (*query.PurchaseOrderHistory, error) {
	var purchaseOrderHistory query.PurchaseOrderHistory
	err := repo.db.Model(&query.PurchaseOrderHistory{}).Where(&query.PurchaseOrderHistory{ItemId: orderId}).First(&purchaseOrderHistory).Error
	return &purchaseOrderHistory, err
}
