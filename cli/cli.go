package cli

import (
	"flag"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"microblockchain/common"
	"os"
	"runtime"
	"strconv"

	"microblockchain/blockchain"
	"microblockchain/blockchain/wallet"
)

type CommandLine struct{}

//printUsage will display what options are availble to the user
func (cli *CommandLine) printUsage() {
	fmt.Println(Green("----------------------------------------------------------------------------------------"))
	fmt.Println(Green("Micro:Blockchain usage: "))
	fmt.Println(Green("----------------------------------------------------------------------------------------"))
	fmt.Println(BrightGreen("getbalance -address ADDRESS"), Green("- get balance for ADDRESS"))
	fmt.Println(BrightGreen("createblockchain -address ADDRESS"), Green("- creates a blockchain and rewards the mining fee for ADDRESS"))
	fmt.Println(BrightGreen("mineblock -address ADDRESS"), Green("- creates a block and rewards the mining fee for ADDRESS"))
	fmt.Println(BrightGreen("printchain"), Green("- prints the blocks in the chain"))
	fmt.Println(BrightGreen("send -from FROM -to TO -amount AMOUNT"), Green("- send AMOUNT of coins from one ADDRESS to another ADDRESS"))
	fmt.Println(BrightGreen("createwallet"), Green("- creates a new wallet"))
	fmt.Println(BrightGreen("listaddresses"), Green("- lists the addresses in the wallet file"))
	fmt.Println(Green("----------------------------------------------------------------------------------------"))
	fmt.Println(Green("Example: <go run main.go/microblockchain> -colors true --help"))
	fmt.Println(Green("----------------------------------------------------------------------------------------"))
}

//validateArgs ensures the cli was given valid input
func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		//go exit will exit the application by shutting down the goroutine
		// if you were to use os.exit you might corrupt the data
		runtime.Goexit()
	}
}

//printChain will display the entire contents of the blockchain
func (cli *CommandLine) printChain() {
	chain := blockchain.ContinueBlockChain("")
	defer chain.Database.Close()
	iterator := chain.Iterator()

	for {
		block := iterator.Next()
		fmt.Print(Red(fmt.Sprintf("Previous hash: %x\n", block.PrevHash)))
		fmt.Print(Red(fmt.Sprintf("Hash: %x\n", block.Hash)))
		pow := blockchain.NewProofOfWork(block)
		fmt.Print(Red(fmt.Sprintf("Pow: %s\n", strconv.FormatBool(pow.Validate()))))
		fmt.Println()
		// This works because the Genesis block has no PrevHash to point to.
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

//listAddresses will list all addresses in the wallet file
func (cli *CommandLine) listAddresses() {
	wallets, _ := wallet.CreateWallets()
	addresses := wallets.GetAllAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}

}

//createWallet will create a wallet in the wallet file
func (cli *CommandLine) createWallet() {
	wallets, _ := wallet.CreateWallets()
	address := wallets.AddWallet()
	wallets.SaveFile()

	fmt.Print(Red(fmt.Sprintf("New address is: %s\n", address)))

}

//Creates a blockchain and awards address the coinbase
func (cli *CommandLine) createBlockChain(address string) {
	newChain := blockchain.InitBlockChain(address)
	newChain.Database.Close()
	fmt.Print(Red(fmt.Sprintf("Finished creating chain")))
}

func (cli *CommandLine) mineBlock(address string) {

	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()
	tx := blockchain.CoinbaseTx(address, "")
	chain.AddBlock([]*blockchain.Transaction{tx})

	fmt.Print(Red(fmt.Sprintf("Finished adding block for address %v\n", address)))
}

func (cli *CommandLine) getBalance(address string) {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Print(Red(fmt.Sprintf("Balance of %s: %d\n", address, balance)))
}

func (cli *CommandLine) send(from, to string, amount int) {
	chain := blockchain.ContinueBlockChain(from)
	defer chain.Database.Close()

	tx := blockchain.NewTransaction(from, to, amount, chain)

	chain.AddBlock([]*blockchain.Transaction{tx})
	fmt.Println("Success!")

}

//Run will start up the command line
func (cli *CommandLine) Run() {
	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	mineBlockCmd := flag.NewFlagSet("mineblock", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	mineBlock := mineBlockCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			common.Logr.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			common.Logr.Panic(err)
		}
	case "mineblock":
		err := mineBlockCmd.Parse(os.Args[2:])
		if err != nil {
			common.Logr.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			common.Logr.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			common.Logr.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			common.Logr.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			common.Logr.Panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			runtime.Goexit()
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			runtime.Goexit()
		}
		cli.createBlockChain(*createBlockchainAddress)
	}

	if mineBlockCmd.Parsed() {
		cli.mineBlock(*mineBlock)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			runtime.Goexit()
		}

		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
	if listAddressesCmd.Parsed() {
		cli.listAddresses()
	}
	if createWalletCmd.Parsed() {
		cli.createWallet()
	}
}
