# Word Frequency Server

This project aims to build a simple Go-based HTTP server that counts the frequency of a given word in the downloaded text file from Gutenberg project.

This server is built as a part of the [Load Testing Utility](https://github.com/siddhant-vij/Load-Testing-Utility), which is used to test the performance of this server in terms of Latency and Throughput w.r.t the number of server threads used to process the frequency count requests.

<br>

## Table of Contents

1. [Server Implementation](#server-implementation)
1. [Concurrent Design](#concurrent-design)
1. [Notes](#notes)
1. [Contributing](#contributing)
1. [License](#license)

<br>

## Server Implementation

- The server is implemented using Go's Networking API (no framework is used).
- It listens for HTTP GET requests and processes them to return the frequency of a specified word in `resources/war_and_peace.txt`.

<br>

## Concurrent Design
- The server is implemented using `numServerThreads` goroutines to process the requests concurrently - as an input given by the Load Testing Utility for server threads.
- The server waits for all the goroutines to finish their execution and then, returns the frequency of the word in the text file.
- The server implements the Producer-Consumer design pattern with a channel to handle the tasks.
- Basic throttling is implemented to check for and avoid channel overflow, maintaining a high throughput.
- This design is used to achieve the Load Testing Utility's requirement to test the server's performance.

<br>

## Notes
- Ensure that this server is running before starting the load testing utility.
- The server and load testing utility are designed for educational purposes to demonstrate the effects of multithreading on performance.

<br>

## Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.
1. **Fork the Project**
2. **Create your Feature Branch**: 
    ```bash
    git checkout -b feature/AmazingFeature
    ```
3. **Commit your Changes**: 
    ```bash
    git commit -m 'Add some AmazingFeature'
    ```
4. **Push to the Branch**: 
    ```bash
    git push origin feature/AmazingFeature
    ```
5. **Open a Pull Request**

<br>

## License

Distributed under the MIT License. See [`LICENSE`](https://github.com/siddhant-vij/Word-Frequency-Server/blob/main/LICENSE) for more information.