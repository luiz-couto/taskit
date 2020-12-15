#!/bin/bash
echo "Removing last Docker Container...."
sudo docker stop taskit-webserver
sudo docker rm taskit-webserver

echo "Starting Webserver Docker Container...."
sudo docker build -t taskit-webserver .
sudo docker volume create taskit
sudo docker run -v taskit:/data -p 49160:8080 -d --restart always --name taskit-webserver taskit-webserver