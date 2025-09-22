```bash
# root和user zsh 安装
# 配置区分ssh和local
# 远端.zshrc添加
# SSH 登录时的前缀，cyan色
PROMPT='%F{cyan}[SSH:%n@%m]%f %~ %# '

# 免密码su切换账户
sudo vim /etc/pam.d/su
# 若无，添加如下
# Debian没有wheel组，通常用sudo组。
# 下面这行的意思是：**sudo组里的用户 su root 时免密**。
auth sufficient pam_wheel.so trust use_uid group=sudo   
# 查看用户
cat /etc/passwd
# 用户加入sudo组
sudo usermod -aG sudo {user}
```

```bash
docker pull --platform=linux/amd64 ubuntu:20.04 
docker save -o ubuntu_20.04.tar ubuntu:20.04 
scp ubuntu_20.04.tar devbox:/your/target/path/ 
docker load -i ubuntu_20.04.tar 
docker images 
docker run --network fl-demo-net -it --name agw ubuntu:20.04 /bin/bash
```
