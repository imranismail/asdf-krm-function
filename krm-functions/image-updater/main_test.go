package main

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/registry"
	"github.com/google/go-containerregistry/pkg/v1/random"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func TestFilter(t *testing.T) {
	s := httptest.NewServer(registry.New())
	defer s.Close()

	u, err := url.Parse(s.URL)

	if err != nil {
		t.Fatal(err)
	}

	dst, err := name.ParseReference(fmt.Sprintf("%s/test/image", u.Host))
	if err != nil {
		t.Fatal(err)
	}

	img, err := random.Image(1024, 5)
	if err != nil {
		t.Fatal(err)
	}

	if err := remote.Write(dst, img); err != nil {
		t.Fatal(err)
	}

	raw := fmt.Sprintf(`
apiVersion: image-updater.imranismail.dev/v
kind: AutoUpdateImage
metadata:
  name: test
spec:
  target:
    registry: %s
  filterTags:
    pattern: ^latest$
`, dst.Context().RegistryStr())
	testCase := []*yaml.RNode{
		yaml.MustParse(fmt.Sprintf(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-deployment
spec:
  template:
    spec:
      containers:
        - name: test-container
          image: %s
`, dst.String())),
	}

	api := AutoUpdateImage{}

	err = yaml.Unmarshal([]byte(raw), &api)
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}

	result, err := filter(&api)(testCase)
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1, got '%d'", len(result))
	}

	image, err := result[0].Pipe(yaml.Lookup("spec", "template", "spec", "containers", "[name=test-container]", "image"))
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}

	if image.YNode().Value != dst.String() {
		t.Errorf("Expected %s, got '%s'", dst.String(), image.YNode().Value)
	}
}
