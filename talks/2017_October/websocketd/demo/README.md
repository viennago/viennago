## Websocketd Demo Time

The demo for this talk was based on the websocketd ["Ten second tutorial"](https://github.com/joewalnes/websocketd/wiki/Ten-second-tutorial).

We also had a look at using tail -f with websocketd, and showed how we can append to the file once the websocket is open, and it will be reflected in the open websocket session, on the client side.

I've included the count.sh script from the tutorial, as well as the example tail -f log file, for convenience.

Command used to demonstrate websocketd with tail -f:
`websocketd --port 1234 tail -f logfile.txt`
