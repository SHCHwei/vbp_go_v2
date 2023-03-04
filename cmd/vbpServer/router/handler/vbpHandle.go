package handler


import (
	"time"
	"strconv"
	"net/http"
    "vbp/internal/vbpHandle"
)



//讀取全部隊伍
type LoadTeam struct {
	BaseHandler
}

func (mh *LoadTeam) ExecLoad(rw http.ResponseWriter, req *http.Request) {
    resultData, err := vbpHandle.ShowList()
	mh.responseWith(rw, resultData, err)
}


//新增
type AddTeam struct {
	BaseHandler
}

func (mh *AddTeam) ExecAdd(rw http.ResponseWriter, req *http.Request) {

    var teamInfo vbpHandle.Team

    teamInfo.Title = req.FormValue("title")
    teamInfo.Team = req.FormValue("team")
    teamInfo.How_many, _ = strconv.ParseInt(req.FormValue("how_many"), 10, 64)
    teamInfo.BuildTime = time.Now().Unix()
    teamInfo.Status = 0


	resultData, err := vbpHandle.AddTeam(teamInfo)
	mh.responseWith(rw, resultData, err)
}

//刪除 status值為 2
type DeleteTeam struct {
	BaseHandler
}

func (mh *DeleteTeam) ExecDelete(rw http.ResponseWriter, req *http.Request) {
    tid, _ := strconv.ParseInt(req.FormValue("id"), 10, 64)
	resultData, err := vbpHandle.DeleteTeam(tid)
	mh.responseWith(rw, resultData, err)
}

//放棄 status值為 3
type GivenUp struct{
    BaseHandler
}

func (mh *GivenUp) ExecGivenUp(rw http.ResponseWriter, req *http.Request) {
    tid, _ := strconv.ParseInt(req.FormValue("id"), 10, 64)
	resultData, err := vbpHandle.GivenUp(tid)
	mh.responseWith(rw, resultData, err)
}

//更新
type UpdateTeam struct {
	BaseHandler
}

func (mh *UpdateTeam) ExecUpdate(rw http.ResponseWriter, req *http.Request) {

    var teamInfo vbpHandle.Team
    teamInfo.ID, _ = strconv.ParseInt(req.FormValue("id"), 10, 64)
    teamInfo.Title = req.FormValue("title")
    teamInfo.Team = req.FormValue("team")
    teamInfo.How_many, _ = strconv.ParseInt(req.FormValue("how_many"), 10, 64)

	resultData, err := vbpHandle.UpdateTeam(teamInfo)
	mh.responseWith(rw, resultData, err)
}
