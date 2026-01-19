package model

import "github.com/gocql/gocql"

// Structure  for admin panel
//admin can only register through database

type Users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*voter struct includes unique ID that will be primary key in database,
firtName,lastName ,email,password of type string
verified and voted are of type bool
verified is to check the voter is authorised or not
voted is to check
*/
type Voter struct {
	ID        gocql.UUID `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Verified  bool       `json:"verified"`
	Voted     bool       `json:"voted"`
}

/* Position struct include position which is string priority of int to display in
ballotposition*/

type Position struct {
	ID       gocql.UUID
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

/* Candidate Struct includes caditates full name of type String
photo in term of byte, Bio type is string to store the bio of
the candidate Position for which position he is participating
*/

type Candidate struct {
	ID       gocql.UUID
	FullName string `json:"fullname"`
	Bio      string `json:"bio"`
	Position string `json:"position"`
}

//votes table to get all the voter whom they are voted and for which position

type Votes struct {
	Voter     gocql.UUID `json:"voterid"`
	Position  string     `json:"position"`
	Candidate string     `json:"candidate"`
}
