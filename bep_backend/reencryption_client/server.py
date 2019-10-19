from flask import Flask, request
import re_encryption
from flask_cors import CORS
import json

app = Flask(__name__)
CORS(app, supports_credentials=True)

@app.route('/generateKey', methods=["GET"])
def gen_key():
    res = re_encryption.ReEncryption.generateKey()
    return json.dumps(res)

@app.route('/encryptData', methods=["POST"])
def upload():
    params = json.loads(request.get_data(as_text=True))
    res = re_encryption.ReEncryption.encryptInfo(params['msg'], params['private_key'])
    return json.dumps(res)


@app.route('/generateKfrags', methods=["POST"])
def generateKfrags():
    params = json.loads(request.get_data(as_text=True))
    res = re_encryption.ReEncryption.generateKfrags(params['private_key'], params['signing_key'], params['receiving_pubkey'], int(params['threshold']), int(params['n']))
    return json.dumps(res)

@app.route('/reencrypt', methods=["POST"])
def reencryption():
    params = json.loads(request.get_data(as_text=True))
    pubkey = params['public_key']
    verikey = params['verifying_key']
    recv_pubkey = params['receiving_pubkey']
    kfrags_list = params['kfrags']
    capsule = params['capsule']
    res = re_encryption.ReEncryption.reencrypt(pubkey, verikey, recv_pubkey, kfrags_list, capsule)
    return json.dumps(res)

@app.route('/decrypt', methods=["POST"])
def decrypt():
    params = json.loads(request.get_data(as_text=True))
    pubkey = params['public_key']
    verikey = params['verifying_key']
    recv_privkey = params['receiving_privkey']
    ciphertext = params['ciphertext']
    cfrags_list = params['cfrags']
    capsule = params['capsule']
    res = re_encryption.ReEncryption.decrypt(pubkey, verikey, recv_privkey, ciphertext, cfrags_list, capsule)
    return res

if __name__ == '__main__':
    app.run(debug = True)
