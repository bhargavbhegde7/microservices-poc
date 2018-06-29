#stop all the containers
docker stop $(docker ps -a -q)
#remove all the containers.
docker rm $(docker ps -a -q)
#remove all the images
docker rmi $(docker images -a -q)
#show all the containers
docker ps -a
#show all the images
docker images -a
