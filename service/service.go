package service

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
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

		fmt.Println("order book id:",msg.LastUpdateId)
		fmt.Println("Asks count:",len(msg.Asks))
		fmt.Println("Bids count:",len(msg.Bids))
		for i, v := range msg.Asks{
			fmt.Printf("Ask #%v: price:%v, quantity:%v\n", i+1, v[0], v[1])
		}
		for i, v := range msg.Bids{
			fmt.Printf("Bid #%v: price:%v, quantity:%v\n", i+1, v[0], v[1])
		}
	}
}
