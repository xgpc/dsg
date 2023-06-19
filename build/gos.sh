# /bin/bash

# 获取参数
method=$1
name=$2

echo "$method"

# 创建Log文件夹
if [ ! -d "./logs" ]; then
  mkdir ./logs
fi


# 启动
if [ $method == start ]
then
    # nohup ./wx-account.sh >> logs/wx-account.log 2>&1 & echo $! > wx-account.pid
    chmod a+x "$name".sh
    nohup ./"$name".sh >> logs/"$name".log 2>&1 & echo $! > "$name".PID  
fi

# 关闭
if [ $method == stop ]
then
    kill -9 `cat "$name".PID`
fi


# restart
if [ $method == restart ]
then
    kill -9 `cat "$name".PID`
    chmod a+x "$name".sh
    nohup ./"$name".sh >> logs/"$name".log 2>&1 & echo $! > "$name".PID
fi

