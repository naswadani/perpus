package siswa

import "time"

type Siswa struct {
	IdUser        int
	KodeUser      string
	NIS           string
	Fullname      string
	Username      string
	Password      string
	Kelas         string
	Alamat        string
	Verif         string
	Role          string
	JoinDate      time.Time
	TerakhirLogin time.Time
}

func (Siswa) TableName() string {
	return "user"
}
