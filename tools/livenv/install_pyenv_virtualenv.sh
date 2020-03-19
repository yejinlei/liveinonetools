#apt-get install curl git-core
#apt-get install libbz2
#apt-get install libbz2-dev
#apt-get install libssl-dev
#apt-get install libreadline6 libreadline6-dev
#apt-get install libsqlite3-dev

#curl -L https://raw.github.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash
#or
#git clone git://github.com/yyuu/pyenv.git ～/.pyenv
#echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bashrc
#echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bashrc
#echo 'eval "$(pyenv init -)"' >> ~/.bashrc
#source ~/.bashrc

#git clone https://github.com/yyuu/pyenv-virtualenv.git ~/.pyenv/plugins/pyenv-virtualenv
#exec "$SHELL"

pyenv install 2.7.13 -v
pyenv virtualenv 2.7.13 env_2.7.13
pyenv local env_2.7.13

#pyenv deactivate
#pyenv uninstall env_2.7.13
#rm -rf ~/.pyenv/versions/env_2.7.13/

