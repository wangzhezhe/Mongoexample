FROM google/golang

MAINTAINER dockerlover@zju.edu.cn

ADD ./mongo_store.go /gopath/
ADD tutorial/tuto_a.md /gopath/
ADD ./Mongoexample /gopath/

EXPOSE 8080
CMD /gopath/Mongoexample && tail -f 
