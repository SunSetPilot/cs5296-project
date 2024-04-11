package dal

import (
	"context"

	"github.com/SunSetPilot/cs5296-project/model"
	"github.com/SunSetPilot/cs5296-project/model/table"
)

var TableTask _TableTask

type _TableTask struct{}

func (*_TableTask) GetTaskBySrcPodUID(ctx context.Context, srcPodUID string) ([]*table.TaskModel, error) {
	var tasks []*table.TaskModel
	err := DB.NewRequest(ctx).Where("src_pod_uid = ? and task_status = ?", srcPodUID, model.TASK_STATUS_CREATED).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (*_TableTask) BatchCreate(ctx context.Context, data []*table.TaskModel) error {
	err := DB.NewRequest(ctx).CreateInBatches(data, len(data)).Error
	if err != nil {
		return err
	}
	return nil
}

func (*_TableTask) UpdateByTaskID(ctx context.Context, data *table.TaskModel) error {
	err := DB.NewRequest(ctx).Where("task_id = ?", data.TaskID).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (*_TableTask) GetTaskByTaskID(ctx context.Context, taskID string) (*table.TaskModel, error) {
	var task table.TaskModel
	err := DB.NewRequest(ctx).Where("task_id = ?", taskID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}
