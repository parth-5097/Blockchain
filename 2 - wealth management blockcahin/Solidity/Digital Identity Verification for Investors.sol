pragma solidity ^0.8.0;

contract IdentityVerification {
    struct Investor {
        address investorAddress;
        string fullName;
        string identification;
        bool verified;
    }
    
    mapping(address => Investor) public investors;
    
    event InvestorRegistered(address investorAddress, string fullName, string identification);
    event InvestorVerified(address investorAddress);
    
    function registerInvestor(string memory _fullName, string memory _identification) external {
        require(bytes(_fullName).length > 0, "Invalid full name");
        require(bytes(_identification).length > 0, "Invalid identification");
        require(investors[msg.sender].investorAddress == address(0), "Investor already registered");
        
        investors[msg.sender] = Investor(msg.sender, _fullName, _identification, false);
        emit InvestorRegistered(msg.sender, _fullName, _identification);
    }
    
    function verifyInvestor(address _investorAddress) external {
        require(investors[_investorAddress].investorAddress != address(0), "Investor not found");
        require(!investors[_investorAddress].verified, "Investor already verified");
        
        investors[_investorAddress].verified = true;
        emit InvestorVerified(_investorAddress);
    }
    
    function getInvestor(address _investorAddress) external view returns (address, string memory, string memory, bool) {
        Investor memory investor = investors[_investorAddress];
        return (investor.investorAddress, investor.fullName, investor.identification, investor.verified);
    }
}


// verificaiton library
pragma solidity ^0.8.0;

library Verification {
    struct Document {
        string documentHash;
        bool verified;
    }
    
    mapping(address => Document) public verificationRecords;
    
    function submitDocument(address _userAddress, string memory _documentHash) internal {
        require(bytes(_documentHash).length > 0, "Invalid document hash");
        
        verificationRecords[_userAddress] = Document(_documentHash, false);
    }
    
    function verifyDocument(address _userAddress) internal {
        require(bytes(verificationRecords[_userAddress].documentHash).length > 0, "Document not found");
        require(!verificationRecords[_userAddress].verified, "Document already verified");
        
        verificationRecords[_userAddress].verified = true;
    }
}
