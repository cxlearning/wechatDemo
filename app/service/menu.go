package service

import (
	"errors"
	"fmt"
	"github.com/wechatDemo/app/model"
	"github.com/wechatDemo/library/http"
)

// CreateMenu 创建菜单
func CreateMenu(buttons []model.Button) (err error) {
	if len(buttons) > 3 {
		return errors.New("too many first level menu, must less than 3")
	}
	for _, sub := range buttons {
		if len(sub.SubButton) > 5 {
			return errors.New("too many second level menu, must less than 5")
		}
	}

	menu := struct {
		Button []model.Button `json:"button"`
	}{buttons}

	url := fmt.Sprintf(model.MenuCreateURL, AccessToken())
	fmt.Println(url)
	return http.Post(url, menu, nil)
}

// GetMenu 查询菜单
func GetMenu() (all *model.AllMenu, err error) {
	url := fmt.Sprintf(model.MenuGetURL, AccessToken())
	all = &model.AllMenu{}
	err = http.Get(url, all)
	return all, err
}

// DeleteMenu 删除菜单
func DeleteMenu() (err error) {
	url := fmt.Sprintf(model.MenuDeleteURL, AccessToken())
	return http.Get(url, nil)
}

// GetMenuInfo 获取自定义菜单配置
func GetMenuInfo() (mi *model.MenuInfo, err error) {
	url := fmt.Sprintf(model.MenuInfoURL, AccessToken())
	mi = &model.MenuInfo{}
	err = http.Get(url, mi)
	return mi, err
}
