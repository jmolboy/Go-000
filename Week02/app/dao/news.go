package dao

import (
	"database/sql"
	"errors"
	"jmolboy/homework/week02/app/model"
	"log"
)

type NewsDao struct {
	Id int
}

func (newsDao *NewsDao) Find(id int) (*model.News, error) {
	_, err := query("select * from news where id=1 limit 1")
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, &noRecordErr{Err: err, Msg: "未发现记录"}
	}

	if err != nil {
		//记录日志
		log.Fatal(err)

		// 返回查询失败queryFail类型error
		return nil, &queryFailErr{Err: err, Msg: "访问异常"}
	}

	news := &model.News{}
	// 数据绑定
	// values绑定到news
	// ...

	return news, nil
}
