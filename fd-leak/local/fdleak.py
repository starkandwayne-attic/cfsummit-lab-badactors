from flask import Flask, render_template 
import sys, os, time


app = Flask(__name__)


@app.route('/')
def home():
    return ("hi there")

@app.route('/leak')
def leak():
    fdlist=[]
    for i in range(2000):
        fd= open("files/test"+str(i)+".txt", "w")
        fd.write("hi")
        fdlist.append(fd) 
    time.sleep(99999)
    return "goodbye"

@app.route('/exit')
def stop():
    os._exit(0)
    return "goodbye"

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0')