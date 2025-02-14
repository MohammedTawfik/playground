// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTransfer = "Transfers"

// Transfer mapped from table <Transfers>
type Transfer struct {
	TransferID    int64     `gorm:"column:TransferId;primaryKey;default:nextval('"Transfers_TransferId_seq"'::regclass)" json:"TransferId"`
	FromAccountID int64     `gorm:"column:FromAccountId;not null" json:"FromAccountId"`
	ToAccountID   int64     `gorm:"column:ToAccountId;not null" json:"ToAccountId"`
	Amount        int64     `gorm:"column:Amount;not null;comment:Amount must be positive" json:"Amount"` // Amount must be positive
	CreatedAt     time.Time `gorm:"column:Created_At;not null;default:2023-09-25 15:33:04.536172+00" json:"Created_At"`
}

// TableName Transfer's table name
func (*Transfer) TableName() string {
	return TableNameTransfer
}
