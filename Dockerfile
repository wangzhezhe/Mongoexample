FROM google/golang

MAINTAINER dockerlover@zju.edu.cn

ADD ./Mongoexample .
ADD tutorial/tuto_a.md .

CMD /Mongoexample && tail -f 
