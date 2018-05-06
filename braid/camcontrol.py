#!/usr/bin/env python
from pygame import camera as cammod
import pygame
import StringIO
import os
from PIL import Image, ImageDraw
import asciified as asc
import subprocess
import sys
import time
username = os.getlogin()
def main():
    pygame.init()
    cammod.init()
    cam1 = cammod.list_cameras()
    #print(cam1[0])

    cam = cammod.Camera('/dev/video0', (1024, 768))
    count = 0
    cam.start()
    ready = False
    thumbnail = 32, 32
    while True :

        #exists = os.open("/home/turanga/1101/101/image"+str(count)+".png", os.O_CREAT)

        while ready != True:
            ready = cam.query_image()
        if ready:
            #print("Camera ready for picture "+ str(count))
            path = "/home/"+username+"/1101/101/image.png"
            image = cam.get_image()
            pygame.image.save(image, path)
            pathText = asc.asciified(path)
            ready = cam.query_image()
            #return pathText
        
    cam.stop()

main()
