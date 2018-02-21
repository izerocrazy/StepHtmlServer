package Step

import (
	"KLog"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var EventManagerInstance EventManage
var PlayerPsthInstance PlayerPsth
var PlayerInfoInstace PlayerInfo

type Manager struct {
}

func (m *Manager) Init() {
	var szFileName = "./db.xml"
	PlayerPsthInstance.LoadFromFile(szFileName)
	EventManagerInstance.LoadFromFile()
	PlayerInfoInstace.Init()
}

func (m *Manager) GetAllEventTypeJsonView() []byte {
	return EventManagerInstance.ToJson()
}

func (m *Manager) GetAllEventJsonView() ([]byte, error) {
	v := CreateJsonShow(PlayerPsthInstance.PsthList)

	buf, err := json.Marshal(v.PsthList)
	KLog.CheckErr(err)
	return buf, err
}

func (m *Manager) GetPlayerInfoJsonView() []byte {
	return PlayerInfoInstace.ToJson()
}

func (m *Manager) ShowHtml(w http.ResponseWriter) {
	v := CreateWebShow(PlayerPsthInstance.PsthList)
	v.EventTypeList = EventManagerInstance.EventList

	t, err := template.ParseFiles("./index.html")
	KLog.CheckErr(err)

	err = t.Execute(w, v)
	KLog.CheckErr(err)
}

func (m *Manager) AddStep(nTypeId int) {
	PlayerPsthInstance.AddStep(nTypeId)
	PlayerPsthInstance.SaveToFile("./db.xml")
}

func (m *Manager) UpdateOneStepTime(szStepId string, StartTime string, EndTime string) error {
	step, err := PlayerPsthInstance.GetOneStepById(szStepId)
	if err != nil {
		return err
	}

	step.UpdateTime(StartTime, EndTime)
	PlayerPsthInstance.SaveToFile("./db.xml")

	return nil
}

func (m *Manager) RemoveStepById(StepId string) error {
	fmt.Println("Remove Step By Id", StepId)
	err := PlayerPsthInstance.RemoveStepById(StepId)

	return err
}
