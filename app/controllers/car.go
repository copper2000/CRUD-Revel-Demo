package controllers

import (
	"github.com/revel/revel"
	"math"
	"revelSampleApp/app"
	"revelSampleApp/app/models"
	"revelSampleApp/app/services"
	"strconv"
)

type Cars struct {
	*revel.Controller
}

func (c Cars) CreateForm() revel.Result {
	return c.Render()
}

func (c Cars) EditForm() revel.Result {

	// get id param from url
	var id = c.Params.Get("id")

	// convert id to int
	var idConvert, _ = strconv.Atoi(id)
	car := services.Get(idConvert)

	return c.Render(car)
	//return c.RenderJSON(car)
}

func (c Cars) RemoveForm() revel.Result {
	return c.Render()
}

func (c Cars) GetAll() revel.Result {
	db := app.DB

	if db != nil {
		cars := services.List()
		return c.Render(cars)
	}

	return nil
}

func (c Cars) GetById() revel.Result {
	db := app.DB

	// get id param from url
	var id = c.Params.Get("id")

	// convert id to int
	var idConvert, _ = strconv.Atoi(id)

	if db != nil {
		car := services.Get(idConvert)
		return c.Render(car)
	}
	return nil
}

func (c Cars) Create() revel.Result {

	var car models.Car
	// bind json body => revel can read it
	_ = c.Params.BindJSON(&car)

	cars := services.Create(car)

	return c.RenderJSON(cars)
}

func (c Cars) Update() revel.Result {

	var car models.Car
	// bind json body => revel can read it
	_ = c.Params.BindJSON(&car)

	cars := services.Update(car)

	return c.RenderJSON(cars)
}

func (c Cars) Delete() revel.Result {

	var car models.Car
	// bind json body => revel can read it
	_ = c.Params.BindJSON(&car)

	cars := services.Delete(car)

	return c.RenderJSON(cars)
}

func (c Cars) Paging() revel.Result {

	// get id param from url
	var pageSize = c.Params.Get("pageSize")
	var pageIndex = c.Params.Get("pageIndex")

	// convert id to int
	var pageSizeConvert, _ = strconv.Atoi(pageSize)
	var pageIndexConvert, _ = strconv.Atoi(pageIndex)

	cars, _ := services.Paging(pageSizeConvert, pageIndexConvert)

	lstCar := services.List()

	var totalItems = len(lstCar)

	totalPage := int(math.Ceil(float64(totalItems / pageSizeConvert)))

	var listPage []int

	for i := 1; i <= totalPage; i++ {
		listPage = append(listPage, i)
	}

	// listPage = append(listPage, pageSizeConvert)
	return c.Render(cars, listPage, pageSizeConvert)

}
