package app

import (
	"mattermost-server/model"
	"mattermost-server/store/sqlstore"
)

func (a *App) CreateCompany(org *model.Org) (map[string]string, *model.AppError) {
	result, err := sqlstore.RedisStoreSaveCompany(org)
	return result, err
}
func (a *App) GetCompany(key string) (map[string]string, *model.AppError) {
	result, err := sqlstore.RedisStoreGetCompany(key)
	return result, err

}
func (a *App) CreateOrg(parent string, name string) (map[string]string, *model.AppError) {
	result, err := sqlstore.RedisStoreSaveOrg(parent, name)
	return result, err
}
func (a *App) GetOrgNode(key string) (map[string]string, *model.AppError) {
	result, err := sqlstore.RedisStoreGetOrgNode(key)
	return result, err
}
func (a *App) DelOrg(key string) (map[string]string, *model.AppError) {
	result, err := sqlstore.RedisStoreDelOrgNode(key)
	return result, err
}
func (a *App) UpdateOrg(key string, name string) (map[string]string, *model.AppError) {
	result, err := sqlstore.RedisStoreUpdateOrg(key, name)
	return result, err
}
