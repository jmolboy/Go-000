package service

import "jmolboy/homework/week02/app/dao"


func IsNoRecordErr(err error) bool {
	return dao.IsNoRecordErr(err)
}