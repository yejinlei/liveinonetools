pyenv local miniconda3-4.3.30
pyenv virtualenv miniconda3-4.3.30 mxnet
source activate mxnet
#pip install --upgrade pip -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
pip install -r mxnet_requestments.txt -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
