FROM freepascal/fpc:latest-minimal

WORKDIR /projektowanie-obiektowe
COPY ./1.pas ./1.pas

RUN fpc ./1.pas

ENTRYPOINT [ "./1" ]
