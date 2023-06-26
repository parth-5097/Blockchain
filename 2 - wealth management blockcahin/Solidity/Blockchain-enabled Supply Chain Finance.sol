pragma solidity ^0.8.0;

contract SupplyChainFinance {
    struct Invoice {
        uint256 invoiceId;
        address seller;
        address buyer;
        uint256 amount;
        bool financed;
    }
    
    mapping(uint256 => Invoice) public invoices;
    uint256 public invoiceCount;
    
    event InvoiceCreated(uint256 invoiceId, address seller, address buyer, uint256 amount);
    event InvoiceFinanced(uint256 invoiceId);
    
    function createInvoice(address _buyer, uint256 _amount) external {
        require(_buyer != address(0), "Invalid buyer address");
        require(_amount > 0, "Invalid invoice amount");
        
        invoices[invoiceCount] = Invoice(invoiceCount, msg.sender, _buyer, _amount, false);
        emit InvoiceCreated(invoiceCount, msg.sender, _buyer, _amount);
        invoiceCount++;
    }
    
    function financeInvoice(uint256 _invoiceId) external payable {
        Invoice storage invoice = invoices[_invoiceId];
        require(invoice.invoiceId == _invoiceId, "Invoice not found");
        require(!invoice.financed, "Invoice already financed");
        require(invoice.seller == msg.sender, "Only the seller can finance the invoice");
        require(msg.value == invoice.amount, "Incorrect financing amount");
        
        invoice.financed = true;
        
        emit InvoiceFinanced(_invoiceId);
    }
    
    function getInvoice(uint256 _invoiceId) external view returns (uint256, address, address, uint256, bool) {
        Invoice memory invoice = invoices[_invoiceId];
        require(invoice.invoiceId == _invoiceId, "Invoice not found");
        
        return (invoice.invoiceId, invoice.seller, invoice.buyer, invoice.amount, invoice.financed);
    }
}
