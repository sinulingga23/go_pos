package payload


type (
	CloudStorage struct {
		BucketName string
		FileName string
		FileBytes []byte
	}
)