package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"time"
)

var MyEmail = "your e-mail address"
var MyPassword = "your e-mail password"
var TargetEmail = "e-mail address, where you'll check your IP"

func main()  {
	var ip string

	auth := smtp.PlainAuth("", MyEmail, MyPassword, "smtp.gmail.com")

	for i:=0; i<1; i=i{
		res, _ := http.Get("https://myexternalip.com/raw")
		bd, _ := ioutil.ReadAll(res.Body)
		sbd := string(bd)

		if ip!=sbd&&sbd!=""{
			ip = sbd
			err := smtp.SendMail("smtp.gmail.com:587", auth, MyEmail, []string{TargetEmail}, bd)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Printf("New IP: %s\n", ip)
		}
		time.Sleep(time.Minute*3)

	}
}