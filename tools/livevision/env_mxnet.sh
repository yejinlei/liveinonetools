pyenv virtualenv 3.6.0 mxnet
source activate mxnet
pip install --upgrade pip
pip install -r mxnet_requestments.txt -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
