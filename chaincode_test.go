package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// SKH is a high level smart contract that
type SKH struct {

}

// EmployeeDetails is for storing Person Details

type EmployeeDetails struct{	
	EmployeeId string `json:"employeeId"`
	Title string `json:"title"`
	Gender string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Dob string `json:"dob"`
	Email string `json:"email"`
	Country string `json:"country"`
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
	CreatedBy string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
}

type QualificationDetails struct{	
	QualificationId string `json:"qualificationId"`
	EmployeeId string `json:"employeeId"`
	NameExam string `json:"nameExam"`
	NameBoardUniversity string `json:"nameBoardUniversity"`
	YearPassing string `json:"yearPassing"`
	MarksPercentage string `json:"marksPercentage"`
	CreatedBy string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
}

type ExperienceDetails struct{	
	ExperienceId string `json:"experienceId"`
	EmployeeId string `json:"employeeId"`
	NameOrganization string `json:"nameOrganization"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	CurrentPosition string `json:"currentPosition"`
	CreatedBy string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
}

type CertificationDetails struct{	
	CertificateId string `json:"certificateId"`
	EmployeeId string `json:"employeeId"`
	NameCertificate string `json:"nameCertificate"`
	CertificationDate string `json:"certificationDate"`
	CertAuthority string `json:"certAuthority"`
	CreatedBy string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
}
// Transaction is for storing transaction Details

type Transaction struct{	
	TrxId string `json:"trxId"`
	TimeStamp string `json:"timeStamp"`
	PersonId string `json:"PersonId"`
	Source string `json:"source"`
	Skill string `json:"skill"`
	Trxntype string `json:"trxntype"`
	TrxnSubType string `json:"trxnSubType"`
	Remarks string `json:"remarks"`
}

// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}
	



// Init initializes the smart contracts
func (t *SKH) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("EmployeeDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("EmployeeDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "employeeId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdDate", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating EmployeeDetails.")
	}

	// Check if QualificationDetails table already exists
	_, err = stub.GetTable("QualificationDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create QualificationDetails Table
	err = stub.CreateTable("QualificationDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "qualificationId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "employeeId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nameExam", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nameBoardUniversity", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "yearPassing", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "marksPercentage", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdDate", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating QualificationDetails.")
	}
	
	// Check if ExperienceDetails table already exists
	_, err = stub.GetTable("ExperienceDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	
	// Create ExperienceDetails Table
	err = stub.CreateTable("ExperienceDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "experienceId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "employeeId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nameOrganization", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "startDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "endDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "currentPosition", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdDate", Type: shim.ColumnDefinition_STRING, Key: false},
	})
		if err != nil {
		return nil, errors.New("Failed creating ExperienceDetails.")
	}

	// Check if CertificationDetails table already exists
	_, err = stub.GetTable("CertificationDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	// Create CertificationDetails Table
	err = stub.CreateTable("CertificationDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "certificateId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "employeeId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nameCertificate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "certificationDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "certAuthority", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdDate", Type: shim.ColumnDefinition_STRING, Key: false},
	})
		})
		if err != nil {
		return nil, errors.New("Failed creating CertificationDetails.")
	}

	// Check if table already exists
	_, err = stub.GetTable("Transaction")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	
	// Create application Table
	err = stub.CreateTable("Transaction", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "trxId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "timeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PersonId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "skill", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxntype", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxnSubType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "remarks", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}

	// setting up the users role
	stub.PutState("user_type1_1", []byte("University"))
	stub.PutState("user_type1_2", []byte("Institute"))
	stub.PutState("user_type1_3", []byte("Organization1"))
	stub.PutState("user_type1_4", []byte("Organization2"))	
	
	return nil, nil
}
	

	
//registerEmployee to register a person
func (t *SKH) registerEmployee(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 13 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 13 . Got: %d.", len(args))
		}
		
		PersonId:=args[0]
		title:=args[1]
		gender:=args[2]
		firstName:=args[3]
		lastName:=args[4]
		dob:=args[5]
		email:=args[6]
		country:=args[7]
		address:=args[8]
		city:=args[9]
		zip:=args[10]
		
		/*assignerOrg1, err := stub.GetState(args[11])
		assignerOrg := string(assignerOrg1)
		createdBy:=assignerOrg*/
		
		createdBy:=args[11]
		createdDate:=args[12]


		// Insert a row
		ok, err := stub.InsertRow("EmployeeDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: employeeId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: createdDate}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

//addQualification - Add Qualification of employee
func (t *SKH) addQualification(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 8 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 8 . Got: %d.", len(args))
		}
		qualificationId:=args[0]
		employeeId:=args[1]
		nameExam:=args[2]
		nameBoardUniversity:=args[3]
		yearPassing:=args[4]
		marksPercentage:=args[5]
		createdBy:=args[6]
		createdDate:=args[7]

		// Insert a row
		ok, err := stub.InsertRow("QualificationDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: qualificationId}},
				&shim.Column{Value: &shim.Column_String_{String_: employeeId}},
				&shim.Column{Value: &shim.Column_String_{String_: nameExam}},
				&shim.Column{Value: &shim.Column_String_{String_: nameBoardUniversity}},
				&shim.Column{Value: &shim.Column_String_{String_: yearPassing}},
				&shim.Column{Value: &shim.Column_String_{String_: marksPercentage}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: createdDate}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		return nil, nil
}

//addExperience - Add Experience of employee
func (t *SKH) addExperience(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 8 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 8 . Got: %d.", len(args))
		}
		experienceId:=args[0]
		employeeId:=args[1]
		nameOrganization:=args[2]
		startDate:=args[3]
		endDate:=args[4]
		currentPosition:=args[5]
		createdBy:=args[6]
		createdDate:=args[7]

		// Insert a row
		ok, err := stub.InsertRow("ExperienceDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: experienceId}},
				&shim.Column{Value: &shim.Column_String_{String_: employeeId}},
				&shim.Column{Value: &shim.Column_String_{String_: nameOrganization}},
				&shim.Column{Value: &shim.Column_String_{String_: startDate}},
				&shim.Column{Value: &shim.Column_String_{String_: endDate}},
				&shim.Column{Value: &shim.Column_String_{String_: currentPosition}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: createdDate}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		return nil, nil
}

//addCertification - Add Certification of employee
func (t *SKH) addCertification(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 8 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 7 . Got: %d.", len(args))
		}
		certificateId:=args[0]
		employeeId:=args[1]
		nameCertificate:=args[2]
		certificationDate:=args[3]
		certAuthority:=args[4]
		createdBy:=args[5]
		createdDate:=args[6]

		// Insert a row
		ok, err := stub.InsertRow("CertificationDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: certificateId}},
				&shim.Column{Value: &shim.Column_String_{String_: employeeId}},
				&shim.Column{Value: &shim.Column_String_{String_: nameCertificate}},
				&shim.Column{Value: &shim.Column_String_{String_: certificationDate}},
				&shim.Column{Value: &shim.Column_String_{String_: certAuthority}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: createdDate}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		return nil, nil
}

// add the transaction(irrespective of org)
func (t *SKH) addTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8.")
	}

	trxId := args[0]
	timeStamp:=args[1]
	PersonId := args[2]
	
	/*assignerOrg1, err := stub.GetState(args[3])
	assignerOrg := string(assignerOrg1)
	
	source := assignerOrg*/
	source := args[3]
	skill := args[4]
	trxntype := args[5]
	trxnSubType := args[6]
	remarks := args[7]
	
	/*newPoints, _ := strconv.ParseInt(points, 10, 0)
	
	//whether ADD_PENDING, DELETE_PENDING 
	if trxnSubType == "ADD_PENDING" || trxnSubType == "DELETE_PENDING"{
		newPoints = 0
	}
	
*/
	// Get the row pertaining to this Personid
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: PersonId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("EmployeeDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving user with Personid %s. Error %s", PersonId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}
/*
	newRoyaltyPoint := row.Columns[12].GetString_()
	
	if trxntype=="add"{
		earlierMile:=row.Columns[12].GetString_()
		earlierRoyalty, _:=strconv.ParseInt(earlierMile, 10, 0)
		newRoyaltyPoint = strconv.Itoa(int(earlierRoyalty) + int(newPoints))
	}else if trxntype=="delete"{
	
		earlierMile:=row.Columns[12].GetString_()
		earlierRoyalty, _:=strconv.ParseInt(earlierMile, 10, 0)
		newRoyaltiPointtoTest := int(earlierRoyalty) - int(newPoints)
		
		if newRoyaltiPointtoTest < 0 {
			return nil, errors.New("can't deduct as the resulting royalty becoming less than zero.")
		}
		newRoyaltyPoint = strconv.Itoa(int(earlierRoyalty) - int(newPoints))
	}else{
		return nil, fmt.Errorf("Error: Failed retrieving user with ffid %s. Error %s", ffId, err.Error())
	}
	
	
	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this ffid
	err = stub.DeleteRow(
		"EmployeeDetails",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}
	
	employeeId := row.Columns[0].GetString_()
	
	title := row.Columns[1].GetString_()
	gender := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	dob := row.Columns[5].GetString_()
	email := row.Columns[6].GetString_()
	country := row.Columns[7].GetString_()
	address := row.Columns[8].GetString_()
	city := row.Columns[9].GetString_()
	zip := row.Columns[10].GetString_()
	createdBy := row.Columns[11].GetString_()
	createdDate := row.Columns[12].GetString_()


		// Insert a row
		ok, err := stub.InsertRow("EmployeeDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: employeeId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: createdDate}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
*/
		
		//inserting the transaction
		
		// Insert a row
	ok, err := stub.InsertRow("Transaction", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: trxId}},
				&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: PersonId}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: skill}},
				&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
				&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}	
	return nil, nil

}

/*
//get the miles against the ffid (irrespective of org)
func (t *SKH) getMile(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting PersonId to query")
	}

	ffId := args[0]
	

	// Get the row pertaining to this PesonId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: PersonId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("EmployeeDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the PersonId " + PersonId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the PersonId " + PersonId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := GetMile{}
	
	res2E.TotalPoint = row.Columns[12].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}
*/


//get QualificationDetails against the EmployeeId 
func (t *SKH) getQualification(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting EmployeeId to query")
	}

	EmployeeId := args[0]
	//assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("QualificationDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	//assignerOrg1, err := stub.GetState(assignerRole)
	//assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*QualificationDetails{}	
	
	for row := range rows {		
		newApp:= new(QualificationDetails)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.EmployeeId = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.Skill = row.Columns[4].GetString_()
		newApp.Trxntype = row.Columns[5].GetString_()
		newApp.TrxnSubType = row.Columns[6].GetString_()
		newApp.Remarks = row.Columns[7].GetString_()
		
		//if newApp.EmployeeId == EmployeeId && newApp.Source == assignerOrg{
		if newApp.EmployeeId == EmployeeId{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get All transaction against Personid (irrespective of org)
func (t *SKH) getAllTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting PersonId to query")
	}

	PersonId := args[0]
	//assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Transaction", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	//assignerOrg1, err := stub.GetState(assignerRole)
	//assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Transaction{}	
	
	for row := range rows {		
		newApp:= new(Transaction)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.PersonId = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.Skill = row.Columns[4].GetString_()
		newApp.Trxntype = row.Columns[5].GetString_()
		newApp.TrxnSubType = row.Columns[6].GetString_()
		newApp.Remarks = row.Columns[7].GetString_()
		
		if newApp.PersonId == PersonId{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


// to get the deatils of a user against ffid (for internal testing, irrespective of org)
func (t *SKH) getEmployee(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting EmployeeId to query")
	}

	EmployeeId := args[0]
	

	// Get the row pertaining to this EmployeeId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: EmployeeId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("EmployeeDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + EmployeeId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + EmployeeId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	res2E := EmployeeDetails{}
	
	res2E.EmployeeId = row.Columns[0].GetString_()
	res2E.Title = row.Columns[1].GetString_()
	res2E.Gender = row.Columns[2].GetString_()
	res2E.FirstName = row.Columns[3].GetString_()
	res2E.LastName = row.Columns[4].GetString_()
	res2E.Dob = row.Columns[5].GetString_()
	res2E.Email = row.Columns[6].GetString_()
	res2E.Country = row.Columns[7].GetString_()
	res2E.Address = row.Columns[8].GetString_()
	res2E.City = row.Columns[9].GetString_()
	res2E.Zip = row.Columns[10].GetString_()
	res2E.CreatedBy = row.Columns[11].GetString_()
	res2E.CreatedDate = row.Columns[12].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


// verify the user is present or not (for internal testing, irrespective of org)
func (t *SKH) verifyPerson(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting PersonId to query")
	}

	PersonId := args[0]
	dob := args[1]
	

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: PersonId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("EmployeeDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + PersonId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + PersonId + "\"}"
		return nil, errors.New(jsonResp)
	}

	userDob := row.Columns[5].GetString_()
	
	res2E := VerifyU{}
	
	if dob == userDob{
		res2E.Result="success"
	}else{
		res2E.Result="failed"
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}



// Invoke invokes the chaincode
func (t *SKH) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerEmployee" {
		t := SKH{}
		return t.registerEmployee(stub, args)	
	}else if function == "addQualification" { 
		t := SKH{}
		return t.addQualification(stub, args)
	}else if function == "addExperience" { 
		t := SKH{}
		return t.addExperience(stub, args)
	}else if function == "addCertification" { 
		t := SKH{}
		return t.addCertification(stub, args)
	}
	/*else if function == "addTransaction" { 
		t := SKH{}
		return t.addTransaction(stub, args)
	}*/

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *SKH) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getQualification" { 
		t := SKH{}
		return t.getQualification(stub, args)
	}else if function == "getAllTransaction" { 
		t := SKH{}
		return t.getAllTransaction(stub, args)
	} else if function == "getEmployee" { 
		t := SKH{}
		return t.getEmployee(stub, args)
	}else if function == "verifyPerson" { 
		t := SKH{}
		return t.verifyPerson(stub, args)
	}
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(SKH))
	if err != nil {
		fmt.Printf("Error starting SKH: %s", err)
	}
} 
