package orm

import (
	"gorm.io/gorm"
)

type ORMFunctions struct {
	db *gorm.DB
}

func NewORMFunctions(db *gorm.DB) *ORMFunctions {
	return &ORMFunctions{db: db}
}

func (o *ORMFunctions) CreateOrReturnUserByUserId(id uint64) (*User, error) {
	var user *User
	err := o.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user = &User{
				Id:          id,
				LanguageISO: "en",
			}
			err = o.db.Create(user).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return user, nil
}

func (o *ORMFunctions) ChangeLanguageByUserId(userId uint64, language string) error {
	return o.db.Model(&User{}).Where("id = ?", userId).Update("LanguageISO", language).Error
}

func (o *ORMFunctions) AddToTPByUserId(userId uint64, totp string, name string) error {
	return o.db.Create(&ToTP{
		Id:     totp,
		Name:   name,
		UserId: userId,
	}).Error
}

func (o *ORMFunctions) FindAllToTPByUserId(userId uint64) ([]ToTP, error) {
	var totps []ToTP
	err := o.db.Model(&ToTP{}).Where("user_id = ?", userId).Find(&totps).Error
	if err != nil {
		return nil, err
	}
	return totps, nil
}

func (o *ORMFunctions) RemoveToTPByUserId(userId uint64, totp string) error {
	return o.db.Model(&ToTP{}).Where("user_id = ? AND id = ?", userId, totp).Delete(&ToTP{}).Error
}
