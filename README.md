# broker
Tool for connect gVisor and web interface
## Usage
1. Run gVisor
2. Run RabbitMQ
3. Run this [tool](#config)
4. Run connect [web](#client) interface to soket

## Config

### Run with command line params
> Arguments:
```
  -h  --help     Print help information

  -a  --address  Socket address (for example "localhost:9988" or "/tmp/runsc/log.unix")

  -t  --type     Socket type ("tcp" or "unix")

  -r  --rabbit   Rabbit address

  -q  --queue    Rabbit queue

  -p  --port     WebSoket port
```

### Run with .env config

Example `.env` file:

```env
ADDRESS=localhost:9988
TYPE=tcp
WS_PORT=8080
RABBIT_ADDRESS=amqp://guest:guest@localhost:5672/
QUEUE=main
```

### Default config
if .env doesn't exists and no command line arguments, use default config

For more information see example `.env`

## Client

If you want try this tool without web interface, with postman, for example:

1. Connect to `ws://<host>:<ws_port_from_config>`
2. Send pull-message: `pull`
3. Read your messages!
