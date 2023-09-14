# aws-global-accelerator-custom-routing-workshop

This project is mainly used in AWS Global Accelerator Custom Routing Workshop. The project includes the following echo-server and echo-cli

## echo-server

To monitor different protocols on multiple ports, the relevant configuration files are as follows:

Protocol | Default value of Port number | open by default
-|-|-
TCP|8080|true
UDP|8081|true
HTTP|8082|true
Websocket|8083|true
GRPC|8084|true

The default parameters can be modified by modifying the relevant configuration items in the config.yaml file.

### echo-cli

By default, access will be based on the default port number of the echo server. You can modify the access port and the number of tests through the configuration file.
