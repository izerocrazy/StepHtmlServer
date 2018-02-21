package Step

import (
	"KLog"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	/*
		"time"
	*/)

type EventValueConfig struct {
	EventType int `xml:"type_id"`
	UnitValue int `xml:"unit_value"`
	UnitTime  int `xml:"unit_time"`
}

type ValueConfigList struct {
	ConfigList []EventValueConfig `xml:"value_info"`
}

func (v *ValueConfigList) Init(szValueName string) {
	KLog.Log("ValueConfigList Init %s", szValueName)
	var szFileName = "./" + szValueName + "_value_config.xml"

	buf, err := ioutil.ReadFile(szFileName)
	KLog.Asset(len(buf) != 0, "read file fail: %s", szFileName)
	KLog.CheckErr(err)

	err = xml.Unmarshal(buf, &v)
	KLog.CheckErr(err)
}

func (v *ValueConfigList) CalcValue(type_id int, time_m int) int {
	KLog.Log("ValueConfigList CalcValue %d %d", type_id, time_m)

	var ret = 0
	KLog.Asset(type_id > 0, "type id is error: %d", type_id)
	KLog.Asset(time_m >= 0, "time is error: %d", time_m)

	for i := 0; i < len(v.ConfigList); i++ {
		config := v.ConfigList[i]
		if config.EventType == type_id {
			ret = time_m / config.UnitTime * config.UnitValue
			KLog.Log("ValueConfigList CalcValue, the calc value is %d", ret)
			break
		}
	}

	return ret
}

//////////////////////////////////////////////
type PlayerInfo struct {
	HealthValue      int `xml:"health_value" json:"health_value"`
	HealthConfigList ValueConfigList

	PhysicalValue      int `xml:"physical_value" json:"physical_value"`
	PhysicalConfigList ValueConfigList
}

func (p *PlayerInfo) Init() {
	KLog.Log("PlayerInfo Init")

	// Load From File
	var szFileName = "./player_info.xml"
	p.LoadFromFile(szFileName)

	p.HealthConfigList.Init("health")
	p.PhysicalConfigList.Init("physical")
}

func (p *PlayerInfo) LoadFromFile(szFileName string) {
	buf, err := ioutil.ReadFile(szFileName)
	KLog.Asset(len(buf) != 0, "read file fail: %s", szFileName)
	KLog.CheckErr(err)

	err = xml.Unmarshal(buf, &p)
	KLog.CheckErr(err)
}

func (p *PlayerInfo) SaveToFile(szFileName string) {
	buf, err := xml.Marshal(p)
	KLog.CheckErr(err)

	ioutil.WriteFile(szFileName, buf, 0777)
}

func (p *PlayerInfo) InsertStep(s *Step) {
	KLog.Log("PlayerInfo InsertStep Type is %d, EndTime %s, StartTime %s", s.TypeId, s.EndTime, s.StartTime)

	// 通过 Step 获取到 Event
	type_id := s.TypeId
	// 通过 Step 获取到 时间
	time := int(s.EndTime.Sub(s.StartTime).Minutes())

	p.HealthValue += p.HealthConfigList.CalcValue(type_id, time)
	p.PhysicalValue += p.PhysicalConfigList.CalcValue(type_id, time)
	KLog.Log("PlayerInfo Health is %d, Physical is %d", p.HealthValue, p.PhysicalValue)

	var szFileName = "./player_info.xml"
	p.SaveToFile(szFileName)
}

func (p *PlayerInfo) ToJson() []byte {
	buf, err := json.Marshal(p)

	KLog.CheckErr(err)

	return buf
}
