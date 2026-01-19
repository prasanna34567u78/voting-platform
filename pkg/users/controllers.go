package users

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	database "github.com/onlinevoting/pkg/Database"
	"github.com/onlinevoting/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	user      model.Users
	voter     model.Voter
	position  model.Position
	candidate model.Candidate
)

func UserController(c *gin.Context) {
	c.String(200, "hello world")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func PostController(c *gin.Context) {
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	temp := GetAllEmail("user")
	if contains(temp, user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exits"})
	} else {
		addUser(&user)
		successResponse := gin.H{
			"message": "successfully register go to login page",
		}
		failureResponse := gin.H{
			"message": "Failure",
			"error":   err,
		}
		if err == nil {
			c.JSON(200, successResponse)
		} else {
			c.JSON(http.StatusBadRequest, failureResponse)
		}
	}

}

func addUser(user *model.Users) {

	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `INSERT INTO users (email,password) VALUES (?,?)`
	newcontroller.ExcuteQuery(query, user.Email, user.Password)

}

func Login(c *gin.Context) {

	c.HTML(200, "login.html", nil)
}
func Signup(c *gin.Context) {

	c.HTML(200, "signup.html", nil)
}

func GetAllEmail(user string) []string {
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `SELECT email from ` + user + ` ALLOW FILTERING`
	data, err := newcontroller.GetData(query)
	if err != nil {
		log.Fatal(err)
	}

	return data

}

// function to get all password to check whether the user enter valid password or not
func GetAllPassword(user *model.Users) []string {
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `SELECT password from users where email = ? ALLOW FILTERING`
	data, err := newcontroller.GetData(query, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	return data

}

// login api to enter the user to vote the caditates

// func LoginAPi(c *gin.Context) {
// 	err := c.BindJSON(&user)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tempEmail := GetAllEmail("user")
// 	tempPassword := GetAllPassword(&user)

// 	if contains(tempEmail, user.Email) {
// 		if contains(tempPassword, user.Password) {
// 			c.JSON(200, gin.H{
// 				"message": "Successfully loged in",
// 			})
// 		} else {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": "Password is incorrect",
// 			})
// 		}
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Email is not register please go and register",
// 		})
// 	}

// }

// dashboard page
func Dashboad(c *gin.Context) {
	c.HTML(200, "dashboad.html", gin.H{
		"title": "DashBoard page",
	})
}

// voter page to display all voters
func Voter(c *gin.Context) {

	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `SELECT id,firstname,lastname,email,verified from voter  ALLOW FILTERING`
	voter, err := newcontroller.GetVoter(query)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(voter)
	c.HTML(200, "voters.html", gin.H{
		"title": "Voter Page",
		"voter": voter,
	})
}

// function used to generate uuid
func GenerateUUID() gocql.UUID {
	return gocql.TimeUUID()
}

/*
api to add to the voter to the database so that he can vote to any candidates
gocl.uuid which is generate by the function called generateUUID
all the values are take from frontend through json object
*/
func AddVoterApi(c *gin.Context) {
	vote := c.BindJSON(&voter)
	if vote != nil {
		log.Fatal(vote)
	}
	tempEmail := GetAllEmail("voter")
	if contains(tempEmail, voter.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exits",
		})
	} else {

		addVote(&voter)
		fmt.Println("added")
		c.JSON(200, gin.H{
			"message": "Voter added succfully",
		})
	}

}

/*
function to add vote by generating unique id
*/
func addVote(vote *model.Voter) {
	vote.ID = GenerateUUID()
	vote.Verified = false
	vote.Voted = false
	hash, err1 := bcrypt.GenerateFromPassword([]byte(vote.Password), 10)
	if err1 != nil {
		log.Fatal(err1)

	}
	vote.Password = string(hash)
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO voter (id,firstname,lastname,email,password,verified,voted) VALUES (?,?,?,?,?,?,?)`
	newcontroller.ExcuteQuery(query, vote.ID, vote.FirstName, vote.LastName, vote.Email, vote.Password, vote.Verified, vote.Voted)
}

//api to delete the row from the voter

func DeleteVote(c *gin.Context) {
	id := c.Param("id")

	deleteVoterRecord(id)

	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})

}

// function to delete the record
func deleteVoterRecord(ID string) {
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}

	query := `DELETE from voter where id = ? `
	newcontroller.ExcuteQuery(query, ID)

}

// function to all the position name
func GetAllPosition(user string) []string {
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `SELECT name from ` + user + ` ALLOW FILTERING`
	data, err := newcontroller.GetData(query)
	if err != nil {
		log.Fatal(err)
	}

	return data

}

// api to add position
func AddPositionApi(c *gin.Context) {
	vote := c.BindJSON(&position)
	if vote != nil {
		log.Fatal(vote)
	}
	tempPositon := GetAllPosition("position")
	if contains(tempPositon, position.Name) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Position  name already exits",
		})
	} else {

		addPosition(&position)
		// fmt.Println("added")
		c.JSON(200, gin.H{
			"message": "Position added succfully",
		})
	}

}

// fuction to add position in te voting system
func addPosition(position *model.Position) {
	position.ID = GenerateUUID()
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO position (id,name,priority) VALUES (?,?,?)`
	newcontroller.ExcuteQuery(query, position.ID, position.Name, position.Priority)
}

func Position(c *gin.Context) {

	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `SELECT id,name,priority from position  ALLOW FILTERING`
	position, err := newcontroller.GetPosition(query)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "position.html", gin.H{
		"title":    "position page",
		"position": position,
	})
}

//api to delete the row from the voter

func DeletePosition(c *gin.Context) {
	id := c.Param("id")

	deletePositionRecord(id)

	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})

}

// function to delete the record
func deletePositionRecord(ID string) {
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}

	query := `DELETE from position where id = ?`
	newcontroller.ExcuteQuery(query, ID)

}

// function to all the candidates name
func GetAllCandidates(user string) []string {
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `SELECT fullname from ` + user + ` ALLOW FILTERING`
	data, err := newcontroller.GetData(query)
	if err != nil {
		log.Fatal(err)
	}

	return data

}

// api to add Candidates
func AddCandidateApi(c *gin.Context) {
	vote := c.BindJSON(&candidate)

	if vote != nil {
		log.Fatal(vote)
	}

	tempCandidates := GetAllCandidates("candidate")

	if contains(tempCandidates, candidate.FullName) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Candidate name already exits",
		})
	} else {

		addCandidates(&candidate)
		// fmt.Println("added")
		c.JSON(200, gin.H{
			"message": "Candidates added succfully",
		})
	}

}

// fuction to add position in te voting system
func addCandidates(candidate *model.Candidate) {
	candidate.ID = GenerateUUID()

	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO candidate (id,bio,fullname,position) VALUES (?,?,?,?)`
	newcontroller.ExcuteQuery(query, candidate.ID, candidate.Bio, candidate.FullName, candidate.Position)
}

func Candidate(c *gin.Context) {

	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `SELECT id,name,priority from position  ALLOW FILTERING`
	query1 := `SELECT id,fullname,bio,position from candidate  ALLOW FILTERING`
	position, err1 := newcontroller.GetPosition(query)
	candidate, err := newcontroller.GetCandidate(query1)
	if err1 != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "candidates.html", gin.H{
		"title":     "candidate page",
		"position":  position,
		"candidate": candidate,
	})
}

func DeleteCandidate(c *gin.Context) {
	id := c.Param("id")

	deleteCandidateRecord(id)

	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})

}

// function to delete the record
func deleteCandidateRecord(ID string) {
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}

	query := `DELETE from candidate where id = ?`
	newcontroller.ExcuteQuery(query, ID)

}

func Ballot(c *gin.Context) {

	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}

	query1 := `SELECT id,fullname,bio,position from candidate  ALLOW FILTERING`

	candidate, err := newcontroller.GetCandidate(query1)

	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "ballot.html", gin.H{
		"title": "Ballot page",

		"candidate": candidate,
	})
}

func SignupPage(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash Password",
		})
		return

	}
	temp := GetAllEmail("users")
	if contains(temp, body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exits"})
	} else {
		user := model.Users{Email: body.Email, Password: string(hash)}
		addUser(&user)
		c.JSON(200, gin.H{
			"message": "success",
		})

	}
}

func LoginAPi(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	tempEmail := GetAllEmail("users")
	// fmt.Println(body.Email)
	user, err1 := newcontroller.GetUserByEmail(body.Email)
	// fmt.Println(user)
	if err != nil {
		// fmt.Println("prasanna")
		log.Fatal(err1)
	}
	if contains(tempEmail, body.Email) {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Password",
			})
			return
		}

		tokenString, err := generateToken(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create token",
			})
			return
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

		c.JSON(200, gin.H{
			"message": "Login successfully",
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invaalid email or password",
		})
	}

}

func generateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECERET")))
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"message": user,
	})
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorisation", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{
		"message": "logout successfully",
	})
}

func LoginAPiVoter(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	tempEmail := GetAllEmail("voter")
	// fmt.Println(body.Email)
	voter, err1 := newcontroller.GetVoterByEmail(body.Email)
	// fmt.Println(user)
	if err != nil {
		// fmt.Println("prasanna")
		log.Fatal(err1)
	}
	if contains(tempEmail, body.Email) {
		err := bcrypt.CompareHashAndPassword([]byte(voter.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Password",
			})
			return
		}

		tokenString, err := generateVoterToken(voter.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create token",
			})
			return
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorisation", tokenString, 3600*24*30, "", "", false, true)

		c.JSON(200, gin.H{
			"message": "Login successfully",
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invaalid email or password",
		})
	}

}

func generateVoterToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECERET")))
}

func Home(c *gin.Context) {
	voter, _ := c.Get("voter")
	c.HTML(200, "home.html", gin.H{
		"voter": voter,
	})
}

func Profile(c *gin.Context) {
	voter, _ := c.Get("voter")
	c.HTML(200, "profile.html", gin.H{
		"title": "Profile page",
		"voter": voter,
	})
}

func VoterBallot(c *gin.Context) {
	voter, _ := c.Get("voter")
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}

	query1 := `SELECT id,fullname,bio,position from candidate  ALLOW FILTERING`

	candidate, err := newcontroller.GetCandidate(query1)

	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "voterballot.html", gin.H{
		"title":     "Voter Ballot page",
		"voter":     voter,
		"candidate": candidate,
	})
}
func VoterVerify(c *gin.Context) {
	voter, _ := c.Get("voter")

	c.HTML(200, "verify.html", gin.H{
		"title": "Voter Verify page",
		"voter": voter,
	})
}

func VoterVerifyApi(c *gin.Context) {
	var voters struct {
		id       gocql.UUID
		email    string
		verified bool
	}
	// voter, _ := c.Get("voter")
	if c.BindJSON(&voters) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cant read body",
		})
	}
	// fmt.Println(voter)

	// vote.verified = true
	newcontroller, err := database.NewMyController()
	if err != nil {
		log.Fatal(err)
	}
	query := `UPDATE voter set verified = ? where id = ? and email = ?`
	newcontroller.ExcuteQuery(query, voters.verified, voters.id, voters.email)
	c.JSON(200, gin.H{
		"message": "verified successfully",
	})

}
