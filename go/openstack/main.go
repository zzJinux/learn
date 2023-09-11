package main

import (
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
	"github.com/gophercloud/utils/openstack/clientconfig"
)

func main() {
	var err error
	provider, err := clientconfig.AuthenticatedClient(&clientconfig.ClientOpts{})
	client, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	// err = zzcopy(client, "", "adler-test", "go3.mod", "adler-test", "go.mod")
	err = zzget(client, "sadf", "adler-test")
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
}

func zzget(client *gophercloud.ServiceClient, container, object string) error {
	_, err := objects.Get(client, container, object, objects.GetOpts{
		Newest:          false,
		Expires:         "",
		Signature:       "",
		ObjectVersionID: "",
	}).Extract()
	return err
}

type ObjectsCopyOpts struct {
	Destination        string `h:"Destination" required:"true"`
	DestinationAccount string `h:"Destination-Account"`
}

type ObjectsCopyHeader struct {
}

func (opts ObjectsCopyOpts) ToObjectCopyMap() (map[string]string, error) {
	h, err := gophercloud.BuildHeaders(opts)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func zzcopy(client *gophercloud.ServiceClient, destAccount, destContainer, destObject, srcContainer, srcObject string) error {
	r := objects.Copy(client, srcContainer, srcObject, ObjectsCopyOpts{
		Destination:        destContainer + "/" + destObject,
		DestinationAccount: destAccount,
	})
	return r.Err
}
