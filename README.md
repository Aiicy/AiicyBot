
[![Build Status](https://travis-ci.org/Aiicy/mybot.svg?branch=master)](https://travis-ci.org/Aiicy/mybot)

# mybot

# config.yaml 

```yaml
slacktoken: YOU_SLACK_BOT_API_TOKEN_HERE
ReportChannel: '#chanel_name'
botfather:  U3K9E0HRA  #botfather_id, use @botname whoami to get it
BIJIN_TIMEZONE: 'Asia/Shanghai'
```

# how to build it 
	$git clone https://github.com/Aiicy/mybot.git
	$cd mybot
	$go get -d -v -t
	$go build -o slackbot

# how to run 
	1. after build # use config.yaml to run the app
	   $./slackbot start
	2. you can use [setenv.sh](/setenv.sh) to run the app
	$source ./setenv.sh && ./slackbot start

# set up the sent sexy mm func
	1. python pic_dl.py
	2. copy all the file after download && put to images folder
	



