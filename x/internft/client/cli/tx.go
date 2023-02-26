package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

// NewTxCmd returns the transaction commands for the module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        internft.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", internft.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		// NewTxCmdUpdateParams(),
		NewTxCmdSend(),
		NewTxCmdNewClass(),
		NewTxCmdUpdateClass(),
		NewTxCmdMintNFT(),
		NewTxCmdBurnNFT(),
		NewTxCmdUpdateNFT(),
	)

	return txCmd
}

// func NewTxCmdUpdateParams() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "update-params [authority] [params-json]",
// 		Args:  cobra.ExactArgs(2),
// 		Short: "Update the module parameters",
// 		Example: `
// Example of the content of params-json:

// {
// }
// `,
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			if err := validateGenerateOnly(cmd); err != nil {
// 				return err
// 			}

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			params, err := parseParams(clientCtx.Codec, args[1])
// 			if err != nil {
// 				return err
// 			}

// 			msg := internft.MsgUpdateParams{
// 				Authority: args[0],
// 				Params:    *params,
// 			}
// 			if err := msg.ValidateBasic(); err != nil {
// 				return err
// 			}

// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }

func NewTxCmdSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [sender] [recipient] [id]",
		Args:  cobra.ExactArgs(3),
		Short: "Send an nft from one account to another account",
		Example: `
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			sender := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, sender); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nft, err := internft.NFTFromString(args[2])
			if err != nil {
				return err
			}

			msg := internft.MsgSend{
				Sender:    sender,
				Recipient: args[1],
				Nft:       *nft,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdNewClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "new-class [owner] [traits-json]",
		Args:    cobra.ExactArgs(2),
		Short:   "Create a class",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			traits, err := ParseTraits(clientCtx.Codec, args[1])
			if err != nil {
				return err
			}

			msg := internft.MsgNewClass{
				Owner:  owner,
				Traits: traits,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdUpdateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-class [class-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Update a class",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			classID := args[0]
			if err := internft.ValidateClassID(classID); err != nil {
				return err
			}

			owner := internft.ClassOwner(classID).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := internft.MsgUpdateClass{
				ClassId: classID,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mint-nft [class-id] [properties-json] [recipient]",
		Args:    cobra.ExactArgs(3),
		Short:   "Mint an nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			classID := args[0]
			if err := internft.ValidateClassID(classID); err != nil {
				return err
			}

			owner := internft.ClassOwner(classID).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			properties, err := ParseProperties(clientCtx.Codec, args[1])
			if err != nil {
				return err
			}

			msg := internft.MsgMintNFT{
				ClassId:    classID,
				Properties: properties,
				Recipient:  args[2],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdBurnNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-nft [owner] [id]",
		Args:  cobra.ExactArgs(2),
		Short: "Burn an nft",
		Example: `
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nft, err := internft.NFTFromString(args[1])
			if err != nil {
				return err
			}

			msg := internft.MsgBurnNFT{
				Owner: owner,
				Nft:   *nft,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdUpdateNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-nft [id] [properties-json]",
		Args:    cobra.ExactArgs(2),
		Short:   "Update a property of an nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			nft, err := internft.NFTFromString(args[0])
			if err != nil {
				return err
			}

			owner := internft.ClassOwner(nft.ClassId).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			properties, err := ParseProperties(clientCtx.Codec, args[1])
			if err != nil {
				return err
			}

			msg := internft.MsgUpdateNFT{
				Nft:        *nft,
				Properties: properties,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
