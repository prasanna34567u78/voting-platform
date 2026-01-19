package database

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/onlinevoting/pkg/model"
)

type MyController struct {
	session *gocql.Session
}

func NewMyController() (*MyController, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = "voting_app"
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &MyController{session}, nil
}

func (c *MyController) ExcuteQuery(query string, values ...interface{}) {
	if err := c.session.Query(query).Bind(values...).Exec(); err != nil {
		log.Fatal(err)

	}
}

func (c *MyController) Close() {
	c.session.Close()
}

func (c *MyController) GetData(query string, values ...interface{}) ([]string, error) {

	iter := c.session.Query(query, values...).Iter()
	var result []string

	var column1Type string

	for iter.Scan(&column1Type) {
		result = append(result, column1Type)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

// function to get all the voter from the database
func (c *MyController) GetVoter(query string, values ...interface{}) ([]model.Voter, error) {

	iter := c.session.Query(query, values...).Iter()
	var result []model.Voter

	var firstname, lastName, email string
	var verified bool
	var id gocql.UUID
	for iter.Scan(&id, &firstname, &lastName, &email, &verified) {
		newVote := model.Voter{
			ID:        id,
			FirstName: firstname,
			LastName:  lastName,
			Email:     email,
			Verified:  verified,
		}
		result = append(result, newVote)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

// function to get all the Position from the database
func (c *MyController) GetPosition(query string, values ...interface{}) ([]model.Position, error) {

	iter := c.session.Query(query, values...).Iter()
	var result []model.Position

	var name string
	var priority int
	var id gocql.UUID
	for iter.Scan(&id, &name, &priority) {
		newVote := model.Position{
			ID:       id,
			Name:     name,
			Priority: priority,
		}
		result = append(result, newVote)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *MyController) GetCandidate(query string, values ...interface{}) ([]model.Candidate, error) {

	iter := c.session.Query(query, values...).Iter()
	var result []model.Candidate

	var fullnameme, bio, position string

	var id gocql.UUID
	for iter.Scan(&id, &fullnameme, &bio, &position) {
		newVote := model.Candidate{
			ID:       id,
			FullName: fullnameme,
			Bio:      bio,
			Position: position,
		}
		result = append(result, newVote)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *MyController) GetUserByID(ID gocql.UUID) (*model.Voter, error) {

	var result model.Voter
	iter := c.session.Query(`SELECT id, email, password FROM voter WHERE id = ?`, ID).Scan(&result.ID, &result.Email, &result.Password)
	if iter != nil {
		log.Fatal(iter)
	}
	return &result, nil
}

func (c *MyController) ValidateCredentials(username, password string) bool {

	var Password string
	if err := c.session.Query(`SELECT password FROM voter WHERE email =? ALLOW FILTERING`, username).Scan(&Password); err != nil {
		// fmt.Println("P")
		return false
	}

	// err := bcrypt.CompareHashAndPassword(Password, password)
	// return err == nil
	return password == Password
	// return false
}
func (c *MyController) GetByEmail(username string) gocql.UUID {

	var id gocql.UUID
	if err := c.session.Query(`SELECT id FROM voter WHERE email =? ALLOW FILTERING`, username).Scan(&id); err != nil {
		// fmt.Println("P")
		log.Fatal(err)
	}

	// err := bcrypt.CompareHashAndPassword(Password, password)
	// return err == nil
	return id
	// return false
}

func (c *MyController) GetUserByEmail(email string) (*model.Users, error) {

	var result model.Users
	iter := c.session.Query(`SELECT  email,password FROM  users  WHERE email = ? ALLOW FILTERING`, email).Scan(&result.Email, &result.Password)
	if iter != nil {
		log.Fatal(iter)
	}
	return &result, nil
}

func (c *MyController) GetVoterByEmail(email string) (*model.Voter, error) {

	var result model.Voter
	iter := c.session.Query(`SELECT id,firstname,lastname, email, password,verified,voted FROM voter WHERE email = ? ALLOW FILTERING`, email).Scan(&result.ID, &result.FirstName, &result.LastName, &result.Email, &result.Password, &result.Verified, &result.Verified)
	if iter != nil {
		log.Fatal(iter)
	}
	return &result, nil
}

func (c *MyController) GetALLVoter(query string, values ...interface{}) ([]model.Voter, error) {

	iter := c.session.Query(query, values...).Iter()
	var result []model.Voter

	var firstname, lastName, email string
	var verified, voted bool
	var id gocql.UUID
	for iter.Scan(&id, &firstname, &lastName, &email, &verified, &voted) {
		newVote := model.Voter{
			ID:        id,
			FirstName: firstname,
			LastName:  lastName,
			Email:     email,
			Verified:  verified,
			Voted:     voted,
		}
		result = append(result, newVote)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return result, nil
}
