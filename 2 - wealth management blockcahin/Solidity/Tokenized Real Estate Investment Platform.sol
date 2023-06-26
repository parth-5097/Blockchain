pragma solidity ^0.8.0;

contract RealEstateTokenization {
    struct Property {
        uint256 propertyId;
        string name;
        address owner;
        uint256 totalSupply;
        uint256 tokenPrice;
        mapping(address => uint256) balances;
    }
    
    mapping(uint256 => Property) public properties;
    uint256 public propertyCount;
    
    event PropertyCreated(uint256 propertyId, string name, address owner, uint256 totalSupply, uint256 tokenPrice);
    event TokensPurchased(uint256 propertyId, address buyer, uint256 amount);
    event TokensSold(uint256 propertyId, address seller, uint256 amount);
    
    function createProperty(string memory _name, uint256 _totalSupply, uint256 _tokenPrice) external {
        require(bytes(_name).length > 0, "Invalid property name");
        require(_totalSupply > 0, "Invalid total supply");
        require(_tokenPrice > 0, "Invalid token price");
        
        properties[propertyCount] = Property(propertyCount, _name, msg.sender, _totalSupply, _tokenPrice);
        emit PropertyCreated(propertyCount, _name, msg.sender, _totalSupply, _tokenPrice);
        propertyCount++;
    }
    
    function purchaseTokens(uint256 _propertyId, uint256 _amount) external payable {
        Property storage property = properties[_propertyId];
        require(property.propertyId == _propertyId, "Property not found");
        require(property.owner != msg.sender, "Cannot purchase tokens from your own property");
        require(_amount > 0, "Invalid token amount");
        require(msg.value == _amount * property.tokenPrice, "Insufficient payment");
        require(property.totalSupply >= _amount, "Insufficient token supply");
        
        property.balances[msg.sender] += _amount;
        property.totalSupply -= _amount;
        
        emit TokensPurchased(_propertyId, msg.sender, _amount);
    }
    
    function sellTokens(uint256 _propertyId, uint256 _amount) external {
        Property storage property = properties[_propertyId];
        require(property.propertyId == _propertyId, "Property not found");
        require(property.balances[msg.sender] >= _amount, "Insufficient token balance");
        
        uint256 saleValue = _amount * property.tokenPrice;
        property.balances[msg.sender] -= _amount;
        property.totalSupply += _amount;
        
        (bool success, ) = msg.sender.call{value: saleValue}("");
        require(success, "Failed to send payment");
        
        emit TokensSold(_propertyId, msg.sender, _amount);
    }
    
    function getProperty(uint256 _propertyId) external view returns (uint256, string memory, address, uint256, uint256) {
        Property memory property = properties[_propertyId];
        require(property.propertyId == _propertyId, "Property not found");
        
        return (property.propertyId, property.name, property.owner, property.totalSupply, property.tokenPrice);
    }
    
    function getTokenBalance(uint256 _propertyId, address _owner) external view returns (uint256) {
        Property memory property = properties[_propertyId];
        require(property.propertyId == _propertyId, "Property not found");
        
        return property.balances[_owner];
    }
}
