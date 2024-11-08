package generated

import "strings"
import "github.com/autonity/autonity/accounts/abi"
import "github.com/autonity/autonity/common"

var NonStakeableVestingBytecode = common.Hex2Bytes("608060405234801561001057600080fd5b50604051611eaf380380611eaf83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b611e1c806100936000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a803c8f41161008c578063bc47a07411610066578063bc47a07414610227578063c3aeff1f1461023a578063def25f381461024d578063fca78d111461026057600080fd5b8063a803c8f4146101e1578063a9f45b62146101f4578063aad557261461020757600080fd5b80635558c922116100c85780635558c9221461014e578063635bf9331461016157806380e53d931461019757806381170628146101ce57600080fd5b8063213fe2b7146100ef57806325078446146101185780633577a8f814610139575b600080fd5b6101026100fd3660046119ef565b610273565b60405161010f9190611a0c565b60405180910390f35b61012b610126366004611a91565b6103f4565b60405190815260200161010f565b61014c610147366004611abd565b610411565b005b61014c61015c366004611abd565b610456565b61012b61016f3660046119ef565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b6101aa6101a5366004611abd565b61070c565b6040805182518152602080840151908201529181015115159082015260600161010f565b61014c6101dc366004611ad6565b610770565b61012b6101ef366004611a91565b610806565b61014c610202366004611af8565b610837565b61021a610215366004611a91565b610946565b60405161010f9190611b3a565b61014c610235366004611b7f565b6109ff565b61014c610248366004611abd565b610eac565b61012b61025b366004611a91565b6110bc565b61012b61026e366004611a91565b6110d0565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260016020526040812080546060929067ffffffffffffffff8111156102b6576102b6611bba565b60405190808252806020026020018201604052801561032257816020015b61030f6040518060c0016040528060008152602001600081526020016000815260200160008152602001600081526020016000151581525090565b8152602001906001900390816102d45790505b50905060005b82548110156103ec57600483828154811061034557610345611be9565b90600052602060002001548154811061036057610360611be9565b60009182526020918290206040805160c081018252600690930290910180548352600181015493830193909352600283015490820152600382015460608201526004820154608082015260059091015460ff16151560a082015282518390839081106103ce576103ce611be9565b602002602001018190525080806103e490611c47565b915050610328565b509392505050565b600061040861040384846110d0565b61117b565b90505b92915050565b600061041d33836110d0565b90506104516004828154811061043557610435611be9565b906000526020600020906006020161044c8361129e565b61137f565b505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f7866ee36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e59190611c7f565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105645760405162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207472656173757279206163636f756e74000060448201526064015b60405180910390fd5b600080546040517f7264c4da0000000000000000000000000000000000000000000000000000000081523060048201526024810184905273ffffffffffffffffffffffffffffffffffffffff90911690637264c4da9060440160a060405180830381865afa1580156105da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105fe9190611c9c565b9050806060015181604001516106149190611d33565b8160800151101561068d5760405162461bcd60e51b815260206004820152602760248201527f7363686564756c6520746f74616c206475726174696f6e206e6f74206578706960448201527f7265642079657400000000000000000000000000000000000000000000000000606482015260840161055b565b6000828152600360205260409020600281015460ff166106e0576106e0818360000151815560020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b6106fd33826001015483600001546106f89190611d33565b61141a565b60008082556001909101555050565b610732604051806060016040528060008152602001600081526020016000151581525090565b506000908152600360209081526040918290208251606081018452815481526001820154928101929092526002015460ff1615159181019190915290565b600061077c33846110d0565b90506107878161129e565b8211156107d65760405162461bcd60e51b815260206004820152601960248201527f6e6f7420656e6f75676820756e6c6f636b65642066756e647300000000000000604482015260640161055b565b610800600482815481106107ec576107ec611be9565b90600052602060002090600602018361137f565b50505050565b6000600561081484846110d0565b8154811061082457610824611be9565b9060005260206000200154905092915050565b600054604080517fe7f43c680000000000000000000000000000000000000000000000000000000081529051339273ffffffffffffffffffffffffffffffffffffffff169163e7f43c689160048083019260209291908290030181865afa1580156108a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ca9190611c7f565b73ffffffffffffffffffffffffffffffffffffffff161461092d5760405162461bcd60e51b815260206004820152601a60248201527f63616c6c6572206973206e6f7420746865206f70657261746f72000000000000604482015260640161055b565b600061093984846110d0565b9050610800848284611507565b6109816040518060c0016040528060008152602001600081526020016000815260200160008152602001600081526020016000151581525090565b600461098d84846110d0565b8154811061099d5761099d611be9565b60009182526020918290206040805160c081018252600690930290910180548352600181015493830193909352600283015490820152600382015460608201526004820154608082015260059091015460ff16151560a0820152905092915050565b600054604080517fe7f43c680000000000000000000000000000000000000000000000000000000081529051339273ffffffffffffffffffffffffffffffffffffffff169163e7f43c689160048083019260209291908290030181865afa158015610a6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a929190611c7f565b73ffffffffffffffffffffffffffffffffffffffff1614610af55760405162461bcd60e51b815260206004820152601a60248201527f63616c6c6572206973206e6f7420746865206f70657261746f72000000000000604482015260640161055b565b600080546040517f7264c4da0000000000000000000000000000000000000000000000000000000081523060048201526024810185905273ffffffffffffffffffffffffffffffffffffffff90911690637264c4da9060440160a060405180830381865afa158015610b6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8f9190611c9c565b600084815260036020526040902060028101549192509060ff16610be657610be6818360000151815560020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b8054851115610c5d5760405162461bcd60e51b815260206004820152603860248201527f6e6f7420656e6f7567682066756e647320746f206372656174652061206e657760448201527f20636f6e747261637420756e646572207363686564756c650000000000000000606482015260840161055b565b6000610c68876116bf565b6004549091508114610cbc5760405162461bcd60e51b815260206004820152601360248201527f696e76616c696420636f6e747261637420696400000000000000000000000000604482015260640161055b565b6000610cd18460200151856000015189611713565b90506000610cf589610ce3848b611d46565b87604001518989606001516000611732565b6004805460018082018355600092835283517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b6006938402908101919091556020808601517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c8301556040808701517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19d84015560608701517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19e84015560808701517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19f84015560a08701517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd1a090930180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016931515939093179092556005805493840190557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0909201879055878452919052812089905585549192508991869190610e81908490611d46565b9250508190555081846001016000828254610e9c9190611d33565b9091555050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f7866ee36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3b9190611c7f565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fb55760405162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207472656173757279206163636f756e740000604482015260640161055b565b6000818152600360205260409020600281015460ff166110a357600080546040517f7264c4da0000000000000000000000000000000000000000000000000000000081523060048201526024810185905273ffffffffffffffffffffffffffffffffffffffff90911690637264c4da9060440160a060405180830381865afa158015611045573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110699190611c9c565b5190506110a18282815560020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b505b6110b133826001015461141a565b600060019091015550565b60006104086110cb84846110d0565b61129e565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604081205482106111445760405162461bcd60e51b815260206004820152601360248201527f696e76616c696420636f6e747261637420696400000000000000000000000000604482015260640161055b565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260016020526040902080548390811061082457610824611be9565b60008054828252600660205260408083205490517f7264c4da0000000000000000000000000000000000000000000000000000000081523060048201526024810191909152829173ffffffffffffffffffffffffffffffffffffffff1690637264c4da9060440160a060405180830381865afa1580156111ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112239190611c9c565b90506005838154811061123857611238611be9565b90600052602060002001546004848154811061125657611256611be9565b9060005260206000209060060201600101546112838360200151846000015161127e886118c9565b611713565b61128d9190611d46565b6112979190611d46565b9392505050565b600080600483815481106112b4576112b4611be9565b9060005260206000209060060201905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166389c614b86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561132f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113539190611d59565b816003015482600201546113679190611d33565b11156113765750600092915050565b6112978361117b565b81546000908211156113ad5782546113979083611d46565b90506113a88333856000015461192b565b6113be565b81156113be576113be83338461192b565b60005473ffffffffffffffffffffffffffffffffffffffff16337feed10c470424824e4a4309075162f10b9989088b23fbed2349698cedd44493fb6114038486611d46565b60405190815260200160405180910390a392915050565b600080546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590529091169063a9059cbb906044016020604051808303816000875af1158015611494573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b89190611d72565b9050806104515760405162461bcd60e51b815260206004820152601360248201527f4e544e206e6f74207472616e7366657272656400000000000000000000000000604482015260640161055b565b73ffffffffffffffffffffffffffffffffffffffff831660009081526001602081905260408220805490929161153c91611d46565b67ffffffffffffffff81111561155457611554611bba565b60405190808252806020026020018201604052801561157d578160200160208202803683370190505b5090506000805b835481101561160b57858482815481106115a0576115a0611be9565b906000526020600020015403156115f9578381815481106115c3576115c3611be9565b90600052602060002001548383806115da90611c47565b9450815181106115ec576115ec611be9565b6020026020010181815250505b8061160381611c47565b915050611584565b5073ffffffffffffffffffffffffffffffffffffffff8616600090815260016020908152604090912083516116429285019061196a565b5073ffffffffffffffffffffffffffffffffffffffff808516600081815260016020818152604080842080549384018155845292200188905551918816917f893ca4c0017fb7a30186cb3f7c82b127e989d3079f8473989c10e06edf1cf738906116af9089815260200190565b60405180910390a3505050505050565b6002805473ffffffffffffffffffffffffffffffffffffffff831660009081526001602081815260408320805492830181558352822001829055825490928361170783611c47565b90915550909392505050565b6000826117208386611d94565b61172a9190611dab565b949350505050565b61176d6040518060c0016040528060008152602001600081526020016000815260200160008152602001600081526020016000151581525090565b73ffffffffffffffffffffffffffffffffffffffff87166117f65760405162461bcd60e51b815260206004820152602260248201527f62656e65666963696172792063616e6e6f74206265207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161055b565b600086116118465760405162461bcd60e51b815260206004820152601960248201527f616d6f756e742073686f756c6420626520706f73697469766500000000000000604482015260640161055b565b8383116118955760405162461bcd60e51b815260206004820152601e60248201527f656e64206d7573742062652067726561746572207468616e20636c6966660000604482015260640161055b565b506040805160c0810182529586526000602087015285019390935260608401919091526080830152151560a0820152919050565b600080600483815481106118df576118df611be9565b906000526020600020906006020190506005838154811061190257611902611be9565b9060005260206000200154816001015482600001546119219190611d33565b6112979190611d33565b8083600001600082825461193f9190611d46565b925050819055508083600101600082825461195a9190611d33565b909155506104519050828261141a565b8280548282559060005260206000209081019282156119a5579160200282015b828111156119a557825182559160200191906001019061198a565b506119b19291506119b5565b5090565b5b808211156119b157600081556001016119b6565b73ffffffffffffffffffffffffffffffffffffffff811681146119ec57600080fd5b50565b600060208284031215611a0157600080fd5b8135611297816119ca565b6020808252825182820181905260009190848201906040850190845b81811015611a8557611a72838551805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a0810151151560a08301525050565b9284019260c09290920191600101611a28565b50909695505050505050565b60008060408385031215611aa457600080fd5b8235611aaf816119ca565b946020939093013593505050565b600060208284031215611acf57600080fd5b5035919050565b60008060408385031215611ae957600080fd5b50508035926020909101359150565b600080600060608486031215611b0d57600080fd5b8335611b18816119ca565b9250602084013591506040840135611b2f816119ca565b809150509250925092565b60c0810161040b8284805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a0810151151560a08301525050565b60008060008060808587031215611b9557600080fd5b8435611ba0816119ca565b966020860135965060408601359560600135945092505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c7857611c78611c18565b5060010190565b600060208284031215611c9157600080fd5b8151611297816119ca565b600060a08284031215611cae57600080fd5b60405160a0810181811067ffffffffffffffff82111715611cf8577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b806040525082518152602083015160208201526040830151604082015260608301516060820152608083015160808201528091505092915050565b8082018082111561040b5761040b611c18565b8181038181111561040b5761040b611c18565b600060208284031215611d6b57600080fd5b5051919050565b600060208284031215611d8457600080fd5b8151801515811461129757600080fd5b808202811582820484141761040b5761040b611c18565b600082611de1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea2646970667358221220d1c9db76f1ebe9eabc251b9814dbe9cd20fe4678a95b2bcf0a2f5e9cfd6b6b3a64736f6c63430008150033")

var NonStakeableVestingAbi, _ = abi.JSON(strings.NewReader(`[
   {
      "inputs" : [
         {
            "internalType" : "address payable",
            "name" : "_autonity",
            "type" : "address"
         }
      ],
      "stateMutability" : "nonpayable",
      "type" : "constructor"
   },
   {
      "anonymous" : false,
      "inputs" : [
         {
            "indexed" : true,
            "internalType" : "address",
            "name" : "newBeneficiary",
            "type" : "address"
         },
         {
            "indexed" : true,
            "internalType" : "address",
            "name" : "oldBeneficiary",
            "type" : "address"
         },
         {
            "indexed" : false,
            "internalType" : "uint256",
            "name" : "contractID",
            "type" : "uint256"
         }
      ],
      "name" : "BeneficiaryChanged",
      "type" : "event"
   },
   {
      "anonymous" : false,
      "inputs" : [
         {
            "indexed" : true,
            "internalType" : "address",
            "name" : "to",
            "type" : "address"
         },
         {
            "indexed" : true,
            "internalType" : "address",
            "name" : "token",
            "type" : "address"
         },
         {
            "indexed" : false,
            "internalType" : "uint256",
            "name" : "amount",
            "type" : "uint256"
         }
      ],
      "name" : "FundsReleased",
      "type" : "event"
   },
   {
      "inputs" : [
         {
            "internalType" : "address",
            "name" : "_beneficiary",
            "type" : "address"
         },
         {
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         },
         {
            "internalType" : "address",
            "name" : "_recipient",
            "type" : "address"
         }
      ],
      "name" : "changeContractBeneficiary",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "address",
            "name" : "_beneficiary",
            "type" : "address"
         },
         {
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         }
      ],
      "name" : "getContract",
      "outputs" : [
         {
            "components" : [
               {
                  "internalType" : "uint256",
                  "name" : "currentNTNAmount",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "withdrawnValue",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "start",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "cliffDuration",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "totalDuration",
                  "type" : "uint256"
               },
               {
                  "internalType" : "bool",
                  "name" : "canStake",
                  "type" : "bool"
               }
            ],
            "internalType" : "struct ContractBase.Contract",
            "name" : "",
            "type" : "tuple"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "address",
            "name" : "_beneficiary",
            "type" : "address"
         }
      ],
      "name" : "getContracts",
      "outputs" : [
         {
            "components" : [
               {
                  "internalType" : "uint256",
                  "name" : "currentNTNAmount",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "withdrawnValue",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "start",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "cliffDuration",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "totalDuration",
                  "type" : "uint256"
               },
               {
                  "internalType" : "bool",
                  "name" : "canStake",
                  "type" : "bool"
               }
            ],
            "internalType" : "struct ContractBase.Contract[]",
            "name" : "",
            "type" : "tuple[]"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "address",
            "name" : "_beneficiary",
            "type" : "address"
         },
         {
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         }
      ],
      "name" : "getExpiredFunds",
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
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         }
      ],
      "name" : "getScheduleTracker",
      "outputs" : [
         {
            "components" : [
               {
                  "internalType" : "uint256",
                  "name" : "unsubscribedAmount",
                  "type" : "uint256"
               },
               {
                  "internalType" : "uint256",
                  "name" : "expiredFromContract",
                  "type" : "uint256"
               },
               {
                  "internalType" : "bool",
                  "name" : "initialized",
                  "type" : "bool"
               }
            ],
            "internalType" : "struct NonStakeableVesting.ScheduleTracker",
            "name" : "",
            "type" : "tuple"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "address",
            "name" : "_beneficiary",
            "type" : "address"
         },
         {
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         }
      ],
      "name" : "getUniqueContractID",
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
            "name" : "_beneficiary",
            "type" : "address"
         },
         {
            "internalType" : "uint256",
            "name" : "_amount",
            "type" : "uint256"
         },
         {
            "internalType" : "uint256",
            "name" : "_scheduleID",
            "type" : "uint256"
         },
         {
            "internalType" : "uint256",
            "name" : "_cliffDuration",
            "type" : "uint256"
         }
      ],
      "name" : "newContract",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "uint256",
            "name" : "_scheduleID",
            "type" : "uint256"
         }
      ],
      "name" : "releaseAllFundsForTreasury",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         }
      ],
      "name" : "releaseAllNTN",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "uint256",
            "name" : "_scheduleID",
            "type" : "uint256"
         }
      ],
      "name" : "releaseExpiredFundsForTreasury",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         },
         {
            "internalType" : "uint256",
            "name" : "_amount",
            "type" : "uint256"
         }
      ],
      "name" : "releaseNTN",
      "outputs" : [],
      "stateMutability" : "nonpayable",
      "type" : "function"
   },
   {
      "inputs" : [
         {
            "internalType" : "address",
            "name" : "_beneficiary",
            "type" : "address"
         }
      ],
      "name" : "totalContracts",
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
            "name" : "_beneficiary",
            "type" : "address"
         },
         {
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         }
      ],
      "name" : "vestedFunds",
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
            "name" : "_beneficiary",
            "type" : "address"
         },
         {
            "internalType" : "uint256",
            "name" : "_id",
            "type" : "uint256"
         }
      ],
      "name" : "withdrawableVestedFunds",
      "outputs" : [
         {
            "internalType" : "uint256",
            "name" : "",
            "type" : "uint256"
         }
      ],
      "stateMutability" : "view",
      "type" : "function"
   }
]
`))