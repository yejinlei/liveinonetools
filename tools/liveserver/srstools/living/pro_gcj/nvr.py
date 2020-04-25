# -*- coding: utf-8 -*-
'''
The MIT License (MIT)

Copyright (c) 2013-2015 SRS(ossrs)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
'''

"""
the api-server is a default demo server for srs to call
when srs get some event, for example, when client connect
to srs, srs can invoke the http api of the api-server
"""
from kafka import *

from multiprocessing import Process, Queue
from Queue import Empty
import sys
import imp
# reload sys model to enable the getdefaultencoding method.
imp.reload(sys)
sys.setdefaultencoding('utf-8') 
# set the default encoding to utf-8
# using exec to set the encoding, to avoid error in IDE.
# exec("sys.setdefaultencoding('utf-8')")
# assert sys.getdefaultencoding().lower() == "utf-8"

import os, json, time, datetime, threading
import configparser



import subprocess

CONST_RESERVER_FLV_COUNT = 5

TARGET_FILE_EXT = ".flv"

file_log = None


# simple log functions.
def trace(msg, is_write_file=False):
    date = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    print "[main_process][%s][trace] %s"%(date, msg)
    
    if is_write_file is False :
        return    

    global file_log
    if file_log is None :
        file_log = open('nvr_rtsp2rtmp.log', mode='w')
    file_log.write('[' + date + ']' + '[trace]' + msg + '\n')
    file_log.flush()



#FFMPEG_FULL_PATH = 'ffmpeg'
FFMPEG_FULL_PATH = '/usr/local/bin/ffmpeg'

class task_item(object):

    old_msg_dict = {}

    rtmp_str = ''
    rtsp_str = ''

    is_start = 0
    start_timestamp = None

    is_thread_need_exit = 0

    process_obj = None
    cmd_stdout_lines = ''
    cmd_stderr_lines = ''

    run_fail = 0
    already_ret = 0

    def __init__(self):
        pass

    def is_same_obj(self, other_task_item_obj):
        if self.rtmp_str != other_task_item_obj.rtmp_str:
            return False

        if self.rtsp_str != other_task_item_obj.rtsp_str:
            return False

        if self.is_start != other_task_item_obj.is_start:
            return False

        if self.is_thread_need_exit != other_task_item_obj.is_thread_need_exit:
            return False

        return True

    def stop_task(self):
        self.is_start = 0

    def start_task(self):
        self.is_start = 1

    def is_stop_task(self):
        return self.is_start == 0

    def set_thread_flag(self):
        self.is_thread_need_exit = 1

    def is_exit_flag(self):
        return self.is_thread_need_exit == 1

    def set_url(self, rtmp_str, rtsp_str):
        if len(rtmp_str) <= 0 or len(rtsp_str) <= 0:
            return False

        self.rtmp_str = rtmp_str
        self.rtsp_str = rtsp_str

        return True

    def get_url(self):
        return self.rtsp_str, self.rtsp_str

def process_thread(call_back_obj) :
    if call_back_obj is None:
        return

    while True:
        ret = call_back_obj._run_loop()
        if ret is False:
            return

class process_nvr_rtsp2rtmp(object):
    log_obj = None

    task_queue = Queue()

    send_msg_kafka = None
    task_max_exec_time = -1

    def __init__(self, log_obj, send_msg_kafka, task_max_exec_time):
        self.log_obj = log_obj
        self.send_msg_kafka = send_msg_kafka
        self.task_max_exec_time = task_max_exec_time

    def process_msg(self, msg_dict_dict):
        task_item_dict_tmp = {}

        if len(msg_dict_dict) > 0:
            for key, dict_obj in msg_dict_dict.items():
                task_item_obj = task_item()
                task_item_obj.set_url(dict_obj['rtmp'], dict_obj['rtsp'])
                str_op = dict_obj['op'].strip().lower()
                if str_op == "close":
                    task_item_obj.stop_task()
                elif str_op == "open":
                    task_item_obj.start_task()
                else:
                    return
                task_item_obj.old_msg_dict = dict_obj
                task_item_dict_tmp[dict_obj['rtmp']] = task_item_obj

        if len(task_item_dict_tmp) > 0:
            self.task_queue.put_nowait(task_item_dict_tmp)

        pass

    def _run_loop(self):

        task_item_dict_local = {}

        while True:
            task_item_exec_result_dict = {}
            task_item_dict_obj = self._get_task()
            if self._do_task_run(task_item_dict_local, task_item_dict_obj, task_item_exec_result_dict) is False:
                break

            #build result info
            for key, obj_tmp in task_item_exec_result_dict.items():
                if obj_tmp.already_ret == 1:
                    continue

                obj_tmp.already_ret = 1

                str_dict = obj_tmp.old_msg_dict
                if obj_tmp.run_fail == 1:
                    str_dict['status'] = "0"
                else:
                    str_dict['status'] = "1"

                str_dict['apISource'] = "NVR_RTSP_RSP"
                msg = json.dumps(str_dict).encode('utf-8')
                self.send_msg_kafka.send('TOPIC_NVR_RTSP_RSP_CHANGE', key=key.encode('utf-8'), value=msg, partition=0)

        return True

    def _get_task(self):
        dict_dict_obj = {}

        try:
            dict_dict_obj = self.task_queue.get(block=True, timeout=0.1)
        except Empty:
            pass

        return dict_dict_obj

    def _do_task_run(self, task_item_dict_local, task_item_dict_obj, task_item_exec_result_dict):
        for key, obj_tmp in task_item_dict_obj.items():
            if key not in task_item_dict_local:
                task_item_dict_local[key] = obj_tmp
                if obj_tmp.is_exit_flag():
                    task_item_dict_local = {}
                    return False
            else:

                dict_cur = task_item_dict_local[key]
                if dict_cur.is_same_obj(obj_tmp):
                    continue

                if obj_tmp.is_exit_flag():
                    task_item_dict_local = {}
                    return False

                dict_cur.is_start = obj_tmp.is_start
                obj_tmp = dict_cur

            self._run_cmd(obj_tmp)
            task_item_exec_result_dict[key] = obj_tmp
            if obj_tmp.is_stop_task():
                task_item_dict_local.pop(key)

        #
        self._check_process(task_item_dict_local, task_item_exec_result_dict)

        return True

    def _run_cmd(self, task_item_obj):
        task_item_obj.already_ret = 0
        if task_item_obj.process_obj is not None:
            task_item_obj.process_obj.kill()
            task_item_obj.process_obj.wait()
            task_item_obj.process_obj = None
            task_item_obj.start_timestamp = None

        if task_item_obj.is_stop_task():
            task_item_obj.old_msg_dict["op"] = "close"
            trace("recv stop task req, do stop " + task_item_obj.rtsp_str + ", " + task_item_obj.rtmp_str, True)
            return True

        cmd_list = [FFMPEG_FULL_PATH, '-rtsp_transport', 'tcp', '-re', '-i',
                    task_item_obj.rtsp_str,
                    '-vcodec', 'copy', '-an', '-f', 'flv',
                    task_item_obj.rtmp_str]

        #cmd_list = FFMPEG_FULL_PATH + ' -rtsp_transport tcp -re -i \'' + task_item_obj.rtsp_str + '\' -vcodec copy -an -f flv ' + task_item_obj.rtmp_str
        #trace(cmd_list)
        trace("start task " + task_item_obj.rtsp_str + ", " + task_item_obj.rtmp_str, True)

        task_item_obj.process_obj = subprocess.Popen(cmd_list, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        time.sleep(1)
        task_item_obj.start_timestamp = datetime.datetime.now()

        return True

    def _check_process(self, task_item_dict_local, task_item_exec_result_dict):
        death_process = []

        for key, task_item_obj in task_item_dict_local.items():
            if task_item_obj.process_obj is None:
                task_item_dict_local.pop(key)
                continue

            task_item_obj.process_obj.poll()

            # need stop task?
            now_time = datetime.datetime.now()
            cost_time = (now_time - task_item_obj.start_timestamp).seconds
            if self.task_max_exec_time != -1 and cost_time >= self.task_max_exec_time:
                task_item_obj.stop_task()
                task_item_obj.process_obj.kill()
                task_item_obj.process_obj.wait()
                death_process.append(key)
                task_item_obj.process_obj = None
                task_item_obj.start_timestamp = None
                trace("auto stop task, " + task_item_obj.rtsp_str + ", " + task_item_obj.rtmp_str, True)
                continue

            #cmd_lines = task_item_obj.process_obj.stdout.readlines()
            #for line in cmd_lines:
            #   task_item_obj.cmd_stdout_lines += line

            #if len(cmd_lines) > 0:
            #   trace(cmd_lines)

            #cmd_lines = task_item_obj.process_obj.stderr.readlines()
            #for line in cmd_lines:
            #    task_item_obj.cmd_stderr_lines += line

            #if len(cmd_lines) > 0:
            #    trace(cmd_lines)

            if task_item_obj.process_obj.returncode is not None:
                task_item_obj.stop_task()
                death_process.append(key)
                cmd_lines = task_item_obj.process_obj.stderr.readlines()
                print_str = ''
                for line in cmd_lines:
                   task_item_obj.cmd_stderr_lines += line
                   print_str += line

                if len(print_str) > 0:
                   trace(print_str)
                continue

        for death_process_item in death_process:
            task_item_dict_local.pop(death_process_item)
            if death_process_item in task_item_exec_result_dict:
                task_item_exec_result_dict[death_process_item].run_fail = 1
                pass

        pass

class msg_pump(object):

    srv_str = ''
    recv_topic = ''

    recv_obj = None
    send_obj = None

    def __init__(self, srv_str, recv_topic):
        self.srv_str = srv_str
        self.recv_topic = recv_topic

        self.recv_obj = KafkaConsumer(recv_topic, bootstrap_servers=srv_str,
                                           api_version=(0, 10, 1), auto_offset_reset='latest')

        self.send_obj = KafkaProducer(bootstrap_servers=srv_str, api_version=(0, 10, 1))
        pass

    def do_pump_and_send_back(self):

        msg_dict_dict = {}
        msg_dict_tmp = {}

        try:
            new_msg = self.recv_obj.poll(timeout_ms=0.05)
        except Exception as e:
            #self.log_obj.trace_thread_safe("self.consumer_obj.poll have some err", True)
            return False, msg_dict_dict

        icount = 0
        for data in new_msg.values():
            for line in data:
                msg_dict_tmp = json.loads(line.value)
                msg_dict_dict[str(icount)] = msg_dict_tmp
                icount += 1

        return True, msg_dict_dict

'''
main code start.
'''
# donot support use this module as library.
if __name__ == "__main__":
    config = configparser.ConfigParser()
    config.read('/nvr.conf')

    task_max_exec_time = config['common']['task_max_exec_time']
    task_max_exec_time = int(task_max_exec_time)

    msg_pump_obj = msg_pump(srv_str=config['kafka']['bootstrap_servers'], recv_topic="TOPIC_NVR_RTSP_REQ_CHANGE")

    process_obj = process_nvr_rtsp2rtmp(None, msg_pump_obj.send_obj, task_max_exec_time)

    thread_for_process = threading.Thread(target=process_thread, args=(process_obj,))
    thread_for_process.start()

    while True:
        ret, msg_dict_dict = msg_pump_obj.do_pump_and_send_back()
        process_obj.process_msg(msg_dict_dict)


