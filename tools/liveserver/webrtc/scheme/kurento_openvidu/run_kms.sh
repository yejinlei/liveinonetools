docker run -d -p 4443:4443 -p 8888:8888 --rm \
    -e DOMAIN_OR_PUBLIC_IP=192.168.1.14 \
    -e KMS_STUN_IP="121.37.143.208" \
    -e KMS_STUN_PORT="3478" \
    -e OPENVIDU_SECRET=pass \
    -e CERTIFICATE_TYPE=selfsigned \
    -e KMS_TURN_URL=usr:pass@121.37.143.208:3478 \
    -v /var/run/docker.sock:/var/run/docker.sock \
yejinlei/kurento-media-server
