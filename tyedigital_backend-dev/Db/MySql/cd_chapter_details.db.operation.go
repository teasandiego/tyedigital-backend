package mysql

import (
	"database/sql"
	"log"

	model "github.com/Kushmanda-Tech/tyedigital_backend/Model"
)

func ValidateChapterId(chapter_id string) bool {
	var dbplanetID int
	SqlQuery := `SELECT chapter_id FROM cd_lookup_chapter WHERE chapter_id = ?`
	if err := DB.QueryRow(SqlQuery, chapter_id).Scan(&dbplanetID); err != nil && err == sql.ErrNoRows {
		log.Println("ValidateChapterId QueryRow ", err.Error())
		return false
	}
	return true
}
func GetChapterDetailsFromDb(chapter_id int) []model.Chapter {
	var chapterDetailsValue model.Chapter
	chapterDetailsValue_array := []model.Chapter{}
	SqlQuery := `SELECT chapter_id, chapter_name, country,city,full_address,zip_code,latidude,longitude  FROM cd_chapter_details WHERE chapter_id = ?`

	row, err := DB.Query(SqlQuery, chapter_id)

	if err != nil {
		log.Println(err.Error())
		return chapterDetailsValue_array
	}
	for row.Next() {
		err := row.Scan(&chapterDetailsValue.ChapterId, &chapterDetailsValue.ChapterName, &chapterDetailsValue.Country, &chapterDetailsValue.City, &chapterDetailsValue.FullAddress, &chapterDetailsValue.ZipCode, &chapterDetailsValue.Latidude, &chapterDetailsValue.Longitude)
		if err != nil {
			log.Println(err.Error())
			return chapterDetailsValue_array
		}
		chapterDetailsValue_array = append(chapterDetailsValue_array, chapterDetailsValue)
	}
	defer row.Close()
	return chapterDetailsValue_array
}
