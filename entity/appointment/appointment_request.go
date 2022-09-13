package appointment

import (
	"time"

	"github.com/anggi-susanto/go-exercise/entity"
)

type AppointmentRequest struct {
	ID                  *uint32
	Type                string
	RequesterID         uint32
	ApproverID          uint32
	Date                time.Time
	Status              string
	RescheduleRequestID *uint32
}

func NewAppointmentRequest(types string, requesterID uint32, approverID uint32, date time.Time, status string, rescheduleRequestID uint32) (*AppointmentRequest, error) {
	ar := &AppointmentRequest{
		Type:                types,
		RequesterID:         requesterID,
		ApproverID:          approverID,
		Date:                date,
		Status:              status,
		RescheduleRequestID: &rescheduleRequestID,
	}

	err := ar.Validate()

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return ar, nil
}

func (ar *AppointmentRequest) Validate() error {
	zeroTime := time.Time{}

	if ar.Type == "" || ar.RequesterID < 1 || ar.ApproverID < 1 || ar.Date == zeroTime || ar.Status == "" {
		return entity.ErrInvalidEntity
	}
	return nil
}
