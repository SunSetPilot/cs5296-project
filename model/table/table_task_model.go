package table

import "time"

type TaskModel struct {
	ID         uint32    `gorm:"column:id;type:int(11) unsigned;not null;primary_key;auto_increment;comment:'主键'"`
	TaskID     string    `gorm:"column:task_id;type:varchar(255);not null;default:'';comment:'任务ID'"`
	SrcPodUID  string    `gorm:"column:src_pod_uid;type:varchar(255);not null;default:'';comment:'源Pod唯一标识'"`
	SrcPodIP   string    `gorm:"column:src_pod_ip;type:varchar(255);not null;default:'';comment:'源Pod IP地址'"`
	DstPodUID  string    `gorm:"column:dst_pod_uid;type:varchar(255);not null;default:'';comment:'目标Pod唯一标识'"`
	DstPodIP   string    `gorm:"column:dst_pod_ip;type:varchar(255);not null;default:'';comment:'目标Pod IP地址'"`
	TaskParam  string    `gorm:"column:task_param;type:varchar(255);not null;default:'';comment:'任务参数'"`
	TaskType   string    `gorm:"column:task_type;type:varchar(255);not null;default:'';comment:'任务类型'"`
	TaskStatus uint8     `gorm:"column:task_status;type:tinyint(3) unsigned;not null;default:1;comment:'任务状态 1已创建 2运行中 3已结束'"`
	TaskResult string    `gorm:"column:task_result;type:longtext;comment:'任务结果'"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'任务创建时间'"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'任务更新时间'"`
}

func (TaskModel) TableName() string {
	return "table_task"
}
