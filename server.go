package main

import (
	"KLog"
	"Step"
	"fmt"
	"net/http"
	"strconv"
)

var StepManager Step.Manager

func main() {
	HtmlServer := http.FileServer(http.Dir("."))
	http.Handle("/", HtmlServer)
	//http.HandleFunc("/test", testFun)

	StepManager.Init()

	http.HandleFunc("/index", Index)
	http.HandleFunc("/BeginEvent", testBegin)

	http.HandleFunc("/all_event", getAllEvent)
	http.HandleFunc("/all_event_type", getAllEventType)
	http.HandleFunc("/begin_event", beginEvent)

	http.HandleFunc("/update_step_time", updateStepTime)
	http.HandleFunc("/delete_step", deleteStep)

	http.HandleFunc("/player_info", getPlayerInfo)

	err := http.ListenAndServe(":8000", nil)
	CheckErr(err)
}

func getAllEventType(w http.ResponseWriter, r *http.Request) {
	KLog.Log("getAllEventType")
	fmt.Fprintf(w, "%q", StepManager.GetAllEventTypeJsonView())
}

func getAllEvent(w http.ResponseWriter, r *http.Request) {
	KLog.Log("getAllEvent")

	buf, err := StepManager.GetAllEventJsonView()
	KLog.CheckErr(err)
	fmt.Fprintf(w, "%q", buf)
}

func getPlayerInfo(w http.ResponseWriter, r *http.Request) {
	KLog.Log("getPlayerInfo")

	fmt.Fprintf(w, "%q", StepManager.GetPlayerInfoJsonView())
}

func beginEvent(w http.ResponseWriter, r *http.Request) {
	szTypeId := r.URL.Query()["TypeId"][0]
	fmt.Println(r.URL.Query())
	nTypeId, _ := strconv.Atoi(szTypeId)
	StepManager.AddStep(nTypeId)

	fmt.Fprintf(w, "%q", "{ret:0}")
}

func updateStepTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	szStepId := r.URL.Query()["StepId"][0]
	szStartTime := r.URL.Query()["StartTime"][0]
	szEndTime := r.URL.Query()["EndTime"][0]

	err := StepManager.UpdateOneStepTime(szStepId, szStartTime, szEndTime)
	KLog.CheckErr(err)

	fmt.Fprintf(w, "%q", "{ret:0}")
}

func deleteStep(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	szStepId := r.URL.Query()["StepId"][0]

	err := StepManager.RemoveStepById(szStepId)
	KLog.CheckErr(err)

	fmt.Fprintf(w, "%q", "{ret:0}")
}

func testBegin(w http.ResponseWriter, r *http.Request) {
	szTypeId := r.URL.Query()["TypeId"][0]
	fmt.Println(r.URL.Query())
	nTypeId, _ := strconv.Atoi(szTypeId)
	StepManager.AddStep(nTypeId)

	testFun(w, r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	StepManager.ShowHtml(w)
}

func testFun(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.RawQuery))
	//fmt.Fprintf(w, "hello, %q", r.URL.RawQuery)
	//fmt.Println(r.URL.Query())

	StepManager.ShowHtml(w)
}

func CheckErr(e error) {
	if e != nil {
		fmt.Println("error :", e.Error())
	}
}
