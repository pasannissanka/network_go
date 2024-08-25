### GO Network

This project is a network communication tool written in Go. It consists of two main components: a master node and slave nodes. The master node listens for UDP messages from slave nodes, and upon receiving a message, it responds with its TCP address and IP. The slave node scans a specified IP and port range, sends a UDP message to the master node, and then establishes a TCP connection with the master node upon receiving the response.

### Concepts

- **Master Node**: Listens for UDP messages from slave nodes and responds with its TCP address and IP.
- **Slave Node**: Scans a specified IP and port range, sends a UDP message to the master node, and establishes a TCP connection upon receiving the master's response.
- **UDP Communication**: Used by the slave node to send initial messages to the master node.
- **TCP Communication**: Used to establish a reliable connection between the master and slave nodes after the initial UDP handshake.

### Usage

#### Running the Master Node

Environment variables can be set to configure the master node. The following environment variables can be set:

```env
ID=1          # Node ID
TCP_PORT=8888 # TCP port of the node
UDP_PORT=8881 # UDP port of the node

MASTER=true   # Flag to indicate that this is the master node

```
Save the above configuration in a file named `.env.{node_ID}` in the root directory of the project. For example, for the master node, save the configuration in a file named `.env.1`.


To run the master node, use the following command:

```sh
./network_go 1
```

The master node will start listening for UDP messages from slave nodes.

#### Running the Slave Node

To run the slave node, use the following command:

```sh
./network_go 2
```

The slave node will start scanning the specified IP and port range, send a UDP message to the master node, and establish a TCP connection upon receiving the master's response.

### Guidelines

1. **Configuration**: Ensure that the IP and port ranges are correctly configured in the slave node before running the program.
2. **Network Setup**: Make sure that the master and slave nodes are on the same network or can communicate with each other over the network.
3. **Firewall Settings**: Ensure that the necessary ports are open on both the master and slave nodes to allow UDP and TCP communication.
4. **Error Handling**: Check the logs for any errors or issues during the execution of the nodes. The logs will provide information about the status of the nodes and any communication issues.

### Example

1. **Start the Master Node**:

   ```sh
   ./network_go 1
   ```

2. **Start the Slave Node**:

   ```sh
   ./network_go 2
   ```

3. **Logs**:
   - The master node will log when it receives a UDP message and when it sends its TCP address and IP.
   - The slave node will log the IP and port range it is scanning, when it sends the UDP message, and when it establishes the TCP connection.

### Code Structure

- **client**: Contains the code for the slave node, including IP and port scanning, sending UDP messages, and establishing TCP connections.
- **server**: Contains the code for the master node, including listening for UDP messages and responding with TCP address and IP.

### Dependencies

- **Go**: Ensure that Go is installed on your system to build and run the project.
- **Network Configuration**: Proper network configuration to allow communication between master and slave nodes.

### Building the Project

To build the project, navigate to the project directory and run:

```sh
go build -o network_go
```

This will create an executable named `network_go` that can be used to run the master and slave nodes.

### TODO

- Implement additional features such as data exchange between master and slave nodes.
- Enhance the network discovery mechanism to support dynamic node addition and removal.
- Master node selection and failover mechanisms for improved reliability.
- Connection pooling and load balancing for handling multiple slave nodes.
- Security improvements such as encryption and authentication for secure communication.
