pragma solidity ^0.8.0;

contract TradeSettlementSystem {
    struct Trade {
        address buyer;
        address seller;
        uint256 amount;
        bool settled;
    }
    
    mapping(uint256 => Trade) public trades;
    uint256 public tradeCount;
    
    event TradeCreated(uint256 tradeId, address buyer, address seller, uint256 amount);
    event TradeSettled(uint256 tradeId);
    
    function createTrade(address _buyer, address _seller, uint256 _amount) external {
        require(_buyer != address(0), "Invalid buyer address");
        require(_seller != address(0), "Invalid seller address");
        require(_amount > 0, "Invalid trade amount");
        
        trades[tradeCount] = Trade(_buyer, _seller, _amount, false);
        emit TradeCreated(tradeCount, _buyer, _seller, _amount);
        tradeCount++;
    }
    
    function settleTrade(uint256 _tradeId) external {
        require(_tradeId < tradeCount, "Invalid trade ID");
        require(!trades[_tradeId].settled, "Trade has already been settled");
        
        trades[_tradeId].settled = true;
        emit TradeSettled(_tradeId);
    }
    
    function getTrade(uint256 _tradeId) external view returns (address, address, uint256, bool) {
        require(_tradeId < tradeCount, "Invalid trade ID");
        
        Trade memory trade = trades[_tradeId];
        return (trade.buyer, trade.seller, trade.amount, trade.settled);
    }
}
