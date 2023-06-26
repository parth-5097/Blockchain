pragma solidity ^0.8.0;

contract ProxyVoting {
    struct Proposal {
        uint256 proposalId;
        string title;
        string description;
        uint256 yesVotes;
        uint256 noVotes;
        bool executed;
        mapping(address => bool) hasVoted;
    }
    
    mapping(uint256 => Proposal) public proposals;
    uint256 public proposalCount;
    
    event ProposalCreated(uint256 proposalId, string title, string description);
    event VoteCasted(uint256 proposalId, address voter, bool supports);
    event ProposalExecuted(uint256 proposalId);
    
    function createProposal(string memory _title, string memory _description) external {
        require(bytes(_title).length > 0, "Invalid proposal title");
        require(bytes(_description).length > 0, "Invalid proposal description");
        
        proposals[proposalCount] = Proposal(proposalCount, _title, _description, 0, 0, false);
        emit ProposalCreated(proposalCount, _title, _description);
        proposalCount++;
    }
    
    function castVote(uint256 _proposalId, bool _supports) external {
        Proposal storage proposal = proposals[_proposalId];
        require(proposal.proposalId == _proposalId, "Proposal not found");
        require(!proposal.executed, "Proposal already executed");
        require(!proposal.hasVoted[msg.sender], "Already voted");
        
        if (_supports) {
            proposal.yesVotes++;
        } else {
            proposal.noVotes++;
        }
        
        proposal.hasVoted[msg.sender] = true;
        
        emit VoteCasted(_proposalId, msg.sender, _supports);
    }
    
    function executeProposal(uint256 _proposalId) external {
        Proposal storage proposal = proposals[_proposalId];
        require(proposal.proposalId == _proposalId, "Proposal not found");
        require(!proposal.executed, "Proposal already executed");
        require(proposal.yesVotes > proposal.noVotes, "Proposal not approved");
        
        proposal.executed = true;
        
        emit ProposalExecuted(_proposalId);
    }
    
    function getProposal(uint256 _proposalId) external view returns (uint256, string memory, string memory, uint256, uint256, bool) {
        Proposal memory proposal = proposals[_proposalId];
        require(proposal.proposalId == _proposalId, "Proposal not found");
        
        return (proposal.proposalId, proposal.title, proposal.description, proposal.yesVotes, proposal.noVotes, proposal.executed);
    }
    
    function hasVoted(uint256 _proposalId, address _voter) external view returns (bool) {
        Proposal memory proposal = proposals[_proposalId];
        require(proposal.proposalId == _proposalId, "Proposal not found");
        
        return proposal.hasVoted[_voter];
    }
}
