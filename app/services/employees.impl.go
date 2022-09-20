package services

import (
	"context"
	"errors"
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeServiceImpl struct {
	employeeCollection *mongo.Collection
	ctx                context.Context
}

func NewEmployeeService(employeeCollection *mongo.Collection, ctx context.Context) EmployeeService {

	return &EmployeeServiceImpl{
		employeeCollection: employeeCollection,
		ctx:                ctx,
	}

}

func (emp EmployeeServiceImpl) CreateEmployee(employee *models.Employee) error {
	_, err := emp.employeeCollection.InsertOne(emp.ctx, employee)
	return err
}

func (emp EmployeeServiceImpl) GetEmployee(id *primitive.ObjectID) (*models.Employee, error) {

	var employee *models.Employee
	query := bson.D{bson.E{Key: "_id", Value: id}}

	err := emp.employeeCollection.FindOne(emp.ctx, query).Decode(&employee)

	return employee, err

}

func (emp EmployeeServiceImpl) GetAll() ([]*models.Employee, error) {

	var employees []*models.Employee

	cursor, err := emp.employeeCollection.Find(emp.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(emp.ctx) {
		var employee models.Employee
		err := cursor.Decode(&employee)

		if err != nil {
			return nil, err
		}
		employees = append(employees, &employee)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(emp.ctx)

	if len(employees) == 0 {
		return nil, errors.New("documents not found")
	}

	return employees, err

}

func (emp EmployeeServiceImpl) UpdateEmployee(employee *models.Employee) error {

	filter := bson.D{bson.E{Key: "_id", Value: employee.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: employee.Name}, bson.E{Key: "position", Value: employee.Position}}}}

	result, _ := emp.employeeCollection.UpdateOne(emp.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched s found for update")
	}

	return nil
}

func (emp EmployeeServiceImpl) DeleteEmployee(id *primitive.ObjectID) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	result, _ := emp.employeeCollection.DeleteOne(emp.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no matched documents found to delete")
	}

	return nil

}
