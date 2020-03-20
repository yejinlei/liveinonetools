#apt-get install curl git-core
#apt-get install libbz2
#apt-get install libbz2-dev
#apt-get install libssl-dev
#apt-get install libreadline6 libreadline6-dev
#apt-get install libsqlite3-dev

#curl -L https://raw.github.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash
#or
#git clone git://github.com/yyuu/pyenv.git ~/.pyenv
#echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bashrc
#echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.zshrc
#echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bashrc
#echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.zshrc
#echo 'eval "$(pyenv init -)"' >> ~/.bashrc
#echo 'eval "$(pyenv init -)"' >> ~/.zshrc
#exec $SHELL -l

#git clone https://github.com/yyuu/pyenv-virtualenv.git ~/.pyenv/plugins/pyenv-virtualenv
#exec $SHELL -l

#echo 'eval "$(pipenv --completion)"'  >> ~/.bashrc
#echo 'eval "$(pipenv --completion)"'  >> ~/.zshrc
#exec $SHELL -l

#cd ~/.pyenv/plugins/python-build/../.. && git pull && cd -

v=$1
#wget http://mirrors.sohu.com/python/$v/Python-$v.tar.xz -P ~/.pyenv/cache/
axel -n 8 -o ~/.pyenv/cache/ http://mirrors.sohu.com/python/$v/Python-$v.tar.xz 
pyenv install $v
#pyenv install $1 -v

pyenv rehash
pyenv virtualenv $v dev_$v
pyenv activate dev_$v
pyenv rehash
pip install pip --upgrade

#pyenv deactivate
#pyenv uninstall env_$1
#rm -rf ~/.pyenv/versions/env_$1/
