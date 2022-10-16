#1.执行命令拉取并加载httpserver镜像
docker run -it  -p 8080:80 honework3/httpserver

#2.获取httpserver的PID
  #获取container id
  docker ps 
  #根据container id 获取pid
  docker inspect container_id | grep -i pid

#3.根据pid进入容器执行命令
   nsenter -t pid -n ip a