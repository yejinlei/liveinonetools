#snap install minikube
#snap install kubectl
#curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube && mv minikube /usr/local/bin/
axel -n 8 https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 -o ./minikube && chmod +x minikube && mv minikube /usr/local/bin/
