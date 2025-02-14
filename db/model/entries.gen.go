// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEntry = "Entries"

// Entry mapped from table <Entries>
type Entry struct {
	EntryID   int64     `gorm:"column:EntryId;primaryKey;default:nextval('"Entries_EntryId_seq"'::regclass)" json:"EntryId"`
	AccountID int64     `gorm:"column:AccountId;not null" json:"AccountId"`
	Amount    int64     `gorm:"column:Amount;not null;comment:Amount can be positive or negative" json:"Amount"` // Amount can be positive or negative
	CreatedAt time.Time `gorm:"column:Created_At;not null;default:2023-09-25 15:33:04.536172+00" json:"Created_At"`
}

// TableName Entry's table name
func (*Entry) TableName() string {
	return TableNameEntry
}
