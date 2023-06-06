package siswa

import "gorm.io/gorm"

type Repository interface {
	Save(user Siswa) (Siswa, error)
	FindByID(id int) (Siswa, error)
	Update(user Siswa) (Siswa, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) Save(user Siswa) (Siswa, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (Siswa, error) {
	var user Siswa
	err := r.db.Where("id_user =?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *repository) Update(user Siswa) (Siswa, error) {
	err := repo.db.Model(&Siswa{}).Where("id_user = ?", user.IdUser).Updates(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
