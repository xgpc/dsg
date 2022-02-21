package models

type UserSession struct {
	UserID        int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	AccountStatus uint8  `gorm:"NOT NULL;DEFAULT:1"`
	UserName      string `gorm:"NOT NULL"`

	Mobile   string `gorm:"NOT NULL"`
	OpenId   string `gorm:"NOT NULL"`
	UnionId  string `gorm:"NOT NULL"`
	NickName string `gorm:"NOT NULL"`
	UserImg  string `gorm:"NOT NULL"`
	Gender   uint8  `gorm:"NOT NULL;DEFAULT:0"`
	Country  string `gorm:"NOT NULL"`
	Province string `gorm:"NOT NULL"`
	City     string `gorm:"NOT NULL"`
}
