package entity

import "time"

// Collection of saved links and other collections
type Collection struct {
	id       int64
	name     string
	icon     string
	parent   *Collection
	parentId int64
	createAt int64
	updateAt int64
}

// Returns an instance of a new collection
func NewCollection(id int64, name string) *Collection {
	time := time.Now().Unix()

	return &Collection{
		id:       id,
		name:     name,
		createAt: time,
		updateAt: time,
	}
}

// Returns the unique identifier
func (c *Collection) GetID() int64 {
	return c.id
}

func (c *Collection) GetParentID() int64 {
	return c.parentId
}

func (c *Collection) GetParent() *Collection {
	return c.parent
}

// Get folder name
func (c *Collection) GetName() string {
	return c.name
}

// Remove collection from another collection
func (c *Collection) RemoveParent() {
	c.parent = nil
}

// Get collection creation date
func (c *Collection) GetUnixCreateAt() int64 {
	return c.createAt
}

// Get the date of the last change in the collection
func (c *Collection) GetUnixUpdateAt() int64 {
	return c.updateAt
}

// Get collection creation date
func (c *Collection) GetStringCreateAt() string {
	return time.Unix(c.createAt, 0).Format("01-02-2006 15:04:05")
}

// Get the date of the last change in the collection
func (c *Collection) GetStringUpdateAt() string {
	return time.Unix(c.updateAt, 0).Format("01-02-2006 15:04:05")
}

func (c *Collection) AddId(id int64) *Collection {
	c.id = id

	return c
}

func (c *Collection) AddCreateAt(at int64) *Collection {
	c.createAt = at

	return c
}

func (c *Collection) AddUpdateAt(at int64) *Collection {
	c.updateAt = at

	return c
}

func (c *Collection) AddParentId(id int64) *Collection {
	if id == 0 {
		return c
	}

	c.parentId = id

	return c
}

// Move collection into another collection
func (c *Collection) AddParent(parent *Collection) *Collection {
	if parent == nil {
		return c
	}

	c.parentId = parent.id
	c.parent = parent

	return c
}

// Change folder name
func (c *Collection) AddName(newName string) *Collection {
	if newName == "" {
		return c
	}

	c.name = newName

	return c
}

func (c *Collection) AddIcon(icon string) *Collection {
	if icon == "" {
		return c
	}

	c.icon = icon

	return c
}

func (c *Collection) GetIcon() string {
	return c.icon
}
