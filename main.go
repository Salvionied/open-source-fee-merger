package main

import "open-source-fee-merger/merger"

func main() {
	merger := merger.NewMerger("wallet")
	merger.Loop()
}
