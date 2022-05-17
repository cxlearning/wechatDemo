package service

import (
	"errors"
	"fmt"
	"github.com/wechatDemo/app/model"
	"github.com/wechatDemo/library/http"
)



// CreateConditionalMenu 创建个性化菜单
func CreateConditionalMenu(cm *model.ConditionalMenu) (menuId string, err error) {
	if len(cm.Button) > 3 {
		return "", errors.New("too many first level menu, must less than 3")
	}
	for _, sub := range cm.Button {
		if len(sub.SubButton) > 5 {
			return "", errors.New("too many second level menu, must less than 5")
		}
	}

	url := fmt.Sprintf(model.MenuCreateConditionalURL, AccessToken())

	wapper := &struct {
		http.WXError
		MenuId string `json:"menuid"`
	}{}
	err = http.Post(url, cm, wapper)
	return wapper.MenuId, err
}

// DeleteConditionalMenu 删除个性化菜单，menuId 为菜单 id，可以通过自定义菜单查询接口获取
func DeleteConditionalMenu(menuId int) (err error) {
	url := fmt.Sprintf(model.MenuDeleteConditionalURL, AccessToken())
	js := fmt.Sprintf(`{"menuid":"%d"}`, menuId)
	return http.Post(url, []byte(js), nil)
}

// TryMatchConditionalMenu 测试个性化菜单匹配结果，userId 可以是粉丝的 OpenID，也可以是粉丝的微信号
func TryMatchConditionalMenu(userId string) (buttons []model.Button, err error) {
	url := fmt.Sprintf(model.MenuTryMatchConditionalMenuURL, AccessToken())
	js := fmt.Sprintf(`{"user_id":"%s"}`, userId)

	wapper := &struct {
		http.WXError
		Button []model.Button `json:"button"`
	}{}
	err = http.Post(url, []byte(js), wapper)
	return wapper.Button, err
}
