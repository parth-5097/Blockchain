pragma solidity ^0.8.0;

contract CompliancePlatform {
    struct Customer {
        uint256 customerId;
        string name;
        string identification;
        bool kycVerified;
        bool amlVerified;
    }
    
    mapping(address => Customer) public customers;
    uint256 public customerCount;
    
    event CustomerRegistered(address customerAddress, uint256 customerId, string name, string identification);
    event KYCVerified(address customerAddress, uint256 customerId);
    event AMLVerified(address customerAddress, uint256 customerId);
    
    function registerCustomer(string memory _name, string memory _identification) external {
        require(bytes(_name).length > 0, "Invalid customer name");
        require(bytes(_identification).length > 0, "Invalid customer identification");
        
        customers[msg.sender] = Customer(customerCount, _name, _identification, false, false);
        emit CustomerRegistered(msg.sender, customerCount, _name, _identification);
        customerCount++;
    }
    
    function verifyKYC(address _customerAddress) external {
        require(customers[_customerAddress].customerId != 0, "Customer not found");
        require(!customers[_customerAddress].kycVerified, "KYC already verified");
        
        customers[_customerAddress].kycVerified = true;
        emit KYCVerified(_customerAddress, customers[_customerAddress].customerId);
    }
    
    function verifyAML(address _customerAddress) external {
        require(customers[_customerAddress].customerId != 0, "Customer not found");
        require(customers[_customerAddress].kycVerified, "KYC not verified");
        require(!customers[_customerAddress].amlVerified, "AML already verified");
        
        customers[_customerAddress].amlVerified = true;
        emit AMLVerified(_customerAddress, customers[_customerAddress].customerId);
    }
    
    function getCustomer(address _customerAddress) external view returns (uint256, string memory, string memory, bool, bool) {
        Customer memory customer = customers[_customerAddress];
        return (customer.customerId, customer.name, customer.identification, customer.kycVerified, customer.amlVerified);
    }
}


// KYCAML

pragma solidity ^0.8.0;

library KYCAML {
    struct KYC {
        string kycDocument;
        bool verified;
    }
    
    mapping(address => KYC) public kycRecords;
    
    function submitKYCDocument(address _customerAddress, string memory _document) internal {
        require(bytes(_document).length > 0, "Invalid KYC document");
        
        kycRecords[_customerAddress] = KYC(_document, false);
    }
    
    function verifyKYC(address _customerAddress) internal {
        require(bytes(kycRecords[_customerAddress].kycDocument).length > 0, "KYC document not found");
        require(!kycRecords[_customerAddress].verified, "KYC already verified");
        
        kycRecords[_customerAddress].verified = true;
    }
    
    struct AML {
        string amlDocument;
        bool verified;
    }
    
    mapping(address => AML) public amlRecords;
    
    function submitAMLDocument(address _customerAddress, string memory _document) internal {
        require(bytes(_document).length > 0, "Invalid AML document");
        
        amlRecords[_customerAddress] = AML(_document, false);
    }
    
    function verifyAML(address _customerAddress) internal {
        require(bytes(amlRecords[_customerAddress].amlDocument).length > 0, "AML document not found");
        require(!amlRecords[_customerAddress].verified, "AML already verified");
        
        amlRecords[_customerAddress].verified = true;
    }
}
