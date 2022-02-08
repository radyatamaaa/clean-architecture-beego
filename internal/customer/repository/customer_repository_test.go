package repository_test

import (
	customerRepo "clean-architecture-beego/internal/customer/repository"
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/helpers/test"
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository_Fetch(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//resultMock
	mockCustomer := domain.CustomerTest{}
	err = faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	limit := 5
	offset := 0

	//rowAndColumnMock
	row, fields := test.GetValueAndColumnStructToDriverValue(mockCustomer)

	querySelect := "*"

	querySelectCount := "count(*)"

	t.Run("success", func(t *testing.T) {
		rowsCount := sqlmock.NewRows([]string{"count"}).
			AddRow(limit)
		queryCount := "SELECT " + querySelectCount + " FROM `customers`"
		queryRegexCount := regexp.QuoteMeta(queryCount)
		mock.ExpectQuery(queryRegexCount).
			WillReturnRows(rowsCount)

		rows := sqlmock.NewRows(fields).
			AddRow(row...)
		query := "SELECT " + querySelect + " FROM `customers`" + " LIMIT 5"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).
			WillReturnRows(rows)
		a := customerRepo.NewCustomerRepository(db)

		fetch, err := a.Fetch(context.TODO(), limit, offset)
		assert.NoError(t, err)
		assert.NotNil(t, fetch)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		rowsCount := sqlmock.NewRows([]string{"count"}).
			AddRow(limit)
		queryCount := "SELECT " + querySelectCount + " FROM `customers`"
		queryRegexCount := regexp.QuoteMeta(queryCount)
		mock.ExpectQuery(queryRegexCount).
			WillReturnRows(rowsCount)

		rows := sqlmock.NewRows(fields).
			AddRow(row...)
		query := "SELECT " + querySelect + " FROM `customers`" + " LIMIT 10"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).
			WillReturnRows(rows)
		a := customerRepo.NewCustomerRepository(db)

		_, err := a.Fetch(context.TODO(), limit, offset)
		assert.Error(t, err)
	})

}

func TestCustomerRepository_FindByID(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//resultMock
	mockCustomer := domain.CustomerTest{}
	err = faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	id := mockCustomer.Id

	//rowAndColumnMock
	row, fields := test.GetValueAndColumnStructToDriverValue(mockCustomer)

	querySelect := "*"

	queryWhere := "id =?"

	queryOrder := "ORDER BY `customers`.`id`"

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(fields).
			AddRow(row...)
		query := "SELECT " + querySelect + " FROM `customers` WHERE " + queryWhere + " " + queryOrder + " LIMIT 1"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).
			WithArgs(id).
			WillReturnRows(rows)
		a := customerRepo.NewCustomerRepository(db)

		findByID, err := a.FindByID(context.TODO(), id)
		assert.NoError(t, err)
		assert.NotNil(t, findByID)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		rows := sqlmock.NewRows(fields).
			AddRow(row...)
		query := "SELECT " + querySelect + " FROM `customers` WHERE " + queryWhere + " LIMIT 1"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).
			WithArgs(id).
			WillReturnRows(rows)
		a := customerRepo.NewCustomerRepository(db)

		_, err := a.FindByID(context.TODO(), id)
		assert.Error(t, err)
	})

}

func TestCustomerRepository_Update(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//requestMock
	mockCustomerTest := domain.CustomerTest{}
	err = faker.FakeData(&mockCustomerTest)
	assert.NoError(t, err)

	mockCustomer := mockCustomerTest.ToCustomer()

	querySet := "`customer_name`=?,`phone`=?,`email`=?,`address`=?,`created_at`=?,`updated_at`=?"

	queryWhere := "`id` = ?"

	t.Run("success", func(t *testing.T) {
		query := "UPDATE `customers` SET " + querySet + " WHERE " + queryWhere
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectBegin()
		mock.ExpectExec(queryRegex).
			WithArgs(
				mockCustomer.CustomerName,
				mockCustomer.Phone,
				mockCustomer.Email,
				mockCustomer.Address,
				mockCustomer.CreatedAt,
				mockCustomer.UpdatedAt,
				mockCustomer.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		a := customerRepo.NewCustomerRepository(db)

		err := a.Update(context.TODO(), mockCustomer)
		assert.NoError(t, err)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		query := "UPDATE `customers` SET " + querySet + " " + queryWhere
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectBegin()
		mock.ExpectExec(queryRegex).
			WithArgs(
				mockCustomer.CustomerName,
				mockCustomer.Phone,
				mockCustomer.Email,
				mockCustomer.Address,
				mockCustomer.CreatedAt,
				mockCustomer.UpdatedAt,
				mockCustomer.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		a := customerRepo.NewCustomerRepository(db)

		err := a.Update(context.TODO(), mockCustomer)
		assert.Error(t, err)
	})

}

func TestCustomerRepository_Store(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//requestMock
	mockCustomerTest := domain.CustomerTest{}
	err = faker.FakeData(&mockCustomerTest)
	assert.NoError(t, err)

	mockCustomer := mockCustomerTest.ToCustomer()

	querySet := "`customer_name`,`phone`,`email`,`address`,`created_at`,`updated_at`,`id`"

	t.Run("success", func(t *testing.T) {
		query := "INSERT INTO `customers` (" + querySet + ") VALUES (?,?,?,?,?,?,?)"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectBegin()
		mock.ExpectExec(queryRegex).
			WithArgs(
				mockCustomer.CustomerName,
				mockCustomer.Phone,
				mockCustomer.Email,
				mockCustomer.Address,
				mockCustomer.CreatedAt,
				mockCustomer.UpdatedAt,
				mockCustomer.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		a := customerRepo.NewCustomerRepository(db)

		err := a.Store(context.TODO(), mockCustomer)
		assert.NoError(t, err)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		query := "INSERT `customers` SET " + querySet
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectBegin()
		mock.ExpectExec(queryRegex).
			WithArgs(
				mockCustomer.CustomerName,
				mockCustomer.Phone,
				mockCustomer.Email,
				mockCustomer.Address,
				mockCustomer.CreatedAt,
				mockCustomer.UpdatedAt,
				mockCustomer.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		a := customerRepo.NewCustomerRepository(db)

		err := a.Store(context.TODO(), mockCustomer)
		assert.Error(t, err)
	})

}

func TestCustomerRepository_Delete(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//resultMock
	mockCustomer := domain.CustomerTest{}
	err = faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	id := mockCustomer.Id

	queryWhere := "id =?"

	t.Run("success", func(t *testing.T) {
		query := "delete from customers where " + queryWhere
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectExec(queryRegex).
			WithArgs(mockCustomer.Id).
			WillReturnResult(sqlmock.NewResult(int64(mockCustomer.Id), 1))
		a := customerRepo.NewCustomerRepository(db)

		err := a.Delete(context.TODO(), id)
		assert.NoError(t, err)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		query := "DELETE FROM `customers` WHERE " + queryWhere
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectExec(queryRegex).
			WithArgs(mockCustomer.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		a := customerRepo.NewCustomerRepository(db)

		err := a.Delete(context.TODO(), id)
		assert.Error(t, err)
	})

}
