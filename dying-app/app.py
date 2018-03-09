from flask import Flask, render_template


app = Flask(__name__)


@app.route('/')
def landing():
    return ("hi there")

@app.route('/exit')
def stop():
    exit()
    return ("goodbye")

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0')