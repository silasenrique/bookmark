package entity

type BookmarkDTO struct {
	Id       string
	Url      string
	Title    string
	Path     string
	Cover    string
	CreateAt int64
	UpdateAt int64
}

func DtoToReader(read *BookmarkDTO) *Bookmark {
	return &Bookmark{
		id:       ID(read.Id),
		url:      read.Url,
		title:    read.Title,
		path:     ID(read.Path),
		cover:    read.Cover,
		createAt: read.CreateAt,
		updateAt: read.UpdateAt,
	}
}
