# Simple Automate Absent with Go

* https://www.easycron.com/
* https://heroku.com

<p>because using the free version of the heroku account so many restrictions are obtained, therefore additional tools use easycron to run the scheduler.</p>

## ENV
```
EMAIL=xxx
PASSWORD=xx
DEVICE_ID=xxx
LATITUDE=-6.3443904
LONGITUDE=106.8705516
SECRET_KEY=xxx          //for middleware service 
```

## Setup
* SetUp heroku 
* SetUp easycron
* SetUp Env file
* Change file selfie.JPG to your own photo

## API
```
/ping --> ping service
/checkin --> checkin absent
/checkout --> checkout absent
```

## Running with Your Cron
```go
//by default serve rest api
// CRON_CHECKIN = "* 08 * * *"
// CRON_CHECKOUT = "* 20 * * *"
func main() {
	log.Println("Start")
	helper.Dispatcher.Run()
	//UNCOMMENT THIS FUNCTION IF YOU RUNNING WITH YOUR OWN CRON
	cron.RunJob()

	//running with rest api
	//gate.Route()

	log.Println("Finish")
}
```
