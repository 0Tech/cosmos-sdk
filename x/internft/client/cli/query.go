package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/x/internft"
)

// NewQueryCmd returns the query commands for the module
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        internft.ModuleName,
		Short:                      fmt.Sprintf("%s query subcommands", internft.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewQueryCmdParams(),
		NewQueryCmdClass(),
		NewQueryCmdClasses(),
		NewQueryCmdNFT(),
		NewQueryCmdNFTs(),
		NewQueryCmdOwner(),
	)

	return cmd
}

func NewQueryCmdParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the module parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := internft.QueryParamsRequest{}

			queryClient := internft.NewQueryClient(clientCtx)
			res, err := queryClient.Params(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class [class-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query a class",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			classID := args[0]
			if err := internft.ValidateClassID(classID); err != nil {
				return err
			}

			req := internft.QueryClassRequest{
				ClassId: classID,
			}

			queryClient := internft.NewQueryClient(clientCtx)
			res, err := queryClient.Class(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdClasses() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "classes",
		Args:  cobra.NoArgs,
		Short: "Query the classes",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := internft.QueryClassesRequest{}

			queryClient := internft.NewQueryClient(clientCtx)
			res, err := queryClient.Classes(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query an nft",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nft, err := internft.NFTFromString(args[0])
			if err != nil {
				return err
			}

			req := internft.QueryNFTRequest{
				ClassId: nft.ClassId,
				Id:      nft.Id.String(),
			}

			queryClient := internft.NewQueryClient(clientCtx)
			res, err := queryClient.NFT(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdNFTs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nfts [class-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the nfts",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			classID := args[0]
			if err := internft.ValidateClassID(classID); err != nil {
				return err
			}

			req := internft.QueryNFTsRequest{
				ClassId: classID,
			}

			queryClient := internft.NewQueryClient(clientCtx)
			res, err := queryClient.NFTs(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the owner of an nft",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nft, err := internft.NFTFromString(args[0])
			if err != nil {
				return err
			}

			req := internft.QueryOwnerRequest{
				ClassId: nft.ClassId,
				Id:      nft.Id.String(),
			}

			queryClient := internft.NewQueryClient(clientCtx)
			res, err := queryClient.Owner(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
