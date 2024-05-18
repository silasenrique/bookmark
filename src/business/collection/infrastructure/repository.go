package infrastructure

import (
	"database/sql"
	"go-read-list/src/business/collection/domain/repository"
)

type collectionPersistence struct {
	*sql.DB
}

func NewCollectionRepository(db *sql.DB) repository.FolderWriterReader {
	return &collectionPersistence{db}
}

func (f *collectionPersistence) GetNextId() (int64, error) {
	id := 0

	err := f.QueryRow("select id + 1 from collection_seq").Scan(&id)
	if err != nil {
		return 0, err
	}

	_, err = f.Exec("update collection_seq set id = ?", id)
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}
