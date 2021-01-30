FROM centos:7

RUN yum install -y kde-l10n-Chinese glibc-common epel-release iproute telnet nmap-ncat bind-utils  net-tools bash-completion vim wget curl tree sudo \
    && yum install -y supervisor openssh-clients openssh-server cronie rsyslog \
    && yum clean all \
    && localedef -c -f UTF-8 -i zh_CN zh_CN.utf8 && echo LANG="zh_CN.UTF-8" >> /etc/profile \
    && echo -e "set encoding=utf-8\nset termencoding=utf-8\nset fileencodings=utf-8,gbk,latin1" >> /etc/vimrc \
    && ssh-keygen -t rsa -P "" -f /etc/ssh/ssh_host_key \
    && ssh-keygen -t rsa -P "" -f /etc/ssh/ssh_host_rsa_key \
    && ssh-keygen -t dsa -P "" -f /etc/ssh/ssh_host_dsa_key \
    && ssh-keygen -t ecdsa -P "" -f /etc/ssh/ssh_host_ecdsa_key \
    && ssh-keygen -t ed25519 -P "" -f /etc/ssh/ssh_host_ed25519_key \
    && cd /opt/ \
    && wget https://repo.huaweicloud.com/java/jdk/8u202-b08/jdk-8u202-linux-x64.tar.gz \
    && tar xf jdk-8u202-linux-x64.tar.gz && rm -f jdk-8u202-linux-x64.tar.gz \
    && ln -sv jdk1.8.0_202/ java \
    && echo -e """export JAVA_HOME=/opt/java\nexport PATH=\$JAVA_HOME/bin:\$PATH""" > /etc/profile.d/java.sh 

ARG ZK_VERSION=3.5.9
ARG HADOOP_VERSION=3.2.1
ARG HBASE_VERSION=2.2.6

RUN groupadd -g 501 hadoop && useradd -u 501 -g 501 hadoop  \
    && echo "hadoop ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers \
    && mkdir -p /home/hadoop/.ssh/ \
    && ssh-keygen -t rsa -C "hadoop@hbase" -P "" -f /home/hadoop/.ssh/id_rsa \
    && cat /home/hadoop/.ssh/id_rsa.pub >> /home/hadoop/.ssh/authorized_keys && chmod 600 /home/hadoop/.ssh/authorized_keys \
    && cd /home/hadoop \
    && wget https://mirrors.tuna.tsinghua.edu.cn/apache/zookeeper/zookeeper-${ZK_VERSION}/apache-zookeeper-${ZK_VERSION}-bin.tar.gz \
    && wget https://mirrors.tuna.tsinghua.edu.cn/apache/hadoop/common/hadoop-${HADOOP_VERSION}/hadoop-${HADOOP_VERSION}.tar.gz \
    && wget https://mirrors.tuna.tsinghua.edu.cn/apache/hbase/${HBASE_VERSION}/hbase-${HBASE_VERSION}-bin.tar.gz \
    && tar xf apache-zookeeper-${ZK_VERSION}-bin.tar.gz && rm -f apache-zookeeper-${ZK_VERSION}-bin.tar.gz \
    && tar xf hadoop-${HADOOP_VERSION}.tar.gz && rm -f hadoop-${HADOOP_VERSION}.tar.gz \
    && tar xf hbase-${HBASE_VERSION}-bin.tar.gz && rm -f hbase-${HBASE_VERSION}-bin.tar.gz \
    && ln -sv apache-zookeeper-${ZK_VERSION}-bin zookeeper-current \
    && ln -sv hadoop-${HADOOP_VERSION} hadoop-current \
    && ln -sv hbase-${HBASE_VERSION} hbase-current \
    && echo "export HADOOP_HOME=/home/hadoop/hadoop-current" >> /home/hadoop/.bashrc \
    && echo "export HBASE_HOME=/home/hadoop/hbase-current" >> /home/hadoop/.bashrc \
    && echo "export ZK_HOME=/home/hadoop/zookeeper-current" >> /home/hadoop/.bashrc \
    && echo "export PATH=\$HADOOP_HOME/bin:\$HADOOP_HOME/sbin:\$HBASE_HOME/bin:\$ZK_HOME/bin:\$PATH" >> /home/hadoop/.bashrc \
    && chown -R hadoop:hadoop /home/hadoop

COPY conf/* /etc/supervisord.d/
COPY scripts/* /opt/

COPY template/hadoop/* /home/hadoop/hadoop-current/etc/hadoop/
COPY template/hbase/* /home/hadoop/hbase-current/conf/
COPY template/zk/* /home/hadoop/zookeeper-current/conf/

WORKDIR /home/hadoop

ENTRYPOINT ["/opt/entrypoint.sh"]
CMD ["/usr/bin/supervisord","-n","-c","/etc/supervisord.conf"]
