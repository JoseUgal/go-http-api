package mysql

import (
	"context"
	"database/sql"
	"fmt"

	mooc "github.com/JoseUgal/go-http-api/internal"
	"github.com/huandu/go-sqlbuilder"
)

// CourseRepository is a MySQL mooc.CouseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *CourseRepository) Save( ctx context.Context, course mooc.Course) error {

	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))

	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID: 	  course.ID().String(),
		Name: 	  course.Name().String(),
		Duration: course.Duration().String(),	
	}).Build()

	fmt.Println(query)
	fmt.Println(args)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}
