package customer

import "github.com/cristianpontes/teamwork-go-challenge/pkg/mail"

type Customers []*Customer

type Customer struct {
	FirstName string     `json:"first_name" csv:"first_name"`
	LastName  string     `json:"last_name" csv:"last_name"`
	Email     mail.Email `json:"email" csv:"email"`
	Gender    string     `json:"gender" csv:"gender"`
	IPAddress string     `json:"ip_address" csv:"ip_address"`
}
