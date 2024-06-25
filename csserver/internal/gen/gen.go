package gen

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"

	log "github.com/sirupsen/logrus"
)

type Manifest struct {
	ServicesDir string    `yaml:"services_dir"`
	Services    []Service `yaml:"services"`
}

type Service struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func GenServices(manifestPath string) error {
	log.Info("Generating services")
	manifest := LoadManifest(manifestPath)

	for _, service := range manifest.Services {
		log.Printf("Generating service: %s\n", service.Name)
		props := GenProps{
			ServiceName: service.Name,
			OutputPath:  path.Join(manifest.ServicesDir, service.Path),
		}

		err := GenService(props)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

// LoadManifest loads the manifest file and returns a slice of ObjectMap structs.
func LoadManifest(path string) Manifest {
	log.Printf("Loading manifest file: %s\n", path)
	yfile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var manifest Manifest
	err = yaml.Unmarshal(yfile, &manifest)
	if err != nil {
		log.Fatal(err)
	}

	return manifest
}
