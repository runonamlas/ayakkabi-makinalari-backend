package repository

import (
	"github.com/my-way-teams/my_way_backend/entity"
	"gorm.io/gorm"
)

type RouteRepository interface {
	InsertRoute(r entity.Route) entity.Route
	UpdateRoute(r entity.Route) entity.Route
	DeleteRoute(r entity.Route)
	AllRoute(cityID uint64) []entity.Route
	AllRoutes() []entity.Route
	FindRouteByID(routeID uint64) entity.Route
}

type routeConnection struct {
	connection *gorm.DB
}

func NewRouteRepository(dbConn *gorm.DB) RouteRepository {
	return &routeConnection{
		connection: dbConn,
	}
}

func (db *routeConnection) InsertRoute(r entity.Route) entity.Route {
	db.connection.Save(&r)
	db.connection.Preload("Places.City.Country").Preload("City.Country").Find(&r)
	return r
}

func (db *routeConnection) UpdateRoute(r entity.Route) entity.Route {
	db.connection.Save(&r)
	db.connection.Preload("Places.City.Country").Preload("City.Country").Find(&r)
	return r
}

func (db *routeConnection) DeleteRoute(r entity.Route) {
	db.connection.Delete(&r)
}

func (db *routeConnection) FindRouteByID(routeID uint64) entity.Route {
	var route entity.Route
	db.connection.Preload("Places.City.Country").Preload("City.Country").Find(&route, routeID)
	return route
}

func (db *routeConnection) AllRoute(cityID uint64) []entity.Route {
	var routes []entity.Route
	db.connection.Preload("Places.City.Country").Preload("City.Country").Where("city_id = ?", cityID).Find(&routes)
	return routes
}

func (db *routeConnection) AllRoutes() []entity.Route {
	var routes []entity.Route
	db.connection.Preload("Places.City.Country").Preload("City.Country").Find(&routes)
	return routes
}