package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 { return 98000000 /*...*/ }

// Email sender configuration/
// Note: never put passwords in source code!
const sender = "notification@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = "Warning: you are using %d bytes of storage, %d%% of your quota"

// gopl.io/ch11/storage1
// func CheckQuota(username string) {
// 	used := bytesInUse(username)
// 	const quota = 1000000000 // 1GB
// 	percent := 100 * used / quota
// 	if percent < 90 {
// 		return
// 	}
// 	msg := fmt.Sprintf(template, used, percent)
// 	auth := smtp.PlainAuth("", sender, password, hostname)
// 	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
// 	if err != nil {
// 		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
// 	}
// }

// gopl.io/ch11/storage2
// implementing notifyUser as we don't want the test to send out real email
var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 100000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
