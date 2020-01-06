git clone https://github.com/gwuhaolin/livego livego_rep
cd livego_rep
go build
cp livego ../
cd ..
rm -rf livego_rep
echo '{' > .livego.json
echo '"server": [' >> .livego.json
echo    '{' >>.livego.json
echo '     "appname": "live",' >> .livego.json
echo '      "liveon": "on",' >> .livego.json
echo '      "hlson": "on"' >> .livego.json
echo '    }' >> .livego.json
echo '  ]' >> .livego.json
echo '}' >> .livego.json
