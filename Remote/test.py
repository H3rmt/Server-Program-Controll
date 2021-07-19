
from time import sleep
import sys
import threading

def printshit():
    print("hi1")
    sys.stdout.flush()

    sleep(0.5)
    print("hi2")
    sys.stdout.flush()

    sleep(1)
    print("hi3")
    sys.stdout.flush()

    sleep(1.5)
    print("hi4")
    sys.stdout.flush()

    sleep(0.5)
    raise AttributeError("This is an err")

if __name__ == '__main__':
    thr = threading.Thread(target=printshit)
    thr.start()
    sleep(0.025)
    #thr2 = threading.Thread(target=printshit)
    #thr2.start()