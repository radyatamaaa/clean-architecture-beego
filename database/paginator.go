package database

import (
	"context"
	"gorm.io/gorm"
	"math"
)

// Paginator structure containing pagination information and result records.
// Can be sent to the client directly.
type Paginator struct {
	db *gorm.DB

	BackPage int
	NextPage     int
	MaxPage     int64
	Total       int64
	PageSize    int
	CurrentPage int
	Records     interface{}
}

func paginateScope(ctx context.Context,page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.WithContext(ctx).Offset(offset).Limit(pageSize)
	}
}

// NewPaginator create a new Paginator.
//
// Given DB transaction can contain clauses already, such as WHERE, if you want to
// filter results.
//
//  articles := []model.Article{}
//  tx := database.Conn().Where("title LIKE ?", "%"+helper.EscapeLike(search)+"%")
//  paginator := database.NewPaginator(tx, page, pageSize, &articles)
//  result := paginator.Find()
//  if response.HandleDatabaseError(result) {
//      response.JSON(http.StatusOK, paginator)
//  }
//
func NewPaginator(db *gorm.DB, page, pageSize int, dest interface{}) *Paginator {
	return &Paginator{
		db:          db,
		CurrentPage: page,
		PageSize:    pageSize,
		Records:     dest,
	}
}

func (p *Paginator) updatePageInfo() {
	count := int64(0)
	if err := p.db.Model(p.Records).Count(&count).Error; err != nil {
		panic(err)
	}
	p.Total = count
	p.MaxPage = int64(math.Ceil(float64(count) / float64(p.PageSize)))
	if p.MaxPage == 0 {
		p.MaxPage = 1
	}
	prev := p.CurrentPage
	next := p.CurrentPage
	if p.CurrentPage != 1 {
		prev = p.CurrentPage - 1
	}

	if p.CurrentPage != int(p.MaxPage)  {
		next = p.CurrentPage + 1
	}

	p.NextPage = next
	p.BackPage = prev
}

// Find requests page information (total records and max page) and
// executes the transaction. The Paginate struct is updated automatically, as
// well as the destination slice given in NewPaginate().
func (p *Paginator) Find(ctx context.Context) *gorm.DB {
	p.updatePageInfo()
	return p.db.Scopes(paginateScope(ctx,p.CurrentPage, p.PageSize)).Find(p.Records)
}
