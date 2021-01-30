#!/bin/bash

MASTER1=`echo $MASTER_HOSTS | awk -F"," '{print $1}' `
MASTER2=`echo $MASTER_HOSTS | awk -F"," '{print $2}' `
MASTER3=`echo $MASTER_HOSTS | awk -F"," '{print $3}' `
MY_INDEX=`echo "${MY_HOST:0-1}+1"|bc`

role=`echo $MY_HOST|grep -iq master && echo master || echo slave`

generateZkConfig() {

    echo "replace zk conf"
    sed -i -e "s@hadoop1@${MASTER1}@g" -e "s@hadoop2@${MASTER2}@g" -e "s@hadoop3@${MASTER3}@g" `grep -E 'hadoop[1|2|3]' -rl  /home/hadoop/zookeeper-current/conf/`
    echo -e "export ZOO_LOG_DIR=/home/hadoop/logs/zookeeper" >> /home/hadoop/zookeeper-current/bin/zkEnv.sh
    if [ ! -d /home/hadoop/cluster-data/zookeeper ]
    then
        mkdir -p /home/hadoop/cluster-data/zookeeper
        chown -R hadoop:hadoop /home/hadoop/cluster-data/zookeeper
    fi
    echo "$MY_INDEX" > /home/hadoop/cluster-data/zookeeper/myid
}

generateMasterConfig() {

    echo "replace hadoop conf"
    sed -i -e "s@hadoop1@${MASTER1}@g" -e "s@hadoop2@${MASTER2}@g" -e "s@hadoop3@${MASTER3}@g" `grep -E 'hadoop[1|2|3]' -rl  /home/hadoop/hadoop-current/etc/hadoop/`
    echo -e 'export HADOOP_LOG_DIR=/home/hadoop/logs/hadoop' >> /home/hadoop/hadoop-current/etc/hadoop/hadoop-env.sh



    echo "replacehbase conf"
    sed -i -e "s@hadoop1@${MASTER1}@g" -e "s@hadoop2@${MASTER2}@g" -e "s@hadoop3@${MASTER3}@g" `grep -E 'hadoop[1|2|3]' -rl  /home/hadoop/hbase-current/conf/`
    echo -e "export JAVA_HOME=/opt/java/\nexport HBASE_OPTS=\"\$HBASE_OPTS -XX:+UseConcMarkSweepGC\"\nexport HBASE_LOG_DIR=/home/hadoop/logs/hbase\nexport HBASE_MANAGES_ZK=false" >> /home/hadoop/hbase-current/conf/hbase-env.sh
}

generateSlaveConfig() {
    dirs=${DN_DIRS%?}
    sed -i -e "s@DNDIR@${dirs}@g" /home/hadoop/hadoop-current/etc/hadoop/hdfs-site.xml
    echo "${MY_HOST}.hbase-slave-service" > /home/hadoop/hadoop-current/etc/hadoop/workers

}

generateSSHDConfig() {
cat > /etc/supervisord.d/sshd.ini << EOF
[program:sshd]
command=/usr/sbin/sshd -D -p $1
autostart=true       ; 在supervisord启动的时候也自动启动
startsecs=10         ; 启动10秒后没有异常退出，就表示进程正常启动了，默认为1秒
autorestart=true     ; 程序退出后自动重启,可选值：[unexpected,true,false]，默认为unexpected，表示进程意外杀死后才重启
startretries=3       ; 启动失败自动重试次数，默认是3
user=root            ; 用哪个用户启动进程，默认是root
priority=1           ; 进程启动优先级，默认999，值小的优先启动 redirect_stderr=true ; 把stderr重定向到stdout，默认false
stdout_logfile_maxbytes=20MB  ; stdout 日志文件大小，默认50MB
stdout_logfile_backups = 20   ; stdout 日志文件备份数，默认是10
; stdout 日志文件，需要注意当指定目录不存在时无法正常启动，所以需要手动创建目录（supervisord 会自动创建日志文件）
stdout_logfile=/tmp/sshd.log
stopasgroup=false     ;默认为false,进程被杀死时，是否向这个进程组发送stop信号，包括子进程
killasgroup=false     ;默认为false，向进程组发送kill信号，包括子进程
EOF

cat > /home/hadoop/.ssh/config << EOF
Host *
  StrictHostKeyChecking no
Host *.hbase-master-service
  Port 2022

Host *.hbase-slave-service
  Port 2222
EOF

}

generateMasterScipt() {
cat > /opt/service.sh << EOF
#!/bin/bash
set -e

EOF
echo $COMPONENT|grep -q hdfs || (echo "hdfs is requested" && exit 1)
echo "/home/hadoop/zookeeper-current/bin/zkServer.sh stop" >> /opt/service.sh
echo "/home/hadoop/zookeeper-current/bin/zkServer.sh start" >> /opt/service.sh
echo "/home/hadoop/hadoop-current/bin/hdfs --daemon stop journalnode" >> /opt/service.sh
echo "/home/hadoop/hadoop-current/bin/hdfs --daemon start journalnode" >> /opt/service.sh
case "$MY_INDEX" in
    1)
        cat >> /opt/service.sh << EOF
[ -d /home/hadoop/cluster-data/dfs/name/current ] && echo "namenode has been formated" || /home/hadoop/hadoop-current/bin/hadoop namenode -format
flag="true"
while \$flag;do
    /home/hadoop/zookeeper-current/bin/zkCli.sh  ls /hadoop-ha/mycluster && flag="false" || /home/hadoop/hadoop-current/bin/hdfs zkfc -formatZK
    sleep 10
done

EOF
        echo "/home/hadoop/hadoop-current/bin/hdfs --daemon stop namenode" >> /opt/service.sh
        echo "/home/hadoop/hadoop-current/bin/hdfs --daemon start namenode" >> /opt/service.sh
        echo "/home/hadoop/hadoop-current/bin/hdfs --daemon stop zkfc" >> /opt/service.sh
        echo "/home/hadoop/hadoop-current/bin/hdfs --daemon start zkfc" >> /opt/service.sh
        if `echo $COMPONENT|grep -q mapreduce`
        cat >> /opt/service.sh << EOF
    flag="true"
    while \$flag;do
        /home/hadoop/hadoop-current/bin/hdfs dfs -ls / && flag="false" || sleep 10
    done
EOF
        then
            echo "/home/hadoop/hadoop-current/bin/yarn --daemon stop resourcemanager" >> /opt/service.sh
            echo "/home/hadoop/hadoop-current/bin/yarn --daemon start resourcemanager" >> /opt/service.sh
        fi
        ;;
    2)
        cat >> /opt/service.sh << EOF
[ -d /home/hadoop/cluster-data/dfs/name/current ] && echo "namenode has been synced" || {
    mkdir -p /home/hadoop/cluster-data/dfs/name/
    flag="true"
    while \$flag;do
        ssh hbase-master-0.hbase-master-service ls /home/hadoop/cluster-data/dfs/name/ && scp -r hbase-master-0.hbase-master-service:/home/hadoop/cluster-data/dfs/name/* /home/hadoop/cluster-data/dfs/name/ && flag="false" || sleep 10
    done
}

EOF
        echo "/home/hadoop/hadoop-current/bin/hdfs --daemon stop namenode" >> /opt/service.sh
        echo "/home/hadoop/hadoop-current/bin/hdfs --daemon start namenode" >> /opt/service.sh
        echo "/home/hadoop/hadoop-current/bin/hdfs --daemon stop zkfc" >> /opt/service.sh
        echo "/home/hadoop/hadoop-current/bin/hdfs --daemon start zkfc" >> /opt/service.sh
        cat >> /opt/service.sh << EOF
    flag="true"
    while \$flag;do
        /home/hadoop/hadoop-current/bin/hdfs dfs -ls / && flag="false" || sleep 10
    done
EOF
        if `echo $COMPONENT|grep -q mapreduce`
        then
            echo "/home/hadoop/hadoop-current/bin/yarn --daemon stop resourcemanager" >> /opt/service.sh
            echo "/home/hadoop/hadoop-current/bin/yarn --daemon start resourcemanager" >> /opt/service.sh
        fi
        if `echo $COMPONENT|grep -q hbase`
        then
            echo "/home/hadoop/hbase-current/bin/hbase-daemon.sh stop master --backup" >> /opt/service.sh
            echo "/home/hadoop/hbase-current/bin/hbase-daemon.sh start master --backup" >> /opt/service.sh
        fi
        ;;
    3)
        cat >> /opt/service.sh << EOF
    flag="true"
    while \$flag;do
        /home/hadoop/hadoop-current/bin/hdfs dfs -ls / && flag="false" || sleep 10
    done
EOF
        if `echo $COMPONENT|grep -q mapreduce` 
        then
            echo "/home/hadoop/hadoop-current/bin/mapred --daemon stop historyserver" >> /opt/service.sh
            echo "/home/hadoop/hadoop-current/bin/mapred --daemon start historyserver" >> /opt/service.sh
        fi
        if `echo $COMPONENT|grep -q hbase` 
        then
            echo "/home/hadoop/hbase-current/bin/hbase-daemon.sh stop master" >> /opt/service.sh
            echo "/home/hadoop/hbase-current/bin/hbase-daemon.sh start master" >> /opt/service.sh
        fi
        ;;
esac
}

generateSlaveScipt() {
    cat > /opt/service.sh << EOF
#!/bin/bash
set -e

EOF
echo $COMPONENT|grep -q hdfs || (echo "hdfs is requested" && exit 1)
    echo "/home/hadoop/hadoop-current/bin/hdfs --daemon stop datanode" >> /opt/service.sh
    echo "/home/hadoop/hadoop-current/bin/hdfs --daemon start datanode" >> /opt/service.sh
if `echo $COMPONENT|grep -q mapreduce`
then
    echo "/home/hadoop/hadoop-current/bin/yarn --daemon stop nodemanager" >> /opt/service.sh
    echo "/home/hadoop/hadoop-current/bin/yarn --daemon start nodemanager" >> /opt/service.sh
fi
if `echo $COMPONENT|grep -q hbase`
then
    echo "/home/hadoop/hbase-current/bin/hbase-daemon.sh stop regionserver" >> /opt/service.sh
    echo "/home/hadoop/hbase-current/bin/hbase-daemon.sh start regionserver" >> /opt/service.sh
fi
}

case "$role" in
    master)
        generateZkConfig
        generateMasterConfig
        generateSSHDConfig 2022
        generateMasterScipt
        ;;
    slave)
        generateMasterConfig
        generateSlaveConfig
        generateSSHDConfig 2222
        generateSlaveScipt
        ;;
esac

chmod +x /opt/service.sh
chown -R hadoop:hadoop /home/hadoop

