// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAccount = "accounts"

// Account mapped from table <accounts>
type Account struct {
	ID       string `gorm:"column:id;primaryKey" json:"id"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Password string `gorm:"column:password" json:"password"`
	Salt     string `gorm:"column:salt" json:"salt"`
}

// TableName Account's table name
func (*Account) TableName() string {
	return TableNameAccount
}