package repo

import (
	"fmt"
	"miniproject/model"
)

func (r *Repo) FindBooking(userid int) ([]model.Bookings, error) {
	var bookings []model.Bookings

	if err := r.DB.Where("user_id = ?", userid).Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *Repo) Addboking(userid, roomid, totalDays int, status string) (model.Bookings, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var booking model.Bookings

	var room model.Rooms
	if err := tx.Where("id = ?", roomid).First(&room).Error; err != nil {
		tx.Rollback()
		return booking, fmt.Errorf("room not found")
	}

	var user model.Users
	userID := userid

	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return booking, fmt.Errorf("user not found")
	}

	totalAmount := room.Price * totalDays

	if user.Deposit_amount < totalAmount {
		tx.Rollback()
		return booking, fmt.Errorf("balance is not enough")
	}

	if !room.Availibility {
		tx.Rollback()
		return booking, fmt.Errorf("room is not avaliable")
	}

	user.Deposit_amount -= totalAmount

	if err := tx.Model(&user).Update("deposit_amount", user.Deposit_amount).Error; err != nil {
		tx.Rollback()
		return booking, err
	}

	if err := tx.Model(&room).Update("availibility", false).Error; err != nil {
		tx.Rollback()
		return booking, err
	}

	inputBooking := model.Bookings{User_id: userID, Room_id: roomid, Total_day: totalDays, Total_Price: totalAmount}
	if err := tx.Create(&inputBooking).Error; err != nil {
		tx.Rollback()
		return booking, err
	}

	tx.Commit()

	return inputBooking, nil

}

func (r *Repo) EditBooking(bookingid, userid int, roomid, totalDays int, status string) (model.Bookings, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var booking model.Bookings
	if err := tx.Where("id = ? AND user_id = ?", bookingid, userid).First(&booking).Error; err != nil {
		tx.Rollback()
		return booking, fmt.Errorf("booking not found")
	}

	var user model.Users
	userID := userid

	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return booking, fmt.Errorf("user not found")
	}

	var room model.Rooms
	if err := tx.Where("id = ?", roomid).First(&room).Error; err != nil {
		tx.Rollback()
		return booking, fmt.Errorf("room not found")
	}

	totalAmount := room.Price * totalDays

	// Kembalikan saldo pengguna
	user.Deposit_amount += booking.Total_Price
	if err := tx.Model(&user).Update("deposit_amount", user.Deposit_amount).Error; err != nil {
		tx.Rollback()
		return booking, err
	}

	// Kembalikan status ketersediaan kamar jika sebelumnya tidak tersedia
	if room.Availibility {
		if err := tx.Model(&room).Update("availibility", true).Error; err != nil {
			tx.Rollback()
			return booking, err
		}
	}

	user.Deposit_amount -= totalAmount

	if err := tx.Model(&user).Update("deposit_amount", user.Deposit_amount).Error; err != nil {
		tx.Rollback()
		return booking, err
	}

	if err := tx.Model(&room).Update("availibility", false).Error; err != nil {
		tx.Rollback()
		return booking, err
	}

	// Update status booking
	booking.Status = status
	booking.Room_id = roomid
	booking.Total_day = totalDays
	booking.Total_Price = totalAmount

	if err := tx.Save(&booking).Error; err != nil {
		tx.Rollback()
		return booking, err
	}

	tx.Commit()

	return booking, nil
}
