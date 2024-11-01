// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract SimpleStorage {
    constructor(){

    }
    uint256 private storedData;

    function set(uint256 x) public {
        storedData = x;
    }

    function get() public view returns (uint256) {
        return storedData;
    }
}
