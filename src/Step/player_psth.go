package Step

import (
	"KLog"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"time"
)

const temp = `TypeId={{.TypeId}}, StartTime={{.StartTime}}, EndTime = {{.EndTime}}`

type PlayerPsth struct {
	PsthList []Step `xml:"psth_list"`
}

// Create MD5
func (p *PlayerPsth) CreateMD5ByTime(t time.Time) string {
	const layout = time.ANSIC
	temp := t.Format(layout)

	return p.CreateMD5ByString(temp)
}

func (p *PlayerPsth) CreateMD5ByString(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	cipherStr := h.Sum(nil)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果

	return hex.EncodeToString(cipherStr)
}

func (p *PlayerPsth) AddStep(nTypeId int) {
	i := len(p.PsthList)
	now := time.Now()
	md5 := p.CreateMD5ByTime(now)
	minutes := EventManagerInstance.GetRecommendTimeByType(nTypeId)
	newPsth := Step{
		Id:        md5,
		TypeId:    nTypeId,
		StartTime: now,
		EndTime:   now.Add(time.Duration(minutes) * time.Minute), // 默认五分钟后结束
	}

	// 自动将最后一个设置为当前结束
	if i > 0 {
		p.PsthList[i-1].EndTime, _ = time.Parse("2006/1/2 15:04:05", now.Format("2006/1/2 15:04:05"))
		PlayerInfoInstace.InsertStep(&p.PsthList[i-1])
	}

	p.PsthList = append(p.PsthList, newPsth)
}

func (p *PlayerPsth) LoadFromFile(szFileName string) {
	buf, err := ioutil.ReadFile(szFileName)
	if len(buf) == 0 {
		return
	}

	KLog.CheckErr(err)

	err = xml.Unmarshal(buf, &p)

	// 如果没有 id 则主动加上 Id
	for i := 0; i < len(p.PsthList); i++ {
		psth := p.PsthList[i]
		if psth.Id == "" {
			p.PsthList[i].Id = p.CreateMD5ByTime(psth.StartTime)
		}
	}

	KLog.CheckErr(err)
}

func (p *PlayerPsth) SaveToFile(szFileName string) {
	buf, err := xml.Marshal(p)
	KLog.CheckErr(err)

	ioutil.WriteFile(szFileName, buf, 0777)
}

func (p *PlayerPsth) ShowTemplate() {
	for i := 0; i < len(p.PsthList); i++ {
		psth := p.PsthList[i]

		t := template.New("Psth template")
		t, err := t.Parse(temp)
		KLog.CheckErr(err)

		err = t.Execute(os.Stdout, psth)
		KLog.CheckErr(err)
	}
}

func (p *PlayerPsth) GetOneStepById(StepId string) (step *Step, err error) {
	for i := 0; i < len(p.PsthList); i++ {
		if p.PsthList[i].Id == StepId {
			return &p.PsthList[i], nil
		}
	}

	return nil, errors.New("Step No Exist")
}

func (p *PlayerPsth) GetOneStepIndexById(StepId string) (index int, err error) {
	for i := 0; i < len(p.PsthList); i++ {
		if p.PsthList[i].Id == StepId {
			return i, nil
		}
	}

	return -1, errors.New("Step No Exist")
}

func (p *PlayerPsth) RemoveStepById(StepId string) (err error) {
	index, err := p.GetOneStepIndexById(StepId)
	KLog.CheckErr(err)

	p.PsthList = append(p.PsthList[:index], p.PsthList[index+1:]...)
	return nil
}

func PsthTest() {
	var MyPsthList PlayerPsth
	var szFileName = "./db.xml"

	MyPsthList.LoadFromFile(szFileName)
	fmt.Println(len(MyPsthList.PsthList))

	if len(MyPsthList.PsthList) == 0 {
		MyPsthList.AddStep(1)
	}

	MyPsthList.ShowTemplate()
	MyPsthList.SaveToFile(szFileName)
}
