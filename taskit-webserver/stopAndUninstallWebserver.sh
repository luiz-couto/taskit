#!/bin/bash
echo "Removing last Docker Container...."
sudo docker stop taskit-webserver
sudo docker rm taskit-webserver
sudo docker volume rm taskit