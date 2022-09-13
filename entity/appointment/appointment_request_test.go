package appointment_test

import (
	"testing"
	"time"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/appointment"
	"github.com/stretchr/testify/assert"
)

func TestNewAppointmentRequest(t *testing.T) {
	time := time.Now()
	ar, err := appointment.NewAppointmentRequest("a", 1, 1, time, "a", 1)

	assert.Nil(t, err)
	assert.Equal(t, ar.Type, "a")
	assert.Equal(t, ar.RequesterID, uint32(1))
	assert.Equal(t, ar.ApproverID, uint32(1))
	assert.Equal(t, ar.Date, time)
	assert.Equal(t, ar.Status, "a")
}

func TestAppointmentRequestValidate(t *testing.T) {
	type test struct {
		types               string
		requesterID         uint32
		approverID          uint32
		date                time.Time
		status              string
		rescheduleRequestID uint32
		want                error
	}

	zeroTime := time.Time{}

	tests := []test{
		{
			types:               "a",
			requesterID:         1,
			approverID:          1,
			date:                time.Now(),
			status:              "a",
			rescheduleRequestID: 1,
			want:                nil,
		},
		{
			types:               "",
			requesterID:         1,
			approverID:          1,
			date:                time.Now(),
			status:              "a",
			rescheduleRequestID: 1,
			want:                entity.ErrInvalidEntity,
		},
		{
			types:               "a",
			requesterID:         0,
			approverID:          1,
			date:                time.Now(),
			status:              "a",
			rescheduleRequestID: 1,
			want:                entity.ErrInvalidEntity,
		},
		{
			types:               "a",
			requesterID:         1,
			approverID:          0,
			date:                time.Now(),
			status:              "a",
			rescheduleRequestID: 1,
			want:                entity.ErrInvalidEntity,
		},
		{
			types:               "a",
			requesterID:         1,
			approverID:          1,
			date:                zeroTime,
			status:              "a",
			rescheduleRequestID: 1,
			want:                entity.ErrInvalidEntity,
		},
		{
			types:               "a",
			requesterID:         1,
			approverID:          1,
			date:                time.Now(),
			status:              "",
			rescheduleRequestID: 1,
			want:                entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := appointment.NewAppointmentRequest(tc.types, tc.requesterID, tc.approverID, tc.date, tc.status, tc.rescheduleRequestID)
		assert.Equal(t, err, tc.want)
	}
}
