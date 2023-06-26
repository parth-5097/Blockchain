pragma solidity ^0.8.0;

contract AssetCustody {
    struct Asset {
        uint256 assetId;
        string name;
        address custodian;
    }
    
    mapping(uint256 => Asset) public assets;
    uint256 public assetCount;
    
    event AssetRegistered(uint256 assetId, string name, address custodian);
    event AssetTransferred(uint256 assetId, address previousCustodian, address newCustodian);
    
    function registerAsset(string memory _name) external {
        require(bytes(_name).length > 0, "Invalid asset name");
        
        assets[assetCount] = Asset(assetCount, _name, msg.sender);
        emit AssetRegistered(assetCount, _name, msg.sender);
        assetCount++;
    }
    
    function transferAsset(uint256 _assetId, address _newCustodian) external {
        require(assets[_assetId].custodian == msg.sender, "Not the custodian of the asset");
        require(_newCustodian != address(0), "Invalid new custodian address");
        
        address previousCustodian = assets[_assetId].custodian;
        assets[_assetId].custodian = _newCustodian;
        emit AssetTransferred(_assetId, previousCustodian, _newCustodian);
    }
    
    function getAsset(uint256 _assetId) external view returns (uint256, string memory, address) {
        Asset memory asset = assets[_assetId];
        return (asset.assetId, asset.name, asset.custodian);
    }
}
