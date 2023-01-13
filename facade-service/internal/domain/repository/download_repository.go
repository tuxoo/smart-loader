package repository

const downloadTable = "download"

type DownloadRepository struct {
	db *PostgresDB
}

func NewDownloadRepository(db *PostgresDB) *DownloadRepository {
	return &DownloadRepository{
		db: db,
	}
}

func (r *DownloadRepository) FindBy() {

}
