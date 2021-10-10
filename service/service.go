package service

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"strconv"
	"wsTest/configs"
	"wsTest/model"
)

func BinanceWSService(cfg *configs.Config, errCh chan error)  {
	//connecting to binance stream api
	ws, wsErr := websocket.Dial(fmt.Sprintf("%v%v@%v", cfg.WsUrl, cfg.OrdersInBook, cfg.TimeRange), "", fmt.Sprintf("%v%v@%v", cfg.WsUrl, cfg.OrdersInBook, cfg.TimeRange))
	if wsErr != nil{
		log.Printf("dial error: %v", wsErr)
		errCh <- wsErr
	}

	for  {
		//model for binance messages
		msg := &model.Msg{}
		//unmarshalling received messages
		jsonErr := websocket.JSON.Receive(ws, msg)
		if jsonErr != nil{
			log.Printf("msg receive error: %v", jsonErr)
			errCh <- jsonErr
		}

		//counting total price and total quantity
		totalAsksPrice, totalAsksQuantity := totalCount(msg.Asks)
		totalBidsPrice, totalBidsQuantity := totalCount(msg.Bids)


		fmt.Println("order book id:",msg.LastUpdateId)
		fmt.Printf("Asks count: %v, Asks total price: %v, Asks total quantity: %v\n",len(msg.Asks), totalAsksPrice, totalAsksQuantity)
		fmt.Printf("Bids count: %v, Bids total price: %v, Bids total quantity: %v\n",len(msg.Bids), totalBidsPrice, totalBidsQuantity)

		for i, v := range msg.Asks{
			fmt.Printf("Ask #%v: price:%v, quantity:%v\n", i+1, v[0], v[1])
		}
		for i, v := range msg.Bids{
			fmt.Printf("Bid #%v: price:%v, quantity:%v\n", i+1, v[0], v[1])
		}
	}
}

func totalCount(arr []model.StrArr) (string, string)  {
	var totalPrice float64
	var totalQuantity float64

	for _, v := range arr{
		priceF, priceErr := strconv.ParseFloat(v[0], 64)
		if priceErr != nil{
			log.Fatalf("cannot parse price: %v", priceErr)
		}
		quantityF, quantityErr := strconv.ParseFloat(v[1], 64)
		if quantityErr != nil{
			log.Fatalf("cannot parse quantity: %v", quantityErr)
		}
		totalPrice += priceF
		totalQuantity += quantityF
	}

	return fmt.Sprintf("%f", totalPrice), fmt.Sprintf("%f", totalQuantity)
}
