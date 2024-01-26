### What?
Prints a random port following all 3 of these rules to stdout
- the random port should be an integer between 1024-65000
- the port should not be in use on localhost already
- the port should not be one of the known ports listed here (https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers)

### Why?
GoLang practice

### How?
`git clone https://github.com/virtualbeck/random-ephemeral-port.git`\
`cd random-ephemeral-port`\
`go build`\
`chmod +x random-ephemeral-port`\
`./random-ephemeral-port`
