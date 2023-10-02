resource "aws_s3_bucket" "insta-media" {
  bucket = "insta-clone-s3-bucket"

  tags = {
    Name = "insta-clone-s3-bucket"
  }
}

resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.insta-media.id

  policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Effect   = "Allow",
        Principal= "*",
        Action   = ["s3:GetObject", "s3:PutObject"],
        Resource =[ "${aws_s3_bucket.insta-media.arn}/*"]
      },
    ],
  })
}