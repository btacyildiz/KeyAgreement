package main

import (
    "fmt"
    "math"
    "bytes"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

func getPublicVariable(stub shim.ChaincodeStubInterface, key string) (int, error){
        value, err := stub.GetState(key)
        if err != nil{
                return -1, fmt.Errorf("Failed to get asset: %s with error: %s", key, err)
        }

        // convert to integer 
        intValue, err := strconv.Atoi(string(value)) 
        if err != nil{
                return -1, fmt.Errorf("Failed to convert to integer asset: %s with error: %s", value, err)
        }

        return intValue, nil
}

// Init is called during chaincode instantiation to initialize any data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {

 // Get the args from the transaction proposal
 // number of arguments should be equal to 3
 args := stub.GetStringArgs()
 if len(args) != 4{
   return shim.Error("Incorrect arguments. Expecting value p, q and a")
 }

p, err := strconv.Atoi(args[0])
q, err := strconv.Atoi(args[1])
a, err := strconv.Atoi(args[2])
participantCount, err := strconv.Atoi(args[3])

fmt.Printf("Initial Key Agreement Values Are...\n")
fmt.Printf("p: %d q: %d a: %d participantCount:%d\n", p, q, a, participantCount)

var res = int(math.Pow(float64(a),float64(q)))

// check if values fulfill the requirement of a^q mod p = 1
if (res % p != 1) {
        fmt.Printf("Moduation error: %d\n", res)
        return shim.Error(fmt.Sprintf("Failed to fulfill the requirement of keyagreement initial values."))
}

// set p 
err = stub.PutState("p", []byte(args[0]))
  if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset p"))
  }

// set q
err = stub.PutState("q", []byte(args[1]))
  if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset q"))
  }

  // set a
err = stub.PutState("a", []byte(args[2]))
if err != nil {
  return shim.Error(fmt.Sprintf("Failed to create asset a"))
}

// set participant count 
err = stub.PutState("participantCount", []byte(args[3]))
if err != nil {
  return shim.Error(fmt.Sprintf("Failed to create asset participantCount"))
}
  return shim.Success(nil)

}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    // Extract the function and args from the transaction proposal
        fn, args := stub.GetFunctionAndParameters()
	var result string
        var err error
        fmt.Printf("Invoke function: %s called\n", fn)
        if fn == "set" {
                result, err = set(stub, args)
        } else if fn == "get"{
                result, err = get(stub, args)
        } else if fn == "getAll"{
                result, err = getAll(stub, args)
        } else if fn == "setZ"{
                result, err = setZ(stub, args)
        } else if fn == "getZ"{
                result, err = getZ(stub, args)
        }else if fn == "setX"{
                result, err = setX(stub, args)
        }else if fn == "getX" {
                result, err = getX(stub, args)
        }

        if err != nil {
                return shim.Error(err.Error())
        }
        // Return the result as success payload
        return shim.Success([]byte(result))
}

/*SET AND VERIFY PARTICIPANT PUBLIC VARIABLE: Zj && Zj ^q  mod p == 1*/ 
func setZ(stub shim.ChaincodeStubInterface, args []string) (string, error){
        if len(args) != 2 {
                return "", fmt.Errorf("Incorrect arguments. Expecting a key")
        }

        participantID := args[0]

        z, err := strconv.Atoi(args[1])
        if err != nil{
                return "", fmt.Errorf("Failed to set asset (converting to int): %s", args[0])
        }

        // get public variables
        p, err := getPublicVariable(stub, "p")
        if err != nil{
                return "", fmt.Errorf("Failed get public variable p error: %s", err)
        }

        q, err := getPublicVariable(stub, "q")
        if err != nil{
                return "", fmt.Errorf("Failed get public variable q error: %s", err)
        }

        // verify the variable
        res := int(math.Pow(float64(z), float64(q)))
        if (res % p != 1){
                fmt.Printf("Moduation error: %d\n", res)
                return "", fmt.Errorf("SetZ Modulation error")
        }
        var buffer bytes.Buffer
        buffer.WriteString("Z")
        buffer.WriteString(participantID) //for example z1 z2 etc ...
        
        // set variable for participant
        err = stub.PutState(buffer.String(), []byte(args[1]))
        if err != nil {
                return "", fmt.Errorf("Failed to set asset: %s", args[0])
        }
        return args[1], nil
}

func getZ(stub shim.ChaincodeStubInterface, args []string)(string, error){
                // get participant count
        
        participantCount, err := getPublicVariable(stub, "participantCount")
        if err != nil{
                return "", fmt.Errorf("Failed get public variable participantCount error: %s", err)
        }

        start, end := "Z1", "Z" + strconv.Itoa(participantCount+1)

        return getValues(stub, start, end)
}

/*SET Xi Variable, Assign for each participant*/
func setX(stub shim.ChaincodeStubInterface, args []string) (string, error){
        if len(args) != 3 {
                return "", fmt.Errorf("Incorrect arguments. Expecting a key")
        }

        participantID := args[0]

        var bufferX bytes.Buffer
        bufferX.WriteString("X")
        bufferX.WriteString(participantID) 

        var bufferY bytes.Buffer
        bufferY.WriteString("Y")
        bufferY.WriteString(participantID) 

        
        // set variable for participant
        err := stub.PutState(bufferX.String(), []byte(args[1]))
        if err != nil {
                return "", fmt.Errorf("Failed to set asset: %s", args[0])
        }

        // set variable for participant
        err = stub.PutState(bufferY.String(), []byte(args[2]))
        if err != nil {
                return "", fmt.Errorf("Failed to set asset: %s", args[0])
        }
        return args[1], nil
}

func getX(stub shim.ChaincodeStubInterface, args []string)(string, error){
        // get participant count
        
        participantCount, err := getPublicVariable(stub, "participantCount")
        if err != nil{
                return "", fmt.Errorf("Failed get public variable participantCount error: %s", err)
        }

        start, end := "X1", "Y" + strconv.Itoa(participantCount+1)

        return getValues(stub, start, end)
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 2 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
    }

    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
            return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }
    return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 1 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key")
    }
    value, err := stub.GetState(args[0])
    if err != nil {
            return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if value == nil {
            return "", fmt.Errorf("Asset not found: %s", args[0])
    }
    return string(value), nil
}

func getValues(stub shim.ChaincodeStubInterface, startKey string, endKey string)(string, error){
        fmt.Printf("StartKey: %s EndKey: %s\n", startKey, endKey)
        resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return "", fmt.Errorf("GetStateByRange err")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("{")
        
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return "", fmt.Errorf("Iterator next err")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("\""+ queryResponse.Key +"\":")
		buffer.WriteString("\"")
                buffer.WriteString(string(queryResponse.Value))
                buffer.WriteString("\"")
                bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("}")

	fmt.Printf("- queryAllAssets:\n%s\n", buffer.String())
	return buffer.String(), nil 
}


func getAll(stub shim.ChaincodeStubInterface, args []string)(string, error){
        startKey := ""
        endKey := ""
        return getValues(stub, startKey, endKey)
}

// main function starts up the chaincode in the container during instantiate
func main() {
    if err := shim.Start(new(SimpleAsset)); err != nil {
            fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
    }
}