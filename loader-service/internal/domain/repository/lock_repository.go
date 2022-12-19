package repository

type LockRepository struct {
	db *PostgresDB
}

func NewLockRepository(db *PostgresDB) *LockRepository {
	return &LockRepository{
		db: db,
	}
}

func (r *LockRepository) ChangeState(types, value string, state bool) error {

	return nil
}
