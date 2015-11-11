FROM golang:1.4.3

ADD ./scripts/bootstrap /scripts/bootstrap
RUN /scripts/bootstrap
