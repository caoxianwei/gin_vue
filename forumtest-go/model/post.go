package model

type Post struct {
	//ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	ID        uint   `json:"id" gorm:"primary_key"`
	UserId uint `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	//Category   []Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp"`
}

//新版gorm不支持gorm.Scope了，BeforeCreate的uuid
//func (post *Post) BeforeCreate(scope *gorm.Scope) error {
//	return scope.SetColumn("ID", uuid.NewV4())
//}

