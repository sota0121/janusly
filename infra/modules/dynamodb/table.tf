resource "aws_dynamodb_table" "janusly_main" {
	name           = "janusly_main"
	billing_mode   = "PROVISIONED"
	hash_key       = "id"
	read_capacity = 1
	write_capacity = 1

	attribute {
		name = "id"
		type = "S"
	}
}
