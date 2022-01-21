package repository_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"

	//"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"clean-architecture-beego/helper"
	"clean-architecture-beego/helper/logger"
	"clean-architecture-beego/models"
	sampleModuleRepo "clean-architecture-beego/sample_module/repository"
	"regexp"
	"testing"
)

var (
	mockSampleModuleList = []models.SampleModule{}
)

func TestSampleModuleRepository_List(t *testing.T) {
	_, mock, err := helper.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//resultMock
	mockSampleModule := models.SampleModule{}
	mockSampleModule = mockSampleModule.MappingExpampleData()

	mockSampleModuleList = append(mockSampleModuleList, mockSampleModule)

	//rowAndColumnMock
	row, fields := helper.GetValueAndColumnStructToDriverValue(mockSampleModuleList[0])

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(fields).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...)
		query := `SELECT * FROM "YOUR TABLE" WHERE (region=?) AND (mainbr=?) AND (branch=?) AND (pernr=?) AND (kategori_churn = ?) AND (acctno = ?)`
		queryRegex := regexp.QuoteMeta(query)
		//fmt.Println(matched)
		mock.ExpectQuery(queryRegex).WillReturnRows(rows)
		a := sampleModuleRepo.NewSampleModuleRepository(nil, logger.L)

		sampleModule, err := a.List(context.TODO(), 10, 0)
		assert.NoError(t, err)
		assert.NotNil(t, sampleModule)
		assert.Len(t, sampleModule, 10)
	})

}

func TestSampleModuleRepository_Count(t *testing.T) {
	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//dbgorm, err := gorm.Open("avatica", db)

	//resultMock
	mockSampleModule := models.SampleModule{}
	mockSampleModule = mockSampleModule.MappingExpampleData()

	mockSampleModuleList = append(mockSampleModuleList, mockSampleModule)

	//rowAndColumnMock
	row, fields := helper.GetValueAndColumnStructToDriverValue(mockSampleModuleList[0])

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(fields).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...).
			AddRow(row...)
		query := `SELECT * FROM "YOUR TABLE" WHERE (region=?) AND (mainbr=?) AND (branch=?) AND (pernr=?) AND (kategori_churn = ?) AND (acctno = ?)`
		queryRegex := regexp.QuoteMeta(query)
		//fmt.Println(matched)
		mock.ExpectQuery(queryRegex).WillReturnRows(rows)
		a := sampleModuleRepo.NewSampleModuleRepository(nil, logger.L)

		sampleModule, err := a.Count(context.TODO())
		assert.NoError(t, err)
		assert.NotNil(t, sampleModule)
	})

}
