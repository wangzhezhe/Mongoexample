FROM google/golang

MAINTAINER dockerlover@zju.edu.cn

ADD ./mongo_store.go /gopath/
ADD tutorial/tuto_a.md /gopath/
ADD ./mongo_store.go /gopath/
#ADD ./gopkg.in /gopath/src/gopkg.in/

WORKDIR /gopath/

RUN go get gopkg.in/mgo.v2

RUN go build mongo_store.go


EXPOSE 8080
CMD /gopath/mongo_store && tail -f 
