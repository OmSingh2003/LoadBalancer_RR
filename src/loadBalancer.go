package src
import (
	""
)
type loadBalancer struct {
	port string  
	roundRobinCount int 
	servers []Servers 
}
func new loadBalancer( port string ,  servers []Servers) {
	 return &loadBalancer{
		roundRobinCount : 0
		 port : port 
		servers : servers
	}
}
