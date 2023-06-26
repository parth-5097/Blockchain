pragma solidity ^0.8.0;

contract DataPrivacySecurity {
    struct Client {
        address clientAddress;
        string name;
        string email;
        bytes32 encryptedData;
    }
    
    mapping(address => Client) public clients;
    
    event ClientRegistered(address clientAddress, string name, string email, bytes32 encryptedData);
    event DataUpdated(address clientAddress, bytes32 encryptedData);
    
    function registerClient(string memory _name, string memory _email, bytes32 _encryptedData) external {
        require(bytes(_name).length > 0, "Invalid client name");
        require(bytes(_email).length > 0, "Invalid client email");
        require(_encryptedData != bytes32(0), "Invalid encrypted data");
        require(clients[msg.sender].clientAddress == address(0), "Client already registered");
        
        clients[msg.sender] = Client(msg.sender, _name, _email, _encryptedData);
        
        emit ClientRegistered(msg.sender, _name, _email, _encryptedData);
    }
    
    function updateData(bytes32 _encryptedData) external {
        require(_encryptedData != bytes32(0), "Invalid encrypted data");
        require(clients[msg.sender].clientAddress != address(0), "Client not registered");
        
        clients[msg.sender].encryptedData = _encryptedData;
        
        emit DataUpdated(msg.sender, _encryptedData);
    }
    
    function getClient(address _clientAddress) external view returns (address, string memory, string memory, bytes32) {
        Client memory client = clients[_clientAddress];
        require(client.clientAddress != address(0), "Client not found");
        
        return (client.clientAddress, client.name, client.email, client.encryptedData);
    }
}
