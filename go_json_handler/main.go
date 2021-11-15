package main

func main() {
	rootCmd := NewRootCmd()
	rootCmd.createConfigDat()
	rootCmd.listConfigDat()
}
