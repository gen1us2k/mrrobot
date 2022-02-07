# Mr. Robot a greeter bot for your slack community

A simple greeter bot for your slack communities. The main reason of this bot to have better onboarding experience for newly joined users. This bot built using [slack events api](https://api.slack.com/apis/connections/events-api) and hosted on [AWS lambda](https://aws.amazon.com/lambda/)

## Local Development
### Clonning the repo 

```
git clone git@github.com:gen1us2k/mrrobot
cd mrrobot
```

The bot has make commands to speedup development

```
   build_docker                   Build docker image
   build_linux                    Build executable for linux system
   lint                           Runs linter against the code
   test                           Run tests locally
   zip                            Build and create a zip archive for deploying to AWS lambda
```

## Deploying to AWS

1. Signin to AWS Console and go to the Lambda service
2. Create a new function from scratch using Go 1.x SDK
3. Run `make zip` to create a zip archive to upload to the AWS
4. Upload a zipfile to run your code

### Configuring AWS API Gateway

1. Create a new REST API.

2. In the "Resources" section create a new `ANY` method to handle requests to `/` (check "Use Lambda Proxy Integration").

    ![API Gateway index](https://akrylysov.github.io/algnhsa/apigateway-index.png)

3. Add a catch-all `{proxy+}` resource to handle requests to every other path (check "Configure as proxy resource").

    ![API Gateway catch-all](https://akrylysov.github.io/algnhsa/apigateway-catchall.png)
4. Deploy your API gateway and copy gateway url    
    
    
### Configuring slack

1. Create a [new application](https://api.slack.com/apps)
2. Copy signing secret from `Basic information page` (at the bottom of the page)
3. Open Oauth and permissions page of your application and add `channels:history`, `chat:write` and `users:read` to both `Bot Token scopes` and `User token scopes`
4. Go to the `Basic information` page and open `Event subscriptions`
5. Turn on `Enable events` and paste your API Gateway URL

### Environment variables

```
# ENV variable can have development or production. production uses aws lambda sdk to run. development uses standard net/http handler. 
# Default: development
ENV=development

# BIND_ADDR configures bind address for local development
# Default: :12022
BIND_ADDR=:12022

# SLACK_SIGNING_SECRET stores signing secret from slack. You can get one on `Basic information` of your app
# Default: unset
SLACK_SIGNING_SECRET=

# SLACK_BOT_TOKEN can start with xoxp or xoxb. Put xoxp if you want to send message from your personal account. xoxb sends message from the bot username
# Default: unset
SLACK_BOT_TOKEN=xoxp...

# WELCOME_MESSAGE stores welcome message for your onboring experience
# Default: unset
WELCOME_MESSAGE=Hello
```
