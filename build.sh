#!/bin/bash

if [[ $# <= 1 ]]
do
 echo "please use "$0" swarm-node-number"
 exit 1
done



NODES=$1

for (( c=1; c<=NODES; c++ ))
do
   echo "Welcome $c times"
done
 

exit 1

echo "creating dev virtualbox"
sudo docker-machine create -d virtualbox dev

echo "loading env:";docker-machine env local
eval "$(docker-machine env local)"

echo "creating token"
export TOK=`docker run swarm create`
echo -n "- token:"$TOK

echo "Create the Swarm master:"
docker-machine create \
    -d virtualbox \
    --swarm \
    --swarm-master \
    --swarm-discovery token://$TOK \
    swarm-master

# maybe can be parallelized
echo "Creating Swarm nodes"
INAME="swarm-node-00"
echo "creating node 00"
sudo docker-machine create \
    -d virtualbox \
    --swarm \
    --swarm-discovery token://$TOK \
    $INAME
echo "- created: "$INAME

echo "Connecting nodes to master..."
eval "$(docker-machine env --swarm swarm-master)"
docker info

echo "Composing ..."
docker-compose build

echo "Compose up"
docker-compose up -d

echo "Compose ps"
docker-compose ps

echo "test with curl"
curl "http://$(docker-machine ip):80/test"


# docker pull liuggio/golang-stupid-microservice
#$(docker-machine env swarm-master)