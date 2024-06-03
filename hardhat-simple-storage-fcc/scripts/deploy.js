//imports
const {ethers,run,network} = require("hardhat")

//async main
async function main(){
  const simpleStorageFactory = await ethers.getContractFactory(
    "SimpleStorage"
  )
  console.info("delpoying...")
  const simpleStorage = await simpleStorageFactory.deploy()
  await simpleStorage.deployed()
  console.info(`deployed contract to :${simpleStorage.address}`)
  console.log(network.config)
  if(network.config.chainId === 11155111 && process.env.ETHERSCAN_API_KEY){
    await simpleStorage.deployTransaction.wait(6)
    await verify(simpleStorage.address,[])
  }

  const currentValue = await simpleStorage.retrieve()
  console.log(`current value is:${currentValue}`)
  const transactionRes = await simpleStorage.store(7)
  await transactionRes.wait(1)
  const updatedValue = await simpleStorage.retrieve()
  console.log(`uodeted value is:${updatedValue}`)
}

async function verify(contractAddress,args){

  console.log("verifying contract...")
  try{
  await run("verify:verify",{
    address:contractAddress,
    constructorArguments:args
  })
}catch(e){
  if (e.message.toLowerCase().includes("already verified")) {
    console.log("Already Verified!")
  } else {
    console.log(e)
  }
}
}
// main
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })