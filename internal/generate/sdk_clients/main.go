// Inspired by hashicorp/terraform-provider-aws's internal/generate/servicepackages/main.go

package main

import (
	_ "embed"
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ServiceDatum struct {
	ServiceName	string
}

type TemplateData struct {
	PackageName string
	Services    []ServiceDatum
}

func main() {
	clientFiles, err := filepath.Glob("../../submodules/aws/aws-sdk-go-v2/service/*/api_client.go")
	if err != nil {
		fmt.Println("Failed to glob AWS SDK for list of clients: %+v", err)
		os.Exit(1)
	}

	tmplData := TemplateData{
		PackageName: "provider",
	}
	for _, clientFile := range clientFiles {
		clientFileParts := strings.Split(clientFile, "/")
		service := ServiceDatum{
			ServiceName: clientFileParts[len(clientFileParts)-2],
		}
		tmplData.Services = append(tmplData.Services, service)
	}

	tmpl, err := template.New("sdk_clients").Parse(tmplBody)
	if err != nil {
		fmt.Println("Failed to create template: %+v", err)
		os.Exit(1)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, tmplData)
	if err != nil {
		fmt.Println("Failed to execute template: %+v", err)
		os.Exit(1)
	}

	body, err := format.Source(buffer.Bytes())
	if err != nil {
		fmt.Println("Failed to gofmt: %+v", err)
		os.Exit(1)
	}

	var bodyBuilder strings.Builder
	_, err = bodyBuilder.Write(body)
	if err != nil {
		fmt.Println("Failed to convert to string: %+v", err)
		os.Exit(1)
	}

	file, err := os.OpenFile("../../internal/provider/sdk_clients_gen.go", os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open file to write: %+v", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(bodyBuilder.String())
	if err != nil {
		fmt.Println("Failed to write to file: %+v", err)
		os.Exit(1)
	}
}

//go:embed file.tmpl
var tmplBody string
