package common

import (
	"io"
	"mime/multipart"
	"os"
)

// FormBuilder 表单建造者
type FormBuilder interface {
	createFormFile(fieldName string, file *os.File) error
	writeField(fieldName, value string) error
	close() error
	formDataContentType() string
}

// DefaultFormBuilder 默认的表单建造者
type DefaultFormBuilder struct {
	writer *multipart.Writer
}

// NewDefaultFormBuilder 创建默认表单建造者
func NewDefaultFormBuilder(body io.Writer) *DefaultFormBuilder {
	return &DefaultFormBuilder{
		writer: multipart.NewWriter(body),
	}
}

// createFormFile 创建表单文件
func (b *DefaultFormBuilder) createFormFile(fieldName string, file *os.File) error {
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
func (b *DefaultFormBuilder) writeField(fieldName, value string) error {
	return b.writer.WriteField(fieldName, value)
}

// close 关闭表单
func (b *DefaultFormBuilder) close() error {
	return b.writer.Close()
}

// formDataContentType 表单内容类型
func (b *DefaultFormBuilder) formDataContentType() string {
	return b.writer.FormDataContentType()
}
