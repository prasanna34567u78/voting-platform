package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onlinevoting/pkg/users"
)

func AddRouter(router *gin.Engine) {
	router.POST("/add", users.PostController)

}

func LoginRouter(router *gin.Engine) {
	router.GET("/login", users.Login)
}

// func SignupRouter(router *gin.Engine) {
// 	router.GET("/signup", users.Signup)
// }

func LoginApiRouter(router *gin.Engine) {
	router.POST("/api/login", users.LoginAPi)
}

func Dashboad(router *gin.Engine) {
	router.GET("/dash", users.Dashboad)
}

func VoterPage(router *gin.Engine) {
	router.GET("/voter", users.Voter)
}

func VoterApiRouter(router *gin.Engine) {
	router.POST("/api/voter", users.AddVoterApi)
}

// delete api route
func VoterDeleteApiRouter(router *gin.Engine) {
	router.DELETE("/api/delete/:id", users.DeleteVote)
}

// router for position
func PositionPage(router *gin.Engine) {
	router.GET("/position", users.Position)
}

func PositionApiRouter(router *gin.Engine) {
	router.POST("/api/position", users.AddPositionApi)
}

func PositionDeleteApiRouter(router *gin.Engine) {
	router.DELETE("/api/position_delete/:id", users.DeletePosition)
}

func CandidatePage(router *gin.Engine) {
	router.GET("/candidate", users.Candidate)
}

func CandidateApiRouter(router *gin.Engine) {
	router.POST("/api/candidate", users.AddCandidateApi)
}
func CandidateDeleteApiRouter(router *gin.Engine) {
	router.DELETE("/api/candidate_delete/:id", users.DeleteCandidate)
}

func BallotPage(router *gin.Engine) {
	router.GET("/ballot", users.Ballot)
}

func Signup(router *gin.Engine) {
	router.POST("/signup", users.SignupPage)
}

func Validate1(router *gin.Engine) {
	router.GET("/validate", users.Validate)
}

func Logout(router *gin.Engine) {
	router.POST("/logout", users.Logout)
}

func LoginVoterApi(router *gin.Engine) {
	router.POST("/voter/login", users.LoginAPiVoter)
}

func Home(router *gin.Engine) {
	router.GET("/home", users.Home)
}
func Profile(router *gin.Engine) {
	router.GET("/profile", users.Profile)
}
func VoterBallot(router *gin.Engine) {
	router.GET("/voterballot", users.VoterBallot)
}
func VoterVerify(router *gin.Engine) {
	router.GET("/verify", users.VoterVerify)
}

func VoterVerifyApi(router *gin.Engine) {
	router.PUT("/voterverify", users.VoterVerifyApi)
}
