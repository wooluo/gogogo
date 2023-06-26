wget https://github.com/zan8in/afrog/releases/download/v2.5.6/afrog_2.5.6_linux_amd64.zip
unzip afrog_2.5.6_linux_amd64.zip
chmod +x afrog
rm afrog_2.5.6_linux_amd64.zip

wget https://github.com/projectdiscovery/nuclei/releases/download/v2.9.6/nuclei_2.9.6_linux_amd64.zip

unzip nuclei_2.9.6_linux_amd64.zip
chmod +x nuclei
rm nuclei_2.9.6_linux_amd64.zip

bash <(curl -sS -L http://oss.yaklang.io/install-latest-yak.sh)

ufw allow 64333
ufw allow 8085
ufw allow 8087
ufw reload

nohup yak bridge --secret cisco888 2>1 &


nohup yak grpc --host 0.0.0.0 --port 8087 --secret youR-aw0some-PA5sss --tls 2>1 &
