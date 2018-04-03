from flask import Flask, render_template 
import sys, os


app = Flask(__name__)

port = int(os.getenv("PORT", 5000))

@app.route('/')
def home():
    return ("hi there")

@app.route('/exit')
def stop():
    os._exit(0)
    return "goodbye"

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0',port=port)