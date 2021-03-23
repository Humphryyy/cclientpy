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



Credit:
https://github.com/x04/cclient
