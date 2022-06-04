package mysql

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mooc "github.com/JoseUgal/go-http-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CourseRepository_Save_RepositoryError(t *testing.T) {
	courseID, courseName, courseDuration := "c6675339-8706-4068-a549-eca2208d223a", "Curso React: De 0 a Super-héroe", "3 months"
	
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO (id, name, duration) VALUES (?, ?, ?)").
			WithArgs(courseID, courseName, courseDuration).
			WillReturnError(errors.New("something-failed"))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.Error(t, err)
}

func Test_CourseRepository_Save_Succeed(t *testing.T) {
	courseID, courseName, courseDuration := "c6675339-8706-4068-a549-eca2208d223a", "Curso React: De 0 a Super-héroe", "3 months"
	
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO (id, name, duration) VALUES (?, ?, ?)").
			WithArgs(courseID, courseName, courseDuration).
			WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.Error(t, err)
}
