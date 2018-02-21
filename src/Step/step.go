package Step

import (
	"KLog"
	"fmt"
	// "strconv"
	"time"
)

type Step struct {
	Id        string    `xml:"id"`
	TypeId    int       `xml:"type_id"`
	StartTime time.Time `xml:"start_time"`
	EndTime   time.Time `xml:"end_time"`
}

func (s *Step) UpdateTime(StartTime string, EndTime string) {
	fmt.Printf("Update Time: %s:%s\n", StartTime, EndTime)

	t, err := time.Parse("2006-01-02 15:04:05", StartTime)
	KLog.CheckErr(err)
	KLog.Log("Start Time: %s", t)
	s.StartTime = t

	t, err = time.Parse("2006-01-02 15:04:05", EndTime)
	KLog.CheckErr(err)
	KLog.Log("End Time: %s", t)
	s.EndTime = t
}
