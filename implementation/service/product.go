package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/sinulingga23/go-pos/api/repository"
	service "github.com/sinulingga23/go-pos/api/service/third-party"
	thirdparty_service "github.com/sinulingga23/go-pos/api/service/third-party"
	"github.com/sinulingga23/go-pos/definition"
	"github.com/sinulingga23/go-pos/domain"
	"github.com/sinulingga23/go-pos/payload"
	thirdparty_payload "github.com/sinulingga23/go-pos/payload/third-party"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productService struct {
	productRepository repository.ProductRepository
	categoryProdductRepositry repository.CategoryProductRepository
	
	cloudStorageService thirdparty_service.CloudStorageService
}

func NewProductService(
	productRepository repository.ProductRepository, 
	categoryProductRepository repository.CategoryProductRepository,
	cloudStorageService service.CloudStorageService,
	) *productService {
	
		return &productService{
		productRepository: productRepository,
		categoryProdductRepositry: categoryProductRepository,
		cloudStorageService: cloudStorageService,
	}
}

func (p productService) Create(createProductRequest payload.CreateProductRequest) (
	*payload.Product, 
	error,
) {
	images := createProductRequest.Images["images"]
	if len(images) == 0 ||
		len(createProductRequest.CategoryProductIds) == 0 {
		return nil, definition.ErrBadRequest
	}


	if len(strings.Trim(createProductRequest.ProductName, " ")) == 0 ||
		len(strings.Trim(createProductRequest.Description, " ")) == 0 {
		return nil, definition.ErrBadRequest
	}
	createProductRequest.ProductName = strings.Trim(createProductRequest.ProductName, " ")
	createProductRequest.Description = strings.Trim(createProductRequest.Description, " ")

	if createProductRequest.Price <= float64(0)  ||
		createProductRequest.Stock <= 0 {
		return nil, definition.ErrBadRequest
	}

	categoryProductIds := []primitive.ObjectID{}
	for i := 0;  i < len(createProductRequest.CategoryProductIds); i++ {
		categoryProductId, err := primitive.ObjectIDFromHex(createProductRequest.CategoryProductIds[i])
		if err != nil {
			return nil, definition.ErrBadRequest
		}

		categoryProductIds = append(categoryProductIds, categoryProductId)
	}

	categoryProducts, err := p.categoryProdductRepositry.FindByIds(context.TODO(), categoryProductIds)
	if err != nil {
		return nil, err
	}

	if len(categoryProductIds) != len(categoryProducts) {
		// there are some category product not exists
		return nil, definition.ErrDataNotFound // 404 or 400 ?
	}

	createdProduct, err := p.productRepository.Create(context.TODO(), domain.Product{
		Id: primitive.NewObjectID(),
		CategoryProductIds: categoryProductIds,
		ProductName: createProductRequest.ProductName,
		Description: createProductRequest.Description,
		Stock: createProductRequest.Stock,
		Price: createProductRequest.Price,
		UrlImages: []domain.UrlImage{},
	})
	if err != nil {
		return nil, definition.ErrInternalServer
	}

	createdProductId := createdProduct.Id
	bucketName := os.Getenv("BUCKET_NAME")
	var mu sync.Mutex
	for _, image := range images {
		go func(image *multipart.FileHeader) {
			theimage, err := image.Open()
			if err != nil {
				log.Printf("[%v]: %v\n", image.Filename, err.Error())
				return
			}
			defer theimage.Close()

			bytes, err := io.ReadAll(theimage)
			if err != nil {
				log.Printf("[%v]: %v\n", image.Filename, err.Error())
				return
			}

			extention := filepath.Ext(image.Filename)
			fileName := fmt.Sprintf("%v-%s-%s%s", time.Now().Unix(), strings.Split(image.Filename, extention)[0], "sinulingga", extention)
		
			mu.Lock()
			urlImage, err := p.cloudStorageService.AddFile(context.TODO(), thirdparty_payload.CloudStorage{
				BucketName: bucketName, 
				FileName: fileName,
				FileBytes: bytes,
			})
			if err != nil {
				mu.Unlock()
				log.Printf("[%v]: %v\n", image.Filename, err.Error())
				return
			}

			if err := p.productRepository.AddUrlImageToProduct(context.TODO(), createdProductId, domain.UrlImage{
				UrlImageId: primitive.NewObjectID(),
				Url: urlImage,
			}); err != nil {
				mu.Unlock()
				log.Printf("[%v]: %v\n", image.Filename, err.Error())
				return
			}

			mu.Unlock()
		}(image)
	}

	payloadCategoryProducts := []payload.CategoryProduct{}
	for _, categoryProduct := range categoryProducts {
		payloadCategoryProducts = append(payloadCategoryProducts, payload.CategoryProduct{
			Id: categoryProduct.Id.Hex(),
			CategoryName: categoryProduct.CategoryName,
			Description: categoryProduct.Description,
		})
	}

	return &payload.Product{
		Id: createdProduct.Id.Hex(),
		CategoryProducts: payloadCategoryProducts,
		ProductName: createdProduct.ProductName,
		Description: createdProduct.Description,
		Stock: createdProduct.Stock,
		Price: createdProduct.Price,
		UrlImages: []payload.UrlImage{},
	}, nil
}

func (p productService) AddImagesToProduct(productId string, addImageToProductRequest payload.AddImageToProductRequest) error {
	if strings.Compare(productId, addImageToProductRequest.ProudctId) != 0 ||
	len(strings.Trim(productId, " ")) == 0 || len(strings.Trim(addImageToProductRequest.ProudctId, " ")) == 0 {
		return definition.ErrBadRequest
	}
	
	images := addImageToProductRequest.Images["images"]
	if len(images) == 0 {
		return definition.ErrBadRequest
	}

	bucketName := os.Getenv("BUCKET_NAME")
	var mu sync.Mutex
	for _, image := range images {
		go func(image *multipart.FileHeader){
			theimage, err := image.Open()
			if err != nil {
				log.Printf("[%v]: %v\n", image.Filename, err.Error())
				return
			}
			defer theimage.Close()

			bytes, err := io.ReadAll(theimage)
			if err != nil {
				log.Printf("[%s]: %s\n", image.Filename, err.Error())
				return
			}

			extention := filepath.Ext(image.Filename)
			fileName := fmt.Sprintf("%v-%s-%s%s", time.Now().Unix(), strings.Split(image.Filename, extention)[0], "sinulingga", extention)

			mu.Lock()
			urlImage, err := p.cloudStorageService.AddFile(context.TODO(), thirdparty_payload.CloudStorage{
				BucketName: bucketName,
				FileName: fileName,
				FileBytes: bytes,
			})
			if err != nil {
				mu.Unlock()
				log.Printf("[%s]: %s\n", image.Filename, err.Error())
				return
			}

			productOID, err := primitive.ObjectIDFromHex(productId)
			if err != nil {
				mu.Unlock()
				log.Printf("[%s]: %s\n", image.Filename, err.Error())
				return
			}

			if err := p.productRepository.AddUrlImageToProduct(context.TODO(), productOID, domain.UrlImage{
				UrlImageId: primitive.NewObjectID(),
				Url: urlImage,
			}); err != nil {
				mu.Unlock()
				log.Printf("[%s]: %s\n", image.Filename, err.Error())
				return
			}
			mu.Unlock()
		}(image)
	}

	return nil
}