package service

import (
	"jmolboy/homework/week02/app/dao"
	"jmolboy/homework/week02/app/model"
)

type NewsService struct{}

func (newsService *NewsService) Find(id int) (news *model.News, err error) {
	newsDao := dao.NewsDao{}
	return newsDao.Find(id)
}
