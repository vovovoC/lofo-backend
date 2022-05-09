package transfer

type PostUpdateDTO struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	Image       string `gorm:"type:text" json:"image"`
	Place       string `gorm:"type:text" json:"place"`
	Category    string `gorm:"type:varchar(255)" json:"category"`
	Status      string `gorm:"type:varchar(255)" json:"status"`
	Datetime    string `gorm:"type:varchar(255)" json:"datetime"`
}

type PostCreateDTO struct {
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	Image       string `gorm:"type:text" json:"image"`
	Place       string `gorm:"type:text" json:"place"`
	Category    string `gorm:"type:varchar(255)" json:"category"`
	Status      string `gorm:"type:varchar(255)" json:"status"`
	Datetime    string `gorm:"type:varchar(255)" json:"datetime"`
}
