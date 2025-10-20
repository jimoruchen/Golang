package models

type ClassModel struct {
	ID          int
	Name        string         `gorm:"size:12"`
	StudentList []StudentModel `gorm:"foreignkey:ClassID"`
}

type StudentModel struct {
	ID         int
	Name       string `gorm:"size:12"`
	ClassID    int
	ClassModel *ClassModel `gorm:"foreignkey:ClassID"`
}
