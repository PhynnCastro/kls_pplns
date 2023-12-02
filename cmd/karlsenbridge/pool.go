func connectToPool() {
  conn, err := net.Dial("tcp", config.Pool.Host+":"+config.Pool.Port)

  // handle connection  
}
