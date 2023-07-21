package dbcreationscript

var TABLE_CREATION_SQLQUERY string = `

CREATE TABLE IF NOT EXISTS sd_lookup_lesson (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  lesson varchar(100) NOT NULL
);
CREATE TABLE IF NOT EXISTS cd_chapter_details (
  chapter_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  chapter_name varchar(20) NOT NULL,
  country varchar(20) NOT NULL,
  city varchar(20) NOT NULL,
  full_address varchar(100) NOT NULL,
  zip_code int NOT NULL,
  latidude float NOT NULL,
  longitude float NOT NULL,
  created_by int NOT NULL,
  last_updated_by int NOT NULL,
  creation_date timestamp NOT NULL,
  last_updated_date timestamp NOT NULL,
  KEY (chapter_id)
);

CREATE TABLE IF NOT EXISTS pd_project_details (
  project_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  project_name varchar(50) NOT NULL,
  project_link varchar(50) NOT NULL,
  project_year year NOT NULL,
  chapter_id int DEFAULT NULL,
  problem_statement varchar(100) NOT NULL,
  deliverables varchar(50) NOT NULL,
  created_by int NOT NULL,
  last_updated_by int NOT NULL,
  creation_date timestamp NOT NULL,
  last_updated_date timestamp NOT NULL,
  KEY (project_id),
  KEY (chapter_id),
  CONSTRAINT pd_project_details_ibfk_1 FOREIGN KEY (chapter_id) REFERENCES cd_chapter_details (chapter_id) ON DELETE SET NULL ON UPDATE SET NULL
);


CREATE TABLE IF NOT EXISTS sd_lookup_role (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  role_name varchar(20) NOT NULL
);


CREATE TABLE IF NOT EXISTS sd_lookup_signature (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  signature_by int NOT NULL,
  signature_link varchar(50) NOT NULL
);


CREATE TABLE IF NOT EXISTS sd_lookup_team (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  team_name varchar(20) NOT NULL
);



CREATE TABLE IF NOT EXISTS sd_student_details (
  user_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_name varchar(50) NOT NULL,
  email varchar(50) NOT NULL,
  role_id int DEFAULT NULL,
  spent_hours int NOT NULL,
  raised int NOT NULL COMMENT 'K$',
  user_team_id int DEFAULT NULL,
  project_id int DEFAULT NULL,
  signature_id int DEFAULT NULL,
  created_by int NOT NULL,
  last_updated_by int NOT NULL,
  creation_date timestamp NOT NULL,
  last_updated_date timestamp NOT NULL,
  KEY (project_id),
  KEY (signature_id),
  KEY (role_id),
  KEY (user_team_id),
  KEY (user_id),
  CONSTRAINT sd_student_details_ibfk_1 FOREIGN KEY (project_id) REFERENCES pd_project_details (project_id) ON DELETE SET NULL ON UPDATE SET NULL,
  CONSTRAINT sd_student_details_ibfk_2 FOREIGN KEY (role_id) REFERENCES sd_lookup_role (id) ON DELETE SET NULL ON UPDATE SET NULL,
  CONSTRAINT sd_student_details_ibfk_3 FOREIGN KEY (user_team_id) REFERENCES sd_lookup_team (id) ON DELETE SET NULL ON UPDATE SET NULL,
  CONSTRAINT sd_student_details_ibfk_4 FOREIGN KEY (signature_id) REFERENCES sd_lookup_signature (id) ON DELETE SET NULL ON UPDATE SET NULL
);
CREATE TABLE IF NOT EXISTS sd_student_profile (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id int DEFAULT NULL,
  linkedin_url varchar(50) NOT NULL,
  social_media_link varchar(100) ,
  photos_uploaded varchar(500) ,
  school_attended varchar(100) ,
  business_ideas varchar(500) ,
  impactful_experience varchar(500) ,
  future_vision varchar(500) ,
  current_project_involvement varchar(2000) ,
  success_definition varchar(500) ,
  tye_experience_impact varchar(500) ,
  biggest_goals varchar(500) ,
  memorable_experience_tye varchar(500) ,
  hobbies varchar(500) ,
  favorite_song varchar(100) ,
  favorite_movie varchar(100) ,
  deserted_island_item varchar(500) ,
  favorite_food varchar(100) ,
  additional_comments varchar(500) ,
  new_learned_skills varchar(500) ,
  program_challenges_overcome varchar(1000) ,
  team_role varchar(500) ,
  visited_companies varchar(100) NOT NULL,
  your_product varchar(500) ,
  weekly_project_hours varchar(100) ,
  problem_solved varchar(500) ,
  tye_business_name varchar(100) NOT NULL,
  created_by int NOT NULL,
  last_updated_by int NOT NULL,
  creation_date timestamp NOT NULL,
  last_updated_date timestamp NOT NULL,
  KEY (user_id),
  CONSTRAINT sd_student_profile_ibfk_1 FOREIGN KEY (user_id) REFERENCES sd_student_details (user_id) ON DELETE SET NULL ON UPDATE SET NULL
);

CREATE TABLE IF NOT EXISTS sd_map_user_lesson (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id int DEFAULT NULL,
  lesson_id int DEFAULT NULL,
  created_by int NOT NULL,
  last_updated_by int NOT NULL,
  creation_date timestamp NOT NULL,
  last_updated_date timestamp NOT NULL,
  KEY (user_id),
  KEY (lesson_id),
  KEY (id),
  CONSTRAINT sd_map_user_lesson_ibfk_1 FOREIGN KEY (lesson_id) REFERENCES sd_lookup_lesson (id) ON DELETE SET NULL ON UPDATE SET NULL
);
  

`
