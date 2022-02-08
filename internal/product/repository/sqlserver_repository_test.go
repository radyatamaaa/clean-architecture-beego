package repository_test

import (
	"clean-architecture-beego/internal/domain"
	productRepo "clean-architecture-beego/internal/product/repository"
	"clean-architecture-beego/pkg/helpers/test"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestProductRepository_Fetch(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	//resultMock
	mockProduct := domain.ProductTest{}
	err = faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	limit  := 5
	offset := 0

	//rowAndColumnMock
	row, fields := test.GetValueAndColumnStructToDriverValue(mockProduct)

	querySelect := "*"

	querySelectCount := "count(*)"

	t.Run("success", func(t *testing.T) {
		rowsCount := sqlmock.NewRows([]string{"count"}).
			AddRow(limit)
		queryCount := "SELECT " + querySelectCount + " FROM `products`"
		queryRegexCount := regexp.QuoteMeta(queryCount)
		mock.ExpectQuery(queryRegexCount).
			WillReturnRows(rowsCount)

		rows := sqlmock.NewRows(fields).
			AddRow(row...)
		query := "SELECT " + querySelect + " FROM `products`" + " LIMIT 5"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).
			WillReturnRows(rows)
		a := productRepo.NewProductRepository(db)

		fetch, err := a.Fetch(context.TODO(),limit,offset)
		assert.NoError(t, err)
		assert.NotNil(t, fetch)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		rowsCount := sqlmock.NewRows([]string{"count"}).
			AddRow(limit)
		queryCount := "SELECT " + querySelectCount + " FROM `products`"
		queryRegexCount := regexp.QuoteMeta(queryCount)
		mock.ExpectQuery(queryRegexCount).
			WillReturnRows(rowsCount)

		rows := sqlmock.NewRows(fields).
			AddRow(row...)
		query := "SELECT " + querySelect + " FROM `products`" + " LIMIT 10"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).
			WillReturnRows(rows)
		a := productRepo.NewProductRepository(db)


		_, err := a.Fetch(context.TODO(),limit,offset)
		assert.Error(t, err)
	})


}

func TestProductRepository_FindByID(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	//resultMock
	mockProduct := domain.ProductTest{}
	err = faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	id  := mockProduct.Id

	//rowAndColumnMock
	row, fields := test.GetValueAndColumnStructToDriverValue(mockProduct)

	querySelect := "*"

	queryWhere := "id =?"

	queryOrder := "ORDER BY `products`.`id`"

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(fields).
			AddRow(row...)
		query := "SELECT " + querySelect + " FROM `products` WHERE " + queryWhere + " " + queryOrder + " LIMIT 1"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).
			WithArgs(id).
			WillReturnRows(rows)
		a := productRepo.NewProductRepository(db)

		findByID, err := a.FindByID(context.TODO(),id)
		assert.NoError(t, err)
		assert.NotNil(t, findByID)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		rows := sqlmock.NewRows(fields).
			AddRow(row...)
		query := "SELECT " + querySelect + " FROM `products` WHERE " + queryWhere + " LIMIT 1"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).
			WithArgs(id).
			WillReturnRows(rows)
		a := productRepo.NewProductRepository(db)

		_, err := a.FindByID(context.TODO(),id)
		assert.Error(t, err)
	})


}

func TestProductRepository_Update(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	//requestMock
	mockProductTest := domain.ProductTest{}
	err = faker.FakeData(&mockProductTest)
	assert.NoError(t, err)

	mockProduct := mockProductTest.ToProduct()

	querySet := "`product_name`=?,`product_price`=?,`stock`=?,`created_at`=?,`updated_at`=?"

	queryWhere := "`id` = ?"


	t.Run("success", func(t *testing.T) {
		query := "UPDATE `products` SET " + querySet + " WHERE " + queryWhere
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectBegin()
		mock.ExpectExec(queryRegex).
			WithArgs(
			mockProduct.ProductName,
			mockProduct.Price,
			mockProduct.Stock,
			mockProduct.CreatedAt,
			mockProduct.UpdatedAt,
			mockProduct.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		a := productRepo.NewProductRepository(db)

		err := a.Update(context.TODO(),mockProduct)
		assert.NoError(t, err)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		query := "UPDATE `products` SET " + querySet + " " + queryWhere
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectBegin()
		mock.ExpectExec(queryRegex).
			WithArgs(
				mockProduct.ProductName,
				mockProduct.Price,
				mockProduct.Stock,
				mockProduct.CreatedAt,
				mockProduct.UpdatedAt,
				mockProduct.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		a := productRepo.NewProductRepository(db)

		err := a.Update(context.TODO(),mockProduct)
		assert.Error(t, err)
	})


}

func TestProductRepository_Store(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	//requestMock
	mockProductTest := domain.ProductTest{}
	err = faker.FakeData(&mockProductTest)
	assert.NoError(t, err)

	mockProduct := mockProductTest.ToProduct()

	querySet := "`product_name`,`product_price`,`active_sale`,`stock`,`created_at`,`updated_at`,`id`"


	t.Run("success", func(t *testing.T) {
		query := "INSERT INTO `products` (" + querySet + ") VALUES (?,?,?,?,?,?,?)"
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectBegin()
		mock.ExpectExec(queryRegex).
			WithArgs(
				mockProduct.ProductName,
				mockProduct.Price,
				mockProduct.ActiveSale,
				mockProduct.Stock,
				mockProduct.CreatedAt,
				mockProduct.UpdatedAt,
				mockProduct.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		a := productRepo.NewProductRepository(db)

		err := a.Store(context.TODO(),mockProduct)
		assert.NoError(t, err)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		query := "INSERT `products` SET " + querySet
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectBegin()
		mock.ExpectExec(queryRegex).
			WithArgs(
				mockProduct.ProductName,
				mockProduct.Price,
				mockProduct.Stock,
				mockProduct.CreatedAt,
				mockProduct.UpdatedAt,
				mockProduct.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		a := productRepo.NewProductRepository(db)

		err := a.Store(context.TODO(),mockProduct)
		assert.Error(t, err)
	})



}

func TestProductRepository_Delete(t *testing.T) {
	err := test.NewMockEnv()
	db, mock, err := test.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	//resultMock
	mockProduct := domain.ProductTest{}
	err = faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	id  := mockProduct.Id

	queryWhere := "id =?"

	t.Run("success", func(t *testing.T) {
		query := "delete from products where " + queryWhere
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectExec(queryRegex).
			WithArgs(mockProduct.Id).
			WillReturnResult(sqlmock.NewResult(int64(mockProduct.Id), 1))
		a := productRepo.NewProductRepository(db)

		 err := a.Delete(context.TODO(),int(id))
		assert.NoError(t, err)
	})

	t.Run("error-regex-query", func(t *testing.T) {
		query := "DELETE FROM `products` WHERE " + queryWhere
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectExec(queryRegex).
			WithArgs(mockProduct.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		a := productRepo.NewProductRepository(db)

		 err := a.Delete(context.TODO(),int(id))
		assert.Error(t, err)
	})


}
