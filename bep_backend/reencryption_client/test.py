import requests as req
import json

base_url = "http://127.0.0.1:5000"
def generate_keys():
    path = "/generateKeys"
    url = base_url + path
    r = req.get(url)
    return json.loads(r.text)

def encrypt_msg(msg, privkey):
    path = "/encryptData"
    url = base_url + path
    payload = {
            "msg": msg,
            "private_key": privkey
    }
    r = req.post(url, data = json.dumps(payload))
    return json.loads(r.text)

def generate_kfrags(privkey, signkey, recv_pubkey, threshold, N):
    path = "/generateKfrags"
    url = base_url + path
    payload = {
            "private_key": privkey,
            "signing_key": signkey,
            "receiving_pubkey": recv_pubkey,
            "threshold": threshold,
            "n": N,
    }
    r = req.post(url, data = json.dumps(payload))
    return json.loads(r.text)['kfrags']
    
def reencrypt(pubkey, verikey, recv_pubkey, kfrags_list, capsule):
    path = "/reencrypt"
    url = base_url + path
    payload = {
            "public_key": pubkey,
            "verifying_key": verikey,
            "receiving_pubkey": recv_pubkey,
            "kfrags": kfrags_list,
            "capsule": capsule,
    }
    r = req.post(url, data = json.dumps(payload))
    return json.loads(r.text)

def decrypt(pubkey, verikey, recv_privkey, ciphertext, cfrags_list, capsule):
    path = "/decrypt"
    url = base_url + path
    payload = {
            "public_key": pubkey, 
            "verifying_key": verikey,
            "receiving_privkey": recv_privkey,
            "ciphertext": ciphertext,
            "cfrags": cfrags_list,
            "capsule": capsule,
    }
    r = req.post(url, data = json.dumps(payload))
    return r.text

if __name__ == '__main__':
    alice_key = generate_keys()
    alice_sign = generate_keys()
    bob_key = generate_keys()
    
    msg = "hello world"
    enc = encrypt_msg(msg, alice_key['private_key'])
    ciphertext = enc['ciphertext']
    capsule = enc['capsule']
    
    kfrags_list = generate_kfrags(alice_key['private_key'], alice_sign['private_key'], bob_key['public_key'], threshold = 1, N = 2)
    
    cfrags = reencrypt(alice_key['public_key'], alice_sign['public_key'], bob_key['public_key'], kfrags_list, capsule)

    decrypted_msg = decrypt(alice_key['public_key'], alice_sign['public_key'], bob_key['private_key'], ciphertext, cfrags, capsule)
    if decrypted_msg == msg:
        print("test case passed")
    else:
        print("test case failed")
