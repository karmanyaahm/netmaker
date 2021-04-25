#! /bin/bash

sudo docker kill mongodb
sudo docker rm mongodb
sudo docker volume rm mongovol
sudo docker kill netmaker-ui
sudo docker rm netmaker-ui
sudo rm -rf /etc/netmaker
sudo rm /etc/systemd/system/netmaker.service
sudo netclient -c remove -n default
sudo find /etc/systemd/system/ -name 'netclient*' -delete
sudo rm -rf /etc/netclient
sudo systemctl daemon-reload
