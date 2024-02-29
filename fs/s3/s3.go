package s3

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-bamboo/pkg/log"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	allowFileExt         = map[string]int{".png": 1, ".PNG": 1, ".jpg": 1, ".JPG": 1, ".jpeg": 1, ".JPEG": 1, ".gif": 1, ".GIF": 1, ".svg": 1}
	allowAdminFileExt    = map[string]int{".png": 1, ".PNG": 1, ".jpg": 1, ".JPG": 1, ".jpeg": 1, ".JPEG": 1, ".gif": 1, ".GIF": 1, ".mp4": 1, ".svg": 1}
	allowLazyMintFileExt = map[string]string{
		".png": "image", ".PNG": "image", ".jpg": "image", ".JPG": "image",
		".jpeg": "image", ".JPEG": "image",
		".gif": "image", ".GIF": "image",
		".svg": "svg", ".webp": "webp",
		".mp4": "video", ".mp3": "audio",
	}
	contentTypeReg = regexp.MustCompile("(video|image|audio)/.+")
)

type S3Session struct {
	c      *Conf
	s3     *s3.Client
	domain string
}

func New(c *Conf) (s3Sess *S3Session, err error) {
	if c == nil {
		return nil, ErrorConfigNotFound("c is nil")
	}
	if len(c.Key) <= 0 {
		c.Key = os.Getenv("S3_KEY")
		if len(c.Key) <= 0 {
			return nil, ErrorConfigNotFound("key lost")
		}
	}
	if len(c.Secret) <= 0 {
		c.Secret = os.Getenv("S3_SECRET")
		if len(c.Secret) <= 0 {
			return nil, ErrorConfigNotFound("secret lost")
		}
	}
	if len(c.Region) <= 0 {
		c.Region = os.Getenv("S3_REGION")
		if len(c.Region) <= 0 {
			return nil, ErrorConfigNotFound("region lost")
		}
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		//config.WithSharedConfigProfile(opts.profile),
		config.WithRegion(c.Region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     c.Key,
				SecretAccessKey: c.Secret,
				SessionToken:    "",
			},
		}),
	)
	if err != nil {
		return
	}
	s3c := s3.NewFromConfig(cfg)
	s3Sess = &S3Session{
		c:      c,
		s3:     s3c,
		domain: fmt.Sprintf("https://s3.%v.amazonaws.com", c.Region),
	}
	return
}

func (c *S3Session) UploadBytes(ctx context.Context, fileName string, data []byte) (string, error) {
	return c.UploadBytesToDir(ctx, c.c.Dir, fileName, data)
}

func (c *S3Session) UploadBytesToDir(ctx context.Context, dir, fileName string, data []byte) (string, error) {
	return c.UploadBytesToBucketDir(ctx, c.c.Bucket, dir, fileName, data)
}

func (c *S3Session) UploadBytesToBucketDir(ctx context.Context, bucket, dir, fileName string, data []byte) (string, error) {
	contentMd5 := fmt.Sprintf("%x", md5.Sum(data))
	if fileName == "" {
		fileName = contentMd5
	}
	key := fmt.Sprintf("%s/%s", dir, fileName)
	contentType := aws.String(http.DetectContentType(data))
	size := len(data)
	_, err := c.s3.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:        aws.String(bucket),
			Key:           aws.String(key),
			Body:          bytes.NewReader(data),
			ContentLength: int64(size),
			//ContentMD5:    aws.String(contentMd5),
			ContentType: contentType,
		})
	if err != nil {
		return "", err
	}
	//versionId := output.VersionId
	//log.Infof("upload data to s3, version(%v)", *versionId)
	return c.domain + "/" + bucket + "/" + key, nil
}

func (c *S3Session) UploadImage(ctx context.Context, imagePath string, fileName string, proxy *url.URL) (string, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
		},
		Timeout: 2 * time.Second,
	}
	request, err := http.NewRequest("GET", imagePath, nil)
	if err != nil {
		return "", err
	}
	resp, respError := httpClient.Do(request)
	// 默认没有请求到则使用代理进行请求
	if (respError != nil || resp == nil || resp.StatusCode/100 > 3) && proxy != nil {
		proxyURL := http.ProxyURL(proxy)
		httpClientWithProxy := &http.Client{
			Transport: &http.Transport{
				Proxy: proxyURL,
				Dial: (&net.Dialer{
					Timeout:   60 * time.Second,
					Deadline:  time.Now().Add(6 * time.Second),
					KeepAlive: 4 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 4 * time.Second,
			},
			Timeout: 120 * time.Second,
		}
		request, err = http.NewRequest("GET", imagePath, nil)
		if err != nil {
			return "", err
		}
		resp, respError = httpClientWithProxy.Do(request)
	}

	if respError != nil {
		return "", respError
	}
	if resp == nil || resp.StatusCode/100 > 3 {
		return "", errors.New(fmt.Sprintf("Get Resource Error,StatusCode:%d", resp.StatusCode))
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(data)

	if !contentTypeReg.MatchString(contentType) {
		return "", errors.New("Wrong ContentType:" + contentType)
	}
	var objBody *bytes.Reader
	objBody = bytes.NewReader(data)

	contentLength := resp.ContentLength
	if contentLength == -1 {
		contentLength = objBody.Size()
	}

	return c.UploadBytes(ctx, fileName, data)
}

func (c *S3Session) UploadMultipart(ctx context.Context, fileHeader *multipart.FileHeader) (externalUrl string, err error) {
	originFilename := filepath.Base(fileHeader.Filename)
	ext := path.Ext(originFilename)
	if _, ok := allowFileExt[ext]; !ok {
		return "", ErrorNotAllowExt("ext: %v", ext)
	}
	fn, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	buffer, err := io.ReadAll(fn)
	if err != nil {
		return "", err
	}
	imageNameHash := fmt.Sprintf("%x", md5.Sum(buffer))
	key := fmt.Sprintf("%s/%s%s", c.c.Dir, imageNameHash, ext)
	log.Debugf("bucket[%v], key[%v]", c.c.Bucket, key)
	return c.UploadBytes(ctx, key, buffer)
}

func (c *S3Session) UploadApiAvatarDoc(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, dir string, userId, docId int64) (string, error) {
	originFilename := filepath.Base(fileHeader.Filename)
	ext := path.Ext(originFilename)
	if _, ok := allowFileExt[ext]; !ok {
		return "", ErrorNotAllowExt("ext: %v", ext)
	}
	size := fileHeader.Size
	buffer := make([]byte, size)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	sh := md5.New()
	sh.Write(buffer)
	fileName := fmt.Sprintf("%s/%d/%d%s", dir, userId, docId, ext)
	return c.UploadBytes(ctx, fileName, buffer)
}

func (c *S3Session) UploadAdminResource(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	originFilename := filepath.Base(fileHeader.Filename)
	ext := path.Ext(originFilename)
	if _, ok := allowAdminFileExt[ext]; !ok {
		return "", ErrorNotAllowExt("ext: %v", ext)
	}
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	sh := md5.New()
	sh.Write(buffer)
	imageNameHash := hex.EncodeToString(sh.Sum([]byte("")))
	fileName := fmt.Sprintf("images/official/%s%s", imageNameHash, ext)
	return c.UploadBytes(ctx, fileName, buffer)
}

func (c *S3Session) UploadLazyMintFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	originFilename := filepath.Base(fileHeader.Filename)
	ext := strings.ToLower(path.Ext(originFilename))
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	sh := md5.New()
	sh.Write(buffer)
	imageNameHash := hex.EncodeToString(sh.Sum([]byte("")))
	fileName := fmt.Sprintf("images/lazy_mint/%s%s", imageNameHash, ext)
	return c.UploadBytes(ctx, fileName, buffer)
}

func (c *S3Session) CopyObject(ctx context.Context, bucket, dir, filename string, src string) (string, error) {
	key := fmt.Sprintf("%s/%s", dir, filename)
	_, err := c.s3.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(bucket),
		Key:        aws.String(key),
		CopySource: aws.String(src),
	})
	if err != nil {
		return "", err
	}
	outputURI, _ := url.Parse(c.domain)
	return outputURI.JoinPath(bucket, key).String(), nil
}
