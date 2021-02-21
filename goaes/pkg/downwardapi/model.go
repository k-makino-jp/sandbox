// Package downwardapi Pod情報取得処理用パッケージ
package downwardapi

// PodInfo Pod情報
type PodInfo struct {
	PodName   string `json:"pod_name"`
	Namespace string `json:"namespace"`
}
