package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	awselb "github.com/aws/aws-sdk-go/service/elb"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/genevievelesperance/leftovers/app"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/elb"
	"github.com/genevievelesperance/leftovers/aws/iam"
	"github.com/genevievelesperance/leftovers/aws/s3"
	flags "github.com/jessevdk/go-flags"
)

type opts struct {
	NoConfirm bool `short:"n"  long:"no-confirm"  description:"Destroy resources without prompting. THIS DANGEROUS, MAKE GOOD CHOICES!"`

	AWSAccessKeyID     string `long:"aws-access-key-id"     env:"AWS_ACCESS_KEY_ID"     description:"AWS access key id."`
	AWSSecretAccessKey string `long:"aws-secret-access-key" env:"AWS_SECRET_ACCESS_KEY" description:"AWS secret access key."`
	AWSRegion          string `long:"aws-region"            env:"AWS_REGION"            description:"AWS region."`
}

type resource interface {
	Delete() error
}

func main() {
	log.SetFlags(0)

	var c opts
	parser := flags.NewParser(&c, flags.HelpFlag|flags.PrintErrors)
	_, err := parser.ParseArgs(os.Args)
	if err != nil {
		os.Exit(0)
	}

	logger := app.NewLogger(os.Stdout, os.Stdin, c.NoConfirm)

	if c.AWSAccessKeyID == "" {
		log.Fatal("Missing AWS_ACCESS_KEY_ID.")
	}

	if c.AWSSecretAccessKey == "" {
		log.Fatal("Missing AWS_SECRET_ACCESS_KEY.")
	}

	if c.AWSRegion == "" {
		log.Fatal("Missing AWS_REGION.")
	}

	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(c.AWSAccessKeyID, c.AWSSecretAccessKey, ""),
		Region:      aws.String(c.AWSRegion),
	}
	sess := session.New(config)

	iamClient := awsiam.New(sess)
	ec2Client := awsec2.New(sess)
	elbClient := awselb.New(sess)
	s3Client := awss3.New(sess)

	rolePolicies := iam.NewRolePolicies(iamClient, logger)
	userPolicies := iam.NewUserPolicies(iamClient, logger)
	accessKeys := iam.NewAccessKeys(iamClient, logger)

	ro := iam.NewRoles(iamClient, logger, rolePolicies)
	us := iam.NewUsers(iamClient, logger, userPolicies, accessKeys)
	ip := iam.NewInstanceProfiles(iamClient, logger)
	sc := iam.NewServerCertificates(iamClient, logger)

	ke := ec2.NewKeyPairs(ec2Client, logger)
	in := ec2.NewInstances(ec2Client, logger)
	se := ec2.NewSecurityGroups(ec2Client, logger)
	ta := ec2.NewTags(ec2Client, logger)
	vo := ec2.NewVolumes(ec2Client, logger)

	lo := elb.NewLoadBalancers(elbClient, logger)

	bu := s3.NewBuckets(s3Client, logger)

	resources := []resource{ip, ro, us, us, lo, sc, vo, ta, ke, in, se, bu}
	for _, r := range resources {
		if err = r.Delete(); err != nil {
			log.Fatalf("\n\n%s\n", err)
		}
	}
}
