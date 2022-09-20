package controllers

import (
	"gypsyland_farming/app/models"
	"gypsyland_farming/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeController struct {
	EmployeeService services.EmployeeService
}

func NewEmployeeController(employeeService services.EmployeeService) EmployeeController {
	return EmployeeController{
		EmployeeService: employeeService,
	}
}

func (ec EmployeeController) CreateEmployee(ctx *gin.Context) {

	var employee models.Employee

	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ec.EmployeeService.CreateEmployee(&employee)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (ec EmployeeController) GetEmployee(ctx *gin.Context) {

	employeeId, _ := primitive.ObjectIDFromHex(ctx.Param("_id"))
	employee, err := ec.EmployeeService.GetEmployee(&employeeId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": employee})

}

func (ec EmployeeController) GetAll(ctx *gin.Context) {

	employees, err := ec.EmployeeService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, employees)
}

func (ec EmployeeController) UpdateEmployee(ctx *gin.Context) {

	var employee models.Employee

	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ec.EmployeeService.UpdateEmployee(&employee)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": employee})

}

func (ec EmployeeController) DeleteEmployee(ctx *gin.Context) {

	id, _ := primitive.ObjectIDFromHex(ctx.Param("_id"))
	err := ec.EmployeeService.DeleteEmployee(&id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (ec EmployeeController) RegisterUserRoutes(rg *gin.RouterGroup) {

	employeeRoute := rg.Group("/employees")
	employeeRoute.POST("/get/:_id", ec.GetEmployee)
	employeeRoute.POST("/getall", ec.GetAll)
	employeeRoute.POST("/update", ec.UpdateEmployee)
	employeeRoute.POST("/create", ec.CreateEmployee)
	employeeRoute.POST("/delete/:_id", ec.DeleteEmployee)

}
