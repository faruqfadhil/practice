# PUBSUB EXAMPLE

## RUN WITH EMULATOR
- run emulator
```shell
gcloud beta emulators pubsub start --project=<PUBSUB_PROJECT_ID>
```
- you need to set the environment variables each time you start the emulator.
```shell
# run this in your terminal thats your apps will be run.
$(gcloud beta emulators pubsub env-init)
```

for more details: https://cloud.google.com/pubsub/docs/emulator

## ADMIN
To use the emulator, you must have an application built using the Google Cloud Client Libraries. The emulator does not support Cloud Console or gcloud pubsub commands.

prerequisite:
- run emulator
- run emulator environment variable

- go to /pubsub/admin to create new topic and new subscription
```shell
#create topic
go run main.go -action=create-topic -topic=my-topic -project=test-project

#create subsription
go run main.go -action=create-subscription -topic=my-topic -project=test-project -subscription=my-subs
```

## PUBLISH MESSAGE
prerequisite:
- run emulator
- run emulator environment variable
- created project
- created topic
- created subscription

- go to /pubsub/publisher to publish message
```shell
go run main.go -project=<project ID> -topic=<created topic>
```

## SUBSCRIBE & RECEIVE MESSAGE
prerequisite:
- run emulator
- run emulator environment variable
- created project
- created topic
- created subscription

- go to /pubsub/subscriber to subscribe & receive the message
```shell
go run main.go -project=<project ID> -subscription=<created subscription>
```
