package models

type Admin struct {
	IdAdmin uint
	Nama    string
}

type User struct {
	IdUser      uint
	NamaLengkap string
	NoTelpon    int
	Email       string
	Password    string
}
