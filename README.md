Requires [bashsh-0](https://github.com/mcrio/bashsh).

    ./build.sh

produces

    helloworld.exe

To make this available via the network, run

    go run serve_helloworld_exe.go

which will serve the file on `http://<your_ip>:9000/`.
