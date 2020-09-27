#sudo usermod -aG docker $USER && newgrp docker
curl -Lo minikube https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/releases/v1.13.0/minikube-linux-amd64 && chmod +x minikube && chown -R $USER $HOME/.minikube; chmod -R u+wrx $HOME/.minikube
