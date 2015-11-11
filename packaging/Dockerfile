FROM alpine:3.2

MAINTAINER @cloudnautique
ENV COWBELL_RELEASE v0.0.5
ADD https://github.com/cloudnautique/cowbell/releases/download/${COWBELL_RELEASE}/cowbell /cowbell
RUN chmod +x cowbell

ENTRYPOINT ["/cowbell"]
