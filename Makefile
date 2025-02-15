export ROOT_MOD=github.com/Blue-Berrys/Tiktok_e_commerce
#export CWGO=/Users/mac/go/bin/cwgo

#server
.PHONY:gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --I ../../idl --idl ../../idl/user.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


.PHONY:gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --I ../../idl --idl ../../idl/product.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


.PHONY:gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart &&  cwgo server --type RPC --service cart --module ${ROOT_MOD}/app/cart --I ../../idl --idl ../../idl/cart.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


.PHONY:gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module ${ROOT_MOD}/app/payment --I ../../idl --idl ../../idl/payment.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


.PHONY:gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module ${ROOT_MOD}/app/checkout --I ../../idl --idl ../../idl/checkout.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


.PHONY:gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module ${ROOT_MOD}/app/order --I ../../idl --idl ../../idl/order.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY:gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module ${ROOT_MOD}/app/email --I ../../idl --idl ../../idl/email.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


.PHONY:gen-ai
gen-ai:
	@cd rpc_gen && cwgo client --type RPC --service ai --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/ai.proto
	@cd app/ai && cwgo server --type RPC --service ai --module ${ROOT_MOD}/app/ai --I ../../idl --idl ../../idl/ai.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


#需要额外注意:生成frontend代码时type=HTTP
.PHONY:gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && cwgo server --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend -I ../../idl  --idl ../../idl/frontend/auth_page.proto

.PHONY:gen-frontend-home
gen-frontend-home:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/home.proto

.PHONY:gen-frontend-product
gen-frontend-product:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/product_page.proto

.PHONY:gen-frontend-category
gen-frontend-category:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/category_page.proto

.PHONY:gen-frontend-cart
gen-frontend-cart:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/cart_page.proto


.PHONY:gen-frontend-checkout
gen-frontend-checkout:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/checkout_page.proto

.PHONY:gen-frontend-order
gen-frontend-order:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/order_page.proto
