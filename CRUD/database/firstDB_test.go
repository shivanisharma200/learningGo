package database

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	return db, mock
}

// Test for GetById
func Test_GetById(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	testCases := []struct {
		id          int
		emp         *Employee
		mockQuery   interface{}
		expectError error
	}{
		// Success Case
		{
			id:  1,
			emp: &Employee{1, "John", "jhn@yahoo.com", "intern"},
			mockQuery: mock.ExpectQuery("SELECT * from employee WHERE id = ?").
				WithArgs(1).
				WillReturnRows(mock.NewRows([]string{"id", "name", "email", "role"}).
					AddRow(1, "John", "jhn@yahoo.com", "intern")),
			expectError: nil,
		},
		{
			id:  2,
			emp: &Employee{2, "Jane", "jane@yahoo.com", "intern"},
			mockQuery: mock.ExpectQuery("SELECT * from employee WHERE id = ?").
				WithArgs(2).
				WillReturnRows(mock.NewRows([]string{"id", "name", "email", "role"}).
					AddRow(2, "Jane", "jane@yahoo.com", "intern")),
			expectError: nil,
		},
		// Failure
		{
			id:  2,
			emp: nil,
			mockQuery: mock.ExpectQuery("SELECT * from employee WHERE id = ?").
				WithArgs(2).
				WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
		},
	}
	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			emp, err := GetById(testCase.id, db)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error:%v, got:%v", testCase.expectError, err)
			}
			if !reflect.DeepEqual(emp, testCase.emp) {
				t.Errorf("expected userL%v, got:%v", Employee{1, "John", "jhn@yahoo.com", "intern"}, emp)
			}
		})
	}
}

// Test for InsertRow
func Test_InsertRow(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	testCases := []struct {
		id          int
		emp         Employee
		mockQuery   interface{}
		expectError error
	}{
		// Success
		{
			id:  1,
			emp: Employee{1, "John", "jhn@yahoo.com", "intern"},
			mockQuery: mock.ExpectExec("INSERT INTO employee(name, email, role) VALUES (?, ?, ?)").
				WithArgs("John", "jhn@yahoo.com", "intern").
				WillReturnResult(sqlmock.NewResult(1, 1)),
			expectError: nil,
		},
		{
			id:  2,
			emp: Employee{1, "John", "jhn@yahoo.com", "intern"},
			mockQuery: mock.ExpectExec("INSERT INTO employee(name, email, role) VALUES (?, ?, ?)").
				WithArgs("John", "jhn@yahoo.com", "intern").
				WillReturnError(sql.ErrConnDone),
			expectError: sql.ErrConnDone,
		},
	}
	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := InsertRow(testCase.emp, db)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error:%v, got:%v", testCase.expectError, err)
			}
		})
	}
}

// Test for DeleteById
func Test_DeleteById(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	testCases := []struct {
		id          int
		emp         Employee
		mockQuery   interface{}
		expectError error
	}{
		// Success
		{
			id:  1,
			emp: Employee{1, "John", "jhn@yahoo.com", "intern"},
			mockQuery: mock.ExpectPrepare("DELETE FROM employee WHERE id = ?").
				ExpectExec().WithArgs(1).
				WillReturnResult(sqlmock.NewResult(0, 1)),
			expectError: nil,
		},
		// Prepare Error
		{
			id:  1,
			emp: Employee{1, "John", "jhn@yahoo.com", "intern"},

			mockQuery:   mock.ExpectPrepare("DELETE FROM employee WHERE id = ?").WillReturnError(errors.New("prepare query not working")),
			expectError: errors.New("prepare query not working"),
		},
		// Failure
		{
			id:  2,
			emp: Employee{1, "John", "jhn@yahoo.com", "intern"},
			mockQuery: mock.ExpectPrepare("DELETE FROM employee WHERE id = ?").
				ExpectExec().WithArgs(2).
				WillReturnError(sql.ErrTxDone),
			expectError: sql.ErrTxDone,
		},
	}
	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := DeleteById(testCase.id, db)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error:%v, got:%v", testCase.expectError, err)
			}
		})
	}
}

// Test for Update
func Test_Update(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	testCases := []struct {
		id          int
		emp         Employee
		mockQuery   interface{}
		expectError error
	}{
		// Success
		{
			id:  1,
			emp: Employee{1, "John", "jhn@yahoo.com", "intern"},
			mockQuery: mock.ExpectPrepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?").
				ExpectExec().WithArgs("John", "jhn@yahoo.com", "intern", 1).
				WillReturnResult(sqlmock.NewResult(0, 1)),
			expectError: nil,
		},
		// Prepare Error
		{
			id:  1,
			emp: Employee{1, "John", "jhn@yahoo.com", "intern"},

			mockQuery:   mock.ExpectPrepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?").WillReturnError(errors.New("prepare query not working")),
			expectError: errors.New("prepare query not working"),
		},
		// Failure
		{
			id:  2,
			emp: Employee{1, "John", "jhn@yahoo.com", "intern"},
			mockQuery: mock.ExpectPrepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?").
				ExpectExec().WithArgs("John", "jhn@yahoo.com", "intern", 1).
				WillReturnError(sql.ErrConnDone),
			expectError: sql.ErrConnDone,
		},
	}
	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := Update(testCase.emp, db)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error:%v, got:%v", testCase.expectError, err)
			}
		})
	}
}
