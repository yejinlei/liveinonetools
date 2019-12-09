# liveinonetools - tools - capture
1，支持通过kafka和http post触发对rtmp流进行单幅图片快照
2，支持通过kafka将执行结果消息回送
3，支持图片信息上送到fdfs（分布式文件系统）

## 说明
### 调试相关vendor
由于go的库冲突等问题，因此在调试的时候，将工程中vendor、mod相应先移除，然后按照相应代码依赖进行安装对应库即可。

## 配置文件
- my.ini 整个工具的主配置文件
- fdfs_client.conf fdfs相应库需要的配置文件

### my.ini
>[common]
>rtmp_server_addr=10.10.82.66:11936
>image_save_path=D:\my_work\ffmpeg-4.2-win64-static\bin\ 
>
>[ffmpeg]
>bin_full_path=D:\my_work\ffmpeg-4.2-win64-static\bin\ffmpeg
>
>[fdfs]
>server_addr=10.10.82.66:8888
>conf_full_path=D:\my_work\hxdi_code\work_for_git\hxdi\liveinonetools\tools\capture\fdfs_client.conf
>
>[kafka]
>server_addr=10.10.82.66:9092

- rtmp_server_addr, rtmp服务器的地址，支持ip:port或ip
- image_save_path，图片临时存储的绝对路径，图片提交fdfs完成后会被删除
- bin_full_path，ffmpeg可执行文件的绝对路径

- server_addr，服务器地址，支持ip:port或ip
- conf_full_path，配置文件绝对路径

### fdfs_client.conf
fdfs库需要的配置文件按照配置文件来修改对应值就好了