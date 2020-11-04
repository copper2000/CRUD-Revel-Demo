package services

import (
	"revelSampleApp/app"
	"revelSampleApp/app/models"
)

func List() (car []models.Car) {
	sql := "SELECT * FROM car"
	rows, err := app.DB.Query(sql)

	if err != nil {
		return nil
	} else {
		var cars []models.Car

		for rows.Next() {
			var id int64
			var name string
			var types string
			var engine string

			err2 := rows.Scan(&id, &name, &types, &engine)

			if err2 != nil {
				return nil
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
		return cars
	}
}

func Get(carID int) (car []models.Car) {
	sql := "SELECT * FROM car where id=?"
	rows, err := app.DB.Query(sql, carID)

	if err != nil {
		return nil
	} else {
		var cars []models.Car

		for rows.Next() {
			var id int64
			var name string
			var types string
			var engine string

			err2 := rows.Scan(&id, &name, &types, &engine)

			if err2 != nil {
				return nil
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
		return cars
	}
}

func Create(car models.Car) (err error) {

	result, err := app.DB.Exec("INSERT INTO car (name, type, engine) VALUES (?, ?, ?)",
		car.Name, car.Type, car.Engine)

	if err != nil {
		return err
	} else {
		car.ID, _ = result.LastInsertId()
		return nil
	}
}

func Update(car models.Car) (err error) {

	result, err := app.DB.Exec("UPDATE car set name = ?, type = ?, engine = ? WHERE id = ?",
		car.Name, car.Type, car.Engine, car.ID)

	if err != nil {
		return err
	} else {
		_, _ = result.RowsAffected()
		return nil
	}
}

func Delete(car models.Car) (err error) {

	result, err := app.DB.Exec("DELETE FROM car WHERE id = ?", car.ID)

	if err != nil {
		return err
	} else {
		_, _ = result.RowsAffected()
		return nil
	}
}

func Paging(pageSize int, pageIndex int) (car []models.Car, err error) {
	sql := "SELECT * FROM car LIMIT ? OFFSET ?"
	rows, err := app.DB.Query(sql, pageSize, (pageIndex - 1) * pageSize)

	if err != nil {
		return nil, err
	} else {
		var cars []models.Car

		for rows.Next() {
			var id int64
			var name string
			var types string
			var engine string

			err2 := rows.Scan(&id, &name, &types, &engine)

			if err2 != nil {
				return nil, err2
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
		return cars, nil
	}
}
