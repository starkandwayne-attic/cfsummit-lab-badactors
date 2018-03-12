from flask import Flask, render_template 
import sys, os


app = Flask(__name__)


@app.route('/')
def home():
    return ("hi there")

@app.route('/landing')
def landing():
    return ("welcome to our landing page")

@app.route('/exit')
def stop():
    os._exit(0)
    return "goodbye"

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0')