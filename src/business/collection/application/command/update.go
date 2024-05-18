package command

type CollectionUpdateCommand struct {
	Id           int64
	Name         string
	ParentId     int64
	Icon         string
	RemoveParent bool
}
