
from time import sleep
import sys
import threading

def printshit():
    print("IMPORTANT|","hi1")

    sleep(0.5)
    print("NORMAL|","hi2")

    sleep(1)
    print("LOW|","hi3")

    sleep(1.5)
    print("hi4")
    

    sleep(10.5)
    raise AttributeError("This is an err")

if __name__ == '__main__':
    thr = threading.Thread(target=printshit)
    thr.start()
    sleep(0.025)
    #thr2 = threading.Thread(target=printshit)
    #thr2.start()