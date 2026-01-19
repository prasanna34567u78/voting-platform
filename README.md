# E-online voting system 
This E - Voting System Was Developed With golan(gin Framework) and cassandra database. 


This Voting System web application built using golang can serve as the automated voting system for organizations and/or institutions. The system works like the common election manual system of voting whereas the system must be populated by the list of the positions, candidates, and voters. The algomox E-voting system can help a certain organization like algomox or school to minimize the voting time duration because aside providing the voters an online platform to vote, the system will automatically count the votes for each candidate. The system has 2 sides of the user interface which are the administrator and voters side. The admin user is in charge to populate and manage the data of the system and the voter side which is where the voters will choose their candidate and submit their votes.

## Features:
- [x] setting cookies
- [x] jwt-auth
- [x] CRUD voters
- [x] CRUD candidates
- [x] CRUD positions
- [x]storing password in hash
- [x] all operation done using REST FULL API
- [x] Middleware to authicate users

### A. Admin Users Can
1. Manage Voters 
2. Manage Candidates 
3. Manage Positions 
4. Update/Change Ballot Title
5.Admin should register using api call only
api= http://localhoast:8000/signup
go to postman past this above api set method post in the body
pass
{
    "email":"youremail@gmail.com",
    "password":"yourpassword"
}
note : first you should run server 

### B. Voters Can
1. Register through admin only he cant regsitered by him self
2. Login
3. Verify with id 
4. Votes for their favourite candidates
5. View candidates they voted for




### Pre-Requisites:
1.Install golang in window or mac 
 [https://go.dev/doc/install]

2.Cassandra use image in docker 
 [https://docs.docker.com/desktop/install/windows-install/]

### commads to pull the cassandra image from docker engine 
open terminal
1.To pull image from docker
 [docker pull cassandra]

2.Run conatiner using image
 [docker run -p 7000:7000 -p 7001:7001 -p 7199:7199 -p 9042:9042 -p 9106:9106 --name cassandra -d cassandra:latest ]

3.check container is running or not
 [docker ps]

4.Get inside container (you will one id of container plz pasete below)
 [docker exec -it pastehere bash]

5.Enter into cql
 [cqlsh]

6.create key space 
[ create keyspace voting_app with replication = {'class':'SimpleStrategy' ,'replication_factor' :1}; ]

7.To check keyspace is create or not 
 [desc keyspaces]

8.if the key is done come to my folder open file
file path to change the keyspace if you want
 [ onlinevoting->pkg->DatabaseConnection.go ]
 change keyspace in that mycontroller if you want


### Installation
**1. Create a Folder where you want to save the project**
## Install nessary packes open terminal install one by one  
 [go mod init "github.com/onlinevoting"]
 [go get "golang.org/x/crypto"]
 [go get "golang.org/x/net"]
 [go get "github.com/joho/godotenv"]
 [go get "github.com/gocql/gocql" ]
 [go get "github.com/go-chi/chi"]
 [go get "github.com/gin-gonic/gin"]
 [go get "github.com/dgrijalva/jwt-go"]
 [go get "github.com/gin-contrib/sessions" ]

 **2. open the project folder**
1.Initializer foler contains inital database setup and loading env variable 
2.pkg floder contains all required folder like
 -model
 -(users)controller
 -database
 -middleware
 -router
3.template contains all the forntend part
 -html,css,js
 -js calling api

**3. run the project**
1.open terminal and paste it will automatically run and build every time changes
[compiledaemon --command="./onlinevoting"]

2.open browers
 [http://loacalhost:8000/login]
to login there are two choice admin or voter if your voter you should ask that 
from admin for mail and password

for admin login i alread mention above he can sigin using api only for securty
and password are stored in hash value .

**4. Data base model**

//user model or admin
type Users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//voter model
type Voter struct {
	ID        gocql.UUID `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Verified  bool       `json:"verified"`
	Voted     bool       `json:"voted"`
}

//position model
type Position struct {
	ID       gocql.UUID
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

//candidates model

type Candidate struct {
	ID       gocql.UUID
	FullName string `json:"fullname"`
	Bio      string `json:"bio"`
	Position string `json:"position"`
}

//votes model
type Votes struct {
	Voter     gocql.UUID `json:"voterid"`
	Position  string     `json:"position"`
	Candidate string     `json:"candidate"`
}


## How the system works
Administrator is required to have created candidates. 
Before creating candidates, the admin must have created positions
After doing this, the voters can vote (provided that they are registered and verified)


## How do voters get verified ?
UUID  is generated for different voter. 
After login as voter he/she can verified and then he/she can vote


## Enquire
1.Email:prasannakumarhm078@gmailcom
2.phone No:6361477267
3.prasannakumar


## thankyou 
for algomox hiring team for giving this oppturinity doing this i Gained more knowledge and had greate expreience .if there is any mistake or if i forget to do you can suggest me i will do
