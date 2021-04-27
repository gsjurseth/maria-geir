# maria-geir
Sample repo showing Envoy used for several services plus a portal that documents them

## What's in here anyway?
The following have been setup and are run with docker compose
* grpc server written in java running greeter service: normal and pirate mode available
* grpc client written in golang that calls the remote service (and accepts an optional key as an argument)
* a graphql service. This uses all flat files and is simply the book list example. It's self contained
* kafka and zookeepr for said kafka
* a nodejs service that constantly sends messages to a kafka topic called `messages`
* a nodejs service that consumes that topic and then sends those messages to any websocket clients that connect and use a unique groupId
* an envoy service that proxies for all of it.

And a nice image:

![System Architecture](/images/layout.png "System Architecture"){: style="background-color: white;"}

## What's coming
I'm in the process of building out a few more things
* documentation published in the portal for all provided services

The idea here is to create working example where apigee is used in concert with a simple microservice
to translate kafka -> websockets, proxy for gRPC, and proxy for graphql.

## Getting setup
Run the following:

NB:
* The email will be used to create the developer and so needs to be a user with rights. Probably your own email (no password is used)
* The remote url is probably the base url of your Apigee X setup. If you've used the auto-provisioner it will be one that ends in `nip.io`
* You'll need to create a service-account and store the key locally and then refer to that key as outlined below

Run the setup like so:

```bash
./setup.sh -a <path-to-ax-sa-json> -o your-org -e your-environment -t $(gcloud beta auth print-access-token) -r "http://remote-url-for-apigee-x" -u user@email.com
```

With that complete you'll need to add the following `api_header` setting your to your config.yaml file. Simply locate the auth stanza
and add append `api_header: x-target-name` like so:

```yaml
    auth:
      jwt_provider_key: https://1.2.3.4.nip.io/remote-token/token
      append_metadata_headers: true
      api_header: x-target-name
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

## Testing it all
Here I've provide a number of commands I've used to test all of this. One thing I've used is a websocket cli called wscat.
So my examples use it.

Every example I provide here refers to `<apikey>`. You'll need to replace that with the key stored in `my_app.json` which
is created when the setup script finishes.

### Testing the websocket bit:
The point of the websocket is two fold:
1. Show how we might use envoy with the envoy adapter to act as a proxy for websockets
2. Use the websocket server as a mediation point: websocket <-> kafka

The websocket server is itself subscribing to a kafka topic that's being populated by another nodejs-service: a producer.
Provide a `groupId` as shown below or it will default to one ... And that one will only work once of course.

```bash
 wscat -c http://localhost:8080/ws -H x-api-key:<apikey> -H "x-group-id: first"
```

### Testing the graphql bit:
The graphql server is simple. A nodejs service returning a hardcoded list based on a schema for books. You can make a query like so with curl:

```bash
curl -i 'http://localhost:8080/graphql' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' --data-binary '{"query":"{\n  books {\n    title\n    author\n  }\n}"}' -H "x-api-key: <apikey>" -H "host: envoy.local"
```
### Testing the gRPC bit:
We talk a whole lot about proxying for gRPC but what does it look like. This example uses vanilla envoy with the envoy adapter
to proxy a java-springboot gRPC service. I've provided a simple golang-grpc client that you can use to test this service.

Execute that test like so:

```bash
cd go-grpc-client/
GOPATH=$GOPATH:$(pwd) go run main.go -h localhost -p 8080 -k <apikey> gobblybook
```
Or if you'd like a pirate response add the `-pirate` flag like so
```bash
cd go-grpc-client/
GOPATH=$GOPATH:$(pwd) go run main.go -h localhost -p 8080 -pirate -k <apikey> gobblybook
```
