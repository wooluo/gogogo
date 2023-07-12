#!/bin/bash
##########################
#                        #
# install yak for tools  # 
#                        #
##########################
#ubuntu 23.04 linux amd64 
#Linux vultr 6.2.0-23-generic #23-Ubuntu SMP PREEMPT_DYNAMIC Wed May 17 16:55:20 UTC 2023 x86_64 x86_64 x86_64 GNU/Linux

#安装afrog工具,并升级
wget https://github.com/zan8in/afrog/releases/download/v2.5.6/afrog_2.5.6_linux_amd64.zip
unzip afrog_2.5.6_linux_amd64.zip
chmod +x afrog
./afrog
./afrog -un
rm afrog_2.5.6_linux_amd64.zip


#安装nuclei工具,并升级
wget https://github.com/projectdiscovery/nuclei/releases/download/v2.9.6/nuclei_2.9.6_linux_amd64.zip
unzip nuclei_2.9.6_linux_amd64.zip
chmod +x nuclei
./nuclei
./nuclei -up
rm nuclei_2.9.6_linux_amd64.zip

bash <(curl -sS -L http://oss.yaklang.io/install-latest-yak.sh)

#添加防火墙端口，并重新加载生效
ufw allow 64333
ufw allow 8085
ufw allow 8087
ufw reload
#安装iftop，查看发包工具
apt install iftop

#启动全局反连
nohup yak bridge --secret cisco888 2>1 &

#启动远程模式
nohup yak grpc --host 0.0.0.0 --port 8087 --secret youR-aw0some-PA5sss --tls 2>1 &

#查看运行情况
ps -ef |grep yak

#查看连接证书
yak grpc --host 0.0.0.0 --port 8087 --secret youR-aw0some-PA5sss --tls

