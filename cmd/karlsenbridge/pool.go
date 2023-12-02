func connectToPool() {
  conn, err := net.Dial("tcp", config.Pool.Host+":"+config.Pool.Port)
  // handle connection  
}

func submitShare(share string) {
  _, err := conn.Write([]byte(share))
  // handle response
}
