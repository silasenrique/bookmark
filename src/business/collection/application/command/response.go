package command

import (
	"go-read-list/src/business/collection/domain/entity"
	"time"
)

type CollectionResponse struct {
	Id       int64
	Name     string
	Parent   *CollectionResponse
	Icon     string
	CreateAt string
	UpdateAt string
}

func ToCollectionResponse(collection *entity.Collection) *CollectionResponse {
	return &CollectionResponse{
		Id:       collection.GetID(),
		Name:     collection.GetName(),
		Icon:     collection.GetIcon(),
		CreateAt: time.Unix(collection.GetCreateAt(), 0).Format("01-02-2006 15:04:05"),
		UpdateAt: time.Unix(collection.GetUpdateAt(), 0).Format("01-02-2006 15:04:05"),
	}
}

func (f *CollectionResponse) AddParent(folder *entity.Collection) *CollectionResponse {
	f.Parent = ToCollectionResponse(folder)

	return f
}

type FoldersResponse []*CollectionResponse

func ToFoldersResponse(folders []*entity.Collection) *FoldersResponse {
	responses := FoldersResponse{}

	for _, r := range folders {
		responses = append(responses, ToCollectionResponse(r))
	}

	return &responses
}
