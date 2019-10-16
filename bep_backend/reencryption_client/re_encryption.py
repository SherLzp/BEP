from umbral import config
from umbral.curve import SECP256K1
from umbral import keys, signing
from umbral import pre
import base64
import pickle
import os

config.set_default_curve(SECP256K1)

def hex2bytes(hex_str):
    return bytes(bytearray.fromhex(hex_str))

class ReEncryption:

    def generateKeys() -> bytes:
        private_key = keys.UmbralPrivateKey.gen_key()
        public_key = private_key.get_pubkey()
        res = {}
        res['private_key'] = private_key.to_bytes().hex()
        res['public_key'] = public_key.to_bytes().hex()
        return res

    def encryptInfo(msg, privkey) -> str:
        if isinstance(msg, str):
            msg = msg.encode()
        owner_pri_key = keys.UmbralPrivateKey.from_bytes(hex2bytes(privkey))
        owner_pub_key = owner_pri_key.get_pubkey()
        ciphertext, capsule = pre.encrypt(owner_pub_key, msg)
        res = {}
        res['ciphertext'] = ciphertext.hex()
        res['capsule'] = capsule.to_bytes().hex()
        return res

    def generateKfrags(privkey, signkey, recv_pubkey, threshold, N):
        if isinstance(recv_pubkey, str):
            recv_pubkey = keys.UmbralPublicKey.from_bytes(hex2bytes(recv_pubkey))
        owner_pri_key = keys.UmbralPrivateKey.from_bytes(hex2bytes(privkey))
        owner_signing_key = keys.UmbralPrivateKey.from_bytes(hex2bytes(signkey))
        signer = signing.Signer(private_key=owner_signing_key)
        kfrags = pre.generate_kfrags(delegating_privkey=owner_pri_key,
                                     signer=signer,
                                     receiving_pubkey=recv_pubkey,
                                     threshold=threshold,
                                     N=N)
        res = {}
        bytes_kfrags = list()
        for v in kfrags:
            bytes_kfrags.append(v.to_bytes().hex())
        res['kfrags'] = bytes_kfrags
        return res

    def reencrypt(pubkey, verikey, recv_pubkey, kfrags_list, capsule):
        kfrags = list()
        for v in kfrags_list:
            kfrag = pre.KFrag.from_bytes(hex2bytes(v))
            kfrags.append(kfrag)
        a_pub_key = keys.UmbralPublicKey.from_bytes(hex2bytes(pubkey))
        a_ver_key = keys.UmbralPublicKey.from_bytes(hex2bytes(verikey))
        b_pub_key = keys.UmbralPublicKey.from_bytes(hex2bytes(recv_pubkey))
        capsule = pre.Capsule.from_bytes(hex2bytes(capsule), a_pub_key.params)
        capsule.set_correctness_keys(delegating=a_pub_key,
                                     receiving=b_pub_key,
                                     verifying=a_ver_key)
        cfrags = list()
        for kfrag in kfrags:
            cfrag = pre.reencrypt(kfrag=kfrag, capsule=capsule)
            cfrags.append(cfrag.to_bytes().hex())
        return cfrags

    def decrypt(pubkey, verikey, recv_privkey, ciphertext, cfrags_list, capsule) -> str:
        pri_key = keys.UmbralPrivateKey.from_bytes(hex2bytes(recv_privkey))
        pub_key = pri_key.get_pubkey()
        a_pub_key = keys.UmbralPublicKey.from_bytes(hex2bytes(pubkey))
        a_ver_key = keys.UmbralPublicKey.from_bytes(hex2bytes(verikey))
        capsule = pre.Capsule.from_bytes(hex2bytes(capsule), pub_key.params)
        capsule.set_correctness_keys(delegating=a_pub_key, receiving=pub_key, verifying=a_ver_key)
        cfrags = list()
        for cf in cfrags_list:
            cfrag = pre.CapsuleFrag.from_bytes(hex2bytes(cf))
            cfrags.append(cfrag)
        for cfrag in cfrags:
            capsule.attach_cfrag(cfrag)
        if isinstance(ciphertext, str):
            ciphertext = hex2bytes(ciphertext)
        msg_bytes = pre.decrypt(ciphertext=ciphertext,
                                capsule=capsule,
                                decrypting_key=pri_key)
        return msg_bytes.decode()
