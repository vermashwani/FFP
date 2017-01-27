package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// FFP is a high level smart contract that FFPs together business artifact based smart contracts
type FFP struct {

}

// Application is for storing retreived Application

type Application struct{	
	RoaltyPoints string `json:"roaltyPoints"`
}


// Init initializes the smart contracts
func (t *FFP) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("ApplicationTable")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ApplicationTable", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "applicationId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "passportNumber", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "age", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "martialStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nationality", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "panNumber", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "aadharNumber", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "membershipType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "roaltyPoints", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
	return nil, nil
}



func (t *FFP) updateMembership(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	applicationId := args[0]
	newMembershipType := args[1]

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()



	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this applicationId
	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	//applicationId := row.Columns[0].GetString_()
	
	passportNumber := row.Columns[1].GetString_()
	status := row.Columns[2].GetString_()
	title := row.Columns[3].GetString_()
	firstName := row.Columns[4].GetString_()
	lastName := row.Columns[5].GetString_()
	gender := row.Columns[6].GetString_()
	dob := row.Columns[7].GetString_()
	age := row.Columns[8].GetString_()
	martialStatus := row.Columns[9].GetString_()
	nationality := row.Columns[10].GetString_()
	panNumber := row.Columns[11].GetString_()
	aadharNumber := row.Columns[12].GetString_()
	membershipType := newMembershipType
	roaltyPoints := row.Columns[14].GetString_()


	//Insert the row pertaining to this applicationId with new status
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: passportNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: membershipType}},
				&shim.Column{Value: &shim.Column_String_{String_: roaltyPoints}},
		}})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

}

func (t *FFP) addMile(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	applicationId := args[0]
	mileToAdd:=args[1]
	newRoyaltyToAdd, _ := strconv.ParseInt(mileToAdd, 10, 0)

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()



	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this applicationId
	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	
	earlierMile:=row.Columns[14].GetString_()
	earlierRoyalty, _:=strconv.ParseInt(earlierMile, 10, 0)
	newRoyaltyPoint:= strconv.Itoa(int(earlierRoyalty) + int(newRoyaltyToAdd))
	
	
	//applicationId := row.Columns[0].GetString_()
	
	passportNumber := row.Columns[1].GetString_()
	status := row.Columns[2].GetString_()
	title := row.Columns[3].GetString_()
	firstName := row.Columns[4].GetString_()
	lastName := row.Columns[5].GetString_()
	gender := row.Columns[6].GetString_()
	dob := row.Columns[7].GetString_()
	age := row.Columns[8].GetString_()
	martialStatus := row.Columns[9].GetString_()
	nationality := row.Columns[10].GetString_()
	panNumber := row.Columns[11].GetString_()
	aadharNumber := row.Columns[12].GetString_()
	membershipType := row.Columns[13].GetString_()
	roaltyPoints := newRoyaltyPoint


	//Insert the row pertaining to this applicationId with new status
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: passportNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: membershipType}},
				&shim.Column{Value: &shim.Column_String_{String_: roaltyPoints}},
		}})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

}


func (t *FFP) deductMile(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	applicationId := args[0]
	mileToDeduct:=args[1]
	newRoyaltyToDeduct, _ := strconv.ParseInt(mileToDeduct, 10, 0)

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()



	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this applicationId
	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	
	earlierMile:=row.Columns[14].GetString_()
	earlierRoyalty, _:=strconv.ParseInt(earlierMile, 10, 0)
	newRoyaltiPointtoTest:= earlierRoyalty - newRoyaltyToDeduct
	
	if newRoyaltiPointtoTest < 0 {
		return nil, errors.New("can't deduct as the resulting royalty becoming less than zero.")
	}
	newRoyaltyPoint:= strconv.Itoa(int(earlierRoyalty) - int(newRoyaltyToDeduct))
	
	
	//applicationId := row.Columns[0].GetString_()
	
	passportNumber := row.Columns[1].GetString_()
	status := row.Columns[2].GetString_()
	title := row.Columns[3].GetString_()
	firstName := row.Columns[4].GetString_()
	lastName := row.Columns[5].GetString_()
	gender := row.Columns[6].GetString_()
	dob := row.Columns[7].GetString_()
	age := row.Columns[8].GetString_()
	martialStatus := row.Columns[9].GetString_()
	nationality := row.Columns[10].GetString_()
	panNumber := row.Columns[11].GetString_()
	aadharNumber := row.Columns[12].GetString_()
	membershipType := row.Columns[13].GetString_()
	roaltyPoints := newRoyaltyPoint


	//Insert the row pertaining to this applicationId with new status
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: passportNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: membershipType}},
				&shim.Column{Value: &shim.Column_String_{String_: roaltyPoints}},
		}})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

}


func (t *FFP) getMile(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting applicationid to query")
	}

	applicationId := args[0]
	

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + applicationId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + applicationId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := Application{}
	
	res2E.RoaltyPoints = row.Columns[14].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


// Invoke invokes the chaincode
func (t *FFP) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "submitApplication" {
		if len(args) != 15 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 15. Got: %d.", len(args))
		}

		applicationId := args[0]
		passportNumber := args[1]
		status := args[2]
		title := args[3]
		firstName := args[4]
		lastName := args[5]
		gender := args[6]
		dob := args[7]
		age := args[8]
		martialStatus := args[9]
		nationality := args[10]
		panNumber := args[11]
		aadharNumber := args[12]
		membershipType := args[13]
		roaltyPoints := args[14]
		
		
		// Insert a row
		ok, err := stub.InsertRow("ApplicationTable", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: passportNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: membershipType}},
				&shim.Column{Value: &shim.Column_String_{String_: roaltyPoints}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		return nil, err
	} else if function == "updateMembership" { 
		t := FFP{}
		return t.updateMembership(stub, args)
	} else if function == "addMile" { 
		t := FFP{}
		return t.addMile(stub, args)
	} else if function == "deductMile" { 
		t := FFP{}
		return t.deductMile(stub, args)
	}

	return nil, errors.New("Invalid invoke function name.")

}

func (t *FFP) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getMile" {
		if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting applicationid to query")
		}
		t := FFP{}
		return t.getMile(stub, args)		
	}	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(FFP))
	if err != nil {
		fmt.Printf("Error starting FFP: %s", err)
	}
} 
