package main

import "time"

// Edge devices
type edgeDevice interface {
	getInput() string
}

type keyboard struct{}
type touchscreen struct{}

var edgeDeviceTypeMap = map[string]edgeDevice{
	"keyboard":    &keyboard{},
	"touchscreen": &touchscreen{},
}

func createdgeDevice(deviceType string) edgeDevice {
	return edgeDeviceTypeMap[deviceType]
}

func (k keyboard) getInput() string {
	// Get input from keyboard
	return input
}

func (t touchscreen) getInput() string {
	// Get input from touchscreen
	return input
}

// Authentication
type authMethod interface {
	authenticate(string) bool
}

type authCredential struct {
	authMethod	  string
	AccountNumber string
	KeyHash       []byte
}

type pin struct{}
type faceId struct{}
type fingerprintScanner struct{}

var authMethodTypeMap = map[string]authMethod{
	"pin":    &pin{},
	"faceId": &faceId{},
	"fingerprintScanner": &fingerprintScanner{},
}

func CreateAuthMethod(authMethod string) authMethod {
	return authMethodTypeMap[authMethod]
}

func (f fingerprintScanner) authenticate(accountNumber) bool {
}

func (f faceId) authenticate(accountNumber) bool {
}

func (p pin) authenticate(accountNumber) bool {
	// 1. User account number is received from input
	// 2. Get encrypted key of the account - get_key_hash(accountNumber)
	// 3. Encrypt inserted key with sha256
	// 4. Match and get result
	// 5. return depending on result
}

func (a authCredential) get_key_hash(accountNumber, authMethod) []byte {
	// Get the KeyHash belonging to the given account for provided authMethodfrom DB
	return key_hash
}

// Operations
type operation interface {
	perform(string) transaction
}

type withdraw struct{}
type deposit struct{}
type transfer struct{}
type balance struct{}

var operationTypeMap = map[string]operation{
	"withdraw": &withdraw{},
	"deposit":  &deposit{},
	"transfer": &transfer{},
	"balance":  &balance{},
}

func CreateOperation(operationType string) operation {
	return operationTypeMap[operationType]
}

func (w withdraw) perform(accountNumber) transaction {
	// Get transactable amount from user
	// Get account balance get_account_balance(accountNumber)
	// Check if account balance is greater than transactable amount
	// If yes, perform withdraw operation
	// If no, return error
	// Return transaction

	return transaction
}

func (t transfer) perform(accountNumber) transaction {
	// Get transactable amount from user
	// Get destination account number from user
	// Get account balance get_account_balance(accountNumber)
	// Check if account balance is greater than transactable amount
	// If yes, perform transfer operation

	return transaction
}

func (b balance) perform(accountNumber) transaction {
	// Get account balance get_account_balance(accountNumber)
	return transaction
}

func (d deposit) perform(accountNumber) transaction {

	return transaction
}

func get_account_balance(accountNumber) float64 {
	// Get account balance by querying the DB for transactions
	return balance
}

// Transaction
type transaction struct {
	Amount        float64
	Operation     string
	FromAccount   string
	ToAccount     string
	TransactionTime time.Time
}

func storeTransaction(transaction) {
	// Store transaction in DB
}

func main() {
	deviceType := // Get device type based on user input or initiation
	edgeDevice := createdgeDevice(deviceType)

	// User decides on authentication method
	authenticationMethod := edgeDevice.getInput()
	authMethod := CreateAuthMethod(authenticationMethod)

	// User inserts account number
	accountNumber := edgeDevice.getInput()
	// User is authenticated based on the authentication method
	isAuthenticated := authMethod.authenticate(accountNumber)

	if !isAuthenticated {
		// Return error message indicating authentication failure
		return
	}

	// User decides on operation type
	operationType := edgeDevice.getInput()
	operation := CreateOperation(operationType)

	// Perform the selected operation
	trans := operation.perform(accountNumber)

	// Store the transaction in the database
	storeTransaction(trans)

	// Show the transaction result to the user
	showResult(trans)
}
