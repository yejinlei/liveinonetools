#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Auther: Ray
# @Time: 2019/11/5 10:40

from prometheus_client import Gauge, generate_latest, CollectorRegistry
from flask import Response, Flask
import os
import json
import time, datetime
import re

app = Flask(__name__)

# 定义数据仓库
REGISTRY = CollectorRegistry()
media_all = Gauge('media_all', "media_all", ['group'], registry=REGISTRY)

@app.route("/metrics")
def ApiResponse():
    CAMERA_CONF = os.getenv("CAMERA_CONF")
    media_url = 'rtsp://10.10.10.30:8554/live/2'
    cmdLine = "ffprobe -v quiet -print_format json -show_format -show_streams -print_format json=c=1 " + media_url
    #cmdLine = "ffprobe -v quiet -show_format -show_streams " + media_url
    cmdResult = os.popen(cmdLine).readlines()
    media_all.labels(group=cmdResult).set(1)

    return Response(generate_latest(REGISTRY), mimetype="text/plain")


@app.route('/health')
def index():
    return "Media metrics"


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=3999)
