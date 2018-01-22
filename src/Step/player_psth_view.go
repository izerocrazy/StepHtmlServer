package Step

import (
	"time"
)

type JsonShow struct {
	EventTypeList []Event    `json:"event_type_list"`
	PsthList      []StepView `json:"psth_list"`
}

func CreateJsonShow(stepList []Step) JsonShow {
	var ws JsonShow
	nLen := len(stepList)
	nTimeNow := time.Now()
	for k, _ := range stepList {
		v := stepList[nLen-k-1]

		// 超过 24 小时的就忽略掉
		if nTimeNow.Sub(v.EndTime).Hours() > 24 {
			continue
		}

		ws.PsthList = append(ws.PsthList, CreateStepView(v))
	}
	return ws
}

////////////////////////////////////////
type WebShow struct {
	EventTypeList []Event    `json:"event_type_list"`
	PsthList      []StepView `json:"psth_list"`
	NowStep       StepView   `json:"now_step"`
}

func CreateWebShow(stepList []Step) WebShow {
	var ws WebShow
	nLen := len(stepList)
	nTimeNow := time.Now()
	for k, _ := range stepList {
		v := stepList[nLen-k-1]
		// 第一个也就是当前的事件
		if k == 0 {
			ws.NowStep = CreateStepView(v)

			timeSince := time.Since(v.StartTime)
			ws.NowStep.LastTime = ws.NowStep.GetTimeString(timeSince)

			continue
		}

		// 超过 24 小时的就忽略掉
		if nTimeNow.Sub(v.EndTime).Hours() > 24 {
			continue
		}

		ws.PsthList = append(ws.PsthList, CreateStepView(v))
	}
	return ws
}
