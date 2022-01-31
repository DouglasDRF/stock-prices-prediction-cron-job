FROM golang:1.16

WORKDIR $GOPATH/src/stockpredictioncronjob
COPY . .

RUN go get -d -v 
RUN go install -v 


ENV STOCK_PREDICTIONS_API=
ENV STOCK_PREDICTIONS_API_KEY=
ENV STOCK_PREDICTIONS_API_SECRET=
ENV PAST_STOCKS_REF=40

CMD ["stockpredictionscronjob"]
