package model

import "time"

type (
	Lesson struct {
		Id     int    `json:"id"`
		Lesson string `json:"lesson"`
	}
	Project struct {
		ProjectId        int    `json:"project_id"`
		ProjectName      string `json:"project_name"`
		ProjectLink      string `json:"project_link"`
		ProjectYear      string `json:"project_year"`
		ChapterId        int    `json:"chapter_id"`
		ProblemStatement string `json:"problem_statement"`
		Deliverables     string `json:"deliverables"`
	}
	GetProjectDetails struct {
		ProjectId        int    `json:"project_id"`
		ProjectName      string `json:"project_name"`
		ProjectLink      string `json:"project_link"`
		ProjectYear      string `json:"project_year"`
		ChapterId        int    `json:"chapter_id"`
		ChapterName      string `json:"chapter_name"`
		ProblemStatement string `json:"problem_statement"`
		Deliverables     string `json:"deliverables"`
	}
	Role struct {
		Id       int    `json:"id"`
		RoleName string `json:"role_name"`
	}
	Signature struct {
		Id            int    `json:"id"`
		SignatureBy   int    `json:"signature_by"`
		SignatureLink string `json:"signature_link"`
	}
	Team struct {
		Id       int    `json:"id"`
		TeamName string `json:"team_name"`
	}
	StudentDetails struct {
		UserID          int    `json:"user_id"`
		UserName        string `json:"user_name"`
		RoleID          int    `json:"role_id"`
		Email           string `json:"email"`
		SpentHours      int    `json:"spent_hours"`
		Raised          int    `json:"raised"`
		UserTeamId      int    `json:"user_team_id"`
		ProjectId       int    `json:"project_id"`
		SignatureId     int    `json:"signature_id"`
		CreatedBy       int    `json:"created_by"`
		LastUpdatedBy   int    `json:"last_updated_by"`
		CreationDate    int    `json:"creation_date"`
		LastUpdatedDate int    `json:"last_updated_date"`
	}
	StudentProfile struct {
		ID                        int    `json:"id"`
		UserID                    int    `json:"user_id"`
		LinkedinUrl               string `json:"linkedin_url"`
		SocialMediaLink           string `json:"social_media_link"`
		PhotosUploaded            string `json:"photos_uploaded"`
		SchoolAttended            string `json:"school_attended"`
		BusinessIdeas             string `json:"business_ideas"`
		ImpactfulExperience       string `json:"impactful_experience"`
		FutureVision              string `json:"future_vision"`
		CurrentProjectInvolvement string `json:"current_project_involvement"`
		SuccessDefinition         string `json:"success_definition"`
		TyeExperienceImpact       string `json:"tye_experience_impact"`
		BiggestGoals              string `json:"biggest_goals"`
		MemorableExperienceTye    string `json:"memorable_experience_tye"`
		Hobbies                   string `json:"hobbies"`
		FavoriteSong              string `json:"favorite_song"`
		FavoriteMovie             string `json:"favorite_movie"`
		DesertedIslandItem        string `json:"deserted_island_item"`
		FavoriteFood              string `json:"favorite_food"`
		AdditionalComments        string `json:"additional_comments"`
		NewLearnedSkills          string `json:"new_learned_skills"`
		ProgramChallengesOvercome string `json:"program_challenges_overcome"`
		TeamRole                  string `json:"team_role"`
		VisitedCompanies          string `json:"visited_companies"`
		YourProduct               string `json:"your_product"`
		WeeklyProjectHours        int    `json:"weekly_project_hours"`
		ProblemSolved             string `json:"problem_solved"`
		TyeBusinessName           string `json:"tye_business_name"`
	}
	GetStudentDetails struct {
		UserID                    int    `json:"user_id"`
		UserName                  string `json:"user_name"`
		RoleID                    int    `json:"role_id"`
		Email                     string `json:"email"`
		SpentHours                int    `json:"spent_hours"`
		Raised                    int    `json:"raised"`
		UserTeamId                int    `json:"user_team_id"`
		ProjectId                 int    `json:"project_id"`
		SignatureId               int    `json:"signature_id"`
		RoleName                  string `json:"role_name"`
		TeamName                  string `json:"team_name"`
		ProjectName               string `json:"project_name"`
		ProjectLink               string `json:"project_link"`
		ProjectYear               int64  `json:"project_year"`
		SignatureBy               int    `json:"signature_by"`
		SignatureLink             string `json:"signature_link"`
		ProblemStatement          string `json:"problem_statement"`
		Deliverables              string `json:"deliverables"`
		LinkedinUrl               string `json:"linkedin_url"`
		SocialMediaLink           string `json:"social_media_link"`
		PhotosUploaded            string `json:"photos_uploaded"`
		SchoolAttended            string `json:"school_attended"`
		BusinessIdeas             string `json:"business_ideas"`
		ImpactfulExperience       string `json:"impactful_experience"`
		FutureVision              string `json:"future_vision"`
		CurrentProjectInvolvement string `json:"current_project_involvement"`
		SuccessDefinition         string `json:"success_definition"`
		TyeExperienceImpact       string `json:"tye_experience_impact"`
		BiggestGoals              string `json:"biggest_goals"`
		MemorableExperienceTye    string `json:"memorable_experience_tye"`
		Hobbies                   string `json:"hobbies"`
		FavoriteSong              string `json:"favorite_song"`
		FavoriteMovie             string `json:"favorite_movie"`
		DesertedIslandItem        string `json:"deserted_island_item"`
		FavoriteFood              string `json:"favorite_food"`
		AdditionalComments        string `json:"additional_comments"`
		NewLearnedSkills          string `json:"new_learned_skills"`
		ProgramChallengesOvercome string `json:"program_challenges_overcome"`
		TeamRole                  string `json:"team_role"`
		VisitedCompanies          string `json:"visited_companies"`
		YourProduct               string `json:"your_product"`
		WeeklyProjectHours        string `json:"weekly_project_hours"`
		ProblemSolved             string `json:"problem_solved"`
		TyeBusinessName           string `json:"tye_business_name"`
	}
	Chapter struct {
		ChapterId   int    `json:"chapter_id"`
		ChapterName string `json:"chapter_name"`
		Country     string `json:"country"`
		City        string `json:"city"`
		FullAddress string `json:"full_address"`
		ZipCode     int64  `json:"zip_code"`
		Latidude    int64  `json:"latidude"`
		Longitude   int64  `json:"longitude"`
	}
	MapUserLesson struct {
		Id       int `json:"id"`
		UserId   int `json:"user_id"`
		LessonId int `json:"lesson_id"`
	}
	PasetoClaimData struct {
		Exp time.Time `json:"exp"`
		Iat time.Time `json:"iat"`
		Nbf time.Time `json:"nbf"`

		Aud    string `json:"aud"`
		Email  string `json:"email"`
		Iss    string `json:"iss"`
		Jti    string `json:"jti"`
		Sub    string `json:"sub"`
		UserID string `json:"user_id"`
	}
)
