// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAccount = "Accounts"

// Account mapped from table <Accounts>
type Account struct {
	AccountID    int64     `gorm:"column:AccountId;primaryKey;default:nextval('"Accounts_AccountId_seq"'::regclass)" json:"AccountId"`
	AccountOwner string    `gorm:"column:AccountOwner;not null" json:"AccountOwner"`
	Balance      int32     `gorm:"column:Balance;not null" json:"Balance"`
	Currency     string    `gorm:"column:Currency;not null" json:"Currency"`
	CreatedAt    time.Time `gorm:"column:Created_At;not null;default:2023-09-25 15:33:04.536172+00" json:"Created_At"`
}

// TableName Account's table name
func (*Account) TableName() string {
	return TableNameAccount
}
