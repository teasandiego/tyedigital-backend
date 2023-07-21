package mysql

import (
	"log"

	model "github.com/Kushmanda-Tech/tyedigital_backend/Model"
)

func GetProjectDetailsFromDb(project_id int) []model.GetProjectDetails {
	var chapterDetailsValue model.GetProjectDetails
	chapterDetailsValue_array := []model.GetProjectDetails{}
	SqlQuery := `SELECT project_id, project_name, project_link, project_year, pd.chapter_id, problem_statement, deliverables, ccd.chapter_name FROM pd_project_details pd INNER JOIN cd_chapter_details ccd ON ccd.chapter_id = pd.chapter_id WHERE project_id = ?`

	row, err := DB.Query(SqlQuery, project_id)

	if err != nil {
		log.Println(err.Error())
		return chapterDetailsValue_array
	}
	for row.Next() {
		err := row.Scan(&chapterDetailsValue.ProjectId, &chapterDetailsValue.ProjectName, &chapterDetailsValue.ProjectLink, &chapterDetailsValue.ProjectYear, &chapterDetailsValue.ChapterId, &chapterDetailsValue.ProblemStatement, &chapterDetailsValue.Deliverables, &chapterDetailsValue.ChapterName)
		if err != nil {
			log.Println(err.Error())
			return chapterDetailsValue_array
		}
		chapterDetailsValue_array = append(chapterDetailsValue_array, chapterDetailsValue)
	}
	defer row.Close()
	return chapterDetailsValue_array
}
