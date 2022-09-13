package appointment_test

import (
	"testing"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/appointment"
	"github.com/stretchr/testify/assert"
)

func TestNewCoachAvailability(t *testing.T) {
	c, err := appointment.NewCoachAvailability(1, "a", "a", "a", "a")
	assert.Nil(t, err)
	assert.Equal(t, c.UserID, uint32(1))
	assert.Equal(t, c.Timezone, "a")
	assert.Equal(t, c.DayOfWeek, "a")
	assert.Equal(t, c.TimeStart, "a")
	assert.Equal(t, c.TimeEnd, "a")
}

func TestCoachAvailabilityValidate(t *testing.T) {
	type test struct {
		userId    uint32
		timezone  string
		dayOfWeek string
		timeStart string
		timeEnd   string
		want      error
	}

	tests := []test{
		{
			userId:    1,
			timezone:  "a",
			dayOfWeek: "a",
			timeStart: "a",
			timeEnd:   "a",
			want:      nil,
		},
		{
			userId:    0,
			timezone:  "a",
			dayOfWeek: "a",
			timeStart: "a",
			timeEnd:   "a",
			want:      entity.ErrInvalidEntity,
		},
		{
			userId:    1,
			timezone:  "",
			dayOfWeek: "a",
			timeStart: "a",
			timeEnd:   "a",
			want:      entity.ErrInvalidEntity,
		},
		{
			userId:    1,
			timezone:  "a",
			dayOfWeek: "",
			timeStart: "a",
			timeEnd:   "a",
			want:      entity.ErrInvalidEntity,
		},
		{
			userId:    1,
			timezone:  "a",
			dayOfWeek: "a",
			timeStart: "",
			timeEnd:   "a",
			want:      entity.ErrInvalidEntity,
		},
		{
			userId:    1,
			timezone:  "a",
			dayOfWeek: "a",
			timeStart: "a",
			timeEnd:   "",
			want:      entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := appointment.NewCoachAvailability(tc.userId, tc.timezone, tc.dayOfWeek, tc.timeStart, tc.timeEnd)
		assert.Equal(t, err, tc.want)
	}
}
