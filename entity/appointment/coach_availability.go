package appointment

import (
	"github.com/anggi-susanto/go-exercise/entity"
)

type CoachAvailability struct {
	ID        *uint32
	UserID    uint32
	Timezone  string
	DayOfWeek string
	TimeStart string
	TimeEnd   string
}

func NewCoachAvailability(userID uint32, timezone string, dayOfWeek string, timeStart string, timeEnd string) (*CoachAvailability, error) {
	c := &CoachAvailability{
		UserID:    userID,
		Timezone:  timezone,
		DayOfWeek: dayOfWeek,
		TimeStart: timeStart,
		TimeEnd:   timeEnd,
	}

	err := c.Validate()
	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return c, nil
}

func (c *CoachAvailability) Validate() error {
	if c.UserID < 1 || c.Timezone == "" || c.DayOfWeek == "" || c.TimeStart == "" || c.TimeEnd == "" {
		return entity.ErrInvalidEntity
	}
	return nil
}
