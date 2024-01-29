# README: SEEKNET 🚀

SeekNet is a chat application crafted for the command-line interface (CLI), specifically tailored for use within a terminal and bypassing the necessity of a browser interface. While it's not yet ready for production-level deployment, you're encouraged to try it out and have fun with the terminal through the provided setup.

#### 🚀 Blast Off: Get It Running Locally 🛠

##### Linux

### 🚨Setting up the Environment:

1. Install Golang:
   ```bash
   sudo apt install golang-go
   ```
2. Clone the repository:
    ```bash
   git clone [repo_url] && cd [repo_name]
   ```
3. Install the wscat client:
   ```bash
   sudo apt update
   sudo apt install nodejs npm
   sudo npm install -g wscat
   ```
4. Cleaning the go.mod file
   ```bash
   go mod tidy
   ```


 ### 🔥Ignite the Engines:

   - Golang:
     ```bash
     go build
     ./seeknet
     ```
   - Now open up a few terminals and to this address:
     ```bash
     wscat -c ws://localhost:5000/ws
     ```

     ### Now enter the password and username and test it out


     # DEMO:
     
      [demo-seeknet.webm](https://github.com/lla-dane/SeekNet/assets/120122716/401daebb-62f2-452c-9f96-def172175654)

       This project is dedicated to testing and exploring the functionality of websockets in Golang, utilizing the
        ```bash
        github.com/gorilla/websocket
        ```
      The focus is on implementing and fine-tuning websocket features within the project, allowing for comprehensive testing and assessment of the capabilities provided by this particular package in the Go programming language. The goal is to gain       a deeper understanding of websockets and their integration in Golang through practical experimentation and analysis.
 








  
