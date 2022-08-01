package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Education struct {
	Institute string `json:"institute"`
	Course    string `json:"course"`
	Grade     string `json:"grade"`
	Code      string `json:"code"`
}

type WorkExperience struct {
	Organisation string         `json:"organisation"`
	Role         string         `json:"role"`
	StartDate    datatypes.Date `json:"startDate"`
	EndDate      datatypes.Date `json:"endDate"`
}

type EducationSlice []Education
type WorkExperienceSlice []WorkExperience

type User struct {
	gorm.Model     `json:"-"`
	ObjectId       string              `gorm:"" json:"objectId"`
	Name           string              `json:"name"`
	Age            int                 `json:"age"`
	Gender         string              `json:"gender"` // Male, Female, Other
	Role           string              `json:"role"`   // CA, PA, SA
	ImageUrl       string              `json:"imageUrl"`
	Email          string              `json:"email" gorm:"unique"`
	Password       string              `json:"-"`
	Todos          []Todo              `json:"todos"`
	Interest       datatypes.JSON      `json:"interest"`
	Social         datatypes.JSON      `json:"social"`
	Education      EducationSlice      `json:"education"`
	WorkExperience WorkExperienceSlice `json:"workExperience"`
}
