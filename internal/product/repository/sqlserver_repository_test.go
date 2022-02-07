package repository_test

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/helpers/test"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	productRepo "clean-architecture-beego/internal/product/repository"
)

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
