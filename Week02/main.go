package main

import (
	"fmt"
	"jmolboy/homework/week02/app/service"
)

func main() {
	/*
		我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
		为什么，应该怎么做请写出代码？
	*/

	/*  dao层定义noRecordErr，用来统一作为dao层遇到的sql.ErrNoRows时的返回值，并在dao层提供IsNoRecordErr方法
	    用来判定是否是noRecordErr类型，这样的优点：
		1.service等调用dao的层不直接依赖于sql.ErrNoRows实现解耦
		2.dao层切换时调用dao的其它地方可以快速切换

	    noRecordErr定义参考net包
	*/
	newsSvs := service.NewsService{}
	news, err := newsSvs.Find(0)
	if err != nil && service.IsNoRecordErr(err) {
		fmt.Printf("no record found", err)
	} else if (err != nil) {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("news id:%d", news.Id)
	}
}
