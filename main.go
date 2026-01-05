package main

import (
	"cronproject/dbconnection"
	"fmt"
	"log"

"cronproject/maildata"
	
)
func main()  {

	
	// Connect to PostgreSQL
	if err := dbconnection.Connectiontopostgres();err != nil{
		log.Fatal(err)
	}
	fmt.Println("db connected successfully")
	 // close pool on exit
err := maildata.SendMailToAllUsers(
		"System Update",
		"<h3>Hello User</h3><p>This is a cron email.</p>",
)
if err != nil {
log.Fatal(err)
}
fmt.Println("email sent succesfully to all users")
fmt.Println("Cron job completed......")
}
