pragma solidity ^0.8.0;

contract RiskAssessment {
    struct Portfolio {
        uint256 portfolioId;
        string name;
        uint256 totalValue;
        mapping(uint256 => Asset) assets;
        uint256 assetCount;
    }
    
    struct Asset {
        uint256 assetId;
        string name;
        uint256 value;
    }
    
    mapping(address => Portfolio) public portfolios;
    uint256 public portfolioCount;
    
    event PortfolioCreated(uint256 portfolioId, string name);
    event AssetAdded(uint256 portfolioId, uint256 assetId, string name, uint256 value);
    event AssetUpdated(uint256 portfolioId, uint256 assetId, uint256 value);
    event PortfolioValueUpdated(uint256 portfolioId, uint256 totalValue);
    
    function createPortfolio(string memory _name) external {
        require(bytes(_name).length > 0, "Invalid portfolio name");
        
        portfolios[msg.sender] = Portfolio(portfolioCount, _name, 0, 0);
        emit PortfolioCreated(portfolioCount, _name);
        portfolioCount++;
    }
    
    function addAsset(uint256 _portfolioId, string memory _name, uint256 _value) external {
        require(bytes(_name).length > 0, "Invalid asset name");
        require(_value > 0, "Invalid asset value");
        
        Portfolio storage portfolio = portfolios[msg.sender];
        require(portfolio.portfolioId == _portfolioId, "Portfolio not found");
        
        Asset memory newAsset = Asset(portfolio.assetCount, _name, _value);
        portfolio.assets[portfolio.assetCount] = newAsset;
        portfolio.assetCount++;
        portfolio.totalValue += _value;
        
        emit AssetAdded(_portfolioId, newAsset.assetId, _name, _value);
        emit PortfolioValueUpdated(_portfolioId, portfolio.totalValue);
    }
    
    function updateAssetValue(uint256 _portfolioId, uint256 _assetId, uint256 _newValue) external {
        Portfolio storage portfolio = portfolios[msg.sender];
        require(portfolio.portfolioId == _portfolioId, "Portfolio not found");
        require(_assetId < portfolio.assetCount, "Asset not found");
        
        uint256 previousValue = portfolio.assets[_assetId].value;
        portfolio.assets[_assetId].value = _newValue;
        portfolio.totalValue = portfolio.totalValue - previousValue + _newValue;
        
        emit AssetUpdated(_portfolioId, _assetId, _newValue);
        emit PortfolioValueUpdated(_portfolioId, portfolio.totalValue);
    }
    
    function getPortfolio(address _owner, uint256 _portfolioId) external view returns (uint256, string memory, uint256) {
        Portfolio memory portfolio = portfolios[_owner];
        require(portfolio.portfolioId == _portfolioId, "Portfolio not found");
        
        return (portfolio.portfolioId, portfolio.name, portfolio.totalValue);
    }
    
    function getAsset(uint256 _portfolioId, uint256 _assetId) external view returns (uint256, string memory, uint256) {
        Portfolio memory portfolio = portfolios[msg.sender];
        require(portfolio.portfolioId == _portfolioId, "Portfolio not found");
        require(_assetId < portfolio.assetCount, "Asset not found");
        
        Asset memory asset = portfolio.assets[_assetId];
        return (asset.assetId, asset.name, asset.value);
    }
}
