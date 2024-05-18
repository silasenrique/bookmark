package infrastructure

func (f *collectionPersistence) CountDependencies(id int64) (int64, error) {
	count := 0

	sqt, err := f.Prepare("select count(*) from collection where collection_id = ?")
	if err != nil {
		return 0, err
	}

	defer sqt.Close()

	err = sqt.QueryRow(id).Scan(&count)
	if err != nil {
		return 0, nil
	}

	return int64(count), nil
}
