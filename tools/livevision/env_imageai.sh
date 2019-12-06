pyenv virtualenv miniconda3-4.3.30 imageai
source activate imageai
pip install --upgrade pip
pip install -r imageai_requestments.txt -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
