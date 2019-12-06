pyenv local 3.6.0
pyenv virtualenv 3.6.0 imageai
source activate imageai
pip install --upgrade pip -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
pip install -r imageai_requestments_gpu.txt -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
