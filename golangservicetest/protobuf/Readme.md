
Windowa

descargar compilar de https://developers.google.com/protocol-buffers/docs/downloads#release-packages

descomprimit y copiar proto.exe en la carpeta bin del workspace de go

Luego descargar 

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}



protoc.exe --go_out=. *.proto

El comando anterior generar los archivos go necesarias para tabajar con los archivos protobuf
