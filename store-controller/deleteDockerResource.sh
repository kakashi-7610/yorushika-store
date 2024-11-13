#!/bin/sh

read -p "delete docker resource? (y/n):" CHECK_YN

case "$CHECK_YN" in
  [yY])
    docker container prune -f
    docker network prune -f
    docker volume prune -f
    echo "delete complate!" 
    ;;
  [nN])
    echo "delete cancel"
    ;;
esac
