package model

import (
	"context"
	"time"
)

const (
	TASK_STATUS_CREATED = iota + 1
	TASK_STATUS_RUNNING
	TASK_STATUS_FINISHED
)

var TableTask _TableTask

type _TableTask struct{}

type TableTaskModel struct {
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

func (TableTaskModel) TableName() string {
	return "table_task"
}

func (*_TableTask) GetTaskBySrcPodUID(ctx context.Context, srcPodUID string) ([]*TableTaskModel, error) {
	var tasks []*TableTaskModel
	err := DB.NewRequest(ctx).Where("src_pod_uid = ? and task_status = ?", srcPodUID, TASK_STATUS_CREATED).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (*_TableTask) BatchCreate(ctx context.Context, data []*TableTaskModel) error {
	err := DB.NewRequest(ctx).CreateInBatches(data, len(data)).Error
	if err != nil {
		return err
	}
	return nil
}

func (*_TableTask) UpdateByTaskID(ctx context.Context, data *TableTaskModel) error {
	err := DB.NewRequest(ctx).Where("task_id = ?", data.TaskID).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (*_TableTask) GetTaskByTaskID(ctx context.Context, taskID string) (*TableTaskModel, error) {
	var task TableTaskModel
	err := DB.NewRequest(ctx).Where("task_id = ?", taskID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}
