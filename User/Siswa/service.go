package siswa

import (
	"fmt"
	"strconv"
	"time"

	"github.com/iancoleman/strcase"
)

type Service interface {
	RegisterSiswa(input RegisterUserInput) (Siswa, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterSiswa(input RegisterUserInput) (Siswa, error) {
	user := Siswa{}
	user.NIS = input.NIS
	user.Fullname = input.Fullname
	user.Username = input.Username
	user.Password = strcase.ToCamel(input.Username)
	user.Kelas = input.Kelas
	user.Alamat = input.Alamat
	user.Verif = input.Verif
	user.Role = "anggota"
	user.JoinDate = time.Now()
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	//untuk mendapatkan ID untuk kode siswa
	angkaStr := strconv.Itoa(newUser.IdUser)
	newUser.KodeUser = fmt.Sprintf("%s%03s", "AP", angkaStr)

	_, err = s.repository.Update(newUser)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
