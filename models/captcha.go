package models

type Captcha struct {
	Id          int64
	ItemId      string
	CrawlerId   string
	CaptchaFile string
	Result      string `orm:"null"`
}
