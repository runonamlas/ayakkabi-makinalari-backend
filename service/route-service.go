package service

import (
	"fmt"
	"github.com/mashingan/smapping"
	"github.com/my-way-teams/my_way_backend/dto"
	"github.com/my-way-teams/my_way_backend/entity"
	"github.com/my-way-teams/my_way_backend/repository"
	"log"
)

type RouteService interface {
	Insert(r dto.RouteCreateDTO) entity.Route
	Update(r dto.RouteUpdateDTO) entity.Route
	Delete(r entity.Route)
	All(cityID uint64) []entity.Route
	AllRoutes() []entity.Route
	FindByID(routeID uint64) entity.Route
	IsAllowedToEdit(cityID string, routeID uint64) bool
}

type routeService struct {
	routeRepository repository.RouteRepository
}

func NewRouteService(routeRepo repository.RouteRepository) RouteService {
	return &routeService{
		routeRepository: routeRepo,
	}
}

func (service *routeService) Insert(r dto.RouteCreateDTO) entity.Route {
	route := entity.Route{}
	err := smapping.FillStruct(&route, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.routeRepository.InsertRoute(route)
	return res
}

func (service *routeService) Update(r dto.RouteUpdateDTO) entity.Route {
	route := entity.Route{}
	err := smapping.FillStruct(&route, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.routeRepository.UpdateRoute(route)
	return res
}

func (service *routeService) Delete(r entity.Route) {
	service.routeRepository.DeleteRoute(r)
}

func (service *routeService) All(cityID uint64) []entity.Route {
	return service.routeRepository.AllRoute(cityID)
}

func (service *routeService) AllRoutes() []entity.Route {
	return service.routeRepository.AllRoutes()
}

func (service *routeService) FindByID(routeID uint64) entity.Route {
	return service.routeRepository.FindRouteByID(routeID)
}

func (service *routeService) IsAllowedToEdit(cityID string, routeID uint64) bool {
	p := service.routeRepository.FindRouteByID(routeID)
	id := fmt.Sprintf("%v", p.CityID)
	return cityID == id
}