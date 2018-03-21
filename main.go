// main.go

package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	//a.Initialize("root", "root", "doodh_test")
	//a.Initialize("dbuser", "!@#$%^doodhw", "doodhwala.net:3306", "DOODHWALA_BCK_FEB13")
	a.Initialize("root", "root", "localhost:3306", "go_notes")
	//username:password@protocol(address)/dbname?param=value

	a.Run(":8000")
}
