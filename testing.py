from ctypes import cdll
import ctypes, cchardet, json
from bs4 import BeautifulSoup
lib = cdll.LoadLibrary("./main.so")
lib.cclientpy.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
lib.cclientpy.restype = ctypes.POINTER(ctypes.c_ubyte*8)
""" THERE IS A BIG MEMORY LEAK, BEWARE """

def newrequest(path, lister={}):
    try:
        ptr = lib.cclientpy(path.encode("utf-8"), str(lister).encode("utf-8"))
        length = int.from_bytes(ptr.contents, byteorder="little")
        data = bytes(ctypes.cast(ptr,
                ctypes.POINTER(ctypes.c_ubyte*(8 + length))
                ).contents[8:])
        return data
    except:
        pass

headers = {
    "Headers": [
    
    {
        "Name": "accept",
        "Value": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
    },
    {
        "Name": "accept-encoding",
        "Value": "gzip, deflate, br"
    },
    {
        "Name": "user-agent",
        "Value": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36"
    }
    ]
}

headerjson = json.dumps(headers)

x = newrequest("https://ja3er.com/json", headerjson)

print(x)
soup = BeautifulSoup(x, 'lxml')

with open("Page.html", "w+", encoding="utf-8") as f:
    f.write(str(soup))


