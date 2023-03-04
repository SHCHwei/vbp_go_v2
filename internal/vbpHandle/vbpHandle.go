package vbpHandle

import (
    "os"
    "encoding/json"
    "io/ioutil"
)



type Team struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Team string `json:"team"`
	How_many int64 `json:"how_many"`
	BuildTime int64 `json:"buildTime"`
	Status int64 `json:"status"`
}


func ShowList()( []Team , error ) {

    var PlayTeam []Team
    jsonFile, readErr := os.Open("./resource/team.json")

    defer jsonFile.Close()

    if readErr != nil {
            return PlayTeam, readErr
        }

    play, err := ioutil.ReadAll(jsonFile)

    if err != nil {
        return PlayTeam, err
    }

    json.Unmarshal([]byte(play), &PlayTeam)

    return PlayTeam, nil
}

func AddTeam(t Team)( map[string]string, error ) {


    PlayTeam, _ := ShowList()

    responseData := map[string]string{"statue": "false", "message": "建立失敗"}

    if len(PlayTeam) == 0 {
        t.ID = 1
    }else{
        t.ID = PlayTeam[len(PlayTeam) - 1].ID + 1
    }

    //新舊合併
    PlayTeam = append(PlayTeam, t)
	newTeamBytes, err := json.Marshal(PlayTeam)
    if err != nil {
        return responseData, err
    }

    err = ioutil.WriteFile(("./resource/team.json"), newTeamBytes, 0644)
    if err != nil {
        return responseData, err
    }

    responseData["statue"] = "true"
    responseData["message"] = "報隊成功"

    return responseData, nil
}


func DeleteTeam(tid int64)( map[string]string, error) {

    PlayTeam, _ := ShowList()

    responseData := map[string]string{"statue": "false", "message": "建立失敗"}

    var newPlayTeam []Team
	for _, val := range PlayTeam {
		if val.ID == tid {
			val.Status = 2
			newPlayTeam = append(newPlayTeam, val)
		}else{
			newPlayTeam = append(newPlayTeam, val)
		}
	}

    newTeamBytes, err := json.Marshal(newPlayTeam)
    if err != nil {
        return responseData, err
    }

	err = ioutil.WriteFile("./resource/team.json", newTeamBytes, 0644)
    if err != nil {
        return responseData, err
    }

    responseData["statue"] = "true"
    responseData["message"] = "隊伍已清除"

	return responseData, nil
}

func GivenUp(tid int64)( map[string]string, error) {

    PlayTeam, _ := ShowList()

    responseData := map[string]string{"statue": "false", "message": "建立失敗"}

    var newPlayTeam []Team
	for _, val := range PlayTeam {
		if val.ID == tid {
			val.Status = 3
			newPlayTeam = append(newPlayTeam, val)
		}else{
			newPlayTeam = append(newPlayTeam, val)
		}
	}

    newTeamBytes, err := json.Marshal(newPlayTeam)
    if err != nil {
        return responseData, err
    }

	err = ioutil.WriteFile("./resource/team.json", newTeamBytes, 0644)
    if err != nil {
        return responseData, err
    }

    responseData["statue"] = "true"
    responseData["message"] = "隊伍已放棄"

	return responseData, nil
}



func UpdateTeam(t Team)( map[string]string, error) {

    PlayTeam, _ := ShowList()

    responseData := map[string]string{"statue": "false", "message": "建立失敗"}

	var newPlayTeam []Team

	for _, val := range PlayTeam {
		if val.ID == t.ID {
		    val.Title = t.Title
		    val.Team = t.Team
		    val.How_many = t.How_many
			newPlayTeam = append(newPlayTeam, val)
		}else{
			newPlayTeam = append(newPlayTeam, val)
		}

	}

    newTeamBytes, err := json.Marshal(newPlayTeam)
    if err != nil {
        return responseData, err
    }

	err = ioutil.WriteFile("./resource/team.json", newTeamBytes, 0644)
    if err != nil {
        return responseData, err
    }

    responseData["statue"] = "true"
    responseData["message"] = "完成變更"

    return responseData, nil
}
