package table

import "time"

type ClientModel struct {
	ID           uint32    `gorm:"column:id;type:int(11) unsigned;not null;primary_key;auto_increment;comment:'主键'"`
	PodName      string    `gorm:"column:pod_name;type:varchar(255);not null;default:'';comment:'Pod名称'"`
	PodUID       string    `gorm:"column:pod_uid;type:varchar(255);not null;default:'';comment:'Pod唯一标识'"`
	PodIP        string    `gorm:"column:pod_ip;type:varchar(255);not null;default:'';comment:'Pod IP地址'"`
	NodeName     string    `gorm:"column:node_name;type:varchar(255);not null;default:'';comment:'Pod所在节点名称'"`
	NodeIP       string    `gorm:"column:node_ip;type:varchar(255);not null;default:'';comment:'Pod所在节点IP地址'"`
	ClientStatus uint8     `gorm:"column:client_status;type:tinyint(3) unsigned;not null;default:1;comment:'客户端状态 1空闲 2忙碌 3下线'"`
	RegisterTime time.Time `gorm:"column:register_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'首次上线时间'"`
	UpdateTime   time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'记录更新时间'"`
}

func (ClientModel) TableName() string {
	return "table_client"
}
