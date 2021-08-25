package model

import (
	api "server/base-server/api/v1"
	"server/common/dao"
	"time"
)

type BillingOwner struct {
	dao.Model
	OwnerId   string               `gorm:"primaryKey;type:varchar(100);not null;default:'';comment:归属用户id"`
	OwnerType api.BillingOwnerType `gorm:"primaryKey;type:tinyint;not null;default:0;comment:归属用户类型1space 2user"`
	Amount    float64              `gorm:"type:decimal(10,2);not null;default:0;comment:机时数量"`
}

func (BillingOwner) TableName() string {
	return "billing_owner"
}

type BillingPayRecord struct {
	dao.Model
	Id        string                     `gorm:"primaryKey;type:varchar(100);not null;comment:id"`
	OwnerId   string                     `gorm:"type:varchar(100);not null;default:'';uniqueIndex:biz;comment:归属用户id"`
	OwnerType api.BillingOwnerType       `gorm:"type:tinyint;not null;default:0;uniqueIndex:biz;comment:归属用户类型1space 2user"`
	BizId     string                     `gorm:"type:varchar(100);not null;default:'';uniqueIndex:biz;comment:业务关联id"`
	BizType   api.BillingBizType         `gorm:"type:tinyint;not null;default:0;uniqueIndex:biz;comment:业务类型 1trainJob 2notebook"`
	Amount    float64                    `gorm:"type:decimal(10,2);not null;default:0;comment:结算机时"`
	Title     string                     `gorm:"type:varchar(100);not null;default:'';comment:标题"`
	StartedAt *time.Time                 `gorm:"type:datetime(3);comment:计费起始时间"`
	EndedAt   *time.Time                 `gorm:"type:datetime(3);comment:计费截止时间"`
	Status    api.BillingPayRecordStatus `gorm:"type:tinyint;not null;default:1;comment:计费状态 1计费中 2计费完成"`
}

func (BillingPayRecord) TableName() string {
	return "billing_pay_record"
}

type BillingRechargeRecord struct {
	dao.Model
	Id        string               `gorm:"primaryKey;type:varchar(100);not null;comment:id"`
	OwnerId   string               `gorm:"type:varchar(100);not null;index:ownerId;comment:归属用户id"`
	OwnerType api.BillingOwnerType `gorm:"type:tinyint;not null;default:0;comment:归属用户类型1space 2user"`
	Amount    float64              `gorm:"type:decimal(10,2);not null;default:0;comment:充值机时"`
}

func (BillingRechargeRecord) TableName() string {
	return "billing_recharge_record"
}

type BillingPayRecordKey struct {
	OwnerId   string
	OwnerType api.BillingOwnerType
	BizId     string
	BizType   api.BillingBizType
}

type BillingOwnerKey struct {
	OwnerId   string
	OwnerType api.BillingOwnerType
}

type BillingPayRecordQuery struct {
	PageIndex int
	PageSize  int
	SortBy    string
	OrderBy   string
	OwnerId   string
	OwnerType api.BillingOwnerType
	SearchKey string
}

type BillingRechargeRecordQuery struct {
	PageIndex int
	PageSize  int
	SortBy    string
	OrderBy   string
	OwnerId   string
	OwnerType api.BillingOwnerType
}

type BillingOwnerQuery struct {
	PageIndex int
	PageSize  int
	SortBy    string
	OrderBy   string
	OwnerId   string
	OwnerType api.BillingOwnerType
}
