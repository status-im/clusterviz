package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/divan/graph-experiments/graph"
	"github.com/divan/graph-experiments/layout"
	"github.com/gorilla/websocket"
)

type WSServer struct {
	upgrader websocket.Upgrader
	hub      []*websocket.Conn

	Positions []*position
	layout    layout.Layout
	graph     *graph.Graph

	stats   *Stats
	fetcher *Fetcher
}

func NewWSServer(f *Fetcher, updateInterval time.Duration) *WSServer {
	ws := &WSServer{
		upgrader: websocket.Upgrader{},
		stats:    &Stats{},
		fetcher:  f,
	}
	go func() {
		t := time.NewTicker(updateInterval)
		for range t.C {
			ctx, _ := context.WithTimeout(context.Background(), updateInterval)
			if err := ws.refresh(ctx); err != nil {
				// TODO: send error to ws
				continue
			}
			ws.stats.Update(ws.graph)
			ws.broadcastStats()
		}
	}()
	return ws
}

type WSResponse struct {
	Type      MsgType         `json:"type"`
	Positions []*position     `json:"positions,omitempty"`
	Graph     json.RawMessage `json:"graph,omitempty"`
	Stats     Stats           `json:"stats,omitempty"`
}

type WSRequest struct {
	Cmd WSCommand `json:"cmd"`
}

type MsgType string
type WSCommand string

// WebSocket response types
const (
	RespPositions MsgType = "positions"
	RespGraph     MsgType = "graph"
	RespStats     MsgType = "stats"
)

// WebSocket commands
const (
	CmdInit    WSCommand = "init"
	CmdRefresh WSCommand = "refresh"
	CmdStats   WSCommand = "stats"
)

func (ws *WSServer) Handle(w http.ResponseWriter, r *http.Request) {
	c, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	ws.hub = append(ws.hub, c)

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", mt, err)
			break
		}
		ws.processRequest(c, mt, message)
	}
}

func (ws *WSServer) processRequest(c *websocket.Conn, mtype int, data []byte) {
	var cmd WSRequest
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		log.Fatal("unmarshal command", err)
		return
	}

	switch cmd.Cmd {
	case CmdInit:
		ws.sendGraphData(c)
		ws.updatePositions()
		ws.sendPositions(c)
	case CmdRefresh:
		ws.refresh(context.TODO())
		ws.broadcastStats()
	case CmdStats:
		ws.sendStats(c)
	}
}

func (ws *WSServer) sendMsg(c *websocket.Conn, msg *WSResponse) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Println("write:", err)
		return
	}

	err = c.WriteMessage(1, data)
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func (ws *WSServer) refresh(ctx context.Context) error {
	log.Println("Getting peers from Status-cluster")
	g, err := BuildGraph(ctx, ws.fetcher)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch: %s", err)
		return err
	}
	log.Printf("Loaded graph: %d nodes, %d links\n", len(g.Nodes()), len(g.Links()))

	log.Printf("Initializing layout...")
	repelling := layout.NewGravityForce(-50.0, layout.BarneHutMethod)
	springs := layout.NewSpringForce(0.01, 5.0, layout.ForEachLink)
	drag := layout.NewDragForce(0.8, layout.ForEachNode)
	l := layout.New(g, repelling, springs, drag)

	l.CalculateN(400)

	ws.updateGraph(g, l)
	ws.updatePositions()

	return nil
}
