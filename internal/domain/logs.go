package domain

import "time"

type LogItem struct {
	UserId    int       `bson:"user_id"`
	ArticleId int       `bson:"article_id"`
	Action    string    `bson:"action"`
	Time      time.Time `bson:"time"`
}
