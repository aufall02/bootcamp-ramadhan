package utils

import "fmt"

func Pointer() {
    //Mengkopi nilai yang ada 
    var alamat1 string
    alamat1 = "Tangerang banten"


    alamat2 := alamat1

    alamat1 = "Bandung Kabupaten"

    fmt.Println("alamat 1 yaitu adalah:", alamat1)
    fmt.Println("alamat 2 yaitu adalah:", alamat2)

    //----------------Pointer ----------------
    fmt.Println("--------------------")
    var alamat3 string
    alamat3 = "Tangerang banten jaya selalu"


    alamat4 := &alamat3
    alamat3 = "Bandung "

    fmt.Println("alamat 3 yaitu adalah:", alamat3)
    fmt.Println("alamat 4 yaitu adalah:", *alamat4)

}



