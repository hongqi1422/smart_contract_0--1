const {ethers} = require("hardhat")
const {assert,expect} = require("chai")
describe("SimpleStorage",function(){
  //执行测试方法前要做的事情
  let contractFactory,simpleStorage
  beforeEach(async function(){
    contractFactory =  await ethers.getContractFactory("SimpleStorage")
    simpleStorage = await contractFactory.deploy()
  })
  //测试方法1
  it("should start with a favorite number of 0",async function(){
    const currentValue = await simpleStorage.retrieve()
    const expectedValue = "0"
    assert.equal(currentValue.toString(),expectedValue)
  })
  //测试方法2
  it("shoud update when we call store",async function(){
    const expectedValue = "7"
    const txRes = await simpleStorage.store(expectedValue)
    await txRes.wait(1)
    const updatedValue = await simpleStorage.retrieve()
    assert.equal(updatedValue.toString(),expectedValue)
  })
})