#!/bin/bash
#该脚本仅用于滚动发布 su root ./eshort.sh start|stop|restart
APP_NAME=eshort

start() {
  echo "start...."
  go build
  ./${APP_NAME} &
  echo "start success"
}

stop() {
  echo "stop..."
  pgrep ${APP_NAME} | xargs kill -SIGTERM
  echo "stop success"
}

restart() {
  echo "restart...."
  git pull origin master
  go build
  pgrep ${APP_NAME} | xargs kill -SIGUSR2
  echo "restart success"
}

case $1 in
"stop")
  stop
  ;;
"restart")
  restart
  ;;
"start")
  start
  ;;
*)
  echo "仅支持 start,stop,restart 命令"
  exit 1
  ;;
esac
