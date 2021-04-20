# maria-geir
Sample repo showin Envoy used for several services plus a portal that documents them

## What's in here anyway?
The following have been setup and are run with docker compose
* grpc server written in java running greeter service: normal and pirate mode available
* grpc client written in golang that calls the remote service (and accepts an optional key as an argument)
* a graphql service. This uses all flat files and is simply the book list example. It's self contained
* an envoy service that proxies for all of it.

## What's coming
I'm in the process of building out a few more things
* an async api documentation example to be published in the portal
* a kafka image that will auto-create an example topic
* a producer for said topic that spits out messages once per second
* a consumer that subscribes to said topic and then pushes those messages down a websocket once they connect
* a websocket client that can receive these messages
* documentation published in the portal for that same websocket service

The idea here is to create working example where apigee is used in concert with a simple microservice
to translate kafka -> websockets for easier consumption by external parties.

## Getting setup
Run the following:

NB:
* The email will be used to create the developer and so needs to be a user with rights. Probably your own email (no password is used)
* The remote url is probably the base url of your Apigee X setup. If you've used the auto-provisioner it will be one that ends in `nip.io`

```bash
./setup.sh -o your-org -e your-environment -t $(gcloud beta auth print-access-token) -r "http://remote-url-for-apigee-x" -u user@email.com
```


## Building and Runnig it once configured
To get this running you'll need to build it all (after you've run setup mind you) by doing:

```bash
docker-compose build
```

Once it's built you can play with any of the config files and environment variables named
in `docker-compose.yaml`

Finally you can run it by:

```bash
docker-compose up
```
