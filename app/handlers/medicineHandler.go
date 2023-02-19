package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfsmariz/manager-system-backend/app/repository"
	"github.com/lfsmariz/manager-system-backend/app/usecase"
)

type MedicineHandler struct {
	medCase usecase.MedicineUseCase
}

func NewMedicineHandler() *MedicineHandler {
	medRepository := repository.NewMedicinePostgres(repository.GetInstance().Db)

	medCase := usecase.NewMedicineUseCase(medRepository)

	return &MedicineHandler{
		medCase,
	}
}

type medicineResponse struct {
	Animal              string    `json:"animal"`
	MedicineName        string    `json:"medicineName"`
	ApplicationDate     time.Time `json:"applicationDate"`
	NextApplicationDate time.Time `json:"nextApplicationDate"`
}

type medicineRequest struct {
	Animal              string    `json:"animal" binding:"required"`
	MedicineName        string    `json:"medicineName" binding:"required"`
	ApplicationDate     time.Time `json:"applicationDate" binding:"required"`
	NextApplicationDate time.Time `json:"nextApplicationDate" binding:"required"`
}

func mapToMedicineResponse(m *usecase.MedicineDTO) *medicineResponse {
	return &medicineResponse{
		Animal:              m.Animal,
		MedicineName:        m.MedicineName,
		ApplicationDate:     m.ApplicationDate,
		NextApplicationDate: m.ApplicationDate,
	}
}

func (m *medicineRequest) mapToMedicineDTO() *usecase.MedicineDTO {
	return &usecase.MedicineDTO{
		Animal:              m.Animal,
		MedicineName:        m.MedicineName,
		ApplicationDate:     m.ApplicationDate,
		NextApplicationDate: m.ApplicationDate,
	}
}

func (m *MedicineHandler) GetAll(c *gin.Context) {

	response := make([]medicineResponse, 0)

	list := m.medCase.RetrieveAllMedicineList()

	for _, v := range *list {
		response = append(response, *mapToMedicineResponse(&v))
	}
	c.JSON(http.StatusOK, &response)
}

func (m *MedicineHandler) Save(c *gin.Context) {

	req := medicineRequest{}

	c.ShouldBind(&req)

	dto := req.mapToMedicineDTO()

	respDTO, err := m.medCase.SaveMedicine(dto)

	if err == nil {
		c.JSON(http.StatusOK, mapToMedicineResponse(respDTO))
		return
	}

	fmt.Println(err.Error())
	c.JSON(http.StatusInternalServerError, nil)
}
