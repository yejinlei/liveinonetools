# coding:utf-8
#  imageai下载地址：https://github.com/OlafenwaMoses/ImageAI
#  resnet50_coco_best_v2.1.0.h5 模型下载地址：https://github.com/fizyr/keras-retinanet/releases/
from imageai.Detection import ObjectDetection  # 导入了 ImageAI 目标检测类
import cv2
import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'
import matplotlib.pyplot as plt


if __name__=='__main__':
    model_path = ('./resnet50_coco_best_v2.1.0.h5')
    path = os.path.abspath(model_path)
    detector = ObjectDetection()  # 定义了目标检测类
    detector.setModelTypeAsRetinaNet()  # 模型的类型设置为 RetinaNet
    detector.setModelPath(path)  # 将模型路径设置为 RetinaNet 模型的路径
    detector.loadModel()  # 模型加载到的目标检测类
    cap = cv2.VideoCapture(0)
    count = 0
    while (True):  # 逐帧显示
        ret, img = cap.read()
        count += 1
        if count % 100 != 0:
            continue
        detections, objects_path = detector.detectObjectsFromImage(input_image=img, input_type='array',output_type='array')
        for eachObject, eachObjectPath in zip(detections, objects_path):
            print(eachObject["name"] + " : " + eachObject["percentage_probability"] )
            print("Object's image saved in " + eachObjectPath)
            print("--------------------------------")

        cv2.imshow('image', detections[0])
        if cv2.waitKey(1) & 0xFF == ord(' '):
            break
    cap.release()  # 释放摄像头
    cv2.destroyAllWindows()  # 释放窗口资源
