package Step

import (
	"fmt"
	"time"
)

type StepView struct {
	StepId        string `json:"step_id"`
	EventTypeId   int    `json:"event_type_id"`
	EventTypeName string `json:"event_type_name"`
	EventTypeIcon string `json:"event_type_icon"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	LastTime      string `json:"last_time"`
}

func CreateStepView(s Step) StepView {
	var showtem StepView

	showtem.StepId = s.Id
	showtem.EventTypeId = s.TypeId
	event := EventManagerInstance.GetEventInfoByType(s.TypeId)
	showtem.EventTypeName = event.TypeName
	showtem.EventTypeIcon = event.TypeIcon
	const layout = time.ANSIC
	showtem.StartTime = s.StartTime.Format(layout)
	if s.StartTime == s.EndTime {
		s.EndTime = time.Now()
	}
	showtem.EndTime = s.EndTime.Format(layout)
	showtem.LastTime = showtem.GetTimeString(s.EndTime.Sub(s.StartTime))

	return showtem
}

func (sv *StepView) GetTimeString(t time.Duration) string {
	nHour := int(t.Hours())
	nMinute := t.Minutes()
	s := fmt.Sprintf("%d : %d", nHour, int(nMinute)-nHour*60)

	return s
}
