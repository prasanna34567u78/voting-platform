package main

import (
	"github.com/gin-gonic/gin"
	"github.com/onlinevoting/initializers"
	database "github.com/onlinevoting/pkg/Database"
	"github.com/onlinevoting/pkg/middleware"
	"github.com/onlinevoting/pkg/routes"
)

func init() {
	initializers.LoadEnv()
	//setting up db creating table comment below function after calling firsttime

	database.SetUpDB()
}

func main() {

	router := gin.New()
	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")
	router.Static("/images", "./templates/images")
	router.LoadHTMLGlob("templates/*.html")
	routes.LoginRouter(router)
	routes.LoginApiRouter(router)
	routes.LoginVoterApi(router)
	routes.Signup(router)
	//voter
	router.Group("/", func(ctx *gin.Context) {
		router.Use(middleware.AuthVoter)
		routes.Home(router)
		routes.Profile(router)
		routes.VoterBallot(router)
		routes.VoterVerify(router)
		routes.VoterVerifyApi(router)
		routes.Logout(router)
	})
	routes.Logout(router)
	router.Use(middleware.Auth)

	routes.AddRouter(router)

	routes.VoterPage(router)
	routes.BallotPage(router)
	routes.Validate1(router)
	routes.PositionPage(router)
	routes.CandidatePage(router)
	routes.Dashboad(router)
	routes.VoterApiRouter(router)
	routes.VoterDeleteApiRouter(router)
	routes.PositionApiRouter(router)
	routes.PositionDeleteApiRouter(router)
	routes.CandidateDeleteApiRouter(router)
	routes.CandidateApiRouter(router)
	// routes.Logout(router)
	router.Run(":8000")
	// fmt.Println("Prsanna kumar")
}
