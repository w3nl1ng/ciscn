FROM golang:latest

RUN mkdir /app
COPY ./ /app
RUN apt update
RUN apt install nmap -y

WORKDIR /app
ENV GOPROXY "https://goproxy.cn,direct"
RUN go mod tidy
RUN go build -o main .

CMD [ "./main", "-f", "iplist.txt"]