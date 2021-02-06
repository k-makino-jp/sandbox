// azure is sample package
package azure

type azure interface {
	enqueue()
}

type azureImpl struct{}

func (a *azureImpl) enqueue() {
}

func NewAzure() *azureImpl {
	return &azureImpl{}
}
