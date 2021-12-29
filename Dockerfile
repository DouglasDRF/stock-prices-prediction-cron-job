FROM golang:1.16

WORKDIR $GOPATH/src/stockpredictioncronjob
COPY . .

RUN go get -d -v 
RUN go install -v 

ENV PAST_STOCKS_REF=40
ENV STOCK_PREDICTIONS_API=

CMD ["stockpredictionscronjob"]
