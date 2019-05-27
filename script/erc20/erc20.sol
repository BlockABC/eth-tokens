pragma solidity ^0.5.7;

contract ERC20 {
    string public constant name = "";
    string public constant NAME = "";
    string public constant symbol = "";
    string public constant SYMBOL = "";
    uint8 public constant decimals = 0;
    uint8 public constant DECIMAL = 0;
    function totalSupply() public pure returns (uint);
    function balanceOf(address tokenOwner) public pure returns (uint balance);
    function allowance(address tokenOwner, address spender) public pure returns (uint remaining);
    function transfer(address to, uint tokens) public returns (bool success);
    function approve(address spender, uint tokens) public returns (bool success);
    function transferFrom(address from, address to, uint tokens) public returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
