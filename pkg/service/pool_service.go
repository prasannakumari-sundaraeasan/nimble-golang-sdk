// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Pool Service - Manage pools. Pools are an aggregation of arrays.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// PoolService type
type PoolService struct {
	objectSet *client.PoolObjectSet
}

// NewPoolService - method to initialize "PoolService"
func NewPoolService(gs *NsGroupService) *PoolService {
	objectSet := gs.client.GetPoolObjectSet()
	return &PoolService{objectSet: objectSet}
}

// GetPools - method returns a array of pointers of type "Pools"
func (svc *PoolService) GetPools(params *util.GetParams) ([]*model.Pool, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	poolResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return poolResp, nil
}

// CreatePool - method creates a "Pool"
func (svc *PoolService) CreatePool(obj *model.Pool) (*model.Pool, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	poolResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return poolResp, nil
}

// UpdatePool - method modifies  the "Pool"
func (svc *PoolService) UpdatePool(id string, obj *model.Pool) (*model.Pool, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	poolResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return poolResp, nil
}

// GetPoolById - method returns a pointer to "Pool"
func (svc *PoolService) GetPoolById(id string) (*model.Pool, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	poolResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return poolResp, nil
}

// GetPoolByName - method returns a pointer "Pool"
func (svc *PoolService) GetPoolByName(name string) (*model.Pool, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	poolResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(poolResp) == 0 {
		return nil, nil
	}

	return poolResp[0], nil
}

// DeletePool - deletes the "Pool"
func (svc *PoolService) DeletePool(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
