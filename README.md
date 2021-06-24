# µBlockchain - DeFiDB Project

<img src="https://github.com/defidb/docs/blob/main/microblockchain.png">
<br/>

**Introduction**:<br/>
µBlockchain is a Golang based minimalistic blockchain project to demonstrate all functionalities of blockchain/crypto-assets use case.

**Usage**<br/>
```
----------------------------------------------------------------------------------------
Micro:Blockchain usage: 
----------------------------------------------------------------------------------------
getbalance -address ADDRESS - get balance for ADDRESS
createblockchain -address ADDRESS - creates a blockchain and rewards the mining fee for ADDRESS
mineblock -address ADDRESS - creates a block and rewards the mining fee for ADDRESS
printchain - prints the blocks in the chain
send -from FROM -to TO -amount AMOUNT - send AMOUNT of coins from one ADDRESS to another ADDRESS
createwallet - creates a new wallet
listaddresses - lists the addresses in the wallet file
----------------------------------------------------------------------------------------
Example: go run main.go -colors true --help (or go build and then ./microblockchain --help)
----------------------------------------------------------------------------------------
```
**Example**:<br/>
```shell
go build
```
Create genesis block and first wallet (reward: 100):<br/>
```shell
createblockchain -address 1JMQeu1tyCnErcWXjWuh4NfU95pVSqtZLC
# output
# badger 2021/06/24 08:43:40 INFO: All 0 tables opened in 0s
# 00002488975f5c22e3ad33ec5f8a3dc91e390f3fa92b68726b26ff31666d9729
# Genesis block created!
# badger 2021/06/24 08:43:41 DEBUG: Storing value log head: {Fid:0 Len:42 Offset:569}
# badger 2021/06/24 08:43:41 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
# badger 2021/06/24 08:43:41 INFO: Running for level: 0
# badger 2021/06/24 08:43:41 DEBUG: LOG Compact. Added 3 keys. Skipped 0 keys. Iteration took: 145.28µs
# badger 2021/06/24 08:43:41 DEBUG: Discard stats: map[]
# badger 2021/06/24 08:43:41 INFO: LOG Compact 0->1, del 1 tables, add 1 tables, took 46.044145ms
# badger 2021/06/24 08:43:41 INFO: Compaction for level: 0 DONE
# badger 2021/06/24 08:43:41 INFO: Force compaction on level 0 done
# Finished creating chain        
```
Create wallet #2:<br/>
```shell
./microblockchain createwallet
# output
# New address is: 1JMQeu1tyCnErcWXjWuh4NfU95pVSqtZLC
```
Mine a new block with reward 100:<br/>
```shell
./microblockchain mineblock -address 1CVFX4qaBwpti5wRGdJNDxCwxBbzUEWRKG
# output
# badger 2021/06/24 08:44:54 INFO: All 1 tables opened in 0s
# badger 2021/06/24 08:44:54 INFO: Replaying file id: 0 at offset: 611
# badger 2021/06/24 08:44:54 INFO: Replay took: 27.384µs
# badger 2021/06/24 08:44:54 DEBUG: Value log discard stats empty
# 000031a8165a609c802d607459a2c8c267b4ee80c4441fe1c903d070653e86e7
# Finished adding block for address 1CVFX4qaBwpti5wRGdJNDxCwxBbzUEWRKG
# badger 2021/06/24 08:44:58 DEBUG: Storing value log head: {Fid:0 Len:42 Offset:1228}
# badger 2021/06/24 08:44:58 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
# badger 2021/06/24 08:44:58 INFO: Running for level: 0
# badger 2021/06/24 08:44:58 DEBUG: LOG Compact. Added 6 keys. Skipped 0 keys. Iteration took: 372.882µs
# badger 2021/06/24 08:44:58 DEBUG: Discard stats: map[]
# badger 2021/06/24 08:44:58 INFO: LOG Compact 0->1, del 2 tables, add 1 tables, took 45.573949ms
# badger 2021/06/24 08:44:58 INFO: Compaction for level: 0 DONE
# badger 2021/06/24 08:44:58 INFO: Force compaction on level 0 done
```
Check balance of wallet #1:<br/>
```shell
./microblockchain getbalance -address 1JMQeu1tyCnErcWXjWuh4NfU95pVSqtZLC
# output
# badger 2021/06/24 10:04:23 INFO: All 1 tables opened in 0s
# badger 2021/06/24 10:04:23 INFO: Replaying file id: 0 at offset: 1270
# badger 2021/06/24 10:04:23 INFO: Replay took: 16.131µs
# badger 2021/06/24 10:04:23 DEBUG: Value log discard stats empty
# Balance of 1JMQeu1tyCnErcWXjWuh4NfU95pVSqtZLC: 100
# badger 2021/06/24 10:04:23 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
```
Check balance of wallet #2:<br/>
```shell
./microblockchain getbalance -address 1CVFX4qaBwpti5wRGdJNDxCwxBbzUEWRKG
# output
# badger 2021/06/24 10:04:39 INFO: All 1 tables opened in 1ms
# badger 2021/06/24 10:04:39 INFO: Replaying file id: 0 at offset: 1270
# badger 2021/06/24 10:04:39 INFO: Replay took: 6.944µs
# badger 2021/06/24 10:04:39 DEBUG: Value log discard stats empty
# Balance of 1CVFX4qaBwpti5wRGdJNDxCwxBbzUEWRKG: 100
# badger 2021/06/24 10:04:39 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
```
Send crypto from wallet #1 to wallet #2:<br/>
```shell
./microblockchain send -from 1JMQeu1tyCnErcWXjWuh4NfU95pVSqtZLC -to 1CVFX4qaBwpti5wRGdJNDxCwxBbzUEWRKG -amount 20 
# output   
# badger 2021/06/24 10:09:53 INFO: All 1 tables opened in 0s
# badger 2021/06/24 10:09:53 INFO: Replaying file id: 0 at offset: 1270
# badger 2021/06/24 10:09:53 INFO: Replay took: 22.517µs
# badger 2021/06/24 10:09:53 DEBUG: Value log discard stats empty
# 00001063b36a0dce4125017a468800d076897ac396606b38fafe205cf658466a
# Success!
# badger 2021/06/24 10:09:55 DEBUG: Storing value log head: {Fid:0 Len:42 Offset:1962}
# badger 2021/06/24 10:09:55 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
# badger 2021/06/24 10:09:55 INFO: Running for level: 0
# badger 2021/06/24 10:09:55 DEBUG: LOG Compact. Added 7 keys. Skipped 2 keys. Iteration took: 61.681µs
# badger 2021/06/24 10:09:55 DEBUG: Discard stats: map[0:64]
# badger 2021/06/24 10:09:55 INFO: LOG Compact 0->1, del 2 tables, add 1 tables, took 44.76387ms
# badger 2021/06/24 10:09:55 INFO: Compaction for level: 0 DONE
# badger 2021/06/24 10:09:55 INFO: Force compaction on level 0 done
```
Check balance of wallet #1:<br/>
```shell
./microblockchain getbalance -address 1JMQeu1tyCnErcWXjWuh4NfU95pVSqtZLC
# output      
# badger 2021/06/24 10:10:12 INFO: All 1 tables opened in 0s
# badger 2021/06/24 10:10:12 INFO: Replaying file id: 0 at offset: 2004
# badger 2021/06/24 10:10:12 INFO: Replay took: 7.396µs
# badger 2021/06/24 10:10:12 DEBUG: Value log discard stats empty
# Balance of 1JMQeu1tyCnErcWXjWuh4NfU95pVSqtZLC: 80
#  badger 2021/06/24 10:10:12 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
```
Check balance of wallet #1:<br/>
```shell
./microblockchain getbalance -address 1CVFX4qaBwpti5wRGdJNDxCwxBbzUEWRKG                                         
# output
# badger 2021/06/24 10:10:17 INFO: All 1 tables opened in 0s
# badger 2021/06/24 10:10:17 INFO: Replaying file id: 0 at offset: 2004
# badger 2021/06/24 10:10:17 INFO: Replay took: 12.511µs
# badger 2021/06/24 10:10:17 DEBUG: Value log discard stats empty
# Balance of 1CVFX4qaBwpti5wRGdJNDxCwxBbzUEWRKG: 120
# badger 2021/06/24 10:10:17 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
```
Print all blocks:<br/>
```shell
./microblockchain printchain
# output                                            
# badger 2021/06/24 10:10:34 INFO: All 1 tables opened in 0s
# badger 2021/06/24 10:10:34 INFO: Replaying file id: 0 at offset: 2004
# badger 2021/06/24 10:10:34 INFO: Replay took: 7.017µs
# badger 2021/06/24 10:10:34 DEBUG: Value log discard stats empty
# Previous hash: 000031a8165a609c802d607459a2c8c267b4ee80c4441fe1c903d070653e86e7
# Hash: 00001063b36a0dce4125017a468800d076897ac396606b38fafe205cf658466a
# Pow: true
# 
# Previous hash: 00002488975f5c22e3ad33ec5f8a3dc91e390f3fa92b68726b26ff31666d9729
# Hash: 000031a8165a609c802d607459a2c8c267b4ee80c4441fe1c903d070653e86e7
# Pow: true
# 
# Previous hash: 
# Hash: 00002488975f5c22e3ad33ec5f8a3dc91e390f3fa92b68726b26ff31666d9729
# Pow: true
# 
# badger 2021/06/24 10:10:34 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
```