package routes

import (
	"meeting3/database"
	"meeting3/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	//Membuat Inputan JSON
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//Encrip password
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	NoTelpon, _ := strconv.Atoi(data["no_telpon"])

	user := models.User{
		NamaLengkap: data["nama_lengkap"],
		NoTelpon:    NoTelpon,
		Email:       data["email"],
		Password:    string(password),
	}

	var email models.User

	database.DB.Where("email = ?", user.Email).Find(email)

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil insert data d",
		"data":  user,
	})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	//Mengambil data berdasarkan Email
	database.DB.Where("email = ? ", data["email"]).Find(&user)

	if user.IdUser == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Pesan": " Tidak Ada di dalam Database",
		})
	}

	//Compare Hass Password

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Pesan": "Password Tidak Sesuai",
		})
	} else {
		//Kita Akan Buatkan Sebuah JWT Token Barel
        claims := jwt.MapClaims{
            "id_user": user.IdUser,
            "role":    3123,
            "exp":     time.Now().Add(time.Hour * 48).Unix(),
            "status":  "Mencoba JWT",
        }
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        t, err := token.SignedString([]byte("Rahasia"))

        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "Pesan":"Gagal Membuat JWT",
            })
        }
        return c.JSON(fiber.Map{
            "token": t,
        })
    }
	
}
