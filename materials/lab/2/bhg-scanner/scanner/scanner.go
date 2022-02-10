// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Explanation: This program scans for available ports. It will scan for ports by using the worker method to open and test the port, then it will close the port. It will
// maintain a list of open and closed ports as well.
// Usage: to test file run "go test" in scanner directory via terminal. to run program navigate to main directory via terminal and first run "go build". then run "./main".
// {TODO 1: FILL IN}

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)


func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)    
		conn, err := net.DialTimeout("tcp", address, 21 * time.Second) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		if err != nil { 
			results <- -1 * p
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object 
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner(numPorts int) (int, int) {  

	//TODO 3 : ADD closed ports; currently code only tracks open ports
	var openports []int  // notice the capitalization here. access limited!
	var closedports []int

	ports := make(chan int, 111)   // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= numPorts; i++ {
			ports <- i
		}
	}()

	for i := 0; i < numPorts; i++ {
		port := <-results
		if port > 0 {
			openports = append(openports, port)
		}else{
			closedports = append(closedports, -1 * port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)

	//TODO 5 : Enhance the output for easier consumption, include closed ports
	for _, port := range closedports {
		fmt.Printf("%d, closed\n", port)
	}
	for _, port := range openports {
		fmt.Printf("%d, open\n", port)
	}
	fmt.Printf("%d, open ports, %d closed ports\n", len(openports), len(closedports))
	return len(openports), len(closedports) // TODO 6 : Return total number of ports scanned (number open, number closed); 
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
