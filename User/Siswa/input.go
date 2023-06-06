package siswa


type RegisterUserInput struct {
	NIS           string    `json:"nis" binding:"required"`
	Fullname      string    `json:"fullname" binding:"required"`
	Username      string    `json:"username" binding:"required"`
	Kelas         string    `json:"kelas" binding:"required"`
	Alamat        string    `json:"alamat" binding:"required"`
	Verif         string    `json:"verif" binding:"required"`
}
