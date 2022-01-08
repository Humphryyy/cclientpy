There is a large memory leak, use with caution.

You can build the .so file using ```go build -o main.so -buildmode=c-shared main.go models.go```

Output of testing.py: ```b'{"headers":[["Content-Length","318"],["Connection","keep-alive"],["Set-Cookie","visited=b32309a26951912be7dba376398abc3b"],["Access-Control-Allow-Origin","*"],["Server","nginx"],["Date","Sat, 08 Jan 2022 08:33:17 GMT"],["Content-Type","application/json"]],"body":"{\\"ja3_hash\\":\\"b32309a26951912be7dba376398abc3b\\", \\"ja3\\": \\"771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0\\", \\"User-Agent\\": \\"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36\\"}"}'```

Credit to [Carcraftz](github.com/Carcraftz) for the libraries used