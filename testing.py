from ctypes import cdll
import ctypes, json

lib = cdll.LoadLibrary("./main.so")
lib.SendRequest.argtypes = [ctypes.c_char_p]
lib.SendRequest.restype = ctypes.POINTER(ctypes.c_ubyte*8)
""" THERE IS A BIG MEMORY LEAK, BEWARE """

def sendRequest(path, lister={}):
    try:
        ptr = lib.SendRequest(path.encode("utf-8"), str(lister).encode("utf-8"))
        length = int.from_bytes(ptr.contents, byteorder="little")
        data = bytes(ctypes.cast(ptr,
                ctypes.POINTER(ctypes.c_ubyte*(8 + length))
                ).contents[8:])
        return data
    except:
        pass




request = json.dumps({
    "url": "https://ja3er.com/json",
    "method": "GET",
    "headers": [
        ["User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"],
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

resp = sendRequest(request)

print(resp)
