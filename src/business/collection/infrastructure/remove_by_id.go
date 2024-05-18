package infrastructure

func (f *collectionPersistence) RemevoFolderById(id int64) error {
	_, err := f.Exec("delete from collection where id = ?", id)

	return err
}
