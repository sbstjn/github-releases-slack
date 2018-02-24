include .env

clean:
	rm -rf dist

test:
	go test

build: clean
	GOOS=linux go build -o dist/handler ./...

configure:
	@aws s3api create-bucket \
		--bucket $(AWS_BUCKET_NAME) \
		--region $(AWS_REGION)
		--create-bucket-configuration LocationConstraint=$(AWS_REGION) \

package: build
	@aws cloudformation package \
		--template-file infrastructure.yml \
		--s3-bucket $(AWS_BUCKET_NAME) \
		--region $(AWS_REGION) \
		--output-template-file dist/package.yml

deploy: package
	@aws cloudformation deploy \
		--template-file dist/package.yml \
		--region $(AWS_REGION) \
		--capabilities CAPABILITY_IAM \
		--stack-name $(AWS_STACK_NAME)

describe:
	@aws cloudformation describe-stacks \
		--region $(AWS_REGION) \
		--stack-name $(AWS_STACK_NAME)

outputs:
	@make describe | jq -r '.Stacks[0].Outputs'
