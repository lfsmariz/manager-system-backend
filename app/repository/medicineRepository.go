package repository

import (
	"github.com/lfsmariz/manager-system-backend/app/repository/model"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	GetAll() *[]model.Medicine
	Save(*model.Medicine) (*model.Medicine, error)
}

type MedicinePostgres struct {
	db *gorm.DB
}

func NewMedicinePostgres(db *gorm.DB) *MedicinePostgres {
	return &MedicinePostgres{
		db: db,
	}
}

func (m *MedicinePostgres) GetAll() *[]model.Medicine {
	var medicines []model.Medicine

	m.db.Find(&medicines)

	return &medicines
}

func (m *MedicinePostgres) Save(new *model.Medicine) (*model.Medicine, error) {

	result := m.db.Create(new)

	if result.Error == nil {
		return new, nil
	}

	return nil, result.Error
}
