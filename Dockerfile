FROM embeddedenterprises/burrow as builder
RUN apk update && apk add build-base
RUN burrow clone https://github.com/jwuensche/autobahnausfahrt.git
WORKDIR $GOPATH/src/github.com/jwuensche/autobahnausfahrt
RUN burrow e && burrow b && cp bin/autobahnausfahrt /bin

FROM scratch
LABEL service "autobahnausfahrt"

COPY --from=builder /bin/autobahnausfahrt /bin/autobahnausfahrt
ENTRYPOINT ["/bin/autobahnausfahrt"]
CMD []

