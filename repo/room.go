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

// UpdateRoomAvailibility mengubah status availibility kamar.
func (r *Repo) UpdateRoomAvailibility(roomID int, availibility bool) error {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var room model.Rooms
	if err := tx.Where("id = ?", roomID).First(&room).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update status availibility kamar
	room.Availibility = availibility
	if err := tx.Model(&room).Update("availibility", availibility).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
