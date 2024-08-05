package main

import (
	"fmt"
	Aes_crypt "github.com/navy-birds-MRS/Go_Demo/Cryptography"
)

func main() {
	data := "hello everyone"
	encrypted_data, secreat_key, _ := Aes_crypt.EncryptByAes([]byte(data))
	fmt.Println("key :" + secreat_key + "\nvalue :" + encrypted_data+"\nNote: Please store the key value some where safe.")

	message,err_msg, _ := Aes_crypt.DecryptAES(encrypted_data,secreat_key)
	fmt.Println("message : " + string(message), err_msg)
}
