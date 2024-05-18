package entity

import "time"

type Bookmark struct {
	id       ID
	url      string
	title    string
	path     ID
	cover    string
	createAt int64
	updateAt int64
}

func NewRead(id ID, uri, title string) *Bookmark {
	time := time.Now().Unix()

	return &Bookmark{
		id:       id,
		url:      uri,
		title:    title,
		createAt: time,
		updateAt: time,
	}
}

func (u *Bookmark) GetTitle() string {
	return u.title
}

func (u *Bookmark) GetId() string {
	return u.id.String()
}

func (u *Bookmark) AddCover(cover string) {
	if cover == "" {
		return
	}

	u.cover = cover
}

func (u *Bookmark) GetCover() string {
	return u.cover
}

func (u *Bookmark) ChangeTitle(newTitle string) {
	u.title = newTitle
}

func (u *Bookmark) GetUrl() string {
	return u.url
}

func (u *Bookmark) ChangeUrl(newUrl string) {
	u.url = newUrl
}

func (u *Bookmark) AddPath(id ID) {
	u.path = id
}

func (u *Bookmark) RemovePath() {
	u.path = ""
}

func (u *Bookmark) GetPathId() ID {
	if u.path == "" {
		return ""
	}

	return u.path
}

func (u *Bookmark) GetCreateAt() int64 {
	return u.createAt
}

func (u *Bookmark) GetUpdateAt() int64 {
	return u.updateAt
}
