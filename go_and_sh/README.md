# Memo

* [NewReryPolicyFactory](https://pkg.go.dev/github.com/Azure/azure-storage-queue-go/azqueue#NewRetryPolicyFactory)
  * [calcDelay](https://github.com/Azure/azure-storage-queue-go/blob/636801874cdd/azqueue/zc_policy_retry.go#L101)
    * Sleep時間を計算するメソッド
    * 計算式は一般的なExponentialBackoff系