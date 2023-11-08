package main

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/awslabs/amazon-ecr-credential-helper/ecr-login"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type AutoUpdateImage struct {
	metav1.GroupVersionKind `yaml:",inline" json:",inline"`
	Metadata                metav1.ObjectMeta   `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Spec                    AutoUpdateImageSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type AutoUpdateImageSpec struct {
	Target     AutoUpdateImageTarget     `json:"target,omitempty" yaml:"target,omitempty"`
	FilterTags AutoUpdateImageFilterTags `json:"filterTags,omitempty" yaml:"filterTags,omitempty"`
}

type AutoUpdateImageTarget struct {
	Registry string `json:"registry,omitempty" yaml:"registry,omitempty"`
}

type AutoUpdateImageFilterTags struct {
	Pattern regexp.Regexp `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Extract string        `json:"extract,omitempty" yaml:"extract,omitempty"`
}

func init() {
	log.SetOutput(os.Stderr)
}

func filter(api *AutoUpdateImage) kio.FilterFunc {
	return func(items []*yaml.RNode) ([]*yaml.RNode, error) {
		for _, item := range items {
			var containers *yaml.RNode
			var err error

			switch item.GetKind() {
			case "Deployment", "StatefulSet", "DaemonSet", "Job":
				containers, err = item.Pipe(yaml.Lookup("spec", "template", "spec", "containers"))
			case "CronJob":
				containers, err = item.Pipe(yaml.Lookup("spec", "jobTemplate", "spec", "template", "spec", "containers"))
			default:
				continue
			}

			if err != nil {
				return nil, err
			}

			log.Printf("processing %s", item.GetKind())

			err = containers.VisitElements(func(node *yaml.RNode) error {
				image, err := node.GetString("image")

				if err != nil {
					return err
				}

				parts := strings.Split(image, ":")
				repo, err := name.NewRepository(parts[0])

				if err != nil {
					return err
				}

				if api.Spec.Target.Registry != "" && api.Spec.Target.Registry != repo.Registry.String() {
					return nil
				}

				log.Printf("processing %s", repo.String())

				images, err := remote.List(repo,
					remote.WithAuthFromKeychain(
						authn.NewKeychainFromHelper(
							ecr.NewECRHelper(),
						),
					),
				)

				if err != nil {
					return err
				}

				filtered := []string{}
				rgx := api.Spec.FilterTags.Pattern
				cmp := rgx.SubexpIndex(api.Spec.FilterTags.Extract)

				log.Printf("using pattern: %s", api.Spec.FilterTags.Pattern.String())
				for _, image := range images {
					if rgx.MatchString(image) {
						filtered = append(filtered, image)
					}
				}

				log.Printf("sorting by: %s", api.Spec.FilterTags.Extract)
				sort.Slice(filtered, func(i, j int) bool {
					if cmp == -1 {
						return filtered[i] > filtered[j]
					} else {
						mi := rgx.FindStringSubmatch(filtered[i])
						mj := rgx.FindStringSubmatch(filtered[j])

						return mi[cmp] > mj[cmp]
					}
				})

				if len(filtered) > 0 {
					latest := filtered[0]
					log.Printf("latest tag: %s", latest)

					if err := node.PipeE(yaml.SetField("image", yaml.NewStringRNode(repo.Tag(latest).String()))); err != nil {
						return err
					}
				}

				return nil
			})
		}

		return items, nil
	}
}

func main() {
	var api AutoUpdateImage

	p := framework.SimpleProcessor{Config: &api, Filter: kio.FilterFunc(filter(&api))}
	cmd := command.Build(p, command.StandaloneEnabled, false)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}
