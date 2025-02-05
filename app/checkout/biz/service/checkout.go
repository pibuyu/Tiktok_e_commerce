package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/checkout/infra/rpc"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart"
	checkout "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/checkout"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/payment"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// 1.get user's cart product
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, "cart is empty")
	}

	//2.calculate order amount
	var total float32
	for _, cartItem := range cartResult.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product.Price
		cost := p * float32(cartItem.Quantity)

		total += cost
	}

	//3.create virtual orderId replace
	var orderId string

	u, _ := uuid.NewRandom()
	orderId = u.String()

	//4.payment request param and clean cart
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.Error(err.Error())
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
