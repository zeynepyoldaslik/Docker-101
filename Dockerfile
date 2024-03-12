FROM golang:1.22
WORKDIR /firstapp
COPY /FirstDockerfile . 
CMD ["calculator.exe"]


