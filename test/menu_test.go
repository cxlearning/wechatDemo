package test

import (
	"github.com/wechatDemo/app/model"
	"github.com/wechatDemo/app/service"
	"github.com/wechatDemo/config"
	"testing"
	"time"
)

func TestCreateMenu(t *testing.T) {
	service.RefreshAccessToken(config.AppId, config.AppSecret)
	time.Sleep(3*time.Second)
	buttons := []model.Button{
		model.Button{
			Name: "扫码",
			SubButton: []model.Button{
				model.Button{
					Name: "扫码带提示",
					Type: model.MenuTypeScancodeWaitmsg,
					Key:  "rselfmenu_0_0",
				},
				model.Button{
					Name: "扫码推事件",
					Type: model.MenuTypeScancodePush,
					Key:  "rselfmenu_0_1",
				},
			},
		},
		model.Button{
			Name: "发图",
			SubButton: []model.Button{
				model.Button{
					Name: "系统拍照发图",
					Type: model.MenuTypePicSysphoto,
					Key:  "rselfmenu_1_0",
				},
				model.Button{
					Name: "拍照或者相册发图",
					Type: model.MenuTypePicPhotoOrAlbum,
					Key:  "rselfmenu_1_1",
				},
				model.Button{
					Name: "微信相册发图",
					Type: model.MenuTypePicWeixin,
					Key:  "rselfmenu_1_2",
				},
			},
		},
		model.Button{
			Name: "测试",
			SubButton: []model.Button{
				model.Button{
					Name: "腾讯",
					Type: model.MenuTypeView,
					URL:  "http://qq.com",
				},
				model.Button{
					Name: "发送位置",
					Type: model.MenuTypeLocationSelect,
					Key:  "rselfmenu_2_0",
				},
			},
		},
	}

	err := service.CreateMenu(buttons)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestGetMenu(t *testing.T) {
	buttons, err := service.GetMenu()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", buttons)
}

func TestDeleteMenu(t *testing.T) {
	err := service.DeleteMenu()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestGetMenuInfo(t *testing.T) {
	mi, err := service.GetMenuInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%v", mi)
}
