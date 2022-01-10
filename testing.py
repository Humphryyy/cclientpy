from ctypes import cdll
import ctypes
import json

lib = cdll.LoadLibrary("./main.so")
lib.SendRequest.argtypes = [ctypes.c_char_p]
lib.SendRequest.restype = ctypes.c_char_p


request = json.dumps({
    "url": "https://ja3er.com/json",
    "method": "GET",
    "headers": [
        ["User-Agent",
            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"],
    ],
    "body": "",
    "allowRedirect": True,
    #  "proxy": "http://localhost:8888",
    "timeout": 10000,
    "pseudoHeaderOrder": [
        ":method",
        ":authority",
        ":scheme",
        ":path",
    ],
})


while (True):
    resp = lib.SendRequest(request.encode("utf-8"))
    print(resp)
