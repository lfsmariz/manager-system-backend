package usecase

import (
	"time"

	"github.com/lfsmariz/manager-system-backend/app/repository"
	"github.com/lfsmariz/manager-system-backend/app/repository/model"
)

type MedicineDTO struct {
	Animal              string
	MedicineName        string
	ApplicationDate     time.Time
	NextApplicationDate time.Time
}

func mapToMedicineDTO(m *model.Medicine) *MedicineDTO {
	return &MedicineDTO{
		Animal:              m.Animal,
		MedicineName:        m.MedicineName,
		ApplicationDate:     m.ApplicationDate,
		NextApplicationDate: m.ApplicationDate,
	}
}

func (m *MedicineDTO) mapToMedicineModel() *model.Medicine {
	return &model.Medicine{
		Animal:              m.Animal,
		MedicineName:        m.MedicineName,
		ApplicationDate:     m.ApplicationDate,
		NextApplicationDate: m.ApplicationDate,
	}
}

type MedicineUseCase struct {
	medRepository repository.MedicineRepository
}

func NewMedicineUseCase(medRepository repository.MedicineRepository) MedicineUseCase {
	return MedicineUseCase{
		medRepository,
	}
}

func (m MedicineUseCase) RetrieveAllMedicineList() *[]MedicineDTO {
	all := m.medRepository.GetAll()

	response := make([]MedicineDTO, 0)

	for _, medicine := range *all {
		response = append(response, *mapToMedicineDTO(&medicine))
	}

	return &response
}

func (m MedicineUseCase) SaveMedicine(med *MedicineDTO) (*MedicineDTO, error) {
	_, err := m.medRepository.Save(med.mapToMedicineModel())

	if err == nil {
		return med, nil
	}

	return nil, err
}
