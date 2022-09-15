package service

import "tabeldatadotcom/archetype/backend-api/pkg/model"

type Service interface {
	SaveAnEmployee(employee *model.Employee) (*model.Employee, error)

}
