#!/bin/sh

#本脚本只对master进行元数据恢复
hostname=`hostname`
if [[ ! $hostname = *"master"* ]];then
   exit 0
fi
#检测集群是否正常，如果正常，从正常的节点恢复cluser-data，拉起服务

nn1_state=`/home/hadoop/hadoop-current/bin/hdfs haadmin -getServiceState nn1 2>/dev/null`
nn2_state=`/home/hadoop/hadoop-current/bin/hdfs haadmin -getServiceState nn2 2>/dev/null`

nn1=`/home/hadoop/hadoop-current/bin/hdfs getconf -namenodes | awk '{print $1}'`
nn2=`/home/hadoop/hadoop-current/bin/hdfs getconf -namenodes | awk '{print $2}'`
server_id=`cat /home/hadoop/zookeeper-current/conf/zoo.cfg | grep $hostname | awk -F '=' '{print $1}' | awk -F '.' '{print $2}'`

# 如果有一个nn为active状态，表明集群状态正常
if [ "$nn1_state" == 'active' ]; then
   if [ ! -d "/home/hadoop/cluster-data/pids" ]; then
	scp -r $nn1:~/cluster-data/* /home/hadoop/cluster-data/
	echo $server_id > /home/hadoop/cluster-data/zookeeper/myid
   fi
fi

if [ "$nn2_state" == 'active' ]; then
   if [ ! -d "/home/hadoop/cluster-data/pids" ]; then
	scp -r $nn2:~/cluster-data/* /home/hadoop/cluster-data/
	echo $server_id > /home/hadoop/cluster-data/zookeeper/myid
   fi
fi
