package main

import (
	"github.com/sugandhasaxena1911/GetCatFacts/controllers"
	"os"
)

func init() {
	os.Setenv("APIFILEPATH", "../Files/")
	os.Setenv("APIFILENAME", "GetCatFacts.json")
	os.Setenv("PERIOD", "5000")

}

func main() {
	controllers.PeriodCalls()

}
