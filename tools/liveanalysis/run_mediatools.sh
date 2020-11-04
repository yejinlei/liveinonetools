docker stop mediatools
docker run --name='mediatools' --rm --privileged=true --entrypoint="" yejinlei/mediatools:0.1 ping baidu.com
