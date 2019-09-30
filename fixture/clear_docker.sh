sudo docker rm -f $(sudo docker ps -aq) # 清除容器们
sudo docker network prune # 来清理没有再被任何容器引用的networks
sudo docker volume prune  # 清理挂载卷