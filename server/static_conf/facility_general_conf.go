package static_conf

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path"
	"slgserver/config"
	"slgserver/log"
)

//城内设施校场、统帅厅
var FGEN facilityGeneralConf

type level struct {
	Level	int8			`json:"level"`
	Cnt		int8 			`json:"cnt"`
	Need	levelNeedRes	`json:"need"`
}

type general struct {
	Name	string		`json:"name"`
	Des		string		`json:"des"`
	Type	int8		`json:"type"`
	Levels	[]level		`json:"levels"`
}

type facilityGeneralConf struct {
	Title	string		`json:"title"`
	JC		general		`json:"jc"`		//校场
	TST		general		`json:"tst"`	//统帅厅
}

func (this *facilityGeneralConf) Load()  {
	jsonDir := config.File.MustValue("logic", "json_data", "../data/conf/")
	fileName := path.Join(jsonDir, "facility_general.json")
	jdata, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.DefaultLog.Error("facilityGeneralConf load file error", zap.Error(err), zap.String("file", fileName))
		os.Exit(0)
	}

	json.Unmarshal(jdata, this)
	fmt.Println(this)
}

func (this *facilityGeneralConf) MaxLevel(fType int8) int8 {
	if this.JC.Type == fType{
		return int8(len(this.JC.Levels))
	}else if this.TST.Type == fType{
		return int8(len(this.TST.Levels))
	}else{
		return 0
	}
}