FROM google/golang

MAINTAINER dockerlover@zju.edu.cn

ADD ./mongo_store.go /gopath/
ADD tutorial/tuto_a.md /gopath/
ADD ./Mongoexample /gopath/

CMD /gopath/Mongoexample && tail -f 
