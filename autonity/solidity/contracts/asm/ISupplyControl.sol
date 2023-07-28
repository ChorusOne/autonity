pragma solidity ^0.8.19;

/*
      .o.        .oooooo..o ooo        ooooo
     .888.      d8P'    `Y8 `88.       .888'
    .8"888.     Y88bo.       888b     d'888
   .8' `888.     `"Y8888o.   8 Y88. .P  888
  .88ooo8888.        `"Y88b  8  `888'   888
 .8'     `888.  oo     .d8P  8    Y     888
o88o     o8888o 8""88888P'  o8o        o888o

       Auton Stabilization Mechanism
*/

/// @title ASM Supply Control Contract Interface
/// @notice Controls the supply of Auton on the network.
/// @dev Intended to be deployed by the protocol at genesis. The operator is
/// expected to be the Stabilization Contract.
interface ISupplyControl {
    /// Auton was minted.
    /// @param recipient Recipient of the Auton
    /// @param amount Amount of Auton minted
    event Mint(address recipient, uint amount);

    /// Auton was burned.
    /// @param amount Amount of Auton burned
    event Burn(uint amount);

    /// The supply of Auton available for minting.
    function availableSupply() external view returns (uint);

    /// The account that is authorized to mint and burn.
    function operator() external view returns (address);

    /// The total supply of Auton under management.
    function totalSupply() external view returns (uint256);

    /// Update the operator that is authorized to mint and burn.
    /// @param operator The new operator account
    /// @dev Only the admin can update the operator account.
    function setOperator(address operator) external;

    /// Mint Auton and send it to the recipient.
    /// @param recipient Recipient of the Auton
    /// @param amount Amount of Auton to mint (non-zero)
    /// @dev Only the operator is authorized to mint Auton. The recipient
    /// cannot be the operator or the zero address.
    function mint(address recipient, uint amount) external;

    /// Burn Auton by taking it out of circulation.
    /// @dev Only the operator is authorized to burn Auton.
    function burn() external payable;
}