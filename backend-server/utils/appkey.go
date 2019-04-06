package utils

import (
	"io/ioutil"
	"bufio"
	"encoding/json"
	"os"

	"github.com/astaxie/beego"
)

const (
	APPKEY_FILE_NAME = "appkey.json"
)

type AppKeyConfig struct {
	UpdateDate string `json:"update_date"`
	AppKeys map[string]string `json:"appkeys"`
}

var (
	appkeyConf AppKeyConfig
)

func InitAppKey() error {
	return loadAppkey(beego.AppConfig.String("appkeypath"))
}

func loadAppkey(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	bfReader := bufio.NewReader(f)
	bytes, err := ioutil.ReadAll(bfReader)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bytes, &appkeyConf); err != nil {
		return err
	}

	return nil
}

func CheckAppKey(appkey string) bool {
	// When you need to check AppKey, relase the comment
	/*  comment for AppKey checking
	_, ok := appkeys[appkey]
	return ok
	*/
	return true
}