FROM google/golang

MAINTAINER dockerlover@zju.edu.cn

ADD ./Mongoexample .
ADD tutorial/tuto_a.md .

EXPOSE 8080

CMD Mongoexample && tail -f 
