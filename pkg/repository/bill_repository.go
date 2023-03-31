package repository

import (
	"errors"
	"strconv"
	"vatansoft/internal/storage"
	"vatansoft/pkg/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BillRepository struct {
	DB    *gorm.DB
	Redis *storage.RedisClient
}

var (
	billTable = "invoices"
)

func NewBillRepository(db *gorm.DB, redis *storage.RedisClient) *BillRepository {
	return &BillRepository{
		DB:    db,
		Redis: redis,
	}
}
func (r *BillRepository) CreateBill(c echo.Context, bill *model.Invoice) (*model.Invoice, error) {
	if bill.InvoiceNo == "" || bill.ProductID == 0 {
		return nil, errors.New("verilerde eksiklik mevcut")
	}
	if err := r.DB.Table(billTable).Create(&bill).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return bill, nil
}

func (r *BillRepository) UpdateBill(c echo.Context, id string, newBill *model.Invoice) (*model.Invoice, error) {
	temporaryProduct := &model.Invoice{
		Model:      newBill.Model,
		InvoiceNo:  newBill.InvoiceNo,
		ProductID:  newBill.ProductID,
		Quantity:   newBill.Quantity,
		TotalPrice: newBill.TotalPrice,
	}

	if err := r.DB.Table(billTable).Where("id = ?", id).Updates(temporaryProduct).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	// Convert the updated product to a ProductResponse object and return it
	return newBill, nil
}
func (r *BillRepository) DeleteBill(c echo.Context, id string) (*model.Invoice, error) {
	var bill model.Invoice
	result := r.DB.Table(billTable).Where("id = ?", id).Scan(&bill).Delete(id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(result.Error.Error())
		}
		return nil, errors.New(result.Error.Error())
	}
	return &bill, nil
}
func (r *BillRepository) GetAllBills(c echo.Context) ([]*model.Invoice, error) {
	var bills []*model.Invoice
	if err := r.DB.Unscoped().Table(billTable).Find(&bills).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return bills, nil
}

func (r *BillRepository) GetBillById(c echo.Context, id string) (*model.Invoice, error) {
	bill := &model.Invoice{}
	if err := r.DB.Table(billTable).Where("id = ?", id).First(bill).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	newId, _ := strconv.Atoi(id)
	bill.ID = uint(newId)
	return bill, nil
}
