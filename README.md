#Cowbell

### What is this?
Cowbell is an experiment to add a Rancher service scale up webhook trigger. Cowbell is launched inside a Rancher stack and is configured to listen for events, and react accordingly.

### How to use Cowbell.

Cowbell is meant to be deployed as a service inside your application stack. To deploy cowbell, add a service definition to your docker-compose.yml file:

```
  ...
  web:
    image: nginx
  ...
  cowbell:
    image: cloudnautique/cowbell:v0.0.1
    ports:
      - 8888:8088 //pick a public port that works for your setup.
  ...
```

Then configure Cowbell through rancher-metadata in the rancher-compose.yml file.

```
...
cowbell:
  metadata:
    accessKey: "KEY" # Rancher API Access Key
    secretKey: "SECRET" # Rancher API Secret Key
    rancherURL: "http://myrancher.com/v1/projects/1a5" # Rancher URL from API
    services:
      - name: "web" # name of the service to scale
        increment: 1 # number of containers to add per event
        decrement: 1 # does nothing yet...
        token: "set this to something long and url friendly"
        quietTime: 60 # does nothing yet
...
```

Once the service is configured, you should now be able to send a POST and have the service scale up by the increment:

`curl -X POST http://(host):(publicport)/v1-scale/services/web?token=reallylongtokenthatshouldbekeptsecret``

### Todo

* Add quiet period to prevent overloading
* Add Decrement??

### Building

To build cowbell, just run `./scripts/build`