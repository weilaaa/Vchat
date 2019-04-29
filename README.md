# Vchat
Vchat based on TCP/IP as a multiple-users chatting software containing basic functions and can be extended very easily. 
Vserver and Vclient are included.

# Server based on Aliyun
Vchat server is deployed on Aliyun, anyone can use specified client of relevant OS system to access to Vchat server.

# Something about docker
I just think about if I can make the whole server environment into a docker mirror, I just need pack all of dependence and push it up to dockerhub, and deploy fast in Aliyun server. I am looking forward to doing it.

# Dependence 
This whole project based on go1.11 and redis.5.0.4. You need make sure the versions of them are newer then mine or at least keep same.

# Usage
No.1

1.git clone https://github.com/WeilaScofield/Vchat.git

2.go build vServer and vClient

3.make redis running

4.open vServer starting to listen

5.open vClient and have fun

No.2

1.get the client of winClient, macClient or linuxClient

2.enjoy chatting free

# Extension
Vchat developed follows MVC model so that it owns great extensibility. Many interfaces have been exposed for new functions extension later.

# Reference
Some details refer to the project of Han ShunPing.

# Conversation
If you guys have any interests to discuss or figure out any problems, you can cantact me.

Wechat:cai4561168630

e-mail:weilanidaye@gmail.com
