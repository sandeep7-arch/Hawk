import subprocess
import pickle
import hashlib

user_input = input("Enter command: ")

subprocess.call(user_input, shell=True)

data = input("Enter serialized data: ")
obj = pickle.loads(data.encode())

password = input("Password: ")
print(hashlib.md5(password.encode()).hexdigest())
