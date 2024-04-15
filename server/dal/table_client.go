package dal

import (
	"context"
	"time"

	"gorm.io/gorm/clause"

	"github.com/SunSetPilot/cs5296-project/model"
	"github.com/SunSetPilot/cs5296-project/model/table"
)

var TableClient _TableClient

type _TableClient struct{}

func (*_TableClient) CreateOrUpdate(ctx context.Context, data *table.ClientModel) error {
	err := DB.NewRequest(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "pod_uid"}},
		DoUpdates: clause.AssignmentColumns([]string{"pod_name", "pod_ip", "node_name", "node_ip", "client_status", "update_time"}),
	}).Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (*_TableClient) GetOnlineClientList(ctx context.Context) ([]*table.ClientModel, error) {
	var list []*table.ClientModel
	err := DB.NewRequest(ctx).Where("client_status <> ?", model.CLIENT_STATUS_OFFLINE).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (*_TableClient) UpdateOfflineClients(ctx context.Context, threshold int) error {
	clientModel := &table.ClientModel{
		ClientStatus: model.CLIENT_STATUS_OFFLINE,
		UpdateTime:   time.Now(),
	}
	expireTime := clientModel.UpdateTime.Add(-time.Second * time.Duration(threshold))
	err := DB.NewRequest(ctx).
		Where("update_time < ? AND client_status <> ?", expireTime, model.CLIENT_STATUS_OFFLINE).
		Updates(clientModel).Error
	if err != nil {
		return err
	}
	return nil
}
