/*
   @Auther feiwang
   @Time 2019/12/1  11:27
*/
package main

import (
	cron "gocron/PPGo_Job-master/crons"
)
func main() {
	println("dfasdfasd")
	cron.New(cron.WithSeconds())
}
