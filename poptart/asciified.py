#!/usr/bin/env python

from PIL import Image
import os
import sys
import pwd
getlogin = lambda: pwd.getpwuid(os.getuid())[0]
username = getlogin()




asciicode = " -.-,-:-;-i-r-s-X-A-2-5-3-h-M-H-G-S-#-9-B-&-@".split("-")
def asciified(filename):
    if filename != 'asciified.py':
        img = Image.open(filename).convert('LA')
        filename = filename.split('/')
        for x in filename:
            if '.' in x:
                filename = x
                break
        delname = filename
        filename = filename.split('.')
        img = img.resize((16,16))
        argument = ("{0}".format("grey")+filename[0])
        #os.remove('/home/{0}/go/src/localtoast.net/localtoast/bakery/poptart/101/'.format(username)+delname)
        img.save('/home/{0}/go/src/localtoast.net/localtoast/bakery/poptart/101/'.format(username)+argument+".png")

        img = Image.open('/home/{0}/go/src/localtoast.net/localtoast/bakery/poptart/101/grey{1}.png'.format(username, filename[0]))
        #finalimg = Image.new("LA", (32, 32))
        #grey = img.crop((0, 0, 128, 128))
        grey = img.load()
        #os.remove('~/go/src/localtoast.net/localtoast/bakery/poptart/101/grey{0}.png'.format(filename[0]))
        #print grey[0,0]
        # Create the final file to be made
        filename = '/home/{0}/go/src/localtoast.net/localtoast/bakery/poptart/101/{1}.txt'.format(username, filename[0])
        finalascii = open(filename, 'w+')
        #print finalascii
        for column in range(16):
            for row in range(16):
                rowcol = grey[row, column]
                asciinum = rowcol[0]/24

                finalascii.write(asciicode[asciinum])
            finalascii.write("\n")
        finalascii.close()
        return filename
