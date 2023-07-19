package probe

//func TestPortProbe(t *testing.T) {
//	CheckIfPortLive("49.232.214.202", 80)
//	//CheckIfPortLive("165.22.22.193", 80)
//	//CheckIfPortLive("109.95.151.226", 8095)
//}
//
//func TestAllPortProbe(t *testing.T) {
//	startTime := time.Now()
//	for port := 1; port <= 500; port++ {
//		if CheckIfPortLive("49.232.214.202", port) {
//			address := fmt.Sprintf("%s:%d", "49.232.214.202", port)
//			fmt.Println(address)
//		}
//	}
//	rtt := time.Since(startTime)
//	fmt.Printf("Round-Trip Time: %v\n", rtt)
//}
