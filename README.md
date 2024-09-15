# Codename SplitBit
 Cost-keeping app for friends and households built using latest and greatest software stack featuring native clients.

 ## Development setup

### Server (/server)
The server handles the requests and communication with the database. SplitBit uses sqlite3 database as it's primary datastore and before running the server, it needs to have the database file called *test.db* present in the same directory. You can easily initialize such file using the command ```sqlite3 test.db < initdb.sql```. After that in order to build the app you run ```go build``` and ```./server``` runs the server. Now you clients are ready to interact with it.

### Web Client (/web)
Copy file .env.example to .env and adjust the url your client is running at (if you have changed it). Run ```npm install``` to get all the dependencies and ```npm run dev -- --open``` to get the client working. 

### Android Client (/android)
Open the project in Android Studio and you should be good to go. By default it uses a public instance hosted on isdc's servers. If you are developing server features at the same time, you might want to change that to make your phone/emulator do its requests to your computer. To do so, modify baseUrl inside *src/main/java/fi/isdc_helsinki/splitbit/repositories/ApiClient.kt* with the local ip address you get from running commands such as ```ip address``` or ```ifconfig``` AND the port number the server is running on. Then modify the ip address inside *app/src/main/res/xml/network_security_config.xml* to allow android to make these calls.
