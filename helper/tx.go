package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	// check if error occurs
	err := recover()

	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
