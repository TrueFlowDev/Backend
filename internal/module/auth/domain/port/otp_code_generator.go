package port

type OtpCodeGenerator interface {
	Generate() string
}
