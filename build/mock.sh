go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
~/go/bin/mockgen -source=usecase/transaction/interface.go -destination=usecase/transaction/mock/transaction.go -package=mock
~/go/bin/mockgen -source=usecase/transactionItem/interface.go -destination=usecase/transactionItem/mock/transaction_item.go -package=mock