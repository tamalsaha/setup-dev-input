package main

import (
	"fmt"
	cloudv1alpha1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	rsmeta "kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
	"sigs.k8s.io/yaml"
)

func main() {
	var list metav1.List
	list.TypeMeta = metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "List",
	}
	list.Items = []runtime.RawExtension{
		{
			Object: &cloudv1alpha1.Credential{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Credential",
					APIVersion: cloudv1alpha1.SchemeGroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "azure-credentials",
				},
				Spec: cloudv1alpha1.CredentialSpec{},
			},
		},
		{
			Object: &cloudv1alpha1.Credential{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Credential",
					APIVersion: cloudv1alpha1.SchemeGroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "azure-storage-credentials",
				},
				Spec: cloudv1alpha1.CredentialSpec{},
			},
		},
		{
			Object: &rsmeta.ClusterProfile{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ClusterProfile",
					APIVersion: rsmeta.SchemeGroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "aws-dbaas",
				},
				Spec: rsmeta.ClusterProfileSpec{
					Title:               "AWS DBaaS",
					Description:         "",
					Provider:            "",
					RequiredFeatureSets: nil,
				},
			},
		},
	}

	data, err := yaml.Marshal(list)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
