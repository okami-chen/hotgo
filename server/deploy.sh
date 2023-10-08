#! /bin/bash

# 应用程序名称
APP_NAME=server-linux-amd64

function check_if_process_is_running {
    PID=$(ps aux | grep server-linux-amd64 | grep -v grep | awk '{print $2}' )
    if [ -n "$PID" ]; then
        return 0
    else
        return 1
    fi
}

case "$1" in
    start)
        if check_if_process_is_running; then
            echo "$APP_NAME is already running"
            exit 1
        fi
        ./server-linux-amd64 > access.log 2>&1 &
        echo "Starting $APP_NAME..."
        tail -200f access.log
        ;;
    stop)
        if ! check_if_process_is_running; then
            echo "$APP_NAME is not running"
            exit 1
        fi
        kill $(ps aux | grep server-linux-amd64 | grep -v grep | awk '{print $2}' )
        echo "Stopping $APP_NAME..."
        sleep 2s
        ;;
    restart)
        $0 stop
        sleep 1
        $0 start
        ;;
    status)
        if check_if_process_is_running; then
            echo "$APP_NAME is running"
        else
            echo "$APP_NAME is not running"
        fi
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|status}"
        exit 1
        ;;
esac

exit 0
