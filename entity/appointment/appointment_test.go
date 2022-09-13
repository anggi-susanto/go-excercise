package appointment_test

import (
	"testing"
	"time"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/appointment"
	"github.com/stretchr/testify/assert"
)

func TestNewAppointment(t *testing.T) {
	time := time.Now()
	a, err := appointment.NewAppointment(1, 1, time)
	assert.Nil(t, err)
	assert.Equal(t, a.CoachID, uint32(1))
	assert.Equal(t, a.UserID, uint32(1))
	assert.Equal(t, a.Date, time)
}

func TestAppointmentValidate(t *testing.T) {
	type test struct {
		coachID uint32
		userID  uint32
		date    time.Time
		want    error
	}

	zeroTime := time.Time{}

	tests := []test{
		{
			coachID: 1,
			userID:  1,
			date:    time.Now(),
			want:    nil,
		},
		{
			coachID: 0,
			userID:  1,
			date:    time.Now(),
			want:    entity.ErrInvalidEntity,
		},
		{
			coachID: 1,
			userID:  0,
			date:    time.Now(),
			want:    entity.ErrInvalidEntity,
		},
		{
			coachID: 1,
			userID:  1,
			date:    zeroTime,
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := appointment.NewAppointment(tc.coachID, tc.userID, tc.date)
		assert.Equal(t, err, tc.want)
	}
}
