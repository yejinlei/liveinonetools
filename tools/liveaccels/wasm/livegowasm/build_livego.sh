git clone https://github.com/gwuhaolin/livego livego_rep
cd livego_rep
GOOS=js GOARCH=wasm go build -o test.wasm main.go
cp test.wasm ../
cd ..
echo '{' > .livego.json
echo '"server": [' >> .livego.json
echo    '{' >>.livego.json
echo '     "appname": "live",' >> .livego.json
echo '      "liveon": "on",' >> .livego.json
echo '      "hlson": "on"' >> .livego.json
echo '    }' >> .livego.json
echo '  ]' >> .livego.json
echo '}' >> .livego.json
cp $GOROOT/misc/wasm/wasm_exec.{html,js} .
