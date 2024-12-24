/*
Copyright © 2024 AGou-ops <nil@nil.com>
*/

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeconfig string
	output     string
)

var rootCmd = &cobra.Command{
	Use:   "kubectl-crd2sample <crdName>",
	Short: "Get the sample configuration of the specified CRD",
	Long: `This CLI tool helps you to fetch the schema of a specified CustomResourceDefinition (CRD)
from the Kubernetes cluster and generates a sample YAML file.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		crdName := args[0]
		generateCRDSample(crdName, kubeconfig, output)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Define flags
	rootCmd.Flags().StringVarP(&kubeconfig, "kubeconfig", "c", homedir.HomeDir()+"/.kube/config", "Path to the kubeconfig file (optional, defaults to in-cluster configuration)")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Path to the output file (optional, defaults to stdout)")
}

// generateExampleFromSchema recursively generates an example resource from a schema
func generateExampleFromSchema(schema *apiextensionsv1.JSONSchemaProps) interface{} {
	if schema == nil {
		return nil
	}

	switch schema.Type {
	case "object":
		result := make(map[string]interface{})
		for key, prop := range schema.Properties {
			result[key] = generateExampleFromSchema(&prop)
		}
		return result
	case "array":
		if schema.Items != nil {
			return []interface{}{generateExampleFromSchema(schema.Items.Schema)}
		}
		return []interface{}{"<array>"}
	case "string":
		return "<string>"
	case "integer":
		return 0
	case "boolean":
		return true
	case "number":
		return 0.0
	default:
		return "<nil>"
	}
}

func generateCRDSample(crdName, kubeconfig, output string) {
	// Load kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create apiextensions client
	apiextensionsClient, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating apiextensions client: %v", err)
	}

	// Get CRD details
	crd, err := apiextensionsClient.ApiextensionsV1().CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error fetching CRD: %v", err)
	}

	// Use the schema of the first version of the CRD
	if len(crd.Spec.Versions) == 0 || crd.Spec.Versions[0].Schema == nil {
		log.Fatalf("No schema found in CRD")
	}
	schema := crd.Spec.Versions[0].Schema.OpenAPIV3Schema

	// Generate example resource
	example := generateExampleFromSchema(schema)

	parsedCrdName := strings.Split(crdName, ".")

	// Add required metadata for a Kubernetes resource
	resource := map[string]interface{}{
		"apiVersion": crd.Spec.Group + "/" + crd.Spec.Versions[0].Name,
		"kind":       crd.Spec.Names.Kind,
		"metadata": map[string]string{
			"name": parsedCrdName[0] + "-sample",
		},
		"spec": example,
	}

	// Convert to YAML
	yamlData, err := yaml.Marshal(resource)
	if err != nil {
		log.Fatalf("Error marshalling YAML: %v", err)
	}

	// Write to file or print to stdout
	if output != "" {
		err := os.WriteFile(output, yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
		fmt.Printf("Sample configuration written to %s\n", output)
	} else {
		fmt.Println(string(yamlData))
	}
}
