package kernel

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/yu-org/yu/common"
	"github.com/yu-org/yu/core/protocol"
	"github.com/yu-org/yu/core/types"
)

func (k *Kernel) GetBlock(ctx *gin.Context) {
	blockNumStr := ctx.Query("number")
	blockHashStr := ctx.Query("hash")

	if blockNumStr != "" {
		blockNum, err := hexutil.DecodeUint64(blockNumStr)
		if err != nil {
			protocol.RenderError(ctx, protocol.BlockFailure, err)
			return
		}
		block, err := k.Chain.GetBlockByHeight(common.BlockNum(blockNum))
		if err != nil {
			protocol.RenderError(ctx, protocol.BlockFailure, err)
			return
		}
		protocol.RenderSuccess(ctx, block)
		return
	}

	if blockHashStr != "" {
		blockHash := common.HexToHash(blockHashStr)
		block, err := k.Chain.GetBlock(blockHash)
		if err != nil {
			protocol.RenderError(ctx, protocol.BlockFailure, err)
			return
		}
		protocol.RenderSuccess(ctx, block)
		return
	}

	block, err := k.Chain.GetEndBlock()
	if err != nil {
		protocol.RenderError(ctx, protocol.BlockFailure, err)
		return
	}
	protocol.RenderSuccess(ctx, block)
}

func (k *Kernel) GetReceipt(ctx *gin.Context) {
	txHashStr := ctx.Query("tx_hash")
	txHash := common.HexToHash(txHashStr)
	receipt, err := k.TxDB.GetReceipt(txHash)
	if err != nil {
		protocol.RenderError(ctx, protocol.ReceiptFailure, err)
		return
	}
	protocol.RenderSuccess(ctx, receipt)
}

func (k *Kernel) GetReceipts(ctx *gin.Context) {
	receipts, err := k.getReceipts(ctx)
	if err != nil {
		protocol.RenderError(ctx, protocol.ReceiptFailure, err)
		return
	}
	protocol.RenderSuccess(ctx, receipts)
}

func (k *Kernel) GetReceiptsCount(ctx *gin.Context) {
	receipts, err := k.getReceipts(ctx)
	if err != nil {
		protocol.RenderError(ctx, protocol.ReceiptFailure, err)
		return
	}
	protocol.RenderSuccess(ctx, len(receipts))
}

func (k *Kernel) getReceipts(ctx *gin.Context) ([]*types.Receipt, error) {
	var (
		block *types.CompactBlock
		err   error
	)

	blockHashStr := ctx.Query("block_hash")
	blockNumberStr := ctx.Query("block_number")
	if blockNumberStr != "" {
		blockNumber, err := strconv.Atoi(blockNumberStr)
		if err != nil {
			return nil, err
		}
		block, err = k.Chain.GetCompactBlockByHeight(common.BlockNum(blockNumber))
		if err != nil {
			return nil, err
		}
	} else if blockHashStr != "" {
		blockHash := common.HexToHash(blockHashStr)
		block, err = k.Chain.GetCompactBlock(blockHash)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("none request params")
	}

	var receipts []*types.Receipt
	for _, txHash := range block.TxnsHashes {
		receipt, err := k.TxDB.GetReceipt(txHash)
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, receipt)
	}
	return receipts, nil
}
