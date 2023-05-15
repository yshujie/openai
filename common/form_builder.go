package common

import (
	"io"
	"mime/multipart"
	"os"
)

// formBuilder 表单建造者
type formBuilder interface {
	createFormFile(fieldName string, file *os.File) error
	writeField(fieldName, value string) error
	close() error
	formDataContentType() string
}

// defaultFormBuilder 默认的表单建造者
type defaultFormBuilder struct {
	writer *multipart.Writer
}

// NewDefaultFormBuilder 创建默认表单建造者
func NewDefaultFormBuilder(body io.Writer) *defaultFormBuilder {
	return &defaultFormBuilder{
		writer: multipart.NewWriter(body),
	}
}

// createFormFile 创建表单文件
func (b *defaultFormBuilder) createFormFile(fieldName string, file *os.File) error {
	fieldWriter, err := b.writer.CreateFormFile(fieldName, file.Name())
	if err != nil {
		return err
	}

	_, err = io.Copy(fieldWriter, file)
	if err != nil {
		return err
	}

	return nil
}

// writeField 写字段
func (b *defaultFormBuilder) writeField(fieldName, value string) error {
	return b.writer.WriteField(fieldName, value)
}

// close 关闭表单
func (b *defaultFormBuilder) close() error {
	return b.writer.Close()
}

// formDataContentType 表单内容类型
func (b *defaultFormBuilder) formDataContentType() string {
	return b.writer.FormDataContentType()
}
