package appointment

import (
	"time"

	"github.com/anggi-susanto/go-exercise/entity"
)

type Appointment struct {
	ID      *uint32
	CoachID uint32
	UserID  uint32
	Date    time.Time
}

func NewAppointment(coachID uint32, userID uint32, date time.Time) (*Appointment, error) {
	a := &Appointment{
		CoachID: coachID,
		UserID:  userID,
		Date:    date,
	}

	err := a.Validate()

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}

	return a, nil
}

func (a *Appointment) Validate() error {
	zeroTime := time.Time{}
	if a.CoachID < 1 || a.UserID < 1 || a.Date == zeroTime {
		return entity.ErrInvalidEntity
	}
	return nil
}
