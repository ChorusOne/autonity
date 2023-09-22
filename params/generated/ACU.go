package generated

import "strings"
import "github.com/autonity/autonity/accounts/abi"
import "github.com/autonity/autonity/common"

var ACUBytecode = common.Hex2Bytes("60806040523480156200001157600080fd5b5060405162001c1e38038062001c1e83398101604081905262000034916200036d565b858580518251146200005957604051634ff799c560e01b815260040160405180910390fd5b60005b8151811015620000c0576001600160ff1b038282815181106200008357620000836200052c565b60200260200101511115620000ab57604051634ff799c560e01b815260040160405180910390fd5b80620000b78162000558565b9150506200005c565b508751620000d69060039060208b01906200014b565b508651620000ec9060049060208a0190620001a8565b506001869055620000ff86600a62000673565b6002555050600680546001600160a01b039485166001600160a01b03199182161790915560078054938516938216939093179092556008805491909316911617905550620007e3915050565b82805482825590600052602060002090810192821562000196579160200282015b8281111562000196578251829062000185908262000717565b50916020019190600101906200016c565b50620001a4929150620001f4565b5090565b828054828255906000526020600020908101928215620001e6579160200282015b82811115620001e6578251825591602001919060010190620001c9565b50620001a492915062000215565b80821115620001a45760006200020b82826200022c565b50600101620001f4565b5b80821115620001a4576000815560010162000216565b5080546200023a9062000688565b6000825580601f106200024b575050565b601f0160209004906000526020600020908101906200026b919062000215565b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620002af57620002af6200026e565b604052919050565b60006001600160401b03821115620002d357620002d36200026e565b5060051b60200190565b600082601f830112620002ef57600080fd5b81516020620003086200030283620002b7565b62000284565b82815260059290921b840181019181810190868411156200032857600080fd5b8286015b848110156200034557805183529183019183016200032c565b509695505050505050565b80516001600160a01b03811681146200036857600080fd5b919050565b60008060008060008060c087890312156200038757600080fd5b86516001600160401b038111156200039e57600080fd5b8701601f81018913620003b057600080fd5b8051620003c16200030282620002b7565b808282526020820191508b60208460051b8601011115620003e157600080fd5b602084015b60208460051b860101811015620004b85780516001600160401b038111156200040e57600080fd5b8d603f82880101126200042057600080fd5b858101602001516001600160401b038111156200044157620004416200026e565b62000456601f8201601f191660200162000284565b8181528f604083858b01010111156200046e57600080fd5b60005b828110156200049657604081858b010101516020828401015260208101905062000471565b50600060208383010152808652505050602083019250602081019050620003e6565b5060208b0151909950925050506001600160401b03811115620004da57600080fd5b620004e889828a01620002dd565b95505060408701519350620005006060880162000350565b9250620005106080880162000350565b91506200052060a0880162000350565b90509295509295509295565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016200056d576200056d62000542565b5060010190565b600181815b80851115620005b557816000190482111562000599576200059962000542565b80851615620005a757918102915b93841c939080029062000579565b509250929050565b600082620005ce575060016200066d565b81620005dd575060006200066d565b8160018114620005f65760028114620006015762000621565b60019150506200066d565b60ff84111562000615576200061562000542565b50506001821b6200066d565b5060208310610133831016604e8410600b841016171562000646575081810a6200066d565b62000652838362000574565b806000190482111562000669576200066962000542565b0290505b92915050565b6000620006818383620005bd565b9392505050565b600181811c908216806200069d57607f821691505b602082108103620006be57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200071257600081815260208120601f850160051c81016020861015620006ed5750805b601f850160051c820191505b818110156200070e57828155600101620006f9565b5050505b505050565b81516001600160401b038111156200073357620007336200026e565b6200074b8162000744845462000688565b84620006c4565b602080601f8311600181146200078357600084156200076a5750858301515b600019600386901b1c1916600185901b1785556200070e565b600085815260208120601f198616915b82811015620007b45788860151825594840194600190910190840162000793565b5085821015620007d35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61142b80620007f36000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80637adbf97311610076578063b3ab15fb1161005b578063b3ab15fb14610149578063d54d27991461015c578063f51e181a1461017157600080fd5b80637adbf9731461011e578063a2e620451461013157600080fd5b80633fa4f245116100a75780633fa4f245146100f857806344b4708a14610100578063683dd1911461011557600080fd5b806307039ff9146100c3578063146ca531146100e1575b600080fd5b6100cb61017a565b6040516100d89190610b1e565b60405180910390f35b6100ea60005481565b6040519081526020016100d8565b6100ea610253565b61011361010e366004610c45565b610297565b005b6100ea60025481565b61011361012c366004610d95565b610435565b6101396104cd565b60405190151581526020016100d8565b610113610157366004610d95565b61087c565b610164610914565b6040516100d89190610e06565b6100ea60015481565b60606003805480602002602001604051908101604052809291908181526020016000905b8282101561024a5783829060005260206000200180546101bd90610e19565b80601f01602080910402602001604051908101604052809291908181526020018280546101e990610e19565b80156102365780601f1061020b57610100808354040283529160200191610236565b820191906000526020600020905b81548152906001019060200180831161021957829003601f168201915b50505050508152602001906001019061019e565b50505050905090565b60008054600003610290576040517f3a7c017600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060055490565b828280518251146102d4576040517f4ff799c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8151811015610365577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82828151811061031357610313610e6c565b60200260200101511115610353576040517f4ff799c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8061035d81610eca565b9150506102d7565b5060075473ffffffffffffffffffffffffffffffffffffffff1633146103b7576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84516103ca90600390602088019061096c565b5083516103de9060049060208701906109c2565b5060018390556103ef83600a610fca565b6002556040517fdbdcd10543a20811a4a332247f28d03b22686d3281043de35824a06075c06c099061042690879087908790610fd6565b60405180910390a15050505050565b60065473ffffffffffffffffffffffffffffffffffffffff163314610486576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60065460009073ffffffffffffffffffffffffffffffffffffffff163314610521576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006001600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f8743f76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610592573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b6919061100c565b6105c09190611025565b905080600054106105d357600091505090565b600080600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639670c0bc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610643573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610667919061100c565b905060005b600354811015610817576040517f5553442d55534400000000000000000000000000000000000000000000000000602082015260009060270160405160208183030381529060405280519060200120600383815481106106ce576106ce610e6c565b906000526020600020016040516020016106e89190611038565b604051602081830303815290604052805190602001200361070a5750816107ce565b6008546003805460009273ffffffffffffffffffffffffffffffffffffffff1691633c8510fd918991908790811061074457610744610e6c565b906000526020600020016040518363ffffffff1660e01b815260040161076b9291906110cc565b608060405180830381865afa158015610788573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ac919061117d565b905080606001516000146107c7576000965050505050505090565b6020015190505b600482815481106107e1576107e1610e6c565b9060005260206000200154816107f791906111e3565b610801908561122f565b935050808061080f90610eca565b91505061066c565b506108228183611257565b60058190556000849055604080514381524260208201528082018690526060810192909252517f23f161ca67071b3e902d4fa7afade82672c6160677e89d373a830145bdda6d269181900360800190a16001935050505090565b60065473ffffffffffffffffffffffffffffffffffffffff1633146108cd576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6060600480548060200260200160405190810160405280929190818152602001828054801561096257602002820191906000526020600020905b81548152602001906001019080831161094e575b5050505050905090565b8280548282559060005260206000209081019282156109b2579160200282015b828111156109b257825182906109a29082611317565b509160200191906001019061098c565b506109be929150610a09565b5090565b8280548282559060005260206000209081019282156109fd579160200282015b828111156109fd5782518255916020019190600101906109e2565b506109be929150610a26565b808211156109be576000610a1d8282610a3b565b50600101610a09565b5b808211156109be5760008155600101610a27565b508054610a4790610e19565b6000825580601f10610a57575050565b601f016020900490600052602060002090810190610a759190610a26565b50565b600082825180855260208086019550808260051b8401018186016000805b85811015610b10577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080888603018b5283518051808752845b81811015610aea578281018901518882018a01528801610acf565b5086810188018590529b87019b601f019091169094018501935091840191600101610a96565b509198975050505050505050565b602081526000610b316020830184610a78565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610bae57610bae610b38565b604052919050565b600067ffffffffffffffff821115610bd057610bd0610b38565b5060051b60200190565b600082601f830112610beb57600080fd5b81356020610c00610bfb83610bb6565b610b67565b82815260059290921b84018101918181019086841115610c1f57600080fd5b8286015b84811015610c3a5780358352918301918301610c23565b509695505050505050565b600080600060608486031215610c5a57600080fd5b833567ffffffffffffffff80821115610c7257600080fd5b818601915086601f830112610c8657600080fd5b81356020610c96610bfb83610bb6565b82815260059290921b8401810191818101908a841115610cb557600080fd5b8286015b84811015610d6057803586811115610cd15760008081fd5b8701603f81018d13610ce35760008081fd5b84810135604088821115610cf957610cf9610b38565b610d29877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011601610b67565b8281528f82848601011115610d3e5760008081fd5b8282850189830137600092810188019290925250845250918301918301610cb9565b5097505087013592505080821115610d7757600080fd5b50610d8486828701610bda565b925050604084013590509250925092565b600060208284031215610da757600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610b3157600080fd5b600081518084526020808501945080840160005b83811015610dfb57815187529582019590820190600101610ddf565b509495945050505050565b602081526000610b316020830184610dcb565b600181811c90821680610e2d57607f821691505b602082108103610e66577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006000198203610edd57610edd610e9b565b5060010190565b600181815b80851115610f1f578160001904821115610f0557610f05610e9b565b80851615610f1257918102915b93841c9390800290610ee9565b509250929050565b600082610f3657506001610fc4565b81610f4357506000610fc4565b8160018114610f595760028114610f6357610f7f565b6001915050610fc4565b60ff841115610f7457610f74610e9b565b50506001821b610fc4565b5060208310610133831016604e8410600b8410161715610fa2575081810a610fc4565b610fac8383610ee4565b8060001904821115610fc057610fc0610e9b565b0290505b92915050565b6000610b318383610f27565b606081526000610fe96060830186610a78565b8281036020840152610ffb8186610dcb565b915050826040830152949350505050565b60006020828403121561101e57600080fd5b5051919050565b81810381811115610fc457610fc4610e9b565b600080835461104681610e19565b6001828116801561105e5760018114611091576110c0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00841687528215158302870194506110c0565b8760005260208060002060005b858110156110b75781548a82015290840190820161109e565b50505082870194505b50929695505050505050565b82815260006020604081840152600084546110e681610e19565b806040870152606060018084166000811461110857600181146111405761116e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008516838a01528284151560051b8a0101955061116e565b896000528660002060005b858110156111665781548b820186015290830190880161114b565b8a0184019650505b50939998505050505050505050565b60006080828403121561118f57600080fd5b6040516080810181811067ffffffffffffffff821117156111b2576111b2610b38565b8060405250825181526020830151602082015260408301516040820152606083015160608201528091505092915050565b808202600082127f80000000000000000000000000000000000000000000000000000000000000008414161561121b5761121b610e9b565b8181058314821517610fc457610fc4610e9b565b808201828112600083128015821682158216171561124f5761124f610e9b565b505092915050565b60008261128d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60001983147f8000000000000000000000000000000000000000000000000000000000000000831416156112c3576112c3610e9b565b500590565b601f82111561131257600081815260208120601f850160051c810160208610156112ef5750805b601f850160051c820191505b8181101561130e578281556001016112fb565b5050505b505050565b815167ffffffffffffffff81111561133157611331610b38565b6113458161133f8454610e19565b846112c8565b602080601f83116001811461137a57600084156113625750858301515b600019600386901b1c1916600185901b17855561130e565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156113c7578886015182559484019460019091019084016113a8565b50858210156113e55787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220978fbc00b9273a4f791b00db90fd2d4a98a21f0df9a37ee8dbd0164fbd96552a64736f6c63430008150033")

var ACUAbi, _ = abi.JSON(strings.NewReader(`[
   {
      "inputs" : [
         {
            "internalType" : "string[]",
            "name" : "symbols_",
            "type" : "string[]"
         },
         {
            "internalType" : "uint256[]",
            "name" : "quantities_",
            "type" : "uint256[]"
         },
         {
            "internalType" : "uint256",
            "name" : "scale_",
            "type" : "uint256"
         },
         {
            "internalType" : "address",
            "name" : "autonity",
            "type" : "address"
         },
         {
            "internalType" : "address",
            "name" : "operator",
            "type" : "address"
         },
         {
            "internalType" : "address",
            "name" : "oracle",
            "type" : "address"
         }
      ],
      "stateMutability" : "nonpayable",
      "type" : "constructor"
   },
   {
      "inputs" : [],
      "name" : "InvalidBasket",
      "type" : "error"
   },
   {
      "inputs" : [],
      "name" : "NoACUValue",
      "type" : "error"
   },
   {
      "inputs" : [],
      "name" : "Unauthorized",
      "type" : "error"
   },
   {
      "anonymous" : false,
      "inputs" : [
         {
            "indexed" : false,
            "internalType" : "string[]",
            "name" : "symbols",
            "type" : "string[]"
         },
         {
            "indexed" : false,
            "internalType" : "uint256[]",
            "name" : "quantities",
            "type" : "uint256[]"
         },
         {
            "indexed" : false,
            "internalType" : "uint256",
            "name" : "scale",
            "type" : "uint256"
         }
      ],
      "name" : "BasketModified",
      "type" : "event"
   },
   {
      "anonymous" : false,
      "inputs" : [
         {
            "indexed" : false,
            "internalType" : "uint256",
            "name" : "height",
            "type" : "uint256"
         },
         {
            "indexed" : false,
            "internalType" : "uint256",
            "name" : "timestamp",
            "type" : "uint256"
         },
         {
            "indexed" : false,
            "internalType" : "uint256",
            "name" : "round",
            "type" : "uint256"
         },
         {
            "indexed" : false,
            "internalType" : "int256",
            "name" : "value",
            "type" : "int256"
         }
      ],
      "name" : "Updated",
      "type" : "event"
   },
   {
      "inputs" : [
         {
            "internalType" : "string[]",
            "name" : "symbols_",
            "type" : "string[]"
         },
         {
            "internalType" : "uint256[]",
            "name" : "quantities_",
            "type" : "uint256[]"
         },
         {
            "internalType" : "uint256",
            "name" : "scale_",
            "type" : "uint256"
         }
      ],
      "name" : "modifyBasket",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [],
      "name" : "quantities",
      "outputs" : [
         {
            "internalType" : "uint256[]",
            "name" : "",
            "type" : "uint256[]"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   },
   {
      "inputs" : [],
      "name" : "round",
      "outputs" : [
         {
            "internalType" : "uint256",
            "name" : "",
            "type" : "uint256"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   },
   {
      "inputs" : [],
      "name" : "scale",
      "outputs" : [
         {
            "internalType" : "uint256",
            "name" : "",
            "type" : "uint256"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   },
   {
      "inputs" : [],
      "name" : "scaleFactor",
      "outputs" : [
         {
            "internalType" : "uint256",
            "name" : "",
            "type" : "uint256"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "address",
            "name" : "operator",
            "type" : "address"
         }
      ],
      "name" : "setOperator",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "address",
            "name" : "oracle",
            "type" : "address"
         }
      ],
      "name" : "setOracle",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [],
      "name" : "symbols",
      "outputs" : [
         {
            "internalType" : "string[]",
            "name" : "",
            "type" : "string[]"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   },
   {
      "inputs" : [],
      "name" : "update",
      "outputs" : [
         {
            "internalType" : "bool",
            "name" : "status",
            "type" : "bool"
         }
      ],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [],
      "name" : "value",
      "outputs" : [
         {
            "internalType" : "int256",
            "name" : "",
            "type" : "int256"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   }
]
`))