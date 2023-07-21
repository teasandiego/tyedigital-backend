package mysql

import (
	"log"

	model "github.com/Kushmanda-Tech/tyedigital_backend/Model"
)

func Get_Student_Details_By_User_Id_FrDb(user_id int) []model.GetStudentDetails {
	var user_data model.GetStudentDetails
	user_data_array := []model.GetStudentDetails{}
	SqlQuery := "SELECT sd.user_id,sd.user_name,sd.email,sd.role_id,sd.spent_hours,sd.raised,sd.user_team_id,sd.project_id,sd.signature_id,slr.role_name,slt.team_name,ppd.project_id,ppd.project_name,ppd.project_link ,ppd.project_year,ppd.problem_statement,ppd.deliverables,sls.signature_by ,sls.signature_link ,ssp.linkedin_url, ssp.social_media_link, ssp.photos_uploaded, ssp.school_attended, ssp.business_ideas, ssp.impactful_experience, ssp.future_vision, ssp.current_project_involvement, ssp.success_definition, ssp.tye_experience_impact, ssp.biggest_goals, ssp.memorable_experience_tye, ssp.hobbies, ssp.favorite_song, ssp.favorite_movie, ssp.deserted_island_item, ssp.favorite_food, ssp.additional_comments, ssp.new_learned_skills, ssp.program_challenges_overcome, ssp.team_role, ssp.visited_companies, ssp.your_product, ssp.weekly_project_hours, ssp.problem_solved, ssp.tye_business_name FROM sd_student_details sd INNER JOIN sd_lookup_role slr ON sd.role_id = slr.id INNER JOIN sd_lookup_team slt ON sd.user_team_id = slt.id INNER JOIN pd_project_details ppd  ON sd.project_id = ppd.project_id INNER JOIN sd_lookup_signature sls  ON sd.signature_id = sls.id INNER JOIN sd_student_profile ssp ON sd.user_id = ssp.user_id WHERE sd.user_id = ?"
	if err := DB.QueryRow(SqlQuery, user_id).Scan(&user_data.UserID, &user_data.UserName, &user_data.Email, &user_data.RoleID, &user_data.SpentHours, &user_data.Raised, &user_data.UserTeamId, &user_data.ProjectId, &user_data.SignatureId, &user_data.RoleName, &user_data.TeamName, &user_data.ProjectId, &user_data.ProjectName, &user_data.ProjectLink, &user_data.ProjectYear, &user_data.ProblemStatement, &user_data.Deliverables, &user_data.SignatureBy, &user_data.SignatureLink, &user_data.LinkedinUrl, &user_data.SocialMediaLink, &user_data.PhotosUploaded, &user_data.SchoolAttended, &user_data.BusinessIdeas, &user_data.ImpactfulExperience, &user_data.FutureVision, &user_data.CurrentProjectInvolvement, &user_data.SuccessDefinition, &user_data.TyeExperienceImpact, &user_data.BiggestGoals, &user_data.MemorableExperienceTye, &user_data.Hobbies, &user_data.FavoriteSong, &user_data.FavoriteMovie, &user_data.DesertedIslandItem, &user_data.FavoriteFood, &user_data.AdditionalComments, &user_data.NewLearnedSkills, &user_data.ProgramChallengesOvercome, &user_data.TeamRole, &user_data.VisitedCompanies, &user_data.YourProduct, &user_data.WeeklyProjectHours, &user_data.ProblemSolved, &user_data.TyeBusinessName); err != nil {
		log.Println("Get_User_Details_By_User_Id_FrDb Scan ", err)
	}
	// user_data.FCM_Token = fcm_token.String
	user_data_array = append(user_data_array, user_data)
	return user_data_array
}
