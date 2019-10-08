package intro19

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type subscriber struct {
	w  io.Writer
	ch chan []byte
}

func (s *subscriber) send(msg []byte) {
	select {
	case s.ch <- msg:
		// message delivered
	case <-time.After(time.Second):
		// could not deliver message within one second
	}

}

func (s *subscriber) worker(id string, flusher http.Flusher) {
	fmt.Fprintf(s.w, "subscriber %s connected\n", id)
	flusher.Flush()
	defer close(s.ch)

	for msg := range s.ch {
		if _, err := s.w.Write(append(msg, '\n')); err != nil {
			break
		}
		flusher.Flush()
	}
}

type pubSubController struct {
	subMutex    sync.RWMutex
	subscribers map[string]*subscriber
	msgChan     chan string
}

func (c *pubSubController) subscriberHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)

	f, isFlusher := rw.(http.Flusher)
	if !isFlusher {
		http.Error(rw, "remote end does not support flusher interface", http.StatusNotImplemented)
		return
	}

	c.subMutex.Lock()
	s := &subscriber{
		w:  rw,
		ch: make(chan []byte, 1),
	}
	c.subscribers[r.RemoteAddr] = s
	log.Printf("subscriber %s connected\n", r.RemoteAddr)
	c.subMutex.Unlock()

	s.worker(r.RemoteAddr, f)

	log.Printf("subscriber %s disconnected\n", r.RemoteAddr)
	c.subMutex.Lock()
	defer c.subMutex.Unlock()
	delete(c.subscribers, r.RemoteAddr)
	log.Printf("subscriber %s removed\n", r.RemoteAddr)
}

func (c *pubSubController) publishHandler(rw http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "failed to read message: "+err.Error(), http.StatusInternalServerError)
		return
	} else if len(msg) == 0 {
		http.Error(rw, "message is empty", http.StatusBadRequest)
		return
	}

	c.subMutex.RLock()
	defer c.subMutex.RUnlock()
	for _, sub := range c.subscribers {
		go sub.send(msg)
	}
	fmt.Fprintf(rw, "message delivered to %d subscribers", len(c.subscribers))
	log.Printf("delivered message '%s' to %d subscribvers\n", msg, len(c.subscribers))
}

func newPubSubController() *pubSubController {
	return &pubSubController{
		subscribers: make(map[string]*subscriber),
	}
}
