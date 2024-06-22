package main

//import (
//	"os"
//	"sync"
//)
//
//Data model to obtain coupon for a purchase.
//
//type HashKey string
//type Value interface
//
//
//
//
//	app / main.go -> boot
//
//app struct
//{
//hash *HashTable
//}
//// init of services :=
//app.hash = newHashTable
//
//resp
//1. GET "key"         nil, 4xx
//2. SET "key" "val"   _ , 2xx
//3. GET "key"        "val", 2xx
//4. SET "key" "val2" _, 2xx
//5. GET "key"        "val2", 2xx
//
//
//type Peers []string {"ips","ips","ips"}
//
//type HashTable map[HashKey]Value
//
//func newHashTable() {
//	ht := make(HashTable)
//	return &ht
//}
//
//func (ht *HashTable) Get(hashKey string) Value, bool {
//
//val, ok:= ht[hashKey]
//if !ok {
//fileLocation := getFileLocation(hashKey)
//file, nok := os.ReadFile(fileLocation)
//if nok {
//
//}
//
//val := ht.callPeerGet(hashKey)
//return file.String(), true
//}
//}
//
//func (ht *HashTable) Set(hashKey, value string)  {
//// check basis of key
//	// common table or ledger -> right peer -> ( offload -> that )
//// ledger Consensus with algorithm
//
//	res := ht.CallPeerSet(hashKey, value, true) // -> sqs.publish(PeerSet(hashKey, value))
//	if res {
//		ht[hashKey] = value
//		fileLocation := getFileLocation(hashKey //
//		file, ok := os.WriteFile(fileLocation, value) // overriding the file master
//	}
//
//	return 5xx
//
//}
//
//
//
//// future scope :
//// 1. use same file for multiple keys , pos in file can be determined using the hashing
//// 2. synch the ledger with the hashkey and peerlist
//// 3. making all the async requests via sqs, check resp using channels
//// 4. check consensus
////
//
//
//
//
//// CallPeerGet should call the peers and get the values for that hashKey in those peers, depending on
//// consensusScore return back the one with highest score
//func (ht *HashTable) CallPeerGet(hashKey, value string)  {
////wg
//	valChan := make(chan int)
//	for _,ip: range peers {
//		go http.Get(ip+"/peer/get")
//	}
//
//	//consensusScore map[val]int
//   //
//   //val with highest consensusScore
//   //
//	ht[hashKey] = value
//	fileLocation := getFileLocation(hashKey)
//	file, ok := os.WriteFile(fileLocation, value) // overriding the file
//}1
//
//// CallPeerSet should call the peers and set the values for that hashKey in those peers, depending on
//// whether it should wait it can return back or wait for n/2 +1 responses and return
//func (ht *HashTable) CallPeerSet(hashKey, value string, shouldWait bool)  {
//	if shouldWait {
//		// check for peer set success score and then return set here and then return val
//		// get back len(peers)/2 +1 responses
//      var wg sync.WaitGroup
//		wg.Add(len(peers)/2 +1)
//		for _,ip: range peers {
//			go http.Post(ip+"/peer/set", PeerSet(hashKey, value))
//		}
//		wg.Wait()
//		// send back the response
//		return
//	}
//	for _,ip: range peers {
//		go http.Post(ip+"/peer/set", PeerSet(hashKey, value))
//	}
//	ht[hashKey] = value
//	fileLocation := getFileLocation(hashKey)
//	file, ok := os.WriteFile(fileLocation, value) // overriding the file
//}
//
//// PeerGet should get the value for the hashKey from the peer and return the value, comes from another node
//func (ht *HashTable) PeerGet(hashKey, value string)  {
//	// same as peerSet
//	ht[hashKey] = value
//	fileLocation := getFileLocation(hashKey)
//	file, ok := os.WriteFile(fileLocation, value) // overriding the file
//}
//
//// PeerSet should set the value for the hashKey from the peer and return the value, it comes from another node
//func (ht *HashTable) PeerSet(hashKey, value string)  {
//	ht[hashKey] = value
//	fileLocation := getFileLocation(hashKey)
//	file, ok := os.WriteFile(fileLocation, value) // overriding the file
//}
//
//
//
//
//http.Get('/get', GetController)
//http.Post('/set', SetController)
//
//
//GetController -> takes the query params -> passes it to the service which calls app.Hash.Get
//SetController -> takes the req body-> passes it to the service which calls app.Hash.Set
//
//
//
//
//
//
//
//
//
//
//
//type Coupon struct {
//	type CouponTitle string
//	type CouponState string // created, published, active, expired
//	type CouponType string // AmountOffOrder BxGy PercOffOrder etc.
//	type ExpiresAt *time.Time
//	type CreatedAt time.Time
//	type ActiveFrom time.Time // Coupons in draft, -> published , active
//	type CustomerWhitelist CustomerWhiteList
//	type CouponRemainingUsage int
//	type DeletedAt *time.Time
//	type Amount Amount
//	type OfferId string // conditions for redemption are contained within this
//}
//
//
//
//
//
//
//
//type CouponRedemptionRequest struct {
//	type CouponTitle string
//	type CouponType string // AmountOffOrder BxGy PercOffOrder etc.
//	type ExpiresAt time.Time
//	type CreatedAt time.Time
//	type ActiveFrom time.Time // Coupons in draft, -> published , active
//	type CustomerWhitelist CustomerWhiteList
//	type Amount Amount
//	type Customer Customer
//}
//
//type CouponRedemptionResponse struct {
//	type CouponTitle string
//	type Success bool
//	type CouponRemainingUsage int
//}
//
//type Customer struct {
//	type Name string
//	type Email string
//	type Mob string
//}
//
//type Amount struct {
//	type OfferType string // Perc/Fixed
//	type Amount int
//}
//
//type CustomerWhiteList struct {
//	type CustomerWhitelistFile string // file name for a csv file , +mobile numbers /emails
//}
