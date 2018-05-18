package sqlstore

import (
	"fmt"
	"mattermost-server/model"
	"net/http"
	"strings"

	l4g "github.com/alecthomas/log4go"
	"github.com/gomodule/redigo/redis"
)

func RedisStoreSaveCompany(org *model.Org) (map[string]string, *model.AppError) {
	var result = make(map[string]string)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, model.NewAppError("RedisStoreSaveCompany", "api.org.savecompany.app_error", nil, "conn redis fail", http.StatusBadRequest)
	}
	defer c.Close()
	_, err2 := redis.String(c.Do("set", org.Key, org.Name))
	if err2 != nil {
		l4g.Debug("err2=", err2.Error())
		return nil, model.NewAppError("RedisStoreSaveCompany", "api.org.savecompany.app_error", nil, "set company fail", http.StatusBadRequest)
	}
	result[org.Key] = org.Name
	return result, nil
}

func RedisStoreGetCompany(key string) (map[string]string, *model.AppError) {
	result := make(map[string]string)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, model.NewAppError("RedisStoreGetCompany", "api.org.getcompany.app_error", nil, "conn redis fail", http.StatusBadRequest)
	}
	defer c.Close()
	name, err2 := redis.String(c.Do("get", key))
	if err2 != nil {
		return nil, model.NewAppError("RedisStoreGetCompany", "api.org.getcompany.app_error", nil, "there is no value in redis", http.StatusBadRequest)
	}
	result[key] = name
	fmt.Printf("get value :%s\n", name)
	return result, nil
}

func RedisStoreSaveOrg(parent string, name string) (map[string]string, *model.AppError) {

	var result = make(map[string]string)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, model.NewAppError("RedisStoreSaveOrg", "api.org.saveorg.app_error", nil, "conn redis fail", http.StatusBadRequest)
	}
	defer c.Close()
	l4g.Debug("parent", parent)
	l4g.Debug("name=", name)
	value, err2 := redis.String(c.Do("get", parent))
	if err2 != nil || value == "" {
		return nil, model.NewAppError("RedisStoreSaveOrg", "api.org.saveorg.app_error", nil, "org parent is not exist", http.StatusBadRequest)
	} else {
		_, err3 := redis.String(c.Do("set", parent+":"+name, name))
		if err3 != nil {
			return nil, model.NewAppError("RedisStoreSaveOrg", "api.org.saveorg.app_error", nil, "set org is fail", http.StatusBadRequest)
		}
		result[parent+":"+name] = name
		return result, nil
	}

}
func RedisStoreGetOrgNode(key string) (map[string]string, *model.AppError) {
	var result = make(map[string]string)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, model.NewAppError("RedisStoreGetOrgNode", "api.org.saveorg.app_error", nil, "conn redis fail", http.StatusBadRequest)
	}
	defer c.Close()

	keys, err2 := redis.Strings(c.Do("keys", key+":"+"*"))
	if err2 != nil {
		return nil, model.NewAppError("RedisStoreGetOrgNode", "api.org.getorg.app_error", nil, "get allorgnode fail", http.StatusBadRequest)
	}
	for _, keystring := range keys {
		if s := strings.Trim(keystring, key+":"); !strings.Contains(s, ":") {
			value, err3 := redis.String(c.Do("get", keystring))
			if err3 != nil {
				return nil, model.NewAppError("RedisStoreGetOrgNode", "api.org.getorg.app_error", nil, "get orgnode fail", http.StatusBadRequest)
			}
			result[keystring] = value
		}
	}
	return result, nil
}
func RedisStoreDelOrgNode(key string) (map[string]string, *model.AppError) {
	var result = make(map[string]string)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, model.NewAppError("RedisStoreDelOrgNode", "api.org.delorg.app_error", nil, "conn redis fail", http.StatusBadRequest)
	}
	defer c.Close()

	keys, err2 := redis.Strings(c.Do("keys", key+"*"))
	if err2 != nil {
		return nil, model.NewAppError("RedisStoreDelOrgNode", "api.org.delorg.app_error", nil, "keys orgnode fail", http.StatusBadRequest)
	}
	for _, keystring := range keys {
		if value, err3 := redis.String(c.Do("get", keystring)); err3 != nil {
			if _, err4 := redis.String(c.Do("del", keystring)); err4 != nil {
				result[keystring] = value
			} else {
				return nil, model.NewAppError("RedisStoreDelOrgNode", "api.org.delorg.app_error", nil, "del orgnode fail", http.StatusBadRequest)
			}
		} else {
			return nil, model.NewAppError("RedisStoreDelOrgNode", "api.org.delorg.app_error", nil, "get orgnode fail", http.StatusBadRequest)

		}
	}
	return result, nil
}
func RedisStoreUpdateOrg(key string, name string) (map[string]string, *model.AppError) {
	var result = make(map[string]string)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, model.NewAppError("RedisStoreUpdateOrg", "api.org.updateorg.app_error", nil, "conn redis fail", http.StatusBadRequest)
	}
	defer c.Close()

	value, err2 := c.Do("get", key)
	if err2 != nil || value == "" {
		return nil, model.NewAppError("RedisStoreUpdateOrg", "api.org.updateorg.app_error", nil, "get org fail", http.StatusBadRequest)
	} else {
		_, err3 := redis.String(c.Do("set", key, name))
		if err3 != nil {
			return nil, model.NewAppError("RedisStoreUpdateOrg", "api.org.updateorg.app_error", nil, "update org fail", http.StatusBadRequest)
		}
		result[key] = name
		return result, nil
	}
}
