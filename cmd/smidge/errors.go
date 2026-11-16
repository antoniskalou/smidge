package main

import "fmt"

type cliErrorCode string

const (
	errorCodeDatabaseOpen      cliErrorCode = "database.open"
	errorCodeMigrationDir      cliErrorCode = "migration.directory_create"
	errorCodeMigrationWrite    cliErrorCode = "migration.file_write"
	errorCodeMigrationStatus   cliErrorCode = "migration.status"
	errorCodeMigrationDeploy   cliErrorCode = "migration.deploy"
	errorCodeMigrationRollback cliErrorCode = "migration.rollback"
	errorCodeInternalUnknown   cliErrorCode = "internal.unknown"
)

type cliError struct {
	code    cliErrorCode
	message string
	hint    string
	cause   error
}

func (e *cliError) Error() string {
	if e.hint != "" {
		return fmt.Sprintf("%s: %v", e.message, e.cause)
	}

	return e.message
}

func (e *cliError) Unwrap() error {
	return e.cause
}
