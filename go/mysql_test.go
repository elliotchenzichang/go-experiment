package _go

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"
)

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
}

// a successful case
func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// a failing test case
func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	// now we execute our method
	if err = recordStats(db, 2, 3); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

type Employee struct {
	ID        int64
	Name      string
	Age       int
	CreatedAt time.Time
}

func GetAllEmployees(db *sql.DB) ([]Employee, error) {
	rows, err := db.Query("SELECT id, name, age, created_at FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emps []Employee
	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Age, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		emps = append(emps, e)
	}
	return emps, nil
}

func TestGetAllEmployees(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("opening a stub database connection error=%s", err)
	}
	defer db.Close()

	// mock return rows
	rows := sqlmock.NewRows([]string{"id", "name", "age", "created_at"}).
		AddRow(1, "john", 33, time.Now()).
		AddRow(2, "mary", 28, time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnRows(rows)

	emps, err := GetAllEmployees(db)
	count := len(emps)
	expected := 2
	if count != expected {
		t.Errorf("Expected %d, but %d", expected, count)
	}
}

func GetEmployeeByID(db *sql.DB, id int64) (*Employee, error) {
	row := db.QueryRow("SELECT * FROM employee WHERE id = $1 LIMIT 1", id)
	var emp Employee
	err := row.Scan(
		&emp.ID,
		&emp.Name,
		&emp.Age,
		&emp.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func TestGetEmployeeByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("opening a stub database connection error=%s", err)
	}
	defer db.Close()

	createdAt, _ := time.Parse("2006-01-02", "2021-01-14")
	rows := sqlmock.NewRows([]string{"id", "name", "age", "created_at"}).
		AddRow(1, "john", 33, createdAt)

	// mock return rows
	mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id = \\$1 LIMIT 1$").
		WillReturnRows(rows)

	emp, err := GetEmployeeByID(db, 1)

	expected := Employee{
		ID:        1,
		Name:      "john",
		Age:       33,
		CreatedAt: createdAt,
	}
	if *emp != expected {
		t.Errorf("Expected %v, but %v", expected, emp)
	}
}
