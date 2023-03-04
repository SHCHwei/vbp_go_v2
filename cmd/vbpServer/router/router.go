package router

import (
	"net/http"
	"vbp/cmd/vbpServer/router/handler"
	"vbp/internal/middleware"
	"os"
	"fmt"
	"log"
)


func init() {


	f, err := os.Stat("./resource/team.json")
	if err != nil {
        jsonfile, buildErr := os.Create("./resource/team.json")
        if buildErr != nil {
            log.Fatal(buildErr)
        }
        defer jsonfile.Close()
	}

    fmt.Println("f = ", f)

	http.Handle("/showList", middleware.HandlerConv(new(handler.LoadTeam).ExecLoad))
	http.Handle("/addTeam", middleware.HandlerConv(new(handler.AddTeam).ExecAdd))
	http.Handle("/delTeam", middleware.HandlerConv(new(handler.DeleteTeam).ExecDelete))
	http.Handle("/givenUp", middleware.HandlerConv(new(handler.GivenUp).ExecGivenUp))
	http.Handle("/editTeam", middleware.HandlerConv(new(handler.UpdateTeam).ExecUpdate))

}
