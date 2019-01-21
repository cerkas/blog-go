
package root
import (
	"blog-go/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id           string  `json:"id"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
}
type userModel struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Username     string
	Password     string
}
type UserService struct {
	collection *mgo.Collection
}
func(p UserService) Create(u *User) error {
	user := newUserModel(u)
	return p.collection.Insert(&user)
}
func newUserModel(u *User) *userModel {
	return &userModel{
		Username: u.Username,
		Password: u.Password }
}
func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}
func NewUserService(session *config.Session, dbName string, collectionName string) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService {collection}
}