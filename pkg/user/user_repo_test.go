package user

import (
	"fmt"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()
	repo := &UsersSQL{DB: db}
	var reqStr = "user1"
	rows := sqlmock.NewRows([]string{"id", "username", "passwordhash"})

	t.Run("positive test", func(t *testing.T) {

		expect := []*User{
			{"5fgdgds", reqStr, "fsd"},
		}

		for _, item := range expect {
			rows = rows.AddRow(item.Id, item.Username, item.PasswordHash)
		}

		mock.ExpectQuery("SELECT id, username, passwordhash FROM users WHERE").
			WithArgs(reqStr).
			WillReturnRows(rows)

		user, err := repo.GetUser(reqStr)
		if err != nil {
			t.Errorf("unexpected error:%s", err)
			return
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled experations: %s", err)
			return
		}
		if !reflect.DeepEqual(user, expect[0]) {
			t.Errorf("results not match  expected: %v  got: %v", expect[0], user)
			return
		}
	})

	t.Run("negative test", func(t *testing.T) {

		mock.ExpectQuery("SELECT id, username, passwordhash FROM users WHERE").
			WithArgs(reqStr).
			WillReturnError(fmt.Errorf("error db"))

		_, err = repo.GetUser(reqStr)
		if err == nil {
			t.Errorf("expected error got nil")
			return
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled experations: %s", err)
			return
		}

	})
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()
	repo := &UsersSQL{DB: db}

	t.Run("positive test", func(t *testing.T) {

		testUser := &User{
			Id:           "1",
			Username:     "vdffd",
			PasswordHash: "fsddfa",
		}

		mock.ExpectExec("INSERT INTO users ").
			WithArgs(testUser.Id, testUser.Username, testUser.PasswordHash).
			WillReturnResult(sqlmock.NewResult(1, 1))

		user, err := repo.CreateUser(testUser)
		if err != nil {
			t.Errorf("unexpected error:%s", err)
			return
		}

		if user.Id != "1" {
			t.Errorf("bad id")
			return
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled experations: %s", err)
			return
		}

	})

	t.Run("negative test", func(t *testing.T) {
		testUser := &User{
			Id:           "1",
			Username:     "vdffd",
			PasswordHash: "fsddfa",
		}

		mock.ExpectExec("INSERT INTO users").
			WithArgs(testUser.Id, testUser.Username, testUser.PasswordHash).
			WillReturnError(fmt.Errorf("error db"))

		_, err = repo.CreateUser(testUser)
		if err == nil {
			t.Errorf("expected error got nil")
			return
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled experations: %s", err)
			return
		}

	})

}
