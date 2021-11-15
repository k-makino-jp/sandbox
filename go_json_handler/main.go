package main

func main() {
	rootCmd := NewConfigCmd()
	// rootCmd.createConfigDat()
	rootCmd.readConfigDat()
	rootCmd.listConfigDat()
}
