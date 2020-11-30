package dao

import (
	"database/sql"
)

type Error interface {
	error
	IsNoRecord() bool
}

type noRecord interface {
	IsNoRecord() bool
}

type noRecordErr struct{
	Msg string
	Err error
}

func (e *noRecordErr) IsNoRecord() bool { return true }
func (e *noRecordErr) Unwrap() error { return e.Err }
func (e *noRecordErr) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.Msg + " : " + e.Err.Error()
}

// 判定是否是noRecord类型err
func IsNoRecordErr(err error) bool {
	te, ok := err.(noRecord)
	return ok && te.IsNoRecord()
}



// 测试方法用来模拟返回sql.ErrNoRows error
func query(sqlStr string) (values []interface{}, err error) {
	return nil, sql.ErrNoRows
}
