package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/aws"
	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/gcp"
	"github.com/genevieve/leftovers/nsxt"
	"github.com/genevieve/leftovers/vsphere"
	flags "github.com/jessevdk/go-flags"
)

type opts struct {
	Version bool `short:"v"  long:"version"                     description:"Print version."`

	IAAS      string `short:"i"  long:"iaas"        env:"BBL_IAAS"  description:"The IaaS for clean up."  `
	NoConfirm bool   `short:"n"  long:"no-confirm"                  description:"Destroy resources without prompting. This is dangerous, make good choices!"`
	DryRun    bool   `short:"d"  long:"dry-run"                     description:"List all resources without deleting any."`
	Filter    string `short:"f"  long:"filter"                      description:"Filtering resources by an environment name."`
	Type      string `short:"t"  long:"type"                        description:"Type of resource to delete."`

	AWSAccessKeyID       string `long:"aws-access-key-id"        env:"BBL_AWS_ACCESS_KEY_ID"        description:"AWS access key id."`
	AWSSecretAccessKey   string `long:"aws-secret-access-key"    env:"BBL_AWS_SECRET_ACCESS_KEY"    description:"AWS secret access key."`
	AWSRegion            string `long:"aws-region"               env:"BBL_AWS_REGION"               description:"AWS region."`
	AzureClientID        string `long:"azure-client-id"          env:"BBL_AZURE_CLIENT_ID"          description:"Azure client id."`
	AzureClientSecret    string `long:"azure-client-secret"      env:"BBL_AZURE_CLIENT_SECRET"      description:"Azure client secret."`
	AzureTenantID        string `long:"azure-tenant-id"          env:"BBL_AZURE_TENANT_ID"          description:"Azure tenant id."`
	AzureSubscriptionID  string `long:"azure-subscription-id"    env:"BBL_AZURE_SUBSCRIPTION_ID"    description:"Azure subscription id."`
	GCPServiceAccountKey string `long:"gcp-service-account-key"  env:"BBL_GCP_SERVICE_ACCOUNT_KEY"  description:"GCP service account key path."`
	VSphereIP            string `long:"vsphere-vcenter-ip"       env:"BBL_VSPHERE_VCENTER_IP"       description:"vSphere vCenter IP address."`
	VSpherePassword      string `long:"vsphere-vcenter-password" env:"BBL_VSPHERE_VCENTER_PASSWORD" description:"vSphere vCenter password."`
	VSphereUser          string `long:"vsphere-vcenter-user"     env:"BBL_VSPHERE_VCENTER_USER"     description:"vSphere vCenter username."`
	VSphereDC            string `long:"vsphere-vcenter-dc"       env:"BBL_VSPHERE_VCENTER_DC"       description:"vSphere vCenter datacenter."`
	NSXTManagerHost      string `long:"nsxt-manager-host"        env:"BBL_NSXT_MANAGER_HOST"        description:"NSX-T manager IP address or domain name."`
	NSXTUser             string `long:"nsxt-username"            env:"BBL_NSXT_USERNAME"            description:"NSX-T manager username."`
	NSXTPassword         string `long:"nsxt-password"            env:"BBL_NSXT_PASSWORD"            description:"NSX-T manager password."`
}

type leftovers interface {
	Delete(filter string) error
	DeleteType(filter, rType string) error
	List(filter string)
	Types()
}

var Version = "dev"

const (
	AWS     = "aws"
	GCP     = "gcp"
	Azure   = "azure"
	VSphere = "vsphere"
	NSXT    = "nsxt"
)

func main() {
	log.SetFlags(0)

	var o opts
	parser := flags.NewParser(&o, flags.HelpFlag|flags.PrintErrors)
	remaining, err := parser.ParseArgs(os.Args)
	if err != nil {
		return
	}

	command := "destroy"
	if len(remaining) > 1 {
		command = remaining[1]
	}

	if o.Version {
		log.Printf("%s\n", Version)
		return
	}

	logger := app.NewLogger(os.Stdout, os.Stdin, o.NoConfirm)

	var l leftovers

	switch o.IAAS {
	case AWS:
		o = useOtherEnvVars(o, AWS)
		l, err = aws.NewLeftovers(logger, o.AWSAccessKeyID, o.AWSSecretAccessKey, o.AWSRegion)
	case Azure:
		o = useOtherEnvVars(o, Azure)
		l, err = azure.NewLeftovers(logger, o.AzureClientID, o.AzureClientSecret, o.AzureSubscriptionID, o.AzureTenantID)
	case GCP:
		o = useOtherEnvVars(o, GCP)
		l, err = gcp.NewLeftovers(logger, o.GCPServiceAccountKey)
	case NSXT:
		o = useOtherEnvVars(o, NSXT)
		l, err = nsxt.NewLeftovers(logger, o.NSXTManagerHost, o.NSXTUser, o.NSXTPassword)
	case VSphere:
		if o.Filter == "" {
			log.Fatalf("--filter is required for vSphere.")
		}
		if o.NoConfirm {
			log.Fatalf("--no-confirm is not supported for vSphere.")
		}
		l, err = vsphere.NewLeftovers(logger, o.VSphereIP, o.VSphereUser, o.VSpherePassword, o.VSphereDC)
	default:
		err = errors.New("Missing or unsupported BBL_IAAS.")
	}

	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	if command == "types" {
		l.Types()
		return
	}

	if o.DryRun {
		l.List(o.Filter)
		return
	}

	if o.Type != "" {
		err = l.DeleteType(o.Filter, o.Type)
	} else {
		err = l.Delete(o.Filter)
	}
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	if !o.DryRun {
		log.Println(fmt.Sprintf("Try %s to list remaining resources!", fmt.Sprintf(color.BlueString("leftovers --filter %s --dry-run"), o.Filter)))
	}
}

func useOtherEnvVars(o opts, iaas string) opts {
	switch iaas {
	case AWS:
		if o.AWSAccessKeyID == "" {
			o.AWSAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
		}
		if o.AWSSecretAccessKey == "" {
			o.AWSSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
		}
		if o.AWSRegion == "" {
			o.AWSRegion = os.Getenv("AWS_DEFAULT_REGION")
		}
	case Azure:
		if o.AzureClientID == "" {
			o.AzureClientID = os.Getenv("ARM_CLIENT_ID")
		}
		if o.AzureClientSecret == "" {
			o.AzureClientSecret = os.Getenv("ARM_CLIENT_SECRET")
		}
		if o.AzureSubscriptionID == "" {
			o.AzureSubscriptionID = os.Getenv("ARM_SUBSCRIPTION_ID")
		}
		if o.AzureTenantID == "" {
			o.AzureTenantID = os.Getenv("ARM_TENANT_ID")
		}
	case GCP:
		if o.GCPServiceAccountKey == "" {
			o.GCPServiceAccountKey = os.Getenv("GOOGLE_CREDENTIALS")
		}
	case NSXT:
		if o.NSXTManagerHost == "" {
			o.NSXTManagerHost = os.Getenv("NSXT_MANAGER_HOST")
		}
		if o.NSXTUser == "" {
			o.NSXTUser = os.Getenv("NSXT_USERNAME")
		}
		if o.NSXTPassword == "" {
			o.NSXTPassword = os.Getenv("NSXT_PASSWORD")
		}
	case VSphere:
		if o.VSphereUser == "" {
			o.VSphereUser = os.Getenv("VSPHERE_USER")
		}
		if o.VSpherePassword == "" {
			o.VSpherePassword = os.Getenv("VSPHERE_PASSWORD")
		}
		if o.VSphereDC == "" {
			o.VSphereDC = os.Getenv("VSPHERE_DATACENTER")
		}
		if o.VSphereIP == "" {
			o.VSphereIP = os.Getenv("VSPHERE_IP")
		}
	}

	return o
}
