package main

func main() {
	rootCmd := NewConfigCmd()
	rootCmd.createConfigDat()
	rootCmd.listConfigDat()
}
