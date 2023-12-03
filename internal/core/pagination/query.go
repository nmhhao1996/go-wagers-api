package pagination

import "github.com/gin-gonic/gin"

const (
	defaultPage  = 1
	defaultLimit = 10
)

// Query is the query struct for pagination
type Query struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// Offset returns the offset
func (q Query) Offset() int {
	return (q.Page - 1) * q.Limit
}

func (q *Query) adjust() {
	if q.Page <= 0 {
		q.Page = defaultPage
	}

	if q.Limit <= 0 {
		q.Limit = defaultLimit
	}
}

// GetPaginationQueryFromContext gets the pagination query from context
func GetPaginationQueryFromContext(c *gin.Context) (Query, error) {
	var q Query
	if err := c.ShouldBindQuery(&q); err != nil {
		return q, err
	}

	q.adjust()

	return q, nil
}
