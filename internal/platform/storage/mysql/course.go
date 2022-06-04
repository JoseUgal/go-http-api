package mysql

/* DTO for MySQL */

// Is the name of our MySQL table.
const (
	sqlCourseTable = "courses"
)

// DTO for mapping MySQL objects.
type sqlCourse struct {
	ID		 string `db:"id"`
	Name	 string `db:"name"`
	Duration string `db:"duration"`
}