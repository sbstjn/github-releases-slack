PROJECT_NAME ?= github-releases-slack
ENV ?= stable

AWS_BUCKET_NAME ?= $(PROJECT_NAME)-artifacts-$(ENV)
AWS_STACK_NAME ?= $(PROJECT_NAME)-$(ENV)
AWS_REGION ?= eu-west-1
GOOS ?= linux

FILE_TEMPLATE = infrastructure.yml
FILE_PACKAGE = dist/package.yml

clean:
	@ rm -rf dist

install:
	@ dep ensure

test:
	@ go test ./... -v

build: clean
	@ GOOS=$(GOOS) go build -o dist/handler ./...

build-osx: 
	@ GOOS=darwin make build

configure:
	@ aws s3api create-bucket \
		--bucket $(AWS_BUCKET_NAME) \
		--region $(AWS_REGION) \
		--create-bucket-configuration LocationConstraint=$(AWS_REGION)

package:
	@ aws cloudformation package \
		--template-file $(FILE_TEMPLATE) \
		--s3-bucket $(AWS_BUCKET_NAME) \
		--region $(AWS_REGION) \
		--output-template-file $(FILE_PACKAGE)

deploy:
	@ aws cloudformation deploy \
		--template-file $(FILE_PACKAGE) \
		--region $(AWS_REGION) \
		--capabilities CAPABILITY_IAM \
		--stack-name $(AWS_STACK_NAME)

describe:
	@ aws cloudformation describe-stacks \
		--region $(AWS_REGION) \
		--stack-name $(AWS_STACK_NAME)

outputs:
	@ make describe \
		| jq -r '.Stacks[0].Outputs'
