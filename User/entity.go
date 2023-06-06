package user

import "time"

type User struct {
	ID            int      
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

func (User) TableName() string {
	return "user"
}
