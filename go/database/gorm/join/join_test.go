package join

import (
	"testing"

	"learn/database"
	. "learn/database/gorm"
	. "learn/database/u"
)

func TestJoin(t *testing.T) {
	type Department struct {
		ID   string
		Name string
	}
	type Professor struct {
		ID     string
		Name   string
		Salary int

		DepartmentID string
	}
	sess := Connect(database.EnvPort)
	GormMust(sess.Exec("DROP TABLE IF EXISTS department, professor"))

	sess = sess.Debug()
	Must(sess.Migrator().CreateTable(Department{}, Professor{}))

	GormMust(sess.Create([]Department{
		{ID: "1", Name: "Math"},
		{ID: "2", Name: "Physics"},
	}))
	GormMust(sess.Create([]Professor{
		{ID: "1", Name: "John", Salary: 1000, DepartmentID: "1"},
		{ID: "2", Name: "Ferry", Salary: 2000, DepartmentID: "2"},
		{ID: "3", Name: "Terran", Salary: 3000, DepartmentID: "1"},
		{ID: "4", Name: "Bobby", Salary: 4000, DepartmentID: "2"},
	}))

	{
		type JoinedProfessor struct {
			Professor
			Department `gorm:"embedded;embeddedPrefix:department_"`
		}

		var got JoinedProfessor
		GormMust(sess.Model(Professor{}).
			// You cannot omit professor.* here. Gorm complains about ambiguity.
			Select("professor.*, department.id as department_id, department.name as department_name").
			Joins("join department on professor.department_id = department.id").
			Where("professor.id = ?", "3").
			First(&got))

		t.Logf("%+v", got)
	}

	{
		type JoinedProfessor struct {
			Professor

			DepartmentName string
		}

		var got JoinedProfessor
		GormMust(sess.Model(Professor{}).
			Select("*, department.name as department_name").
			Joins("join department on department_id = department.id").
			Where("id = ?", "3").
			First(&got))

		t.Logf("%+v", got)

	}
}
