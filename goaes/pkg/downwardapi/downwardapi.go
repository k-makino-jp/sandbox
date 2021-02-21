// Package downwardapi Pod情報取得処理用パッケージ
package downwardapi

const (
	podNameFilePath   = "/etc/podinfo/podname"
	namespaceFilePath = "/etc/podinfo/namespace"
)

// DownwardAPI Pod情報取得処理用インターフェース
type DownwardAPI interface {
	GetPodInfo() (PodInfo, error)
}

type downwardAPI struct {
	podNameFilePath   string
	namespaceFilePath string
	ioUtil            ioUtil
}

// GetPodInfo Pod情報取得関数
func (d downwardAPI) GetPodInfo() (PodInfo, error) {
	podNameBytes, err := d.ioUtil.ReadFile(d.podNameFilePath)
	if err != nil {
		return PodInfo{}, err
	}
	namespaceBytes, err := d.ioUtil.ReadFile(d.namespaceFilePath)
	if err != nil {
		return PodInfo{}, err
	}
	podInfo := PodInfo{
		PodName:   string(podNameBytes),
		Namespace: string(namespaceBytes),
	}
	return podInfo, err
}

// NewDonwardAPI コンストラクタ
func NewDonwardAPI() *downwardAPI {
	return &downwardAPI{
		podNameFilePath:   podNameFilePath,
		namespaceFilePath: namespaceFilePath,
		ioUtil:            ioUtilImpl{},
	}
}
