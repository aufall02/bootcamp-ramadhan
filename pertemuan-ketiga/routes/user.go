package routes

import (
	
	"meeting3/models"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	//Membuat Inputan JSON 
	var data map[string]string
	if err := c.BodyParser(&data); err != nil{
		return err
	}
	//Encrip password 


	NoTelpon,_ := strconv.Atoi(data["no_telpon"])

	user := models.User{
		NamaLengkap: data["nama_lengkap"],
		NoTelpon: NoTelpon,
		Email: data["email"],
		Password: data["password"],
	}

	return c.JSON(fiber.Map{
		"Pesan":"Berhasil",
		"data": user,
	})
}

func Login(){
	//Menggunakan Jwt menampilkan JWT
}
