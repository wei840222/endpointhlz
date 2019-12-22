package svc

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type Endpoint struct {
	ID        *int64 `xorm:"pk autoincr 'id'" json:"id"`
	Name      string `xorm:"notnull 'name'" json:"name" validate:"required"`
	URL       string `xorm:"notnull 'url'" json:"url" validate:"required"`
	Status    *int   `xorm:"null 'status'" json:"status,omitempty"`
	CreatedAt int64  `xorm:"notnull created 'created_at'" json:"created_at"`
	UpdatedAt int64  `xorm:"notnull updated 'updated_at'" json:"updated_at"`
}

type EndpointSVC struct {
	db *xorm.Engine
}

func NewEndpointSVC() *EndpointSVC {
	engine, err := xorm.NewEngine("sqlite3", "sqlite3.db")
	if err != nil {
		panic(err)
	}
	engine.ShowExecTime(true)
	engine.ShowSQL(true)

	if err := engine.Ping(); err != nil {
		panic(err)
	}
	if err := engine.Sync2(&Endpoint{}); err != nil {
		panic(err)
	}
	return &EndpointSVC{
		db: engine,
	}
}

func (svc *EndpointSVC) InsertEndpoint(e *Endpoint) error {
	e.ID = nil
	if _, err := svc.db.Insert(e); err != nil {
		return err
	}
	return nil
}

func (svc *EndpointSVC) GetAllEndpoint() ([]*Endpoint, error) {
	e := make([]*Endpoint, 0)
	if err := svc.db.Find(&e); err != nil {
		return nil, err
	}
	return e, nil
}

func (svc *EndpointSVC) UpdateEndpointByID(id int64, e *Endpoint) error {
	e.ID = nil
	if _, err := svc.db.ID(id).Update(e); err != nil {
		return err
	}
	return nil
}

func (svc *EndpointSVC) DeleteEndpointByID(id int64) error {
	if _, err := svc.db.ID(id).Delete(&Endpoint{}); err != nil {
		return err
	}
	return nil
}
