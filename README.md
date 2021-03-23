# cclientpy
Calling x04's cclient from Python using ctypes, I take no credit for cclient. 

Pros:

  - Ja3 in python
  - Headers
  - Faster then expected

Cons:
  - Large memory leak, do not use this for constant monitoring.
  - No proxy support
  - Headers are annoying
  - There are probably a lot more cons that I can't think of.


You can build the .so file using ```go build -o main.so -buildmode=c-shared main.go```

Output of testing.py: ```b'{"ja3_hash":"b32309a26951912be7dba376398abc3b", "ja3": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0", "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36"}'```

Credit:
https://github.com/x04/cclient
