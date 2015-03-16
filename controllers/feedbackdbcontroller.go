package controllers

import (
	"93ztc/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func AddNewFeedback(title string, content string, uname string) {
	if title == "" || content == "" || uname == "" {
		return
	}
	o := orm.NewOrm()
	feedBack := new(models.Feedback)
	feedBack.Title = title
	feedBack.Content = content
	feedBack.Owner = uname
	feedBack.CreateTime = time.Now().Unix()
	o.Insert(feedBack)
}

func GetFeedbackList(uname string) ([]*models.Feedback, bool) {
	o := orm.NewOrm()
	var feedbacks []*models.Feedback
	num1, _ := o.Raw("select * from feedback where owner=?", uname).QueryRows(&feedbacks)
	beego.Debug("found feedbacks:", num1)
	if num1 > 0 {
		beego.Debug("Got feedbacks:", feedbacks)
		return feedbacks, true
	} else {
		return feedbacks, false
	}
	return feedbacks, false
}

func GetReadableFeedbackList(uname string) ([]*models.ReadableFeedback, bool) {
	o := orm.NewOrm()
	var feedbacks []*models.Feedback
	var readFeedbacks []*models.ReadableFeedback
	num1, _ := o.Raw("select * from feedback where owner=?", uname).QueryRows(&feedbacks)
	beego.Debug("found feedbacks :", num1)
	if num1 > 0 {
		readFeedbacks = make([]*models.ReadableFeedback, 0, num1)
		for i, _ := range feedbacks {
			rfb := new(models.ReadableFeedback)
			rfb.Id = feedbacks[i].Id
			rfb.Title = feedbacks[i].Title
			rfb.Content = feedbacks[i].Content
			rfb.DevFeedback = feedbacks[i].DevFeedback
			//rfb.CreateTime, _ = fmt.Print(time.Unix(feedbacks[i].CreateTime, 0))
			rfb.CreateTime = time.Unix(feedbacks[i].CreateTime, 0).String()
			rfb.Owner = feedbacks[i].Owner
			rfb.FeedbackTime = time.Unix(feedbacks[i].FeedbackTime, 0).String()
			readFeedbacks = append(readFeedbacks, rfb)
		}
		beego.Debug("Got feedbacks:", readFeedbacks)
		return readFeedbacks, true
	} else {
		return readFeedbacks, false
	}
	return readFeedbacks, false
}

func AddNewSubcribeEmail(email string) {
	if email == "" {
		return
	}
	o := orm.NewOrm()
	subscriber := new(models.MailList)
	subscriber.EMail = email
	subscriber.SubscribeTime = time.Now().Unix()
	o.Insert(subscriber)
}
