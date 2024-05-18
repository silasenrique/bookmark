package command

import (
	"go-read-list/src/business/bookmark/domain/entity"
	"time"
)

type ReadResponse struct {
	Id       string
	Url      string
	Title    string
	Cover    string
	CreateAt string
	UpdateAt string
}

func ToReadResponde(read *entity.Bookmark) *ReadResponse {
	return &ReadResponse{
		Id:       read.GetId(),
		Url:      read.GetUrl(),
		Title:    read.GetTitle(),
		Cover:    read.GetCover(),
		CreateAt: time.Unix(read.GetCreateAt(), 0).Format("01-02-2006 15:04:05"),
		UpdateAt: time.Unix(read.GetUpdateAt(), 0).Format("01-02-2006 15:04:05"),
	}
}

type ReadersResponse []*ReadResponse

func ToReadersResponse(readers []*entity.Bookmark) *ReadersResponse {
	responses := ReadersResponse{}

	for _, r := range readers {
		responses = append(responses, ToReadResponde(r))
	}

	return &responses
}
