package Step

import (
	"KLog"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Event struct {
	TypeId        int    `xml:"type_id" json:"type_id"`
	TypeName      string `xml:"type_name" json:"type_name"`
	TypeIcon      string `xml:"type_icon" json:"type_icon"`
	RecommendTime int    `xml:"type_recommend_time" json:"type_recommend_time"`
}

type EventManage struct {
	EventList []Event `xml:"event_list" json:"event_list"`
}

func (e *EventManage) ToJson() []byte {
	buf, err := json.Marshal(e)

	KLog.CheckErr(err)

	return buf
}

func (e *EventManage) LoadFromFile() {
	buf, err := ioutil.ReadFile("./event.xml")
	if len(buf) == 0 {
		return
	}

	KLog.CheckErr(err)

	err = xml.Unmarshal(buf, &e)

	KLog.CheckErr(err)
}

func (e *EventManage) GetEventInfoByType(nTypId int) *Event {
	for _, v := range e.EventList {
		if nTypId == v.TypeId {
			return &v
		}
	}

	return nil
}

func (e *EventManage) GetRecommendTimeByType(nTypeId int) int {
	for _, v := range e.EventList {
		if nTypeId == v.TypeId {
			return v.RecommendTime
		}
	}

	return 0
}

func EventTest() {
	var em EventManage
	em.LoadFromFile()

	fmt.Println(em.EventList)
}
