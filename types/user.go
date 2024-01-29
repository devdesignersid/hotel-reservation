package types

type User struct {
	Id        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"LastName" json:"lastName"`
}
