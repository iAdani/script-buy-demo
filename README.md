# script-buy project

## Description


## Installation

### Dependencies
* Node.js version 18 LTS from https://nodejs.org/en/ which includes npm by default
* Android Emulator or Android Device with USB Debugging Enabled (We chose Android Studio Emulator)
* Golang from https://golang.org/dl/


1) First thing you need to do is clone the repository to your local machine.
2) Now we will install the dependencies using npm:
    ```bash
    npm install
    ```
3) Now download the app apk from the repository and install it on your device or emulator.

4) if running on windows you may need to run the following command in powershell to allow the app to run:
    ```bash
    Set-ExecutionPolicy RemoteSigned
    ```
5) Now you can run the app using the following command:
    ```bash
   expo start
    ```
   then press 'a' to run on android emulator or device and choose the developer build option.
6) The app should now be running on your device or emulator.
7) To start the server, open a new terminal and type
    ```bash
    cd ./Go_Project
    go run ./server/run_server.go
    ```
8) After the server is on, you can run tests with
   ```bash
   npm test
   ```
