package model

type Model struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

// 任务
type Job struct {
	Model
	Name        string `gorm:"size:128;not null;" json:"name" form:"name"`                         // 任务名
	Type   	    int    `gorm:"int;not null" json:"type" form:"type"`                               // 类型
	Status      int    `gorm:"int;not null" json:"status" form:"status"`                           // 状态
	CreateTime  int64  `json:"createTime" form:"createTime"`                                       // 创建时间
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`                                       // 更新时间
}