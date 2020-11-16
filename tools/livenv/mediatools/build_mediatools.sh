docker stop mediatools
docker rm mediatools
docker build -t yejinlei/mediatools:0.1 -f Dockerfile_mediatools .
docker run --name mediatools --privileged=true --rm yejinlei/mediatools:0.1
docker push yejinlei/mediatools:0.1
