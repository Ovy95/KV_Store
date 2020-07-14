# KV_Store

- server runs on port 8080 
Currently

For now try to get an http server up and running that just prints “GET”, “SET”, “DELETE” instead of calling a store function

Next features to add

Store function info 


They key is a string and the value can be anything it could be a string, image, whatever you want that’s why we store it as a raw slice of bytes

You should be able to grab the key from the url with a simple function that splits the url on the first ‘/‘ and the key is anything after that

