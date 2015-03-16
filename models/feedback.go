package models

type Feedback struct {
	Id           int64
	Title        string
	Content      string
	DevFeedback  string `orm:"null"`
	CreateTime   int64
	Owner        string
	FeedbackTime int64 `orm:"null"`
}

type ReadableFeedback struct { //only for ui, no need to register
	Id           int64
	Title        string
	Content      string
	DevFeedback  string
	CreateTime   string
	Owner        string
	FeedbackTime string
}

type MailList struct {
	Id            int64
	EMail         string
	SubscribeTime int64 `orm:default(0)`
}
