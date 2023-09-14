package repo

import (
	"miniproject/model"
)

func (r *Repo) FindAvailableRooms(input []model.Rooms) ([]model.Rooms, error) {

	if err := r.DB.Where("availibility=?", true).Find(&input).Error; err != nil {
		return nil, err
	}

	return input, nil
}
