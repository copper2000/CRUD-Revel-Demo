package controllers

import (
	"revelSampleApp/app"
	"revelSampleApp/app/models"

	"github.com/revel/revel"
)

type Cars struct {
	*revel.Controller
}

var cars = []models.Car{
	models.Car{
		ID:     1,
		Name:   "Lamboghini Aventador",
		Type:   "Automatic",
		Engine: "6.5L V12 Engine",
	},
	models.Car{
		ID:     2,
		Name:   "Ford Mustang Shelby GT350",
		Type:   "Manual",
		Engine: "5.2L Ti-VCT V8 Engine",
	},
	models.Car{
		ID:     3,
		Name:   "Nissan GTR 2021",
		Type:   "Automatic",
		Engine: "3.8-liter DOHC 24-valve twin-turbocharged V6",
	},
}

func (c Cars) List() revel.Result {
	sql := "SELECT * FROM car"
	rows, err := app.DB.Query(sql)

	if err != nil {
		return c.NotFound("Could not find car")
	} else {
		var cars []models.Car

		for rows.Next() {
			var id int64
			var name string
			var types string
			var engine string

			err2 := rows.Scan(&id, &name, &types, &engine)

			if err2 != nil {
				return c.NotFound("Error", nil)
			} else {
				car := models.Car{
					ID:     id,
					Name:   name,
					Type:   types,
					Engine: engine,
				}
				cars = append(cars, car)
			}
		}
		return c.RenderJSON(cars)
	}
}

func (c Cars) Show(carID int) revel.Result {
	sql := "SELECT * FROM car where id=?"
	rows, err := app.DB.Query(sql, carID)

	if err != nil {
		return c.NotFound("Could not find car")
	} else {
		var cars []models.Car

		for rows.Next() {
			var id int64
			var name string
			var types string
			var engine string

			err2 := rows.Scan(&id, &name, &types, &engine)

			if err2 != nil {
				return c.NotFound("Error", nil)
			} else {
				car := models.Car{
					ID:     id,
					Name:   name,
					Type:   types,
					Engine: engine,
				}
				cars = append(cars, car)
			}
		}
		return c.RenderJSON(cars)
	}
	//var res models.Car
	//for _, car := range cars {
	//	if car.ID == carID {
	//		res = car
	//	}
	//}
	//if res.ID == 0 {
	//	return c.NotFound("Could not find car")
	//}
	//return c.RenderJSON(res)
}
