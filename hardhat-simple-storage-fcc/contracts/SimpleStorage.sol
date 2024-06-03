// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract SimpleStorage{
    // boolean uint int address bytes
    uint256 public  favoriteNumber;

    mapping (string =>uint256) public nameToFavoriteNumber;

    function store(uint256 _favoriteNumber) public virtual {
        favoriteNumber = _favoriteNumber;
    }
    function retrieve() public view returns (uint256){
        return favoriteNumber;
    }
    
    // People public  person = People ({
    //     favoriteNumber:2,name:"Max.G"
    // });
    struct People {
        uint256 favoriteNumber;
        string name;
    }
    People[] public  people;

    function addPerson(string memory _name,uint256 _favoriteNumber) public {
        people.push(People(_favoriteNumber,_name));
        nameToFavoriteNumber[_name] = _favoriteNumber;
    }
}
//0xd9145CCE52D386f254917e481eB44e9943F39138