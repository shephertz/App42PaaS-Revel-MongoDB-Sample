package controllers

import (
  "App42PaaS-Revel-MongoDB-Sample/app/models"
  "github.com/revel/revel"
  "labix.org/v2/mgo/bson"
  "fmt"
)

type App struct {
	ModelController
}

func (c App) New() revel.Result {
	return c.Render()
}

func (c App) Index(name string, email string, description string) revel.Result{
  coll := db.DB("revel_mongo").C("user")
	fmt.Println("Collection: ", coll)

	err := coll.Insert(models.User{name, email, description})
	PanicIf(err)
	
	users := []models.User{}
	err = coll.Find(bson.M{}).All(&users)
	PanicIf(err)

	return c.Render(users)
}
