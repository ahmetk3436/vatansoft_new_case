package service

import (
	"errors"
	"vatansoft/pkg/model"
	"vatansoft/pkg/repository"

	"github.com/labstack/echo/v4"
)

type BillService struct {
	repository *repository.BillRepository
}

func NewBillService(repository *repository.BillRepository) *BillService {
	return &BillService{
		repository: repository,
	}
}

func (c *BillService) CreateBillService(e echo.Context, dto *model.Invoice) (_ *model.Invoice, err error) {
	dto, err = c.repository.CreateBill(e, dto)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return dto, nil
}

func (c *BillService) UpdateBillService(e echo.Context, id string, newBill *model.Invoice) (*model.Invoice, error) {
	newBill, updateErr := c.repository.UpdateBill(e, id, newBill)
	if updateErr != nil {
		return nil, errors.New(updateErr.Error())
	}
	return newBill, nil
}

func (c *BillService) DeleteBillService(e echo.Context, id string) (bill *model.Invoice, err error) {
	bill, err = c.repository.DeleteBill(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return bill, nil
}
func (c *BillService) GetAllBillsService(e echo.Context) (bills []*model.Invoice, err error) {
	bills, err = c.repository.GetAllBills(e)
	if err != nil && len(bills) == 0 {
		return nil, errors.New("sistemde ürün bulunmamaktadır")
	}
	return bills, nil
}

func (c *BillService) GetBillByIdService(e echo.Context, id string) (bill *model.Invoice, err error) {
	bill, err = c.repository.GetBillById(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return bill, nil
}
