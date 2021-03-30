package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"study.go/gorm_demo/global"
)

/*type User struct {
	ID       int              `gorm:"column:id;primary_key" json:"id"` //
	Username string           `gorm:"column:username" json:"username"` //
	Openid   string           `gorm:"column:openid" json:"openid"`     //
	Head     string           `gorm:"column:head" json:"head"`         //
	Email    string           `gorm:"column:email" json:"email"`       //
	Fire     int              `gorm:"column:fire" json:"fire"`         //
	Account  *simplejson.JSON `gorm:"account" json:"account"`
}

func (u *User) TableName() string {
	return "user"
}*/

/*type JSON struct {
	Data interface{}
}

// Value returns json value, implement driver.Valuer interface.
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j.Data)
}
*/
type M map[string]interface{}

type Li []M

func (l Li) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func main() {

	var ml Li
	ml = append(ml, M{
		"name": "zhong",
		"age":  1,
	})

	err := global.DB.Exec("update user set account = ? where id = ?", ml, 1).Error
	fmt.Println(err)

}
