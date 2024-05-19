package mapper

import (
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/domain/entity"
	"time"
)

type CollectionDto struct {
	ID          int64
	Name        string
	Icon        string
	IntParentId int64
	CreateAt    int64
	UpdateAt    int64
}

func DtoToCollection(colle *CollectionDto) *entity.Collection {
	if colle == nil {
		return nil
	}

	return entity.NewCollection(colle.ID, colle.Name).
		AddInternalParentIO(colle.IntParentId).
		AddCreateAt(colle.CreateAt).
		AddUpdateAt(colle.UpdateAt).
		AddIcon(colle.Icon)
}

func CollectionToResponse(colle *entity.Collection) *command.CollectionResponse {
	if colle == nil {
		return nil
	}

	return &command.CollectionResponse{
		Id:       colle.GetID(),
		Name:     colle.GetName(),
		Icon:     colle.GetIcon(),
		Parent:   CollectionToResponse(colle.GetParent()),
		CreateAt: time.Unix(colle.GetUnixCreateAt(), 0).Format("01-02-2006 15:04:05"),
		UpdateAt: time.Unix(colle.GetUnixUpdateAt(), 0).Format("01-02-2006 15:04:05"),
	}
}

type CollectionsResponse []*command.CollectionResponse

func ToCollectionsResponse(colle []*entity.Collection) *CollectionsResponse {
	responses := CollectionsResponse{}

	for _, r := range colle {
		responses = append(responses, CollectionToResponse(r))
	}

	return &responses
}
