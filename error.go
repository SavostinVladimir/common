package common

import "github.com/astaxie/beego/logs"

func CheckError(err error){
	if err != nil {
		logs.Error(err)
	}
}
